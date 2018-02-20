package builder

import (
	ocelog "bitbucket.org/level11consulting/go-til/log"
	"bitbucket.org/level11consulting/ocelot/util/repo/nexus"
	pb "bitbucket.org/level11consulting/ocelot/protos"
	"bitbucket.org/level11consulting/ocelot/util/cred"
	"bufio"
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"strings"
	"fmt"
)

type Docker struct{
	Log	io.ReadCloser
	ContainerId	string
	DockerClient *client.Client
	*Basher
}

func NewDockerBuilder(b *Basher) Builder {
	return &Docker{nil, "", nil, b}
}


func (d *Docker) Setup(logout chan []byte, werk *pb.WerkerTask, rc cred.CVRemoteConfig, werkerPort string) *Result {
	var setupMessages []string

	su := InitStageUtil("setup")

	logout <- []byte(su.GetStageLabel() + "Setting up...")

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	d.DockerClient = cli

	if err != nil {
		return &Result{
			Stage:  su.GetStage(),
			Status: FAIL,
			Error:  err,
		}
	}

	imageName := werk.BuildConf.Image

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		ocelog.IncludeErrField(err).Error("couldn't pull image!")
		return &Result{
			Stage:  su.GetStage(),
			Status: FAIL,
			Error:  err,
			Messages: setupMessages,
		}
	}
	setupMessages = append(setupMessages, fmt.Sprintf("pulled image %s \u2713", imageName))

	defer out.Close()

	bufReader := bufio.NewReader(out)
	d.writeToInfo(su.GetStageLabel(), bufReader, logout)

	logout <- []byte(su.GetStageLabel() + "Creating container...")

	//container configurations
	containerConfig := &container.Config{
		Image: imageName,
		Env: werk.BuildConf.Env,
		Cmd: d.DownloadTemplateFiles(werkerPort),
		AttachStderr: true,
		AttachStdout: true,
		AttachStdin:true,
		Tty:true,
	}

	//homeDirectory, _ := homedir.Expand("~/.ocelot")
	//host config binds are mount points
	hostConfig := &container.HostConfig{
		//TODO: have it be overridable via env variable
		Binds: []string{"/var/run/docker.sock:/var/run/docker.sock"},
		//Binds: []string{ homeDirectory + ":/.ocelot", "/var/run/docker.sock:/var/run/docker.sock"},
		NetworkMode: "host",
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig , hostConfig, nil, "")


	if err != nil {
		return &Result{
			Stage:  su.GetStage(),
			Status: FAIL,
			Error:  err,
			Messages: setupMessages,
		}
	}

	setupMessages = append(setupMessages, fmt.Sprint("created build container \u2713"))
	ocelog.Log().Info("JUST PRINTED build container checkmark MSG")

	for _, warning := range resp.Warnings {
		logout <- []byte(warning)
	}

	logout <- []byte(su.GetStageLabel() + "Container created with ID " + resp.ID)

	d.ContainerId = resp.ID
	ocelog.Log().Debug("starting up container")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return &Result{
			Stage:  su.GetStage(),
			Status: FAIL,
			Error:  err,
			Messages: setupMessages,
		}
	}

	logout <- []byte(su.GetStageLabel()  + "Container " + resp.ID + " started")
	ocelog.Log().Info("JUST PRINTED CONTAINER STARTED")


	//since container is created in setup, log tailing via container is also kicked off in setup
	containerLog, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow: true,
	})

	if err != nil {
		return &Result{
			Stage: su.GetStage(),
			Status: FAIL,
			Error:  err,
			Messages: setupMessages,
		}
	}

	d.Log = containerLog
	bufReader = bufio.NewReader(containerLog)

	d.writeToInfo(su.GetStageLabel() , bufReader, logout)

	ocelog.Log().Info("PLEASE SHOW UP IN THE LOGS")

	downloadCodebase := d.Exec(su.GetStage(), su.GetStageLabel(), []string{}, d.DownloadCodebase(werk), logout)
	if downloadCodebase.Error != nil {
		ocelog.Log().Error("an err happened trying to download codebase", downloadCodebase.Error)
		setupMessages = append(setupMessages, "failed to download codebase")
		downloadCodebase.Messages = append(downloadCodebase.Messages, setupMessages...)
		return downloadCodebase
	}


	logout <- []byte(su.GetStageLabel()  + "Retrieving SSH Key")
	ocelog.Log().Info(fmt.Println("just said that we're retrieving ssh key and fullname is %s", werk.FullName))
	setupMessages = append(setupMessages, fmt.Sprintf("downloading SSH key for %s...", werk.FullName))


	acctName := strings.Split(werk.FullName, "/")[0]
	result := d.Exec(su.GetStage(), su.GetStageLabel(), []string{}, d.DownloadSSHKey(
		werk.VaultToken,
		cred.BuildCredPath(werk.VcsType, acctName, cred.Vcs)), logout)
	if result.Error != nil {
		ocelog.Log().Error("an err happened trying to download ssh key", result.Error)
		result.Messages = append(result.Messages, setupMessages...)
		return result
	}

	setupMessages = append(setupMessages, fmt.Sprintf("successfully downloaded SSH key for %s  \u2713", werk.FullName))

	//only if the build tool is maven do we worry about settings.xml
	if werk.BuildConf.BuildTool == "maven" {
		if settingsXML, err := nexus.GetSettingsXml(rc, strings.Split(werk.FullName, "/")[0]); err != nil {
			_, ok := err.(*nexus.NoCreds)
			if !ok {
				return &Result{
					Stage: su.GetStage(),
					Status: FAIL,
					Error:  err,
				}
			}
		} else {
			ocelog.Log().Debug("writing maven settings.xml")
			result := d.Exec(su.GetStage(), su.GetStageLabel(), []string{}, d.WriteMavenSettingsXml(settingsXML), logout)
			result.Messages = append(result.Messages, setupMessages...)
			return result
		}
	}

	setupMessages = append(setupMessages, "completed setup stage \u2713")
	result.Messages = append(result.Messages, setupMessages...)
	return result
	//return &Result{
	//	Stage:  su.GetStage(),
	//	Status: PASS,
	//	Error:  err,
	//	Messages: setupMessages,
	//}
}

func (d *Docker) Cleanup(logout chan []byte) {
	su := InitStageUtil("cleanup")
	logout <- []byte(su.GetStageLabel() + "Performing build cleanup...")

	//TODO: review, should we be creating new contexts for every stage?
	cleanupCtx := context.Background()
	if d.Log != nil {
		d.Log.Close()
	}
	if err := d.DockerClient.ContainerKill(cleanupCtx, d.ContainerId, "SIGKILL"); err != nil {
		ocelog.IncludeErrField(err).WithField("containerId", d.ContainerId).Error("couldn't kill")
	} else {
		if err := d.DockerClient.ContainerRemove(cleanupCtx, d.ContainerId, types.ContainerRemoveOptions{}); err != nil {
			ocelog.IncludeErrField(err).WithField("containerId", d.ContainerId).Error("couldn't rm")
		}
	}
	d.DockerClient.Close()
}


func (d *Docker) Execute(stage *pb.Stage, logout chan []byte, commitHash string) *Result {
	if len(d.ContainerId) == 0 {
		return &Result {
			Stage: stage.Name,
			Status: FAIL,
			Error: errors.New("no container exists, setup before executing"),
		}
	}

	su := InitStageUtil(stage.Name)
	return d.Exec(su.GetStage(), su.GetStageLabel(), stage.Env, d.CDAndRunCmds(stage.Script, commitHash), logout)
}

func (d *Docker) Exec(currStage string, currStageStr string, env []string, cmds []string, logout chan []byte) *Result {
	var stageMessages []string
	ctx := context.Background()

	resp, err := d.DockerClient.ContainerExecCreate(ctx, d.ContainerId, types.ExecConfig{
		Tty: true,
		AttachStdin: true,
		AttachStderr: true,
		AttachStdout: true,
		Env: env,
		Cmd: cmds,
	})

	if err != nil {
		return &Result{
			Stage:  currStage,
			Status: FAIL,
			Error:  err,
			Messages: stageMessages,
		}
	}

	attachedExec, err := d.DockerClient.ContainerExecAttach(ctx, resp.ID, types.ExecConfig{
		Tty: true,
		AttachStdin: true,
		AttachStderr: true,
		AttachStdout: true,
		Env: env,
		Cmd: cmds,
	})

	defer attachedExec.Conn.Close()

	d.writeToInfo(currStageStr, attachedExec.Reader, logout)
	inspector, err := d.DockerClient.ContainerExecInspect(ctx, resp.ID)


	if inspector.ExitCode != 0 || err != nil {
		stageMessages = append(stageMessages, fmt.Sprintf("failed to complete %s stage \u2717", currStage))
		return &Result{
			Stage: currStage,
			Status: FAIL,
			Error: err,
			Messages: stageMessages,
		}
	}

	stageMessages = append(stageMessages, fmt.Sprintf("completed %s stage \u2713", currStage))
	return &Result{
		Stage:  currStage,
		Status: PASS,
		Error:  nil,
		Messages: stageMessages,
	}
}

func (d *Docker) writeToInfo(stage string, rd *bufio.Reader, infochan chan []byte) {
	scanner := bufio.NewScanner(rd)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		str := string(scanner.Bytes())
		infochan <- []byte(stage + str)
		//our setup script will echo this to stdout, telling us script is finished downloading. This is HACK for keeping container alive
		if strings.Contains(str, "Ocelot has finished with downloading templates") {
			ocelog.Log().Debug("finished with source code, returning out of writeToInfo")
			return
		}
	}
	ocelog.Log().Debug("finished writing to channel for stage ", stage)
	if err := scanner.Err(); err != nil {
		ocelog.IncludeErrField(err).Error("error outputing to info channel!")
		infochan <- []byte("OCELOT | BY THE WAY SOMETHING WENT WRONG SCANNING STAGE INPUT")
	}
}
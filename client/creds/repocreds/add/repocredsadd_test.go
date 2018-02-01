package repocredsadd

import (
	"bitbucket.org/level11consulting/ocelot/admin/models"
	"bitbucket.org/level11consulting/ocelot/client/commandhelper"
	"bytes"
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/cli"
	"testing"
)

// testNew will return the bare minimum. flags and fileloc of yaml will have to be set after instantiation
// or be generated by new functions
func testNew(inputReaderData []byte) *cmd {
	ui := cli.NewMockUi()
	if len(inputReaderData) >= 0 {
		ui.InputReader = bytes.NewReader(inputReaderData)
	}
	c := &cmd{
		UI: ui,
		config: commandhelper.NewTestClientConfig([]string{}),
	}
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(&c.fileloc, "credfile-loc", "",
		"Location of yaml file containing creds to upload")
	return c
}

func Test_cmd_Run_Yaml(t *testing.T) {
	var input []byte
	cmd := testNew(input)
	ctx := context.Background()
	cmd.fileloc = "./test-fixtures/newcreds.yml"
	expectedCreds := &models.RepoCredWrapper{
		Repo: []*models.RepoCreds{
			{
				Username:     "thisBeMyUserName",
				Password:     "SHH-BE-QUIET-ITS-A-SECRET",
				RepoUrl:      map[string]string{"snapshots":"https://ocelot.perf/nexus-yo"},
				AcctName:     "jessishank",
				Type:         "nexus",
			},
		},
	}
	var args []string
	if exit := cmd.Run(args); exit != 0 {
		t.Fatal("should return exit 0")
	}
	actualCreds, err := cmd.config.Client.GetRepoCreds(ctx, &empty.Empty{})
	if err != nil {
		t.Fatal("could not get actual creds from fake guide ocelot client")
	}
	if !models.CompareRepoCredWrappers(expectedCreds, actualCreds) {
		t.Error("expected creds mismatch\n expected: ", expectedCreds, "\n actual: ", actualCreds)
	}


}

func Test_cmd_Run_noYaml(t *testing.T) {
	input := []byte(`thisBeMyUserName
nexus
jessishank
https://ocelot.perf/nexus-yo
SHH-BE-QUIET-ITS-A-SECRET`)
	cmd := testNew(input)
	ctx := context.Background()
	expectedCreds := &models.RepoCredWrapper{
		Repo: []*models.RepoCreds{
			{
				Username:     "thisBeMyUserName",
				Password:     "SHH-BE-QUIET-ITS-A-SECRET",
				RepoUrl:      map[string]string{"snapshots":"https://ocelot.perf/nexus-yo"},
				AcctName:     "jessishank",
				Type:         "nexus",
			},
		},
	}

	var args []string
	exit := cmd.Run(args)
	if exit != 0 {
		t.Error("should return exit code 0, got ", exit)
	}
	sentCreds, err := cmd.config.Client.GetRepoCreds(ctx, &empty.Empty{})
	if err != nil {
		t.Fatal("could not get actual creds from fake guide ocelot client")
	}
	if !models.CompareRepoCredWrappers(expectedCreds, sentCreds) {
		t.Error("expected creds mismatch\n expected: ", expectedCreds, "\n actual: ", sentCreds)
	}
}

package cleaner

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shankj3/ocelot/build"
	"github.com/shankj3/ocelot/common/helpers/sshhelper"
	"github.com/shankj3/ocelot/models"
)

type SSHCleaner struct {
	*models.SSHFacts
}

func (d *SSHCleaner) Cleanup(ctx context.Context, id string, logout chan []byte) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ssh, err := sshhelper.CreateSSHChannel(ctx, d.SSHFacts, id)
	if err != nil {
		return err
	}
	if id == "" {
		return errors.New("id cannot be empty")
	}
	prefix := build.GetOcyPrefixFromWerkerType(models.SSH)
	cloneDir := build.GetCloneDir(prefix, id)
	if logout != nil {
		logout <- []byte("removing build directory " + cloneDir)
	}
	err = ssh.RunAndLog("rm -rf "+cloneDir, []string{}, logout, sshhelper.BasicPipeHandler)
	if err != nil {
		failedCleaning.WithLabelValues("ssh").Inc()
		if logout != nil {
			logout <- []byte("rould not remove build directory! Error: " + err.Error())
		}
		return err
	}
	if logout != nil {
		logout <- []byte("ruccessfully removed build directory.")
	}
	// if the context has been cancelled, then it was killed, as this deferred cleanup function is higher in the stack than the deferred cancel in (*launcher).makeitso
	if ctx.Err() == context.Canceled && logout != nil {
		logout <- []byte("//////////REDRUM////////REDRUM////////REDRUM/////////")
	}
	return nil
}

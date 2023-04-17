package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/DavisRayM/ops-gnome/pkg/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type taskRunner struct {
	StartTime time.Time
	task      config.Task
	workDir   string
}

func NewTaskRunner(task config.Task) *taskRunner {
	return &taskRunner{
		task: task,
	}
}

func (t *taskRunner) Start() error {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return fmt.Errorf("failed to create work directory: %w", err)
	}

	if len(t.task.Repo.URL) > 0 {
		cloneOptions := &git.CloneOptions{
			URL:          t.task.Repo.URL,
			SingleBranch: true,
		}

		if len(t.task.Repo.Branch) > 0 {
			cloneOptions.ReferenceName = plumbing.NewBranchReferenceName(t.task.Repo.Branch)
		}

		if len(t.task.Repo.DeployKey) > 0 {
			auth, err := ssh.NewPublicKeys("git", []byte(t.task.Repo.DeployKey), "")
			if err != nil {
				return fmt.Errorf("failed to load deploy key: %w", err)
			}
			cloneOptions.Auth = auth
		}

		_, err := git.PlainClone(tmpDir, false, cloneOptions)

		if err != nil {
			return err
		}
		t.workDir = tmpDir
	}

	t.StartTime = time.Now()
	return t.run()
}

func (t *taskRunner) run() error {
	var args []string
	if len(t.task.Command) > 1 {
		args = t.task.Command[1:]
	}

	cmd := exec.Command(t.task.Command[0], args...)
	cmd.Dir = t.workDir
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

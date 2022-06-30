package repo_manager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type RepoManager struct {
	repos        []string
	ignoreErrors bool
}

func NewRepoManager(rootDir string, repos []string, ignoreErrors bool) (repoManager *RepoManager, err error) {

	// Check if rootDir is valid
	_, err = os.Stat(rootDir)

	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New(fmt.Sprintf("root directory '%s' does not exist.", rootDir))
		}
	}

	// Set rootDir as an absolute path.
	rootDir, err = filepath.Abs(rootDir)
	if err != nil {
		return
	}

	if rootDir[len(rootDir)-1] != '/' {
		rootDir += "/"
	}

	// Check if repos is empty
	if len(repos) == 0 {
		err = errors.New("'repos', list cannot be empty.")
	}

	// Create RepoManager object
	repoManager = &RepoManager{
		ignoreErrors: ignoreErrors,
	}

	for _, r := range repos {
		path := rootDir + r
		repoManager.repos = append(repoManager.repos, path)
	}

	return

}

func (m *RepoManager) GetRepos() []string {
	return m.repos
}

func (m *RepoManager) Exec(cmd string) (output map[string]string, err error) {

	output = map[string]string{}

	var components []string

	for _, c := range strings.Split(cmd, " ") {
		components = append(components, c)
	}

	//Restore working directory after execution.
	pwd, _ := os.Getwd()
	defer os.Chdir(pwd)
	var out []byte
	for _, r := range m.repos {
		err = os.Chdir(r)
		if err != nil {
			return
		}

		// Execute command
		if cmd == "" {
			log.Fatal("command cannot be empty.")
		}
		out, err = exec.Command("git", components...).CombinedOutput()
		output[r] = string(out)

	}
	return
}

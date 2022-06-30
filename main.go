package main

import (
	"flag"
	"fmt"
	"log"
	"multi-git/pkg/repo_manager"
	"os"
	"strings"
)

func main() {

	command := flag.String("command", "", "The git command.")
	ignoreErrors := flag.Bool("ignore-errors", false, "Ignore all errors.")

	flag.Parse()

	root := os.Getenv("MG_ROOT")
	if root[len(root)-1] != '/' {
		root += "/"
	}

	repos := strings.Split(os.Getenv("MG_REPOS"), ",")

	repoManager, err := repo_manager.NewRepoManager(root, repos, *ignoreErrors)
	if err != nil {
		log.Fatal(err)
	}

	output, _ := repoManager.Exec(*command)
	for repo, out := range output {
		fmt.Printf("[%s] git %s\n", repo, *command)
		fmt.Println(out)
	}

	fmt.Println("DONE")

}

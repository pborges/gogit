package main

import (
	"os"
	"fmt"
	"net/url"
	"strings"
	"path"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", os.Args[0], "<git repo>")
		os.Exit(0)
	}
	repo := os.Args[1]
	repoURL, err := url.Parse(repo)
	if err != nil {
		fmt.Println(err.Error())
	}

	rawPath := strings.Split(repoURL.Path, "/")

	p := path.Join(os.Getenv("GOPATH"), "src", repoURL.Host, path.Join(rawPath...))
	err = os.MkdirAll(p, 0755)
	if err != nil {
		fmt.Println("mkdir:", err.Error())
	}
	err = os.Chdir(p)
	if err != nil {
		fmt.Println("chdir:", err.Error())
	}

	err = os.Chdir(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	cmd := exec.Command("git", "clone", repo, p)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	if err != nil {
		fmt.Println("cmd:", err.Error())
	}
}
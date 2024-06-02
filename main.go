package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var (
	pluginsPath string
)

func init() {
	pluginsPath = fmt.Sprintf("%s/.vim/pack/plugins/start", homeDir())
	if _, err := os.Stat(pluginsPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(pluginsPath, os.ModePerm)
		check(err)
	}
}

func main() {
	switch os.Args[1] {
	case "update", "u":
		updatePlugins()
	case "install", "i":
		installPlugin(os.Args[2])
	case "remove", "r":
		removePlugin(os.Args[2])
	}
}

func installPlugin(url string) {
	cmd := exec.Command("git", "clone", url)
	cmd.Dir = pluginsPath
	stdcmd(cmd)
	urlSplit := strings.Split(url, "/")
	fmt.Printf("[vplug] Installed: %s\n", urlSplit[len(urlSplit)-1])
}

func removePlugin(name string) {
	cmd := exec.Command("rm", "-rf", name)
	cmd.Dir = pluginsPath
	stdcmd(cmd)
	fmt.Printf("[vplug] Removed: %s\n", name)
}

func updatePlugins() {
	plugins, err := os.ReadDir(pluginsPath)
	check(err)
	for _, x := range plugins {
		cmd := exec.Command("git", "pull", "origin", "master")
		cmd.Dir = fmt.Sprintf("%s/%s", pluginsPath, x.Name())
		stdcmd(cmd)
	}
	fmt.Printf("[vplug] All plugins updated\n")
}

func homeDir() string {
	currentUser, err := user.Current()
	check(err)
	return fmt.Sprintf("/home/%s", currentUser.Username)
}

func stdcmd(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

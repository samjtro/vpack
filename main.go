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
	pluginsPath, colorsPath, syntaxPath string
)

func init() {
	pluginsPath = fmt.Sprintf("%s/.vim/pack/plugins/start", homeDir())
	colorsPath = fmt.Sprintf("%s/.vim/pack/colors/start", homeDir())
	syntaxPath = fmt.Sprintf("%s/.vim/pack/syntax/start", homeDir())
	if _, err := os.Stat(pluginsPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(pluginsPath, os.ModePerm)
		check(err)
	}
	if _, err := os.Stat(colorsPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(colorsPath, os.ModePerm)
		check(err)
	}
	if _, err := os.Stat(syntaxPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(syntaxPath, os.ModePerm)
		check(err)
	}
}

func main() {
	switch os.Args[1] {
	case "c":
		switch os.Args[2] {
		case "update", "u":
			update("c")
		case "install", "i":
			install(os.Args[3], "c")
		case "remove", "r":
			remove(os.Args[3], "c")
		}
	case "p":
		switch os.Args[2] {
		case "update", "u":
			update("p")
		case "install", "i":
			install(os.Args[3], "p")
		case "remove", "r":
			remove(os.Args[3], "p")
		}
	case "s":
		switch os.Args[2] {
		case "update", "u":
			update("s")
		case "install", "i":
			install(os.Args[3], "s")
		case "remove", "r":
			remove(os.Args[3], "s")
		}
	}
}

func install(url, arg string) {
	cmd := exec.Command("git", "clone", url)
	switch arg {
	case "c":
		cmd.Dir = colorsPath
	case "p":
		cmd.Dir = pluginsPath
	case "s":
		cmd.Dir = syntaxPath
	}
	stdcmd(cmd)
	urlSplit := strings.Split(url, "/")
	fmt.Printf("[vpack] Installed: %s\n", urlSplit[len(urlSplit)-1])
}

func remove(name, arg string) {
	cmd := exec.Command("rm", "-rf", name)
	switch arg {
	case "c":
		cmd.Dir = colorsPath
	case "p":
		cmd.Dir = pluginsPath
	case "s":
		cmd.Dir = syntaxPath
	}
	stdcmd(cmd)
	fmt.Printf("[vpack] Removed: %s\n", name)
}

func update(arg string) {
	switch arg {
	case "c":
		elements, err := os.ReadDir(colorsPath)
		check(err)
		for _, x := range elements {
			cmd := exec.Command("git", "pull", "origin", "master")
			cmd.Dir = fmt.Sprintf("%s/%s", colorsPath, x.Name())
			stdcmd(cmd)
		}
	case "p":
		elements, err := os.ReadDir(pluginsPath)
		check(err)
		for _, x := range elements {
			cmd := exec.Command("git", "pull", "origin", "master")
			cmd.Dir = fmt.Sprintf("%s/%s", pluginsPath, x.Name())
			stdcmd(cmd)
		}
	case "s":
		elements, err := os.ReadDir(syntaxPath)
		check(err)
		for _, x := range elements {
			cmd := exec.Command("git", "pull", "origin", "master")
			cmd.Dir = fmt.Sprintf("%s/%s", syntaxPath, x.Name())
			stdcmd(cmd)
		}
	}
	fmt.Printf("[vpack] All plugins updated\n")
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

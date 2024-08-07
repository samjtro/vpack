package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	if len(os.Args) == 1 {
		fmt.Println(`welcome to vpack. usage:
vpack c i [https://repository.url] - install color scheme to ~/.vim/pack/colors/start/[name_of_repository]
vpack c r [name_of_repository] - remove color scheme
vpack c u - update all color schemes

vpack p i [https://repository.url] - install plugin to ~/.vim/pack/plugin/start/[name_of_repository]
vpack p r [name_of_repository] - remove plugin
vpack p u - update all plugins

vpack s i [https://repository.url] - install syntax to ~/.vim/pack/syntax/start/[name_of_repository]
vpack s r [name_of_repository] - remove syntax
vpack s u - update all syntaxes`)
	} else {
		switch os.Args[1] {
		case "color", "c":
			switch os.Args[2] {
			case "update", "u":
				update("c")
			case "install", "i":
				install(os.Args[3], "c")
			case "remove", "r":
				remove(os.Args[3], "c")
			}
		case "plugin", "p":
			switch os.Args[2] {
			case "update", "u":
				update("p")
			case "install", "i":
				install(os.Args[3], "p")
			case "remove", "r":
				remove(os.Args[3], "p")
			}
		case "syntax", "s":
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
	var path string
	switch arg {
	case "c":
		path = colorsPath
	case "p":
		path = pluginsPath
	case "s":
		path = syntaxPath
	}
	elements, err := os.ReadDir(path)
	check(err)
	for _, x := range elements {
		cmd := exec.Command("git", "pull", "origin", "master")
		cmd.Dir = fmt.Sprintf("%s/%s", path, x.Name())
		stdcmd(cmd)
	}
	fmt.Printf("[vpack] All plugins updated\n")
}

func homeDir() string {
	homedir, err := os.UserHomeDir()
	check(err)
	return homedir
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

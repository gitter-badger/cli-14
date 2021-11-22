package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func generateController(name, namespace string) string {
	return fmt.Sprintf(`<?php

namespace %s;

use JtF\Http\Response;

class %s
{
    public function index(Response $res)
    {
        $res->ok();
    }
}
`, namespace, name)
}

func runServer(port int, folder string) {
	command := exec.Command("php", "-S", "localhost:"+strconv.Itoa(port), "-t", folder)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = os.Environ()
	command.Env = append(command.Env, "JT_ENVIRONMENT=dev")
	err := command.Run()
	if err != nil {
		fmt.Println("\u274C Opps you have a problem...")
		os.Exit(1)
	}
}

func printLogo() {
	figure.NewColorFigure("JT CLI", "speed", "cyan", true).Print()
	fmt.Println("\n\033[92m Fast and powerfull cli for JT Framework\033[0m")
}

func enableUserSecrets() {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
	}
	path := filepath.Join(dir, ".jt")
	os.MkdirAll(path, os.ModePerm)
	id := randomString(8)
	file, err := os.Create(filepath.Join(path, id+".json"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Your secrets ID: %s\n", id)
	defer file.Close()
}

func createController(name, path, ns string) {
	filePath := filepath.Join(path, name+".php")
	os.MkdirAll(path, os.ModePerm)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("\u274C Unable to create file...")
		os.Exit(1)
	} else {
		_, err := file.WriteString(generateController(name, ns))
		if err != nil {
			fmt.Println("\u274C Oops unknown error when writing to file...")
		} else {
			fmt.Printf("\u2705 Controller %s created\n", name)
		}
	}

	defer file.Close()
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

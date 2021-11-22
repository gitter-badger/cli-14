package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

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

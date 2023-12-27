package genie1

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func genie() {
	projectName := flag.String("name", "new_gen", "the name of the new project")
	path := flag.String("path", "./path_gen", "the name of the new project")
	flag.Parse()

	if *path != "./path_gen" {
		os.Mkdir(*path, 0775)
		err := os.Chdir(*path)
		if err != nil {
			os.Exit(1)
		}
	}
	os.Mkdir(*projectName, 0755)
	os.Chdir(*projectName)

	cmd := exec.Command("mkdir", "-p", "cmd", "build", "assets", "tools", "internal/models", "ui", "pkg")
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	// create Readme

}

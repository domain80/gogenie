package main

import (
	"embed"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:embed assets
var assets embed.FS

func main() {

	projectName := flag.String("name", "new_gen", "the name of the new project")
	path := flag.String("path", "./", "the name of the new project")
	flag.Parse()

	err := os.MkdirAll(*path+"/"+*projectName, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}

	err = os.Chdir(*path + "/" + *projectName)
	if err != nil {
		log.Fatal(err)
	}

	// ? create directories
	directories := []string{
		"-p", "cmd ", "internal ", "package ", "web ", "config ", "build ", "deployments ", "test ", "docs ", "assets ", "website",
	}
	cmd := exec.Command("mkdir", directories...)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// ? create root files
	dir, err := assets.ReadDir("assets/rootFiles")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range dir {
		fc, err := assets.ReadFile("assets/rootFiles/" + f.Name())
		if err != nil {
			log.Println(err)
		}
		os.WriteFile(strings.ReplaceAll(f.Name(), "tmpl_", ""), fc, 0755)
	}

}

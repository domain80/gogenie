/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Generate a new project using the default or provided name and path flags",
	Long: `Generate a new project using the default or provided name and path flags 

	- The default path is the current working directory
	- The default name is genieProject_gen

	- A new git repository is initialized in the new project folder
	- A new go module is also created inside the new generated project 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		nameFlag := cmd.Flag("name").Value.String()
		pathFlag := cmd.Flag("path").Value.String()

		err := os.Chdir(path.Join(pathFlag, nameFlag))
		if err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(path.Join(pathFlag, nameFlag), 0755)
				os.Chdir(path.Join(pathFlag, nameFlag))
			} else {
				log.Fatal(err)
			}
		}

		dir, err := assets.ReadDir("assets/project-layout")
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range dir {
			if d.IsDir() {
				err := os.Mkdir(d.Name(), 0755)
				if err != nil {
					log.Println(err)
					continue
				}
			} else {
				fileContent, err := assets.ReadFile(path.Join("assets", "newProject", d.Name()))
				if err != nil {
					log.Println(err)
					continue
				}
				os.WriteFile(
					strings.ReplaceAll(d.Name(), "template_", ""),
					fileContent, 0755,
				)
			}
		}

		// ? initialize git repo and go module

		log.Println("Initializing git repository...")
		out, err := exec.Command("git", "init").Output()
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(out))
		}

		println()
		log.Println("Initializing go module...")
		out, err = exec.Command("go", "mod", "init", nameFlag).Output()
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(out))
		}
		log.Println("Go module intialized at " + path.Join(pathFlag, nameFlag))

	},
}

func init() {
	rootCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//go:assets assets
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "generate a new model using a given struct",
	Long: `Flow:
	user defines model type
	user specifies /*go:generate genie model <model_name> [[db-type="pg/msql/mongo"]] */  comment above model Gen
	When the generate command is run:
	A new file in the internal/models directory is crated with the name <model_name>.go
	A new model of name <model_name>Model
	CRUD functions around the Model.
	  Create 
	  Read
	    Latest
	    All
	    Paginated
	  Update
	  Delete `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
		nameFlag := cmd.Flag("name")
		// pathFlag := cmd.Flag("path")

		modelTemplate, err := assets.ReadFile("assets/model.tmpl")
		if err != nil {
			log.Fatal(err)
		}

		newModel, err := os.Create(nameFlag.Value.String() + ".go")
		if err != nil {
			log.Fatal(err)
		}
		defer newModel.Close()

		tp, err := template.New("modelTemplate").Parse(string(modelTemplate))
		if err != nil {
			log.Fatal(err)
		}

		err = tp.Execute(newModel, struct {
			Name    string
			DBType  string
			Package string
		}{Name: nameFlag.Value.String(),
			DBType:  "sql",
			Package: os.Getenv("GOPACKAGE"),
		})
		if err != nil {
			log.Println(err)
		}

		// bf := bytes.Buffer{}
		a, _ := os.ReadFile(newModel.Name())
		fmt.Println(string(a))
		a, _ = format.Source(a)
		fmt.Println(string(a))
		// a.Read(bf.Bytes())
		// fmt.Println(bf.String())
		// a.WriteTo(newModel)

	},
}

func init() {
	rootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

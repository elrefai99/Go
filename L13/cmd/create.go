package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CreateCommands() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create commands",
	}

	newCmd := &cobra.Command{
		Use:   "new <config-file>",
		Short: "Create new file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var fileType string
			fmt.Println("What file type u need (yml, sql, json)")
			fmt.Scanln(&fileType)
			path := "data/" + args[0] + "." + fileType
			var data any

			if fileType == "sql" {
				jsonData, err := json.Marshal(data)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to encode json data: %v\n", err)
					return
				}

				if err := os.WriteFile(path, jsonData, 0644); err != nil {
					fmt.Fprintf(os.Stderr, "failed to write file %q: %v\n", path, err)
					return
				}
			}
		},
	}

	createCmd.AddCommand(newCmd)

	return createCmd
}

func init() {
	rootCmd.AddCommand(CreateCommands())
}

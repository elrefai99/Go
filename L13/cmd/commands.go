package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var port int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test command", port)
	},
}

func getCommands() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get <config-file>",
		Short: "get value from json file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := "data/" + args[0] + ".json"
			data, err := os.ReadFile(path)
			if err != nil {
				log.Fatal(err)
				return
			}

			var body any
			err = json.Unmarshal(data, &body)
			if err != nil {
				log.Fatal(err)
				return
			}

			pretty, _ := json.MarshalIndent(body, "", "  ")
			fmt.Println(string(pretty))
		},
	}

	return getCmd
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "port number")
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(getCommands())
}

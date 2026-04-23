package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

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
			t := time.Now().Unix() * 1000
			path := "data/" + "migrate-" + time.Now().String() + "-" + strconv.FormatInt(t, 10) + ".sql"
			var data any

			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to encode json data: %v\n", err)
				return
			}

			if err := os.WriteFile(path, jsonData, 0644); err != nil {
				fmt.Fprintf(os.Stderr, "failed to write file %q: %v\n", path, err)
				return
			}
		},
	}

	createCmd.AddCommand(newCmd)

	return createCmd
}

func init() {
	rootCmd.AddCommand(CreateCommands())
}

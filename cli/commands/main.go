package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zrl",
	Short: "CLI tool to shorten URLs.",
	Run: func(cmd *cobra.Command, args []string) {
		res, _ := http.Get("http://localhost:3000")

		var data = ""
		json.NewDecoder(res.Body).Decode(&data)
		fmt.Println(data)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing zrl '%s'\n", err)
		os.Exit(1)
	}
}

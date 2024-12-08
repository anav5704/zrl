package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete short URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteUrl(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteUrl(short string) {
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3000/api/url/" + short, nil)
    res, _ := http.DefaultClient.Do(req)

    var data string
    json.NewDecoder(res.Body).Decode(&data)
    fmt.Println(data)
}

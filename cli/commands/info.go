package commands

import (
	"encoding/json"
	"net/http"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)   

var infoCmd = &cobra.Command{
    Use: "info",
    Short: "Analytics for short URL",
	Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        infoUrl(args[0])
    },
}

func init() {
    rootCmd.AddCommand(infoCmd)
}

func infoUrl(short string) {
    req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/url/" + short, nil)
    res, _ := http.DefaultClient.Do(req)

    var data Url
    json.NewDecoder(res.Body).Decode(&data)

    tbl := table.New("Short", "Long", "Views")
    tbl.AddRow(data.Short, data.Long, data.Views)
    tbl.Print()
}

package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)   

var listCmd = &cobra.Command{
    Use: "ls",
    Short: "List short URLs",
    Run: func(cmd *cobra.Command, args []string) {
        listUrl()   
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}

type Url struct {   
    Short string `json:"short"` 
    Long string `json:"long"`
    Views int `json:"views"` 
}

func listUrl() {
    res, err := http.Get("http://localhost:3000/api/url/")

    if err != nil {
        fmt.Println("Error listing URLs: ", err)
        return
    }

    var data []Url
    json.NewDecoder(res.Body).Decode(&data)

    tbl := table.New("Short", "Long", "Views")

    for _, url := range data {
        tbl.AddRow(url.Short, url.Long, url.Views)
    }

    tbl.Print()
}

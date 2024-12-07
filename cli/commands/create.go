package commands

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"

    "github.com/spf13/cobra"
)   

var createCmd = &cobra.Command{
    Use: "create",
    Short: "Creatae short URL",
    Args: cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        createUrl(args[0], args[1])
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
}

func createUrl(short, long string) {
    res, err := http.PostForm("http://localhost:3000/api/url/", url.Values{
    "short": {short},
    "long": {long},
})

if err != nil {
    fmt.Println("Error creating URL: ", err)
    return
}

var data = ""
json.NewDecoder(res.Body).Decode(&data)
fmt.Println(data)

}

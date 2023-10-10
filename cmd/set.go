package cmd

import (
	"confluence-poc/src/config"
	"confluence-poc/src/models"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// getCmdGroup is a group for everything todo with get
var setCmdGroup = &cobra.Command{
	Use:   "set",
	Short: "Sets one or many resources",
	Long:  `Sets information about one or more resources`,
}

var setPagesCmd = &cobra.Command{
	Use:   "page",
	Short: "Sets one or many pages",
	Long:  `Sets information about one or more pages`,
	Run: func(cmd *cobra.Command, args []string) {
		setPages(args)
	},
}

func init() {
	rootCmd.AddCommand(setCmdGroup)
	setCmdGroup.AddCommand(setPagesCmd)
}

func setPages(content []string) {
	confluenceClient, err := config.NewConfluenceClient(flags)
	if err != nil {
		fmt.Printf("failed to create confluence client: %s", err)
	}

	contentList := []models.Content{}
	for _, c := range content {
		var contentMap models.Content
		err := json.Unmarshal([]byte(c), &contentMap)
		if err != nil {
			fmt.Printf("failed to marshal content: %s", err)
		}
		contentList = append(contentList, contentMap)
	}

	result := confluenceClient.SetContent(contentList)
	j, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("failed to marshal result: %s", err)
	}
	fmt.Println(string(j))
}

package cmd

import (
	"confluence-poc/src/config"
	"confluence-poc/src/models"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var version int

// getCmdGroup is a group for everything todo with get
var getCmdGroup = &cobra.Command{
	Use:   "get",
	Short: "Gets one or many resources",
	Long:  `Gets information about one or more resources`,
}

var getPagesCmd = &cobra.Command{
	Use:   "page",
	Short: "Gets one or many pages",
	Long:  `Gets information about one or more pages`,
	Run: func(cmd *cobra.Command, args []string) {
		getPages(args)
	},
}

func init() {
	rootCmd.AddCommand(getCmdGroup)
	getCmdGroup.AddCommand(getPagesCmd)

	getPagesCmd.Flags().IntVarP(&version, "version", "v", 0, "version")
}

func getPages(ids []string) {
	confluenceClient, err := config.NewConfluenceClient(flags)
	if err != nil {
		fmt.Printf("failed to create confluence client: %s", err)
	}

	query := models.ContentQuery{
		Expand:  []string{"body.storage", "version"},
		Version: version,
	}

	result := confluenceClient.GetContentByID(ids, query)
	j, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("failed to marshal result: %s", err)
	}
	fmt.Println(string(j))
}

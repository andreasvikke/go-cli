package cmd

import (
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

	// getPagesCmd.Flags().IntVarP(&version, "version", "v", 0, "version")
}

func setPages(ids []string) {
	// confluenceClient, err := config.NewConfluenceClient(flags)
	// if err != nil {
	// 	fmt.Printf("failed to create confluence client: %s", err)
	// }

	// query := models.ContentQuery{
	// 	Expand:  []string{"body.storage", "version"},
	// 	Version: version,
	// }

	// result := confluenceClient.GetContentByID(ids, query)
	// j, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Printf("failed to marshal result: %s", err)
	// }
	// fmt.Println(string(j))
}

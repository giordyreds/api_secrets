package cobra

import (
	"example/apisecrets"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := apisecrets.File(encodingKey, secretsPath())
		keysList, err := v.List()
		if err != nil {
			fmt.Println("Missing encoding key.")
			return
		}
		for _, k := range keysList {
			fmt.Printf("%s\n", k)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

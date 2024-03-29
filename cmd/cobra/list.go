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
		encodingKey, err := getEncodingKey()
		if err != nil {
			fmt.Println("error getting key from prompt: ", err)
			return
		}

		v := apisecrets.File(encodingKey, secretsPath())

		keysList, err := v.List()
		if err != nil {
			fmt.Println("Missing encoding key.")
			return
		}

		fmt.Println("\nkeys list:")
		for _, k := range keysList {
			fmt.Printf("%s\n", k)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

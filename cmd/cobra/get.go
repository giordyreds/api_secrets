package cobra

import (
	"example/apisecrets"
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		encodingKey, err := getEncodingKey()
		if err != nil {
			fmt.Println("error getting key from prompt: ", err)
			return
		}

		v := apisecrets.File(encodingKey, secretsPath())

		key, err := readVisibleKey()
		if err != nil {
			fmt.Println("error reading key: ", err)
			return
		}

		value, err := v.Get(key)
		if err != nil {
			fmt.Println("no value set for key: ", key)
			return
		}

		fmt.Printf("%s: %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}

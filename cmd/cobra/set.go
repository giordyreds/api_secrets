package cobra

import (
	"example/apisecrets"
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
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

		value, err := readHiddenValue()
		if err != nil {
			fmt.Println("error reading value: ", err)
			return
		}

		err = v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Value set.")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}

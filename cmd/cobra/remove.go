package cobra

import (
	"example/apisecrets"
	"fmt"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a key in your secret storage",
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

		err = v.Remove(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s has been removed from secrets.\n", key)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}

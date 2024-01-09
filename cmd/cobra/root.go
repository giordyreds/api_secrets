package cobra

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}

func getEncodingKey() (string, error) {
	fmt.Printf("insert encoding key: ")
	result, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", errors.New("missing encoding key")
	}

	return string(result), nil
}

func readVisibleKey() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("insert secret name: ")
	if scanner.Scan() {
		input := scanner.Text()
		return input, nil
	} else {
		err := scanner.Err()
		return "", err
	}
}

func readHiddenValue() (string, error) {
	fmt.Printf("insert value to hide: ")
	result, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}

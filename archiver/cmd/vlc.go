package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Play file using variable-length code",
	Run:   pack,
}
var ErrEmptyFile = errors.New("path to file is specified")

const packedExtension = "vlc"

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyFile)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	// data -> Encode(data)
	packed := ""
	fmt.Println(string(data)) // TODO: remove

	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(vlcCmd)
}

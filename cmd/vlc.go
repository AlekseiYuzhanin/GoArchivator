package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd  = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: pack,
}

const packedExtension = "vlc"

func pack(_ *cobra.Command, args []string){
	filePath := args[0]

	r,err := os.Open(filePath)
	if err!=nil{
		handleError(err)
	}

	data,err:= ioutil.ReadAll(r)
	if err!= nil{
		handleError(err)
	}

	packed := "" + string(data)

	err = ioutil.WriteFile(packedFileName(filePath),[]byte(packed),0644)
	if err!= nil{
		handleError(err)
	}
}

func packedFileName(path string) string{
	fileName := filepath.Base(path)
	

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init(){
	packCmd.AddCommand(vlcCmd)
}
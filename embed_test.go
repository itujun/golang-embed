package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// ? Embed harus berada diluar function

//go:embed version.txt
var version string

//go:embed version/version.txt
var version2 string

func TestEmbedFileToString(t *testing.T){
	fmt.Println(version);
	fmt.Println(version2);
}

//go:embed favicon.ico 
var favicon []byte

//go:embed logo/logo.png 
var logo []byte

func TestEmbedFileToByteArray(t *testing.T){
	err := os.WriteFile("favicon_new.ico", favicon, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("logo/logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

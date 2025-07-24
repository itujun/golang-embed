package test

import (
	"embed"
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

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultipleFile(t *testing.T){
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a));
	
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b));

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c));
}

//go:embed files/*.txt
var path embed.FS
func TestEmbedPathMatcher(t *testing.T){
	dirEntries, _ := path.ReadDir("files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
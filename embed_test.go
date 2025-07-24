package golang_embed

import (
	_ "embed"
	"fmt"
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
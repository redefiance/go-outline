package main

import (
	"flag"
	"fmt"
	"go/token"
	"os"
	"path"

	"github.com/redefiance/go-outline/outline"
)

var fPath = flag.String("path", "testpkg", "TODO desc")
var fPublic = flag.Bool("public", true, "show only package exports")

var fs = token.NewFileSet()

func main() {
	flag.Parse()

	var doFolder func(string)
	doFolder = func(dirpath string) {
		f, err := os.Open(dirpath)
		if err != nil {
			panic(err)
		}
		fis, err := f.Readdir(0)
		if err != nil {
			panic(err)
		}
		for _, fi := range fis {
			if fi.IsDir() {
				doFolder(path.Join(dirpath, fi.Name()))
			}
		}

		pkg, err := outline.ParsePackage(dirpath, *fPublic)
		if err != nil {
			// fmt.Println(err)
		} else {
			fmt.Print(pkg)
		}
	}
	doFolder(*fPath)
}

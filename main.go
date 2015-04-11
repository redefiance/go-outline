package main

import (
	"flag"
	"fmt"
	"go/token"
	"os"
	"path"

	"github.com/redefiance/go-outline/outline"
)

var fPath = flag.String("path", "/home/stargazer/dev/go/go-outline/testpkg", "TODO desc")
var fPublic = flag.Bool("public", true, "show only package exports")

var fs = token.NewFileSet()

func main() {
	flag.Parse()

	var doFolder func(string)
	doFolder = func(dirpath string) {
		pkg, err := outline.ParsePackage(dirpath, *fPublic)
		if err != nil {
			// fmt.Println(err)
		} else {
			fmt.Println("pkg", pkg.Path)
			for _, file := range pkg.Files {
				fmt.Println("file", file.Path)
				for _, decl := range file.Decls {
					fmt.Printf("%s:%d:%d\n", decl, decl.LineFrom, decl.LineTo)
				}
			}
		}

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
	}
	doFolder(*fPath)
}

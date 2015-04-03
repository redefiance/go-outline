package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"

	"github.com/redefiance/go-outline/outline"
)

var fPath = flag.String("path", "/home/dev/go/go-outline/testpkg", "TODO desc")
var fPublic = flag.Bool("public", true, "show only package exports")

var fs = token.NewFileSet()

func main() {
	flag.Parse()

	var doFolder func(string)
	doFolder = func(dirpath string) {
		pkgs, err := parser.ParseDir(fs, dirpath, nil, 0)
		if err != nil {
			panic(err)
		}

		for pkgname, pkg := range pkgs {
			if strings.HasSuffix(pkgname, "_test") {
				continue
			}

			fmt.Println("pkg", strings.TrimLeft(strings.TrimLeft(dirpath, *fPath), "/"))
			for _, file := range outline.ParsePackage(pkg, *fPublic).Files {
				fmt.Println("file", file.Path)
				for _, decl := range file.Decls {
					fmt.Println(decl)
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

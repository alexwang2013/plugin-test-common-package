package main

import "github.com/alexwang2013/plugin-test-common-package/src/printer"

func main() {
	str := "Hello world"
	p := printer.Printer{}
	p.Str = str
	p.Print()
}

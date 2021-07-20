package main

import printer "github.com/alexwang2013/plugin-test-common-package/src"

func main() {
	str := "Hello world"
	p := printer.Printer{}
	p.Str = str
	p.Print()
}

package printer

import "fmt"

type Printer struct {
	var Str string
}

func(p *Printer)print(){
	fmt.Println(p.Str)
}
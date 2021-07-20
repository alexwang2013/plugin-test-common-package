package printer

import "fmt"

type Printer struct {
	Str string
}

func (p *Printer) Print() {
	fmt.Println(p.Str)
}

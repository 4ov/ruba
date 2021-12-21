package inter

import "fmt"

func (r *Ruba) debug(i interface{}) {
	if r.Options.Debug {
		fmt.Println(i)
	}
}

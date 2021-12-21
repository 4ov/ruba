package exec

import (
	"fmt"
	"os"
)

func NewError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

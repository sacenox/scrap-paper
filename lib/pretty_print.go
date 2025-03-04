package lib

import (
	"fmt"

	"github.com/kr/pretty"
)

func PrettyPrint(v interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))
}

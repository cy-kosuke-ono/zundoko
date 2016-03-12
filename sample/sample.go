package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/zundoko"
)

func main() {
	c := zundoko.Run()
	fmt.Println(c)
	z, d := c.CountZunDoko()
	fmt.Printf(" count \"ZUN\": %d\n", z)
	fmt.Printf("count \"DOKO\": %d\n", d)
	fmt.Printf(" count Total: %d\n", z+d)
}

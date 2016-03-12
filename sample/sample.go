package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/zundoko"
)

func main() {
	c := zundoko.Run()
	fmt.Println(c)
	fmt.Printf("%d回\n", c.Count()-1) //exclude "KIYOSHI" choir
}

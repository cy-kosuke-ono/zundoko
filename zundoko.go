package zundoko

import (
	"math/rand"
	"strings"
	"time"
)

// zundoko-choirs items
const (
	Zun     = "ズン"
	Doko    = "ドコ"
	Kiyoshi = "キ・ヨ・シ！"
)

//Choirs - zundoko-choirs list
type Choirs struct {
	c  []string
	ct int
}

//Push is append choirs
func (c *Choirs) Push(s string) {
	c.c = append(c.c, s)
	c.ct++
}

//Count returns choirs count
func (c *Choirs) Count() int {
	return c.ct
}

func (c *Choirs) String() string {
	return strings.Join(c.c, "")
}

func generate() chan string {
	ch := make(chan string)
	go func() {
		var zd = [2]string{Zun, Doko}
		rand.Seed(time.Now().UnixNano())
		for {
			ch <- zd[rand.Intn(2)]
		}
	}()
	return ch
}

//Run zundoko-choirs
func Run() *Choirs {
	zd := generate()
	var c = &Choirs{make([]string, 0), 0}
	zcount := 0
	for {
		s := <-zd
		c.Push(s)
		if s == Zun {
			zcount++
		} else if zcount >= 4 {
			break
		} else {
			zcount = 0
		}
	}
	c.Push(Kiyoshi)
	return c
}

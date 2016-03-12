package zundoko

import (
	"math/rand"
	"reflect"
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
	c []string
}

//Push is append choirs
func (c *Choirs) Push(s string) {
	c.c = append(c.c, s)
}

var matchingPattern = []string{Zun, Zun, Zun, Zun, Doko}

func (c *Choirs) match() bool {
	if c == nil {
		return false
	}
	if len(c.c) < 5 {
		return false
	}
	return reflect.DeepEqual(c.c[len(c.c)-5:], matchingPattern)
}

//CountZunDoko returns count of "ZUN" and "DOKO" choirs
func (c *Choirs) CountZunDoko() (int, int) {
	z := 0
	d := 0
	if c == nil {
		return z, d
	}
	for _, s := range c.c {
		switch s {
		case Zun:
			z++
		case Doko:
			d++
		}
	}
	return z, d
}

func (c *Choirs) String() string {
	if c == nil {
		return ""
	}
	content := make([]byte, 0, 128)
	for _, s := range c.c {
		content = append(content, s...)
	}
	return string(content)
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
	c := &Choirs{make([]string, 0)}
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

//Run2 zundoko-choirs (another logic)
func Run2() *Choirs {
	zd := generate()
	c := &Choirs{make([]string, 0)}
	for {
		s := <-zd
		c.Push(s)
		if c.match() {
			break
		}
	}
	c.Push(Kiyoshi)
	return c
}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	s := t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	fmt.Printf("%T, %v \n", s, t.Format("15:04 Mon Dec 2 2006"))
	pt, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", s)
	fmt.Println("error", err)
	fmt.Printf("parsed time: %v, %T \n", pt, pt)
	tf := pt.Format("Mon Dec 2 2006")
	fmt.Println(tf)
}

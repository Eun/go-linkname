package main

import (
	"fmt"
	"time"
	_ "unsafe"
)

//go:linkname isLeap time.isLeap
func isLeap(year int) bool

//go:linkname abs time.Time.abs
func abs(t time.Time) uint64

func main() {
	t := time.Now()
	fmt.Println(abs(t))
	fmt.Println(isLeap(2000))
}

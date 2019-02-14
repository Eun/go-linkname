package main

import (
	"fmt"
	_ "strings"
	"time"
	_ "unsafe"
)

//go:linkname TimeIsLeap time.isLeap
func TimeIsLeap(year int) bool

//go:linkname TimeTimeAbs time.Time.abs
func TimeTimeAbs(t time.Time) uint64

//go:linkname StringsStringFinder strings.stringFinder
type StringsStringFinder struct {
	pattern        string
	badCharSkip    [256]int
	goodSuffixSkip []int
}

//go:linkname StringsMakeStringFinder strings.makeStringFinder
func StringsMakeStringFinder(pattern string) *StringsStringFinder

func main() {
	fmt.Println(TimeIsLeap(2000))

	t := time.Now()
	fmt.Println(TimeTimeAbs(t))

	finder := StringsMakeStringFinder("Hello")
	fmt.Printf("%#v", finder)
}

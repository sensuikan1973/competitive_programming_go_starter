package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ref: https://zenn.dev/maitake/articles/go_de_atcoder
const (
	BUFSIZE = 10000000
)

var rdr *bufio.Reader

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	example()
	// solve() // TODO: implement solve()
}

func solve() {
}

func min(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func max(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func s2i(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func i2s(i int) string {
	return strconv.Itoa(i)
}

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readint() int {
	return s2i(readline())
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	return slice
}

// input example
// ----------
// 1 2 3
// aaaa
// bbbb
// cccc
// ----------
func example() {
	fmt.Printf("1st line intgers: %v\n", readIntSlice())
	for {
		fmt.Printf("string line: %s\n", readline())
	}
}

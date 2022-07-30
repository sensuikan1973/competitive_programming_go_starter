package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a int
	var s string

	fmt.Fscan(r, &a)
	fmt.Fscan(r, &s)

	fmt.Println("1st input: ", a)
	fmt.Println("2nd input: ", s)
}

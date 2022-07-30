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

	var d int
	var s string

	fmt.Fscan(r, &d)
	fmt.Fscan(r, &s)

	fmt.Fprintf(w, "1st input: %d, 2nd input: %s \n", d, s)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tmp_s := bufio.NewScanner(os.Stdin)
	for tmp_s.Scan() {
		fmt.Printf(tmp_s.Text() + "\n")
	}
}

package myapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MyLineFilter() {

	call()
}
func call() {
	scanner := bufio.NewScanner(os.Stdin)
	defer func() {
		fmt.Println(recover())
	}()
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		if (ucl) == "EXIT" {
			panic("exit")
		}
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

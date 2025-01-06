package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Fake output v0.1")
		os.Exit(0)
	}

	args := strings.Join(os.Args, " ")

	paramMap := map[string]string{
		"dump badging": badging_out,
		"badging":      badging_out,
		"dump":         dump_out,
	}

	for k, v := range paramMap {
		if strings.Contains(args, k) {
			fmt.Print(v)
			fmt.Println("")
			break
		}
	}
}

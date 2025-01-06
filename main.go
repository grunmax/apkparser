package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func out_(f string) {
	dat, err := os.ReadFile(f)
	check(err)
	fmt.Print(string(dat))
	fmt.Println("")
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Fake output")
		os.Exit(0)
	}

	args := strings.Join(os.Args, " ")

	paramJson, err := os.ReadFile("params")
	check(err)
	paramMap := map[string]string{}
	if err := json.Unmarshal([]byte(paramJson), &paramMap); err != nil {
		check(err)
	}

	for k, v := range paramMap {
		if strings.Contains(args, k) {
			out_(v)
			break
		}
	}
}

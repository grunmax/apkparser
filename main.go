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
		fmt.Println("Fake output v0.2")
		os.Exit(0)
	}

	args := strings.Join(os.Args, " ")

	paramMap := map[string]string{
		"settings_apk-debug.apk":                           settings_apk_debug,
		"qr-code-gen.apk":                                  qr_code_gen,
		"appium-uiautomator2-server-v":                     appium_uiautomator2_server_v,
		"appium-uiautomator2-server-debug-androidTest.apk": appium_uiautomator2_server_debug_androidTest,
		"ring-develop":                                     ring_develop,
	}

	for k, v := range paramMap {
		if strings.Contains(args, k) {
			fmt.Print(v)
			fmt.Println("")
			break
		}
	}
}

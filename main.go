package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/avast/apkparser"
)

func main() {
	apkName := ""
	for i := 0; i < 5; i++ {
		if strings.Contains(os.Args[i], ".apk") {
			apkName = os.Args[i]
			break
		}
	}
	if apkName == "" {
		fmt.Println("APK file is not found in parameters")
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")

	zipErr, resErr, manErr := apkparser.ParseApk(apkName, enc)
	if zipErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to open the APK: %s", zipErr.Error())
		os.Exit(1)
		return
	}

	if resErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse resources: %s", resErr.Error())
	}
	if manErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse AndroidManifest.xml: %s", manErr.Error())
		os.Exit(1)
		return
	}
	fmt.Println()
}

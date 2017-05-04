package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/dracher/autorestoxunit/libs"
)

func main() {
	v := libs.TestCase{}
	v.Name = "Check cockpit test run"
	v.Classname = "cockpit auto"

	v.SystemOut = "i am out"
	v.SystemErr = "i am err"
	v.Properties = []libs.Property{
		libs.Property{
			Name:  "polarion-testcase-id",
			Value: "INTE-87",
		},
		libs.Property{
			Name:  "polarion-testcase-id",
			Value: "INTE-89",
		},
	}

	output, err := xml.MarshalIndent(v, "  ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

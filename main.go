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
			Value: "RHEVM-17788",
		},
	}
	ts := libs.TestSuite{}
	ts.Tests = "2"
	ts.Errors = "0"
	ts.Failures = "0"
	ts.Skipped = "0"
	ts.TestCase = []libs.TestCase{
		v,
	}

	tss := libs.TestSuites{}
	tss.Properties = []libs.Property{
		libs.Property{
			Name:  "polarion-response-myteamsname",
			Value: "rhvhqe",
		},
		libs.Property{
			Name:  "polarion-project-id",
			Value: "RHEVM3",
		},
	}
	tss.TestSuite = []libs.TestSuite{
		ts,
	}

	output, err := xml.MarshalIndent(tss, "  ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

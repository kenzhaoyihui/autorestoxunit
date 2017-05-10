package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func marshalResult(r TestSuites) []byte {
	output, err := xml.MarshalIndent(r, "  ", "  ")
	if err != nil {
		log.Panic(err)
	}
	return output
}

// ParsedResult is
type ParsedResult interface {
	GenTestCases() map[string]string
	GenTestSuite() struct {
		Tests    int
		Errors   int
		Failures int
		Skipped  int
	}
	GenTestSuites() struct {
		ProjectID string
		Title     string
	}
}

func genTestCase(cases map[string]string) []TestCase {
	testCases := []TestCase{}

	for k, v := range cases {
		tc := TestCase{}
		tc.Name = k
		tc.SystemOut = v
		tc.Properties = []Property{
			Property{
				Name:  "polarion-testcase-id",
				Value: k,
			},
		}
		testCases = append(testCases, tc)
	}
	return testCases
}

func genTestSuite(tests, errors, failures, skipped int) TestSuite {
	ts := TestSuite{}
	ts.Tests = strconv.Itoa(tests)
	ts.Errors = strconv.Itoa(errors)
	ts.Failures = strconv.Itoa(failures)
	ts.Skipped = strconv.Itoa(skipped)
	return ts
}

func genTestSuites(projectID, title string) TestSuites {
	tss := TestSuites{}
	tss.Properties = []Property{
		Property{
			Name:  "polarion-response-myteamsname",
			Value: "rhvh-auto-team",
		},
		Property{
			Name:  "polarion-project-id",
			Value: projectID,
		},
		Property{
			Name:  "polarion-testrun-title",
			Value: fmt.Sprintf("4_1_Node_Install_AutoTest_%s", title),
		},
	}
	return tss
}

// RawToXunit is
func RawToXunit(val ParsedResult) []byte {
	tc := genTestCase(val.GenTestCases())
	tmp := val.GenTestSuite()
	ts := genTestSuite(tmp.Tests, tmp.Errors, tmp.Failures, tmp.Skipped)
	tmp2 := val.GenTestSuites()
	tss := genTestSuites(tmp2.ProjectID, tmp2.Title)

	ts.TestCase = tc
	tss.TestSuite = []TestSuite{ts}

	return marshalResult(tss)
}

package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
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
	GenTestSuites(projectID string) struct {
		ProjectID string
		Title     string
	}
}

func genTestCase(cases map[string]string) []interface{} {
	testCases := []interface{}{}

	for k, v := range cases {
		if v == "failed" {
			tc := TestCaseFailed{}
			tc.Name = k
			tc.SystemOut = v
			tc.Properties = []Property{
				Property{
					Name:  "polarion-testcase-id",
					Value: k,
				},
			}
			tc.FailedTestCase = FailedTestCase{
				Type:    "failure",
				Message: fmt.Sprintf("case %s is failed, please check the log to see details", k),
			}
			testCases = append(testCases, tc)
		} else {
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
			Value: title,
		},
	}
	return tss
}

// RawToXunit is
func RawToXunit(val ParsedResult, projectID string) []byte {
	tc := genTestCase(val.GenTestCases())
	tmp := val.GenTestSuite()
	ts := genTestSuite(tmp.Tests, tmp.Errors, tmp.Failures, tmp.Skipped)
	tmp2 := val.GenTestSuites(projectID)
	tss := genTestSuites(tmp2.ProjectID, tmp2.Title)

	ts.TestCase = tc
	tss.TestSuite = []TestSuite{ts}

	return marshalResult(tss)
}

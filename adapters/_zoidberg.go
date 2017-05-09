package adapters

import (
	"encoding/json"
	"fmt"
	"strconv"

	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

func NewZoidbergAdapter(rawResult string) *ZoidbergAdapter {
	z := ZoidbergAdapter{
		SrcFile:   rawResult,
		Parsed:    make(map[string]map[string]map[string]string),
		FlatCases: make(map[string]string),
	}
	z.Parse()
	return &z
}

// ZoidbergAdapter will convert zoidberg results
type ZoidbergAdapter struct {
	SrcFile   string
	Parsed    map[string]map[string]map[string]string
	FlatCases map[string]string
	BuildName []string
	KsFiles   []string
	Total     int
	Passed    int
	Failed    int
}

func (z *ZoidbergAdapter) parseSource() {
	v, err := ioutil.ReadFile(z.SrcFile)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(v, &z.Parsed)
	if err != nil {
		log.Warn(err)
	}
}

// Parse is
func (z *ZoidbergAdapter) Parse() *ZoidbergAdapter {
	z.parseSource()
	for b, v := range z.Parsed {
		z.BuildName = append(z.BuildName, b)
		for ks, vv := range v {
			z.KsFiles = append(z.KsFiles, ks)
			for k, vvv := range vv {
				z.FlatCases[k] = vvv
				if vvv == "passed" {
					z.Passed++
				} else {
					z.Failed++
				}
			}
		}
	}
	z.Total = len(z.FlatCases)
	return z
}

// GenTestCases is
func (z ZoidbergAdapter) GenTestCases() []TestCase {
	testCases := []TestCase{}

	for k, v := range z.FlatCases {
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

// GenTestSuite is
func (z ZoidbergAdapter) GenTestSuite() TestSuite {
	ts := TestSuite{}
	ts.Tests = strconv.Itoa(z.Total)
	ts.Errors = "0"
	ts.Failures = strconv.Itoa(z.Failed)
	ts.Skipped = "0"
	return ts
}

// GenTestSuites is
func (z ZoidbergAdapter) GenTestSuites() TestSuites {
	tss := TestSuites{}
	tss.Properties = []Property{
		Property{
			Name:  "polarion-response-myteamsname",
			Value: "rhvh-auto-team",
		},
		Property{
			Name:  "polarion-project-id",
			Value: "RHEVM3",
		},
		Property{
			Name:  "polarion-testrun-title",
			Value: fmt.Sprintf("4_1_Node_Install_AutoTest_%s", z.BuildName[0]),
		},
	}
	return tss
}

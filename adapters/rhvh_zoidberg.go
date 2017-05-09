package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

const (
	projectID          = "RHEVM3"
	testRunTitlePrefix = "4_1_Node_Install_AutoTest_"
)

// Zoidberg represent zoidberg test results
type Zoidberg struct {
	InputFile string
	details   map[string]map[string]map[string]string
	FlatCases map[string]string
	BuildName []string
	KsFiles   []string
	Total     int
	Passed    int
	Failed    int
}

// NewZoidberg is
func NewZoidberg(inputFile string) *Zoidberg {
	z := &Zoidberg{
		InputFile: inputFile,
		details:   make(map[string]map[string]map[string]string),
		FlatCases: make(map[string]string),
	}
	z.parseInputFile()
	z.getSummary()
	return z
}

func (z *Zoidberg) parseInputFile() {
	fp, err := ioutil.ReadFile(z.InputFile)
	if err != nil {
		log.Panic(err)
	}
	json.Unmarshal(fp, &z.details)
}

func (z *Zoidberg) getSummary() {
	for b, v := range z.details {
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
}

// GenTestCases is
func (z Zoidberg) GenTestCases() map[string]string {
	return z.FlatCases
}

// GenTestSuite is
func (z Zoidberg) GenTestSuite() struct {
	Tests    int
	Errors   int
	Failures int
	Skipped  int
} {
	return struct {
		Tests    int
		Errors   int
		Failures int
		Skipped  int
	}{z.Total, 0, z.Failed, 0}
}

// GenTestSuites is
func (z Zoidberg) GenTestSuites() struct {
	ProjectID string
	Title     string
} {
	return struct {
		ProjectID string
		Title     string
	}{
		projectID,
		fmt.Sprintf("%s%s", testRunTitlePrefix, z.BuildName[0]),
	}
}

// DebugSelf is
func (z *Zoidberg) DebugSelf() {
	log.Info(z.GenTestSuite().Tests)
	log.Info(z.GenTestSuites().Title)
}

package adapters

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type cockpitResult struct {
	Title   string            `json:"title"`
	Results map[string]string `json:"results"`
}

// Cockpit represent zoidberg test results
type Cockpit struct {
	InputFile   string
	details     cockpitResult
	FlatCases   map[string]string
	Total       int
	Passed      int
	Failed      int
	ReportTitle string
}

// NewCockpit is
func NewCockpit(inputFile string) *Cockpit {
	c := &Cockpit{
		InputFile: inputFile,
		details:   cockpitResult{Results: make(map[string]string)},
		FlatCases: make(map[string]string),
	}
	c.parseInputFile()
	c.getSummary()
	return c
}

func (c *Cockpit) parseInputFile() {
	fp, err := ioutil.ReadFile(c.InputFile)
	if err != nil {
		log.Panic(err)
	}
	json.Unmarshal(fp, &c.details)
	c.FlatCases = c.details.Results
	c.ReportTitle = c.details.Title
}

func (c *Cockpit) getSummary() {
	for _, v := range c.details.Results {
		if v == "passed" {
			c.Passed++
		} else if v == "failed" {
			c.Failed++
		}
	}
	c.Total = len(c.details.Results)
}

// GenTestCases is
func (c Cockpit) GenTestCases() map[string]string {
	return c.FlatCases
}

// GenTestSuite is
func (c Cockpit) GenTestSuite() struct {
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
	}{c.Total, 0, c.Failed, 0}
}

// GenTestSuites is
func (c Cockpit) GenTestSuites(projectID string) struct {
	ProjectID string
	Title     string
} {
	return struct {
		ProjectID string
		Title     string
	}{
		projectID,
		c.ReportTitle,
	}
}

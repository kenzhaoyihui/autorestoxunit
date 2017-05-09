package adapters

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// Zoidberg represent zoidberg test results
type Zoidberg struct {
	InputFile string
	details   map[string]map[string]map[string]string
	Summary   struct {
		Build     string
		Error     int
		Errorlist []string
		Failed    string
		Passed    int
		Total     int
	} `json:"sum"`
}

// NewZoidberg is
func NewZoidberg(inputFile string) *Zoidberg {
	z := &Zoidberg{
		InputFile: inputFile,
		details:   make(map[string]map[string]map[string]string),
		Summary: struct {
			Build     string
			Error     int
			Errorlist []string
			Failed    string
			Passed    int
			Total     int
		}{},
	}
	z.parseInputFile()
	return z
}

func (z *Zoidberg) parseInputFile() {
	fp, err := ioutil.ReadFile(z.InputFile)
	if err != nil {
		log.Panic(err)
	}
	json.Unmarshal(fp, &z.details)
}

// PrintSelf is
func (z Zoidberg) PrintSelf() {
	log.Info(z.details)
}

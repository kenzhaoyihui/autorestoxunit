package libs

import (
	"encoding/xml"
	"os"

	log "github.com/Sirupsen/logrus"

	ap "github.com/dracher/autorestoxunit/libs/adapters"
)

func debugOutput(r ap.TestSuites) {
	output, err := xml.MarshalIndent(r, "  ", "  ")
	if err != nil {
		log.Panic(err)
	}
	os.Stdout.Write(output)
}

type RawResults interface {
	GenTestCases() []ap.TestCase
	GenTestSuite() ap.TestSuite
	GenTestSuites() ap.TestSuites
}

func RawToXunit(val RawResults) {
	testCases := val.GenTestCases()
	testSuite := val.GenTestSuite()
	testSuites := val.GenTestSuites()

	testSuite.TestCase = testCases
	testSuites.TestSuite = []ap.TestSuite{testSuite}

	debugOutput(testSuites)
}

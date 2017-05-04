package libs

import "encoding/xml"

// Property is
type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

// TestCase represent single testcase in polarion
type TestCase struct {
	XMLName    xml.Name   `xml:"testcase"`
	Name       string     `xml:"name,attr"`
	Assertions string     `xml:"assertions,attr,omitempty"`
	Time       string     `xml:"time,attr,omitempty"`
	Timestamp  string     `xml:"timestamp,attr,omitempty"`
	Classname  string     `xml:"classname,attr,omitempty"`
	Class      string     `xml:"class,attr,omitempty"`
	Line       string     `xml:"line,attr,omitempty"`
	Log        string     `xml:"log,attr,omitempty"`
	Group      string     `xml:"group,attr,omitempty"`
	URL        string     `xml:"url,attr,omitempty"`
	SystemErr  string     `xml:"system-err,omitempty"`
	SystemOut  string     `xml:"system-out,omitempty"`
	Properties []Property `xml:"properties>property"`
}

// TestSuite is
type TestSuite struct {
	XMLName    xml.Name `xml:"testsuite"`
	Name       string   `xml:"name,attr,omitempty"`
	Tests      string   `xml:"tests,attr"`
	Failures   string   `xml:"failures,attr,omitempty"`
	Errors     string   `xml:"errors,attr,omitempty"`
	Time       string   `xml:"time,attr,omitempty"`
	Disabled   string   `xml:"disabled,attr,omitempty"`
	Skipped    string   `xml:"skipped,attr,omitempty"`
	Skips      string   `xml:"skips,attr,omitempty"`
	Timestamp  string   `xml:"timestamp,attr,omitempty"`
	Hostname   string   `xml:"hostname,attr,omitempty"`
	ID         string   `xml:"id,attr,omitempty"`
	Package    string   `xml:"package,attr,omitempty"`
	Assertions string   `xml:"assertions,attr,omitempty"`
	File       string   `xml:"file,attr,omitempty"`
	Skip       string   `xml:"skip,attr,omitempty"`
	Log        string   `xml:"log,attr,omitempty"`
	URL        string   `xml:"url,attr,omitempty"`
	TestCase   []TestCase
}

// TestSuites is
type TestSuites struct {
	XMLName    xml.Name   `xml:"testsuites"`
	Properties []Property `xml:"properties>property"`
	Name       string     `xml:"name,attr,omitempty"`
	Time       string     `xml:"time,attr,omitempty"`
	Tests      string     `xml:"tests,attr,omitempty"`
	Failures   string     `xml:"failures,attr,omitempty"`
	Disabled   string     `xml:"disabled,attr,omitempty"`
	Errors     string     `xml:"errors,attr,omitempty"`
	TestSuite  []TestSuite
}

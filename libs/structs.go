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

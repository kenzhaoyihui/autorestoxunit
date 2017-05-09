package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"

	"github.com/dracher/autorestoxunit/adapters"
	"github.com/dracher/autorestoxunit/parser"
)

var (
	credential = flag.String("c", "rhevm3_machine:polarion", "credential to use xunit importer")
	file       = flag.String("f", "/tmp/final_results.json", "results to upload, currently support .json")
	resultType = flag.String("t", "zoidberg", "raw results type cuurent support { zoidberg | cockpit }")
	upload     = flag.Bool("up", false, "upload to polarion immediately")
	adapter    parser.ParsedResult
)

const (
	uploadURL = "curl -k -u '%s:%s' -X POST -F file=@./%s https://polarion.engineering.redhat.com/polarion/import/xunit"
)

func init() {
	flag.Parse()

}

func uploadToPolarion() {
	log.Warn("not finish yet")
}

func main() {
	if *resultType == "zoidberg" {
		adapter = adapters.NewZoidberg(*file)
	} else if *resultType == "cockpit" {
		log.Fatal("not finish yet")
	} else {
		log.Panic("Ezzzz")
	}
	parser.RawToXunit(adapter)

	if *upload {
		uploadToPolarion()
	}
}

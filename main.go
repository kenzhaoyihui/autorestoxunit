package main

import (
	"flag"

	"github.com/dracher/autorestoxunit/adapters"
)

var (
	credential = flag.String("c", "rhevm3_machine:polarion", "credential to use xunit importer")
	file       = flag.String("f", "/tmp/final_results.json", "results to upload, currently support .json")
)

const (
	uploadURL = "curl -k -u '%s:%s' -X POST -F file=@./%s https://polarion.engineering.redhat.com/polarion/import/xunit"
)

func init() {
	flag.Parse()
}

func main() {
	r := adapters.NewZoidberg(*file)
	r.PrintSelf()
}

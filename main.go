package main

import (
	"github.com/dracher/autorestoxunit/libs"
	"github.com/dracher/autorestoxunit/libs/adapters"
)

const (
	username  = "rhevm3_machine"
	password  = "polarion"
	uploadURL = "curl -k -u '%s:%s' -X POST -F file=@./%s https://polarion.engineering.redhat.com/polarion/import/xunit"
)

func main() {
	v := adapters.NewZoidbergAdapter("/tmp/final_results.json")
	libs.RawToXunit(v)
}

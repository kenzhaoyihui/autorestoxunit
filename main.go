package main

import (
	"flag"
	"os/exec"

	log "github.com/Sirupsen/logrus"

	"os"

	"io/ioutil"

	"fmt"

	"github.com/dracher/autorestoxunit/adapters"
	"github.com/dracher/autorestoxunit/parser"
)

var (
	credential = flag.String("c", "rhevm3_machine:polarion", "default credential can only upload res to project RHEVM3")
	file       = flag.String("f", "/tmp/final_results.json", "results to upload, currently support .json")
	resultType = flag.String("t", "zoidberg", "raw results type cuurent support { zoidberg | cockpit }")
	upload     = flag.Bool("up", false, "upload to polarion immediately")
	adapter    parser.ParsedResult
)

const (
	uploadCMD = "curl -k -u '%s' -X POST -F file=@%s %s"
	uploadURL = "https://polarion.engineering.redhat.com/polarion/import/xunit"
)

func init() {
	flag.Parse()
}

func uploadToPolarion(content []byte) {
	log.Infof("Start upload %s result to polarion with <%s>", *resultType, *credential)
	if err := ioutil.WriteFile("/tmp/xres.xml", content, 0644); err != nil {
		log.Fatal(err)
	}
	cmd := fmt.Sprintf(uploadCMD, *credential, "/tmp/xres.xml", uploadURL)

	stdoutStderr, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}

func main() {
	if *resultType == "zoidberg" {
		adapter = adapters.NewZoidberg(*file)
	} else if *resultType == "cockpit" {
		log.Fatal("not finish yet")
	} else {
		log.Panic("Ezzzz")
	}
	res := parser.RawToXunit(adapter)

	if *upload {
		uploadToPolarion(res)
	} else {
		os.Stdout.Write(res)
	}
}

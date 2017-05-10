package main

import (
	"crypto/tls"
	"flag"

	log "github.com/Sirupsen/logrus"

	"os"

	"io/ioutil"

	"net/http"

	"strings"

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
	uploadURL = "https://polarion.engineering.redhat.com/polarion/import/xunit"
)

func init() {
	flag.Parse()
}

func uploadToPolarion(content []byte) {
	log.Infof("Start upload %s result to polarion with <%s>", *resultType, *credential)
	resfile, err := ioutil.TempFile("/tmp", "xunit")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(resfile.Name())

	if _, err := resfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := resfile.Close(); err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", uploadURL, resfile)
	if err != nil {
		log.Fatal(err)
	}
	cred := strings.Split(*credential, ":")
	req.SetBasicAuth(cred[0], cred[1])
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
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

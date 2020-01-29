package main

import (
	"sudo-bangbang.com/osg/fetch"
	"sudo-bangbang.com/osg/investigate"
	"github.com/jpillora/opts"
	"fmt"
)

// Commandline options
type config struct {
	Repo string `opts:"mode=arg,name=URI,help=URI\t\t\tURL or local path to a repository (or zip/tar with a repository)"`
}

func main() {
	c := config{}
	opts.Parse(&c)

	file_pipeline := make(chan fetch.File)
	go fetch.StartDataPipeline(&c.Repo, file_pipeline)

	file_audit_pipe := make(chan investigate.FileAudit)
	go investigate.StartAuditPipeline(file_pipeline, file_audit_pipe)

	for audit_report := range file_audit_pipe {
		fmt.Print(audit_report)
	}
}

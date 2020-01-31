package main

import (
	"sudo-bangbang.com/osg/fetch"
	"sudo-bangbang.com/osg/investigate"
	"github.com/jpillora/opts"
	"fmt"
	"os"
)

// Commandline options
type config struct {
	RemoteZip string `opts:"name=remote-zip, help=URL to a repository download that's zipped (zip)"`
	RemoteTar string `opts:"name=remote-tar, help=URL to a repository download that's gzipped (tar.gz)"`

	LocalTar  string `opts:"name=tar,        help=Path to a repository that's gzipped (tar.gz)"`
	LocalZip  string `opts:"name=zip,        help=Path to a repository that's zipped (zip)"`
	LocalDir  string `opts:"name=dir,        help=Path to a repository locally"`
}

func main() {
	c := config{}
	opts.Parse(&c)

	if c.LocalDir  != "" { start(&c.LocalDir,  fetch.LocalRepoDir) ; return }
	if c.LocalTar  != "" { start(&c.LocalTar,  fetch.LocalRepoTar) ; return }
	if c.LocalZip  != "" { start(&c.LocalZip,  fetch.LocalRepoZip) ; return }
	if c.RemoteTar != "" { start(&c.RemoteTar, fetch.RemoteRepoTar); return }
	if c.RemoteZip != "" { start(&c.RemoteZip, fetch.RemoteRepoZip); return }
}

func start(identifier *string, fn func(*string, chan fetch.Option)) {
	files  := make(chan fetch.Option)

	go fn(identifier, files)

	file_audit_pipe := make(chan investigate.Option)
	go investigate.StartAuditPipeline(files, file_audit_pipe)

	for audit_report := range file_audit_pipe {
		if audit_report.Err != nil {
			fmt.Println(audit_report.Err)
			os.Exit(1)
		}

		fmt.Print(audit_report.FileAudit)
	}
}

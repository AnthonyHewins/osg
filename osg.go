package main

import (
	"sudo-bangbang.com/osg/fetch"
	"sudo-bangbang.com/osg/investigate"
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		help(1, "not enough args")
	case 2:
		switch os.Args[1] {
		case "-h", "h", "--help", "help":
			help(0)
		default:
			help(1, "don't understand mode:", os.Args[1])
		}
	case 3:
		normal_audit(os.Args[1], os.Args[2])
	case 4:
		github_audit(os.Args[1], os.Args[2], os.Args[3], nil)
	case 5:
		github_audit(os.Args[1], os.Args[2], os.Args[3], &os.Args[4])
	default:
		help(1, "wrong number of args")
	}
}

func normal_audit(mode, arg string) {
	var fn func(*string, chan fetch.Option)
	switch mode {
	case "dir", "d":
		fn = fetch.LocalRepoDir
	case "zip", "z":
		fn = fetch.LocalRepoZip
	case "tar", "t":
		fn = fetch.RemoteRepoTar
	case "remote-zip", "rz":
		fn = fetch.RemoteRepoZip
	case "remote-tar", "rt":
		fn = fetch.RemoteRepoTar
	case "github", "g":
		help(1, "not enough args for github mode")
	default:
		help(1, "don't understand mode:", mode)
	}

	if arg == "" {
		help(1, "Missing ARG: should be a URL or path")
	}

	files := make(chan fetch.Option)
	go fn(&arg, files)
	start_audit(files)
}

func github_audit(mode, username, repo string, branch *string) {
	if mode != "github" && mode != "g" {
		help(1, "unrecognized mode:", mode)
	}

	files := make(chan fetch.Option)
	go fetch.GithubRepo(username, repo, branch, files)
	start_audit(files)
}

func start_audit(files chan fetch.Option) {
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

func help(exitCode int, extraMessages ...interface{}) {
	if len(extraMessages) != 0 {
		fmt.Println(extraMessages...)
	}
	fmt.Println(helpText)
	os.Exit(exitCode)
}

const helpText =
`osg MODE ARG1 ARG2 [...ARG3]

MODES             ARGS                 DESCRIPTION
-------------------------------------------------------------------------------
help                                   Show this help text

dir, d            REPO                 audit local directory REPO
zip, z            REPO.zip             audit local zip file REPO.zip
tar, t            REPO.tar             audit local tar file REPO.tar

remote-zip, rz    URL                  audit an arbitrary ZIP file at URL
remote-tar, rt    URL                  audit an arbitrary TAR file at URL

github, g         USER REPO [BRANCH]   audit a github repo using the github API`

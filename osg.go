package main

import (
	"sudo-bangbang.com/osg/fetch"
	"sudo-bangbang.com/osg/investigate"
	"fmt"
	"os"
)

// Commandline options
type config struct {
	Mode   string `opts:"mode=arg,  help=One of {remote-zip remote-tar zip tar dir github}"`
	Arg    string `opts:"mode=arg,  help=The path or URL or github username"`

	Repo   string `opts:"mode=flag, help=The repo you'd like to audit in "`
	Branch string `opts:"mode=flag, help=The branch you'd like to audit (default is master)"`
}

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("not enough args")
		help()
		os.Exit(1)
	case 2:
		if os.Args[1] != "h" || os.Args[1] != "help" {
			unrecognized(os.Args[1])
		} else {
			help()
		}
	case 3:
		normal_audit(os.Args[1], os.Args[2] )
	case 4:
		github_audit(os.Args[1], os.Args[2], os.Args[3], nil)
	case 5:
		github_audit(os.Args[1], os.Args[2], os.Args[3], &os.Args[4])
	default:
		fmt.Println("too many args")
		help()
		os.Exit(1)
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
		fmt.Println("not enough args for github mode")
		help()
		os.Exit(1)
	default:
		unrecognized(mode)
	}

	if arg == "" {
		fmt.Println("Missing ARG: should be a URL or path")
		os.Exit(1)
	}

	files := make(chan fetch.Option)
	go fn(&arg, files)
	start_audit(files)
}

func github_audit(mode, username, repo string, branch *string) {
	if mode != "github" && mode != "g" {
		unrecognized(mode)
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

func unrecognized(arg string) {
	fmt.Printf("unrecognized arg: %v\n", arg)
	help()
	os.Exit(1)
}

func help() {
	fmt.Println(
		"osg MODE ARG1 ARG2 [...ARG3]\n\n",
		"MODES             ARGS               DESCRIPTION\n",
		"-----------------------------------------------------------------------------\n",
		"help                                 Show this help text\n\n",

		"dir, d            REPO               audit local directory REPO\n",
		"zip, z            REPO.zip           audit local zip file REPO.zip\n",
		"tar, t            REPO.tar           audit local tar file REPO.tar\n\n",

		"remote-zip, rz    URL                audit an arbitrary ZIP file at URL\n",
		"remote-tar, rt    URL                audit an arbitrary TAR file at URL\n\n",

		"github, g         USER REPO [BRANCH] audit a github repo using the github API\n",
		"                                     if BRANCH isn't specified 'master' is used\n",
	)
}

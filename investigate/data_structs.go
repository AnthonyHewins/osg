package investigate

import (
	"fmt"
	"strings"
	"sudo-bangbang.com/osg/fetch"
)

type Reason = uint8

const (
	INET            = 1 << iota
	SUSPICIOUS_WORD = 1 << iota // examples such as "Analytics" case insensitive appearing in the source file
)

type Option struct {
	FileAudit *FileAudit
	Err error
}

type FileAudit struct {
	File              *fetch.File
	MetadataConcerns   Reason
	SrcConcerns      []SrcConcern
}

func (p FileAudit) String() string {
	var output strings.Builder
	output.WriteString(p.File.Name)

	if SUSPICIOUS_WORD & p.MetadataConcerns > 0 { output.WriteString(" SUSP_WORD,") }
	output.WriteString("\n")

	for _, line := range p.SrcConcerns {
		output.WriteString( fmt.Sprintf("  %v\n", line) )
	}

	return output.String()
}

type SrcConcern struct {
	Lineno uint64
	Line   string
	Reason Reason
}

func (sc SrcConcern) String() string {
	var reason strings.Builder

	if INET            & sc.Reason > 0 { reason.WriteString("INET,")      }
	if SUSPICIOUS_WORD & sc.Reason > 0 { reason.WriteString("SUSP_WORD,") }

	return fmt.Sprintf("%v: %v (%v)", sc.Lineno, sc.Line, reason.String())
}

package investigate

import (
	"fmt"
	"strings"
	"sudo-bangbang.com/osg/fetch"
)

type Reason = uint8

// These are bit flags. They are logically OR'd together to produce a state machine of why
const (
	URL             = 1 << iota
	SUSPICIOUS_WORD = 1 << iota // examples such as "Analytics" case insensitive appearing in the source file
)

type FileAudit struct {
	File              *fetch.File
	MetadataConcerns   Reason
	SrcConcerns      []SrcConcern
}

func (p FileAudit) String() string {
	var output strings.Builder
	output.WriteString(p.File.Name)

	if SUSPICIOUS_WORD & p.MetadataConcerns > 0 { output.WriteString(" SUSP_WORD") }
	output.WriteString("\n")

	for _, line := range p.SrcConcerns {
		output.WriteString( fmt.Sprintf("%v\n", line) )
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

	if URL             & sc.Reason > 0 { reason.WriteString("URL,") }
	if SUSPICIOUS_WORD & sc.Reason > 0 { reason.WriteString("SUSP_WORD,") }

	return fmt.Sprintf("%v: %v (%v)", sc.Lineno, sc.Line, reason.String())
}

// TODO split on the byte value of \n and reduce the cost of audit
func StartAuditPipeline(file_pipe chan fetch.File, audit_pipe chan FileAudit) {
	for file := range file_pipe {
		lines := strings.Split(string(file.Contents), "\n")

		var src_concerns []SrcConcern
		for lineno, line := range lines {
			var bitmask Reason

			bitmask =           contains_url(&line)
			bitmask = bitmask | contains_suspicious_word(&line)

			if bitmask > 0 {
				src_concerns = append(
					src_concerns,
					SrcConcern{uint64(lineno), line, bitmask},
				)
			}
		}

		filename_suspicious := contains_suspicious_word(&file.Name)

		if int(filename_suspicious) | len(src_concerns) <= 0 { continue }

		audit_pipe <- FileAudit{
			File: &file,
			MetadataConcerns: filename_suspicious,
			SrcConcerns: src_concerns,
		}
	}

	close(audit_pipe)
}

func contains_url(line *string) Reason {
	if url_regex.MatchString(*line) {
		return URL
	}

	return 0
}

func contains_suspicious_word(line *string) Reason {
	if suspicious_words.MatchString(*line) {
		return SUSPICIOUS_WORD
	}

	return 0
}

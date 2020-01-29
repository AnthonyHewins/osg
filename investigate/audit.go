package investigate

import (
	"strings"
	"sudo-bangbang.com/osg/fetch"
)

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

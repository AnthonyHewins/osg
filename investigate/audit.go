package investigate

import (
	"strings"
	"sync"
	"sudo-bangbang.com/osg/fetch"
)

// TODO split on the byte value of \n and reduce the cost of audit
func StartAuditPipeline(file_pipe chan fetch.Option, audit_pipe chan Option) {
	defer close(audit_pipe)

	for option := range file_pipe {
		if option.Err != nil {
			audit_pipe <- Option{FileAudit: nil, Err: option.Err}
			return
		}

		file := option.File
		lines := strings.Split(string(file.Contents), "\n")

		var src_concerns []SrcConcern
		for lineno, line := range lines {
			var bitmask Reason
			var wg_for_src sync.WaitGroup

			wg_for_src.Add(2)
			go contains_url(&line, &bitmask, &wg_for_src)
			go contains_suspicious_word(&line, &bitmask, &wg_for_src)
			wg_for_src.Wait()

			if bitmask > 0 {
				src_concerns = append(
					src_concerns,
					SrcConcern{uint64(lineno) + 1, line, bitmask},
				)
			}
		}

		filename_suspicious := uint8(0) //contains_suspicious_word(&file.Name)

		if int(filename_suspicious) | len(src_concerns) <= 0 { continue }

		audit := FileAudit{
			File: file,
			MetadataConcerns: filename_suspicious,
			SrcConcerns: src_concerns,
		}

		audit_pipe <- Option{FileAudit: &audit, Err: nil}
	}
}

func contains_url(line *string, bitmask *Reason, wg *sync.WaitGroup) {
	defer wg.Done()
	if url_regex.MatchString(*line) {
		*bitmask = *bitmask | URL
	}
}

func contains_suspicious_word(line *string, bitmask *Reason, wg *sync.WaitGroup) {
	defer wg.Done()
	if suspicious_words.MatchString(*line) {
		*bitmask = *bitmask | SUSPICIOUS_WORD
	}
}

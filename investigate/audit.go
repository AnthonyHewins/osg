package investigate

import (
	"bytes"
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

		file  := option.File

		// 0. Prepare ourselves for performing two tasks async
		var wg sync.WaitGroup
		wg.Add(2)

		// 1. Scan the source code asynchronously, results in this array
		var src_concerns []SrcConcern
		go scan_src(&file.Contents, &src_concerns, &wg)

		// 2. Scan the metadata asynchronously, results in this Reason (a uint8)
		var metadata_concerns Reason
		go scan_metadata(file, &metadata_concerns, &wg)

		// 3. Wait til 1 and 2 are done, then proceed...
		wg.Wait()

		if int(metadata_concerns) | len(src_concerns) <= 0 { continue }

		audit := FileAudit{
			File: file,
			MetadataConcerns: metadata_concerns,
			SrcConcerns: src_concerns,
		}

		audit_pipe <- Option{FileAudit: &audit, Err: nil}
	}
}

func scan_src(buf *[]byte, src_concerns *[]SrcConcern, wg *sync.WaitGroup) {
	defer wg.Done()

	lines := bytes.Split(*buf, []byte("\n"))
	for lineno, line := range lines {
		var bitmask Reason
		var wg_for_checking_in_parallel sync.WaitGroup

		wg_for_checking_in_parallel.Add(2)
		go contains_url(&line, &bitmask, &wg_for_checking_in_parallel)
		go contains_suspicious_word(&line, &bitmask, &wg_for_checking_in_parallel)
		wg_for_checking_in_parallel.Wait()

		if bitmask > 0 {
			*src_concerns = append(
				*src_concerns,
				SrcConcern{uint64(lineno) + 1, string(line), bitmask},
			)
		}
	}
}

func scan_metadata(f *fetch.File, metadata_concerns *Reason, wg *sync.WaitGroup) {
	byte_array_of_name := []byte(f.Name)
	contains_suspicious_word(&byte_array_of_name, metadata_concerns, wg)
}

func contains_url(line *[]byte, bitmask *Reason, wg *sync.WaitGroup) {
	defer wg.Done()
	if url_regex.Match(*line) {
		*bitmask = *bitmask | URL
	}
}

func contains_suspicious_word(line *[]byte, bitmask *Reason, wg *sync.WaitGroup) {
	defer wg.Done()
	if suspicious_words.Match(*line) {
		*bitmask = *bitmask | SUSPICIOUS_WORD
	}
}

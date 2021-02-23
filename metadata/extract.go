package metadata

import (
	"encoding/json"
	"strings"
)

func Extract(body string, data interface{}) (bool, error) {
	for _, line := range strings.Split(body, "\n") {
		if !strings.HasPrefix(line, embeddedCommentPrefix) {
			continue
		}
		if !strings.HasSuffix(line, embeddedCommentSuffix) {
			continue
		}
		if err := json.Unmarshal([]byte(line[lenEmbeddedCommentPrefix:len(line)-lenEmbeddedCommentSuffix]), &data); err != nil {
			continue
		}
		return true, nil
	}
	return false, nil
}

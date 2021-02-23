package metadata

import (
	"encoding/json"
	"fmt"
)

const (
	embeddedCommentPrefix    = "<!-- github-comment: "
	embeddedCommentSuffix    = " -->"
	lenEmbeddedCommentPrefix = len(embeddedCommentPrefix)
	lenEmbeddedCommentSuffix = len(embeddedCommentSuffix)
)

func Convert(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("marshal an embedded metadata to JSON: %w", err)
	}
	return "\n" + embeddedCommentPrefix + string(b) + embeddedCommentSuffix, nil
}

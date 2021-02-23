package metadata_test

import (
	"testing"

	"github.com/suzuki-shunsuke/github-comment-metadata/metadata"
)

func TestConvert(t *testing.T) {
	t.Parallel()
	data := []struct {
		caseName string
		data     interface{}
		exp      string
		isErr    bool
	}{
		{
			caseName: "normal",
			data:     map[string]string{"foo": "bar"},
			exp: `
<!-- github-comment: {"foo":"bar"} -->`,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.caseName, func(t *testing.T) {
			t.Parallel()
			f, err := metadata.Convert(d.data)
			if d.isErr {
				if err == nil {
					t.Fatal("err should not be nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if d.exp != f {
				t.Fatalf("got %s, want %s", f, d.exp)
			}
		})
	}
}

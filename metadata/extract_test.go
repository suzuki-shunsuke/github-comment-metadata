package metadata_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/github-comment-metadata/metadata"
)

func TestExtract(t *testing.T) {
	t.Parallel()
	data := []struct {
		caseName string
		body     string
		data     interface{}
		exp      interface{}
		expB     bool
		isErr    bool
	}{
		{
			caseName: "normal",
			body:     `<!-- github-comment: {"foo": "bar"} -->`,
			data:     &map[string]string{},
			exp:      &map[string]string{"foo": "bar"},
			expB:     true,
		},
		{
			caseName: "not match",
			body:     `hello`,
			expB:     false,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.caseName, func(t *testing.T) {
			t.Parallel()
			f, err := metadata.Extract(d.body, d.data)
			if d.isErr {
				if err == nil {
					t.Fatal("err should not be nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if d.expB != f {
				t.Fatalf("got %t, want %t", f, d.expB)
			}
			if !reflect.DeepEqual(d.exp, d.data) {
				t.Fatalf("got %t, want %t", d.data, d.exp)
			}
		})
	}
}

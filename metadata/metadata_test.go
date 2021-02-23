package metadata_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/github-comment-metadata/metadata"
)

func TestSetCIData(t *testing.T) { //nolint:funlen
	t.Parallel()
	data := []struct {
		caseName string
		ci       string
		env      map[string]string
		exp      map[string]interface{}
		isErr    bool
	}{
		{
			caseName: "circleci",
			ci:       "circleci",
			env: map[string]string{
				"CIRCLE_JOB":             "build",
				"CIRCLE_WORKFLOW_JOB_ID": "xxx",
			},
			exp: map[string]interface{}{
				"JobName": "build",
				"JobID":   "xxx",
			},
		},
		{
			caseName: "drone",
			ci:       "drone",
			env: map[string]string{
				"DRONE_STATE_NAME": "build",
				"DRONE_STEP_NAME":  "checkout",
			},
			exp: map[string]interface{}{
				"WorkflowName": "build",
				"JobName":      "checkout",
			},
		},
		{
			caseName: "github-actions",
			ci:       "github-actions",
			env: map[string]string{
				"GITHUB_WORKFLOW": "build",
				"GITHUB_JOB":      "checkout",
			},
			exp: map[string]interface{}{
				"WorkflowName": "build",
				"JobName":      "checkout",
			},
		},
		{
			caseName: "codebuild",
			ci:       "codebuild",
			env: map[string]string{
				"CODEBUILD_BUILD_ID": "zzz",
			},
			exp: map[string]interface{}{
				"JobID": "zzz",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.caseName, func(t *testing.T) {
			t.Parallel()
			data := map[string]interface{}{}
			err := metadata.SetCIEnv(d.ci, func(key string) string {
				return d.env[key]
			}, data)
			if d.isErr {
				if err == nil {
					t.Fatal("err should not be nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(d.exp, data) {
				t.Fatalf("got %+v, want %+v", data, d.exp)
			}
		})
	}
}

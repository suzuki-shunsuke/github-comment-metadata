package metadata

func SetCIEnv(ci string, getEnv func(string) string, data map[string]interface{}) error {
	switch ci {
	case "circleci":
		data["JobName"] = getEnv("CIRCLE_JOB")
		data["JobID"] = getEnv("CIRCLE_WORKFLOW_JOB_ID")
	case "drone":
		data["WorkflowName"] = getEnv("DRONE_STATE_NAME")
		data["JobName"] = getEnv("DRONE_STEP_NAME")
	case "github-actions":
		data["WorkflowName"] = getEnv("GITHUB_WORKFLOW")
		data["JobName"] = getEnv("GITHUB_JOB")
	case "codebuild":
		data["JobID"] = getEnv("CODEBUILD_BUILD_ID")
	}
	return nil
}

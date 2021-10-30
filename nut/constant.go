package nut

const (
	Pending Status = "pending"
	Running Status = "running"
	Skipped Status = "skipped"
	Passed  Status = "passed"
	Failed  Status = "failed"

	ParamBody     ParamType = "body_params"
	ParamQuery    ParamType = "url_query_params"
	ParamFormData ParamType = "form_data_params"
)

type Status string

type ParamType string

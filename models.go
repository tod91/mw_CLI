package main

// EvaluateReq ...
//
// used to store the request body when Evaluate is called
type EvaluateReq struct {
	Expression string `json:"expression"`
}

// EvaluateResp ...
//
// used to store the response body when Evaluate is called
type EvaluateResp struct {
	Result string `json:"result"`
}

// ValidateReq ...
//
// used to store the request  body when Validate is called
type ValidateReq struct {
	Expression string `json:"expression"`
}

// ValidateResp ...
//
// used to store the response  body when Validate is called
type ValidateResp struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason"`
}

// ErrorsResp ...
//
// used to store the response  body when errors is called
type ErrorsResp struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}

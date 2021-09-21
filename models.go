package main

type EvaluateReq struct {
	Expression string `json:"expression"`
}

type EvaluateResp struct {
	Result string `json:"result"`
}

type ValidateReq struct {
	Expression string `json:"expression"`
}

type ValidateResp struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason"`
}

type ErrorsResp struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}

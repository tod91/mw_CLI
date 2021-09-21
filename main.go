package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var evalInput string
	var validateInput string
	var errorsFlag bool
	flag.StringVar(&evalInput, "eval", "", "evaluates a math expression. Expression must be in the form "+
		"`{\"expression\":\"<a\nsimple math problem>\"}`")

	flag.StringVar(&validateInput, "validate", "", "validates a given expression. Request must be in the form "+
		"`{\"expression\":\"<a\nsimple math problem>\"}`")
	flag.BoolVar(&errorsFlag, "errors", false, "returns all errors that have occurred during every call. Request must be in the form"+
		" \"`{\\\"expression\\\":\\\"<a\\nsimple math problem>\\\"}`\")")

	flag.Parse()

	if !onlyOneIsSet(evalInput, validateInput, errorsFlag) {
		log.Fatal("only one flag needs to be set at a time")
	}

	switch {
	case evalInput != "":
		res := evaluateExpression(evalInput)
		fmt.Printf("Here's your result: %s\n", res.Result)
	case validateInput != "":
		res := validateExpr(validateInput)
		if res.Valid {
			fmt.Println("The expression is valid")
		} else {
			fmt.Printf("The expression is invalid. Reason: %s\n", res.Reason)
		}
	case errorsFlag:
		res := getErrors()
		fmt.Println("Here's all the errors which occurred:")
		for _, err := range res {
			fmt.Printf(" Expression: %s\n\t Endpoint: %s\n\t Frequency: %d\n\t Error Type: %s\n\n",
				err.Expression, err.Endpoint, err.Frequency, err.Type)
		}
	}
}

func evaluateExpression(expr string) *EvaluateResp {
	reqBody, _ := json.Marshal(&EvaluateReq{expr})
	resp, err := http.Post("http://localhost:8080/evaluate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var res *EvaluateResp
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res
}

func validateExpr(expr string) *ValidateResp {
	reqBody, _ := json.Marshal(&ValidateReq{expr})
	resp, err := http.Post("http://localhost:8080/validate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var res *ValidateResp
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res
}

func getErrors() []*ErrorsResp {
	resp, err := http.Get("http://localhost:8080/errors")
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var res []*ErrorsResp
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res
}

func onlyOneIsSet(f1, f2 string, f3 bool) bool {
	return f1 != "" && f2 == "" && !f3 || f1 == "" && f2 != "" && !f3 || f1 == "" && f2 == "" && f3
}

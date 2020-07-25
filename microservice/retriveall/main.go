package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type word struct {
	Word string `json:"word"`
	Freq int    `json:"freq"`
}

type retriveWordRes struct {
	Words []word `json:"words"`
}

func retriveHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var dummyWords [2]word
	dummyWords[0] = word{Word: "hello", Freq: 2}
	dummyWords[1] = word{Word: "world", Freq: 3}

	var resp, err = json.Marshal(dummyWords)
	if err != nil {
		clientError(http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(resp),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(retriveHandler)
}

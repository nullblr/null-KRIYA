package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type addReqBody struct {
	Word string `name:"word"`
}

func saveWordHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody addReqBody
	err := json.Unmarshal([]byte(req.Body), &reqBody)
	if err != nil || reqBody.Word == "" {
		return clientError(http.StatusBadRequest)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("Saved %s", reqBody.Word),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(saveWordHandler)
}

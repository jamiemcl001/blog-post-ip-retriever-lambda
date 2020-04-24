package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// ResponseBody is the structure of the body message returned
// to the user.
type ResponseBody struct {
	IPAddress string `json:"ipAddress"`
}

var (
	// ErrFailedToReturnIP indicates an error occured when trying to retrieve
	// the IP from CanIHazIp
	ErrFailedToReturnIP = errors.New("Failed to retrieve IP from canihazip")
)

func handler(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get("https://canihazip.com/s")
	if err != nil {
		return events.APIGatewayProxyResponse{}, ErrFailedToReturnIP
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	jsonBody := ResponseBody{
		IPAddress: string(bytes),
	}

	marshalledResponse, err := json.Marshal(jsonBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(marshalledResponse),
	}, nil
}

func main() {
	lambda.Start(handler)
}

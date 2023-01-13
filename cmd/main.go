package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cdgbabies/list-blogs-lambda/service"
)

type handler struct {
	dynamoDbClient service.DynamoDbReadOperationClient
}

func (h *handler) handleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	blogs, err := service.QueryDynamoDB(ctx, h.dynamoDbClient)
	if err != nil {
		return events.LambdaFunctionURLResponse{}, err
	}
	body, err := json.Marshal(blogs)
	if err != nil {
		return events.LambdaFunctionURLResponse{}, err
	}
	return events.LambdaFunctionURLResponse{Body: string(body), StatusCode: 200}, nil

}
func main() {
	h := handler{
		dynamoDbClient: service.NewDynamoDbClient(context.Background(), os.Getenv("region_name")),
	}
	lambda.Start(h.handleRequest)
}

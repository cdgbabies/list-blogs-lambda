package service

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Blog struct {
	Sk          string    `json:"sk" dynamodbav:"sk"`
	Pk          string    `json:"pk" dynamodbav:"pk"`
	Description string    `json:"description" dynamodbav:"description"`
	Title       string    `json:"title" dynamodbav:"title"`
	CreatedDate time.Time `json:"createdDate" dynamodbav:"createdDate"`
	User        string    `json:"user" dynamodbav:"user"`
}
type DynamoDbReadOperationClient interface {
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

func NewDynamoDbClient(ctx context.Context, region string) *dynamodb.Client {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Panic("unable to load SDK config,", err)
	}

	dynamoDbClient := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Region = region

	})

	return dynamoDbClient

}

func QueryDynamoDB(ctx context.Context, client DynamoDbReadOperationClient, tableName string) ([]Blog, error) {

	var blogs []Blog

	keyEx := expression.Key("pk").Equal(expression.Value("blogs"))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	response, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 &tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return nil, err
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &blogs)

		if err != nil {
			log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
			return nil, err
		}
	}
	return blogs, nil
}

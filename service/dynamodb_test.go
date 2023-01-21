package service

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type mockDynamoDbClient struct {
	response dynamodb.QueryOutput
	err      error
}

func (m *mockDynamoDbClient) Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return &m.response, m.err
}

func TestQueryDynamoDB(t *testing.T) {

	tableName := "Test"

	keyEx := expression.Key("pk").Equal(expression.Value("blogs"))
	_, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		t.Error("Error building expression: ", err)
	}
	mockClient := &mockDynamoDbClient{
		response: dynamodb.QueryOutput{
			Items: []map[string]types.AttributeValue{
				{

					"sk":          &types.AttributeValueMemberS{Value: "123"},
					"pk":          &types.AttributeValueMemberS{Value: "blogs"},
					"description": &types.AttributeValueMemberS{Value: "Description"},
					"title":       &types.AttributeValueMemberS{Value: "Test Title"},
					"createdDate": &types.AttributeValueMemberS{Value: "2022-01-01T00:00:00Z"},
				},
			},
		},
		err: nil,
	}
	result, err := QueryDynamoDB(context.TODO(), mockClient, tableName)
	if err != nil {
		t.Error("Error querying DynamoDB: ", err)
	}
	if len(result) != 1 {
		t.Error("Incorrect number of items returned")
	}
	if result[0].Title != "Test Title" {
		t.Error("Incorrect item returned")
	}

	// Test error scenario
	mockClient.err = errors.New("Test error")
	result, err = QueryDynamoDB(context.TODO(), mockClient, tableName)
	if err == nil {
		t.Error("Error scenario not handled correctly")
	}
}

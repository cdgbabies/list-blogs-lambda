package domain

import "time"

type Blog struct {
	Sk          string    `json:"sk" dynamodbav:"sk"`
	Pk          string    `json:"pk" dynamodbav:"pk"`
	Description string    `json:"description" dynamodbav:"description"`
	Title       string    `json:"title" dynamodbav:"title"`
	CreatedDate time.Time `json:"createdDate" dynamodbav:"createdDate"`
}

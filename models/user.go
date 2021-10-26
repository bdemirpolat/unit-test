package models

type User struct {
	Username string `json:"username" dynamodbav:"username"`
	Age      int    `json:"age" dynamodbav:"age"`
}

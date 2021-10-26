package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bdemirpolat/unit-test/models"
)

type UserRepository interface {
	Create(user models.User) error
	Delete(tableName string) error
	Get(username string, age string) (*models.User, error)
}

type UserRepo struct {
	Client *dynamodb.DynamoDB
}

func (u UserRepo) Create(user models.User) error {
	userMap, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      userMap,
		TableName: aws.String("users"),
	}

	_, err = u.Client.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepo) Delete(tableName string) error {
	return nil
}

func (u UserRepo) Get(username string, age string) (*models.User, error) {
	result, err := u.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
			"age": {
				N: aws.String(age),
			},
		},
	})
	user := models.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

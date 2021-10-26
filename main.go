package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/bdemirpolat/unit-test/models"
	"github.com/bdemirpolat/unit-test/repository"
	"github.com/gofiber/fiber/v2"
)

func CreateClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-west-2"), Endpoint: aws.String("http://localhost:8000")},
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}

func DeleteTable(client *dynamodb.DynamoDB) error {
	_, err := client.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String("users")})
	return err
}

func ListTables(client *dynamodb.DynamoDB) ([]*string, error) {
	input := &dynamodb.ListTablesInput{}

	result, err := client.ListTables(input)
	return result.TableNames, err
}

func CreateTable(client *dynamodb.DynamoDB) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("username"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("age"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("username"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("age"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("users"),
	}

	_, err := client.CreateTable(input)
	return err
}

func main() {
	client := CreateClient()
	createTable := flag.Bool("create-table", false, "create table")
	listTable := flag.Bool("list-table", false, "list table")
	deleteTable := flag.Bool("delete-table", false, "delete table")
	flag.Parse()
	if *createTable {
		err := CreateTable(client)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("table created")
		return
	}
	if *listTable {
		tables, err := ListTables(client)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found %d tables\n", len(tables))
		for _, table := range tables {
			fmt.Println(*table)
		}
		return
	}
	if *deleteTable {
		err := DeleteTable(client)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("table deleted")
		return
	}
	userRepo := &repository.UserRepo{Client: client}
	app := fiber.New()
	app.Post("/users", func(c *fiber.Ctx) error {
		user := models.User{}
		err := c.BodyParser(&user)
		if err != nil {
			return err
		}
		err = userRepo.Create(user)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(user)
	})
	app.Get("/users/:username/:age", func(c *fiber.Ctx) error {
		username := c.Params("username")
		if username == "" {
			return errors.New("username can not be empty")
		}
		age := c.Params("age")
		if username == "" {
			return errors.New("age can not be empty")
		}
		user, err := userRepo.Get(username, age)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(user)
	})
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

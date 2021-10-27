## Create users table
```
go run main.go -create-table
```

## Create user
```
curl --location --request POST 'localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"burak"
}'
```
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

# Go Test Basics

Go dilinde testlerimizi yapabilmemiz için biz geliştiricilere standart kütüphanesi içinde built-in gelen testing paketi ve go test komutu sağlıyor

## Go dilinde unit test görünümü aşağıdaki gibidir

```
func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("failed, got: %d, want: %d.", total, 10)
	}
}
```


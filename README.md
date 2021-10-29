# Go Test Basics

Go dilinde testlerimizi yapabilmemiz için biz geliştiricilere standart kütüphanesi içinde built-in gelen testing paketi ve go test komutu sağlıyor

## Go dilinde unit test görünümü aşağıdaki gibidir

```
func TestSum(t *testing.T) {
	total := Sum(2, 3)
	if total != 5 {
		t.Errorf("failed, got: %d, want: %d.", total, 10)
	}
}
```

# Golang test'e başlamadan önce bilmemiz gereken bazı özellikler
* Test fonksiyonu için ihtiyacımız olacak fonksiyon argümanı t *testing.T (built-in)
* Test methodlarımız isimlendirirken metodumuzun ismi büyük harfle Test daha sonra kelime ya da kelime öbeği olmalıdır.
  * Ör: (TestUserService)
* Test dosyasının adı something_test.go olacak şekilde olmalıdır.
* t.Log test sırasında loglamak istediğimiz bilgiler için kullanırız
* t.Error ya da t.Fail ile test hatalarını belirtek için kullanırız

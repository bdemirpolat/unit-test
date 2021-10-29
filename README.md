# Go Test Basics

Go dilinde testlerimizi yapabilmemiz için biz geliştiricilere standart kütüphanesi içinde built-in gelen testing paketi ve go test komutu sağlıyor

## Go dilinde unit test görünümü aşağıdaki gibidir

```
func TestSayHello(t *testing.T) {
	result := sayHello("Yemeksepeti")

	if result != "Hello Yemeksepeti!" {
		t.Errorf("Hello Yemeksepeti failed, expected -> %v, got -> %v", "Hello Yemeksepeti!", result)
	}
}
```

# Golang test'e başlamadan önce bilmemiz gereken bazı özellikler
* Test fonksiyonu için ihtiyacımız olacak fonksiyon argümanı t *testing.T (built-in gelen)
* Test methodlarımız isimlendirirken metodumuzun ismi büyük harfle Test daha sonra kelime ya da kelime öbeği olmalıdır.
  * Ör: (TestUserService)
* Test dosyasının adının sonuna _test.go olacak şekilde olmalıdır.
  * Ör: (something_test.go)
* t.Log test sırasında loglamak istediğimiz bilgiler için kullanırız (non-failing debug information)
* t.Error ya da t.Fail ile test hatalarını belirtek için kullanırız

# t.Testing kütüphanesini tanıyalım

# Test Yazımı

# Test Coverage
Uygulama içinde yazdığımız kodun test yüzdesini sağlar. 
Bu sayede yazdığımız kodun hangi taraflarını test etmediğimiz görebiliriz.

# Golang testleri çalıştırma komutları
* go test . -> geçerli dizindeki testleri run eder
* go test ./calc -> calc dizinindeki testleri run eder
* go test ./... modül içindeki tüm testleri run eder
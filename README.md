# Golang Test Basics

Go dilinde testlerimizi yapabilmemiz için standart kütüphanesi içinde built-in gelen **testing.T** paketi ve **go test** komutu sağlıyor

## Golang'te unit test görünümü aşağıdaki gibidir

```
func TestSayHello(t *testing.T) {
	result := sayHello("Yemeksepeti")

	if result != "Hello Yemeksepeti!" {
		t.Errorf("Hello Yemeksepeti failed, expected -> %v, got -> %v", "Hello Yemeksepeti!", result)
	}
}
```

# Golang test'e başlamadan önce bilmemiz gereken bazı özellikler
* Test fonksiyonu için ihtiyacımız olacak fonksiyon argümanı t *testing.T 
* Test methodlarımız isimlendirirken metodumuzun ismi büyük harfle yani Test daha sonra kelime ya da kelime öbeği olmalıdır.
  * Ör: (TestUserService)
* Test dosyasının adının sonuna _test.go olacak şekilde olmalıdır.
  * Ör: (something_test.go)
* t.Log test sırasında loglamak istediğimiz bilgiler için kullanırız (non-failing debug information)
* t.Error ya da t.Fail ile test hatalarını belirtek için kullanırız

# testing.T kütüphanesini tanıyalım
 * **t.Log** test sırasında loglamak istediğimiz bilgiler için kullanırız (non-failing debug information)
 * **t.Logf** test sırasında loglamak istediğimiz bilgileri istediğimiz formatta dönüştürmek için kullanırız
 * **t.Error** test hatalarını istediğimiz formatta belirtek için kullanırız
 * **t.Errorf** test hatalarını belirtek için kullanırız

# Test Yazımı
Go projemizi oluşturduktan sonra greeting.go adında bir dosya oluşturuyoruz. Test edeceğimiz fonksiyon main package altında olacak
```
package main

import (
	"fmt"
)

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hello Anonymous!"
	}

	return fmt.Sprintf("Hello %s!", name)
}
```

yukarıdaki koda baktığımızda name adında argüman alan ve geriye string dönen basit bir fonksiyon. Fonksiyonun amacı gönderilen parametreyi formatlayarak geriye dönmektedir. Eğer name argümanını boş bir değer gönderdiğimizde geriye `Hello Anonymous!` dönüyor.

## Şimdi **sayHello** fonksyionumuzun unit testini oluşturalım. 
İlk olarak aynı dizinde greeting_test.go adında bir dosya oluşturuyoruz.
```
package main

import (
	"fmt"
	"testing"
)

func TestSayHelloValidArgument(t *testing.T) {
	name := "Yemeksepeti"
	expected := fmt.Sprintf("Hello %s!", name)
	result := sayHello(name)

	if result != expected {
		t.Errorf("\"sayHello(%s)\" failed, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayHello(%s)\" succeded, expected -> %v, got -> %v", name, expected, result)
	}

}
```
Yukarıda ne yaptık?

TestSayHelloValidArgument adında bir unit testimiz var. Bir dosya içerisinde birden fazla test fonksiyonlarımız olabilir. Test fonksiyonumuz (adından da anlaşılacağı üzere) `sayHello` fonksiyonumuzu test edecek olan fonksiyonumuzdur. `sayHello` fonksiyonuna geçerli bir parametre göndererek çağrıdık ve sonucun beklediğimiz gibi olup olmadığını kontrol ettik. Eğer aldığımız sonuç beklediğimiz bir sonuç değilse `t.Errorf` kullanarak testimizin başarısız olduğunu ekrana yazdırdık. Eğer testimiz beklediğimiz bir sonuç ise bilgilendirme amaçlı `t.Logf` kullanarak ekrana testimizin başarılı bir şekilde geçtiğini yazdırdık.

Yazdığımız unit testi sonuçları görmek için terminale `go test .` komutunu kullanıyoruz

```
PASS
ok      github.com/bdemirpolat/unit-test        0.466s

```

`go test -v` komutunu kullanarak daha fazla bilgi ile test sonuçlarını ekrana yazdırabilirsiniz.
```
=== RUN   Test_SayHello_Returns_Name_When_Send_Valid_Argument
    greeting_test.go:49: "sayHello(Yemeksepeti)" succeded, expected -> Hello Yemeksepeti!, got -> Hello Yemeksepeti!
--- PASS: Test_SayHello_Returns_Name_When_Send_Valid_Argument (0.00s)
PASS
ok      github.com/bdemirpolat/unit-test        0.466s
```

Şimdi `sayHello` fonksiyonumuzda bir değişiklik yapalım ve testimizin fail olmasını sağlayalım. *Hello* olan yerleri *Hola* olarak değiştirelim ve tekrar testimizi çalıştırıp sonuçları görelim
```
...

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hola Anonymous!"
	}

	return fmt.Sprintf("Hola %s!", name)
}
```

`go test -v` komutunu terminalde çalıştıralım
```
=== RUN   Test_SayHello_Returns_Name_When_Send_Valid_Argument
    greeting_test.go:47: "sayHello(Yemeksepeti)" failed, expected -> Hello Yemeksepeti!, got -> Hola Yemeksepeti!
--- FAIL: Test_SayHello_Returns_Name_When_Send_Valid_Argument (0.00s)
```

# Test Coverage
Uygulama içinde yazdığımız kodun test yüzdesini sağlar. 
Bu sayede yazdığımız kodun hangi taraflarını test etmediğimiz görebiliriz.

# Golang testleri çalıştırma komutları
* go test . -> geçerli dizindeki testleri run eder
* go test ./calc -> calc dizinindeki testleri run eder
* go test ./... modül içindeki tüm testleri run eder

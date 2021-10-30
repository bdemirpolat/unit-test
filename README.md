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

func Test_SayHello_ValidArgument(t *testing.T) {
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
Yukarıda neler yaptık?

TestSayHelloValidArgument adında bir unit testimiz var. Bir dosya içerisinde birden fazla test fonksiyonlarımız olabilir. Test fonksiyonumuz (adından da anlaşılacağı üzere) `sayHello` fonksiyonumuzu test edecek olan fonksiyonumuzdur. `sayHello` fonksiyonuna geçerli bir parametre göndererek çağrıdık ve sonucun beklediğimiz gibi olup olmadığını kontrol ettik. Eğer aldığımız sonuç beklediğimiz bir sonuç değilse `t.Errorf` kullanarak testimizin başarısız olduğunu ekrana yazdırdık. Eğer testimiz beklediğimiz bir sonuç ise bilgilendirme amaçlı `t.Logf` kullanarak ekrana testimizin başarılı bir şekilde geçtiğini yazdırdık.

Yazdığımız unit testi sonuçları görmek için terminale `go test .` komutunu kullanıyoruz

```
PASS
ok      github.com/bdemirpolat/unit-test        0.466s

```

`go test -v` komutunu kullanarak daha fazla bilgi ile test sonuçlarını ekrana yazdırabilirsiniz. (-v flag'i verbose yani ayrıntılı anlamındaıdır)
```
=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:49: "sayHello(Yemeksepeti)" succeded, expected -> Hello Yemeksepeti!, got -> Hello Yemeksepeti!
--- PASS: Test_SayHello_ValidArgument (0.00s)
PASS
ok      github.com/bdemirpolat/unit-test        0.466s
```

Şimdi `sayHello` fonksiyonumuzda bir değişiklik yapalım ve testimizin fail olmasını sağlayalım. *Hello* olan yerleri *Hola* olarak değiştirelim ve tekrar testimizi çalıştırıp sonuçları görelim :)
```
...

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hola Anonymous!"
	}

	return fmt.Sprintf("Hola %s!", name)
}
```

`go test -v` komutunu terminalde çalıştırdığımızda elde edeceğimiz sonuç
```
=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:47: "sayHello(Yemeksepeti)" failed, expected -> Hello Yemeksepeti!, got -> Hola Yemeksepeti!
--- FAIL: Test_SayHello_ValidArgument (0.00s)
```

Bir dosya içerisinde birden fazla test fonksiyonlarımız olabilir demiştik. greeting.go dosyasına `sayGoodBye` diye bir fonksiyon ekleyelim. Bu fonksiyonda aynı `sayHello` fonksiyonu gibi string türünde bir argüman alacak ve geriye string dönen basit bir işlevi olacak.
```
...

func sayGoodBye(name string) string {
	if len(name) == 0 {
		return "Bye Bye Anonymous!"
	}
	return fmt.Sprintf("Bye Bye %s!", name)
}

```

`sayGoodBye` fonksiyonuzun testini **greeting_test.go** dosyasına ekleyelim. 
```
func Test_SayGoodBye(t *testing.T) {
	name := "Yemeksepeti"
	expected := fmt.Sprintf("Bye Bye %s!", name)
	result := sayGoodBye(name)

	if result != expected {
		t.Errorf("\"sayGoodBye(%s)\" failed, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayGoodBye(%s)\" succeded, expected -> %v, got -> %v", name, expected, result)
	}
}
```

**greeting_test.go** dosyasında birden fazla testimiz oldu. Terminalde `go test -v` komutunu çalıştırdığımızda aşağıdaki sonucu elde edeceğiz.

```
=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:49: "sayHello(Yemeksepeti)" succeded, expected -> Hello Yemeksepeti!, got -> Hello Yemeksepeti!
--- PASS: Test_SayHello_ValidArgument (0.00s)
=== RUN   Test_SayGoodBye
    greeting_test.go:61: "sayGoodBye(Yemeksepeti)" succeded, expected -> Bye Bye Yemeksepeti!, got -> Bye Bye Yemeksepeti!
--- PASS: Test_SayGoodBye (0.00s)
PASS
```

`-run` flag ile test dosyaımızda belirli bir fonksiyonu test edebilmek için kullanabiliriz.

`go test -v -run=Test_SayGoodBye` komutunu terminalde çalıştırdığımızda Test_SayGoodBye fonksiyonun sonuçlarını alırız.
```
=== RUN   Test_SayGoodBye
    greeting_test.go:61: "sayGoodBye(Yemeksepeti)" succeded, expected -> Bye Bye Yemeksepeti!, got -> Bye Bye Yemeksepeti!
--- PASS: Test_SayGoodBye (0.00s)
```

# Table-Driven Test Yaklaşımı
*Unit testimizi birden fazla girdi ve girdilerin beklenen sonuçları ile test etmek isteyebiliriz. Test etmek istediğimiz girdiler için array oluşturup bu arrayın her bir elamanı ile testler yapabilmemizi sağlar. Yazdığımız unit testleri farklı kombinasyonlar deneyerek test edebilmemizi sağlayan bir yaklaşımdır.*

Test_SayHello_Valid_Argument fonksiyonunu Table-Driven Test yaklaşımına çevirelim. Input adında bir struct array tanımlıyoruz ve test etmek istediğimiz girdileri ve girdilerden beklenen sonucu array içine alıyoruz. Artık foor-loop iterasyonu ile her bir elemanı test edebiliriz.

```
func Test_SayHello_Valid_Argument(t *testing.T) {
	inputs := []struct {
		name   string
		result string
	}{
		{name: "Yemeksepeti", result: "Hello Yemeksepeti!"},
		{name: "Banabi", result: "Hello Banabi!"},
		{name: "Yemek", result: "Hello Yemek!"},
	}

	for _, item := range inputs {

		result := sayHello(item.name)
		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v", item.name, item.result, result)
		}
	}
}
```

`go test -v -run=Test_SayHello_Valid_Argument` komutunu çalıştırdıktan sonra aşağıdaki sonucu elde edeceğiz.
```
=== RUN   Test_SayHello_Valid_Argument
    greeting_test.go:70: "sayHello('Yemeksepeti')" succeded, expected -> Hello Yemeksepeti!, got -> Hello Yemeksepeti!
    greeting_test.go:70: "sayHello('Banabi')" succeded, expected -> Hello Banabi!, got -> Hello Banabi!
    greeting_test.go:70: "sayHello('Yemek')" succeded, expected -> Hello Yemek!, got -> Hello Yemek!
--- PASS: Test_SayHello_Valid_Argument (0.00s)
PASS
ok      github.com/bdemirpolat/unit-test        0.473s
```

Belirli bir *_test.go dosyasını test edebiliriz. Ancak derleme sırasında test dosyamızın ihtiyacı olan package varsa onu da dahil etmemiz gerekiyor.
`go test -v greeting_test.go` komutu çalıştırdığımızda hata ile karşılaşacağız
```
# command-line-arguments [command-line-arguments.test]
./greeting_test.go:44:12: undefined: sayHello
./greeting_test.go:66:13: undefined: sayHello
./greeting_test.go:78:12: undefined: sayGoodBye
FAIL    command-line-arguments [build failed]
FAIL
```

**greeting_test.go** dosyası test ettiği fonksiyonlar **greeting.go** dosyasında olduğu için `go test -v greeting_test.go` komutuna **greeting.go** dosyasını dahil etmemiz gerekiyor. 
`go test -v greeting_test.go greeting.go` komutu çalıştırdığımızda testler çalışcak ve aşağıdaki sonucu elde edeceğiz.
```
=== RUN   Test_SayHello_Valid_Argument
    greeting_test.go:70: "sayHello({Yemeksepeti Hello Yemeksepeti!})" succeded, expected -> Hello Yemeksepeti!, got -> Hello Yemeksepeti!
    greeting_test.go:70: "sayHello({Banabi Hello Banabi!})" succeded, expected -> Hello Banabi!, got -> Hello Banabi!
    greeting_test.go:70: "sayHello({Yemek Hello Yemek!})" succeded, expected -> Hello Yemek!, got -> Hello Yemek!
--- PASS: Test_SayHello_Valid_Argument (0.00s)
=== RUN   Test_SayGoodBye
    greeting_test.go:83: "sayGoodBye(Yemeksepeti)" succeded, expected -> Bye Bye Yemeksepeti!, got -> Bye Bye Yemeksepeti!
--- PASS: Test_SayGoodBye (0.00s)
PASS
ok      command-line-arguments  0.209s
```

# Test Coverage
Uygulama içinde yazdığımız kodun test yüzdesinin ölçümüdür. Yazdığımız testlerin kodumuzun ne kadarını kapsadığını bilmek önemlidir. Bu sayede kodun hangi taraflarını test ettiğimizi ve hangi taraflarını test etmediğimiz öğrenebiliriz.

Go built-in gelen özelliği sayesinde kod kapsamını kontrol etmemizi sağlıyor.

`go test -cover` komutunu çalıştırıdğımızda aşağıdaki sonucu alacağız.

```
ok      command-line-arguments  0.206s  coverage: 66.7% of statements
```
Testlerimiz başarılı bir şekilde geçti ama yazdığımızın kodun %66.7 si test tarafından ele alınmış. Bu noktada neyi gözden kaçırdığımızı bilmemiz gerekiyor.

Go `-coverprofile` flag'i ile bizlere test kapsam sonuçlarını bir dosyaya aktarmamızı sağlar. Bunun `go test` komutu ile `-coverprofile` flag'ini birleştirmemiz gerekiyor.

`go test -coverprofile=cover_out` komutunu çalıştırıdğımızda geçerli dizinde cover_out dosyası oluşur.

Go projemizin dosya yapısı
* unit-test
     * cover_out (coverage data)
     * greeting_test.go
     * greeting.go
     * main.go

Oluşan dosyayı bizim için daha anlamlı hale getirebiliriz. Go'nun built-in gelen `go tool` komutu ile oluşan *cover_out* dosyasını html formata dönüştürerek web browser üzerinde görüntülememizi sağlar.

`go tool cover -html=cover_out -o cover_out.html` komutunu çalıştırdığımızda geçerli dizinde *cover_out.html* dosyasının oluştuğunu göreceksiniz. 

Go projemizin dosya yapısı
* unit-test
     * cover_out (coverage data)
     * cover_out.html
     * greeting_test.go
     * greeting.go
     * main.go

Oluşan cover_out.html dosyasını herhangi bir web browser ile açtığınızda kodumuzun hangi taraflarını test ettiğimizi ve hangi taraflarını test etmediğimizi görsel bir şekilde görebilirsiniz.

![Image of Cover](https://github.com/bdemirpolat/unit-test/blob/go-test-basics/images/cover_out_result.png)

Kırmızı ile vurgulananlar test tarafından ele alınmadığını anlamına gelir. Yukarıdaki sonuçta gördüğümüz gibi `sayHello` fonksiyonumuza name argümanını boş gönderdiğimizde olan testi yazmadığımızı bize belirtiyor. Bu sayede code coverage eksik taraflarını görüp test coverage sonuçlarımızı artırabilriiz.


# Golang testleri çalıştırma komutları
* go test . -> geçerli dizindeki testleri run eder
* go test ./calc -> calc dizinindeki testleri run eder
* go test ./... modül içindeki tüm testleri run eder

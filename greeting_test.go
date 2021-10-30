package main

import (
	"fmt"
	"testing"
)

/*
	Tricks:
	1. go get -u github.com/rakyll/gotest paketi ile test output'larını renklendirebiliriz
	   kullanımı için go test yerine direkt -> gotest ile başlamalıyız
	2. Belirli testleri çalıştırmak için --run komutunu kullanırız
		ör: go test -v --run Test_SayHello_Returns_Name
	3. Sadece belirli go dosyasını test edebiliriz. Ama ilgili go dosyasının dependency'si varsa
	   onlarıda belirtmemiz gerekir.
	   ör: go test -v greeting_test.go greeting.go
	4. Eğer main pakette (executable main) main.go ile aynı seviyede test dosyaları varsa
	   go run *.go komutu çalışmaz _test.go doslarını da execute etmek isteyecek ve hata olarak cannot run *_test.go files (greeting_test.go) console yazar
	5. Test Cover ile yazdığımız kodun test kapsamını öğrenebiliriz
		-cover flag ile code coverage
		-coverprofile flag ile test cover belirtilen dosyaya aktarabiliriz
		ör: go test -cover
		ör: go test -coverprofile=cover_out
	6. test coverage dosyaya aktardıktan sonra code coverage daha detaylı görmek için HTML çevirebiliriz
		go tool cover komutu oluşan cover dosyasını hangi tür dosyaya çevirmemizi sağlayacak
		ör: go tool cover -html=cover_out -o cover_out.html
		-html flag ile daha önce oluşturduğumuz cover dosyasını html formata dönüştürüyoruz
		-o flag ile html'i çıkaracağımız dosya


*/

// Burada sayHello fonksiyonumuzu test ediyoruz
// sayHello fonksiyonumuz artık programın unit component'i oldu
// test fonksiyonumuz içinde 'Mert' argümanı ile sayHello fonksiyonunu çağrıyoruz
// sonucun beklediğimiz gibi olup olmadını kontrol ediyoruz
// sonuç beklediğimiz gibi değilse t.Errorf ile testen beklentilerimizi yazdırıyoruz
// Uygulamada yazdığımı tüm testleri go test komutunu ile çalıştırıyoruz
// Daha fazla bilgi ile test sonuçlarını yazdırmak için go test -v (verbose)
// Test içinde loglama yapmak için t.Log ya da t.Logf kullanabiliriz
// func Test_SayHello_Returns_Name_When_Send_Valid_Argument(t *testing.T) {
// 	name := "Yemeksepeti"
// 	expected := fmt.Sprintf("Hello %s!", name)
// 	result := sayHello(name)

// 	if result != expected {
// 		t.Errorf("\"sayHello(%s)\" failed, expected -> %v, got -> %v", name, expected, result)
// 	} else {
// 		t.Logf("\"sayHello(%s)\" succeded, expected -> %v, got -> %v", name, expected, result)
// 	}

// }

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

// Fail'e örnek olsun diye aynı testi bu sefer name lengthi 5 ten büyük bir isim girerek deniyoruz
// func TestSayHelloShrinkName(t *testing.T) {
// 	name := "Yemeksepeti"
// 	expected := fmt.Sprintf("Hello %s!", name)
// 	result := sayHelloShrinkName(name)

// 	if result != expected {
// 		t.Errorf("\"sayHello(%s)\" failed, expected -> %v, got -> %v", name, expected, result)
// 	} else {
// 		t.Logf("\"sayHello(%s)\" succeded, expected -> %v, got -> %v", name, expected, result)
// 	}
// }

// func Test_SayHello_Returns_Anonymous_When_Send_Empty_Argument(t *testing.T) {
// 	name := ""
// 	expected := "Hello Anonymous!"
// 	result := sayHello(name)
// 	if result != expected {
// 		t.Errorf("\"sayHello(%s)\" failed, expected -> %v, got -> %v", name, expected, result)
// 	} else {
// 		t.Logf("\"sayHello(%s)\" succeded, expected -> %v, got -> %v", name, expected, result)
// 	}
// }

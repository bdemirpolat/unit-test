package calc

import "testing"

/*
	Working test in Go Module
	Uygulama içinde oluşturduğumuz farklı modülleri test edebiliriz

	1. oluşturduğumuz module test edebilmek için module olduğu dizine gidip go test execute edebiriz
	ya da ana dizinde go test komutu ile test etmek istediğimiz module relative path vermeliyiz
	unit-test
		- calc
			go test -v
	unit-test
		go test ./calc -v
	2. projedeki bütün testleri çalıştırmak için
		go test ./...
		daha önce öğrendiğimiz -v flag ya da -cover flag bu komut ile birleştirebilriiz
		go test ./.. -v -cover
		go test -coverprofile=cover_out -v
	3. projedeki tüm testleri çalıştırıdığımızda golang başarılı testler cache alıyor
		bunu test sonuçlarından da anlayabilriiz
		ok      github.com/bdemirpolat/unit-test/calc   (cached)
		go testleri run ederken ilk başta binary oluşturuyor ve oluşan binary çalıştırıyor
		bu binary dışarı çıkarmak için -c flag kullanabiliriz
			go test ./calc -c


*/

// Örnek test için şimdilik error atladık
func Test_Math_Sum(t *testing.T) {
	expected := 6
	result, _ := Sum(1, 2, 3)
	if result != 6 {
		t.Errorf("\"Sum(1, 2, 3)\" failed, expected -> %v, got -> %v", expected, result)

	} else {
		t.Logf("\"Sum(1, 2, 3)\" succeded, expected -> %v, got -> %v", expected, result)

	}
}

// error handle etme
func Test_Math_Sum_Returns_Error(t *testing.T) {
	result, err := Sum(1)
	if err == nil {
		t.Errorf("\"Sum(1)\" failed, expected an error, got -> %v", result)

	} else {
		t.Logf("\"Sum(1)\" succeded, expected an error -> %v", err)
	}
}

// Bu tür testleri Test Table yani input data ile çözmek en iyi yaklaşım olur
// Bu yaklaşıma TableDriven Test deniyor
func Test_Math_Sum_Input_Data(t *testing.T) {
	inputData := []struct {
		inputs   []int
		result   int
		hasError bool
	}{
		{[]int{1, 2, 3}, 6, false},
		{[]int{100, 200}, 300, false},
		{[]int{1}, 0, true},
	}

	for _, item := range inputData {
		result, err := Sum(item.inputs...)
		if item.hasError {
			if err == nil {
				t.Errorf("\"Sum()\" failed, expected an error, got -> %v", result)

			} else {
				t.Logf("\"Sum()\" succeded, expected an error -> %v", err)
			}
		} else {
			if result != item.result {
				t.Errorf("\"Sum()\" failed, expected %v, got -> %v", item.result, result)
			} else {
				t.Logf("\"Sum()\" succeded, expected %v, got -> %v", item.result, result)
			}
		}
	}
}

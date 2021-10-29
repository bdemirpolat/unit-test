package dob_test

import (
	"testing"
	"time"

	"github.com/bdemirpolat/unit-test/dob"
)

/*
	Seperation of concery
	Aynı package içerisinde yazdığımız testler dışarıya expose olabilir
	Bu durumdan kaçınmalıyız çünkü paketi başkaları kullanmaya başladığında yazdığımız testlerde onlara açık olmuş olacak
	Bunun için package name'lerin sonuna _test suffix ekleyerek ilgi ve alakaları birbirinden ayırmış oluruz

	Test paketlerini ayırdığımızda test ettiğimiz fonksiyon exported olmalı
		yani fonksiyonun adı büyük harfle başlamalı
*/

func Test_DOB_GetRangeOfDob(t *testing.T) {
	expected := "adult"
	adultDob := time.Date(1991, 2, 15, 22, 30, 0, 0, time.UTC)
	result := dob.GetRangeOfDob(adultDob)
	if result == expected {
		t.Logf("PASSED: expected %s, got %s", expected, result)
	} else {
		t.Errorf("FAILED: expected %s, got %s", expected, result)
	}
}

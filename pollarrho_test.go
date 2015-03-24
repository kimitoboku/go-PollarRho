package pollarrho

import (
	"reflect"
	"testing"
)

func TestFactor(t *testing.T) {
	num := 323911866737
	list := Factor(num)
	testlist := []int{5477, 6311, 9371}
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 10
	list = Factor(num)
	testlist = []int{2, 5}
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)

	}
	num = 100
	list = Factor(num)
	testlist = []int{2, 2, 5, 5}
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

}

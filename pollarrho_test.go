package pollarrho

import (
	"reflect"
	"sort"
	"testing"
)

func TestFactor(t *testing.T) {
	num := 323911866737
	list := Factor(num)
	testlist := []int{5477, 6311, 9371}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 10
	list = Factor(num)
	testlist = []int{2, 5}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 100
	list = Factor(num)
	testlist = []int{2, 2, 5, 5}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 4
	list = Factor(num)
	testlist = []int{2, 2}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}
	num = 32
	list = Factor(num)
	testlist = []int{2, 2, 2, 2, 2}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 120
	list = Factor(num)
	testlist = []int{2, 2, 2, 3, 5}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 6887
	list = Factor(num)
	testlist = []int{97, 71}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 9000007
	list = Factor(num)
	testlist = []int{277, 32491}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}
}

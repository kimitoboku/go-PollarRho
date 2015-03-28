package pollarrho

import (
	"reflect"
	"sort"
	"testing"
)

func TestAllPrime(t *testing.T) {
	if !allPrime([]int{5477, 6311, 9371}) {
		t.Error("All Prime function error ", []int{5477, 6311, 9371})
	}
	if !allPrime([]int{2, 5}) {
		t.Error("All Prime function error ", []int{2, 5})
	}
	if allPrime([]int{4, 5, 5}) {
		t.Error("All Prime function error", []int{4, 5, 5})
	}
	if allPrime([]int{4}) {
		t.Error("All Prime function error", []int{4})
	}

}

func TestFactor(t *testing.T) {
	num := 323911866737
	list := PrimeFactor([]int{num})
	testlist := []int{5477, 6311, 9371}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 10
	list = PrimeFactor([]int{num})
	testlist = []int{2, 5}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 100
	list = PrimeFactor([]int{num})
	testlist = []int{2, 2, 5, 5}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}

	num = 4
	list = PrimeFactor([]int{num})
	testlist = []int{2, 2}
	sort.Ints(testlist)
	sort.Ints(list)
	if !reflect.DeepEqual(list, testlist) {
		t.Error("Prime factorization error ", num, " ", list)
	}
}

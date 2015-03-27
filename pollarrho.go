package pollarrho

import (
	"errors"
	"math"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func multMod(a, b, m int) int {
	res := 0
	a = a % m
	for b != 0 {
		if (b & 1) != 0 {
			res = (res + a) % m
		}
		a = (a * 2) % m
		b = b / 2
	}
	return res
}

func PollarRho(num, c int) (int, error) {
	x := 1
	y := 1
	p := 1

	for p == 1 {
		x = (multMod(x, x, num) + c) % num
		ty := (multMod(y, y, num) + c) % num
		y = (multMod(ty, ty, num) + c) % num
		if x == y {
			if math.Sqrt(float64(num)) < float64(c) {
				return -1, errors.New("PollarRho: factor error")
			}
			return PollarRho(num, c+1)
		}

		p = gcd(num, int(math.Abs(float64(x-y))))

	}

	return p, nil
}

func allPrime(num []int) bool {
	for _, i := range num {
		_, err := PollarRho(i, 1)
		if err != nil {
			return false
		}
	}

	return true
}

func Factor(num int) []int {
	var prims []int
	for num != 1 {
		p, err := PollarRho(num, 1)
		if err != nil {
			prims = append(prims, num)
			break
		}
		num = num / p
		prims = append(prims, p)

	}

	return prims
}

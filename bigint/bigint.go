package bigint

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Bigint struct {
	Value string
}

var ErrorBadInput = errors.New("bad input, please input only number")

func NewInt(num string) (Bigint, error) {
	allowed := "1234567890"
	var err bool

	if strings.HasPrefix(num, "-") {
		num = strings.Replace(num, "-", "", 1)
	}
	if strings.HasPrefix(num, "0") {
		err = true
	}
	arr := strings.Split(num, "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}
	}

	if err {
		return Bigint{Value: num}, ErrorBadInput
	} else {
		return Bigint{Value: num}, nil

	}
}

func (z *Bigint) Set(num string) error {
	//
	// Validate input
	for i := 0; i < len(num); i++ {
		if !unicode.IsDigit(rune(num[i])) {
			return errors.New("not valid number")
		}
	}
	//
	z.Value = num
	return nil
}

func Add(a, b Bigint) Bigint {
	//
	// MATH ADD
	carry := 0
	res := ""
	// check length
	if len(a.Value) > len(b.Value) {
		zeros := ""
		diff := len(a.Value) - len(b.Value)
		for i := 0; i < diff; i++ {
			zeros += "0"
		}
		b.Value = zeros + b.Value
	} else if len(a.Value) < len(b.Value) {
		zeros := ""
		diff := len(b.Value) - len(a.Value)
		for i := 0; i < diff; i++ {
			zeros += "0"
		}
		a.Value = zeros + a.Value
	}

	for i := len(a.Value) - 1; i >= 0; i-- {
		x, err := strconv.Atoi(string(a.Value[i]))
		if err != nil {
			return Bigint{}
		}
		y, err := strconv.Atoi(string(b.Value[i]))
		if err != nil {
			return Bigint{}
		}

		z := x + y + carry
		if z > 9 {
			carry = z / 10
			z = z % 10
		} else {
			carry = 0
		}

		curr := strconv.Itoa(z)
		res += curr
	}

	// reverse answer
	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans += string(res[i])
	}

	//
	return Bigint{Value: ans}
}

func Sub(a, b Bigint) Bigint {
	//
	// MATH SUB
	var (
		sign      string
		res       string
		firstNum  string
		secondNum string
	)
	// check length and find first and second num
	if len(a.Value) > len(b.Value) {
		zeros := ""
		diff := len(a.Value) - len(b.Value)
		for i := 0; i < diff; i++ {
			zeros += "0"
		}
		b.Value = zeros + b.Value
		// =======================
		firstNum = a.Value
		secondNum = b.Value
	} else if len(a.Value) < len(b.Value) {
		zeros := ""
		diff := len(b.Value) - len(a.Value)
		for i := 0; i < diff; i++ {
			zeros += "0"
		}
		a.Value = zeros + a.Value
		// =======================
		sign = "-"
		firstNum = b.Value
		secondNum = a.Value
	} else if len(a.Value) == len(b.Value) {
		for i := 0; i < len(a.Value); i++ {
			x, err := strconv.Atoi(string(a.Value[i]))
			if err != nil {
				return Bigint{}
			}
			y, err := strconv.Atoi(string(b.Value[i]))
			if err != nil {
				return Bigint{}
			}

			if x < y {
				sign = "-"
				firstNum = b.Value
				secondNum = a.Value
				break
			}
		}
	}

	// calculate logic
	for i := len(firstNum) - 1; i >= 0; i-- {
		x, err := strconv.Atoi(string(firstNum[i]))
		if err != nil {
			return Bigint{}
		}
		y, err := strconv.Atoi(string(secondNum[i]))
		if err != nil {
			return Bigint{}
		}

		z := x - y
		if z < 0 {
			j := i - 1
			for j != 0 {
				k, err := strconv.Atoi(string(firstNum[j]))
				if err != nil {
					return Bigint{}
				}
				if k == 0 {
					firstNum = firstNum[:j] + "9" + firstNum[j+1:]
				}
				if k > 0 {
					fmt.Println(firstNum[j])
					k = k - 1
					last := strconv.Itoa(k)
					firstNum = firstNum[:j] + last + firstNum[j+1:]
					break
				}
				j--
			}
			z = (x + 10) - y
		}

		curr := strconv.Itoa(z)
		res += curr
	}

	// reverse answer
	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans += string(res[i])
	}

	//
	return Bigint{Value: sign + ans}

}

func Multiply(a, b Bigint) Bigint {
	//
	// MATH Multiply

	res := ""
	// ==================================================
	var dif string

	for i := len(a.Value) - 1; i >= 0; i-- {
		x, err := strconv.Atoi(string(a.Value[i]))
		if err != nil {
			return Bigint{}
		}
		// ==================================================
		var (
			sum string
			k   int
		)

		for j := len(b.Value) - 1; j >= 0; j-- {
			y, err := strconv.Atoi(string(b.Value[j]))
			if err != nil {
				return Bigint{}
			}

			z := (x * y) + k

			if z > 9 {
				k = z / 10
				z = z % 10
			} else {
				k = 0
			}

			sum += strconv.Itoa(z)
		}

		// reverse sum
		sum2 := ""
		for i := len(sum) - 1; i >= 0; i-- {
			sum2 += string(sum[i])
		}

		// add zero to end
		sum2 = sum2 + dif
		dif += "0"

		// add sum of multiply
		curAns := Bigint{}
		curAns = Add(Bigint{
			Value: res,
		}, Bigint{
			Value: sum2,
		})
		res = curAns.Value

	}

	ans := ""
	for i := 0; i < len(res); i++ {
		ans += string(res[i])
	}

	return Bigint{Value: ans}

}

func Mod(a, b Bigint) Bigint {
	//
	// MATH Mod

	// check length
	if len(a.Value) < len(b.Value) {
		return Bigint{Value: "0"}
	}
	if len(a.Value) == len(b.Value) {
		for i := 0; i < len(a.Value); i++ {
			x, err := strconv.Atoi(string(a.Value[i]))
			if err != nil {
				return Bigint{}
			}
			y, err := strconv.Atoi(string(b.Value[i]))
			if err != nil {
				return Bigint{}
			}

			if x < y {
				return Bigint{Value: "0"}
			}
		}
	}

	// calculation logics of mod
	// delitel
	//================================================================================
	// begin values
	y, err := strconv.Atoi(string(b.Value))
	if err != nil {
		return Bigint{}
	}
	i := len(b.Value)
	x, err := strconv.Atoi(string(a.Value[0:len(b.Value)]))
	if err != nil {
		return Bigint{}
	}
	if x < y {
		i = len(b.Value) + 1
		x, err = strconv.Atoi(string(a.Value[0:(len(b.Value) + 1)]))
		if err != nil {
			return Bigint{}
		}
	}
	//================================================================================
	res := ""
	z := 0
	flag := false
	for {
		z = x / y
		if z == 0 {
			break
		}
		curr := strconv.Itoa(z)
		res += curr
		if flag {
			break
		}
		x = x % y
		if x < y {
			cur := strconv.Itoa(x)
			for x < y {
				if i >= len(a.Value) {
					flag = true
					break
				}
				cur += string(a.Value[i])
				i++
				x, err = strconv.Atoi(cur)
				fmt.Println("last x", x)
				if err != nil {
					return Bigint{}
				}
			}

		}

	}
	//================================================================================
	return Bigint{Value: res}
}

func (x *Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	if x.Value[0] == '+' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{
		Value: x.Value,
	}
}

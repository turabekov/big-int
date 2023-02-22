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

	// remove zeros from begining
	byteStr := []byte{}
	for i := 0; i < len(res); i++ {
		byteStr = append(byteStr, res[i])
	}
	for i := 0; i < len(byteStr); i++ {
		if byteStr[i] >= 49 && byteStr[i] <= 57 {
			break
		}
		byteStr[i] = 0
	}
	ans := ""
	for i := 0; i < len(byteStr); i++ {
		ans += string(byteStr[i])
	}

	return Bigint{Value: ans}

}

func Mod(a, b Bigint) Bigint {
	//
	// MATH Mod
	//
	return Bigint{Value: ""}
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

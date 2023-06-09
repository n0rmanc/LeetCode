package LeetCode

import "fmt"

type IntRoman struct {
	Num   int
	Roman string
}

func intToRoman(num int) string {
	romanMap := []IntRoman{
		{
			Num:   1000,
			Roman: "M",
		},
		{
			Num:   900,
			Roman: "CM",
		},
		{
			Num:   500,
			Roman: "D",
		},
		{
			Num:   400,
			Roman: "CD",
		},
		{
			Num:   100,
			Roman: "C",
		},
		{
			Num:   90,
			Roman: "XC",
		},
		{
			Num:   50,
			Roman: "L",
		},
		{
			Num:   40,
			Roman: "XL",
		},
		{
			Num:   10,
			Roman: "X",
		},
		{
			Num:   9,
			Roman: "IX",
		},
		{
			Num:   5,
			Roman: "V",
		},
		{
			Num:   4,
			Roman: "IV",
		},
		{
			Num:   1,
			Roman: "I",
		},
	}
	for _, v := range romanMap {
		fmt.Println(num, v)
		if num >= v.Num {
			num -= v.Num
			return v.Roman + intToRoman(num)
		}
	}
	return ""
}

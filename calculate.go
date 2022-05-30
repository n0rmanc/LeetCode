package LeetCode

import "fmt"

func calculate(s string) int {
	var operations []int
	var operators []int
	for len(s) > 0 {
		var cur int
		cur, s = getNext(s)
		// * = 2a
		// / = 2f
		// + = 2b
		// - = 2d
		// 0~9 = 30~39
		// " " = 20
		switch cur {
		case 0x2a: // *
			var a int
			a, operations = pop(operations)
			var next int
			next, s = getNext(s)
			c := a * (next - 48)
			operations = push(operations, c)
		case 0x2f: // /
			var a int
			a, operations = pop(operations)
			var next int
			next, s = getNext(s)
			c := a / (next - 48)
			operations = push(operations, c)
		case 0x2b: // +
			operators = push(operators, cur)
		case 0x2d: // -
			operators = push(operators, cur)
		case 0x20:
			continue
		default:
			operations = push(operations, cur-48)
		}
	}
	fmt.Println(operations)
	for _, operator := range operators {
		switch operator {
		case 0x2b: // +
			a := operations[0]
			b := operations[1]
			c := a + b
			operations = append([]int{c}, operations[2:]...)
		case 0x2d: // -
			a := operations[0]
			b := operations[1]
			c := a - b
			operations = append([]int{c}, operations[2:]...)
		}

	}
	return operations[0]
}

func push(a []int, o int) []int {
	return append(a, o)
}

func pop(a []int) (int, []int) {
	return a[len(a)-1], a[:len(a)-1]
}

func popS(a string) (int, string) {
	return int(a[0]), a[1:]
}

func getNext(a string) (int, string) {
	var b int
	for ok := true; ok; ok = b == 0x20 && len(a) > 0 {
		b, a = popS(a)
	}

	if b >= 0x30 && b <= 0x39 {

		count := b - 0x30
		for len(a) > 0 && a[0] >= 0x30 && a[0] <= 0x39 {
			b, a = popS(a)
			count = count*10 + (b - 0x30)
		}
		b = count + 48
	}
	return b, a
}

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func generateNick(firstName, lastName string) string {
	nick := strings.ToLower(firstName[:3] + lastName[:3])
	return strings.Map(func(r rune) rune {
		if r == 'Ą' {
			return 'A'
		}
		if r == 'Ł' {
			return 'L'
		}
		if r == 'Ę' {
			return 'E'
		}
		if r == 'Ż' {
			return 'Z'
		}
		if r == 'Ź' {
			return 'Z'
		}
		if r == 'Ń' {
			return 'N'
		}
		return r
	}, nick)
}

func asciiValues(nick string) []int {
	ascii := make([]int, len(nick))
	for i, char := range nick {
		ascii[i] = int(char)
	}
	return ascii
}

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	fact := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	return fact
}

var fibCalls = make(map[int]int)

func fib(n int) int {
	if n <= 1 {
		fibCalls[n]++
		return n
	}
	fibCalls[n]++
	return fib(n-1) + fib(n-2)
}

func main() {
	var firstName, lastName string

	fmt.Print("Podaj imię: ")
	fmt.Scanln(&firstName)

	fmt.Print("Podaj nazwisko: ")
	fmt.Scanln(&lastName)

	nick := generateNick(firstName, lastName)
	fmt.Println("Nick:", nick)

	ascii := asciiValues(nick)
	fmt.Println("ASCII Values:", ascii)

	var strongNum *big.Int
	n := int64(0)
	for {
		strongNum = factorial(n)
		found := true
		for _, val := range ascii {
			if !strings.Contains(strongNum.String(), strconv.Itoa(val)) {
				found = false
				break
			}
		}
		if found {
			break
		}
		n++
	}

	fmt.Println("Strong Number:", n)

	fib30 := fib(30)
	fmt.Println("Fibonacci(30):", fib30)

	// fmt.Println("Fibonacci Calls:")
	// for arg, calls := range fibCalls {
	// 	fmt.Println("Fibonacci(", arg, "):", calls, "calls")
	// }

	weakNum := findWeakNumber(int(n), fibCalls)
	fmt.Println("Weak Number:", weakNum, "Calls:", fibCalls[weakNum])
}

func findWeakNumber(strongNum int, fibCalls map[int]int) int {
	minDiff := int(strongNum)
	weakNum := 0

	for n, calls := range fibCalls {
		diff := abs(strongNum - calls)
		if diff < minDiff {
			minDiff = diff
			weakNum = n
		}
	}

	return weakNum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

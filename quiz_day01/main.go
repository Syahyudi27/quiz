package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("1. Find Divisor")
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)

	fmt.Println("2. Extract Digit")
	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)

	fmt.Println("3. Triangle Left")
	triangleLeft()

	fmt.Println("Triangle Right")
	triangleRight(5)

	fmt.Println("4. Pyramid")
	piramid(8)

	fmt.Println("5. Angka Deret 1")
	angkaDeret(5)

	fmt.Println("6. Angka Deret 2")
	angkaDeret2(5)

	fmt.Println("8. Is Palindrome (String)")
	isPalindrome("Kasur ini rusak")
	isPalindrome("tamaT")
	isPalindrome("Aku Usa")

	fmt.Println("9. Reverse")
	reverse("ABCD")
	reverse("tamaT")
	reverse("XYnb")

	fmt.Println("10. Check Braces")
	fmt.Println(checkBraces("(()))"))
	fmt.Println(checkBraces("(())"))
	fmt.Println(checkBraces("((()"))
	fmt.Println(checkBraces("(()))((())"))

	fmt.Println("11. Is Palindrome (Number)")
	fmt.Println(isNumberPalindrome(121))
	fmt.Println(isNumberPalindrome(123))
}

// 1
func findDivisor(n int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// 2
func extractDigit(n int) {
	for n > 0 {
		digit := n % 10
		fmt.Print(digit, " ")
		n = n / 10
	}
	fmt.Println()
}

// 3
func triangleLeft() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func triangleRight(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= (n-1)-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 4
func piramid(n int) {
	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		for k := 2; k <= i; k++ {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}
}

// 5
func angkaDeret(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 != 0 {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d ", n-i+1)
			}
		}
		fmt.Println()
	}
}

// 6
func angkaDeret2(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (i+j)%2 != 0 {
				fmt.Printf("%d ", j)
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

// 8
func isPalindrome(kata string) {
	cleanKata := strings.ToLower(kata)
	n := len(cleanKata)
	status := true
	for i := 0; i < n/2; i++ {
		if cleanKata[i] != cleanKata[n-1-i] {
			status = false
			break
		}
	}
	fmt.Println(status)
}

// 9
func reverse(kata string) {
	n := len(kata)
	for i := n - 1; i >= 0; i-- {
		fmt.Print(string(kata[i]))
	}
	fmt.Println()
}

// 10
func checkBraces(s string) bool {
	count := 0
	for _, char := range s {
		switch char {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// 11
func isNumberPalindrome(n int) bool {
	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	return original == reversed
}
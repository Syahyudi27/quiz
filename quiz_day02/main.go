package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := []string{"code", "java", "cool"}
	hasil1 := upperCaseExcept(input1, "java")
	fmt.Println("1.", hasil1) 

	input2 := []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	fmt.Println("2.",findMinMax(input2))

	dataRange := []int{5, 3, 4, 2, 6, 7, 8, 9, 1, 10}
	fmt.Println("3.",findMinRange(dataRange, 0, 10))
	fmt.Println("3.",findMinRange(dataRange, 0, 7))

	dataMax := []int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49}
	fmt.Println("3.",findMaxRange(dataMax, 0, 10))
	fmt.Println("3.",findMaxRange(dataMax, 2, 7))

	dataAcak := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("4.",evenOddOrder(dataAcak))

	arr := []int{12, 15, 1, 5, 20}
	fmt.Println("5.",rotateArray(arr, 1))
	arr2 := []int{12, 15, 1, 5, 20}
	fmt.Println(rotateArray(arr2, 2))
	arr3 := []int{12, 15, 1, 5, 20}
	fmt.Println(rotateArray(arr3, 3))
	fmt.Println()

	drawMatrixPatternRed()
	fmt.Println()

	drawMatrixPatternRed()
	fmt.Println()

	drawMatrixSum()

	kunciJawaban := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}
	dataSiswa := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}
	calculateScores(dataSiswa, kunciJawaban)
	fmt.Println()
}

func upperCaseExcept(arr []string, exception string) []string {
	var result []string
	for _, word := range arr {
		if word == exception {
			result = append(result, word)
		} else {
			upperWord := strings.ToUpper(word)
			result = append(result, upperWord)
		}
	}
	return result
}

func findMinMax(arr []int) []int{
	min := arr[0]
	max := arr[0]
	
	for _, number := range arr {
		if number < min{
			min = number
		}
		
		if number > max {
			max = number
		}
	}

	return []int{min, max}
}


func findMinRange(arr []int, startIndex int, endIndex int) []int {
	minValue := arr[startIndex]
	minIndex := startIndex

	for i := startIndex; i < endIndex; i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}

	return []int{minValue, minIndex}
}

func findMaxRange(arr []int, startIndex int, endIndex int) []int {
	maxValue := arr[startIndex]
	maxIndex := startIndex

	for i := startIndex; i < endIndex; i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
			maxIndex = i
		}
	}

	return []int{maxValue, maxIndex}
}

func evenOddOrder(arr []int) []int {
	var genap []int
	var ganjil []int

	for _, number := range arr {
		if number%2 == 0 {
			genap = append(genap, number)
		} else {
			ganjil = append(ganjil, number)
		}
	}

	return append(genap, ganjil...)
}

func rotateArray(arr []int, n int) []int {
	if len(arr) == 0 {
		return arr
	}
	n = n % len(arr)
	return append(arr[n:], arr[:n]...)
}


func calculateScores(studentAnswers [][]string, key []string) {
	for i, answers := range studentAnswers {
		score := 0
		for j, ans := range answers {
			if ans == key[j] {
				score++
			}
		}
		fmt.Printf("Jawaban Siswa %d yang benar : %d\n", i, score)
	}
}

func drawMatrixPatternRed() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Printf("%2d ", i+1)
			} else if j > i {
				fmt.Printf("%2d ", 10)
			} else {
				fmt.Printf("%2d ", 20)
			}
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Printf("%2d ", n-i)
			} else if j > i {
				fmt.Printf("%2d ", 20)
			} else {
				fmt.Printf("%2d ", 10)
			}
		}
		fmt.Println()
	}
}

func drawMatrixSum() {
	n := 7
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := i + j
			matrix[i][j] = val
			matrix[i][n] += val
			matrix[n][j] += val
			matrix[n][n] += val
		}
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}
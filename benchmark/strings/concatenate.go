package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func concatenateByBytes(first string, second string) (result string, timePass time.Duration) {
	start := time.Now()
	var resultBuffer bytes.Buffer
	firstBytes := []byte(first)
	secondBytes := []byte(second)

	for _, b := range firstBytes {
		resultBuffer.WriteString(string(b))
	}
	for _, b := range secondBytes {
		resultBuffer.WriteString(string(b))
	}
	result = resultBuffer.String()
	end := time.Now()
	timePass = end.Sub(start)
	return
}

func concatenateByPlus(first string, second string) (result string, timePass time.Duration) {
	start := time.Now()
	result = first + second
	end := time.Now()
	timePass = end.Sub(start)
	return
}

func concatenateByCopy(first string, second string) (result string, timePass time.Duration) {
	start := time.Now()
	totalLength := len(first) + len(second)
	now := 0
	resultBytes := make([]byte, totalLength)

	for _, s := range first {
		now += copy(resultBytes[now:], string(s))
	}
	for _, s := range second {
		now += copy(resultBytes[now:], string(s))
	}
	result = string(resultBytes)
	end := time.Now()
	timePass = end.Sub(start)
	return
}

func concatenateByStringBuilder(first string, second string) (result string, timePass time.Duration) {
	start := time.Now()
	var resultBuilder strings.Builder

	for _, s := range first {
		resultBuilder.WriteString(string(s))
	}
	for _, s := range second {
		resultBuilder.WriteString(string(s))
	}
	result = resultBuilder.String()
	end := time.Now()
	timePass = end.Sub(start)
	return
}

func main() {
	first := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco `
	second := `laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	result := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	fmt.Printf("Test string length: %d\n", len(result))
	ans, timePass := concatenateByBytes(first, second)

	if ans != result {
		panic(`Cannot concatenate two strings`)
	} else {
		fmt.Printf("Success to concatenate two strings by bytes, time pass: %v\n", timePass)
	}

	ans, timePass = concatenateByPlus(first, second)

	if ans != result {
		panic(`Cannot concatenate two strings`)
	} else {
		fmt.Printf("Success to concatenate two strings by plus, time pass: %v\n", timePass)
	}

	ans, timePass = concatenateByCopy(first, second)

	if ans != result {
		panic(`Cannot concatenate two strings`)
	} else {
		fmt.Printf("Success to concatenate two strings by copy, time pass: %v\n", timePass)
	}

	ans, timePass = concatenateByStringBuilder(first, second)

	if ans != result {
		panic(`Cannot concatenate two strings`)
	} else {
		fmt.Printf("Success to concatenate two strings by string builder, time pass: %v\n", timePass)
	}
}

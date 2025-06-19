package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func DeferExample() {
	fmt.Println("0")
	funcWithDefer()
	fmt.Println("4")
}

func funcWithDefer() (int, error) {
	fmt.Println("1")
	defer fmt.Println("3")
	return fmt.Println("2")
}

func printFile() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

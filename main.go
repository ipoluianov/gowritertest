package main

import (
	"fmt"

	"github.com/ipoluianov/gowritertest/tester"
)

func main() {
	go tester.Run("data_001")
	go tester.Run("data_002")
	go tester.Run("data_003")
	go tester.Run("data_004")
	go tester.Run("data_005")

	go tester.Run("data_006")
	go tester.Run("data_007")
	go tester.Run("data_008")
	go tester.Run("data_009")
	go tester.Run("data_010")

	fmt.Println("Started")
	fmt.Scanln()
}

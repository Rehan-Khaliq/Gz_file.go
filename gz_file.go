package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Create("GZ_file.csv.gz")

	if err != nil {
		fmt.Println("Error Creating File", err)

	}
	defer file.Close()
	gzwriter := gzip.NewWriter(file)
	defer gzwriter.Close()

	cswriter := csv.NewWriter(gzwriter)
	defer cswriter.Flush()

	// By Applying Random Number Generator
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ { // corrected syntax errors in the loop
		randomnumer := generateRandomNumber()
		if err := cswriter.Write([]string{randomnumer}); err != nil {
			fmt.Println("Error Writing To CSV:", err)

		}
	}
	fmt.Println("1 Million Data written successfully") // corrected spelling error
}

func generateRandomNumber() string {
	firstDigit := rand.Intn(8) + 2
	remainingDigit := rand.Int63n(900000000) + 100000000 // 9 remaining digits
	return strconv.Itoa(firstDigit) + strconv.FormatInt(remainingDigit, 10)
}

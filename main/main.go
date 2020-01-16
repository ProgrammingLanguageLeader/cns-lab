package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := parseArgs()
	fmt.Println("Переданные агрументы:")
	fmt.Println(args.ToString())

	outputFile, err := os.Create(args.Output)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	trafficParams := GenerateTrafficParams(args.Variant, args.StudentsNumber)
	for sampleNumber := uint(0); sampleNumber < args.SamplesNumber; sampleNumber++ {
		uintArray := trafficParams.ToArray()
		stringArray := make([]string, len(uintArray))
		for paramIndex, paramValue := range uintArray {
			stringArray[paramIndex] = strconv.FormatUint(uint64(paramValue), 10)
		}
		if err := writer.Write(stringArray); err != nil {
			log.Fatal("Cannot write to file", err)
		}
		trafficParams.GenerateNext()
	}
}

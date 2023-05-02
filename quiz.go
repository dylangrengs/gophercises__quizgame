package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CSVDataLines struct {
	Column1 string
	Column2 string
}

func main() {
	// create a reader for user i/o
	reader := bufio.NewReader(os.Stdin)

	// declare score counter
	var score int = 0

	// read csv into csvData
	csvData, err := readCSV("problems.csv")
	if err != nil {
		// panic if the file does not exist
		panic(err)
	}

	for _, line := range csvData {
		data := CSVDataLines{
			Column1: line[0],
			Column2: line[1],
		}
		// print out the question
		fmt.Println("Question:" + " " + data.Column1 + "?")
		// read in user answer
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		// check for correct answer
		if strings.Compare(text, data.Column2) == 0 {
			score++
		}

	}
	// print user score
	var response string = "Your Score: "
	result := response + strconv.Itoa(score)
	fmt.Println(result)
}

func readCSV(filename string) ([][]string, error) {
	// this opens the csv file
	content, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	// defer moves the execution until the end of the function
	defer content.Close()

	// read file into a variable
	lines, err := csv.NewReader(content).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

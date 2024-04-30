package fileio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Cutf(filename string, fieldList string, delimiter string) {

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	printContent(*scanner, fieldList, delimiter)

}

func CutfFromStdin(fieldList, delimiter string) {
	scanner := bufio.NewScanner(os.Stdin)

	printContent(*scanner, fieldList, delimiter)
}

func getFieldNum(fieldNumStr string) int {
	var fieldNum int
	_, err := fmt.Sscanf(fieldNumStr, "%d", &fieldNum)
	if err != nil {
		fmt.Println("Invalid field number:", fieldNumStr)
		os.Exit(1)
	}
	return fieldNum
}

func printContent(scanner bufio.Scanner, fieldList string, delimiter string) {

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, delimiter)
		output := []string{}

		for _, fieldNumStr := range strings.FieldsFunc(fieldList, func(r rune) bool { return r == ',' || r == ' ' }) {
			fieldNum := getFieldNum(fieldNumStr)
			if fieldNum >= 1 && fieldNum <= len(fields) {
				output = append(output, fields[fieldNum-1])
			}
		}
		fmt.Println(strings.Join(output, delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)

	}

}

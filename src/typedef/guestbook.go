package typedef

import (
	"bufio"
	//"fmt"
	"log"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures []string
}

func GetStrings(fileName string) []string{
	var lines []string
	//pwd, _ := os.Getwd()
	//fmt.Println("pwd :", pwd)
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
}

func Check(err error) {
	if err!=nil {
		log.Fatal(err)
	}
}

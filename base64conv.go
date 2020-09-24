package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {

	input := flag.String("input", "", "input parameter")
	output := flag.String("output", "", "output parameter")
	flag.Parse()

	base64Conv(*input, *output)
}

func base64Conv(input string, output string) {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	if isBase64(string(data)) {
		log.Println("BASE64")
		base64Decode(input, output)
	} else if !isBase64(string(data)) {
		base64Encode(input, output)
	} else {
		log.Fatalln("base64 -input base64.txt -output tmp.pdf")
	}
}

func base64Decode(input string, output string) {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(output, decodeBytes, 0666)
	if err != nil {
		log.Fatalln(err)
	}
}

func base64Encode(input string, output string) {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	encodeString := base64.StdEncoding.EncodeToString(data)
	err = ioutil.WriteFile(output, []byte(encodeString), 0666)
	if err != nil {
		log.Fatalln(err)
	}
}

func isBase64(str string) bool {
	pattern := "^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$"
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	if !(len(str)%4 == 0 && matched) {
		return false
	}
	return true
}

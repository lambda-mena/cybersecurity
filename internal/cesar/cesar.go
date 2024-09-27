package cesar

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RequestInputs() {
	var fileName string = os.Args[2]
	var stringkey string = os.Args[3]
	key, err := strconv.Atoi(stringkey)

	if err != nil {
		fmt.Println("error parsing key: key is not a integer")
		return
	}

	fileContent, err := os.ReadFile(fileName)

	var textBuild strings.Builder
	for _, byte := range fileContent {
		textBuild.WriteByte(byte)
	}
	rawText := textBuild.String()
	parsedText := strings.ToUpper(rawText)

	if err != nil {
		fmt.Println("error reading file")
		return
	}

	//fmt.Println("Text in uppercase:", textInUpper)
	decypheredText := Decrypt(parsedText, key)

	os.WriteFile("out/decyphered.txt", []byte(decypheredText), 0644)

	fmt.Println(decypheredText)
}

func Encrypt(fileContent string, key int) string {
	var encrypted strings.Builder

	key8 := byte(key%26+26) % 26

	for _, charByte := range []byte(fileContent) {
		if 'A' <= charByte && charByte <= 'Z' {
			encrypted.WriteByte('A' + (charByte-'A'+key8)%26)
		} else {
			encrypted.WriteByte(charByte)
		}
	}

	return encrypted.String()
}

func Decrypt(fileContent string, key int) string {
	return Encrypt(fileContent, 26-key)
}

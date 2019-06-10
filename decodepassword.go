package genpass

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func DecodePassword() error {

	fmt.Println("Insert a title")
	var titlePass string
	fmt.Scanf("%s\n", &titlePass)

	// Open the file csv
	filecsv, _ := os.OpenFile("../pass/pass.csv", os.O_RDWR, 0777)
	defer filecsv.Close()

	// Read all the file csv
	scanner := bufio.NewScanner(filecsv)

	// Read line for line
	for scanner.Scan() {
		line := scanner.Text()
		newline := strings.Split(line, ",")
		// Check the title and take the right decoded password
		if newline[0] == titlePass {
			decoded, _ := base64.URLEncoding.DecodeString(newline[1])
			fmt.Println(string(decoded))
			break
		}
	}

	return nil
}

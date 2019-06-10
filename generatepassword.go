package genpass

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Add folder for the csv password file
func GenerateDir() error {
	// If dosen't exist a dir, make it.
	if _, err := os.Stat("../pass"); os.IsNotExist(err) {
		// create a dir
		os.Mkdir("../pass", os.FileMode(0755))
	}

	return nil
}

// Create the csv password file
func GenerateFileCSV() error {
	// If dosen't exist a csv file, make it.
	if _, err := os.Stat("../pass/pass.csv"); os.IsNotExist(err) {
		// go to the /pass directory
		if err := os.Chdir("../pass"); err != nil {
			return err
		}

		// Create csv file
		filecsv, err := os.Create("pass.csv")
		if err != nil {
			return err
		}
		defer filecsv.Close()
	}

	return nil
}

// Add title, base64 password and date, at the last line of the csv file"
func WriteCSV(title string, n int) error {
	// Open csv file in create mode and append mode
	file, err := os.OpenFile("../pass/pass.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	// Add title, base64 password and current time in data
	data := []string{title, GeneratePassBase64(n), time.Now().String()}
	// Write in csv file
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{data}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush() // https://golang.org/pkg/bufio/#Writer

	return nil
}

// All chars for generate our password
var basicchar = []rune("ABCDEFGHILMNOPQRSTUVZabcdefghilmnopqrstuvz0123456789><;:_ç°§é*^?=)(/!&%$$£")

func RandomPassword(n int) string {
	// slice in r the new rune
	r := make([]rune, n)
	for i := range r {
		//Randomize r
		r[i] = basicchar[rand.Intn(len(basicchar))]
	}
	return string(r)
}

// Only for the password, use the time for the Seed of the random and encode in base64
func GeneratePassBase64(n int) string {
	rand.Seed(time.Now().UnixNano())
	value := base64.URLEncoding.EncodeToString([]byte(RandomPassword(n)))
	return value
}

// Header of the generator password
func GeneratePassword() error {

	fmt.Println("How much longer your password?")
	var longPassword int
	fmt.Scanf("%d\n", &longPassword)

	fmt.Println("A title for your password?")
	var titlePassword string
	fmt.Scanf("%s\n", &titlePassword)

	if err := GenerateDir(); err != nil {
		return err
	}

	if err := GenerateFileCSV(); err != nil {
		return err
	}

	if err := WriteCSV(titlePassword, longPassword); err != nil {
		return err
	}

	return nil
}

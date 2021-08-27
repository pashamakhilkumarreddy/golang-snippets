package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	log.Println("Reading files in Golang")
	filePath := flag.String("filepath", "values.json", "location of the file to be read")
	flag.Parse()
	data, err := ioutil.ReadFile(*filePath) // ioutil.ReadFile("/Users/apasham/Downloads/golang-snippets/reading_files/values.json")

	if err != nil {
		fmt.Println("Error Reading file", err)
		return
	}
	_ = data
	// fmt.Println("Contents of the file are: ", string(data))

	box := packr.New("fileBox", "./assets")
	fileData, err := box.FindString("values.json")
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}
	fmt.Println("contents of the file are", fileData)

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	func() {
		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		err = s.Err()
		if err != nil {
			fmt.Println(err)
			// log.Fatal(err)
		}
		defer fmt.Println("done reading from scan")
	}()

	func() {
		r := bufio.NewReader(f)

		b := make([]byte, 3)
		for {
			n, err := r.Read(b)
			if err != nil {
				fmt.Println("error reading file", err)
				break
			}
			fmt.Println(string(b[:n]))
		}
		defer fmt.Println("done reading from chunks")
	}()

}

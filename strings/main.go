package main

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Printf("\nBytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("\nCharacters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printRunes(s string) {
	fmt.Printf("\nCharacters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytesPosition(s string) {
	for i, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, i)
	}
}

func mutate(s []rune) string {
	s[0] = ' '
	return string(s)
}

func compareStrings(s1, s2 string) bool {
	if s1 == s2 {
		fmt.Printf("Strings %s and %s are equal\n", s1, s2)
		return true
	}
	fmt.Printf("Strings %s and %s are not equal\n", s1, s2)
	return false
}

func main() {
	log.Println("Strings in Golang")
	name := "Hola Mundo"
	nombre := strings.Join([]string{"Hola", "Señor"}, " ")
	fmt.Println(name)
	printBytes(name)
	printChars(name)
	printChars(nombre)
	printRunes(nombre)
	charsAndBytesPosition(nombre)

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	fmt.Println(string(runeSlice))

	fmt.Println(len(nombre))
	fmt.Println(utf8.RuneCountInString(nombre))

	_ = compareStrings(name, nombre)

	fmt.Println("Go " + "is awesome")
	result := fmt.Sprintf("Me llamo es %s", "Marco")
	fmt.Println(result)

	fmt.Println(mutate([]rune(name)))
}

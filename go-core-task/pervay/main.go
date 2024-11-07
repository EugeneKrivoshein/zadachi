package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

var (
	numDecimal     int       = 42       // Десятичная система
	numOctal       int       = 052      // Восьмеричная система
	numHexadecimal int       = 0x2A     // Шестнадцатиричная система
	pi             float64   = 3.14     // Тип float64
	name           string    = "Golang" // Тип string
	isActive       bool      = true     // Тип bool
	complexNum     complex64 = 1 + 2i   // Тип complex64
)

func printType() {
	fmt.Printf("numDecimal type: %s\n", reflect.TypeOf(numDecimal))
	fmt.Printf("numOctal type: %s\n", reflect.TypeOf(numOctal))
	fmt.Printf("numHexadecimal type: %s\n", reflect.TypeOf(numHexadecimal))
	fmt.Printf("pi type: %s\n", reflect.TypeOf(pi))
	fmt.Printf("name type: %s\n", reflect.TypeOf(name))
	fmt.Printf("isActive type: %s\n", reflect.TypeOf(isActive))
	fmt.Printf("complexNum type: %s\n", reflect.TypeOf(complexNum))
}

func ToStringCombined() string {
	return fmt.Sprintf("%d%d%d%f%s%t%g", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
}

func HashRuneWithSalt(str string, salt string) string {
	combinedRunes := []rune(str)
	mid := len(combinedRunes) / 2
	combinedRunes = append(combinedRunes[:mid], append([]rune(salt), combinedRunes[mid:]...)...)

	hash := sha256.Sum256([]byte(string(combinedRunes)))
	return hex.EncodeToString(hash[:])
}

func main() {
	//определение типов и вывод на экран
	printType()

	//преобразует все переменные в строковый тип и объединяет в одну строку
	combinedString := ToStringCombined()
	fmt.Println("combined string:", combinedString)

	//преобразование строки в срез рун, хеширование среза рун sha256
	//с добавление соли `go-2024` в середину и вывод на экран
	salt := "go-2024"
	hashedResult := HashRuneWithSalt(combinedString, salt)
	fmt.Println("Hashed with Salt:", hashedResult)
}

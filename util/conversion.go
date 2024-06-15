package util

import (
	"encoding/base64"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
)

func StringToInt64(s string) int64 {

	var i int64

	if strings.ContainsAny(s, ".,") {
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", "")
	}

	if strings.ContainsAny(s, "ABDCEFGHIJKLMNÑOPQRSTUVWXYZÇabcdefghijklmnñopqrstuvwxyzç!\"·$%^&*()#~€¬|@¡¿?¿¡") {
		return 0
	}

	for _, c := range s {
		i = i*10 + int64(c-'0')
	}
	return i
}

func StringToInt(s string) (int, error) {
	if strings.ContainsAny(s, "\"≠”´!@#$%^&*()_+-=[]{};':,./<>?çæ·") {
		return 0, fmt.Errorf("invalid number, contains special characters")
	}
	if strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyzñABCDEFGHIJKLMNOPQRSTUVWXYZÑ") {
		return 0, fmt.Errorf("invalid number, contains letters")
	}
	var number int

	_, err := fmt.Sscan(s, &number)

	if err != nil {
		return 0, err
	}
	return number, err
}

// format date: 17 Junio, 2023 16:13
func FormatDateTime(date time.Time) string {
	return fmt.Sprintf("%s %s, %d %s:%s", addZero(date.Day()), monthSpnaish(date.Month()), date.Year(), addZero(date.Hour()), addZero(date.Minute()))
}

// format date: 17 de Junio, 2023
func FormatDateWithYear(date time.Time) string {
	return fmt.Sprintf("%s de %s, %d", addZero(date.Day()), monthSpnaish(date.Month()), date.Year())
}

// format date: 17 de Junio
func FormatDateWithoutYear(date time.Time) string {
	return fmt.Sprintf("%s de %s", addZero(date.Day()), monthSpnaish(date.Month()))
}

func monthSpnaish(m time.Month) string {
	switch m {
	case time.January:
		return "Enero"
	case time.February:
		return "Febrero"
	case time.March:
		return "Marzo"
	case time.April:
		return "Abril"
	case time.May:
		return "Mayo"
	case time.June:
		return "Junio"
	case time.July:
		return "Julio"
	case time.August:
		return "Agosto"
	case time.September:
		return "Septiembre"
	case time.October:
		return "Octubre"
	case time.November:
		return "Noviembre"
	case time.December:
		return "Diciembre"
	}
	return ""
}

func addZero(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprintf("%d", n)
}

func FisrtIdentifier(id uuid.UUID) string {
	return identifier(id, 0)
}

func SecondIdentifier(id uuid.UUID) string {
	return identifier(id, 1)
}

func ThirdIdentifier(id uuid.UUID) string {
	return identifier(id, 2)
}

func FourthIdentifier(id uuid.UUID) string {
	return identifier(id, 3)
}

func FifthIdentifier(id uuid.UUID) string {
	return identifier(id, 4)
}

func identifier(id uuid.UUID, index uint) string {
	s := strings.Split(id.String(), "-")
	return s[index]
}

func FileToBase64(file *multipart.FileHeader, encoded chan<- string) {

	pdf, err := file.Open()

	if err != nil {
		log.Printf("Error opening file: %s", err)
		encoded <- ""
	}

	defer pdf.Close()

	buffer := make([]byte, file.Size)

	if _, err = pdf.Read(buffer); err != nil {
		log.Printf("Error reading file: %s", err)
		encoded <- ""
	}

	defer close(encoded)

	encoded <- base64.StdEncoding.EncodeToString(buffer)
}

func StringBase64ToBytes(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		log.Printf("Error decoding base64 string: %s", err)
		return []byte{}, err
	}

	return data, nil
}

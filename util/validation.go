package util

import (
	"fmt"
	"math/rand"
	"net/mail"
	"strings"
	"time"
)

const (
	leters       string = "ABCDEFGHIJKLMNOPQRSTUVWXYZÑÇabcdefghijklmnopqrstuvwxyzñç"
	symbols      string = "\\|\"≠”¨`´!@$%^*()+=[]{};',.<>?çæ·"
	symbolsDate  string = symbols + "-:#"
	symbolsHour  string = symbols + "-/#"
	symbolsOrder string = symbols + "-:/"
	symbolsSKU   string = symbols + ":/#"
	allSymbols   string = symbols + "-:/#"
	numbers      string = "0123456789"
)

func ValidAddress(email string) bool {
	if strings.Contains(email, " ") {
		return false
	}

	if !strings.Contains(email, ".") {
		return false
	}

	if !strings.Contains(email, "@") {
		return false
	}

	_, err := mail.ParseAddress(email)

	return err == nil
}

func ValidCompleteName(name string) bool {
	if !strings.Contains(name, " ") {
		return false
	}
	if strings.Contains(name, "  ") {
		return false
	}
	if strings.ContainsAny(name, allSymbols) {
		return false
	}
	if strings.ContainsAny(name, numbers) {
		return false
	}
	return (len(name) >= 7 && len(name) <= 80)
}

func ValidPhone(phone string) bool {
	if strings.Contains(phone, " ") {
		return false
	}
	if !strings.ContainsAny(phone, numbers) {
		return false
	}
	if strings.ContainsAny(phone, allSymbols) {
		return false
	}
	if _, err := StringToInt(phone); err != nil {
		return false
	}
	return (len(phone) >= 8 && len(phone) <= 15)
}

func ValidName(name string) bool {
	if strings.Contains(name, "  ") {
		return false
	}
	if strings.ContainsAny(name, allSymbols) {
		return false
	}
	if strings.ContainsAny(name, numbers) {
		return false
	}
	return (len(name) >= 2 && len(name) <= 80)
}

func ValidUser(user string) bool {
	if strings.Contains(user, " ") {
		return false
	}
	if strings.ContainsAny(user, symbols) {
		return false
	}
	if !strings.ContainsAny(user, leters) {
		return false
	}
	return (len(user) >= 2 && len(user) <= 80)
}

func ValidBussines(bussines string) bool {
	if strings.Contains(bussines, "  ") {
		return false
	}
	if strings.ContainsAny(bussines, allSymbols) {
		return false
	}
	return (len(bussines) >= 2 && len(bussines) <= 80)
}

func ValidPassword(password string) bool {
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZÑ")
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzñ")
	hasSpecial := strings.ContainsAny(password, "!@#$%^&*()_+{}[]|\\;:'\",.<>?/·")
	hasNumber := strings.ContainsAny(password, numbers)

	if len(password) < 8 || len(password) > 25 {
		return false
	}

	if strings.Contains(password, " ") {
		return false
	}

	return (hasUpper && hasLower && hasSpecial && hasNumber)
}

func ValidDate(date string) bool {
	if !strings.Contains(date, "/") {
		return false
	}
	if strings.ContainsAny(date, leters) {
		return false
	}
	if strings.ContainsAny(date, symbolsDate) {
		return false
	}
	return len(date) == 10
}

func ValidTime(time string) bool {
	if !strings.Contains(time, ":") {
		return false
	}
	if strings.ContainsAny(time, leters) {
		return false
	}
	if strings.ContainsAny(time, symbolsHour) {
		return false
	}
	return len(time) == 5
}

func ValidOrderName(orderName string) bool {
	if strings.Contains(orderName, " ") {
		return false
	}
	if strings.ContainsAny(orderName, symbolsOrder) {
		return false
	}
	if strings.ContainsAny(orderName, leters) {
		return false
	}
	return len(orderName) >= 5 && len(orderName) <= 10 && strings.Contains(orderName, "#")
}

func ValidSKU(sku string) bool {
	if strings.ContainsAny(sku, symbolsSKU) {
		return false
	}
	if !strings.ContainsAny(sku, leters) {
		return false
	}
	return len(sku) >= 5 && strings.Contains(sku, "-")
}

func GetOTP() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9))
}

package util

import (
	"fmt"
	"image/color"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

// GenerateQRCode generates a QR code with the id of the user and saves it to a file with the name of the id of the user in the folder ./qr/
func GenerateQRCode(id string) error {

	// Check if directory exists and create it if not
	if err := checkDir("./qr"); err != nil {
		return err
	}

	// Generate QR code and save it to a file that contains the id of the user
	if err := qrcode.WriteColorFile(id, qrcode.Medium, 256, color.Transparent, color.Black, fmt.Sprintf("./qr/%s.png", id)); err != nil {
		return err
	}

	return nil
}

// GenerateQRCode generates a QR code with the id of the user and saves it to a file with the name of the id of the user in the folder ./wqr/
func GenerateWhiteQRCode(id string) error {

	// Check if directory exists and create it if not
	if err := checkDir("./wqr"); err != nil {
		return err
	}

	// Generate QR code and save it to a file that contains the id of the user
	if err := qrcode.WriteColorFile(id, qrcode.Medium, 256, color.Transparent, color.White, fmt.Sprintf("./wqr/%s.png", id)); err != nil {
		return err
	}

	return nil
}

// GenerateQRCode generates a QR code with the id of the user and saves it to a file with the name of the id of the user in the folder ./gqr/
func GenerateGreyQRCode(id string) error {

	// Check if directory exists and create it if not
	if err := checkDir("./gqr"); err != nil {
		return err
	}

	// Generate QR code and save it to a file that contains the id of the user
	if err := qrcode.WriteColorFile(id, qrcode.Medium, 256, color.Transparent, color.RGBA{96, 96, 96, 255}, fmt.Sprintf("./gqr/%s.png", id)); err != nil {
		return err
	}

	return nil
}

func checkDir(dir string) error {

	// Check directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {

		// Create directory if not exists
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

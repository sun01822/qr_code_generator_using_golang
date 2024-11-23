package main

import (
	"image/color"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

type QRCode struct {
	content       string
	level         qrcode.RecoveryLevel
	versionNumber int
	foreground    color.Color
	background    color.Color
	disableBorder bool
}

func NewQRCode(content string, level qrcode.RecoveryLevel, versionNumber int, foreground, background color.Color, disableBorder bool) *QRCode {
	return &QRCode{
		content:       content,
		level:         level,
		versionNumber: versionNumber,
		foreground:    foreground,
		background:    background,
		disableBorder: disableBorder,
	}
}

func main() {
	// Create a new QRCode instance
	q := NewQRCode("http://github.com/sun01822", qrcode.High, 20, color.Black, color.White, true)

	// Generate the QR code
	qr, err := qrcode.New(q.content, q.level)
	if err != nil {
		log.Fatalf("Failed to create QR code: %v", err)
	}

	// Configure the QR code appearance
	qr.DisableBorder = q.disableBorder
	qr.VersionNumber = q.versionNumber

	// Save the QR code to a file
	err = qr.WriteFile(256, "github_qrcode1.png")
	if err != nil {
		log.Fatalf("Failed to save QR code: %v", err)
	}

	log.Println("QR code generated and saved as github_qrcode.png")
}

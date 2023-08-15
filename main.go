package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

input := "Hello, world!"
qrCode, err := qr.Encode(input, qr.M, qr.Alto)

if err != nil {
	panic(err)
}

qrCode, err = barcode.Scale(qrCode, 200, 200)

file, err := os.Create("qrcode.png")
defer file.Close()

png.Encode(file, qrCode)
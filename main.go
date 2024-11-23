package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main(){
	err := qrcode.WriteFile("http://github.com/sun01822", qrcode.High, 256, "qrcode.png")
	if err != nil {
		fmt.Println(err.Error())
    }
}
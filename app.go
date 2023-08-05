package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"

	comp "github.com/oshimoto/compare2images"
)

// App struct
type App struct {
	ctx context.Context
}

type Results struct {
	ImageData string
	Percent float64
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	//Wwidth, Wheight := runtime.WindowGetSize(a.ctx)
	//println(Wwidth, Wheight)
	}

// Compare the two images
func (a *App) CompareFiles(Image1B64 string, Image2B64 string, Image1Type string, Image2Type string) (Results) {

	var Image1ByteNumber int
	var Image2ByteNumber int

	// Switch depending on image type
	switch (Image1Type) {
		case "image/png": {
			Image1ByteNumber = 22
			break
		}
		case "image/jpeg": {
			Image1ByteNumber = 23
			break
	}
}
	switch (Image2Type) {
		case "image/png": {
			Image2ByteNumber = 22
			break
		}
		case "image/jpeg": {
			Image2ByteNumber = 23
			break
	}
}

	//Get the WebApi File Blob data prefix to be spliced back later
	B64Prefix := Image1B64[0:Image1ByteNumber]
	 
	// Decode the Base64
	image1Byte, err := base64.StdEncoding.DecodeString(Image1B64[Image1ByteNumber:])
	if err != nil {
	fmt.Println("Image 1 Base64 Decoding Error:", err)
	} else {
		fmt.Println("Image 1 Base64 Decoded Successfully",)
	}

	image2Byte, err := base64.StdEncoding.DecodeString(Image2B64[Image2ByteNumber:])
	if err != nil {
	fmt.Println("Image 2 Base64 Decoding Error:", err)
	} else {
		fmt.Println("Image 2 Base64 Decoded Successfully",)
	}

	bufimg1 := bytes.NewBuffer(image1Byte)
	bufimg2 := bytes.NewBuffer(image2Byte)

	Image1, err := png.Decode(bufimg1)
	if err != nil {
		fmt.Println("Image 1 Decoding Error:", err)
		} else {
			fmt.Println("Image 1 Decoded Successfully",)
		}
	Image2, err := png.Decode(bufimg2)
	if err != nil {
		fmt.Println("Image 2 Decoding Error:", err)
		} else {
			fmt.Println("Image 2 Decoded Successfully",)
		}

	CompResult, err := comp.Compare2images(&Image1, &Image2)
	if err != nil {
		fmt.Println("Image Compare Error:", err)
		return Results{"",0}
		} else {
			fmt.Println("Images Compared Successfully",)
		}

	CompResultImgRGBA := CompResult.RedGreen
	var buff bytes.Buffer
	err = png.Encode(&buff, CompResultImgRGBA)
	if err != nil {
		fmt.Println("Image Encoding Error:", err)
		
		} else {
			fmt.Println("Image Encoded Successfully",)
		}
	PercentDiff := 100 - CompResult.Percent
	ResultImageB64 := base64.StdEncoding.EncodeToString(buff.Bytes())

	return Results{B64Prefix+ResultImageB64, PercentDiff}
}

// For Debugging
func (a *App) DebugMe(file string) {


	println("This is what we get:", file[0:30])
	
	_, err := base64.StdEncoding.DecodeString(file[22:])
	if err != nil {
	fmt.Println("Decoding Error:", err)
	} else {
		fmt.Println("File Decoded Successfully",)
	}
	
}



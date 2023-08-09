package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"

	comp "github.com/oshimoto/compare2images"
)

// App struct
type App struct {
	ctx context.Context
}

type Results struct {
	ImageData string
	Percent   float64
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
func (a *App) CompareFiles(Image1B64 string, Image2B64 string, Image1Type string, Image2Type string) Results {

	var Image1ByteNumber int
	var Image2ByteNumber int
	var Image1 image.Image
	var Image2 image.Image

	// Switch depending on image type
	switch Image1Type {
	case "image/png", "image/bmp":
		Image1ByteNumber = 22
	case "image/jpeg", "image/webp":
		Image1ByteNumber = 23
	}

	switch Image2Type {
	case "image/png", "image/bmp":
		Image2ByteNumber = 22
	case "image/jpeg", "image/webp":
		Image2ByteNumber = 23
	}

	//Get the WebApi File Blob data prefix to be spliced back later
	B64Prefix := Image1B64[0:Image1ByteNumber]

	// Decode the Base64 to bytes
	image1Byte, err := base64.StdEncoding.DecodeString(Image1B64[Image1ByteNumber:])
	if err != nil {
		fmt.Println("Image 1 Base64 Decoding Error:", err)
	} else {
		fmt.Println("Image 1 Base64 Decoded Successfully")
	}

	image2Byte, err := base64.StdEncoding.DecodeString(Image2B64[Image2ByteNumber:])
	if err != nil {
		fmt.Println("Image 2 Base64 Decoding Error:", err)
	} else {
		fmt.Println("Image 2 Base64 Decoded Successfully")
	}

	bufimg1 := bytes.NewBuffer(image1Byte)
	bufimg2 := bytes.NewBuffer(image2Byte)

	switch Image1Type {
	case "image/png":
		Image1, err = png.Decode(bufimg1)
		if err != nil {
			fmt.Println("Image 1 Decoding Error:", err)
		} else {
			fmt.Println("Image 1 Decoded Successfully")
		}
	case "image/jpeg":
		Image1, err = jpeg.Decode(bufimg1)
		if err != nil {
			fmt.Println("Image 1 Decoding Error:", err)
		} else {
			fmt.Println("Image 1 Decoded Successfully")
		}
	case "image/webp":
		Image1, err = webp.Decode(bufimg1)
		if err != nil {
			fmt.Println("Image 1 Decoding Error:", err)
		} else {
			fmt.Println("Image 1 Decoded Successfully")
		}
	case "image/bmp":
		Image1, err = bmp.Decode(bufimg1)
		if err != nil {
			fmt.Println("Image 1 Decoding Error:", err)
		} else {
			fmt.Println("Image 1 Decoded Successfully")
		}
	}

	switch Image2Type {
	case "image/png":
		Image2, err = png.Decode(bufimg2)
		if err != nil {
			fmt.Println("Image 2 Decoding Error:", err)
		} else {
			fmt.Println("Image 2 Decoded Successfully")
		}
	case "image/jpeg":
		Image2, err = jpeg.Decode(bufimg2)
		if err != nil {
			fmt.Println("Image 2 Decoding Error:", err)
		} else {
			fmt.Println("Image 2 Decoded Successfully")
		}
	case "image/webp":
		Image2, err = webp.Decode(bufimg2)
		if err != nil {
			fmt.Println("Image 2 Decoding Error:", err)
		} else {
			fmt.Println("Image 2 Decoded Successfully")
		}
	case "image/bmp":
		Image2, err = bmp.Decode(bufimg2)
		if err != nil {
			fmt.Println("Image 2 Decoding Error:", err)
		} else {
			fmt.Println("Image 2 Decoded Successfully")
		}
	}

	CompResult, err := comp.Compare2images(&Image1, &Image2)
	if err != nil {
		fmt.Println("Image Compare Error:", err)
		return Results{"", 0}
	} else {
		fmt.Println("Images Compared Successfully")
	}

	CompResultImgRGBA := CompResult.RedGreen
	var buff bytes.Buffer
	err = png.Encode(&buff, CompResultImgRGBA)
	if err != nil {
		fmt.Println("Image Encoding Error:", err)

	} else {
		fmt.Println("Image Encoded Successfully")
	}
	ResultImageB64 := base64.StdEncoding.EncodeToString(buff.Bytes())
	PercentDiff := 100 - CompResult.Percent

	return Results{B64Prefix + ResultImageB64, PercentDiff}
}

// For Debugging
func (a *App) DebugMe(file string) {

	println("This is what we get:", file[0:30])

	_, err := base64.StdEncoding.DecodeString(file[22:])
	if err != nil {
		fmt.Println("Decoding Error:", err)
	} else {
		fmt.Println("File Decoded Successfully")
	}

}

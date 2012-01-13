package main

import (
	"fmt"
	"http"
	//"io"
	"image"
	"image/png"
	"image/draw"
	"os"
	//"time"
	//"io/ioutil"
)

const WIDTH = 300
const HEIGHT = 409
const STEP_X = 1500
const STEP_Y = 1500

//middle of resort
//const START_X = 50644500
//const START_Y = 15433500
//const FROM_X = -14
//const TO_X = 7
//const FROM_Y = -7
//const TO_Y = 9

//yotei peak
const START_X = 50689500
const START_Y = 15417000
const FROM_X = -11
const TO_X = 12
const FROM_Y = -7
const TO_Y = 8

func main() {

	//saveNow(FROM_X, FROM_Y, 0, 0, `tmp1.png`)
	//saveNow(FROM_X, 1, 0, TO_Y, `tmp2.png`)
	//saveNow(1, FROM_Y, TO_X, 0, `tmp3.png`)
	saveNow(1, 1, TO_X, TO_Y, `tmp4.png`)

}

func saveNow(fromx int, fromy int, tox int, toy int, filename string) {

	img := image.NewRGBA((tox-fromx+1)*WIDTH, (toy-fromy+1)*HEIGHT)

	for xOffset := fromx; xOffset <= tox; xOffset++ {
		for yOffset := fromy; yOffset <= toy; yOffset++ {
			fmt.Println(xOffset, yOffset)
			drawNow(xOffset, yOffset, fromx, fromy, img)
			//time.Sleep(5e8)
		}
	}

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)

	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	png.Encode(f, img)

}

func drawNow(offsetX int, offsetY int, minx int, miny int, img *image.RGBA) {

	draw.Draw(img, image.Rect((offsetX-minx)*WIDTH, (offsetY-miny)*HEIGHT, (offsetX-minx+1)*WIDTH, (offsetY-miny+1)*HEIGHT), get(offsetX, offsetY), image.ZP, draw.Src)

}

func get(offsetX int, offsetY int) image.Image {
	offsetY = -offsetY
	x := START_X + offsetX*STEP_X
	y := START_Y + offsetY*STEP_Y

	url := fmt.Sprint(`http://cyberjapandata.gsi.go.jp/data/15nti/new/`, x, `/`, x, `-`, y, `-img.png`)

	//fmt.Println(url)

	r, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	img, err := png.Decode(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	return img

}

//http://cyberjapandata.gsi.go.jp/data/15nti/new/50644500/50644500-15433500-img.png
//                                               xxxxxxxx xxxxxxxx yyyyyyyy
//x increment by 1500 for next east tile
//y increment by 1500 for next north tile

//300 x 409

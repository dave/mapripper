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
	"io/ioutil"
	//"bufio"
)

const WIDTH = 300
const HEIGHT = 409
const STEP_X = 1500
const STEP_Y = 1500

//middle of resort
const START_X = 50644500
const START_Y = 15433500
//const FROM_X = -14
//const TO_X = 7
//const FROM_Y = -7
//const TO_Y = 9

//yotei detail
//const START_X = 50689500
//const START_Y = 15417000
const FROM_X = 0
const TO_X = 21
const FROM_Y = -8
const TO_Y = 18

//yotei area
//const START_X = 50689500
//const START_Y = 15417000
//const FROM_X = -17
//const TO_X = 17
//const FROM_Y = -12
//const TO_Y = 13

func main() {

	saveNow(FROM_X, FROM_Y, TO_X, TO_Y, `../maps/niseko.png`)
	//	saveNow(FROM_X, FROM_Y, 0, 0, `tmp1.png`)
	//	saveNow(FROM_X, 1, 0, TO_Y, `tmp2.png`)
	//	saveNow(1, FROM_Y, TO_X, 0, `tmp3.png`)
	//saveNow(1, 1, TO_X, TO_Y, `tmp4.png`)

}

func saveNow(fromx int, fromy int, tox int, toy int, filename string) os.Error {

	img := image.NewRGBA((tox-fromx+1)*WIDTH, (toy-fromy+1)*HEIGHT)

	height := HEIGHT
	rowHeight := 0

	for yOffset := fromy; yOffset <= toy; yOffset++ {
		fmt.Print(yOffset-fromy, "/", toy-fromy)
		for xOffset := fromx; xOffset <= tox; xOffset++ {
			//fmt.Println(xOffset, yOffset)
			fmt.Print(".")
			h, err := drawNow(xOffset, yOffset, fromx, fromy, img, rowHeight)
			if h > -1 {
				height = h
			}

			if err != nil {
				fmt.Println(err)
				return err
			}
			//time.Sleep(5e8)
		}
		fmt.Println("")
		rowHeight += height
	}
	fmt.Println("Saving...")
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	png.Encode(f, img)

	return nil
}

func drawNow(offsetX int, offsetY int, minx int, miny int, img *image.RGBA, absy int) (int, os.Error) {

	tile, err := get(offsetX, offsetY)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	h := -1
	if tile != nil {

		h = tile.Bounds().Dy()

		//draw.Draw(img, image.Rect((offsetX-minx)*WIDTH, (offsetY-miny)*HEIGHT, (offsetX-minx+1)*WIDTH, (offsetY-miny+1)*HEIGHT), tile, image.ZP, draw.Src)

		draw.Draw(img, image.Rect((offsetX-minx)*WIDTH, absy, (offsetX-minx+1)*WIDTH, absy+h), tile, image.ZP, draw.Src)

	}

	return h, nil
}

func get(offsetX int, offsetY int) (image.Image, os.Error) {

	offsetY = -offsetY
	x := START_X + offsetX*STEP_X
	y := START_Y + offsetY*STEP_Y

	f, err := os.Open(fmt.Sprint(`../mapcache/img-`, x, `-`, y, `.png`))

	//if err != nil {
	//fmt.Println(err)
	//return nil, err
	//}

	if f == nil {

		f, err = saveToCache(x, y)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

	}

	defer f.Close()

	//f2, err := os.Open(fmt.Sprint(`../mapcache/img-`, x, `-`, y, `.png`))

	//if err != nil {
	//	fmt.Println(err)
	//}

	//defer f2.Close

	//	imgConfig, err := png.DecodeConfig(f)

	//	if err != nil {
	//		fmt.Println(err)
	//		return nil, err
	//	}

	//	fmt.Println(imgConfig.Width, imgConfig.Height)

	//	f.Close()

	//	f, err = os.Open(fmt.Sprint(`../mapcache/img-`, x, `-`, y, `.png`))

	//	if err != nil {
	//		fmt.Println(err)
	//		return nil, err
	//	}

	//	defer f.Close()

	img, err := png.Decode(f)

	if err != nil {
		fmt.Println(err, fmt.Sprint(`../mapcache/img-`, x, `-`, y, `.png`))
		return nil, nil
	}

	return img, nil

}

func saveToCache(x int, y int) (*os.File, os.Error) {

	url := fmt.Sprint(`http://cyberjapandata.gsi.go.jp/data/15nti/new/`, x, `/`, x, `-`, y, `-img.png`)

	//fmt.Println(url)

	r, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = ioutil.WriteFile(fmt.Sprint(`../mapcache/img-`, x, `-`, y, `.png`), data, 0666)

	//f1, err := os.Create()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	f, err := os.Open(fmt.Sprint(`../mapcache/img-`, x, `-`, y, `.png`))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return f, nil
	//defer f1.Close()

	//f1.Write(r.Body)

}

//http://denshikokudo.jmc.or.jp/map.asp?x=140.678444&y=%2042.868297&scl=%2075000&po=9999998,9999997,9999996
//http://cyberjapandata.gsi.go.jp/data/15nti/new/50644500/50644500-15433500-img.png
//                                               xxxxxxxx xxxxxxxx yyyyyyyy
//x increment by 1500 for next east tile
//y increment by 1500 for next north tile

//300 x 409

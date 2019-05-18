package camera


import (
"gocv.io/x/gocv"
)
var flag bool = false
func Newcamera() {

	if flag {
		return
	}
	webcam, _ := gocv.VideoCaptureDevice(0)
	flag = true
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}




package main

import "fmt"

type ElectronicDevicesActs interface {
	Touch()
	TakePhoto()
	ScreenShot()
}
type SmartPhonesActs interface {
	ElectronicDevicesActs
	Call()
}

type LaptopsActs interface {
	ElectronicDevicesActs
	KeyBoard()
}

type Laptop struct {
	model string
	size  int
}

func main() {
	var laptopsActs LaptopsActs = &Laptop{model: "asus", size: 24}

	laptopsActs.Touch()

}

func (*Laptop) Touch() {
	fmt.Println("asus touched")
}

func (*Laptop) TakePhoto() {
	fmt.Println("asus taked photo")

}

func (*Laptop) ScreenShot() {
	fmt.Println("asus taked ScreenShot")

}

func (*Laptop) KeyBoard() {
	fmt.Println("asus KeyBoard")

}

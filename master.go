package main

import (
	"fmt"
)

func main() {
	Iphone11 := eletronic.NewApplePhone("Iphone11")
	HTC12 := eletronic.NewAndroidPhone("HTC", "A12")
	CCCP := eletronic.NewStantionPhone("CCCP", "ЗАЛП", 10)
	printCharacteristics(Iphone11)
	printCharacteristics(HTC12)
	printCharacteristics(CCCP)

}

func printCharacteristics(p eletronic.Phone) {
	fmt.Printf("%s\n%s\n%s\n", p.Brand(), p.Model(), p.Type())
}

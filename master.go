package main

import (
	"fmt"

	eletronic "github.com/helloujin/go_mod_8/strctPhone"
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
	switch p.Type() {
	case "Stacionar":
		p, da := p.(eletronic.StationPhone)
		if da {
			fmt.Printf("%+v\n\n", p.ButtonsCount())
		}
	case "Smart":
		p, da := p.(eletronic.Smartphone)
		if da {
			fmt.Printf("%+v\n\n", p.OS())
		}

	}
}

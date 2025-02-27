package puppy

import (
	"fmt"

	"github.com/NageshJ2014/wilddog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof Woof Woof"
}

func Welcome(s string) string {
	return "Hello I am you Favourite" + s
}

func Wildog_bark() string {
	return wilddog.When_Grownup(Bark())
}

func Wildog_barks() string {
	return wilddog.When_Grownup(Barks())
}

func Fromv11() {
	fmt.Println("I am Puppy From Version v1.0.0")
}

func Fromv12() {
	fmt.Println("I am Puppy From Version v1.2.0")
}

func Fromv13() {
	fmt.Println("I am Puppy From Version v1.3.0")
}

func Fromv14() string {
	fmt.Println("I am Puppy From Version v1.4.0 and Also Returns")
	return "Puppy Version v.1.4.0"
}

func Fromv15() string {
	fmt.Println("I am Puppy From Version v1.5.0 and Also Returns")
	return "Puppy Version v.1.5.0"
}

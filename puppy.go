package puppy

import "github.com/NageshJ2014/wilddog"

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

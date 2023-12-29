package puppy

import (
	"github.com/pop96/dog"
)

func Bark() string {
	return "Wooh!!"
}

func Barks() string {
	return "Wooh, Wooh, Wooh!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

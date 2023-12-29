package puppy

import (
	"fmt"

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

func V100() {
	fmt.Println("This is from v1.0.0")
}

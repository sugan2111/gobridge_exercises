package main

import (
	"fmt"
	"github.com/sugan2111/wwg/animals"
)

func main() {
	kitty := animals.Kitten{}
	kitty.SetName("Ms Snowbell")
	fmt.Println(kitty.GetName() + " is a cat and she " + kitty.Sound())

	doggy := animals.Dog{}
	fmt.Println(doggy.Bark())

}


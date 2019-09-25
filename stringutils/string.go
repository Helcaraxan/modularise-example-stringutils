package stringutils

import (
	"fmt"

	"github.com/Helcaraxan/modularise-example-stringutils/internal/random"
)

const randomStringLenght = 20

func PrintRandomString() {
	fmt.Println(random.GenerateRandomString(randomStringLenght))
}

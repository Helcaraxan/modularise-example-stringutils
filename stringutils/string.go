package stringutils

import (
	"fmt"

	"github.com/Helcaraxan/modularise-example-stringutils/v2/internal/random"
)

func PrintRandomString(stringLenght int) {
	fmt.Println(random.GenerateRandomString(stringLenght))
}

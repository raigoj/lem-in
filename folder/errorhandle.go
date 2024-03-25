package folder

import (
	"fmt"
	"os"
)

func ErrHandler() {
	fmt.Println("\nERROR: invalid data format\n")
	os.Exit(1)
}

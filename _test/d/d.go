package d

import (
	"C"
)

func Atoi(in string) (int, error) {
	return C.atoi(in)
}

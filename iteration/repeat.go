package iteration

import (
	"fmt"
	"strings"
)

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		// repeated = repeated + character
		repeated += character
	}
	return repeated
}

func ExampleClone(str string) string {
	s := str
	clone := strings.Clone(s)
	return clone
}

func ExampleFor() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

package card

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

type ValidationError struct{}

func (m *ValidationError) Error() string {
    return "Card number is invalid"
}

func Anonimize() (string, error) {
	fmt.Println("Enter card number")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	if (validate(s)) {
		return "**** **** **** " + strings.Split(s, " ")[3], nil
	}
	return s, &ValidationError{}
}

func validate(s string) (v bool) {
	arr := strings.Split(s, " ")
	if (len(arr) != 4) {
		return false
	}

	for _, val := range arr {

		a := strings.Split(val, "")

		if (len(a) != 4) {
			return false
		}

		for _, n := range arr {
			_, err := strconv.Atoi(n)
			if (err != nil ) {
				return false
			}
		}
	}

	return true
}
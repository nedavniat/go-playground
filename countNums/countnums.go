package countnums

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func CountNums() byte {
	fmt.Println("Enter comma separated numbers")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	var arr []string = strings.Split(s, ",")
	var sum byte
	for i := 0; i < len(arr); i++ {
		n, err := strconv.Atoi(arr[i])
		if (err != nil) {
			continue
		}
		if (n > 0 && n % 2 == 0) {
			sum += 1
		}
	}
	return sum
}
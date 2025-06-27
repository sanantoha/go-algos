package main

import (
	"fmt"
	"strconv"
	"strings"
)

// O(1) time | O(1) space
func validIPAddresses(str string) []string {

	address := make([]string, 4)

	res := make([]string, 0)

	l := len(str)

	for i := 0; i < min(l, 4); i++ {
		address[0] = str[0:i]

		if !isValid(address[0]) {
			continue
		}

		for j := i + 1; j < min(l, i+4); j++ {
			address[1] = str[i:j]

			if !isValid(address[1]) {
				continue
			}

			for k := j + 1; k < min(l, j+4); k++ {
				address[2] = str[j:k]
				address[3] = str[k:]

				if isValid(address[2]) && isValid(address[3]) {
					res = append(res, strings.Join(address, "."))
				}
			}
		}
	}

	return res
}

func isValid(data string) bool {
	intVal, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if intVal < 0 || intVal > 255 {
		return false
	}
	return strconv.Itoa(intVal) == data
}

func main() {

	str := "1921680"
	// [1.9.216.80 1.92.16.80 1.92.168.0 19.2.16.80 19.2.168.0 19.21.6.80 19.21.68.0 19.216.8.0 192.1.6.80 192.1.68.0 192.16.8.0]
	fmt.Println(validIPAddresses(str))
}

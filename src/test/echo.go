package main

import (
	"bufio"
	"fmt"
	"os"
)

func Echo() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("猜100内的数字")
	number := rand.Intn(100)
	reader := bufio.NewReader(os.Stdin)
	Loop:
	for {
		s, _ := reader.ReadString('\n')
		i, _ := strconv.Atoi(strings.TrimSpace(s))
		switch{
		case i > number:
			fmt.Println("猜大了！")
		case i < number:
			fmt.Println("猜小了！")
		case i == number:
			fmt.Println("猜对了！")
			break Loop
		}
	}
}
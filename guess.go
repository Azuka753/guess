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
	fmt.Println("猜0~1000的数字")
	number := rand.Intn(1000)
	reader := bufio.NewReader(os.Stdin)
	count := 0
	Loop:
	for {
		s, _ := reader.ReadString('\n')
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Println("检查输入，再来一遍")
			continue
		}
		count++
		switch{
		case i > number:
			fmt.Println("猜大了！")
		case i < number:
			fmt.Println("猜小了！")
		case i == number:
			fmt.Println("猜对了！你一共猜了", count, "次")
			break Loop
		}
	}
}
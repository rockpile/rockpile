package main

import ("fmt"
	"io/ioutil"
	"strings"
	)

type Data struct {
	Key int
	Val string
} 

func Read(fileName string) string {
	body, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("no")
		return ""
	}

	return  string(body)
}

func main() {
	fmt.Printf("%q\n", strings.Split(Read("/home/rockpile/go/user.txt"), "\n"))

	arrstr := strings.Split(Read("./user.txt"), "\n")

	fmt.Printf("%d\n", len(arrstr))

	for i := len(arrstr) - 2; i >= 0; i-- {
		item := strings.Split(arrstr[i], " ")
		fmt.Printf("item:%s, %s\n", item[0], item[1])

		if "rockpile" == item[0] && "beyond" == item[1] {
			fmt.Printf("match\n")
		}
	}
}

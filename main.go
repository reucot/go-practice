package main

import (
	ip "go-practice/interface_practice"
	"io"
	"log"
	"os"
	"strings"
)

//Проверка реализует ли тип MockClient интерфейс Client
//var _ Client = (*MockClient)(nil)

func main() {

	// a := "some text"
	// defer func(s string) {
	// 	fmt.Println(a)
	// }(a)
	// a += "another text"

	// item1 := funcs_practice.NewItem()
	// // с применением одной опции
	// item2 := funcs_practice.NewItem(funcs_practice.Option2(70))
	// // или двух
	// item3 := funcs_practice.NewItem(funcs_practice.Option1("unusual"), funcs_practice.Option2(99))
	// // опции можно заявлять в разном порядке
	// item4 := funcs_practice.NewItem(funcs_practice.Option2(88), funcs_practice.Option1("rare"), func(i *funcs_practice.Item) {
	// 	i.NoOption = "ins"
	// })

	// fmt.Println(item1)
	// fmt.Println(item2)
	// fmt.Println(item3)
	// fmt.Println(item4)

	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := ip.LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}

}

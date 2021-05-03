package main

import "fmt"

func a(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func b()  {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
func main()  {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

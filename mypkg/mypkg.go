package mypkg

//import "fmt"  //1

//import stdio "fmt" //2

//import . "fmt" //3
import (
	"fmt"
	"strconv"
	"iotestgo/mypkg/a" //此处执行init
) //4
func Mypkg()  {
	//fmt.Println("hello world!")  //1
	//stdio.Println("stdio hello world!") //2
	//Println(". hello world!") //3

	//4
	var istr int=123456
	fmt.Println(strconv.Itoa(istr))

	fmt.Println("========================1=======================")
	//========================================================================

	a.Testinita() //并非使用的时候才init
}

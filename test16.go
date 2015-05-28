package main 

import (
	"fmt"
)

func main() {
//	var a = [5]int{1, 2, 3, 4, 5};
//	for key,val := range a {
//		fmt.Println(key, "=", val)
//	}
//	
//	var mySlice = a[1:]
//	
//	for _, val := range mySlice {
//		fmt.Println(val)
//	}
//	
//	mySlice = append(mySlice, 200, 300)
//	
//	a[2] = 100
//	
//	for key, val := range mySlice {
//		fmt.Println(key, "=", val)
//	}


	type test struct {
		name string `lijianwei`
		flag bool
		a int
	}
	
	var t test
	t.name = `"dsfas"` + `sadfad`
	fmt.Println(t.name, t.flag, t.a)
}



package main

import "fmt"

func main(){
	var list_data = []int{1,2,3,4,5,6}
	for i ,d := range list_data{  // range 不同的是（冒号，等号）
		if d == 4 {
			fmt.Println(i)
		}
	}

}
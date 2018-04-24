package main

import "fmt"

func cou() []string{
	coun := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	need := coun[:len(coun)-2]
	countcp := make([]string, len(need))
	copy(countcp, need)

	return countcp
}

func main()  {
	coued := cou()
	fmt.Println(coued)
}

//func main() {
//	cars := []string{"jack", "helly", "marry"}
//	caa := []string{"jamesss", "job"}
//	caart := append(cars, caa...)
//	fmt.Println("caar:", caart)
//	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
//	cars = append(caart, "jamessss")
//	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
//}


//
//func printarray(a [3][2]string){
//	for _, v1 := range a{
//		for _, v2 := range v1{
//			fmt.Printf("%s ", v2)
//		}
//		fmt.Printf("\n")
//	}
//}
//
//func main()  {
//	a := [3][2]string{
//		{"james", "jack"},
//		{"marry", "lucy"},
//		{"hanm", "kitty"},
//	}
//	printarray(a)
//	var b [3][2]string
//	b[0][0] = "apple"
//	b[0][1] = "samsung"
//	b[1][0] = "microsoft"
//	b[1][1] = "google"
//	b[2][0] =  "AT&t"
//	b[2][1] =  "T-Mobile"
//	fmt.Printf("\n")
//	printarray(b)
//}



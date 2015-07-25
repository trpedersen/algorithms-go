package algo //main

import (
"fmt"
)

func gcd(u int, v int) int {

	var t int
	if u < 0 {
		u = -u
	}
	if v < 0 {
		v = -v
	}

	for ; u > 0; {
		if u < v {
			t = u
			u = v
			v = t
			fmt.Println("swap")
		}
		u = u % v
		fmt.Println(u, v)
	}
	return v
}

type fraction struct {
	numerator int
	denominator int
}

func fractions(){

	var x, y int

	fmt.Print("x y: ")
	for {
		if _, err := fmt.Scanf("%d %d", &x, &y); err != nil{
			fmt.Println(err.Error())
			break
		}
		f1 := fraction{x, y}
		g := gcd(x, y)
		f2 := fraction{x/g, y/g}
		//fmt.Printf("gcd: %d\n", gcd(x, y))
		fmt.Println(f1, f2)
		fmt.Print("x y: ")
	}
}

func threeGcds(){
	var x, y, z int
	fmt.Print("x y z: ")
	for {
		if _, err := fmt.Scanf("%d %d %d", &x, &y, &z); err != nil{
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("gcd: %d\n", gcd( gcd(x, y), z))
		fmt.Print("x y z: ")
	}
}

func binary(x int) {
	var str = ""
	if( x == 0){
		str = "0"
	} else {
		for ; x != 0 ; {
			bit := 1 & x
			str = fmt.Sprintf("%d", bit) + str	
			x = x >> 1	
		}
	}
	fmt.Println(str)
}

func main3(){
	//fractions()
	//for i:= 1; i < 10000; i++ {
	//	binary(i)
	//}
	//binary(1)
	threeGcds()
}
package main

import "basicmath"
import "formattedhello"
import "fmt"

func main(){

	var x ,y int=9,3

	
	var name string="GO"




	formattedhello.GetHelloMessage(name)

	fmt.Println(" x+y =" , basicmath.Add(x,y) )

	fmt.Println(" x+y =" , basicmath.Substract(x,y) )

	fmt.Println(" x+y =" , basicmath.Multiply(x,y) )

	fmt.Println(" x+y =" , basicmath.Divide(x,y) )
	
}

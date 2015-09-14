package main

import "fmt"

var x = [3]float64{1,2,3}
func main() {
	mainForLoopNoIndex() 
	mainForLoopFuncy()
	mainArray() 
	mainSwitch()
	mainForLoopOdd()
	mainForLoop()
	mainForLoopFunny()
	mainStdIn() 
	mainVariablesAndConstants() 
} 

func slices() {
	
}

func mainForLoopNoIndex() {
	
	var total float64 = 0
	// no need i index then use underscore
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

}

func mainForLoopFuncy() {
			
	var total float64 = 0
	for i, value := range x {
		total += value
		fmt.Println(i)
	}
	fmt.Println(total / float64(len(x)))
}

func mainArray() {
	
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func mainSwitch() {
	for i := 0; i<4; i++ {
		switch i {
			case 0: fmt.Println("zero")
			case 1: fmt.Println("One")
			default: fmt.Println("What?")
		}	
	}	
}

func mainForLoopOdd() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func mainForLoop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func mainForLoopFunny() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func mainStdIn() {
	fmt.Print("Enter a tempurture: ")
	var f float64
	fmt.Scanf("%f", &f)

	c := (f - 32) * 5/9
	fmt.Printf("%.2f\n", c)
}

func mainVariablesAndConstants() {
    fmt.Println("hello, world\n")
    fmt.Println("1 + 1 =", 1 + 1)
    fmt.Println("1 + 1 =", 1.0 + 3)
    fmt.Println(len("string length"))
    fmt.Println("string length"[0])
    fmt.Println("The"+" World")
    var myName string = "Ninja" + " Jet"
    myName += " inc"

    var counter int = 1
    counter++
    fmt.Println(counter)
    fmt.Println(myName)

    justNum := 89
    justStr := "oo"
    fmt.Printf("num: %d, str: %s\n", justNum, justStr)

    var combo = 9.0
    fmt.Printf("%.2f\n", combo)

    // constants
    const x string = "cannot change this"
    // x = "oops"
}
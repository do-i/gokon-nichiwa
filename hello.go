package main

import "fmt"

var x = [3]float64{1,2,3}
func main() {
	nestedMaps()
	// coolMap()
	// goodMap()
	// slices()
	// mainForLoopNoIndex() 
	// mainForLoopFuncy()
	// mainArray() 
	// mainSwitch()
	// mainForLoopOdd()
	// mainForLoop()
	// mainForLoopFunny()
	// mainStdIn() 
	// mainVariablesAndConstants() 
} 

func nestedMaps() {
	elements := map[string]map[string]string{
		"H": map[string]string{
		"name":"Hydrogen",
		"state":"gas",
		},
		"He": map[string]string{
		"name":"Helium",
		"state":"gas",
		},
		"Li": map[string]string{
		"name":"Lithium",
		"state":"solid",
		},		
		"Be": map[string]string{
		"name":"Beryllium",
		"state":"solid",
		},
		"B": map[string]string{
		"name":"Boron",
		"state":"solid",
		},
		"C": map[string]string{
		"name":"Carbon",
		"state":"solid",
		},
		"N": map[string]string{
		"name":"Nitrogen",
		"state":"gas",
		},
		"O": map[string]string{
		"name":"Oxygen",
		"state":"gas",
		},
		"F": map[string]string{
		"name":"Fluorine",
		"state":"gas",
		},
		"Ne": map[string]string{
		"name":"Neon",
		"state":"gas",
		},
	}
	// First try to get a value from map, then if it's 
	// successful we run the code inside of the ok block.
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}		
}

func coolMap() {
	elements := map[string]string {
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func goodMap() {
	// make instanciates map type
	mmm := make(map[string]int)
	mmm["one"] = 1
	mmm["two"] = 2
	fmt.Println(mmm)
	delete(mmm, "two")
	fmt.Println(mmm)
}

func badMap() {
	// Runtime Error: panic: assignment to entry in nil map
	var mmm map[string]int
	mmm["one"] = 1
	mmm["two"] = 2
	fmt.Println(mmm)	
}

func slices() {
	// use make keyword to create Slices
	x := make([]float64, 5, 10)
	for i, val := range x {
		fmt.Printf("[%d]=%.2f\n",i, val)
	}

	// use [low:high]
	array := []int{1,2,3,4,5}
	tinySlice := array[1:3] // low inclusive, high exclusive
	fmt.Println(tinySlice)
	fmt.Println(len(tinySlice))

	// append
	biggerSlice := tinySlice[:]
	for i := 0; i < 100; i++ {
		biggerSlice = append(biggerSlice, i)
	}
	fmt.Println(len(biggerSlice))
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
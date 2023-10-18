package main

type Person struct {
	name string
	age  int
}

func main() {

	//Here we are declaring a variable and then initialising it
	var a int
	a = 10
	println(a)

	const b int = 10 //here const keyword is used to define a constant

	//displayLoops()
	//displayConditionalStatements()
	//displayFunctions()
	//displayPointers()
	//displayStructs()
	//displayArrays()
	//displaySlices()
	//displayMaps()
	//displayKeywords()

}

//Code Examples

func displayLoops() {

	//for loop
	for i := 0; i < 10; i++ {
		println(i)
	}
	println("End of Loop")

	//while loop
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	println("End of Loop")

	//infinite loop
	// for {
	// 	println("Hello")
	// }

	//for loop with range
	s := []int{1, 2, 3}
	for i, v := range s {
		println(i, v)
	}
	println("End of Loop")

	//for loop with range and ignoring index
	for _, v := range s {
		println(v)
	}
	println("End of Loop")

	//for loop with range and ignoring value
	for i, _ := range s {
		println(i)
	}
	println("End of Loop")

	//for loop with range and ignoring both index and value
	for range s {
		println("Hello")
	}
}

func displayConditionalStatements() {

	//if statement
	if true {
		println("Hello")
	}

	//if statement with else
	if true {
		println("Hello")
	} else {
		println("World")
	}

	//if statement with else if
	if true {
		println("Hello")
	} else if true {
		println("World")
	} else {
		println("Bye")
	}

	//switch statement
	switch 1 {
	case 1:
		println("Hello")
	case 2:
		println("World")
	}

	//switch statement with fallthrough
	switch 1 {
	case 1:
		println("Hello")
		fallthrough
	case 2:
		println("World")
	default:
		println("Bye")
	}

	//switch statement with default
	switch 1 {
	case 1:
		println("Hello")
	case 2:
		println("World")
	default:
		println("Bye")
	}
}

func displayFunctions() {
	//Here we are declaring a function and then calling it
	f := func() {
		println("Hello World")
	}
	f()

	//Here we are declaring a function and then calling it with arguments
	f2 := func(a int, b int) {
		println(a + b)
	}
	f2(1, 2)

	//Here we are declaring a function and then calling it with arguments and return type
	f3 := func(a int, b int) int { //use of return type is optional here as it is inferred by compiler
		return a + b
	}
	println(f3(1, 2))

	//Here we are declaring a function and then calling it with arguments and return type and name
	f4 := func(a int, b int) (r int) { //use of (r int ) here is that we are naming the return value and we can use it in the function body
		r = a + b
		return
	}
	println(f4(1, 2))

	//Here we are declaring a function and then calling it with arguments and return type and name and multiple return values
	f5 := func(a int, b int) (r1 int, r2 int) {
		r1 = a + b
		r2 = a - b
		return
	}
	println(f5(1, 2))

	//Here we are declaring a function and then calling it with arguments and return type and name and multiple return values and named return values
	f6 := func(a int, b int) (r1 int, r2 int) {
		r1 = a + b
		r2 = a - b
		return r2, r1 //we can also return like this as we have named the return values and we can use them in the function body
	}
	println(f6(1, 2))

	///////////////////////////////Defer/////////////////////////////////////////

	//Here we are declaring a function and then calling it
	f7 := func() {
		println("Hello World")
	}
	defer f7() //defer keyword is used to delay the execution of the function until the surrounding function returns
	println("Hello")

	//Here we are declaring a function and then calling it with arguments
	f8 := func(a int, b int) {
		println(a + b)
	}
	f8(1, 2)
}

func displayPointers() {

	//Creating a Pointer
	var p6 *int //here * is used to create a pointer
	println(p6) //this will print nil as we have not initialised it

	//Creating a Pointer and then initialising it
	var p7 *int
	p7 = new(int) //here new keyword is used to initialise the pointer
	println(p7)   //this will print the address of the variable

	//Creating a Pointer and then initialising it with value
	var p8 *int
	p8 = new(int)
	*p8 = 10 //here * is used to dereference the pointer
	println(*p8)

	//Creating a Pointer and then initialising it with value
	var p9 *int
	p9 = new(int)
	*p9 = 10
	println(p9)  //this will print the address of the variable
	println(*p9) //this will print the value of the variable
}

func displayStructs() {

	//Here First we declare variable and type explicitly and then we are initialising it
	var p Person
	p.name = "John"
	p.age = 25
	println(p.name, p.age)

	//Here we are declaring and initialising variable at the same time
	p2 := Person{"Tom", 30}
	println(p2.name, p2.age)

	//Here we are declaring and initialising variable at the same time but with field names
	p3 := Person{name: "Jane", age: 20}
	println(p3.name, p3.age)

	//Here we are declaring using new keyword and then initialising it
	p4 := new(Person)
	p4.name = "Mike"
	p4.age = 18
	println(p4.name, p4.age)

	//Here we are using pointer to access the fields
	p5 := &Person{"Bob", 50}
	println(p5.name, p5.age)
}

func displayArrays() {

	//Here we are declaring and initialising array at the same time
	a := [3]int{1, 2, 3}
	println(a[0], a[1], a[2])

	//Here we are declaring array and then initialising it
	var a2 [3]int
	a2 = [3]int{1, 2, 3}
	println(a2[0], a2[1], a2[2])

	//Here we are declaring array and then initialising it
	var a3 [3]int
	a3 = [3]int{1, 2, 3}
	println(a3[0], a3[1], a3[2])

}

func displayChannels() {

	//Here we are declaring and initialising channel at the same time
	c := make(chan int)
	println(c)

	//Here we are declaring channel and then initialising it
	var c2 chan int
	c2 = make(chan int)
	println(c2)

	//Here we are declaring channel and then initialising it
	var c3 chan int
	c3 = make(chan int)
	println(c3)

	//Here we are declaring channel and then initialising it with capacity
	var c4 chan int
	c4 = make(chan int, 3)
	println(c4)
}

func displaySlices() {

	//Here we are declaring and initialising slice at the same time
	s := []int{1, 2, 3}
	println(s[0], s[1], s[2])

	//Here we are declaring slice and then initialising it
	var s2 []int
	s2 = []int{1, 2, 3}
	println(s2[0], s2[1], s2[2])

	//Here we are declaring slice and then initialising it but with make keyword
	var s3 []int
	s3 = make([]int, 3)
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	println(s3[0], s3[1], s3[2])

	//Here we are declaring slice and then initialising it but with make keyword and with capacity
	var s4 []int
	s4 = make([]int, 3, 5)
	s4[0] = 1
	s4[1] = 2
	s4[2] = 3
	println(s4[0], s4[1], s4[2])

	//--------functions on slices----------//

	//append
	s5 := []int{1, 2, 3}
	s5 = append(s5, 4, 5, 6)
	println(s5[0], s5[1], s5[2], s5[3], s5[4], s5[5])

	//copy
	s6 := []int{1, 2, 3}
	s7 := make([]int, 3)
	copy(s7, s6)
	println(s7[0], s7[1], s7[2])
}

func displayMaps() {
	//Here we are declaring and initialising map at the same time
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	println(m["one"], m["two"], m["three"])

	//Here we are declaring map and then initialising it
	var m2 map[string]int
	m2 = map[string]int{"one": 1, "two": 2, "three": 3}
	println(m2["one"], m2["two"], m2["three"])

	//Here we are declaring map and then initialising it but with make keyword
	var m3 map[string]int
	m3 = make(map[string]int)
	m3["one"] = 1
	m3["two"] = 2
	m3["three"] = 3
	println(m3["one"], m3["two"], m3["three"])

	//Here we are declaring map and then initialising it but with make keyword and with capacity
	var m4 map[string]int
	m4 = make(map[string]int, 3)
	m4["one"] = 1
	m4["two"] = 2
	m4["three"] = 3
	println(m4["one"], m4["two"], m4["three"])

	//---------------------functions on maps-----------------------//

	//delete
	m5 := map[string]int{"one": 1, "two": 2, "three": 3}
	delete(m5, "two")
	println(m5["one"], m5["three"])

	//len
	m6 := map[string]int{"one": 1, "two": 2, "three": 3}
	println(len(m6))

	//range
	m7 := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m7 {
		println(k, v)
	}

}

func displayKeywords() {
	///////////////////////////////Panic/////////////////////////////////////////

	// Use of Panic is to stop the execution of the program

	//Here we are declaring a function and then calling it
	f9 := func() {
		panic("End it now")
	}
	f9()

	println("This will not be executed")

	///////////////////////////////Recover/////////////////////////////////////////

	// Use of Recover is to recover from the panic

	//Here we are declaring a function and then calling it
	f10 := func() {
		defer func() {
			if r := recover(); r != nil { //recover() returns the value passed to panic() and if there is no panic then it returns nil
				println("Recovered from panic")
			}
		}()
		panic("End it now")
	}
	f10() //this will recover from the panic
}

//Functions
//1. func name() {}
//2. func name() {return}
//3. func name() int {return 0}
//4. func name() (int, int) {return 0,0}
//5. func name() (r1 int, r2 int) {return}
//6. func name() (r1 int, r2 int) {return 0,0}

//7. func name(a int, b int) {}
//8. func name(a int, b int) {return}
//9. func name(a int, b int) int {return 0}
//10. func name(a int, b int) (int, int) {return 0,0}
//11. func name(a int, b int) (r1 int, r2 int) {return}
//12. func name(a int, b int) (r1 int, r2 int) {return 0,0}

//13. func name(a, b int) {}
//14. func name(a, b int) {return}
//15. func name(a, b int) int {return 0}
//16. func name(a, b int) (int, int) {return 0,0}
//17. func name(a, b int) (r1 int, r2 int) {return}
//18. func name(a, b int) (r1 int, r2 int) {return 0,0}

//19. func name(a, b int) (r1, r2 int) {return}
//20. func name(a, b int) (r1, r2 int) {return 0,0}
//21. func name(a, b int) (r1, r2 int) {return r2, r1}

//Data Types
//1. bool
//2. string
//3. int
//4. int8
//5. int16
//6. int32
//7. int64
//8. uint
//9. uint8
//10. uint16
//11. uint32
//12. uint64
//13. uintptr
//14. byte
//15. rune
//16. float32
//17. float64
//18. complex64
//19. complex128

//Keywords and Use
//1. break - To break out of the loop
//2. case - To define a case in switch statement
//3. chan - To define a channel
//4. const - To define a constant
//5. continue - To continue to the next iteration of the loop
//6. default - To define a default case in switch statement
//7. defer - To delay the execution of the function until the surrounding function returns
//8. else - To define a else block in if statement
//9. fallthrough - To fallthrough to the next case in switch statement
//10. for - To define a for loop
//11. func - To define a function
//12. go - To start a goroutine
//13. goto - To jump to a label
//14. if - To define a if statement
//15. import - To import a package
//16. interface - To define a interface
//17. map - To define a map
//18. package - To define a package
//19. range - To iterate over a collection
//20. return - To return from a function
//21. select - To define a select statement
//22. struct - To define a struct
//23. switch - To define a switch statement
//24. type - To define a type
//25. var - To define a variable

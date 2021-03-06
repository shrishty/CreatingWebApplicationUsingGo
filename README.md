# CreatingWebApplicationUsingGo
Web application using golang. Done as part of this course 

https://app.pluralsight.com/library/courses/creating-web-applications-go-update/table-of-contents

Taken from the course
To install local assets:

* install NodeJs (https://nodejs.org)
* install `grunt-cli` globally via `npm install -g grunt-cli`
* run `npm install` from the root directory of the project
* run `grunt` from the root directory of the project to compile CSS

To run the application:
* install the app with `go install <name-of-your-webapp>`
* run with `bin/webapp`
* navigate to `http://localhost:8000/html/home.html`


# Why GO?
* Go addresses challenges which is not addressed in any of the existing languages
* Fully Compiled - Similar to C++
* Simple - Python (No clutter, only 25 keyword)
* Concurrency - Used by Erlang which handles 1,00,000 of sequential processes effortlessly (scaling is better)
* Statically typed language - variables need not be defined before they are used. 

# To start
* `play.golang.go`
* `golang.org`
* Tour of Go in the golang website - TODO: Add link
* effective go, TODO: Add link
* Packages - `golang.org/pkg/`
* The Project - `golang.org/project/` (what all has changed as language has evolved)
* Blog - `blog.golang.org` (Best Prctices are available here)
* gocode - `github.com/nsf/gocode` (Code complete package)

# Package Declaration 

* "package main" - Entry point of the go program
* "func main()" - First function which runs when program is executed
* functions are first class citizens in go

# Our First program

```
package main

func main() {
    println("Hello Go!!")
}

```
# Running the program
* `go run src/main.go`
or
* `go install src/example`
* run the exe file which is generated in the bin folder

# The const keyword 
```
const (
    message = "Hello Go!!"
)
```

* const variables cant be changed

# Variables
```
var (
    message string = "Hello Go!!"
    newMessage string                  //  if not assigned returns empty sting. Go assigns zero value to the variables if not assinged
)
```


# The init() function

* Runs for the first time whenever a module references another module
* init function is called after variable initialization but before any other function is called

# The import keyword 

* Tells go that this module needs to use funtions which are present in another module
* There are two ways to use import

```
import "fmt"
```

or

```
import (
    "fmt"
)
```

* if we are importing a package and not using it we will be compilation error

# Data Types and Operators in Go

## Primitive data types

```
var myInt int             // simplify the parsing of the source code
myInt = 5

var myFloat float32 = 42. // directly assigning during declaration

str := "Hello Go!!"       // short variable declaration

complexNum := complex(2, 3) // first class support for complex numbers, First argument is real and second one is imaginary.

// To fetch the real part use real(complexNum)
// To fetch the imaginary part use imag(complexNum)

```

## Constants

```
const (
    first = "first value"
)
```

```
const (
    first = iota           // assign 0 to first, 1 to second and 2 to third
    second
    third
)
```

### Constant expression

```
const (
    first = 1 << (10 * iota)  // 2^0
    second                    // 2^10
    third                     // 2^100
)
```

## Collections

### Array

```
myArray := [3]int{}
myArray[0] = 42
myArray[1] = 27
myArray[2] = ll

fmt.Println(myArray)

// Another way
myArray := [...]int{42, 27, 99}

// to print the size of the array
sizeOFArray := len(myArray)
```
### Slice
* More flexible than array

```
mySlice := myArray[:]                 // [1:], [:3]
mySlice = append(mySlice, 100)

mySlice  := []float32{14., 15., 16.}  
// if more data gets added, slice has to copy the contents every time new data is added which is inefficient 
fmt.Println(len(mySlice))

// make takes 2 or three arguments
// If two arguments then the second argument is the size of the slice and array both
// If three arguments, second argument = size of the slice, third arguemnt = size of the array
mySlice := make([]float32, 100)        

```

### Map
* Customizable keys
* both key and value can be arbitrary data types

```
myMap := make(map[int]string)         // int is the type of the key, string is type of the value
myMap[42] = "Foo"
myMap[12] = "Bar"

fmt.Println(myMap)
fmt.Println(myMap[42])
```

## Arithmetic Operations
* Addition Operators or string concatenation operator
* subtraction operator
* remainder operator 
* increment operator 
* decrement operator
* Augmented assignment operator

```
add := 1 + 2
str := "Foo" + "Bar"
sub := 3 - 1
remainder := 12 % 4
div : 12 / 4

inc := 1
inc++

dec := 5
dec--            // these are considered as statements and cannot be added with other operators 

dec += 5         // augmented operations
```

# Branching and looping

## Branching

### if statement

```
if foo == 1 {
    println("bar")
} else {
    println("lala")
}

// moving the initialization of foo inside if statement

if foo := 1; foo == 1 {
    println("foo is 1")
}

```

### Switch statement

```
switch foo := 1; foo {
    case 1:
        println("one")
    case 2: 
        println("two")
}

switch {
    case foo == 1:
        println("one")
    case foo > 2:
        println("5")
}
```

### Loops

```
for i:=0; i < 5; i++ {
    println(i)
}

// while kind of for loop

for {
        i++
        println(i)

        if i > 5 {
            break
        }
}

// looping over collection

s := []string{"foo", "bar", "buz"}

for idx, v := range s {
    println(idx, v)
}

myMap := make(map[int]string)         // int is the type of the key, string is type of the value
myMap[42] = "Foo"
myMap[12] = "Bar"

for k, v := range m {
    println(k, v)
}

```

### To take input from the console
```
var option string

fmt.Scanln(&option)               // sends the memory address to populate the value

```

# Function and Parameters

* Pass by Value
* Pass by reference
* Variadic function

```
// defining a function
func sayHello() {
    println("Hello")
}

// calling a function
func main() {
    sayHello()
}

// function with parameters, pass by value

func sayHello(message string) {
    println(message)
}

// func with parameters, pass by reference
func sayHello(message *string) {
    println(*message)                  // dereferencing of pointer
    *message = "bye"
}

// variadic functions

func sayHello (messages ...string) {
    for _, message := range messages {
        println(message)
    }
}

```

## Return values

* Single return value
* Multiple return value
* Named return parameters

```
func add(terms ...int) int {
    result := 0
    for _, term := range terms {
        result += term
    }

    return result
}
```

## Multiple return values

```
func add(terms ...int) (int, int) {
    result := 0
    for _, term := range terms {
        result += term
    }

    return len(terms), result
}

// calling function with multiple return values
numTerms, sum := add(1, 2, 3, 4)
```

## Named return values

```
func add(terms ...int) (numTerms int, sum int) {
    sum := 0
    for _, term := range terms {
        sum += term
    }

    numTerms = len(terms)

    return
}

// calling function with multiple return values
numTerms, sum := add(1, 2, 3, 4)
```

# Anonymous Functions
* Functions are first class citizen

```
func main() {

    addFunc := func(terms ...int) (numTerms int, sum int) {
        sum := 0
        for _, term := range terms {
            sum += term
        }

        numTerms = len(terms)

        return
    }

    numTerms, sum := addFunc(1, 3)
}
```

# Object Oriented Programming 

## Structs and Fields

```
 
type myStruct struct {
    myField string
}

// instantiating a struct variable
foo := myStruct {}
foo.myField = "bar"

foo := mystruct{"bar"}

// to create objects on heap
foo := new(myStruct)
foo.myField = "bar"
```

## Constructor Functions

```
type myStruct struct {
    myMap map[string]string
}

func newMyStruct() *myStruct {
    result := myStruct{}
    result.myMap = map[string]string{}           // this is instantiating the map

    return &result
}

// instantiation
foo := newMyStruct()

foo.myMap["bar"] = "baz"
```

## Methods

* Data is totally separated from the methods
* We can declare the methods anywhere in the package

```
func main() {
    mp := messagePrinter{"Foo"}
    mp.printMessage()
}


type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() {
    println(mp.message)
}
```

## Object Composition 

* instead of inheritance we have composition

```
func main() {
    emp := enhancedMessagePrinter{}
    emp.message = "foo"
    emp.printMessage()

    // or

    emp := enhancedMessagePrinter{messagePrinter{"Foo"}}
}

type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() {
    println(mp.message)
}

type enhancedMessagePrinter struct {
    messagePrinter            // Anonymous field of messagePrinter type
}
```

# Asynchronous Programming

* goroutines - go's thread like construct to run code/functions concurrently 
* channels - sharing data between threads 

## Goroutines

* concurrency is not parallelism 
* Concurrency - only one thing is being run at a time
* Parallelism - Multiple tasks run at same time
* Parallel programming can be done by adding `runtime.GOMAXPROCS(8)` in the code

```
func main() {
    abcGen()           // prints abc in order
    go abcGen()        // if only this is written, nothing will be printer as main has terminated 

    // so we must apply sleep for sometime to see the output
    time.Sleep(100 * time.Millisecond)
}

func abcGen() {
    for l := byte('a'); l < byte('z'); l++ {
        go println(string(l))
    }
}
```

## Channels
* Allows messages to be passed in go routines .
* Ensures data is safely transferred 

```
func main() {
    runtime.GOMAXPROCS(8)

    ch := make(chan string)
    doneCh := make(chan bool)

    go abcGen(ch)         
    go printer(ch, doneCh)

    <-doneCh
}

func abcGen(ch chan string) {
    for l := byte('a'); l < byte('z'); l++ {
        ch <- string(l)
    }
    close(ch)
}

func printer(ch chan string, doneCh chan bool) {
    for l := range ch {
        println(<-ch)
    }

    doneCh <- true
}
```
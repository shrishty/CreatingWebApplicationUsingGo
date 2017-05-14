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
* Concurrency - Used by Erlang which handles 1,00,000 of sequential processes effortlessly (scaleing is better)
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

`
package main

func main() {
    println("Hello Go!!")
}

`
# Running the program
* `go run src/main.go`
or
* `go install src/example`
* run the exe file which is generated in the bin folder

# The const keyword 

const (
    message = "Hello Go!!"
)

* const variables cant be changed

# Variables

var (
    message string = "Hello Go!!"
    newMessage string                  //  if not assigned returns empty sting. Go assigns zero value to the variables if not assinged
)


# The init() function

* Runs for the first time whenever a module references another module
* init function is called after variable initialization but before any other function is called

# The import keyword 

* Tells go that this module needs to use funtions which are present in another module
* There are two ways to use import

`
import "fmt"
`

or

`
import (
    "fmt"
)
`

* if we are importing a package and not using it we will be compilation error

# Data Types and Operators in Go

## Primitive data types

`var myInt int` ---  simplify the parsing of the source code







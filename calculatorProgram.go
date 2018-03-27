package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "math"
)

  func main() {
// allows user to input first value, sets equal to "i"
    fmt.Print("value 1: ")
    var i float64
    fmt.Scan(&i)

//detects which operator the user inputs in the command line under "operator"
    plus := func(r rune) bool {
        return r == '+'
    } //addition

    minus := func(r rune) bool {
        return r == '-'
    } //subtraction

    multiply := func(r rune) bool {
        return r == '*' || r == 'x'
    } //multiplication

    divide := func(r rune) bool {
        return r == '/'
    } //division

    exponent := func(r rune) bool {
      return r == '^'
    } //exponential

    sqroot := func(r rune) bool {
      return r == '√'
    } //square root

// allows for the input of a
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Operator: ")
    operator, _ := reader.ReadString('\n')

// each operator corresponds to its respective command
// e.g. inputting "+" will execute addition
    if strings.IndexFunc(operator, plus) != -1 {
      fmt.Print("value 2: ")
// for the equations that require a second value, the user inputs another number, set to var "j"
      var j float64
      fmt.Scan(&j)
      fmt.Print("= ")
      fmt.Println(i+j)

  } else if strings.IndexFunc(operator, minus) != -1 {
      fmt.Print("value 2: ")
      var j float64
      fmt.Scan(&j)
      fmt.Print("= ")
      fmt.Println(i-j)

  } else if strings.IndexFunc(operator, multiply) != -1 {
      fmt.Print("value 2: ")
      var j float64
      fmt.Scan(&j)
      fmt.Print("= ")
      fmt.Println(i*j)

  } else if strings.IndexFunc(operator, divide) != -1 {
      fmt.Print("value 2: ")
      var j float64
      fmt.Scan(&j)
      fmt.Print("= ")
      fmt.Println(i/j)

  } else if strings.IndexFunc(operator, exponent) != -1 {
      fmt.Print("value 2: ")
      var j float64
      fmt.Scan(&j)
      fmt.Print("= ")
      fmt.Println(math.Pow(i, j))


  } else if strings.IndexFunc(operator, sqroot) != -1 {
      fmt.Print("= ")
      fmt.Println(math.Sqrt(i))


//error message if the user inputs something other than a number or operator
          } else {
            fmt.Println("Error")


          }
  }

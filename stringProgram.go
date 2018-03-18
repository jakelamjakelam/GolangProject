package main

import (
      "bufio"
      "fmt"
      "os"
      "strings"
)

  func main() {
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Enter Text: ")
      text, _ := reader.ReadString('\n')
      fmt.Println(strings.Title(strings.ToLower(text)))

    }

package main

import (
  "shivammittal.in/greetings"
    "github.com/fatih/color"
)

func main() {
    // Use the local module
    message := greetings.Greet("Gopher")
    
    // Use the external module
    color.Blue(message)
    
    color.Green("This text is green!")
    color.Red("This text is red!")
}

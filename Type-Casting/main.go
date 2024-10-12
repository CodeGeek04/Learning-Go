package main

import ("fmt"
"strconv"
)

func main () {
fmt.Println("Conversion of int to float32")
  var a int = 10
  f := float32(a)
  fmt.Println(f)

  fmt.Println("Conversion of string to int and reverse")
  var s string = "100"
  i, err := strconv.Atoi(s)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(i)

  fmt.Println("Conversion of int to string")
  var j int = 200
  str := strconv.Itoa(j)
  fmt.Println(str)


  fmt.Println("Conversion of string to byte array")
  var str1 string = "Hello, World!"
  b := []byte(str1)

  fmt.Println(b)

  fmt.Println("Conversion of byte array to string")
  str2 := string(b)
  fmt.Println(str2)
}

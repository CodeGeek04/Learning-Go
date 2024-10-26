package main
import (
	"fmt"
)
//importing the bytes package so that buffer can be used
import (
	"bytes"
)
func main() {
//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Shivam")
	strBuffer.WriteString(" Mittal")
	fmt.Println("The string buffer output is",strBuffer.String())
}

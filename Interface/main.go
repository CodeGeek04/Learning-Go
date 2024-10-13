package main

import (
  "fmt"
)

type circle struct {
  radius float64
}

type rectangle struct {
  length float64
  width  float64
}

func (c circle) area() float64 {
  return 3.14 * c.radius * c.radius
}

func (r rectangle) area() float64 {
  return r.length * r.width
}

type shape interface {
  area() float64
}

func main() {
  circle1 := circle{radius: 10}
  rectangle1 := rectangle{length: 10, width: 5}
  shapes := []shape{circle1, rectangle1}

  for _, shape := range shapes {
    fmt.Println(shape.area())
  }
}

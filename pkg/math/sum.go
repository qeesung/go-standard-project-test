package mymath

import (
  "github.com/go-standard-project-test/internal/check"
  "fmt"
)

func Sum(a, b int) int {
  if check.CheckInt(a) {
    fmt.Println("a is not int")
  }
  return a+b
}

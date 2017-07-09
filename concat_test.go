package main

import (
  "fmt"
  "strings"
  "testing"
)

func TestConcatTemplates(t *testing.T) {
  b := ConcatTemplates("./templates/services")
  t.Log(string(b))
  fmt.Println(strings.Contains(b, "redis") && strings.Contains(b, "web"))
}

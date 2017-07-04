package main

import (
  "io/ioutil"
  "testing"
)

func TestConcatTemplates(t *testing.T) {
  b := ConcatTemplates("./templatedir")
  t.Log(string(b))
  ioutil.WriteFile("all.tmpl", b, 0644)
}

package main

import (
//  "errors"
  "flag"
  "fmt"
  "os"
//  "path/filepath"
//  "strings"
//  "text/template"
)

const Version = "x.y.z" //DELETE ME WHEN YOU START COMPILING

func main() {
  flag.Usage = func() {
    fmt.Printf("Usage of %s:\n", os.Args[0])
    fmt.Printf("  composiler <environment>\n")
    flag.PrintDefaults()
  }

  var environment string
  var printVersion bool

  flag.BoolVar(&printVersion, "version", false, "Print version and exit merrily.")

  flag.Parse()

  if flag.NArg() == 0 {
      if printVersion {
        fmt.Printf("composiler, %s\n", Version)
        os.Exit(0)
      }
      flag.Usage()
      os.Exit(1)
  } else if flag.NArg() == 1 {
      environment = flag.Arg(0)
  } else {
      flag.Usage()
      os.Exit(1)
  }
  fmt.Printf("Environment is: %s\n", environment)
}

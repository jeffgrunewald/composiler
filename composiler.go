package main

import (
  "bytes"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
//  "text/template"
)

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

/*
  type Config struct {
    Environment string
  }
  c := Config{Environment: environment}
  t, _ := template.New("compose.tmpl").ParseFiles("compose.tmpl")
  t.Execute(os.Stdout, c)
*/

  servs := ConcatTemplates("./templatedir")
  vols := ConcatTemplates("./volumedir")
  s := [][]byte{servs, vols}
  allModules := bytes.Join(s, []byte("\n"))
  ioutil.WriteFile("docker-compose.yml", allModules, 0644)

  fmt.Printf(environment)
}

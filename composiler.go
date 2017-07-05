package main

import (
  "flag"
  "fmt"
//  "io/ioutil"
  "os"
  "text/template"
)

func main() {
  flag.Usage = func() {
    fmt.Printf("Usage of %s:\n", os.Args[0])
    fmt.Printf("  composiler <environment>\n")
    flag.PrintDefaults()
  }

  var composeVersion string
  var conf string
  var environment string
  var printVersion bool

  flag.StringVar(&composeVersion, "compose-version", "3.3", "Set the version of the compose file format.")
  flag.StringVar(&conf, "conf", "/composiler", "Set location of the configs and templates.")
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

  srvs := ConcatTemplates(conf + "/templates/services")
  vols := ConcatTemplates(conf + "/templates/volumes")
  nets := ConcatTemplates(conf + "/templates/networks")
  secs := ConcatTemplates(conf + "/templates/secrets")

  type BaseTemplate struct {
    Srvs string
    Nets string
    Vols string
    Secs string
    ComposeVersion string
  }
  baseTemplate := BaseTemplate{Srvs: srvs, Nets: nets, Vols: vols, Secs: secs, ComposeVersion: composeVersion}
  t, _ := template.New("base-compose.tmpl").ParseFiles("base-compose.tmpl")
  t.Execute(os.Stdout, baseTemplate)

  fmt.Printf("Environment is: %s\n", environment)
}

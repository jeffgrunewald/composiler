package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "text/template"
)

func main() {
  flag.Usage = func() {
    fmt.Printf("Usage of %s:\n", os.Args[0])
    fmt.Printf("  composiler <environment>\n")
    flag.PrintDefaults()
  }

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

  if err := initConfig(); err != nil {
    log.Fatal(err.Error())
  }

  if config.Service == "all" {
    templateSkeleton.Services = ConcatTemplates(conf + "/templates/services")
  } else {
    fmt.Printf("Must set at least one service definition to compose\n")
    os.Exit(1)
  }

  if config.Network == "all" {
    templateSkeleton.Networks = ConcatTemplates(conf + "/templates/networks")
  } else {
    fmt.Printf("Must set at least one network definition to compose\n")
    os.Exit(1)
  }

  if config.Secret == "all" {
    templateSkeleton.Secrets = ConcatTemplates(conf + "/templates/secrets")
  } else {
    fmt.Printf("Must set at least one secret definition to compose\n")
    os.Exit(1)
  }

  if config.Volume == "all" {
    templateSkeleton.Volumes = ConcatTemplates(conf + "/templates/volumes")
  } else {
    fmt.Printf("Must set at least one volume definition to compose\n")
    os.Exit(1)
  }
  
  BuildCombinedTemplate()

  t, err := template.New("combined template").Parse(combinedTemplate)
  if err != nil {
    log.Print(err)
    return
  }
  composeFile, err := os.Create(config.OutFile)
  if err != nil {
    log.Println("create file: ", err)
    return
  }
  err = t.Execute(composeFile, envStruct)
  if err != nil {
    log.Print("execute: ", err)
    return
  }
  composeFile.Close()

  fmt.Printf("Environment is: %s\n", environment)
}

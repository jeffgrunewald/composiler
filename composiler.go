package main

import (
  "flag"
  "fmt"
  "log"
  "os"
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

  servicePath := config.Conf + "/templates/services"
  if config.Service == "all" {
    templateSkeleton.Services = ConcatAllTemplates(servicePath)
  } else {
    templateSkeleton.Services = ConcatSelectTemplates(servicePath, serviceList)
  }

  networkPath := config.Conf + "/templates/networks"
  if config.Network == "all" {
    templateSkeleton.Networks = ConcatAllTemplates(networkPath)
  } else {
    fmt.Printf("Must set at least one network definition to compose\n")
    os.Exit(1)
  }

  secretPath := config.Conf + "/templates/secrets"
  if config.Secret == "all" {
    templateSkeleton.Secrets = ConcatAllTemplates(secretPath)
  } else {
    fmt.Printf("Must set at least one secret definition to compose\n")
    os.Exit(1)
  }

  volumePath := config.Conf + "/templates/volumes"
  if config.Volume == "all" {
    templateSkeleton.Volumes = ConcatAllTemplates(volumePath)
  } else {
    fmt.Printf("Must set at least one volume definition to compose\n")
    os.Exit(1)
  }
  
  combinedTemplate := BuildCombinedTemplate(BaseTemplate, templateSkeleton)

  fmt.Printf("pulling config from environment: %s\n", config.Environment)
  jsonData := DecodeJsonConfig(config.Conf + "/configs/" + config.Environment + ".json")

  fmt.Printf("writing out final compose file to: %s\n", config.OutFile)
  if err := BuildOutputTemplate(combinedTemplate, jsonData, config.OutFile); err != nil {
    log.Fatal(err.Error())
  }
}

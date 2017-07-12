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
  templateSkeleton.Services = ConcatTemplates(servicePath, config.Service)

  networkPath := config.Conf + "/templates/networks"
  templateSkeleton.Networks = ConcatTemplates(networkPath, config.Network)

  secretPath := config.Conf + "/templates/secrets"
  templateSkeleton.Secrets = ConcatTemplates(secretPath, config.Secret)

  volumePath := config.Conf + "/templates/volumes"
  templateSkeleton.Volumes = ConcatTemplates(volumePath, config.Volume)

  combinedTemplate := BuildCombinedTemplate(BaseTemplate, templateSkeleton)

  finalTemplate := RemoveExtraWhitespace(combinedTemplate)

  fmt.Printf("pulling config from environment: %s\n", config.Environment)
  jsonData := DecodeJsonConfig(config.Conf + "/configs/" + config.Environment + ".json")

  fmt.Printf("writing out final compose file to: %s\n", config.OutFile)
  if err := BuildOutputTemplate(finalTemplate, jsonData, config.OutFile); err != nil {
    log.Fatal(err.Error())
  }
}

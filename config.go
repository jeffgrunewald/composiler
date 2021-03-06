package main

import (
  "flag"
  "strings"
)

var (
  composeVersion   string
  conf             string
  config           Config
  environment      string
  network          string
  outFile          string
  printVersion     bool
  secret           string
  service          string
  templateSkeleton TemplateSkeleton
  volume           string
)

type Config struct {
  ComposeVersion string
  Conf           string
  Environment    string
  Network        []string
  OutFile        string
  Secret         []string
  Service        []string
  Volume         []string
}

type TemplateSkeleton struct {
  ComposeVersion string
  Services       string
  Networks       string
  Volumes        string
  Secrets        string
}

func init() {
  flag.StringVar(&composeVersion, "compose-version", "3.3", "Set the version of the compose file format.")
  flag.StringVar(&conf, "conf", "/composiler", "Set location of the configs and templates.")
  flag.StringVar(&network, "network", "all", "Networks to be bundled into the compose file.")
  flag.StringVar(&outFile, "out", "docker-compose.yml", "Name of output file.")
  flag.BoolVar(&printVersion, "version", false, "Print verion and exit merrily.")
  flag.StringVar(&secret, "secret", "all", "Secrets to be bundled into the compose file.")
  flag.StringVar(&service, "service", "all", "Services to be bundled into the compose file.")
  flag.StringVar(&volume, "volume", "all", "Volumes to be bundled into the compose file.")
}

func initConfig() error {

  config = Config{
    ComposeVersion: "3.3",
    Conf:           "/composiler",
    Environment:    environment,
    Network:        []string{},
    OutFile:        "docker-compose.yml",
    Secret:         []string{},
    Service:        []string{},
    Volume:         []string{},
  }

  processFlags()

  templateSkeleton = TemplateSkeleton{
    ComposeVersion: config.ComposeVersion,
  }

  return nil
}

func processFlags() {
  flag.Visit(setFlagConfig)
}

func setFlagConfig(f *flag.Flag) {
  switch f.Name {
    case "compose-version":
      config.ComposeVersion = composeVersion
    case "conf":
      config.Conf = conf
    case "network":
      config.Network = strings.Split(network, ",")
    case "out":
      config.OutFile = outFile
    case "secret":
      config.Secret = strings.Split(secret, ",")
    case "service":
      config.Service = strings.Split(service, ",")
    case "volume":
      config.Volume = strings.Split(volume, ",")
  }
}

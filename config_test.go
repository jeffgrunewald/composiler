package main

import (
  "reflect"
  "testing"
)

func TestInitConfigDefault(t *testing.T) {
  expected := Config{
    ComposeVersion: "3.3",
    Conf:           "/composiler",
    Environment:    "",
    OutFile:        "docker-compose.yml",
    Network:        "all",
    Secret:         "all",
    Service:        "all",
    Volume:         "all",
  }
  if err := initConfig(); err != nil {
    t.Errorf(err.Error())
  }
  if !reflect.DeepEqual(expected, config) {
    t.Errorf("initConfig() = %v, desired %v", config, expected)
  }
}
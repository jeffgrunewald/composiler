package main

import (
  "encoding/json"
  "log"
  "os"
)

func DecodeJsonConfig(filePath string) (data map[string]interface{}) {
  envFile, err := os.Open(filePath)
  if err != nil {
    log.Print(err)
    return
  }
  defer envFile.Close()

  decoder := json.NewDecoder(envFile)
  data = make(map[string]interface{})
  err = decoder.Decode(&data)
  if err != nil {
    log.Print("Failed to decode json config file: ", err)
    return
  }

  return
}
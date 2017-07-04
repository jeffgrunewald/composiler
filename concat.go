package main

import (
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
)

func ConcatTemplates(dirPath string) []byte {
  var combinedTemplates []byte
  filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() && strings.HasSuffix(info.Name(), ".tmpl") {
      println("adding component " + path + " ...")
      b, err := ioutil.ReadFile(path)
      if err != nil {
        return err
      }
      combinedTemplates = append(combinedTemplates, b...)
    }
    return nil
  })
  return combinedTemplates
}

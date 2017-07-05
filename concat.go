package main

import (
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
)

func ConcatTemplates(dirPath string) string {
  var combinedTemplateBytes []byte
  filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() && strings.HasSuffix(info.Name(), ".tmpl") {
      println("adding component " + path + " ...")
      b, err := ioutil.ReadFile(path)
      if err != nil {
        return err
      }
      combinedTemplateBytes = append(combinedTemplateBytes, b...)
    }
    return nil
  })
  combinedTemplateString := string(combinedTemplateBytes[:])

  finalTemplateString := strings.Replace(combinedTemplateString, "\n", "\n  ", -1)

  return finalTemplateString
}

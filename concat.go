package main

import (
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
)

func ConcatTemplates(dirPath string) (formattedTemplate string) {

  var combinedTemplateBytes []byte

  filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() && strings.HasSuffix(info.Name(), ".tmpl") {
      println("adding component " + path + " ...")
      templateBytes, err := ioutil.ReadFile(path)
      if err != nil {
        return err
      }
      combinedTemplateBytes = append(combinedTemplateBytes, templateBytes...)
    }
    return nil
  })

  formattedTemplate = strings.Replace(string(combinedTemplateBytes[:]), "\n", "\n  ", -1)

  return
}

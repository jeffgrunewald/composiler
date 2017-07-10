package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
)

func ConcatTemplates(dirPath string, components []string) (formattedTmpl string) {
  var combinedTmplBytes []byte

  if len(components) == 0 {
    filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
      if !info.IsDir() && strings.HasSuffix(info.Name(), ".tmpl") {
        fmt.Printf("adding component %s ...\n", path)
        tmplBytes, err := ioutil.ReadFile(path)
        if err != nil {
          return err
        }
        combinedTmplBytes = append(combinedTmplBytes, tmplBytes...)
      }
      return nil
    })
  } else {
    for _, component := range components {
      filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() && strings.EqualFold(info.Name(), component + ".tmpl") {
          fmt.Printf("adding component %s ...\n", path)
          tmplBytes, err := ioutil.ReadFile(path)
          if err != nil {
            return err
          }
          combinedTmplBytes = append(combinedTmplBytes, tmplBytes...)
        }
        return nil
      })
    }
  }

  formattedTmpl = strings.Replace(string(combinedTmplBytes[:]), "\n", "\n  ", -1)
  return
}

func RemoveExtraWhitespace (input string) (output string) {
  input = strings.Replace(input, "  \n", "\n", -1)
  input = strings.Replace(input, "\n\n\n", "\n\n", -1)
  output = strings.TrimSpace(input)
  return
}

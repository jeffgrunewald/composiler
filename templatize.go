package main

import (
  "bytes"
  "log"
  "os"
  "text/template"
)

func BuildCombinedTemplate(inTmpl string, tmplData TemplateSkeleton) (outTmpl string) {
  tmpl, err := template.New("base template").Parse(inTmpl)
  if err != nil {
    log.Print(err)
  }

  var tmplBytes bytes.Buffer
  err = tmpl.Execute(&tmplBytes, tmplData)
  if err != nil {
    log.Print("execute: ", err)
  }

  outTmpl = tmplBytes.String()
  return
}

func BuildOutputTemplate(inTmpl string, tmplData map[string]interface{}, file string) error {
  tmpl, err := template.New("combined template").Parse(inTmpl)
  if err != nil {
    log.Print(err)
  }

  outFile, err := os.Create(file)
  if err != nil {
    log.Println("create file: ", err)
  }

  err = tmpl.Execute(outFile, tmplData)
  if err != nil {
    log.Print("execute: ", err)
  }
  outFile.Close()
  return nil
}
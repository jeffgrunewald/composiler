package main

import (
  "bytes"
  "log"
  "text/template"
)

func BuildCombinedTemplate() {

  t, err := template.New("base template").Parse(BaseTemplate)
  if err != nil {
    log.Print(err)
    return
  }

  var combinedTemplateBytes bytes.Buffer
  err = t.Execute(&combinedTemplateBytes, templateSkeleton)
  if err != nil {
    log.Print("execute: ", err)
    return
  }

  combinedTemplate = combinedTemplateBytes.String()
  return
}
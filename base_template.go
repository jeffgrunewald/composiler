package main

const BaseTemplate = `---
version: "{{.ComposeVersion}}"

services:
  {{.Services}}
{{if .Networks}}networks:
  {{.Networks}}{{end}}
{{if .Volumes}}volumes:
  {{.Volumes}}{{end}}
{{if .Secrets}}secrets:
  {{.Secrets}}{{end}}`
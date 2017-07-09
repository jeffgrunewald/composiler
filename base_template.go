package main

const BaseTemplate = `---
version: "{{.ComposeVersion}}"

services:
  {{.Services}}
networks:
  {{.Networks}}
volumes:
  {{.Volumes}}
secrets:
  {{.Secrets}}`
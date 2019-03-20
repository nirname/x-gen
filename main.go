package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	"text/template"
	"strings"
	"path/filepath"
)

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n" + pad, -1)
}

func underscore(v string) string {
  return strings.Replace(strings.ToLower(v), "-", "_", -1)
}

func header(v string) string {
  return strings.Title(strings.Replace(v, "_", "-", -1))
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"indent": indent,
    "underscore": underscore,
    "header": header,
	}
}

type Auth struct {
  Request string
  Headers []string
  Service bool
}

type Location struct {
	Location string // routing path, `/` for example
	Proxy string // service name (upstream name) to proxy to
	Auth Auth // service name for auth_request
	Custom string // cutom configuration that will be in inserted as it is
}

type Context struct {
	Services map[string] []string
	Locations []Location
	RawData string
}

func parseConfig(name string, contextPtr *Context) {
  path := os.Getenv(strings.ToUpper(name) + "PATH")
  if path == "" {
    path = name + ".yml"
  }
  fmt.Fprintln(os.Stderr, "Using " + path)
  data, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err)
  }

  switch ext := filepath.Ext(path); ext {
	case ".yml":
		yaml.Unmarshal(data, contextPtr)
	default:
		contextPtr.RawData += string(data)
	}
}

func main() {
	context := Context{}

	parseConfig("services", &context)
	parseConfig("locations", &context)

	tmpl, err := template.New("nginx.tmpl").Funcs(funcMap()).ParseGlob("*.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, context)
}


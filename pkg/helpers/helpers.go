package helpers

import (
	"bytes"
	"reflect"
	"regexp"
	"strconv"
	"text/template"
)

// KubernetesResourceName returns the resource name for other components
func KubernetesResourceName(name string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9_-]+")
	return reg.ReplaceAllString(name, "-")
}

// Stringify will create a string based on the params
func Stringify(attr interface{}) string {
	switch reflect.TypeOf(attr).Name() {
	case "bool":
		return strconv.FormatBool(attr.(bool))
	case "string":
		return attr.(string)
	case "int":
		return strconv.Itoa(attr.(int))
	}
	return ""
}

// Templatize returns the proper values based on the templating
func Templatize(tempStr string, data interface{}) (resp string, err error) {
	t := template.New("templating")
	t, err = t.Parse(string(tempStr))
	if err != nil {
		return
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	return tpl.String(), err
}

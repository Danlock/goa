package codegen

import (
	"goa.design/goa/pkg"
)

// Header returns a Go source file header section template.
func Header(title, pack string, imports []*ImportSpec) *SectionTemplate {
	return &SectionTemplate{
		Name:   "source-header",
		Source: headerT,
		Data: map[string]interface{}{
			"Title":       title,
			"ToolVersion": pkg.Version(),
			"Pkg":         pack,
			"Imports":     imports,
		},
	}
}

// AddImports adds imports to a section template that was generated with
// Header.
func AddImports(section *SectionTemplate, imprts ...*ImportSpec) {
	if len(imprts) == 0 {
		return
	}
	var specs []*ImportSpec
	if data, ok := section.Data.(map[string]interface{}); ok {
		if imports, ok := data["Imports"]; ok {
			specs = imports.([]*ImportSpec)
		}
		data["Imports"] = append(specs, imprts...)
	}
}

const (
	headerT = `{{if .Title}}// Code generated by goa {{.ToolVersion}}, DO NOT EDIT.
//
// {{.Title}}
//
// Command:
{{comment commandLine}}

{{end}}package {{.Pkg}}

{{if .Imports}}import {{if gt (len .Imports) 1}}(
{{end}}{{range .Imports}}	{{.Code}}
{{end}}{{if gt (len .Imports) 1}})
{{end}}
{{end}}`
)

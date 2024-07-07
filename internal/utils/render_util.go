package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func dict(values ...interface{}) map[string]interface{} {
	if len(values)%2 != 0 {
		panic(fmt.Errorf("Invalid dict call: %v", values))
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic(fmt.Errorf("Dict keys must be strings: %v", values))
		}
		dict[key] = values[i+1]
	}
	return dict
}

func (t *Templates) renderDynamicTemplate(name string, data interface{}) template.HTML {
	var buf bytes.Buffer
	err := t.templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return template.HTML(fmt.Sprintf("Error rendering template: %v", err))
	}
	return template.HTML(buf.String())
}

func NewTemplates() *Templates {
	baseTemplate := template.New("")

	myTemplate := &Templates{
		templates: baseTemplate,
	}

	funcMap := template.FuncMap{
		"dict":                  dict,
		"renderDynamicTemplate": myTemplate.renderDynamicTemplate,
	}

	baseTemplate = baseTemplate.Funcs(funcMap)

	// Parse the base template first
	baseTemplate, err := baseTemplate.ParseFiles("views/Base.html")
	if err != nil {
		panic(err)
	}

	// Then parse all other templates
	parsedTemplates, err := baseTemplate.ParseGlob("views/**/*.html")
	if err != nil {
		panic(err)
	}

	myTemplate.templates = parsedTemplates

	return myTemplate
}

func (template *Templates) Render(writer io.Writer, name string, data interface{}, context echo.Context) error {
	context.Logger().Infof("Rendering template: %s", name)
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		dataMap = make(map[string]interface{})
	}
	dataMap["Layout"] = name

	if context.Request().Header.Get("Accept") == "text/vnd.turbo-stream.html" {
		var buf bytes.Buffer
		err := template.templates.ExecuteTemplate(&buf, name, dataMap)
		if err != nil {
			context.Logger().Errorf("Error rendering turbo stream template %s: %v", name, err)
			return err
		}
		writer.Write([]byte(buf.String()))
		return nil
	}

	err := template.templates.ExecuteTemplate(writer, "base", dataMap)
	if err != nil {
		context.Logger().Errorf("Error rendering template %s: %v", name, err)
	}
	return err
}

func RenderPage(context echo.Context, title, contentBlock, layout string) error {
	data := map[string]interface{}{
		"Title":        title,
		"ContentBlock": contentBlock,
		"echoContext":  context,
	}
	return context.Render(http.StatusOK, layout, data)
}

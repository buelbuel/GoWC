package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Templates is a struct that holds the template environment.
type Templates struct {
	templates *template.Template
}

// dict is a helper function that creates a map from a list of key-value pairs.
// This function is particularly useful in templates where creating maps on the fly is not possible.
// It allows for dynamic creation of data structures within templates, used especially for component
// rendering.
//
// Parameters:
//   - values: A variadic list of interface{} values, expected to be in key-value pairs.
//
// Returns:
//   - A map[string]interface{} containing the key-value pairs.
//
// Panics if the number of arguments is odd or if a key is not a string.
func dict(values ...interface{}) map[string]interface{} {
	if len(values)%2 != 0 {
		panic(fmt.Errorf("invalid dict call: %v", values))
	}

	dict := make(map[string]interface{}, len(values)/2)

	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic(fmt.Errorf("dict keys must be strings: %v", values))
		}
		dict[key] = values[i+1]
	}

	return dict
}

// renderDynamicTemplate is a method of Templates that renders a dynamic template.
// This is useful for rendering nested templates or partial views within a larger template.
//
// Parameters:
//   - name: The name of the template to render.
//   - data: The data to pass to the template.
//
// Returns:
//   - template.HTML containing the rendered template.
//
// If there's an error during rendering, it returns an error message wrapped in HTML.
func (templ *Templates) renderDynamicTemplate(name string, data interface{}) template.HTML {
	var buf bytes.Buffer
	err := templ.templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return template.HTML(fmt.Sprintf("Error rendering template: %v", err))
	}
	return template.HTML(buf.String())
}

// NewTemplates creates and initializes a new Templates instance.
// This function sets up the template environment, including:
//   - Creating a base template
//   - Setting up custom template functions
//   - Parsing the base template file
//   - Parsing all other template files in the views directory
//
// The function uses a specific directory structure:
//   - Base template: "resources/views/Base.html"
//   - Other templates: "resources/views/**/*.html"
//
// Returns:
//   - A pointer to a new Templates instance.
//
// Panics if there's an error parsing the templates.
func NewTemplates() *Templates {
	baseTemplate := template.New("")

	viewTemplate := &Templates{
		templates: baseTemplate,
	}

	funcMap := template.FuncMap{
		"dict":                  dict,
		"renderDynamicTemplate": viewTemplate.renderDynamicTemplate,
	}

	baseTemplate = baseTemplate.Funcs(funcMap)

	// Parse the base template first
	baseTemplate, err := baseTemplate.ParseFiles("resources/views/Base.html")
	if err != nil {
		panic(err)
	}

	// Then parse all other templates
	parsedTemplates, err := baseTemplate.ParseGlob("resources/views/**/*.html")
	if err != nil {
		panic(err)
	}

	viewTemplate.templates = parsedTemplates

	return viewTemplate
}

// Render is the implementation of echo.Renderer interface.
// It renders a template with the given name and data.
//
// Parameters:
//   - writer: The io.Writer to write the rendered template to.
//   - name: The name of the template to render.
//   - data: The data to pass to the template.
//   - context: The echo.Context for the current request.
//
// Returns:
//   - An error if the template rendering fails.
//
// This method adds the "Layout" key to the data map, which can be used
// in the base template to determine which content block to render.
func (templ *Templates) Render(writer io.Writer, name string, data interface{}, context echo.Context) error {
	context.Logger().Infof("Rendering template: %s", name)
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		dataMap = make(map[string]interface{})
	}

	dataMap["Layout"] = name

	err := templ.templates.ExecuteTemplate(writer, "base", dataMap)
	if err != nil {
		context.Logger().Errorf("Error rendering template %s: %v", name, err)
	}

	return err
}

// RenderPage is a helper function to render a full page with a consistent structure.
// It sets up the data for rendering, including the page title, content block, and layout.
//
// Parameters:
//   - context: The echo.Context for the current request.
//   - title: The title of the page.
//   - description: The description of the page.
//   - contentBlock: The name of the content block to render within the layout.
//   - layout: The name of the layout template to use.
//
// Returns:
//   - An error if the rendering fails.
//
// This function is typically used in route handlers to render full pages with a consistent structure.
func RenderPage(context echo.Context, title, description, contentBlock, layout string) error {
	data := map[string]interface{}{
		"Title":        title,
		"Description":  description,
		"ContentBlock": contentBlock,
		"echoContext":  context,
	}

	return context.Render(http.StatusOK, layout, data)
}

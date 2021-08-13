package swagger

import (
	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/long2ice/fastgo/router"
	"net/http"
	"reflect"
	"time"
)

type Swagger struct {
	Title          string
	Description    string
	Version        string
	DocsUrl        string
	RedocUrl       string
	OpenAPIUrl     string
	Routers        map[string]map[string]*router.Router
	Servers        openapi3.Servers
	TermsOfService string
	Contact        *openapi3.Contact
	License        *openapi3.License
}

func Default(title, description, version string, options ...Option) *Swagger {
	swagger := &Swagger{Title: title, Description: description, Version: version, DocsUrl: "/docs", RedocUrl: "/redoc", OpenAPIUrl: "/openapi.json"}
	for _, option := range options {
		option(swagger)
	}
	return swagger
}
func (swagger *Swagger) getParametersByModel(model interface{}) openapi3.Parameters {
	parameters := openapi3.NewParameters()
	type_ := reflect.TypeOf(model).Elem()
	value_ := reflect.ValueOf(model).Elem()
	for i := 0; i < type_.NumField(); i++ {
		field := type_.Field(i)
		value := value_.Field(i)
		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			panic(err)
		}
		parameter := &openapi3.Parameter{}
		var schema *openapi3.Schema
		switch value.Interface().(type) {
		case int, int8, int16:
			schema = openapi3.NewIntegerSchema()
		case int32:
			schema = openapi3.NewInt32Schema()
		case int64:
			schema = openapi3.NewInt64Schema()
		case string:
			schema = openapi3.NewStringSchema()
		case time.Time:
			schema = openapi3.NewDateTimeSchema()
		}
		formTag, err := tags.Get(QUERY)
		if err == nil {
			parameter.In = QUERY
			parameter.Name = formTag.Name
		}
		nameTag, err := tags.Get(DESCRIPTION)
		if err == nil {
			parameter.Description = nameTag.Name
		}
		bindingTag, err := tags.Get(BINDING)
		if err == nil {
			parameter.Required = bindingTag.Name == "required"
		}
		defaultTag, err := tags.Get(DEFAULT)
		if err == nil {
			schema.Default = defaultTag.Name
			parameter.Schema = &openapi3.SchemaRef{
				Value: schema,
			}
		}
		parameters = append(parameters, &openapi3.ParameterRef{
			Value: parameter,
		})
	}

	return parameters
}
func (swagger *Swagger) paths() openapi3.Paths {
	paths := make(openapi3.Paths)
	for path, m := range swagger.Routers {
		pathItem := &openapi3.PathItem{}
		for method, r := range m {
			model := r.Model
			if method == http.MethodGet {
				pathItem.Get = &openapi3.Operation{
					Tags:        r.Tags,
					Summary:     r.Summary,
					Description: r.Description,
					Deprecated:  r.Deprecated,
					Responses:   openapi3.NewResponses(),
					Parameters:  swagger.getParametersByModel(model),
				}
			}
		}
		paths[path] = pathItem
	}
	return paths
}
func (swagger *Swagger) MarshalJSON() ([]byte, error) {
	t := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:          swagger.Title,
			Description:    swagger.Description,
			TermsOfService: swagger.TermsOfService,
			Contact:        swagger.Contact,
			License:        swagger.License,
			Version:        swagger.Version,
		},
		Paths:   swagger.paths(),
		Servers: swagger.Servers,
	}
	return t.MarshalJSON()
}

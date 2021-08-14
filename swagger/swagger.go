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
	OpenAPI        *openapi3.T
}

func Default(title, description, version string, options ...Option) *Swagger {
	swagger := &Swagger{Title: title, Description: description, Version: version, DocsUrl: "/docs", RedocUrl: "/redoc", OpenAPIUrl: "/openapi.json"}
	for _, option := range options {
		option(swagger)
	}
	return swagger
}
func (swagger *Swagger) getSchemaByType(t interface{}) *openapi3.Schema {
	var schema *openapi3.Schema
	var m float64
	m = float64(0)
	switch t.(type) {
	case int, int8, int16:
		schema = openapi3.NewIntegerSchema()
	case uint, uint8, uint16:
		schema = openapi3.NewIntegerSchema()
		schema.Min = &m
	case int32:
		schema = openapi3.NewInt32Schema()
	case uint32:
		schema = openapi3.NewInt32Schema()
		schema.Min = &m
	case int64:
		schema = openapi3.NewInt64Schema()
	case uint64:
		schema = openapi3.NewInt64Schema()
		schema.Min = &m
	case string:
		schema = openapi3.NewStringSchema()
	case time.Time:
		schema = openapi3.NewDateTimeSchema()
	case float32, float64:
		schema = openapi3.NewFloat64Schema()
	case bool:
		schema = openapi3.NewBoolSchema()
	case []byte:
		schema = openapi3.NewBytesSchema()
	default:
		schema = openapi3.NewStringSchema()
	}
	return schema
}
func (swagger *Swagger) getRequestBodyByModel(model interface{}) *openapi3.RequestBodyRef {
	body := &openapi3.RequestBodyRef{
		Value: openapi3.NewRequestBody(),
	}
	schema := openapi3.NewObjectSchema()
	schema.Properties = openapi3.Schemas{}
	type_ := reflect.TypeOf(model).Elem()
	value_ := reflect.ValueOf(model).Elem()
	for i := 0; i < type_.NumField(); i++ {
		field := type_.Field(i)
		value := value_.Field(i)
		fieldSchema := swagger.getSchemaByType(value.Interface())
		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			panic(err)
		}
		formTag, err := tags.Get(FORM)
		if err != nil {
			continue
		}
		descriptionTag, err := tags.Get(DESCRIPTION)
		if err == nil {
			fieldSchema.Description = descriptionTag.Name
		}
		bindingTag, err := tags.Get(BINDING)
		if err == nil {
			if bindingTag.Name == "required" {
				schema.Required = append(schema.Required, formTag.Name)
			}
		}
		defaultTag, err := tags.Get(DEFAULT)
		if err == nil {
			fieldSchema.Default = defaultTag.Name
		}
		schema.Properties[formTag.Name] = openapi3.NewSchemaRef("", fieldSchema)
	}
	body.Value.Required = true
	body.Value.Content = openapi3.NewContentWithJSONSchema(schema)
	return body
}
func (swagger *Swagger) getParametersByModel(model interface{}) openapi3.Parameters {
	parameters := openapi3.NewParameters()
	type_ := reflect.TypeOf(model).Elem()
	value_ := reflect.ValueOf(model).Elem()
	for i := 0; i < type_.NumField(); i++ {
		field := type_.Field(i)
		value := value_.Field(i)
		schema := swagger.getSchemaByType(value.Interface())
		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			panic(err)
		}
		parameter := &openapi3.Parameter{}
		queryTag, err := tags.Get(QUERY)
		if err == nil {
			parameter.In = openapi3.ParameterInQuery
			parameter.Name = queryTag.Name
		}
		pathTag, err := tags.Get(PATH)
		if err == nil {
			parameter.In = openapi3.ParameterInPath
			parameter.Name = pathTag.Name
		}
		headerTag, err := tags.Get(HEADER)
		if err == nil {
			parameter.In = openapi3.ParameterInHeader
			parameter.Name = headerTag.Name
		}
		cookieTag, err := tags.Get(COOKIE)
		if err == nil {
			parameter.In = openapi3.ParameterInCookie
			parameter.Name = cookieTag.Name
		}
		if parameter.In == "" {
			continue
		}
		descriptionTag, err := tags.Get(DESCRIPTION)
		if err == nil {
			parameter.Description = descriptionTag.Name
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
			if r.Exclude {
				continue
			}
			model := r.API.NewModel()
			operation := &openapi3.Operation{
				Tags:        r.Tags,
				OperationID: r.OperationID,
				Summary:     r.Summary,
				Description: r.Description,
				Deprecated:  r.Deprecated,
				Responses:   openapi3.NewResponses(),
				Parameters:  swagger.getParametersByModel(model),
			}
			requestBody := swagger.getRequestBodyByModel(model)
			if method == http.MethodGet {
				pathItem.Get = operation
			} else if method == http.MethodPost {
				pathItem.Post = operation
				operation.RequestBody = requestBody
			} else if method == http.MethodDelete {
				pathItem.Delete = operation
			} else if method == http.MethodPut {
				pathItem.Put = operation
				operation.RequestBody = requestBody
			} else if method == http.MethodPatch {
				pathItem.Patch = operation
			} else if method == http.MethodHead {
				pathItem.Head = operation
			} else if method == http.MethodPatch {
				pathItem.Patch = operation
			} else if method == http.MethodOptions {
				pathItem.Options = operation
			} else if method == http.MethodConnect {
				pathItem.Connect = operation
			} else if method == http.MethodTrace {
				pathItem.Trace = operation
			}
		}
		paths[path] = pathItem
	}
	return paths
}
func (swagger *Swagger) BuildOpenAPI() {
	swagger.OpenAPI = &openapi3.T{
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
}
func (swagger *Swagger) MarshalJSON() ([]byte, error) {

	return swagger.OpenAPI.MarshalJSON()
}

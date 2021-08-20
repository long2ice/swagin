package swagger

import (
	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin/binding"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/security"
	"mime/multipart"
	"net/http"
	"reflect"
	"regexp"
	"time"
)

const (
	DEFAULT     = "default"
	BINDING     = "binding"
	DESCRIPTION = "description"
	QUERY       = "query"
	FORM        = "form"
	URI         = "uri"
	HEADER      = "header"
	COOKIE      = "cookie"
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
	SwaggerOptions map[string]interface{}
	RedocOptions   map[string]interface{}
}

func New(title, description, version string, options ...Option) *Swagger {
	swagger := &Swagger{Title: title, Description: description, Version: version, DocsUrl: "/docs", RedocUrl: "/redoc", OpenAPIUrl: "/openapi.json"}
	for _, option := range options {
		option(swagger)
	}
	return swagger
}
func (swagger *Swagger) getSecurityRequirements(securities []security.ISecurity) *openapi3.SecurityRequirements {
	securityRequirements := openapi3.NewSecurityRequirements()
	for _, s := range securities {
		provide := s.Provider()
		swagger.OpenAPI.Components.SecuritySchemes[provide] = &openapi3.SecuritySchemeRef{
			Value: s.Scheme(),
		}
		securityRequirements.With(openapi3.NewSecurityRequirement().Authenticate(provide))
	}
	return securityRequirements
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
	case *multipart.FileHeader:
		schema = openapi3.NewStringSchema()
		schema.Format = "binary"
	default:
		schema = openapi3.NewStringSchema()
	}
	return schema
}
func (swagger *Swagger) getRequestBodyByModel(model interface{}, contentType string) *openapi3.RequestBodyRef {
	body := &openapi3.RequestBodyRef{
		Value: openapi3.NewRequestBody(),
	}
	if model == nil {
		return body
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
		tag, err := tags.Get(FORM)
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
				schema.Required = append(schema.Required, tag.Name)
			}
		}
		defaultTag, err := tags.Get(DEFAULT)
		if err == nil {
			fieldSchema.Default = defaultTag.Name
		}
		schema.Properties[tag.Name] = openapi3.NewSchemaRef("", fieldSchema)
	}
	body.Value.Required = true
	if contentType == "" {
		contentType = binding.MIMEJSON
	}
	body.Value.Content = openapi3.NewContentWithSchema(schema, []string{contentType})
	return body
}
func (swagger *Swagger) getParametersByModel(model interface{}) openapi3.Parameters {
	parameters := openapi3.NewParameters()
	if model == nil {
		return parameters
	}
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
		uriTag, err := tags.Get(URI)
		if err == nil {
			parameter.In = openapi3.ParameterInPath
			parameter.Name = uriTag.Name
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
		}
		parameter.Schema = &openapi3.SchemaRef{
			Value: schema,
		}
		parameters = append(parameters, &openapi3.ParameterRef{
			Value: parameter,
		})
	}
	return parameters
}

// /:id -> /{id}
func (swagger *Swagger) fixPath(path string) string {
	reg := regexp.MustCompile("/:([0-9a-zA-Z]+)")
	return reg.ReplaceAllString(path, "/{${1}}")
}
func (swagger *Swagger) getPaths() openapi3.Paths {
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
				Responses:   r.Responses,
				Parameters:  swagger.getParametersByModel(model),
				Security:    swagger.getSecurityRequirements(r.Securities),
			}
			requestBody := swagger.getRequestBodyByModel(model, r.ContentType)
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
			} else if method == http.MethodOptions {
				pathItem.Options = operation
			} else if method == http.MethodConnect {
				pathItem.Connect = operation
			} else if method == http.MethodTrace {
				pathItem.Trace = operation
			}
		}
		paths[swagger.fixPath(path)] = pathItem
	}
	return paths
}
func (swagger *Swagger) BuildOpenAPI() {
	components := openapi3.NewComponents()
	components.SecuritySchemes = openapi3.SecuritySchemes{}
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
		Servers:    swagger.Servers,
		Components: components,
	}
	swagger.OpenAPI.Paths = swagger.getPaths()
}
func (swagger *Swagger) MarshalJSON() ([]byte, error) {

	return swagger.OpenAPI.MarshalJSON()
}

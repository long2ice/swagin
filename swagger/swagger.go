package swagger

type Swagger struct {
	Title       string
	Description string
	Version     string
	DocsUrl     string
	RedocUrl    string
}

func Default(title, description, version string) *Swagger {
	return &Swagger{Title: title, Description: description, Version: version, DocsUrl: "/docs", RedocUrl: "/redoc"}
}

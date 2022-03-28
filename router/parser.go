package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	_ "unsafe"
)

var Query = queryBinding{}

func CookiesParser(c *gin.Context, model interface{}) error {
	params := make(map[string][]string)
	for _, cookie := range c.Request.Cookies() {
		params[cookie.Name] = append(params[cookie.Name], cookie.Value)
	}
	return copier.Copy(model, params)
}

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

//go:linkname mapFormByTag github.com/gin-gonic/gin/binding.mapFormByTag
func mapFormByTag(ptr interface{}, form map[string][]string, tag string) error

func (queryBinding) Bind(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	if err := mapFormByTag(obj, values, "query"); err != nil {
		return err
	}
	return nil
}

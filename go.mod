module github.com/long2ice/swagin

go 1.16

require (
	github.com/fatih/structtag v1.2.0
	github.com/getkin/kin-openapi v0.72.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace github.com/gin-gonic/gin v1.7.4 => github.com/long2ice/gin v1.7.2-0.20210925064857-6665ad834758

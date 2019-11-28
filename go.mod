module github.com/varandrew/gin-product

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.48.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/shirou/gopsutil v2.19.10+incompatible
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20191125084936-ffdde1057850 // indirect
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e // indirect
	golang.org/x/tools v0.0.0-20191126055441-b0650ceb63d9 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.2 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	github.com/varandrew/gin-product/app/middlewares => ./app/middlewares
	github.com/varandrew/gin-product/app/models => ./app/models
	github.com/varandrew/gin-product/app/pkg/setting => ./app/pkg/setting
	github.com/varandrew/gin-product/app/routers => ./app/routers
	github.com/varandrew/gin-product/conf => ./app/pkg/conf
)

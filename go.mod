module github.com/varandrew/gin-product

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/astaxie/beego v1.12.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.48.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/shirou/gopsutil v2.19.10+incompatible
	github.com/unknwon/com v1.0.1
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	github.com/varandrew/gin-product/app/middlewares => ./app/middlewares
	github.com/varandrew/gin-product/app/models => ./app/models
	github.com/varandrew/gin-product/app/pkg/setting => ./app/pkg/setting
	github.com/varandrew/gin-product/app/routers => ./app/routers
	github.com/varandrew/gin-product/conf => ./app/pkg/conf
)

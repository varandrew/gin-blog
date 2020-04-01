package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/varandrew/gin-product/app/models"
	"github.com/varandrew/gin-product/app/pkg/logging"
	"github.com/varandrew/gin-product/app/pkg/setting"
	"github.com/varandrew/gin-product/app/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,       // 监听的TCP地址
		Handler:        routersInit,    // http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
		ReadTimeout:    readTimeout,    // 允许读取的最大时间
		WriteTimeout:   writeTimeout,   // 允许写入的最大时间
		MaxHeaderBytes: maxHeaderBytes, // 请求头的最大字节数
	}

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("[info] Start to listening the incoming requests on http address: %s", endPoint)
	log.Printf(server.ListenAndServe().Error())

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:" + strconv.Itoa(setting.ServerSetting.HttpPort) + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

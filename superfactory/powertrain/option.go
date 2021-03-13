package powertrain

import (
	"log"
	"net/http"
	"time"
)

type Options struct {
	// 服务初始化时候优先执行的方法
	// 这里进行全局的 panic 处理
	InitFunc func()
	RegisterRouter func(mux *http.ServeMux)
	HttpServerOptions *http.Server
}

type HttpServerOptions struct {
	Addr string
	ReadTimeout time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout      time.Duration
}

func defaultOptions() *Options {
	initFunc := func() {
		if err := recover(); err != nil {
			log.Printf("server Initfunc: %s", err)
		}
	}

	return &Options{
		InitFunc: initFunc,
		HttpServerOptions: &http.Server{
			Addr:              ":8080",
			ReadTimeout:       30 * time.Second,
			ReadHeaderTimeout: 30 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       30 * time.Second,
		},
	}
}

package powertrain

import (
	"amusingx.fit/amusingx/superfactory/basicmatter"
	"amusingx.fit/amusingx/superfactory/planet"
)

type Option func(o *Options)

func Run(configObj basicmatter.Master, options ...Option) {

	 err := basicmatter.New().Unmarshal(configObj, "")
	 if err != nil {
	 	panic(err)
	 }

	 opts := defaultOptions()
	 for _, o := range options {
		o(opts)
	 }
	 opts.InitFunc()

	 // http 初始化，路由注册，监听端口
	 opts.HttpServerOptions.Addr = configObj.GetAddr()
	 planet.Init().HttpServerInit(opts.HttpServerOptions, opts.RegisterRouter)
}
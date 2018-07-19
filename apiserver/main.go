package main

import (
	"github.com/gin-gonic/gin"
	"./router"
	"./router/middleware"
	"net/http"
	"./config"
	"./model"
	"github.com/lexkong/log"
	"time"
	"errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func pingServer() error {

	for i := 0;i < viper.GetInt("max_ping_count") ; i++  {

		resp,err :=  http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)

	}
	return errors.New("Cannot connect to the router.")


}


func main() {


	pflag.Parse()

	if err := config.Init(*cfg);err != nil {
		panic(err)
	}


	//初始化数据库

	model.DB.Init()
	defer model.DB.Close()

    // 设置模式
	gin.SetMode(viper.GetString("runmode"))

    // 实例化
	g := gin.New()

	// 加载中间件
	router.Load(
		g,
		middleware.Logging(),
		middleware.RequestId(),
	)


	go func() {

		if err := pingServer(); err != nil  {

			log.Fatal("The router has no response, or it might took too long to start up.", err)

		}
		log.Info("The router has been deployed successfully.")

	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
	
}

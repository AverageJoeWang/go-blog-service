package main

import (
	"go-blog-service/global"
	"go-blog-service/internal/model"
	"go-blog-service/internal/routers"
	"go-blog-service/pkg/logger"
	"go-blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init setupLogger err: %v", err)
	}
}

func main() {
	// r := gin.Default()
	// r.GET("/PING", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"mess": "pong"})
	// })
	// r.Run()
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	//err = s.ReadSection("JWT", &global.JWTSetting)
	//if err != nil {
	//	return err
	//}
	//err = s.ReadSection("Email", &global.EmailSetting)
	//if err != nil {
	//	return err
	//}

	global.AppSetting.DefaultContextTimeout *= time.Second
	//global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	//if port != "" {
	//	global.ServerSetting.HttpPort = port
	//}
	//if runMode != "" {
	//	global.ServerSetting.RunMode = runMode
	//}

	return nil
}
//初始化数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

//初始化logger
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
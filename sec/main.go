package main

import (
	//"rabbitgather/article/global"
	"fmt"
	"github.com/meowalien/rabbitgather-lib/math"
)

var IsDebug = ""
//func init()  {
//	initConfig()
//	global.Logger = logger.NewLoggerWrapper()
//
//}

//var ServePath *url.URL
//func init() {
//	type Config struct {
//		ServePath string
//	}
//	var config Config
//	err := json.ParseFileJsonConfig(&config, "config/article.config.json")
//	if err != nil {
//		panic(err.Error())
//	}
//	ServePath, err = url.Parse(config.ServePath)
//	if err != nil {
//		panic(err.Error())
//	}
//}
//
//
//type ArticleServices struct {
//	serverInst *http.Server
//	ginEngine  *gin.Engine
//}
//
//var log *logger.LoggerWrapper
//
//func init() {
//	//logger.LogLevelMask  = logger.ALL
//	log  = logger.NewLoggerWrapper("ArticleServices","log/")
//}
//
//func (s *ArticleServices) Startup(ctx context.Context) error {
//	addr :=   ":" + ServePath.Port()
//	log.DEBUG.Printf("ArticleServices listen on : \"%s\"\n",addr)
//
//	s.ginEngine = gin.Default()
//	s.serverInst = &http.Server{
//		Addr:   addr,
//		Handler: s.ginEngine,
//		TLSConfig: &tls.Config{
//			ClientAuth: tls.NoClientCert,
//		},
//	}
//
//	s.MountService(ctx)
//	go func() {
//		if err := s.serverInst.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.ERROR.Println(err.Error())
//		}
//	}()
//	log.DEBUG.Println("ArticleServices Started .")
//	return nil
//}
//
//func (s *ArticleServices) MountService(ctx context.Context) {
//	s.ginEngine.NoRoute(func(c *gin.Context) {
//		c.File("assets/index.html")
//	})
//}

func main() {
	fmt.Println("IsDebug: ",IsDebug)
	fmt.Println("round: ",math.Round(0.5))
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//webserver := ArticleServices{}
	//err := webserver.Startup(ctx)
	//if err != nil {
	//	cancel()
	//	panic(err.Error())
	//}
	//
	//quitSignal := make(chan os.Signal)
	//signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	//select {
	//case <-ctx.Done():
	//	fmt.Println("Shutdown with Context done")
	//case <-quitSignal:
	//	cancel()
	//	fmt.Println("Shutdown with OS QuitSignal")
	//}
}

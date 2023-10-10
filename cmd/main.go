package main

import (
	"database/sql"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/rn-consider/compuBackend/cmd/initialize"
	"github.com/rn-consider/compuBackend/config"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func runServer(r *gin.Engine) {
	address := ":9785"
	s := initServer(address, r)

	if err := s.ListenAndServe(); err != nil {
		config.GVA_LOG.Error(err.Error())
	}
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
func main() {
	// 初始化配置读取
	config.GVA_VP = initialize.Viper()

	// 初始化日志
	config.GVA_LOG = initialize.Zap()
	zap.ReplaceGlobals(config.GVA_LOG)

	// 初始化数据库
	config.GVA_DB = initialize.Gorm() // gorm连接数据库
	if config.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := config.GVA_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				zap.L().Error("数据库关闭失败", zap.Error(err)) // 使用 Zap 打印数据库关闭失败的错误信息
			}
		}(db)
	} else {
		zap.L().Error("数据库启动失败，是否创建了指定数据库？...") // 使用 Zap 打印数据库关闭失败的错误信息
		return
	}

	// 初始化Router
	router := initialize.Routers()
	if router == nil {
		return
	}

	runServer(router)

	//err := dao.InitMySQL()
	//file, err := os.Create("./issucc ess.txt")
	//if err != nil {
	//	panic(err)
	//}
	//file.Write([]byte("success"))
	//
	//defer file.Close()
	//defer dao.Close()
	//dao.DB.AutoMigrate()
	//user1 := models.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//models.CreateUser(&user1)
	//models.DeleteUser(int(user1.ID))
	/*
		//增加数据做法
		user1 := models.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		models.CreateUser(&user1)
	*/
	/*
		//删除数据做法
		models.DeleteUser(1)
	*/
	/*
		//查找数据做法
		allUser, _ := models.GetAllUser()
		//要打印复杂对象不能直接使用fmt,下面这种做法将结构体列表遍历后将结构体作为json格式输出
		for _, value := range allUser {
			// json encoding
			fmt.Printf("---\njson encoding\n")
			jsonData, err := json.Marshal(&value)
			if err != nil {
			log.Fatal(err)
			}
			fmt.Println(string(jsonData))
		}
	*/
	/*
		//更新user做法
		user1 := models.User{ID: 1, Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		models.UpdatedAtUser(&user1)
	*/
}

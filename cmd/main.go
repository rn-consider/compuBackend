package main

import (
	"github.com/rn-consider/compuBackend/dao"
	"os"
)

func main() {
	err := dao.InitMySQL()
	file, err := os.Create("./issucc ess.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("success"))

	defer file.Close()
	defer dao.Close()
	dao.DB.AutoMigrate()
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

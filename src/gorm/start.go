package main

import (
	"debug/dwarf"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"k8s.io/klog"
	"strconv"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

type User struct {
	Name string
	Age int
	ID int
	CreateAt dwarf.Data
}

func (p *Product) BeforeCreate(tx * gorm.DB) (err error) {
	fmt.Println(p.Code)
	if p.Code == "jinzhu2" {
		return errors.New("invalid p.code")
	}
	return 
}

func main()  {
	dsn:="root:YANGq@123@tcp(118.31.68.29:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("failed connect db")
	}

	product := Product{Code: "D50", Price: 101}
	result := db.Create(&product)

	fmt.Println(product.Code)
	fmt.Println(product.Price)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	db.Select("Code", "Price").Create(&product)

	var users = []Product{{Code: "jinzhu1"}, {Code: "jinzhu2"}, {Code: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}



	fmt.Println("succeed")
}

func connectMysql() {
	klog.Info("init mysql")
	dbSource, logMode, err := GetMysqlDBSource()
	if err != nil{
		panic("error")
	}
	db, err := gorm.Open(sqlite.Open(dbSource))
	if err != nil{
		panic("error")
	}

	fmt.Println(logMode)

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"});
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.Delete(&product, 1)
}

func GetMysqlDBSource() (dbSource, logMode string, err error)  {
	host := "127.0.0.1"
	port, err := strconv.Atoi("3306")
	if err != nil{
		return "","",err
	}
	db:="local"
	user:="root"
	password:="090909"

	dbSource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", user, password, host, port, db)
	klog.Info("mysql datasource: %s", dbSource)
	logMode = "debug"
	return dbSource, logMode, nil
}


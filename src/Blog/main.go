package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//数据库指针
var db *gorm.DB

//用户表名
var tablename string

//用户信息
type Userinfo struct {
	Id       int64  `gorm:"primary_key"`
	Username string `form:"username"` //不能缩进妈耶
	Password string `form:"password"`
}

//文章信息
type Artcle struct {
	Id      int64  `gorm:"primary_key"`
	Author  string `gorm:"type:varchar(30)"`
	Title   string `gorm:"type:varchar(50)" form:"title"`
	Content string `gorm:"type:longtext" form:"content"`
}

//错误处理
func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//首页面
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//展示登录页面
func loginpage(c *gin.Context) {
	if _, err := c.Cookie("username"); err == nil {
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

//处理登录请求
func login(c *gin.Context) {
	var Use Userinfo
	var Res Userinfo
	//绑定结构体
	c.ShouldBind(&Use)
	// fmt.Println("Username is ", Use.Username)

	db.Debug().Where("Username = ? and Password = ?", Use.Username, Use.Password).Find(&Res)

	if Use.Username == "" || Use.Password == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "用户名和密码不能为空！",
		})
	} else if Res.Username != Use.Username {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "用户名或密码错误!",
		})
	} else {
		c.SetCookie("username", Use.Username, 60*60, "/", "127.0.0.1", false, true)
		tablename = string(Res.Username + "_" + strconv.FormatInt(Res.Id, 10))
		c.SetCookie("tablename", tablename, 60*60, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}

//展示注册页面
func signuppage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

//处理注册请求
func signup(c *gin.Context) {
	var Use Userinfo
	var Res Userinfo
	//绑定结构体
	c.ShouldBind(&Use)
	fmt.Println("Username is ", Use.Username)

	db.Debug().Where("Username = ?", Use.Username).Find(&Res)

	//用户名已被占用
	if Res.Username == Use.Username && Res.Username != "" {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"err": "用户名已被占用！",
		})
	} else if Use.Username == "" || Use.Password == "" { //用户名和密码为空

		c.HTML(http.StatusOK, "signup.html", gin.H{
			"err": "用户名和密码不能为空！",
		})
	} else {
		db.Create(&Use)
		db.Debug().Where("Username = ?", Use.Username).Find(&Res)
		db.Debug().Table(string(Res.Username + "_" + strconv.FormatInt(Res.Id, 10))).CreateTable(&Artcle{})
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

// 登录中间件
func alreadylogin(c *gin.Context) {
	var err error
	tablename, err = c.Cookie("tablename")
	//没有登录
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	c.Next()
}

//展示添加文章页面
func addpage(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", nil)
}

//处理添加文章请求
func addartcle(c *gin.Context) {
	var Artc Artcle
	c.ShouldBind(&Artc)
	Author, err := c.Cookie("username")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	fmt.Println(Artc.Title)
	Artc.Author = Author
	db.Debug().Table(tablename).Create(&Artc)

	c.Redirect(http.StatusMovedPermanently, "/home")
}

//用户主页
func home(c *gin.Context) {
	//获取用户文章
	var artcles []Artcle

	db.Debug().Table(tablename).Find(&artcles)

	for _, val := range artcles {
		fmt.Println(val.Title)
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"artcles": artcles,
	})
}

//展示修改页面
func modifypage(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	checkerr(err)

	var Res Artcle

	db.Debug().Table(tablename).Where("id = ?", id).Find(&Res)

	c.HTML(http.StatusOK, "modify.html", gin.H{
		"artcle": Res,
	})
}

//处理修改请求
func modifyartcle(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	checkerr(err)

	var Modart Artcle
	c.ShouldBind(&Modart)

	db.Debug().Table(tablename).Where("id = ?", id).Update(Modart)

	c.Redirect(http.StatusMovedPermanently, "/home")
}

//处理删除请求
func delete(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	checkerr(err)
	db.Debug().Table(tablename).Where("id = ?", id).Delete(&Artcle{})

	c.Redirect(http.StatusMovedPermanently, "/home")
}

func main() {
	//连接本地数据库
	var err error
	db, err = gorm.Open("mysql", "root:439956461@(127.0.0.1:3306)/user")
	checkerr(err)
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("./Web/source/*")
	r.Static("/src", "./Web/source/")
	r.Static("/pic", "./Web/images")

	//处理错误访问
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "wrong.html", nil)
	})

	r.GET("/", index)
	r.GET("/login", loginpage)
	r.POST("/login", login)
	r.GET("/signup", signuppage)
	r.POST("/signup", signup)
	r.GET("/home", alreadylogin, home)
	r.GET("/addartcle", alreadylogin, addpage)
	r.POST("/addartcle", addartcle)
	r.GET("/modifyartcle", alreadylogin, modifypage)
	r.POST("/modifyartcle", modifyartcle)
	r.GET("/delete", delete)

	r.Run("127.0.0.1:8080")
}

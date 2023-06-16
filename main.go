package main

import (
	"context"
	"fmt"
	"golang/connection"
	"golang/middleware"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Blog struct {
	ID          int
	Title       string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	NodeJs      bool
	ReactJs     bool
	NextJs      bool
	PHP         bool
	Image       string
	Author      string
	Start       string
	End         string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

// var dataBlog = []Blog{
// 	{
// 		ID:    0,
// 		Name:  "ibbn",
// 		Image: "imgg.jpg",
// 	},
// }

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	//e = echo package
	// GET/POST = run the method
	// "/" = enpoint/routing
	// helloWorld = function that will run if the routes are opened

	e.Static("/public", "public")
	e.Static("/uploads", "uploads")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/post", addproject)
	e.GET("/testimonials", testimonials)
	e.GET("/posts/:id", blogDetail)
	e.GET("/register", register)
	e.GET("/updateproject/:id", updateproject)
	e.GET("/login", loginform)

	e.POST("/update/:id", middleware.UploadFile(update))
	e.POST("/add-blog", middleware.UploadFile(addBlog))
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/addlogin", login)
	e.POST("/addregister", registerform)
	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start(":5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func home(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_blogs.id, title, start_date, end_date, description, nodejs, reactjs, nextjs, php, image, tb_users.name AS author FROM tb_blogs JOIN tb_users ON tb_blogs.authorid = tb_users.id ORDER BY tb_blogs.id DESC")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.StartDate, &each.EndDate, &each.Description, &each.NodeJs, &each.ReactJs, &each.NextJs, &each.PHP, &each.Image, &each.Author)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Start = each.StartDate.Format("2 January 2006")
		each.End = each.EndDate.Format("2 January 2006")
		result = append(result, each)

	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	datas := map[string]interface{}{
		"Blogs":        result,
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}

func contact(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	data := map[string]interface{}{
		"DataSession": userData,
	}
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addproject(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	datas := map[string]interface{}{
		"DataSession": userData,
	}

	var tmpl, err = template.ParseFiles("views/blogproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}

func testimonials(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	data := map[string]interface{}{
		"DataSession": userData,
	}

	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return tmpl.Execute(c.Response(), data)
}

func blogDetail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT tb_blogs.id, title, start_date, end_date, description, nodejs, reactjs, nextjs, php, image, tb_users.name AS author FROM tb_blogs JOIN tb_users ON tb_blogs.authorid = tb_users.id WHERE tb_blogs.id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.StartDate, &BlogDetail.EndDate, &BlogDetail.Description, &BlogDetail.NodeJs, &BlogDetail.ReactJs, &BlogDetail.NextJs, &BlogDetail.PHP, &BlogDetail.Image, &BlogDetail.Author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	BlogDetail.Start = BlogDetail.StartDate.Format("2 January 2006")
	BlogDetail.End = BlogDetail.EndDate.Format("2 January 2006")

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	datas := map[string]interface{}{
		"Blog":        BlogDetail,
		"DataSession": userData,
	}

	var tmpl, errTemplate = template.ParseFiles("views/post.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}

func addBlog(c echo.Context) error {

	title := c.FormValue("inputTitle")
	start_date := c.FormValue("inputStartDate")
	end_date := c.FormValue("inputEndDate")
	description := c.FormValue("inputDescription")

	nodejs := c.FormValue("nodejs")
	if nodejs == "true" {
		var NJS = nodejs
		var njs, err = strconv.ParseBool(NJS)
		if err == nil {
			fmt.Println(njs)
		}
	}

	reactjs := c.FormValue("reactjs")
	if reactjs == "true" {
		var RJS = reactjs
		var rjs, err = strconv.ParseBool(RJS)
		if err == nil {
			fmt.Println(rjs)
		}
	}

	nextjs := c.FormValue("nodejs")
	if nextjs == "true" {
		var NXJ = nextjs
		var nxj, err = strconv.ParseBool(NXJ)
		if err == nil {
			fmt.Println(nxj)
		}
	}

	php := c.FormValue("php")
	if php == "true" {
		var PHP = php
		var Php, err = strconv.ParseBool(PHP)
		if err == nil {
			fmt.Println(Php)
		}
	}

	image := c.Get("dataFile").(string)

	sess, _ := session.Get("session", c)

	author := sess.Values["id"].(int)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blogs (title, start_date, end_date, description, nodejs, reactjs, nextjs, php, image, authorid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, start_date, end_date, description, nodejs, reactjs, nextjs, php, image, author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID: ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blogs WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func register(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func loginform(c echo.Context) error {

	sess, _ := session.Get("session", c)
	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	var tmpl, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT *FROM tb_users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorect!", false, "/login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "password incorrect", false, "/login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}

func update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("inputTitle")
	start_date := c.FormValue("inputStartDate")
	end_date := c.FormValue("inputEndDate")
	description := c.FormValue("inputDescription")

	nodejs := c.FormValue("nodejs")
	if nodejs == "true" {
		var NJS = nodejs
		var njs, err = strconv.ParseBool(NJS)
		if err == nil {
			fmt.Println(njs)
		}
	}

	reactjs := c.FormValue("reactjs")
	if reactjs == "true" {
		var RJS = reactjs
		var rjs, err = strconv.ParseBool(RJS)
		if err == nil {
			fmt.Println(rjs)
		}
	}

	nextjs := c.FormValue("nodejs")
	if nextjs == "true" {
		var NXJ = nextjs
		var nxj, err = strconv.ParseBool(NXJ)
		if err == nil {
			fmt.Println(nxj)
		}
	}

	php := c.FormValue("php")
	if php == "true" {
		var PHP = php
		var Php, err = strconv.ParseBool(PHP)
		if err == nil {
			fmt.Println(Php)
		}
	}

	image := c.Get("dataFile").(string)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_blogs SET title=$1, start_date=$2, end_date=$3, description=$4, nodejs=$5, reactjs=$6, nextjs=$7, php=$8, image=$9 WHERE id=$10", title, start_date, end_date, description, nodejs, reactjs, nextjs, php, image, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateproject(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, description, nodejs, nextjs, reactjs, php, image FROM tb_blogs WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.StartDate, &BlogDetail.EndDate, &BlogDetail.Description, &BlogDetail.NodeJs, &BlogDetail.NextJs, &BlogDetail.ReactJs, &BlogDetail.PHP, &BlogDetail.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/updateproject.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func registerform(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_users(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Sorry your registration is failed", false, "/register")
	}

	return redirectWithMessage(c, "Congratulations!!! Your register is success", true, "/login")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")

}

package main

import (
	"context"
	"fmt"
	"golang/connection"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jackc/pgtype"
	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID           int
	Name         string
	StartDate    pgtype.Date
	EndDate      pgtype.Date
	Description  pgtype.Text
	Technologies pgtype.VarcharArray
	Image        string
}

var dataBlog = []Blog{
	{
		ID:    0,
		Name:  "ibbn",
		Image: "imgg.jpg",
	},
}

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	//e = echo package
	// GET/POST = run the method
	// "/" = enpoint/routing
	// helloWorld = function that will run if the routes are opened

	e.Static("/public", "public")
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/post", addproject)
	e.GET("/testimonials", testimonials)
	e.GET("/posts/:id", blogDetail)

	e.POST("/add-blog", addBlog)
	e.POST("/blog-delete/:id", deleteBlog)

	e.Logger.Fatal(e.Start(":5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func home(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_blogs")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Name, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		result = append(result, each)

	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addproject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/blogproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	for i, data := range dataBlog {
		if id == i {
			BlogDetail = Blog{
				Name:         data.Name,
				StartDate:    data.StartDate,
				EndDate:      data.EndDate,
				Description:  data.Description,
				Technologies: data.Technologies,
				Image:        data.Image,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, err = template.ParseFiles("views/post.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("inputTitle")
	content := c.FormValue("inputContent")

	println("Title : " + title)
	println("Content : " + content)

	var newBlog = Blog{
		ID:    0,
		Name:  "",
		Image: "",
	}

	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index :", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

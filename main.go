package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
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
	e.GET("/post/:id", blogDetail)

	e.POST("/add-blog", addBlog)

	e.Logger.Fatal(e.Start(":5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
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

	data := map[string]interface{}{
		"id":      id,
		"Title":   "Judul 1",
		"Content": "voluptate ad ipsum commodi itaque. Aut impedit molestiae delectus dicta perspiciatis debitis numquam, nihil sit ut quam exercitationem reiciendis quisquam blanditiis ex et ducimus hic eius provident repudiandae unde. Velit eius cupiditate unde ratione perferendis architecto pariatur modi officiis veniam, adipisci mollitia praesentium culpa, vel, alias at aspernatur quod optio perspiciatis nulla magnam quidem nemo sint! Earum adipisci facilis tenetur, illo obcaecati provident tempora. Reprehenderit fugiat nam ad neque laborum illum quibusdam! Aliquam ab qui voluptates quidem fugiat ratione nemo voluptas, quas atque repellendus numquam quo saepe dolores. Dolore fugiat molestias minima impedit dolor incidunt tenetur magni dicta quas harum sunt corporis alias voluptatum, rerum excepturi nulla error fuga mollitia aspernatur, quisquam, provident repellendus animi delectus consequuntur. Facere animi alias natus deleniti reprehenderit, nesciunt eum consequatur minus consectetur nostrum cupiditate explicabo dolor voluptas labore laborum culpa consequuntur nam dolorum obcaecati laboriosam eveniet accusamus omnis voluptate. Asperiores, beatae! A, ut excepturi dolorum ducimus nobis iusto, suscipit culpa ipsam dolorem nam voluptates voluptatibus voluptatum quo reiciendis maxime, adipisci rerum illum dolores itaque quos! Aut, expedita eius. Eveniet dolore doloremque assumenda amet quidem reprehenderit debitis distinctio suscipit rem veniam, quae doloribus asperiores molestiae incidunt sit tempora quos. Neque quam illo temporibus fugiat eveniet odio explicabo corporis rerum non at provident, accusamus, repellat earum iusto, modi tenetur reprehenderit repellendus nostrum nemo. Vero delectus pariatur animi, aliquid dolore modi maiores suscipit cupiditate amet odit, perspiciatis ratione recusandae, id libero excepturi inventore aut qui minima architecto perferendis. Maiores illum dolorum aut voluptatibus similique doloremque quam ut debitis sint alias nam atque et quo, maxime accusamus hic tempore sit tenetur. Mollitia illo ipsa recusandae suscipit illum minus aliquam beatae, ducimus doloribus eligendi dicta omnis expedita nesciunt a voluptatem vero iure molestias. Quis rem, repellendus necessitatibus maiores accusantium quisquam ipsa officiis! Unde distinctio, nihil natus repellendus repellat, quo libero consectetur laborum itaque neque quos provident qui. Inventore mollitia maiores rem perspiciatis incidunt, atque tempore, assumenda natus illo obcaecati sint!",
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

	return c.Redirect(http.StatusMovedPermanently, "/")
}

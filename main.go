package main

import "fmt"
import "github.com/gofiber/fiber/v2"


func main(){
	fmt.Println("hello")
	app:= fiber.New()

	log.Fatal(app.Listen(":4000"))

}

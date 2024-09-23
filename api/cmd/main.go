package main

import (
	"cash/sample-cash/api/controller"
	"cash/sample-cash/api/infra"
	"cash/sample-cash/api/router"
	"cash/sample-cash/api/usecase"
	"log"
)

func main() {

	db := infra.NewDB()
	infra := infra.NewInfra(db)
	usecase := usecase.NewUsecase(infra)
	c := controller.NewController(usecase)

	e := router.NewRouter(c)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

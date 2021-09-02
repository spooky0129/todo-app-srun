package main

import (
	"fmt"
	"todo-app-srun/app/controllers"
	"todo-app-srun/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}

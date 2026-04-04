package main 

import (

	"github.com/phantom5803/ticket-booking-project/internal/routes"
    "github.com/phantom5803/ticket-booking-project/pkg/database"

    "github.com/gin-gonic/gin"
)


func main(){
//initialize database connection
database.ConnectDB()

r := gin.Default()

routes.SetupRoutes(r)

r.Run(":8080")
	
}
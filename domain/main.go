package main

import(
	"github.com/jinzhu/gorm"
	"github.com/lucasvilarinho/projeto/application/grpc"
	"github.com/lucasvilarinho/projeto/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
package main

import(
	"github.com/jinzhu/gorm"
	"github.com/codeedu/imersao/codepix-go/application/grpc"
	"github.com/codeedu/imersao/codepix-go/infrastructure/dp"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv( key: "env" ))

	grpc.StartGrpcServer(database, port:"50051")
}
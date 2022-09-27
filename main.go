package main

import (
	"context"
	"fmt"
	"log"
	"multilingual-new/models"
	"multilingual-new/pb/pb"
	"multilingual-new/server"
	"net"
	"os"

	"cloud.google.com/go/translate"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("could not load env: %v", err)
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Language{}, &models.TextContent{}, &models.Translation{})
	client, err := translate.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	server := &server.MultiLingualServer{Db: db, Client: client}
	grpcServer := grpc.NewServer()
	pb.RegisterMultiLingualServiceServer(grpcServer, server)
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started on port 50051")
	grpcServer.Serve(listener)
}

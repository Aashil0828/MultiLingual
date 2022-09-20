package main

import (
	"database/sql"
	"fmt"
	"log"
	"multilingual-new/pb/pb"
	"multilingual-new/server"
	"net"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)
var db *sql.DB

// This function will make a connection to the database only once.

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("could not load env: %v", err)
	}
	psqlconn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err = sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}

func main() {
	server := &server.Server{Db: db}
	grpcServer := grpc.NewServer()
	pb.RegisterMultiLingualServiceServer(grpcServer, server)
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("cannot start server : %v", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server : %v", err)
	}
}

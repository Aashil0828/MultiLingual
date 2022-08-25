package main

import (
	"database/sql"
	"fmt"
	"log"
	"multilingual-new/pb/pb"
	"multilingual-new/server"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Akshah089#"
	dbname   = "multilingual"
)

var db *sql.DB

// This function will make a connection to the database only once.

func init() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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

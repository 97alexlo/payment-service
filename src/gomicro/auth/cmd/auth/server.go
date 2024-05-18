package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	auth "github.com/97alexlo/gomicro/internal/implementation"
	pb "github.com/97alexlo/gomicro/proto"
	"google.golang.org/grpc"
)

const (
	dbDriver   = "mysql"
	dbUser     = "a"
	dbPassword = "p"
	dbName     = "n"
)

var db *sql.DB

func main() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)

	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing db: %s", err)
		}
	}()

	// ping the db to ensure connection is valid
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// grpc server setup
	grpcServer := grpc.NewServer()
	authServerImplementation := auth.NewAuthImplementation(db)
	pb.RegisterAuthServiceServer(grpcServer, authServerImplementation)

	// listen and serve
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

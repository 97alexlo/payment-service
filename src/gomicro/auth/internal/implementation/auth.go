package auth

import (
	"context"
	"database/sql"

	pb "github.com/97alexlo/gomicro/proto"
)

type Implementation struct {
	db *sql.DB
	pb.UnimplementedAuthServiceServer
}

func NewAuthImplementation(db *sql.DB) *Implementation {
	return &Implementation{db: db}
}

func (this *Implementation) GetToken(ctx context.Context, credentials *pb.Credentials) (*pb.Token, error) {
	// implementation
	return &pb.Token{}, nil
}

func (this *Implementation) ValidateToken(ctx context.Context, token *pb.Token) (*pb.User, error) {
	return *pb.User{}, nil
}
package server

import (
	"context"

	"github.com/gofrs/uuid"
	usersv1 "github.com/sai-bhargav/grpc-https-gateway/proto/client"
)

// Backend implements the protobuf interface
type Backend struct {
	usersv1.UnimplementedClientServiceServer
}

// New initializes a new Backend struct.
func NewBackend() usersv1.ClientServiceServer {
	return &Backend{}
}

// AddUser adds a user to the in-memory store.
func (b *Backend) AddUser(ctx context.Context, _ *usersv1.StringMessage) (*usersv1.StringMessage, error) {

	vvv := uuid.Must(uuid.NewV4()).String()
	// user := &usersv1.StringMessage{
	// 	Value: vvv,
	// }
	// b.users = append(b.users, user)

	return &usersv1.StringMessage{
		Value: vvv,
	}, nil
}

func (b *Backend) CreateMenu(ctx context.Context, _ *usersv1.StringMessage) (*usersv1.StringMessage, error) {

	vvv := uuid.Must(uuid.NewV4()).String()
	// user := &usersv1.StringMessage{
	// 	Value: vvv,
	// }
	// b.users = append(b.users, user)

	return &usersv1.StringMessage{
		Value: vvv,
	}, nil
}

package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to retrieve comment")
	ErrNotImplemented  = errors.New("not implemented yet")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// store interface
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// struct of service - all business logic here
type Service struct {
	Store Store
}

// constructor function - return a pointer to the new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// get comment
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// update comment
func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

// delete comment
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

// create comment
func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

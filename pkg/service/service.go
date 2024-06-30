package service

import "github.com/msound/example001/pkg/model"

type Service struct {
	storer model.Storer
}

func NewService(storer model.Storer) *Service {
	service := Service{storer: storer}

	return &service
}

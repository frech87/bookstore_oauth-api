package service

import (
	"github.com/frech87/bookstore_oauth-api/src/domain/access_token"
	"github.com/frech87/bookstore_oauth-api/src/repository/db"
	"github.com/frech87/bookstore_oauth-api/src/repository/rest"
	"github.com/frech87/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	//CreateService(token access_token.AccessToken) *errors.RestErr
	Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTimeService(token access_token.AccessToken) *errors.RestErr
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err = s.repository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

type service struct {
	repository   db.DbRepository
	restUserRepo rest.RestUsersRepository
}

func NewService(usersRepo rest.RestUsersRepository, repository db.DbRepository) *service {
	return &service{
		restUserRepo: usersRepo,
		repository:   repository,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

//func (s *service) CreateService(token access_token.AccessToken) *errors.RestErr {
//	if err := token.Validate(); err != nil {
//		return err
//	}
//	return s.repository.Create(token)
//}

func (s *service) UpdateExpirationTimeService(token access_token.AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(token)
}

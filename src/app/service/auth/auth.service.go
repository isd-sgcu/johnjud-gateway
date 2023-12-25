package auth

import (
	"github.com/isd-sgcu/johnjud-gateway/src/app/dto"
	auth_proto "github.com/isd-sgcu/johnjud-go-proto/johnjud/auth/auth/v1"
)

type Service struct {
	client auth_proto.AuthServiceClient
}

func NewService(client auth_proto.AuthServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Signup(signup *dto.SignupRequest) (*dto.SignupResponse, *dto.ResponseErr) {
	// call authClient.Signup()
	// handle error: conflict data, internal error and unavailable service
	return nil, nil
}

func (s *Service) SignIn(signIn *dto.SignIn) (*auth_proto.Credential, *dto.ResponseErr) {
	// call authClient.SignIn()
	// handle error: forbidden, internal error and unavailable service
	return nil, nil
}

func (s *Service) SignOut(token string) (bool, *dto.ResponseErr) {
	return false, nil
}

func (s *Service) Validate(token string) (*dto.TokenPayloadAuth, *dto.ResponseErr) {
	// call authClient.Validate()
	// handle error: unauthorized, internal error and unavailable service
	return nil, nil
}

func (s *Service) RefreshToken(token string) (*auth_proto.Credential, *dto.ResponseErr) {
	// call authClient.Validate()
	// handle error: unauthorized, internal error and unavailable service
	return nil, nil
}

package tc

import (
	"errors"
	"net/rpc"

	"github.com/sirupsen/logrus"
)

// Service interface
type Service interface {
	GenerateTokens(userID, count int) (*GenerateTokensResponse, error)
}

type service struct {
	address string
	logger  *logrus.Logger
}

// New returns a new TokenService
func New(address string, logger *logrus.Logger) Service {
	return &service{address, logger}
}

func (s *service) GenerateTokens(userID, count int) (*GenerateTokensResponse, error) {
	client, err := rpc.DialHTTP("tcp", s.address)
	if err != nil {
		return &GenerateTokensResponse{}, err
	}

	var response GenerateTokensResponse
	request := GenerateTokensRequest{
		UserID: userID,
		Count:  count,
	}

	err = client.Call("RPC.CheckUserAuth", request, &response)
	s.logger.Infof("RPC response %v", response)
	if response.Success == false {
		return &GenerateTokensResponse{}, errors.New("Could not generate tokens on in the account service")
	}
	if err != nil {
		return &GenerateTokensResponse{}, err
	}

	return &response, nil
}

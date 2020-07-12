package tc

import (
	"errors"
	"net/rpc"

	"github.com/sirupsen/logrus"
)

// Service interface
type Service interface {
	GenerateTokens(userID, productID int, email string, count int) (*GenerateTokensResponse, error)
}

type service struct {
	address string
	logger  *logrus.Logger
}

// New returns a new TokenService
func New(address string, logger *logrus.Logger) Service {
	return &service{address, logger}
}

func (s *service) GenerateTokens(userID, productID int, email string, count int) (*GenerateTokensResponse, error) {
	client, err := rpc.DialHTTP("tcp", s.address)
	if err != nil {
		return &GenerateTokensResponse{}, err
	}

	var response GenerateTokensResponse
	request := GenerateTokensRequest{
		UserID:    userID,
		ProductID: productID,
		Email:     email,
		Count:     count,
	}

	err = client.Call("RPC.GenerateTokens", request, &response)
	s.logger.Infof("RPC response %v", response)
	if response.Success == false {
		return &GenerateTokensResponse{}, errors.New("Could not generate tokens in the tokens service")
	}
	if err != nil {
		return &GenerateTokensResponse{}, err
	}

	return &response, nil
}

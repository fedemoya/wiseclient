package wise

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
	"wiseclient/pkg/wise/models"
)

const baseURL = "https://api.sandbox.transferwise.tech"

type clientTestSuite struct {
	suite.Suite

	token           string
	client          *Client
	profileID       int64
	targetAccountID int64
	url             *url.URL
}

func (s *clientTestSuite) SetupTest() {
	token, found := os.LookupEnv("WISE_TOKEN")
	s.Require().True(found, "WISE_TOKEN not defined")

	url, err := url.Parse(baseURL)
	s.Require().NoError(err, "base url is invalid")
	s.url = url

	authClient := NewAuthClient(token, http.DefaultClient)
	s.client = NewClient(url, authClient)

	profileID, found := os.LookupEnv("PROFILE_ID")
	s.Require().True(found, "PROFILE_ID not defined")

	pID, err := strconv.ParseInt(profileID, 10, 64)

	s.Require().NoError(err)

	s.profileID = pID

	targetAccountID, found := os.LookupEnv("TARGET_ACCOUNT_ID")
	s.Require().True(found, "TARGET_ACCOUNT_ID not defined")

	taID, err := strconv.ParseInt(targetAccountID, 10, 64)

	s.Require().NoError(err)

	s.targetAccountID = taID
}

func (s *clientTestSuite) TestCreateQuoteAndTransferSucceed() {

	createQuoteResponse, err := s.client.CreateQuote(
		context.Background(),
		s.profileID,
		models.NewCreateQuoteRequestBodyBuilder().
			SetSourceAmount(10).
			SetSourceCurrency("GBP").
			SetPayOut(models.QuotePayoutTypeBalance).
			SetTargetCurrency("EUR").
			Build(),
	)

	quoteId := createQuoteResponse.Id

	s.Require().NoError(err)
	s.Require().NotEmpty(quoteId)

	quoteUUID, err := uuid.Parse(quoteId)
	s.Require().NoError(err)

	createTransferResponse, err := s.client.CreateTransfer(
		context.Background(),
		models.NewCreateTransferRequestBodyBuilder().
			TargetAccount(s.targetAccountID).
			QuoteUuid(quoteUUID).
			CustomerTransactionId(uuid.New()).
			Build(),
	)

	s.Require().NoError(err)
	s.Require().NotEmpty(createTransferResponse.Id)
}

func (s *clientTestSuite) TestCreateQuoteFailsWithAuthenticationError() {

	authClient := NewAuthClient("wrongToken", http.DefaultClient)
	client := NewClient(s.url, authClient)

	_, err := client.CreateQuote(
		context.Background(),
		s.profileID,
		models.NewCreateQuoteRequestBodyBuilder().
			SetSourceAmount(10).
			SetSourceCurrency("GBP").
			SetPayOut(models.QuotePayoutTypeBalance).
			SetTargetCurrency("EUR").
			Build(),
	)

	s.Require().Error(err)

	log.Info().Msgf("create quote authentication error: %s", err.Error())

	authError, ok := err.(AuthenticationError)

	s.Require().True(ok)

	s.Require().NotEmpty(authError.Err)
	s.Require().NotEmpty(authError.ErrorDescription)
}

func (s *clientTestSuite) TestCreateQuoteFailsWithValidationError() {

	_, err := s.client.CreateQuote(
		context.Background(),
		s.profileID,
		models.NewCreateQuoteRequestBodyBuilder().
			SetSourceAmount(-1).
			SetSourceCurrency("XXX").
			SetPayOut(models.QuotePayoutTypeBalance).
			SetTargetCurrency("EUR").
			Build(),
	)

	s.Require().Error(err)

	log.Info().Msgf("create quote validation error: %s", err.Error())

	validationError, ok := err.(ValidationError)

	s.Require().True(ok)

	s.Require().NotEmpty(validationError.Errors)
}

func TestClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

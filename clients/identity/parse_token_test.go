package identity_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hightidecrm/frontegg/authenticator"
	"github.com/hightidecrm/frontegg/clients/identity"
	"github.com/hightidecrm/frontegg/config"
	"github.com/hightidecrm/frontegg/internal/http_client"
	"github.com/stretchr/testify/suite"
)

type IdentityTestSuite struct {
	suite.Suite

	identityClient *identity.IdentityClient
}

func (suite *IdentityTestSuite) SetupSuite() {
	// set StartingNumber to one
	cfg, err := config.NewFronteggConfig()
	if err != nil {
		panic(err)
	}

	auth := authenticator.NewFronteggAuthenticator(cfg)

	httpClient := http_client.NewFronteggHttpClient(auth, cfg)

	suite.identityClient = identity.NewIdentityClient(cfg, auth, httpClient)
}

func (suite *IdentityTestSuite) getTenantKey() (string, error) {
	reqUrl := fmt.Sprintf("%s/identity/resources/auth/v2/api-token", os.Getenv("FRONTEGG_BASE_URL"))
	var data = fmt.Sprintf(`{
	  "clientId": "%s",
	  "secret": "%s"
	}`, os.Getenv("FRONTEGG_TENANT_CLIENT_ID"), os.Getenv("FRONTEGG_TENANT_SECRET"))
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var body = make(map[string]interface{})

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body["access_token"].(string), nil
}

func (suite *IdentityTestSuite) TestParseTenantToken() {
	// test GetUser
	tenantKey, err := suite.getTenantKey()
	if err != nil {
		panic(err)
	}

	value, err := suite.identityClient.ParseToken(context.Background(), tenantKey)

	suite.NoError(err)

	suite.NotEmpty(value)

	if ok, _ := value.IsTenantToken(); !ok {
		suite.FailNow(fmt.Sprintf("Expected tenant token, got %t", value))
	}
}

func (suite *IdentityTestSuite) getUserKey() (string, error) {
	reqUrl := fmt.Sprintf("%s/identity/resources/auth/v2/api-token", os.Getenv("FRONTEGG_BASE_URL"))
	var data = fmt.Sprintf(`{
	  "clientId": "%s",
	  "secret": "%s"
	}`, os.Getenv("FRONTEGG_USER_CLIENT_ID"), os.Getenv("FRONTEGG_USER_SECRET"))
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var body = make(map[string]interface{})

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body["access_token"].(string), nil
}

func (suite *IdentityTestSuite) TestParseUserToken() {
	// test GetUser
	tenantKey, err := suite.getUserKey()
	if err != nil {
		panic(err)
	}

	value, err := suite.identityClient.ParseToken(context.Background(), tenantKey)

	suite.NoError(err)

	suite.NotEmpty(value)

	if ok, _ := value.IsUserToken(); !ok {
		suite.FailNow(fmt.Sprintf("Expected user token, got %t", value))
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIdentityClient(t *testing.T) {
	suite.Run(t, new(IdentityTestSuite))
}

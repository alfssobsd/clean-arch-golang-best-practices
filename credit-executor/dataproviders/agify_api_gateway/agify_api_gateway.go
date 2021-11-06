package agify_api_gateway

import (
	"clean-arch-golang-best-practices/credit-library/httpclient"
	"encoding/json"
	"go.uber.org/zap"
	"net/url"
)

const apiURL = "https://api.agify.io/"

type AgifyApiGateway struct {
	logger       *zap.SugaredLogger
}

type IAgifyApiGateway interface {
	PredicateAgeOfName(name string, country2AlphaId string)
}

func NewAgifyApiGateway(logger *zap.SugaredLogger) *AgifyApiGateway {
	return &AgifyApiGateway{logger: logger}
}

func (g *AgifyApiGateway) PredicateAgeOfName(name string, country2AlphaId string) (*PredicateAgeResponseDto, error) {
	queryParams := url.Values{}
	queryParams.Add("name", name)
	queryParams.Add("country_id", country2AlphaId)

	req, _ := httpclient.NewApiHttpRequest("GET", apiURL, queryParams, nil)
	response, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var dto = PredicateAgeResponseDto{}
	jsonDecoder := json.NewDecoder(response.Body)
	err = jsonDecoder.Decode(&dto)
	if err != nil {
		return nil, err
	}

	return &dto,nil
}

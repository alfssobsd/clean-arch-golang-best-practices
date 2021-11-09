package agify_api_gateway

import (
	"clean-arch-golang-best-practices/credit-library/httpclient"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
	"encoding/json"
	"net/url"
)

const apiURL = "https://api.agify.io/"

type AgifyApiGateway struct {
	logger *loggerhelper.CustomLogger
}

type IAgifyApiGateway interface {
	PredicateAgeOfName(ctx context.Context, name string, country2AlphaId string) (*PredicateAgeResponseDto, error)
}

func NewAgifyApiGateway(logger *loggerhelper.CustomLogger) *AgifyApiGateway {
	return &AgifyApiGateway{logger: logger}
}

func (g *AgifyApiGateway) PredicateAgeOfName(ctx context.Context, name string, country2AlphaId string) (*PredicateAgeResponseDto, error) {
	g.logger.SugarWithTracing(ctx).Debug("PredicateAgeOfName")
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

	return &dto, nil
}

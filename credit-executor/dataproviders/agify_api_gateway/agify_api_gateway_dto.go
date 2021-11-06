package agify_api_gateway

type PredicateAgeResponseDto struct {
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Count *int    `json:"count"`
}

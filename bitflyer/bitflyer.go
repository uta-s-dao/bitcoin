package bitflyer

const baseURL = "https://api.bitflyer.com.v1"

type APIClient struct {
	key         string
	secret      string
	httpClient *http.Client
}
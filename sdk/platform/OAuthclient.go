package platform

// OAuthClient holds OAuth Client details
type OAuthClient struct {
	Config *PixelbinConfig
	Token  string
}

// NewOAuthClient return OAuthClient instance
func NewOAuthClient(config *PixelbinConfig) *OAuthClient {
	return &OAuthClient{
		Config: config,
		Token:  config.ApiSecret,
	}
}

// GetAccessToken returns the access token
func (o *OAuthClient) GetAccessToken() string {
	return o.Token
}

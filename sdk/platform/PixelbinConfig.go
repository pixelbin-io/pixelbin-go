package platform

// PixelbinConfig provides configuration to a service client instance.
type PixelbinConfig struct {
	ApiSecret   string
	Domain      string
	OAuthClient *OAuthClient
}

//NewPixelbinConfig provides pixelbin configuration
func NewPixelbinConfig(apiSecret, domain string) *PixelbinConfig {
	return &PixelbinConfig{apiSecret, domain, &OAuthClient{}}
}

//SetOAuthClient sets OAuthClient into pixelbin configuration
func (p *PixelbinConfig) SetOAuthClient() {
	p.OAuthClient = NewOAuthClient(p)

}

//GetAccessToken returns the access token
func (p *PixelbinConfig) GetAccessToken() string {
	return p.OAuthClient.GetAccessToken()
}

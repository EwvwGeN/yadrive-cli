package models

type DeviceAccessCodeResponse struct {
	ExpIn           int64  `json:"expires_in"`
	Interval        int    `json:"interval"`
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUrl string `json:"verification_url"`
}

type OAuthTokenResponse struct {
	ExpIn        int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
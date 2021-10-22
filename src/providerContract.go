package src

import "github.com/ArtisanCloud/PowerSocialite/src/contracts"

type ProviderInterface interface {
	Redirect(redirectURL string) (string, error)
	User(token *contracts.AccessTokenInterface)  (*User, error)
	//UserFromCode(code string) (*User, error)
	//UserFromToken(token string) (*User, error)
}

package main

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// var (
// 	oauthConf = &clientcredentials.Config{
// 		ClientID:     "my-first-app",
// 		ClientSecret: "eabe72cc-83c4-4662-81d5-15fab5146bb3",
// 		TokenURL:     "http://localhost:8080/auth/realms/Minio/protocol/openid-connect/token",
// 		Scopes:       []string{"openid", "profile", "email", "offline_access"},
// 		//EndpointParams: "",
// 	}
// 	// random string for oauth2 API calls to protect against CSRF
// 	oauthStateString = "thisshouldberandom"
// )

var (
	oauthConf = &clientcredentials.Config{
		ClientID:     "fy2TvqkILON1nfsqL6zaLL6C0m4a",
		ClientSecret: "9Aon6fkYoEBeGBawwg4fWbqpg6Aa",
		TokenURL:     "https://localhost:9443/oauth2/token",
		Scopes:       []string{"openid", "profile", "email", "offline_access"},
		//EndpointParams: "",
	}
	// random string for oauth2 API calls to protect against CSRF
	oauthStateString = "thisshouldberandom"
)

func getAccessToken() (string, error) {
	token, err := oauthConf.Token(oauth2.NoContext)
	if err != nil {
		fmt.Println("ERROR IN Token Request")
		return "", err
	}
	return token.AccessToken, nil
}

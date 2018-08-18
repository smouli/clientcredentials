package main

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/minio/minio/pkg/auth"
)

const (
	certEndPoint = "http://localhost:8080/auth/realms/Minio/protocol/openid-connect/certs"
)

func GetMinioToken(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("GetMinioToken:ParseForm(): %v", err)
		http.Error(w, "Parse Error", 400)
	}
	//fmt.Println(r.Form)

	//Read the AccessToken from the request form
	token := r.FormValue("AccessToken")
	fmt.Printf("GetMinioToken: %s\n", token)

	//Validate the token just read
	ok, err := validateAccessToken(token)
	if !ok {
		fmt.Println("GetMinioToken:ValidateToken(): %v", err)
		http.Error(w, "Authentication Failed", 401)
	}

	// Read the credentials from the Minio json file
	cred, err := parseConfig()
	if err != nil {
		fmt.Println("GetMinioToken:ParseConfig(): %v", err)
	}

	//Marshal the credentials back to the client as JSON
	b, _ := json.Marshal(cred)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

//Parse the JWT token and validate it
func validateAccessToken(accessToken string) (bool, error) {
	rsaPubkey, err := fetchJWKKey()
	if err != nil {
		fmt.Printf("validateAccessToken:fetchJWKKey(): %v\n", err)
		return false, err
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return rsaPubkey, nil
	})
	if err != nil {
		fmt.Printf("validateAccessToken:ParseWithClaims(): %v\n", err)
		return false, err
	}
	if !token.Valid {
		return false, errors.New("Token is invalid")
	}
	fmt.Printf("Claims = %v\n", claims)
	return true, nil
}

//Makes a call to certs endpoint and returns RSA public key
func fetchJWKKey() (*rsa.PublicKey, error) {
	set, err := jwk.Fetch(certEndPoint)
	if err != nil {
		fmt.Printf("fetchJWKKey: jwk.Fetch(): %v\n", err)
		return nil, err
	}
	key, err := set.Keys[0].Materialize()
	if err != nil {
		fmt.Printf("Materialize: %v\n", err)
		return nil, err
	}
	rsaPubkey := key.(*rsa.PublicKey)
	//fmt.Println("RSA PUB KEY IS: %v", rsaPubkey)
	return rsaPubkey, nil

}
func parseConfig() (*auth.Credentials, error) {
	content, _ := ioutil.ReadFile("/Users/sanatmouli/.minio/config.json")
	var result map[string]interface{}
	json.Unmarshal(content, &result)
	cred := result["credential"].(map[string]interface{})
	credmap := make(map[string]string, 2)
	for key, value := range cred {
		credmap[key] = value.(string)
		//fmt.Println(key, value.(string))
	}
	authcred := &auth.Credentials{
		AccessKey: credmap["accessKey"],
		SecretKey: credmap["secretKey"],
	}

	//fmt.Printf("accessKey: %s, secretKey: %s\n", accessKey, secretKey)
	//return accessKey, secretKey, nil
	return authcred, nil
}

func main() {
	http.HandleFunc("/getminiotoken", GetMinioToken)
	fmt.Println("Listening on port 4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}

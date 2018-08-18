package main

import (
	"encoding/json"
	"github.com/minio/minio/pkg/auth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getMinioCred(accessToken string) (*auth.Credentials, error) {
	minioTokenUrl := "http://localhost:4000"
	resource := "/getminiotoken"
	u, err := url.ParseRequestURI(minioTokenUrl)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := u.String()
	data := url.Values{}
	//accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ4SjZqaWFPN0l2SUxrR05iWVU1TWpraUxPVXdkT2RGSFNKUGozQnpDYl9BIn    0.eyJqdGkiOiIxNzU3YWE0MC00ZmQyLTRhMDgtYWEyNS1lNGRmNzk0M2ZlYzIiLCJleHAiOjE1Mjg0MTAxNzIsIm5iZiI6MCwiaWF0IjoxNTI4NDA5ODcyLCJpc3MiOiJo    dHRwOi8vbG9jYWxob3N0OjgwODAvYXV0aC9yZWFsbXMvZGVtbyIsImF1ZCI6Im15LWZpcnN0LWFwcCIsInN1YiI6IjViNWRiNGRjLWMyOTYtNDdhOC1hN2EzLTJlM2QxNW    Y2N2E4YSIsInR5cCI6IkJlYXJlciIsImF6cCI6Im15LWZpcnN0LWFwcCIsImF1dGhfdGltZSI6MTUyODQwOTg3Miwic2Vzc2lvbl9zdGF0ZSI6ImZjMjIxNmRjLWM3NmYt    NDZlZS05YTgwLTQ3ZTYxMGJiYjUyNiIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOltdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsidW1hX2F1dGhvcml6YXRpb2    4iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUi    XX19LCJuYW1lIjoiYXNob2sgbW91bGkiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhc2hvay1kZW1vIiwiZ2l2ZW5fbmFtZSI6ImFzaG9rIiwiZmFtaWx5X25hbWUiOiJtb3    VsaSIsImVtYWlsIjoiYXNob2ttb3VsaUB5YWhvby5jb20ifQ.Dt5XBW2BzMNZJiaFN1Y-wOpyQgDDchH7aoMNtAGam8cprlwJ_h-1hEV1zyttf8SkQewUmd2mUEPgbWjQF    zUbkV5fEyxyKGYzxV2B5Ptn48HG9nkeNN79N1p2QUzIVg-ZhdcsTNM4ye7cZF6MrnHd1ogtE_1T-rH2PFe_V0jlS05uTMGzenPe5Lf1s6vHV7I-H_GLvUk7LQUVxxz4u2U    vFw6ZO7WxPRM2NVMm_8fRIP4WJBilktfuhkEc4gqtWGaOenpO_i1xpUI2ZNybbHb4VrlNHua1c8R3mGJb1MQkXF79wu-qdixTivN2DUe3owJ0lIypy_z_cU6ngFqRz8mYpg"
	data.Add("AccessToken", accessToken)

	client := &http.Client{}
	r, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cred := &auth.Credentials{}
	json.Unmarshal(body, cred)
	defer resp.Body.Close()
	return cred, nil
}

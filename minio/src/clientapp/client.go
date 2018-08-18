package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func main() {
	//Returns Access Token
	accessToken, err := getAccessToken() //authenticate.go
	if err != nil {
		log.Fatalln(err)
	}
	//accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ4SjZqaWFPN0l2SUxrR05iWVU1TWpraUxPVXdkT2RGSFNKUGozQnpDYl9BIn0.eyJqdGkiOiIxNzU3YWE0MC00ZmQyLTRhMDgtYWEyNS1lNGRmNzk0M2ZlYzIiLCJleHAiOjE1Mjg0MTAxNzIsIm5iZiI6MCwiaWF0IjoxNTI4NDA5ODcyLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvYXV0aC9yZWFsbXMvZGVtbyIsImF1ZCI6Im15LWZpcnN0LWFwcCIsInN1YiI6IjViNWRiNGRjLWMyOTYtNDdhOC1hN2EzLTJlM2QxNWY2N2E4YSIsInR5cCI6IkJlYXJlciIsImF6cCI6Im15LWZpcnN0LWFwcCIsImF1dGhfdGltZSI6MTUyODQwOTg3Miwic2Vzc2lvbl9zdGF0ZSI6ImZjMjIxNmRjLWM3NmYtNDZlZS05YTgwLTQ3ZTYxMGJiYjUyNiIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOltdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJuYW1lIjoiYXNob2sgbW91bGkiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhc2hvay1kZW1vIiwiZ2l2ZW5fbmFtZSI6ImFzaG9rIiwiZmFtaWx5X25hbWUiOiJtb3VsaSIsImVtYWlsIjoiYXNob2ttb3VsaUB5YWhvby5jb20ifQ.Dt5XBW2BzMNZJiaFN1Y-wOpyQgDDchH7aoMNtAGam8cprlwJ_h-1hEV1zyttf8SkQewUmd2mUEPgbWjQFzUbkV5fEyxyKGYzxV2B5Ptn48HG9nkeNN79N1p2QUzIVg-ZhdcsTNM4ye7cZF6MrnHd1ogtE_1T-rH2PFe_V0jlS05uTMGzenPe5Lf1s6vHV7I-H_GLvUk7LQUVxxz4u2UvFw6ZO7WxPRM2NVMm_8fRIP4WJBilktfuhkEc4gqtWGaOenpO_i1xpUI2ZNybbHb4VrlNHua1c8R3mGJb1MQkXF79wu-qdixTivN2DUe3owJ0lIypy_z_cU6ngFqRz8mYpg"

  //Exchange Access Token from IDP for Minio Credentials
	cred, err := getMinioCred(accessToken) //minioCred.go
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("AccessKey: %s SecretKey: %s\n", cred.AccessKey, cred.SecretKey)
	endpoint := "127.0.0.1:9000"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, cred.AccessKey, cred.SecretKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	location := "demo"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// Upload the zip file
	objectName := "code2.java"
	filePath := "/Users/sanatmouli/Downloads/code2.java"
	contentType := "application.txt"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

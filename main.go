package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	s3Bucket:="whodis-test-infra-event-images"
	fmt.Printf("Hola Mundo\n")

	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile: "saml",
	}))
	_, err := s.Config.Credentials.Get()
	if err != nil {
		fmt.Println("Error", err)
	}
		// listbuckets(s)
	listContentBucket(s,s3Bucket)
}

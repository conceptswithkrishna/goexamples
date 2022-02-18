package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	sess := session.Must(session.NewSession())

	svc := s3.New(sess)

	_ = svc

}

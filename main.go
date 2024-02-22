package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	sesClient := ses.NewFromConfig(cfg)
	mailTo := os.Getenv("MAIL_TO")
	mailFrom := os.Getenv("MAIL_FROM")
	_, err = sesClient.SendEmail(context.TODO(), &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{mailTo},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data:    aws.String("Hello, <strong>World</strong>"),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &types.Content{
				Data:    aws.String("Hello, World"),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(mailFrom),
	})
	if err != nil {
		panic(err)
	}

}

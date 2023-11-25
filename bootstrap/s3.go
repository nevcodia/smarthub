package bootstrap

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func NewS3Client(env *Env) *s3.Client {
	region := env.S3Region
	if region == "" {
		region = "aws-global"
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(env.S3AccessKey, env.S3SecretKey, "")),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(env.S3HostAddr)
	})

	return client
}

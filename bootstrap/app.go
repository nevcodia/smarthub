package bootstrap

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Application struct {
	Env *Env
	S3  *s3.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.S3 = NewS3Client(app.Env)
	return *app
}

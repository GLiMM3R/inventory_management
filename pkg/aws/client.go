package aws_service

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Config holds all S3 configuration parameters
type Config struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
}

// S3Client wraps the S3 client and its operations
type S3Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
}

// type resolverV2 struct{}

// func (*resolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (
// 	smithyendpoints.Endpoint, error,
// ) {
// 	// fallback to default
// 	return s3.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
// }

// NewS3Client creates a new S3 client with the specified configuration.
func NewS3Client(cfg Config) (*S3Client, error) {
	// Load AWS configuration with explicit credentials and custom endpoint resolver
	// awsCfg, err := config.LoadDefaultConfig(context.TODO(),
	// 	config.WithRegion(cfg.Region),
	// 	config.WithCredentialsProvider(
	// 		credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
	// 	),
	// )

	// if err != nil {
	// 	log.Fatalf("unable to load SDK config, %v", err)
	// 	return nil, err
	// }

	// client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
	// 	o.UsePathStyle = true
	// 	o.BaseEndpoint = aws.String(cfg.Endpoint)
	// 	o.EndpointResolverV2 = &resolverV2{}
	// })

	options := s3.Options{
		Region:       cfg.Region,
		Credentials:  aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
		BaseEndpoint: aws.String(cfg.Endpoint),
	}

	client := s3.New(options, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	presignClient := s3.NewPresignClient(client)

	return &S3Client{
		client:        client,
		presignClient: presignClient,
	}, nil
}

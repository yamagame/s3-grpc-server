package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.TODO()

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID {
			return aws.Endpoint{URL: "http://localhost:5000", HostnameImmutable: true}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("us-west-1"),
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	output0, err := client.ListBuckets(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*output0.Buckets[1].Name)

	// output1, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
	// 	Bucket: aws.String("testBucket"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(output1)

	// output1, err := client.CreateObject(ctx, &s3.CreateObjectInput{
	// 	Bucket: aws.String("test"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(output1)

	// client.PutObject(ctx, &s3.PutObjectInput{
	// 	Bucket: aws.String("test"),
	// 	Key:    aws.String("test"),
	// 	Body:   strings.NewReader("hello world"),
	// })

	// output2, err := client.GetObject(ctx, &s3.GetObjectInput{
	// 	Bucket: aws.String("test"),
	// 	Key:    aws.String("test"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// streamToString := func(stream io.Reader) string {
	// 	buf := new(bytes.Buffer)
	// 	buf.ReadFrom(stream)
	// 	return buf.String()
	// }

	// log.Println(streamToString(output2.Body))
}

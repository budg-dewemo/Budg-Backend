package repository

import (
	"BudgBackend/src/config"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"mime/multipart"
	"net/http"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	envConfig     s3config
)

type s3config struct {
	accessKey string
	secretKey string
	region    string
	bucket    string
}

func init() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	envConfig.accessKey = config.AwsS3AccessKeyId
	envConfig.secretKey = config.AwsS3SecretKey
	envConfig.region = config.AwsS3Region
	envConfig.bucket = config.AwsS3Bucket

	InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

const (
	AWS_S3_REGION = "" // Region
	AWS_S3_BUCKET = "" // Bucket

)

func ListBuckets() {
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func createConnection() *s3.S3 {
	creds := credentials.NewStaticCredentials(envConfig.accessKey, envConfig.secretKey, "")
	_, err := creds.Get()

	if err != nil {
		fmt.Println("bad credentials")
	}
	//WithRegion("us-west-1")
	cfg := aws.NewConfig().WithCredentials(creds).WithRegion(envConfig.region)
	svc := s3.New(session.New(), cfg)

	return svc

}

func PutFile(fileHandler *multipart.FileHeader, file multipart.File, transactionid int64) (string, error) {

	svc := createConnection()
	var size int64 = fileHandler.Size

	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	//create path
	path := fmt.Sprintf("/media/transactions/%d-%s", transactionid, fileHandler.Filename)
	params := &s3.PutObjectInput{
		Bucket:        aws.String(envConfig.bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	_, err := svc.PutObject(params)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com%s", envConfig.bucket, envConfig.region, path)
	return url, nil
	//fmt.Printf("response %s", awsutil.StringValue(resp))
}

//func GetFile(transactionid int64) (string, error) {
//	svc := createConnection()
//	path := fmt.Sprintf("/media/transactions/%d-%s", transactionid, "test.jpg")
//	params := &s3.GetObjectInput{
//		Bucket: aws.String(envConfig.bucket),
//		Key:    aws.String(path),
//	}
//	resp, err := svc.GetObject(params)
//
//	if err != nil {
//		return "", err
//	}
//	fmt.Printf("response %s", awsutil.StringValue(resp))
//	return "", nil
//}

package upload

import (
	"bytes"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"echo/common/logger"
	"echo/conf"
	"mime/multipart"
	"time"
)

type AliyunOSS struct{}

func (*AliyunOSS) MyUploadFile(decryptFile []byte, fileName, fileName2, ext string) (string, error) {
	bucket, err := NewBucket()
	if err != nil {
		logger.AdminLog.Error("[APP] "+"function AliyunOSS.NewBucket() Failed", zap.Error(err))
		return "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	filePath := conf.Config.AliyunOSS.BasePath + "/" + time.Now().Format("2006-01-02") + "/"
	if fileName != "" {
		filePath += fileName + ext
	} else {
		filePath += fileName2
	}
	reader := bytes.NewReader(decryptFile)
	err = bucket.PutObject(filePath, reader)
	if err != nil {
		logger.AdminLog.Error("[APP] "+"function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return "/" + filePath, nil
}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader, fileName, ext string) (string, error) {
	bucket, err := NewBucket()
	if err != nil {
		logger.AdminLog.Error("[APP] "+"function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	f, openError := file.Open()
	if openError != nil {
		logger.AdminLog.Error("[APP] "+"function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close()
	filePath := conf.Config.AliyunOSS.BasePath + "/" + time.Now().Format("2006-01-02") + "/"
	if fileName != "" {
		filePath += fileName + ext
	} else {
		filePath += file.Filename
	}

	err = bucket.PutObject(filePath, f)
	if err != nil {
		logger.AdminLog.Error("[APP] "+"function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return "/" + filePath, nil
}

func (*AliyunOSS) DeleteFile(filePath string) error {
	bucket, err := NewBucket()
	if err != nil {
		logger.AdminLog.Error("[APP] "+"function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	err = bucket.DeleteObject(filePath)
	if err != nil {
		logger.AdminLog.Error("[APP] "+"function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	client, err := oss.New(conf.Config.AliyunOSS.Endpoint, conf.Config.AliyunOSS.AccessKeyId, conf.Config.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(conf.Config.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

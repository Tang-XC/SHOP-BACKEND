package common

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

func WrapFunc(f interface{}, args ...interface{}) gin.HandlerFunc {
	//使用反射reflect获取函数f的value
	fn := reflect.ValueOf(f)

	//判断传入的函数f的参数数量和args传入的数量是否相同
	if fn.Type().NumIn() != len(args) {
		panic(fmt.Sprintf("invalid input parameters of function %v", fn.Type()))
	}
	//判断传入的函数f的返回参数是否为0，这里规定函数不可返回空
	if fn.Type().NumOut() == 0 {
		panic(fmt.Sprintf("invalid output parameters of function %v,at least one", fn.Type()))
	}

	//反射所有参数args传过来的值
	inputs := make([]reflect.Value, len(args))
	for k, v := range args {
		inputs[k] = reflect.ValueOf(v)
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				FailedResponse(c, http.StatusInternalServerError, fmt.Errorf("%v", err))
			}
		}()

		outputs := fn.Call(inputs)
		if len(outputs) > 1 {
			err, ok := outputs[len(outputs)-1].Interface().(error)
			if ok && err != nil {
				FailedResponse(c, http.StatusInternalServerError, err)
				return
			}
		}
		c.JSON(http.StatusOK, outputs[0].Interface())
	}
}

// 创建桶
func CreateBucket(bucketName string, client *minio.Client) error {
	if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false}); err != nil {
		exists, _ := client.BucketExists(context.Background(), bucketName)
		if exists {
			fmt.Println("桶已创建!")
			return nil
		}
		return err
	}
	fmt.Println("桶创建成功!")
	return nil
}
func FileUploader(bucketName string, fileName string, reader io.Reader, size int64, client *minio.Client) (string, error) {
	contextType := "application/text"
	object, err := client.PutObject(context.Background(), bucketName, fileName, reader, size, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("成功上传%s,大小为%d字节 \n", fileName, object.Size)
	return message, nil
}

// 获取文件访问链接
func FileURL(bucketName string, fileName string, client *minio.Client) (string, error) {
	//获取文件访问链接
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")
	presignedURL, err := client.PresignedGetObject(context.Background(), bucketName, fileName, time.Second*24*60*60, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

// 获取文件的信息
func FileInfo(bucketName string, fileName string, client *minio.Client) (minio.ObjectInfo, error) {
	object, err := client.StatObject(context.Background(), bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		return minio.ObjectInfo{}, err
	}
	return object, nil
}

// 创建文件夹
func CreateFolder(bucketName string, folderName string, client *minio.Client) error {
	//判断文件夹是否存在
	_, err := client.StatObject(context.Background(), bucketName, folderName, minio.StatObjectOptions{})
	if err != nil {
		if err.Error() == "The specified key does not exist." {
			//创建文件夹
			_, err := client.PutObject(context.Background(), bucketName, folderName+"/", nil, 0, minio.PutObjectOptions{})
			if err != nil {
				return err
			}
			fmt.Println("文件夹创建成功!")
			return nil
		}
		return err
	}
	return nil
}

// 删除文件
func DeleteFile(bucketName string, fileName string, client *minio.Client) error {
	err := client.RemoveObject(context.Background(), bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	fmt.Println("文件删除成功!")
	return nil
}

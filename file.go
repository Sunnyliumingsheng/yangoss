package yangoss

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
)

func UploadFile(file multipart.File, handler *multipart.FileHeader) (url string, err error) {
	var NewFilename string
	NewFilename = Hash(handler.Filename)
	extension := filepath.Ext(handler.Filename)

	dst := Config.StoragePath + NewFilename + extension
	newFile, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(newFile, file)
	if err != nil {
		return "", err
	}
	url = Config.ServerPath + "/" + NewFilename + extension
	return url, nil
}
func RemoveFile(fileUrl string) (err error) {
	parsedURL, err := url.Parse(fileUrl)
	if err != nil {
		log.Println("严重错误，解析照片url位dst失败")
		return errors.New("Error parsing URL:" + err.Error())
	}

	// 提取文件名
	fileName := filepath.Base(parsedURL.Path)

	// 构建文件的完整路径
	dst := filepath.Join(Config.ServerPath, fileName)
	err = os.Remove(dst)
	return err
}

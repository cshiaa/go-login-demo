package storage

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/source/config"
	"go.uber.org/zap"
)


type FileStorage struct {
	Storage config.Storage `json: "storage"`
}

//检查文件存储是否可用
func (s *FileStorage) Check() error {
	return nil
}

//文件上传
func (s *FileStorage) UploadFile(file *multipart.FileHeader) (error){

	global.RY_LOG.Info("文件存储类型",zap.String("storagetype", global.RY_CONFIG.Storage.Type))

	if (global.RY_CONFIG.Storage.Type == "local") {

		localFilePath := global.RY_CONFIG.Storage.Local.FilePath

		_,err := os.Stat(localFilePath)
		if err != nil {
			//isnotexist来判断，是不是不存在的错误
			if os.IsNotExist(err){  //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
				global.RY_LOG.Info("目录不存在，创建目录", zap.String("filedir", localFilePath))
				err := os.Mkdir(localFilePath, os.ModePerm)
				if err != nil {
					global.RY_LOG.Error("创建目录失败", zap.String("filedir", localFilePath))
					return err
				}
			}
		}

		// 打开文件头以获取文件句柄
		f, err := file.Open()
		if err != nil {
            global.RY_LOG.Error("打开文件句柄失败")
			return err
		}
		defer f.Close()

		filename := localFilePath + "/" +  file.Filename
		dst, err := os.Create(filename)
		if err != nil {
            global.RY_LOG.Error("创建文件失败", zap.String("filedir", filename))
            return err
        }
		defer dst.Close()

		_, err = io.Copy(dst, f)
		if err != nil {
            global.RY_LOG.Error("文件拷贝失败")
			return err
		}
		global.RY_LOG.Info("更改Kubernetes配置文件路径")
		global.RY_VP.Set("kubernetes.path", filename)
		err = global.RY_VP.WriteConfig()
		if err != nil {
			global.RY_LOG.Error("更改Kubernetes配置失败")
			return err
		}
	}
	global.RY_LOG.Info("上传文件成功")
	return nil

}

func (s *FileStorage) GetFile(filename string) error {
	
	return nil
}
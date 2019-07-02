package service

import (
	"fmt"
	"mall/app/api/web/system/model"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func (s *Service) UploadToken() string {
	var qiniu model.AttachQiniu
	remote := s.d.LoadRemote()
	if remote.Type != model.ATTACH_QINIU {
		remote.Qiniu.AccessKey = "XQaioWVSLln3UuxZGgzNIfwR1M50fo15Hhl47nrM"
		remote.Qiniu.SecretKey = "sKyN9nwDVlDnvUITunZkcLXcP8AyMGubcyZ8DC4i"
		remote.Qiniu.Bucket = "wanshang-official"
		remote.Qiniu.Url = "https://cdn.wenluokeji.com"
	}

	qiniu = remote.Qiniu
	fmt.Println(qiniu)

	mac := qbox.NewMac(qiniu.AccessKey, qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: qiniu.Bucket,
	}

	return putPolicy.UploadToken(mac)

}

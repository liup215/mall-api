package model

const (
	ATTACH_QINIU = 3
)

type Remote struct {
	Type  int         `json:"type"`
	Qiniu AttachQiniu `json:"qiniu"`
}

type AttachFTP struct {
	Ssl      string `json:"ssl"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:password`
	Pasv     string `json:"pasv"`
	Dir      string `json:"dir"`
	Url      string `json:"url"`
}

type AttachQiniu struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Url       string `json:"url"`
}

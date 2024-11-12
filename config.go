package yangoss

var Config OssConfig

type OssConfig struct {
	StoragePath string
	ServerPath  string
}

func init() {
	Config.StoragePath = "~/temp/picture"
	Config.ServerPath = "http://localhost:80/img"
}

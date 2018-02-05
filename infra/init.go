package infra

func Init() {
	InitLogger()
	InitTime()
	InitFileFactory()
	InitConfig("config.toml")
	InitClient()
	InitHost()
}

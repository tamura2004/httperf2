package infra

func Init() {
	InitTime()
	InitFileFactory()
	InitLogger()
	InitConfig("config.toml")
	InitClient()
	InitHost()
}

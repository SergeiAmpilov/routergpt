package providers

var ProviderList = make(map[string]string, 0)

func init() {
	ProviderList["openai"] = "OpenAI"
	ProviderList["yandex"] = "Yandex"
	ProviderList["sber"] = "Sber Cloud"
	ProviderList["alibaba"] = "Alibaba Cloud"
}

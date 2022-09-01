package service

type BannerService struct {
}

func (b BannerService) Id() string {
	return "banner"
}

func (b BannerService) Provide(userId uint64) interface{} {
	return []string{
		"hello",
		"world",
		"feego",
	}
}

func (b BannerService) StrongDepend() bool {
	return false
}

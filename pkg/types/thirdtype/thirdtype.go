package thirdtype

type AppConfig struct {
	ClientID     string
	ClientSecret string
}

type HttpOptions struct {
	Timeout int
	Verify  bool
}

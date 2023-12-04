package config

import (
	"os"
	"sync"

	T "mirror/tool"

	"gopkg.in/yaml.v2"
)

var (
	once   sync.Once
	config *Config
)

func GetConfig() *Config {
	if config != nil {
		return config
	}
	once.Do(loadConfig)
	return config
}

var data = `
enable_ssl: True
handle_cookie: True

host:
  self: google-mirror-nangua.vercel.app
  proxy: www.startpage.com

replaced_urls:
  - old: www.startpage.com
    new: google-mirror-nangua.vercel.app

header_token_key: X-AUTH-TOKEN
`

func loadConfig() {
	config = new(Config)
	err := yaml.Unmarshal([]byte(data), config)
	config.Token = os.Getenv("X_AUTH_TOKEN")
	T.CheckErr(err)
	if config.EnableSSL {
		config.Protocol = "https://"
	} else {
		config.Protocol = "http://"
	}
}

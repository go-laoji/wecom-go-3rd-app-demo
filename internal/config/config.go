package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	CorpId string
	SuiteId string
	SuiteSecret string
	SuiteToken string
	SuiteEncodingAesKey string
}

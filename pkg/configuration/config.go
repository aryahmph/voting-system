package configuration

import (
	"github.com/joho/godotenv"
	"os"
)

type Configuration interface {
	Get(key string) string
}

type configurationImpl struct {
}

func (c *configurationImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfigurationImpl(filenames ...string) *configurationImpl {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil
	}
	return &configurationImpl{}
}

package littlewarden

import (
	"errors"
)

type Config struct {
	ApiKey string
}

func (c *Config) Client() (interface{}, error) {
	err := c.validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c Config) validate() error {
	if c.ApiKey == "" {
		return errors.New("api_key must be configured for the Little Warden provider")
	}

	return nil
}

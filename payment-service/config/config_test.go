package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("from .env file", func(t *testing.T) {
		cfg, err := LoadConfig()

		assert.Nil(t, err)
		assert.Equal(t, "8088", cfg.Port)
	})
}

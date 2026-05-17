package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	t.Setenv("APP_ENCRYPTION_KEY", "?E(G+KbPeShVmYq3t6w9z$C&F)J@McQf")
	_, err := GetConfig()
	require.NoError(t, err)

	t.Setenv("APP_ENVIRONMENT", "abc")
	cfg, err := GetConfig()
	require.NoError(t, err)
	assert.Equal(t, environment("abc"), cfg.App.Environment)
}

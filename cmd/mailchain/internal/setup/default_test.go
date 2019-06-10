package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultSentStorage(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
	}{
		{
			"success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultSentStorage()
			if !assert.NotNil(got) {
				t.Error("want got != nil")
			}
			if !assert.NotNil(got.selectItemSkipable) {
				t.Error("want got.selectItemSkipable != nil")
			}
			if !assert.NotNil(got.sentStoreSetter) {
				t.Error("want got.requiredInput != nil")
			}
			if !assert.NotNil(got.viper) {
				t.Error("want got.viper != nil")
			}
		})
	}
}

func TestDefaultKeystore(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
	}{
		{
			"success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultKeystore()
			if !assert.NotNil(got) {
				t.Error("want got != nil")
			}
			if !assert.NotNil(got.selectItemSkipable) {
				t.Error("want got.selectItemSkipable != nil")
			}
			if !assert.NotNil(got.keystoreSetter) {
				t.Error("want got.keystoreSetter != nil")
			}
			if !assert.NotNil(got.viper) {
				t.Error("want got.viper != nil")
			}
		})
	}
}

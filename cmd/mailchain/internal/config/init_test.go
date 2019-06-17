// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {
	type args struct {
		viper    *viper.Viper
		cfgFile  string
		logLevel string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{viper.New(), "./testdata/.empty.yaml", "DEBUG"},
			false,
		},
		{
			"err-invalid-file",
			args{viper.New(), "./testdata/.invalid.yaml", "DEBUG"},
			true,
		},
		{
			"err-no-file",
			args{viper.New(), "./testdata/.no-file.yaml", "DEBUG"},
			true,
		},
		{
			"invalid-level",
			args{viper.New(), "./testdata/.empty.yaml", "INVALID"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.viper, tt.args.cfgFile, tt.args.logLevel); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
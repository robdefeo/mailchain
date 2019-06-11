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

package stores

import (
	"reflect"
	"testing"
)

func TestSentStoreNames(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			"success",
			[]string{"mailchain", "s3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SentStoreNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentStoreNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

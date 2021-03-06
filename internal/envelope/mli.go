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

package envelope

const (
	MLIEmpty     uint64 = 0
	MLIMailchain uint64 = 1
)

// // LocationCode maps the location to the code
// func LocationCode() map[string]uint64 {
// 	return map[string]uint64{
// 		locationMailchain: CodeMailchain,
// 	}
// }

// MLIToAddress maps code to a location
func MLIToAddress() map[uint64]string {
	return map[uint64]string{
		MLIMailchain: mliMailchain,
	}
}

const (
	mliMailchain = "https://mcx.mx"
)

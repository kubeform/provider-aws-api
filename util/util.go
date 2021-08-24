/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"github.com/google/go-cmp/cmp"
	"github.com/qri-io/jsonpointer"
	"strconv"
)

func CheckIfAnyDifference(soFarKey string, keySplit []string, pos, length int, checkIfAnyDiff *bool, temp, tempOld, tempNew map[string]interface{}) {
	soFarKey = soFarKey + keySplit[pos]
	if pos == length-1 {
		ptr, _ := jsonpointer.Parse(soFarKey)
		valueOld, _ := ptr.Eval(tempOld)
		valueNew, _ := ptr.Eval(tempNew)

		if !cmp.Equal(valueNew, valueOld) {
			*checkIfAnyDiff = true
		}
		return
	}

	ptr, _ := jsonpointer.Parse(soFarKey)
	value, _ := ptr.Eval(temp)

	if value == nil {
		return
	}

	tmpValue := value.([]interface{})
	sz := len(tmpValue)

	for i := 0; i < sz; i++ {
		CheckIfAnyDifference(soFarKey+"/"+strconv.Itoa(i), keySplit, pos+1, length, checkIfAnyDiff, temp, tempOld, tempNew)
	}
}
/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package runtime

import (
	"github.com/onflow/cadence"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/onflow/cadence/runtime/tests/utils"
)

func executeScript(script string, runtimeInterface Interface) (cadence.Value, error) {
	runtime := newTestInterpreterRuntime()

	return runtime.ExecuteScript(
		Script{
			Source: []byte(script),
		},
		Context{
			Interface: runtimeInterface,
			Location:  utils.TestLocation,
		},
	)
}

func TestAssert(t *testing.T) {

	t.Parallel()

	script := `
        pub fun main() {
          Test.assert(false, "condition not satisfied")
        }
    `

	storage := newTestLedger(nil, nil)

	runtimeInterface := &testRuntimeInterface{
		storage: storage,
	}

	_, err := executeScript(script, runtimeInterface)
	require.NoError(t, err)
}

func TestBlockchain(t *testing.T) {

	t.Parallel()

	script := `
        pub fun main() {
          var bc = Test.Blockchain()
          bc.executeScript("pub fun foo() {}")
        }
    `

	storage := newTestLedger(nil, nil)

	runtimeInterface := &testRuntimeInterface{
		storage: storage,
	}

	_, err := executeScript(script, runtimeInterface)
	require.NoError(t, err)
}

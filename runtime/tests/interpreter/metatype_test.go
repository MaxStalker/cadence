/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
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

package interpreter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/cadence/runtime/interpreter"
)

func TestInterpretMetaType(t *testing.T) {

	t.Parallel()

	t.Run("constructor", func(t *testing.T) {
		t.Parallel()

		inter := parseCheckAndInterpret(t, `
           let equal = Type<Int>() == Type<Int>()
           let unequal = Type<Int>() == Type<String>()
           let optional = Type<Int>() == Type<Int?>()
        `)

		assert.Equal(t,
			interpreter.BoolValue(true),
			inter.Globals["equal"].Value,
		)

		assert.Equal(t,
			interpreter.BoolValue(false),
			inter.Globals["unequal"].Value,
		)

		assert.Equal(t,
			interpreter.BoolValue(false),
			inter.Globals["optional"].Value,
		)
	})

	t.Run("identifier", func(t *testing.T) {

		t.Parallel()

		inter := parseCheckAndInterpret(t, `
          let type = Type<[Int]>()
          let identifier = type.identifier
        `)

		assert.Equal(t,
			interpreter.NewStringValue("[Int]"),
			inter.Globals["identifier"].Value,
		)
	})
}
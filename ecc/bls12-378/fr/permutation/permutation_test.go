// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package permutation

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/consensys/gnark-crypto/ecc/bls12-378/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-378/fr/kzg"
)

func TestProof(t *testing.T) {

	kzgSrs, err := kzg.NewSRS(64, big.NewInt(13))
	assert.NoError(t, err)

	a := make([]fr.Element, 8)
	b := make([]fr.Element, 8)

	for i := 0; i < 8; i++ {
		a[i].SetUint64(uint64(4*i + 1))
	}
	for i := 0; i < 8; i++ {
		b[i].Set(&a[(5*i)%8])
	}

	// correct proof
	{
		proof, err := Prove(kzgSrs.Pk, a, b)
		if err != nil {
			t.Fatal(err)
		}

		err = Verify(kzgSrs.Vk, proof)
		if err != nil {
			t.Fatal(err)
		}
	}

	// wrong proof
	{
		a[0].SetRandom()
		proof, err := Prove(kzgSrs.Pk, a, b)
		if err != nil {
			t.Fatal(err)
		}

		err = Verify(kzgSrs.Vk, proof)
		if err == nil {
			t.Fatal(err)
		}
	}

}

func BenchmarkProver(b *testing.B) {

	srsSize := 1 << 15
	polySize := 1 << 14

	kzgSrs, _ := kzg.NewSRS(uint64(srsSize), big.NewInt(13))
	a := make([]fr.Element, polySize)
	c := make([]fr.Element, polySize)

	for i := 0; i < polySize; i++ {
		a[i].SetUint64(uint64(i))
	}
	for i := 0; i < polySize; i++ {
		c[i].Set(&a[(5*i)%(polySize)])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Prove(kzgSrs.Pk, a, c)
	}

}

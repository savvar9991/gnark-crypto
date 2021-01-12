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

// Code generated by gurvy DO NOT EDIT

package twistededwards

import (
	"io"
	"math/big"
	"math/bits"

	"github.com/consensys/gurvy/bls377/fr"
)

// size in byte of (x,y)
const sizeOfPointAffine = fr.Limbs * 16

// Bytes returns the point as bytes array x||y, where
// x and y are in big endian.
func (p *PointAffine) Bytes() [sizeOfPointAffine]byte {

	var res [sizeOfPointAffine]byte
	x := p.X.Bytes()
	y := p.Y.Bytes()
	copy(res[:], x[:])
	copy(res[fr.Limbs*8:], y[:])
	return res
}

// Marshal converts p to a byte slice
func (p *PointAffine) Marshal() []byte {
	b := p.Bytes()
	return b[:]
}

// SetBytes sets p from the buf, where bug is interpreted
// as a sizeOfPointAffine+ byte slice, where the first 32 bytes
// are interpreted as x in big endian, the next 32 bytes are
// interpreted as y in big endian. Returns the number of read bytes
// and an error if the buffer is too short.
// Returns the number of bytes read.
func (p *PointAffine) SetBytes(buf []byte) (int, error) {

	if len(buf) < sizeOfPointAffine {
		return 0, io.ErrShortBuffer
	}

	p.X.SetBytes(buf[:fr.Limbs*8])
	p.Y.SetBytes(buf[fr.Limbs*8:])
	return 64, nil
}

// Unmarshal alias to SetBytes()
func (p *PointAffine) Unmarshal(b []byte) error {
	_, err := p.SetBytes(b)
	return err
}

// PointAffine point on a twisted Edwards curve
type PointAffine struct {
	X, Y fr.Element
}

// PointProj point in projective coordinates
type PointProj struct {
	X, Y, Z fr.Element
}

// Set sets p to p1 and return it
func (p *PointProj) Set(p1 *PointProj) *PointProj {
	p.X.Set(&p1.X)
	p.Y.Set(&p1.Y)
	p.Z.Set(&p1.Z)
	return p
}

// Set sets p to p1 and return it
func (p *PointAffine) Set(p1 *PointAffine) *PointAffine {
	p.X.Set(&p1.X)
	p.Y.Set(&p1.Y)
	return p
}

// Equal returns true if p=p1 false otherwise
func (p *PointAffine) Equal(p1 *PointAffine) bool {
	return p.X.Equal(&p1.X) && p.Y.Equal(&p1.Y)
}

// Equal returns true if p=p1 false otherwise
// If one point is on the affine chart Z=0 it returns false
func (p *PointProj) Equal(p1 *PointProj) bool {
	if p.Z.IsZero() || p1.Z.IsZero() {
		return false
	}
	var pAffine, p1Affine PointAffine
	pAffine.FromProj(p)
	p1Affine.FromProj(p1)
	return pAffine.Equal(&p1Affine)
}

// NewPointAffine creates a new instance of PointAffine
func NewPointAffine(x, y fr.Element) PointAffine {
	return PointAffine{x, y}
}

// IsOnCurve checks if a point is on the twisted Edwards curve
func (p *PointAffine) IsOnCurve() bool {

	ecurve := GetEdwardsCurve()

	var lhs, rhs, tmp fr.Element

	tmp.Mul(&p.Y, &p.Y)
	lhs.Mul(&p.X, &p.X).
		Mul(&lhs, &ecurve.A).
		Add(&lhs, &tmp)

	tmp.Mul(&p.X, &p.X).
		Mul(&tmp, &p.Y).
		Mul(&tmp, &p.Y).
		Mul(&tmp, &ecurve.D)
	rhs.SetOne().Add(&rhs, &tmp)

	return lhs.Equal(&rhs)
}

// Add adds two points (x,y), (u,v) on a twisted Edwards curve with parameters a, d
// modifies p
func (p *PointAffine) Add(p1, p2 *PointAffine) *PointAffine {

	ecurve := GetEdwardsCurve()

	var xu, yv, xv, yu, dxyuv, one, denx, deny fr.Element
	pRes := new(PointAffine)
	xv.Mul(&p1.X, &p2.Y)
	yu.Mul(&p1.Y, &p2.X)
	pRes.X.Add(&xv, &yu)

	xu.Mul(&p1.X, &p2.X).Mul(&xu, &ecurve.A)
	yv.Mul(&p1.Y, &p2.Y)
	pRes.Y.Sub(&yv, &xu)

	dxyuv.Mul(&xv, &yu).Mul(&dxyuv, &ecurve.D)
	one.SetOne()
	denx.Add(&one, &dxyuv)
	deny.Sub(&one, &dxyuv)

	p.X.Div(&pRes.X, &denx)
	p.Y.Div(&pRes.Y, &deny)

	return p
}

// Double doubles point (x,y) on a twisted Edwards curve with parameters a, d
// modifies p
func (p *PointAffine) Double(p1 *PointAffine) *PointAffine {
	p.Add(p1, p1)
	return p
}

// Neg negates point (x,y) on a twisted Edwards curve with parameters a, d
// modifies p
func (p *PointAffine) Neg(p1 *PointAffine) *PointAffine {
	p.X.Neg(&p1.X)
	return p
}

// FromProj sets p in affine from p in projective
func (p *PointAffine) FromProj(p1 *PointProj) *PointAffine {
	p.X.Div(&p1.X, &p1.Z)
	p.Y.Div(&p1.Y, &p1.Z)
	return p
}

// FromAffine sets p in projective from p in affine
func (p *PointProj) FromAffine(p1 *PointAffine) *PointProj {
	p.X.Set(&p1.X)
	p.Y.Set(&p1.Y)
	p.Z.SetOne()
	return p
}

// Add adds points in projective coordinates
// cf https://hyperelliptic.org/EFD/g1p/auto-twisted-projective.html
func (p *PointProj) Add(p1, p2 *PointProj) *PointProj {

	var res PointProj

	ecurve := GetEdwardsCurve()

	var A, B, C, D, E, F, G, H, I fr.Element
	A.Mul(&p1.Z, &p2.Z)
	B.Square(&A)
	C.Mul(&p1.X, &p2.X)
	D.Mul(&p1.Y, &p2.Y)
	E.Mul(&ecurve.D, &C).Mul(&E, &D)
	F.Sub(&B, &E)
	G.Add(&B, &E)
	H.Add(&p1.X, &p1.Y)
	I.Add(&p2.X, &p2.Y)
	res.X.Mul(&H, &I).
		Sub(&res.X, &C).
		Sub(&res.X, &D).
		Mul(&res.X, &p1.Z).
		Mul(&res.X, &F)
	H.Mul(&ecurve.A, &C)
	res.Y.Sub(&D, &H).
		Mul(&res.Y, &p.Z).
		Mul(&res.Y, &G)
	res.Z.Mul(&F, &G)

	p.Set(&res)
	return p
}

// Double adds points in projective coordinates
// cf https://hyperelliptic.org/EFD/g1p/auto-twisted-projective.html
func (p *PointProj) Double(p1 *PointProj) *PointProj {

	var res PointProj

	ecurve := GetEdwardsCurve()

	var B, C, D, E, F, H, J, tmp fr.Element

	B.Add(&p1.X, &p1.Y).Square(&B)
	C.Square(&p1.X)
	D.Square(&p1.Y)
	E.Mul(&ecurve.A, &C)
	F.Add(&E, &D)
	H.Square(&p1.Z)
	tmp.Double(&H)
	J.Sub(&F, &tmp)
	res.X.Sub(&B, &C).
		Sub(&res.X, &D).
		Mul(&res.X, &J)
	res.Y.Sub(&E, &D).Mul(&res.Y, &F)
	res.Z.Mul(&F, &J)

	p.Set(&res)
	return p
}

// Neg sets p to -p1 and returns it
func (p *PointProj) Neg(p1 *PointProj) *PointProj {
	p.X.Neg(&p1.X)
	return p
}

// ScalarMul scalar multiplication of a point
// p1 points on the twisted Edwards curve
// c parameters of the twisted Edwards curve
// scal scalar NOT in Montgomery form
// modifies p
//func (p *PointAffine) ScalarMul(p1 *PointAffine, scalar fr.Element) *PointAffine {
func (p *PointAffine) ScalarMul(p1 *PointAffine, scalar *big.Int) *PointAffine {

	var resProj, p1Proj PointProj
	resProj.X.SetZero()
	resProj.Y.SetOne()
	resProj.Z.SetOne()

	p1Proj.FromAffine(p1)

	const wordSize = bits.UintSize

	sWords := scalar.Bits()

	for i := len(sWords) - 1; i >= 0; i-- {
		ithWord := sWords[i]
		for k := 0; k < 64; k++ { // TODO it assumes 64 bits arch, add a check to change 64 to 32 if necessary
			resProj.Double(&resProj)
			kthBit := (ithWord >> (63 - k)) & 1
			if kthBit == 1 {
				resProj.Add(&resProj, &p1Proj)
			}
		}
	}

	p.FromProj(&resProj)

	return p
}

package fq12over6over2

// Fq12 ...
const Fq12 = `

import (
	"math/big"
)

// e12 is a degree two finite field extension of fp6
type e12 struct {
	C0, C1 e6
}

// Equal returns true if z equals x, fasle otherwise
func (z *e12) Equal(x *e12) bool {
	return z.C0.Equal(&x.C0) && z.C1.Equal(&x.C1)
}

// String puts e12 in string form
func (z *e12) String() string {
	return (z.C0.String() + "+(" + z.C1.String() + ")*w")
}

// SetString sets a e12 from string
func (z *e12) SetString(s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11 string) *e12 {
	z.C0.SetString(s0, s1, s2, s3, s4, s5)
	z.C1.SetString(s6, s7, s8, s9, s10, s11)
	return z
}

// Set copies x into z and returns z
func (z *e12) Set(x *e12) *e12 {
	z.C0 = x.C0
	z.C1 = x.C1
	return z
}

// SetOne sets z to 1 in Montgomery form and returns z
func (z *e12) SetOne() *e12 {
	*z = e12{}
	z.C0.B0.A0.SetOne()
	return z
}

// ToMont converts to Mont form
func (z *e12) ToMont() *e12 {
	z.C0.ToMont()
	z.C1.ToMont()
	return z
}

// FromMont converts from Mont form
func (z *e12) FromMont() *e12 {
	z.C0.FromMont()
	z.C1.FromMont()
	return z
}

// Add set z=x+y in e12 and return z
func (z *e12) Add(x, y *e12) *e12 {
	z.C0.Add(&x.C0, &y.C0)
	z.C1.Add(&x.C1, &y.C1)
	return z
}

// Sub sets z to x sub y and return z
func (z *e12) Sub(x, y *e12) *e12 {
	z.C0.Sub(&x.C0, &y.C0)
	z.C1.Sub(&x.C1, &y.C1)
	return z
}

// Double sets z=2*x and returns z
func (z *e12) Double(x *e12) *e12 {
	z.C0.Double(&x.C0)
	z.C1.Double(&x.C1)
	return z
}

// SetRandom used only in tests
func (z *e12) SetRandom() *e12 {
	z.C0.B0.A0.SetRandom()
	z.C0.B0.A1.SetRandom()
	z.C0.B1.A0.SetRandom()
	z.C0.B1.A1.SetRandom()
	z.C0.B2.A0.SetRandom()
	z.C0.B2.A1.SetRandom()
	z.C1.B0.A0.SetRandom()
	z.C1.B0.A1.SetRandom()
	z.C1.B1.A0.SetRandom()
	z.C1.B1.A1.SetRandom()
	z.C1.B2.A0.SetRandom()
	z.C1.B2.A1.SetRandom()
	return z
}

// Mul set z=x*y in e12 and return z
func (z *e12) Mul(x, y *e12) *e12 {
	var a, b, c e6
	a.Add(&x.C0, &x.C1)
	b.Add(&y.C0, &y.C1)
	a.Mul(&a, &b)
	b.Mul(&x.C0, &y.C0)
	c.Mul(&x.C1, &y.C1)
	z.C1.Sub(&a, &b).Sub(&z.C1, &c)
	z.C0.MulByNonResidue(&c).Add(&z.C0, &b)
	return z
}

// Square set z=x*x in e12 and return z
func (z *e12) Square(x *e12) *e12 {

	//Algorithm 22 from https://eprint.iacr.org/2010/354.pdf
	var c0, c2, c3 e6
	c0.Sub(&x.C0, &x.C1)
	c3.MulByNonResidue(&x.C1).Neg(&c3).Add(&x.C0, &c3)
	c2.Mul(&x.C0, &x.C1)
	c0.Mul(&c0, &c3).Add(&c0, &c2)
	z.C1.Double(&c2)
	c2.MulByNonResidue(&c2)
	z.C0.Add(&c0, &c2)

	return z
}

// squares an element a+by interpreted as an Fp4 elmt, where y**2= non_residue_e2
func fp4Square(a, b, c, d *e2) {
	var tmp e2
	c.Square(a)
	tmp.Square(b).MulByNonResidue(&tmp)
	c.Add(c, &tmp)
	d.Mul(a, b).Double(d)
}

// CyclotomicSquare https://eprint.iacr.org/2009/565.pdf, 3.2
func (z *e12) CyclotomicSquare(x *e12) *e12 {

	var res, b, a e12
	var tmp e2

	// A
	fp4Square(&x.C0.B0, &x.C1.B1, &b.C0.B0, &b.C1.B1)
	a.C0.B0.Set(&x.C0.B0)
	a.C1.B1.Neg(&x.C1.B1)

	// B
	tmp.MulByNonResidueInv(&x.C1.B0)
	fp4Square(&x.C0.B2, &tmp, &b.C0.B1, &b.C1.B2)
	b.C0.B1.MulByNonResidue(&b.C0.B1)
	b.C1.B2.MulByNonResidue(&b.C1.B2)
	a.C0.B1.Set(&x.C0.B1)
	a.C1.B2.Neg(&x.C1.B2)

	// C
	fp4Square(&x.C0.B1, &x.C1.B2, &b.C0.B2, &b.C1.B0)
	b.C1.B0.MulByNonResidue(&b.C1.B0)
	a.C0.B2.Set(&x.C0.B2)
	a.C1.B0.Neg(&x.C1.B0)

	res.Set(&b)
	b.Sub(&b, &a).Double(&b)
	z.Add(&res, &b)

	return z
}

// Inverse set z to the inverse of x in e12 and return z
func (z *e12) Inverse(x *e12) *e12 {
	// Algorithm 23 from https://eprint.iacr.org/2010/354.pdf

	var t0, t1, tmp e6
	t0.Square(&x.C0)
	t1.Square(&x.C1)
	tmp.MulByNonResidue(&t1)
	t0.Sub(&t0, &tmp)
	t1.Inverse(&t0)
	z.C0.Mul(&x.C0, &t1)
	z.C1.Mul(&x.C1, &t1).Neg(&z.C1)

	return z
}

// Exp sets z=x**e and returns it
func (z *e12) Exp(x *e12, e big.Int) *e12 {
	var res e12
	res.SetOne()
	b := e.Bytes()
	for i := range b {
		w := b[i]
		mask := byte(0x80)
		for j := 7; j >= 0; j-- {
			res.Square(&res)
			if (w&mask)>>j != 0 {
				res.Mul(&res, x)
			}
			mask = mask >> 1
		}
	}
	z.Set(&res)
	return z
}

// InverseUnitary inverse a unitary element
func (z *e12) InverseUnitary(x *e12) *e12 {
	return z.Conjugate(x)
}

// Conjugate set z to x conjugated and return z
func (z *e12) Conjugate(x *e12) *e12 {
	*z = *x 
	z.C1.Neg(&z.C1)
	return z
}


{{- $sizeOfFp := mul .Fp.NbWords 8}}

// sizeGT represents the size in bytes that a GT element need in binary form
const sizeGT = {{ $sizeOfFp }} * 12

// Bytes returns the regular (non montgomery) value 
// of z as a big-endian byte slice.
// z.C1.B2.A1 | z.C1.B2.A0 | z.C1.B1.A1 | ...
func (z *e12) Bytes() []byte {
	var r [sizeGT]byte

	_z := *z
	_z.FromMont()

	{{- $offset := mul $sizeOfFp 11}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C0.B0.A0"}}
	
	{{- $offset := mul $sizeOfFp 10}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C0.B0.A1"}}

	{{- $offset := mul $sizeOfFp 9}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C0.B1.A0"}}

	{{- $offset := mul $sizeOfFp 8}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C0.B1.A1"}}

	{{- $offset := mul $sizeOfFp 7}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C0.B2.A0"}}
	
	{{- $offset := mul $sizeOfFp 6}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C0.B2.A1"}}

	{{- $offset := mul $sizeOfFp 5}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C1.B0.A0"}}

	{{- $offset := mul $sizeOfFp 4}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C1.B0.A1"}}

	{{- $offset := mul $sizeOfFp 3}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C1.B1.A0"}}

	{{- $offset := mul $sizeOfFp 2}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C1.B1.A1"}}

	{{- $offset := mul $sizeOfFp 1}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C1.B2.A0"}}

	{{- $offset := mul $sizeOfFp 0}}
	{{- template "putFp" dict "all" . "OffSet" $offset "From" "_z.C1.B2.A1"}}

	return r[:]
}


// SetBytes interprets e as the bytes of a big-endian GT 
// sets z to that value (in Montgomery form), and returns z.
// size(e) == {{ $sizeOfFp }} * 12
// z.C1.B2.A1 | z.C1.B2.A0 | z.C1.B1.A1 | ...
func (z *e12) SetBytes(e []byte) error {
	if len(e) != sizeGT {
		return errors.New("invalid buffer size")
	}

	{{- $offset := mul $sizeOfFp 11}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C0.B0.A0"}}
	
	{{- $offset := mul $sizeOfFp 10}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C0.B0.A1"}}

	{{- $offset := mul $sizeOfFp 9}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C0.B1.A0"}}

	{{- $offset := mul $sizeOfFp 8}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C0.B1.A1"}}

	{{- $offset := mul $sizeOfFp 7}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C0.B2.A0"}}
	
	{{- $offset := mul $sizeOfFp 6}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C0.B2.A1"}}

	{{- $offset := mul $sizeOfFp 5}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C1.B0.A0"}}

	{{- $offset := mul $sizeOfFp 4}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C1.B0.A1"}}

	{{- $offset := mul $sizeOfFp 3}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C1.B1.A0"}}

	{{- $offset := mul $sizeOfFp 2}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C1.B1.A1"}}

	{{- $offset := mul $sizeOfFp 1}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C1.B2.A0"}}

	{{- $offset := mul $sizeOfFp 0}}
	{{- template "readFp" dict "all" . "OffSet" $offset "To" "z.C1.B2.A1"}}

	z.ToMont()

	return nil
}

{{define "putFp"}}
	{{- range $i := reverse .all.Fp.NbWordsIndexesFull}}
			{{- $j := mul $i 8}}
			{{- $j := add $j $.OffSet}}
			{{- $k := sub $.all.Fp.NbWords 1}}
			{{- $k := sub $k $i}}
			{{- $jj := add $j 8}}
			binary.BigEndian.PutUint64(r[{{$j}}:{{$jj}}], {{$.From}}[{{$k}}])
	{{- end}}
{{end}}

{{define "readFp"}}
	{{- range $i := reverse .all.Fp.NbWordsIndexesFull}}
			{{- $j := mul $i 8}}
			{{- $j := add $j $.OffSet}}
			{{- $k := sub $.all.Fp.NbWords 1}}
			{{- $k := sub $k $i}}
			{{- $jj := add $j 8}}
			{{$.To}}[{{$k}}] = binary.BigEndian.Uint64(e[{{$j}}:{{$jj}}])
	{{- end}}
{{end}}

`

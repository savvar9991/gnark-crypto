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

package fp

// expBySqrtExp is equivalent to z.Exp(x, /Users/gbotrel/dev/go/src/github.com/consensys/gnark-crypto/internal/generator/addchain/2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017fa01)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expBySqrtExp(x *Element) *Element {
	// addition chain:
	//
	//	_10       = 2*1
	//	_11       = 1 + _10
	//	_101      = _10 + _11
	//	_110      = 1 + _101
	//	_1011     = _101 + _110
	//	_1101     = _10 + _1011
	//	_10001    = _110 + _1011
	//	_10011    = _10 + _10001
	//	_11001    = _110 + _10011
	//	_11011    = _10 + _11001
	//	_11101    = _10 + _11011
	//	_11111    = _10 + _11101
	//	_100001   = _10 + _11111
	//	_100011   = _10 + _100001
	//	_100101   = _10 + _100011
	//	_101001   = _110 + _100011
	//	_101011   = _10 + _101001
	//	_101111   = _110 + _101001
	//	_110101   = _110 + _101111
	//	_110111   = _10 + _110101
	//	_111011   = _110 + _110101
	//	_111101   = _10 + _111011
	//	_111111   = _10 + _111101
	//	_1111110  = 2*_111111
	//	_1111111  = 1 + _1111110
	//	_10011000 = _11001 + _1111111
	//	i51       = ((_10011000 << 7 + _100011) << 3 + _101) << 13
	//	i68       = ((_101011 + i51) << 5 + _1011) << 9 + _11011
	//	i88       = ((i68 << 5 + _1011) << 5 + _101) << 8
	//	i105      = ((_1101 + i88) << 8 + _111111) << 6 + _11101
	//	i129      = ((i105 << 7 + _10011) << 9 + _101111) << 6
	//	i147      = ((_101001 + i129) << 6 + _11111) << 9 + _101111
	//	i168      = ((i147 << 7 + _101011) << 6 + _111101) << 6
	//	i194      = ((_111011 + i168) << 10 + _11101) << 13 + _110101
	//	i220      = ((i194 << 6 + _10001) << 8 + _111101) << 10
	//	i237      = ((_110101 + i220) << 2 + _11) << 12 + _100001
	//	i261      = ((i237 << 10 + _111101) << 5 + _11001) << 7
	//	i279      = ((_111011 + i261) << 7 + _100101) << 8 + _101011
	//	i304      = ((i279 << 6 + _110111) << 6 + _100101) << 11
	//	i317      = ((_10011 + i304) << 8 + _1111111) << 2 + 1
	//	i338      = 2*((i317 << 10 + 1) << 8 + _1111111)
	//	return      ((1 + i338) << 2 + 1) << 9 + 1
	//
	// Operations: 288 squares 64 multiplies

	// Allocate Temporaries.
	var (
		t0  = new(Element)
		t1  = new(Element)
		t2  = new(Element)
		t3  = new(Element)
		t4  = new(Element)
		t5  = new(Element)
		t6  = new(Element)
		t7  = new(Element)
		t8  = new(Element)
		t9  = new(Element)
		t10 = new(Element)
		t11 = new(Element)
		t12 = new(Element)
		t13 = new(Element)
		t14 = new(Element)
		t15 = new(Element)
		t16 = new(Element)
		t17 = new(Element)
		t18 = new(Element)
		t19 = new(Element)
		t20 = new(Element)
		t21 = new(Element)
	)

	// var t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15,t16,t17,t18,t19,t20,t21 Element
	// Step 1: z = x^0x2
	z.Square(x)

	// Step 2: t8 = x^0x3
	t8.Mul(x, z)

	// Step 3: t17 = x^0x5
	t17.Mul(z, t8)

	// Step 4: t4 = x^0x6
	t4.Mul(x, t17)

	// Step 5: t18 = x^0xb
	t18.Mul(t17, t4)

	// Step 6: t16 = x^0xd
	t16.Mul(z, t18)

	// Step 7: t10 = x^0x11
	t10.Mul(t4, t18)

	// Step 8: t0 = x^0x13
	t0.Mul(z, t10)

	// Step 9: t5 = x^0x19
	t5.Mul(t4, t0)

	// Step 10: t19 = x^0x1b
	t19.Mul(z, t5)

	// Step 11: t11 = x^0x1d
	t11.Mul(z, t19)

	// Step 12: t13 = x^0x1f
	t13.Mul(z, t11)

	// Step 13: t7 = x^0x21
	t7.Mul(z, t13)

	// Step 14: t20 = x^0x23
	t20.Mul(z, t7)

	// Step 15: t1 = x^0x25
	t1.Mul(z, t20)

	// Step 16: t14 = x^0x29
	t14.Mul(t4, t20)

	// Step 17: t3 = x^0x2b
	t3.Mul(z, t14)

	// Step 18: t12 = x^0x2f
	t12.Mul(t4, t14)

	// Step 19: t9 = x^0x35
	t9.Mul(t4, t12)

	// Step 20: t2 = x^0x37
	t2.Mul(z, t9)

	// Step 21: t4 = x^0x3b
	t4.Mul(t4, t9)

	// Step 22: t6 = x^0x3d
	t6.Mul(z, t4)

	// Step 23: t15 = x^0x3f
	t15.Mul(z, t6)

	// Step 24: z = x^0x7e
	z.Square(t15)

	// Step 25: z = x^0x7f
	z.Mul(x, z)

	// Step 26: t21 = x^0x98
	t21.Mul(t5, z)

	// Step 33: t21 = x^0x4c00
	for s := 0; s < 7; s++ {
		t21.Square(t21)
	}

	// Step 34: t20 = x^0x4c23
	t20.Mul(t20, t21)

	// Step 37: t20 = x^0x26118
	for s := 0; s < 3; s++ {
		t20.Square(t20)
	}

	// Step 38: t20 = x^0x2611d
	t20.Mul(t17, t20)

	// Step 51: t20 = x^0x4c23a000
	for s := 0; s < 13; s++ {
		t20.Square(t20)
	}

	// Step 52: t20 = x^0x4c23a02b
	t20.Mul(t3, t20)

	// Step 57: t20 = x^0x984740560
	for s := 0; s < 5; s++ {
		t20.Square(t20)
	}

	// Step 58: t20 = x^0x98474056b
	t20.Mul(t18, t20)

	// Step 67: t20 = x^0x1308e80ad600
	for s := 0; s < 9; s++ {
		t20.Square(t20)
	}

	// Step 68: t19 = x^0x1308e80ad61b
	t19.Mul(t19, t20)

	// Step 73: t19 = x^0x2611d015ac360
	for s := 0; s < 5; s++ {
		t19.Square(t19)
	}

	// Step 74: t18 = x^0x2611d015ac36b
	t18.Mul(t18, t19)

	// Step 79: t18 = x^0x4c23a02b586d60
	for s := 0; s < 5; s++ {
		t18.Square(t18)
	}

	// Step 80: t17 = x^0x4c23a02b586d65
	t17.Mul(t17, t18)

	// Step 88: t17 = x^0x4c23a02b586d6500
	for s := 0; s < 8; s++ {
		t17.Square(t17)
	}

	// Step 89: t16 = x^0x4c23a02b586d650d
	t16.Mul(t16, t17)

	// Step 97: t16 = x^0x4c23a02b586d650d00
	for s := 0; s < 8; s++ {
		t16.Square(t16)
	}

	// Step 98: t15 = x^0x4c23a02b586d650d3f
	t15.Mul(t15, t16)

	// Step 104: t15 = x^0x1308e80ad61b59434fc0
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 105: t15 = x^0x1308e80ad61b59434fdd
	t15.Mul(t11, t15)

	// Step 112: t15 = x^0x98474056b0daca1a7ee80
	for s := 0; s < 7; s++ {
		t15.Square(t15)
	}

	// Step 113: t15 = x^0x98474056b0daca1a7ee93
	t15.Mul(t0, t15)

	// Step 122: t15 = x^0x1308e80ad61b59434fdd2600
	for s := 0; s < 9; s++ {
		t15.Square(t15)
	}

	// Step 123: t15 = x^0x1308e80ad61b59434fdd262f
	t15.Mul(t12, t15)

	// Step 129: t15 = x^0x4c23a02b586d650d3f7498bc0
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 130: t14 = x^0x4c23a02b586d650d3f7498be9
	t14.Mul(t14, t15)

	// Step 136: t14 = x^0x1308e80ad61b59434fdd262fa40
	for s := 0; s < 6; s++ {
		t14.Square(t14)
	}

	// Step 137: t13 = x^0x1308e80ad61b59434fdd262fa5f
	t13.Mul(t13, t14)

	// Step 146: t13 = x^0x2611d015ac36b2869fba4c5f4be00
	for s := 0; s < 9; s++ {
		t13.Square(t13)
	}

	// Step 147: t12 = x^0x2611d015ac36b2869fba4c5f4be2f
	t12.Mul(t12, t13)

	// Step 154: t12 = x^0x1308e80ad61b59434fdd262fa5f1780
	for s := 0; s < 7; s++ {
		t12.Square(t12)
	}

	// Step 155: t12 = x^0x1308e80ad61b59434fdd262fa5f17ab
	t12.Mul(t3, t12)

	// Step 161: t12 = x^0x4c23a02b586d650d3f7498be97c5eac0
	for s := 0; s < 6; s++ {
		t12.Square(t12)
	}

	// Step 162: t12 = x^0x4c23a02b586d650d3f7498be97c5eafd
	t12.Mul(t6, t12)

	// Step 168: t12 = x^0x1308e80ad61b59434fdd262fa5f17abf40
	for s := 0; s < 6; s++ {
		t12.Square(t12)
	}

	// Step 169: t12 = x^0x1308e80ad61b59434fdd262fa5f17abf7b
	t12.Mul(t4, t12)

	// Step 179: t12 = x^0x4c23a02b586d650d3f7498be97c5eafdec00
	for s := 0; s < 10; s++ {
		t12.Square(t12)
	}

	// Step 180: t11 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d
	t11.Mul(t11, t12)

	// Step 193: t11 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a000
	for s := 0; s < 13; s++ {
		t11.Square(t11)
	}

	// Step 194: t11 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a035
	t11.Mul(t9, t11)

	// Step 200: t11 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d40
	for s := 0; s < 6; s++ {
		t11.Square(t11)
	}

	// Step 201: t10 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d51
	t10.Mul(t10, t11)

	// Step 209: t10 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d5100
	for s := 0; s < 8; s++ {
		t10.Square(t10)
	}

	// Step 210: t10 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d
	t10.Mul(t6, t10)

	// Step 220: t10 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f400
	for s := 0; s < 10; s++ {
		t10.Square(t10)
	}

	// Step 221: t9 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435
	t9.Mul(t9, t10)

	// Step 223: t9 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d4
	for s := 0; s < 2; s++ {
		t9.Square(t9)
	}

	// Step 224: t8 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d7
	t8.Mul(t8, t9)

	// Step 236: t8 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d7000
	for s := 0; s < 12; s++ {
		t8.Square(t8)
	}

	// Step 237: t7 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d7021
	t7.Mul(t7, t8)

	// Step 247: t7 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c08400
	for s := 0; s < 10; s++ {
		t7.Square(t7)
	}

	// Step 248: t6 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843d
	t6.Mul(t6, t7)

	// Step 253: t6 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087a0
	for s := 0; s < 5; s++ {
		t6.Square(t6)
	}

	// Step 254: t5 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b9
	t5.Mul(t5, t6)

	// Step 261: t5 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dc80
	for s := 0; s < 7; s++ {
		t5.Square(t5)
	}

	// Step 262: t4 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dcbb
	t4.Mul(t4, t5)

	// Step 269: t4 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5d80
	for s := 0; s < 7; s++ {
		t4.Square(t4)
	}

	// Step 270: t4 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da5
	t4.Mul(t1, t4)

	// Step 278: t4 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da500
	for s := 0; s < 8; s++ {
		t4.Square(t4)
	}

	// Step 279: t3 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52b
	t3.Mul(t3, t4)

	// Step 285: t3 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694ac0
	for s := 0; s < 6; s++ {
		t3.Square(t3)
	}

	// Step 286: t2 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694af7
	t2.Mul(t2, t3)

	// Step 292: t2 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bdc0
	for s := 0; s < 6; s++ {
		t2.Square(t2)
	}

	// Step 293: t1 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5
	t1.Mul(t1, t2)

	// Step 304: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef2800
	for s := 0; s < 11; s++ {
		t1.Square(t1)
	}

	// Step 305: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef2813
	t0.Mul(t0, t1)

	// Step 313: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef281300
	for s := 0; s < 8; s++ {
		t0.Square(t0)
	}

	// Step 314: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f
	t0.Mul(z, t0)

	// Step 316: t0 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dcbb4a57bca04dfc
	for s := 0; s < 2; s++ {
		t0.Square(t0)
	}

	// Step 317: t0 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dcbb4a57bca04dfd
	t0.Mul(x, t0)

	// Step 327: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f400
	for s := 0; s < 10; s++ {
		t0.Square(t0)
	}

	// Step 328: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f401
	t0.Mul(x, t0)

	// Step 336: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f40100
	for s := 0; s < 8; s++ {
		t0.Square(t0)
	}

	// Step 337: z = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017f
	z.Mul(z, t0)

	// Step 338: z = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5026fe802fe
	z.Square(z)

	// Step 339: z = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5026fe802ff
	z.Mul(x, z)

	// Step 341: z = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694af79409bfa00bfc
	for s := 0; s < 2; s++ {
		z.Square(z)
	}

	// Step 342: z = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694af79409bfa00bfd
	z.Mul(x, z)

	// Step 351: z = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017fa00
	for s := 0; s < 9; s++ {
		z.Square(z)
	}

	// Step 352: z = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017fa01
	z.Mul(x, z)

	return z
}

// expByLegendreExp is equivalent to z.Exp(x, /Users/gbotrel/dev/go/src/github.com/consensys/gnark-crypto/internal/generator/addchain/2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017fa0180000)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expByLegendreExp(x *Element) *Element {
	// addition chain:
	//
	//	_10       = 2*1
	//	_11       = 1 + _10
	//	_101      = _10 + _11
	//	_110      = 1 + _101
	//	_1011     = _101 + _110
	//	_1101     = _10 + _1011
	//	_10001    = _110 + _1011
	//	_10011    = _10 + _10001
	//	_11001    = _110 + _10011
	//	_11011    = _10 + _11001
	//	_11101    = _10 + _11011
	//	_11111    = _10 + _11101
	//	_100001   = _10 + _11111
	//	_100011   = _10 + _100001
	//	_100101   = _10 + _100011
	//	_101001   = _110 + _100011
	//	_101011   = _10 + _101001
	//	_101111   = _110 + _101001
	//	_110101   = _110 + _101111
	//	_110111   = _10 + _110101
	//	_111011   = _110 + _110101
	//	_111101   = _10 + _111011
	//	_111111   = _10 + _111101
	//	_1111110  = 2*_111111
	//	_1111111  = 1 + _1111110
	//	_10011000 = _11001 + _1111111
	//	i51       = ((_10011000 << 7 + _100011) << 3 + _101) << 13
	//	i68       = ((_101011 + i51) << 5 + _1011) << 9 + _11011
	//	i88       = ((i68 << 5 + _1011) << 5 + _101) << 8
	//	i105      = ((_1101 + i88) << 8 + _111111) << 6 + _11101
	//	i129      = ((i105 << 7 + _10011) << 9 + _101111) << 6
	//	i147      = ((_101001 + i129) << 6 + _11111) << 9 + _101111
	//	i168      = ((i147 << 7 + _101011) << 6 + _111101) << 6
	//	i194      = ((_111011 + i168) << 10 + _11101) << 13 + _110101
	//	i220      = ((i194 << 6 + _10001) << 8 + _111101) << 10
	//	i237      = ((_110101 + i220) << 2 + _11) << 12 + _100001
	//	i261      = ((i237 << 10 + _111101) << 5 + _11001) << 7
	//	i279      = ((_111011 + i261) << 7 + _100101) << 8 + _101011
	//	i304      = ((i279 << 6 + _110111) << 6 + _100101) << 11
	//	i317      = ((_10011 + i304) << 8 + _1111111) << 2 + 1
	//	i338      = 2*((i317 << 10 + 1) << 8 + _1111111)
	//	i353      = ((1 + i338) << 2 + 1) << 10 + _11
	//	return      i353 << 19
	//
	// Operations: 308 squares 64 multiplies

	// Allocate Temporaries.
	var (
		t0  = new(Element)
		t1  = new(Element)
		t2  = new(Element)
		t3  = new(Element)
		t4  = new(Element)
		t5  = new(Element)
		t6  = new(Element)
		t7  = new(Element)
		t8  = new(Element)
		t9  = new(Element)
		t10 = new(Element)
		t11 = new(Element)
		t12 = new(Element)
		t13 = new(Element)
		t14 = new(Element)
		t15 = new(Element)
		t16 = new(Element)
		t17 = new(Element)
		t18 = new(Element)
		t19 = new(Element)
		t20 = new(Element)
		t21 = new(Element)
	)

	// var t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15,t16,t17,t18,t19,t20,t21 Element
	// Step 1: t0 = x^0x2
	t0.Square(x)

	// Step 2: z = x^0x3
	z.Mul(x, t0)

	// Step 3: t17 = x^0x5
	t17.Mul(t0, z)

	// Step 4: t5 = x^0x6
	t5.Mul(x, t17)

	// Step 5: t18 = x^0xb
	t18.Mul(t17, t5)

	// Step 6: t16 = x^0xd
	t16.Mul(t0, t18)

	// Step 7: t10 = x^0x11
	t10.Mul(t5, t18)

	// Step 8: t1 = x^0x13
	t1.Mul(t0, t10)

	// Step 9: t6 = x^0x19
	t6.Mul(t5, t1)

	// Step 10: t19 = x^0x1b
	t19.Mul(t0, t6)

	// Step 11: t11 = x^0x1d
	t11.Mul(t0, t19)

	// Step 12: t13 = x^0x1f
	t13.Mul(t0, t11)

	// Step 13: t8 = x^0x21
	t8.Mul(t0, t13)

	// Step 14: t20 = x^0x23
	t20.Mul(t0, t8)

	// Step 15: t2 = x^0x25
	t2.Mul(t0, t20)

	// Step 16: t14 = x^0x29
	t14.Mul(t5, t20)

	// Step 17: t4 = x^0x2b
	t4.Mul(t0, t14)

	// Step 18: t12 = x^0x2f
	t12.Mul(t5, t14)

	// Step 19: t9 = x^0x35
	t9.Mul(t5, t12)

	// Step 20: t3 = x^0x37
	t3.Mul(t0, t9)

	// Step 21: t5 = x^0x3b
	t5.Mul(t5, t9)

	// Step 22: t7 = x^0x3d
	t7.Mul(t0, t5)

	// Step 23: t15 = x^0x3f
	t15.Mul(t0, t7)

	// Step 24: t0 = x^0x7e
	t0.Square(t15)

	// Step 25: t0 = x^0x7f
	t0.Mul(x, t0)

	// Step 26: t21 = x^0x98
	t21.Mul(t6, t0)

	// Step 33: t21 = x^0x4c00
	for s := 0; s < 7; s++ {
		t21.Square(t21)
	}

	// Step 34: t20 = x^0x4c23
	t20.Mul(t20, t21)

	// Step 37: t20 = x^0x26118
	for s := 0; s < 3; s++ {
		t20.Square(t20)
	}

	// Step 38: t20 = x^0x2611d
	t20.Mul(t17, t20)

	// Step 51: t20 = x^0x4c23a000
	for s := 0; s < 13; s++ {
		t20.Square(t20)
	}

	// Step 52: t20 = x^0x4c23a02b
	t20.Mul(t4, t20)

	// Step 57: t20 = x^0x984740560
	for s := 0; s < 5; s++ {
		t20.Square(t20)
	}

	// Step 58: t20 = x^0x98474056b
	t20.Mul(t18, t20)

	// Step 67: t20 = x^0x1308e80ad600
	for s := 0; s < 9; s++ {
		t20.Square(t20)
	}

	// Step 68: t19 = x^0x1308e80ad61b
	t19.Mul(t19, t20)

	// Step 73: t19 = x^0x2611d015ac360
	for s := 0; s < 5; s++ {
		t19.Square(t19)
	}

	// Step 74: t18 = x^0x2611d015ac36b
	t18.Mul(t18, t19)

	// Step 79: t18 = x^0x4c23a02b586d60
	for s := 0; s < 5; s++ {
		t18.Square(t18)
	}

	// Step 80: t17 = x^0x4c23a02b586d65
	t17.Mul(t17, t18)

	// Step 88: t17 = x^0x4c23a02b586d6500
	for s := 0; s < 8; s++ {
		t17.Square(t17)
	}

	// Step 89: t16 = x^0x4c23a02b586d650d
	t16.Mul(t16, t17)

	// Step 97: t16 = x^0x4c23a02b586d650d00
	for s := 0; s < 8; s++ {
		t16.Square(t16)
	}

	// Step 98: t15 = x^0x4c23a02b586d650d3f
	t15.Mul(t15, t16)

	// Step 104: t15 = x^0x1308e80ad61b59434fc0
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 105: t15 = x^0x1308e80ad61b59434fdd
	t15.Mul(t11, t15)

	// Step 112: t15 = x^0x98474056b0daca1a7ee80
	for s := 0; s < 7; s++ {
		t15.Square(t15)
	}

	// Step 113: t15 = x^0x98474056b0daca1a7ee93
	t15.Mul(t1, t15)

	// Step 122: t15 = x^0x1308e80ad61b59434fdd2600
	for s := 0; s < 9; s++ {
		t15.Square(t15)
	}

	// Step 123: t15 = x^0x1308e80ad61b59434fdd262f
	t15.Mul(t12, t15)

	// Step 129: t15 = x^0x4c23a02b586d650d3f7498bc0
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 130: t14 = x^0x4c23a02b586d650d3f7498be9
	t14.Mul(t14, t15)

	// Step 136: t14 = x^0x1308e80ad61b59434fdd262fa40
	for s := 0; s < 6; s++ {
		t14.Square(t14)
	}

	// Step 137: t13 = x^0x1308e80ad61b59434fdd262fa5f
	t13.Mul(t13, t14)

	// Step 146: t13 = x^0x2611d015ac36b2869fba4c5f4be00
	for s := 0; s < 9; s++ {
		t13.Square(t13)
	}

	// Step 147: t12 = x^0x2611d015ac36b2869fba4c5f4be2f
	t12.Mul(t12, t13)

	// Step 154: t12 = x^0x1308e80ad61b59434fdd262fa5f1780
	for s := 0; s < 7; s++ {
		t12.Square(t12)
	}

	// Step 155: t12 = x^0x1308e80ad61b59434fdd262fa5f17ab
	t12.Mul(t4, t12)

	// Step 161: t12 = x^0x4c23a02b586d650d3f7498be97c5eac0
	for s := 0; s < 6; s++ {
		t12.Square(t12)
	}

	// Step 162: t12 = x^0x4c23a02b586d650d3f7498be97c5eafd
	t12.Mul(t7, t12)

	// Step 168: t12 = x^0x1308e80ad61b59434fdd262fa5f17abf40
	for s := 0; s < 6; s++ {
		t12.Square(t12)
	}

	// Step 169: t12 = x^0x1308e80ad61b59434fdd262fa5f17abf7b
	t12.Mul(t5, t12)

	// Step 179: t12 = x^0x4c23a02b586d650d3f7498be97c5eafdec00
	for s := 0; s < 10; s++ {
		t12.Square(t12)
	}

	// Step 180: t11 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d
	t11.Mul(t11, t12)

	// Step 193: t11 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a000
	for s := 0; s < 13; s++ {
		t11.Square(t11)
	}

	// Step 194: t11 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a035
	t11.Mul(t9, t11)

	// Step 200: t11 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d40
	for s := 0; s < 6; s++ {
		t11.Square(t11)
	}

	// Step 201: t10 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d51
	t10.Mul(t10, t11)

	// Step 209: t10 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d5100
	for s := 0; s < 8; s++ {
		t10.Square(t10)
	}

	// Step 210: t10 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d
	t10.Mul(t7, t10)

	// Step 220: t10 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f400
	for s := 0; s < 10; s++ {
		t10.Square(t10)
	}

	// Step 221: t9 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435
	t9.Mul(t9, t10)

	// Step 223: t9 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d4
	for s := 0; s < 2; s++ {
		t9.Square(t9)
	}

	// Step 224: t9 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d7
	t9.Mul(z, t9)

	// Step 236: t9 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d7000
	for s := 0; s < 12; s++ {
		t9.Square(t9)
	}

	// Step 237: t8 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d7021
	t8.Mul(t8, t9)

	// Step 247: t8 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c08400
	for s := 0; s < 10; s++ {
		t8.Square(t8)
	}

	// Step 248: t7 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843d
	t7.Mul(t7, t8)

	// Step 253: t7 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087a0
	for s := 0; s < 5; s++ {
		t7.Square(t7)
	}

	// Step 254: t6 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b9
	t6.Mul(t6, t7)

	// Step 261: t6 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dc80
	for s := 0; s < 7; s++ {
		t6.Square(t6)
	}

	// Step 262: t5 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dcbb
	t5.Mul(t5, t6)

	// Step 269: t5 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5d80
	for s := 0; s < 7; s++ {
		t5.Square(t5)
	}

	// Step 270: t5 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da5
	t5.Mul(t2, t5)

	// Step 278: t5 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da500
	for s := 0; s < 8; s++ {
		t5.Square(t5)
	}

	// Step 279: t4 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52b
	t4.Mul(t4, t5)

	// Step 285: t4 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694ac0
	for s := 0; s < 6; s++ {
		t4.Square(t4)
	}

	// Step 286: t3 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694af7
	t3.Mul(t3, t4)

	// Step 292: t3 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bdc0
	for s := 0; s < 6; s++ {
		t3.Square(t3)
	}

	// Step 293: t2 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5
	t2.Mul(t2, t3)

	// Step 304: t2 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef2800
	for s := 0; s < 11; s++ {
		t2.Square(t2)
	}

	// Step 305: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef2813
	t1.Mul(t1, t2)

	// Step 313: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef281300
	for s := 0; s < 8; s++ {
		t1.Square(t1)
	}

	// Step 314: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f
	t1.Mul(t0, t1)

	// Step 316: t1 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dcbb4a57bca04dfc
	for s := 0; s < 2; s++ {
		t1.Square(t1)
	}

	// Step 317: t1 = x^0x98474056b0daca1a7ee9317d2f8bd5fbd83a03544f435c0843dcbb4a57bca04dfd
	t1.Mul(x, t1)

	// Step 327: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f400
	for s := 0; s < 10; s++ {
		t1.Square(t1)
	}

	// Step 328: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f401
	t1.Mul(x, t1)

	// Step 336: t1 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f40100
	for s := 0; s < 8; s++ {
		t1.Square(t1)
	}

	// Step 337: t0 = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017f
	t0.Mul(t0, t1)

	// Step 338: t0 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5026fe802fe
	t0.Square(t0)

	// Step 339: t0 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5026fe802ff
	t0.Mul(x, t0)

	// Step 341: t0 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694af79409bfa00bfc
	for s := 0; s < 2; s++ {
		t0.Square(t0)
	}

	// Step 342: t0 = x^0x1308e80ad61b59434fdd262fa5f17abf7b07406a89e86b81087b97694af79409bfa00bfd
	t0.Mul(x, t0)

	// Step 352: t0 = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5026fe802ff400
	for s := 0; s < 10; s++ {
		t0.Square(t0)
	}

	// Step 353: z = x^0x4c23a02b586d650d3f7498be97c5eafdec1d01aa27a1ae0421ee5da52bde5026fe802ff403
	z.Mul(z, t0)

	// Step 372: z = x^0x2611d015ac36b2869fba4c5f4be2f57ef60e80d513d0d70210f72ed295ef28137f4017fa0180000
	for s := 0; s < 19; s++ {
		z.Square(z)
	}

	return z
}

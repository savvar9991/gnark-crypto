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

package fr

// expBySqrtExp is equivalent to z.Exp(x, fbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff3a5fddaa08b0000265228)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expBySqrtExp(x Element) *Element {
	// addition chain:
	//
	//	_10      = 2*1
	//	_11      = 1 + _10
	//	_101     = _10 + _11
	//	_110     = 1 + _101
	//	_1001    = _11 + _110
	//	_1011    = _10 + _1001
	//	_1100    = 1 + _1011
	//	_1101    = 1 + _1100
	//	_10001   = _101 + _1100
	//	_10011   = _10 + _10001
	//	_10101   = _10 + _10011
	//	_11011   = _110 + _10101
	//	_11101   = _10 + _11011
	//	_100011  = _110 + _11101
	//	_100111  = _1100 + _11011
	//	_101001  = _10 + _100111
	//	_110101  = _1100 + _101001
	//	_110111  = _10 + _110101
	//	_111001  = _10 + _110111
	//	_111011  = _10 + _111001
	//	_111101  = _10 + _111011
	//	_111111  = _10 + _111101
	//	_1111100 = _111101 + _111111
	//	_1111111 = _11 + _1111100
	//	i39      = ((_1111100 << 4 + _11101) << 3 + _11) << 6
	//	i57      = ((1 + i39) << 9 + _1011) << 6 + _1101
	//	i78      = ((i57 << 9 + _10011) << 7 + _100011) << 3
	//	i98      = ((1 + i78) << 11 + _101001) << 6 + _111001
	//	i117     = ((i98 << 7 + _110101) << 4 + _1101) << 6
	//	i138     = ((_1001 + i117) << 12 + _111011) << 6 + _10001
	//	i162     = ((i138 << 11 + _111101) << 6 + _101) << 5
	//	i184     = ((1 + i162) << 11 + _1011) << 8 + _111101
	//	i205     = ((i184 << 6 + _11011) << 8 + _100011) << 5
	//	i227     = ((_10001 + i205) << 12 + _100011) << 7 + _10011
	//	i257     = ((i227 << 6 + _10011) << 13 + _110111) << 9
	//	i279     = ((_11011 + i257) << 9 + _1101) << 10 + _101001
	//	i299     = ((i279 << 8 + _100111) << 2 + 1) << 8
	//	i311     = ((_1111111 + i299) << 2 + _11) << 7 + _11101
	//	i331     = ((i311 << 3 + 1) << 8 + _1111111) << 7
	//	i350     = ((_111011 + i331) << 6 + _10101) << 10 + _10001
	//	i386     = ((i350 << 3 + _11) << 23 + _10011) << 8
	//	return     ((_101001 + i386) << 6 + _101) << 3
	//
	// Operations: 330 squares 67 multiplies

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
	)

	// var t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15,t16,t17,t18 Element
	// Step 1: t6 = x^0x2
	t6.Square(&x)

	// Step 2: t2 = x^0x3
	t2.Mul(&x, t6)

	// Step 3: z = x^0x5
	z.Mul(t6, t2)

	// Step 4: t0 = x^0x6
	t0.Mul(&x, z)

	// Step 5: t15 = x^0x9
	t15.Mul(t2, t0)

	// Step 6: t14 = x^0xb
	t14.Mul(t6, t15)

	// Step 7: t5 = x^0xc
	t5.Mul(&x, t14)

	// Step 8: t9 = x^0xd
	t9.Mul(&x, t5)

	// Step 9: t3 = x^0x11
	t3.Mul(z, t5)

	// Step 10: t1 = x^0x13
	t1.Mul(t6, t3)

	// Step 11: t4 = x^0x15
	t4.Mul(t6, t1)

	// Step 12: t10 = x^0x1b
	t10.Mul(t0, t4)

	// Step 13: t7 = x^0x1d
	t7.Mul(t6, t10)

	// Step 14: t12 = x^0x23
	t12.Mul(t0, t7)

	// Step 15: t8 = x^0x27
	t8.Mul(t5, t10)

	// Step 16: t0 = x^0x29
	t0.Mul(t6, t8)

	// Step 17: t16 = x^0x35
	t16.Mul(t5, t0)

	// Step 18: t11 = x^0x37
	t11.Mul(t6, t16)

	// Step 19: t17 = x^0x39
	t17.Mul(t6, t11)

	// Step 20: t5 = x^0x3b
	t5.Mul(t6, t17)

	// Step 21: t13 = x^0x3d
	t13.Mul(t6, t5)

	// Step 22: t6 = x^0x3f
	t6.Mul(t6, t13)

	// Step 23: t18 = x^0x7c
	t18.Mul(t13, t6)

	// Step 24: t6 = x^0x7f
	t6.Mul(t2, t18)

	// Step 28: t18 = x^0x7c0
	for s := 0; s < 4; s++ {
		t18.Square(t18)
	}

	// Step 29: t18 = x^0x7dd
	t18.Mul(t7, t18)

	// Step 32: t18 = x^0x3ee8
	for s := 0; s < 3; s++ {
		t18.Square(t18)
	}

	// Step 33: t18 = x^0x3eeb
	t18.Mul(t2, t18)

	// Step 39: t18 = x^0xfbac0
	for s := 0; s < 6; s++ {
		t18.Square(t18)
	}

	// Step 40: t18 = x^0xfbac1
	t18.Mul(&x, t18)

	// Step 49: t18 = x^0x1f758200
	for s := 0; s < 9; s++ {
		t18.Square(t18)
	}

	// Step 50: t18 = x^0x1f75820b
	t18.Mul(t14, t18)

	// Step 56: t18 = x^0x7dd6082c0
	for s := 0; s < 6; s++ {
		t18.Square(t18)
	}

	// Step 57: t18 = x^0x7dd6082cd
	t18.Mul(t9, t18)

	// Step 66: t18 = x^0xfbac1059a00
	for s := 0; s < 9; s++ {
		t18.Square(t18)
	}

	// Step 67: t18 = x^0xfbac1059a13
	t18.Mul(t1, t18)

	// Step 74: t18 = x^0x7dd6082cd0980
	for s := 0; s < 7; s++ {
		t18.Square(t18)
	}

	// Step 75: t18 = x^0x7dd6082cd09a3
	t18.Mul(t12, t18)

	// Step 78: t18 = x^0x3eeb0416684d18
	for s := 0; s < 3; s++ {
		t18.Square(t18)
	}

	// Step 79: t18 = x^0x3eeb0416684d19
	t18.Mul(&x, t18)

	// Step 90: t18 = x^0x1f75820b34268c800
	for s := 0; s < 11; s++ {
		t18.Square(t18)
	}

	// Step 91: t18 = x^0x1f75820b34268c829
	t18.Mul(t0, t18)

	// Step 97: t18 = x^0x7dd6082cd09a320a40
	for s := 0; s < 6; s++ {
		t18.Square(t18)
	}

	// Step 98: t17 = x^0x7dd6082cd09a320a79
	t17.Mul(t17, t18)

	// Step 105: t17 = x^0x3eeb0416684d19053c80
	for s := 0; s < 7; s++ {
		t17.Square(t17)
	}

	// Step 106: t16 = x^0x3eeb0416684d19053cb5
	t16.Mul(t16, t17)

	// Step 110: t16 = x^0x3eeb0416684d19053cb50
	for s := 0; s < 4; s++ {
		t16.Square(t16)
	}

	// Step 111: t16 = x^0x3eeb0416684d19053cb5d
	t16.Mul(t9, t16)

	// Step 117: t16 = x^0xfbac1059a1346414f2d740
	for s := 0; s < 6; s++ {
		t16.Square(t16)
	}

	// Step 118: t15 = x^0xfbac1059a1346414f2d749
	t15.Mul(t15, t16)

	// Step 130: t15 = x^0xfbac1059a1346414f2d749000
	for s := 0; s < 12; s++ {
		t15.Square(t15)
	}

	// Step 131: t15 = x^0xfbac1059a1346414f2d74903b
	t15.Mul(t5, t15)

	// Step 137: t15 = x^0x3eeb0416684d19053cb5d240ec0
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 138: t15 = x^0x3eeb0416684d19053cb5d240ed1
	t15.Mul(t3, t15)

	// Step 149: t15 = x^0x1f75820b34268c829e5ae920768800
	for s := 0; s < 11; s++ {
		t15.Square(t15)
	}

	// Step 150: t15 = x^0x1f75820b34268c829e5ae92076883d
	t15.Mul(t13, t15)

	// Step 156: t15 = x^0x7dd6082cd09a320a796ba481da20f40
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 157: t15 = x^0x7dd6082cd09a320a796ba481da20f45
	t15.Mul(z, t15)

	// Step 162: t15 = x^0xfbac1059a1346414f2d74903b441e8a0
	for s := 0; s < 5; s++ {
		t15.Square(t15)
	}

	// Step 163: t15 = x^0xfbac1059a1346414f2d74903b441e8a1
	t15.Mul(&x, t15)

	// Step 174: t15 = x^0x7dd6082cd09a320a796ba481da20f450800
	for s := 0; s < 11; s++ {
		t15.Square(t15)
	}

	// Step 175: t14 = x^0x7dd6082cd09a320a796ba481da20f45080b
	t14.Mul(t14, t15)

	// Step 183: t14 = x^0x7dd6082cd09a320a796ba481da20f45080b00
	for s := 0; s < 8; s++ {
		t14.Square(t14)
	}

	// Step 184: t13 = x^0x7dd6082cd09a320a796ba481da20f45080b3d
	t13.Mul(t13, t14)

	// Step 190: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf40
	for s := 0; s < 6; s++ {
		t13.Square(t13)
	}

	// Step 191: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf5b
	t13.Mul(t10, t13)

	// Step 199: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf5b00
	for s := 0; s < 8; s++ {
		t13.Square(t13)
	}

	// Step 200: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf5b23
	t13.Mul(t12, t13)

	// Step 205: t13 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6460
	for s := 0; s < 5; s++ {
		t13.Square(t13)
	}

	// Step 206: t13 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6471
	t13.Mul(t3, t13)

	// Step 218: t13 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6471000
	for s := 0; s < 12; s++ {
		t13.Square(t13)
	}

	// Step 219: t12 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6471023
	t12.Mul(t12, t13)

	// Step 226: t12 = x^0x1f75820b34268c829e5ae92076883d14202cf5b23881180
	for s := 0; s < 7; s++ {
		t12.Square(t12)
	}

	// Step 227: t12 = x^0x1f75820b34268c829e5ae92076883d14202cf5b23881193
	t12.Mul(t1, t12)

	// Step 233: t12 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464c0
	for s := 0; s < 6; s++ {
		t12.Square(t12)
	}

	// Step 234: t12 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d3
	t12.Mul(t1, t12)

	// Step 247: t12 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a6000
	for s := 0; s < 13; s++ {
		t12.Square(t12)
	}

	// Step 248: t11 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a6037
	t11.Mul(t11, t12)

	// Step 257: t11 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e00
	for s := 0; s < 9; s++ {
		t11.Square(t11)
	}

	// Step 258: t10 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b
	t10.Mul(t10, t11)

	// Step 267: t10 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc3600
	for s := 0; s < 9; s++ {
		t10.Square(t10)
	}

	// Step 268: t9 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d
	t9.Mul(t9, t10)

	// Step 278: t9 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83400
	for s := 0; s < 10; s++ {
		t9.Square(t9)
	}

	// Step 279: t9 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429
	t9.Mul(t0, t9)

	// Step 287: t9 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d8342900
	for s := 0; s < 8; s++ {
		t9.Square(t9)
	}

	// Step 288: t8 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d8342927
	t8.Mul(t8, t9)

	// Step 290: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49c
	for s := 0; s < 2; s++ {
		t8.Square(t8)
	}

	// Step 291: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d
	t8.Mul(&x, t8)

	// Step 299: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d00
	for s := 0; s < 8; s++ {
		t8.Square(t8)
	}

	// Step 300: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7f
	t8.Mul(t6, t8)

	// Step 302: t8 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275fc
	for s := 0; s < 2; s++ {
		t8.Square(t8)
	}

	// Step 303: t8 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff
	t8.Mul(t2, t8)

	// Step 310: t8 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff80
	for s := 0; s < 7; s++ {
		t8.Square(t8)
	}

	// Step 311: t7 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d
	t7.Mul(t7, t8)

	// Step 314: t7 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce8
	for s := 0; s < 3; s++ {
		t7.Square(t7)
	}

	// Step 315: t7 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce9
	t7.Mul(&x, t7)

	// Step 323: t7 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce900
	for s := 0; s < 8; s++ {
		t7.Square(t7)
	}

	// Step 324: t6 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce97f
	t6.Mul(t6, t7)

	// Step 331: t6 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bf80
	for s := 0; s < 7; s++ {
		t6.Square(t6)
	}

	// Step 332: t5 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb
	t5.Mul(t5, t6)

	// Step 338: t5 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feec0
	for s := 0; s < 6; s++ {
		t5.Square(t5)
	}

	// Step 339: t4 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed5
	t4.Mul(t4, t5)

	// Step 349: t4 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5400
	for s := 0; s < 10; s++ {
		t4.Square(t4)
	}

	// Step 350: t3 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411
	t3.Mul(t3, t4)

	// Step 353: t3 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff3a5fddaa088
	for s := 0; s < 3; s++ {
		t3.Square(t3)
	}

	// Step 354: t2 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff3a5fddaa08b
	t2.Mul(t2, t3)

	// Step 377: t2 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed5045800000
	for s := 0; s < 23; s++ {
		t2.Square(t2)
	}

	// Step 378: t1 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed5045800013
	t1.Mul(t1, t2)

	// Step 386: t1 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed504580001300
	for s := 0; s < 8; s++ {
		t1.Square(t1)
	}

	// Step 387: t0 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed504580001329
	t0.Mul(t0, t1)

	// Step 393: t0 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca40
	for s := 0; s < 6; s++ {
		t0.Square(t0)
	}

	// Step 394: z = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca45
	z.Mul(z, t0)

	// Step 397: z = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff3a5fddaa08b0000265228
	for s := 0; s < 3; s++ {
		z.Square(z)
	}

	return z
}

// expByLegendreExp is equivalent to z.Exp(x, 1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca4510000000000)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expByLegendreExp(x Element) *Element {
	// addition chain:
	//
	//	_10      = 2*1
	//	_11      = 1 + _10
	//	_101     = _10 + _11
	//	_110     = 1 + _101
	//	_1001    = _11 + _110
	//	_1011    = _10 + _1001
	//	_1100    = 1 + _1011
	//	_1101    = 1 + _1100
	//	_10001   = _101 + _1100
	//	_10011   = _10 + _10001
	//	_10101   = _10 + _10011
	//	_11011   = _110 + _10101
	//	_11101   = _10 + _11011
	//	_100011  = _110 + _11101
	//	_100111  = _1100 + _11011
	//	_101001  = _10 + _100111
	//	_110101  = _1100 + _101001
	//	_110111  = _10 + _110101
	//	_111001  = _10 + _110111
	//	_111011  = _10 + _111001
	//	_111101  = _10 + _111011
	//	_111111  = _10 + _111101
	//	_1111100 = _111101 + _111111
	//	_1111111 = _11 + _1111100
	//	i39      = ((_1111100 << 4 + _11101) << 3 + _11) << 6
	//	i57      = ((1 + i39) << 9 + _1011) << 6 + _1101
	//	i78      = ((i57 << 9 + _10011) << 7 + _100011) << 3
	//	i98      = ((1 + i78) << 11 + _101001) << 6 + _111001
	//	i117     = ((i98 << 7 + _110101) << 4 + _1101) << 6
	//	i138     = ((_1001 + i117) << 12 + _111011) << 6 + _10001
	//	i162     = ((i138 << 11 + _111101) << 6 + _101) << 5
	//	i184     = ((1 + i162) << 11 + _1011) << 8 + _111101
	//	i205     = ((i184 << 6 + _11011) << 8 + _100011) << 5
	//	i227     = ((_10001 + i205) << 12 + _100011) << 7 + _10011
	//	i257     = ((i227 << 6 + _10011) << 13 + _110111) << 9
	//	i279     = ((_11011 + i257) << 9 + _1101) << 10 + _101001
	//	i299     = ((i279 << 8 + _100111) << 2 + 1) << 8
	//	i311     = ((_1111111 + i299) << 2 + _11) << 7 + _11101
	//	i331     = ((i311 << 3 + 1) << 8 + _1111111) << 7
	//	i350     = ((_111011 + i331) << 6 + _10101) << 10 + _10001
	//	i386     = ((i350 << 3 + _11) << 23 + _10011) << 8
	//	i399     = ((_101001 + i386) << 6 + _101) << 4 + 1
	//	return     i399 << 40
	//
	// Operations: 371 squares 68 multiplies

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
	)

	// var t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15,t16,t17,t18 Element
	// Step 1: t6 = x^0x2
	t6.Square(&x)

	// Step 2: t2 = x^0x3
	t2.Mul(&x, t6)

	// Step 3: z = x^0x5
	z.Mul(t6, t2)

	// Step 4: t0 = x^0x6
	t0.Mul(&x, z)

	// Step 5: t15 = x^0x9
	t15.Mul(t2, t0)

	// Step 6: t14 = x^0xb
	t14.Mul(t6, t15)

	// Step 7: t5 = x^0xc
	t5.Mul(&x, t14)

	// Step 8: t9 = x^0xd
	t9.Mul(&x, t5)

	// Step 9: t3 = x^0x11
	t3.Mul(z, t5)

	// Step 10: t1 = x^0x13
	t1.Mul(t6, t3)

	// Step 11: t4 = x^0x15
	t4.Mul(t6, t1)

	// Step 12: t10 = x^0x1b
	t10.Mul(t0, t4)

	// Step 13: t7 = x^0x1d
	t7.Mul(t6, t10)

	// Step 14: t12 = x^0x23
	t12.Mul(t0, t7)

	// Step 15: t8 = x^0x27
	t8.Mul(t5, t10)

	// Step 16: t0 = x^0x29
	t0.Mul(t6, t8)

	// Step 17: t16 = x^0x35
	t16.Mul(t5, t0)

	// Step 18: t11 = x^0x37
	t11.Mul(t6, t16)

	// Step 19: t17 = x^0x39
	t17.Mul(t6, t11)

	// Step 20: t5 = x^0x3b
	t5.Mul(t6, t17)

	// Step 21: t13 = x^0x3d
	t13.Mul(t6, t5)

	// Step 22: t6 = x^0x3f
	t6.Mul(t6, t13)

	// Step 23: t18 = x^0x7c
	t18.Mul(t13, t6)

	// Step 24: t6 = x^0x7f
	t6.Mul(t2, t18)

	// Step 28: t18 = x^0x7c0
	for s := 0; s < 4; s++ {
		t18.Square(t18)
	}

	// Step 29: t18 = x^0x7dd
	t18.Mul(t7, t18)

	// Step 32: t18 = x^0x3ee8
	for s := 0; s < 3; s++ {
		t18.Square(t18)
	}

	// Step 33: t18 = x^0x3eeb
	t18.Mul(t2, t18)

	// Step 39: t18 = x^0xfbac0
	for s := 0; s < 6; s++ {
		t18.Square(t18)
	}

	// Step 40: t18 = x^0xfbac1
	t18.Mul(&x, t18)

	// Step 49: t18 = x^0x1f758200
	for s := 0; s < 9; s++ {
		t18.Square(t18)
	}

	// Step 50: t18 = x^0x1f75820b
	t18.Mul(t14, t18)

	// Step 56: t18 = x^0x7dd6082c0
	for s := 0; s < 6; s++ {
		t18.Square(t18)
	}

	// Step 57: t18 = x^0x7dd6082cd
	t18.Mul(t9, t18)

	// Step 66: t18 = x^0xfbac1059a00
	for s := 0; s < 9; s++ {
		t18.Square(t18)
	}

	// Step 67: t18 = x^0xfbac1059a13
	t18.Mul(t1, t18)

	// Step 74: t18 = x^0x7dd6082cd0980
	for s := 0; s < 7; s++ {
		t18.Square(t18)
	}

	// Step 75: t18 = x^0x7dd6082cd09a3
	t18.Mul(t12, t18)

	// Step 78: t18 = x^0x3eeb0416684d18
	for s := 0; s < 3; s++ {
		t18.Square(t18)
	}

	// Step 79: t18 = x^0x3eeb0416684d19
	t18.Mul(&x, t18)

	// Step 90: t18 = x^0x1f75820b34268c800
	for s := 0; s < 11; s++ {
		t18.Square(t18)
	}

	// Step 91: t18 = x^0x1f75820b34268c829
	t18.Mul(t0, t18)

	// Step 97: t18 = x^0x7dd6082cd09a320a40
	for s := 0; s < 6; s++ {
		t18.Square(t18)
	}

	// Step 98: t17 = x^0x7dd6082cd09a320a79
	t17.Mul(t17, t18)

	// Step 105: t17 = x^0x3eeb0416684d19053c80
	for s := 0; s < 7; s++ {
		t17.Square(t17)
	}

	// Step 106: t16 = x^0x3eeb0416684d19053cb5
	t16.Mul(t16, t17)

	// Step 110: t16 = x^0x3eeb0416684d19053cb50
	for s := 0; s < 4; s++ {
		t16.Square(t16)
	}

	// Step 111: t16 = x^0x3eeb0416684d19053cb5d
	t16.Mul(t9, t16)

	// Step 117: t16 = x^0xfbac1059a1346414f2d740
	for s := 0; s < 6; s++ {
		t16.Square(t16)
	}

	// Step 118: t15 = x^0xfbac1059a1346414f2d749
	t15.Mul(t15, t16)

	// Step 130: t15 = x^0xfbac1059a1346414f2d749000
	for s := 0; s < 12; s++ {
		t15.Square(t15)
	}

	// Step 131: t15 = x^0xfbac1059a1346414f2d74903b
	t15.Mul(t5, t15)

	// Step 137: t15 = x^0x3eeb0416684d19053cb5d240ec0
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 138: t15 = x^0x3eeb0416684d19053cb5d240ed1
	t15.Mul(t3, t15)

	// Step 149: t15 = x^0x1f75820b34268c829e5ae920768800
	for s := 0; s < 11; s++ {
		t15.Square(t15)
	}

	// Step 150: t15 = x^0x1f75820b34268c829e5ae92076883d
	t15.Mul(t13, t15)

	// Step 156: t15 = x^0x7dd6082cd09a320a796ba481da20f40
	for s := 0; s < 6; s++ {
		t15.Square(t15)
	}

	// Step 157: t15 = x^0x7dd6082cd09a320a796ba481da20f45
	t15.Mul(z, t15)

	// Step 162: t15 = x^0xfbac1059a1346414f2d74903b441e8a0
	for s := 0; s < 5; s++ {
		t15.Square(t15)
	}

	// Step 163: t15 = x^0xfbac1059a1346414f2d74903b441e8a1
	t15.Mul(&x, t15)

	// Step 174: t15 = x^0x7dd6082cd09a320a796ba481da20f450800
	for s := 0; s < 11; s++ {
		t15.Square(t15)
	}

	// Step 175: t14 = x^0x7dd6082cd09a320a796ba481da20f45080b
	t14.Mul(t14, t15)

	// Step 183: t14 = x^0x7dd6082cd09a320a796ba481da20f45080b00
	for s := 0; s < 8; s++ {
		t14.Square(t14)
	}

	// Step 184: t13 = x^0x7dd6082cd09a320a796ba481da20f45080b3d
	t13.Mul(t13, t14)

	// Step 190: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf40
	for s := 0; s < 6; s++ {
		t13.Square(t13)
	}

	// Step 191: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf5b
	t13.Mul(t10, t13)

	// Step 199: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf5b00
	for s := 0; s < 8; s++ {
		t13.Square(t13)
	}

	// Step 200: t13 = x^0x1f75820b34268c829e5ae92076883d14202cf5b23
	t13.Mul(t12, t13)

	// Step 205: t13 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6460
	for s := 0; s < 5; s++ {
		t13.Square(t13)
	}

	// Step 206: t13 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6471
	t13.Mul(t3, t13)

	// Step 218: t13 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6471000
	for s := 0; s < 12; s++ {
		t13.Square(t13)
	}

	// Step 219: t12 = x^0x3eeb0416684d19053cb5d240ed107a284059eb6471023
	t12.Mul(t12, t13)

	// Step 226: t12 = x^0x1f75820b34268c829e5ae92076883d14202cf5b23881180
	for s := 0; s < 7; s++ {
		t12.Square(t12)
	}

	// Step 227: t12 = x^0x1f75820b34268c829e5ae92076883d14202cf5b23881193
	t12.Mul(t1, t12)

	// Step 233: t12 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464c0
	for s := 0; s < 6; s++ {
		t12.Square(t12)
	}

	// Step 234: t12 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d3
	t12.Mul(t1, t12)

	// Step 247: t12 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a6000
	for s := 0; s < 13; s++ {
		t12.Square(t12)
	}

	// Step 248: t11 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a6037
	t11.Mul(t11, t12)

	// Step 257: t11 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e00
	for s := 0; s < 9; s++ {
		t11.Square(t11)
	}

	// Step 258: t10 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b
	t10.Mul(t10, t11)

	// Step 267: t10 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc3600
	for s := 0; s < 9; s++ {
		t10.Square(t10)
	}

	// Step 268: t9 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d
	t9.Mul(t9, t10)

	// Step 278: t9 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83400
	for s := 0; s < 10; s++ {
		t9.Square(t9)
	}

	// Step 279: t9 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429
	t9.Mul(t0, t9)

	// Step 287: t9 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d8342900
	for s := 0; s < 8; s++ {
		t9.Square(t9)
	}

	// Step 288: t8 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d8342927
	t8.Mul(t8, t9)

	// Step 290: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49c
	for s := 0; s < 2; s++ {
		t8.Square(t8)
	}

	// Step 291: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d
	t8.Mul(&x, t8)

	// Step 299: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d00
	for s := 0; s < 8; s++ {
		t8.Square(t8)
	}

	// Step 300: t8 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7f
	t8.Mul(t6, t8)

	// Step 302: t8 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275fc
	for s := 0; s < 2; s++ {
		t8.Square(t8)
	}

	// Step 303: t8 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff
	t8.Mul(t2, t8)

	// Step 310: t8 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff80
	for s := 0; s < 7; s++ {
		t8.Square(t8)
	}

	// Step 311: t7 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d
	t7.Mul(t7, t8)

	// Step 314: t7 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce8
	for s := 0; s < 3; s++ {
		t7.Square(t7)
	}

	// Step 315: t7 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce9
	t7.Mul(&x, t7)

	// Step 323: t7 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce900
	for s := 0; s < 8; s++ {
		t7.Square(t7)
	}

	// Step 324: t6 = x^0x3eeb0416684d19053cb5d240ed107a284059eb647102326980dc360d0a49d7fce97f
	t6.Mul(t6, t7)

	// Step 331: t6 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bf80
	for s := 0; s < 7; s++ {
		t6.Square(t6)
	}

	// Step 332: t5 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb
	t5.Mul(t5, t6)

	// Step 338: t5 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feec0
	for s := 0; s < 6; s++ {
		t5.Square(t5)
	}

	// Step 339: t4 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed5
	t4.Mul(t4, t5)

	// Step 349: t4 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5400
	for s := 0; s < 10; s++ {
		t4.Square(t4)
	}

	// Step 350: t3 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411
	t3.Mul(t3, t4)

	// Step 353: t3 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff3a5fddaa088
	for s := 0; s < 3; s++ {
		t3.Square(t3)
	}

	// Step 354: t2 = x^0xfbac1059a1346414f2d74903b441e8a10167ad91c408c9a60370d83429275ff3a5fddaa08b
	t2.Mul(t2, t3)

	// Step 377: t2 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed5045800000
	for s := 0; s < 23; s++ {
		t2.Square(t2)
	}

	// Step 378: t1 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed5045800013
	t1.Mul(t1, t2)

	// Step 386: t1 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed504580001300
	for s := 0; s < 8; s++ {
		t1.Square(t1)
	}

	// Step 387: t0 = x^0x7dd6082cd09a320a796ba481da20f45080b3d6c8e20464d301b86c1a1493aff9d2feed504580001329
	t0.Mul(t0, t1)

	// Step 393: t0 = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca40
	for s := 0; s < 6; s++ {
		t0.Square(t0)
	}

	// Step 394: z = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca45
	z.Mul(z, t0)

	// Step 398: z = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca450
	for s := 0; s < 4; s++ {
		z.Square(z)
	}

	// Step 399: z = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca451
	z.Mul(&x, z)

	// Step 439: z = x^0x1f75820b34268c829e5ae92076883d14202cf5b238811934c06e1b068524ebfe74bfbb5411600004ca4510000000000
	for s := 0; s < 40; s++ {
		z.Square(z)
	}

	return z
}
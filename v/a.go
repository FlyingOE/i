package main

import (
	. "github.com/ktye/wg/module"
)

func init() {
	Memory(1)
	Memory2(1)
	Data(132, "\x00\x01@\x01\x01\x01\x01\t\x10`\x01\x01\x01\x01\x01\t\xc4\xc4\xc4\xc4\xc4\xc4\xc4\xc4\xc4\xc4\x01 \x01\x01\x01\x01\x01BBBBBBBBBBBBBBBBBBBBBBBBBB\x10\t`\x01\x01\x00\xc2\xc2\xc2\xc2\xc2\xc2BBBBBBBBBBBBBBBBBBBB\x10\x01`\x01") // k_test.go: TestClass
	Data(227, ":+-*%&|<>=~!,^#_$?@.':/:\\:vbcisfzldtmdplx00BCISFZLDT0")
	Export(nyi, Asn, Atx, Cal, cs, dx, Kc, Kf, Ki, kinit, l2, mk, nn, repl, rx, sc, src, tp, trap, Val)

	//            0    :    +    -    *    %    &    |    <    >    =10   ~    !    ,    ^    #    _    $    ?    @    .20  '    ':   /    /:   \    \:                  30                       35                       40                       45
	Functions(00, nul, Idy, Flp, Neg, Fst, Sqr, Wer, Rev, Asc, Dsc, Grp, Not, Til, Enl, Srt, Cnt, Flr, Str, Unq, Typ, Val, ech, nyi, rdc, nyi, scn, nyi, lst, Kst, Out, nyi, nyi, Abs, Img, Cnj, Ang, nyi, nyi, nyi, Tok, Fwh, Las, Exp, Log, Sin, Cos, Prs)
	Functions(64, Asn, Dex, Add, Sub, Mul, Div, Min, Max, Les, Mor, Eql, Mtc, Key, Cat, Cut, Tak, Drp, Cst, Fnd, Atx, Cal, Ech, nyi, Rdc, nyi, Scn, nyi, com, prj, Otu, In, Find, Hyp, Cpx, fdl, Rot, Enc, Dec, nyi, nyi, Bin, Mod, Pow, Lgn)
	Functions(193, tchr, tnms, tvrb, tpct, tvar, tsym, pop)
	Functions(211, Amd, Dmd)

	Functions(220, negi, negf, negz)
	Functions(223, absi, absf, nyi)
	
	Functions(226, addi, addf, addz)
	Functions(229, mini, minf, minz)  //swap sub<->min div<->max
	Functions(232, muli, mulf, mulz)  
	Functions(235, maxi, maxf, maxz)
	Functions(238, subi, subf, subz) 
	Functions(241, divi, divf, divz) 
	
	Functions(244, modi, sqrf, nyi)

	Functions(247, cmi, cmi, cmi, cmF, cmZ, cmC, cmI, cmI, cmF, cmZ, cmL)
	Functions(258, sum, rd0, prd, rd0, min, max)
	Functions(264, mtC, mtC, mtC, mtF, mtF, mtL)
	Functions(270, inC, inI, inI, inF, inZ)
	Functions(275, exp1, log1, sin1, cos1, pow2)
	
	Functions(281, negI, negF, negF)
	Functions(284, absI, absF, nyi)
	Functions(287, ltC, eqC, gtC, ltI, eqI, gtI, ltI, eqI, gtI)
	Functions(296, ltcC, eqcC, gtcC, ltiI, eqiI, gtiI, ltiI, eqiI, gtiI)
	
	Functions(305, sqrF)
	Functions(306, addI, subI, mulI, nyi, minI, maxI)
	Functions(406, addiI, subiI, muliI, nyi, miniI, maxiI)

}

func trap() {
	s := src()
	if srcp == 0 {
		write(Ku(2608)) // 0\n
	} else {
		a := maxi(srcp-30, 0)
		b := mini(nn(s), srcp+30)
		for i := a; i < b; i++ {
			if I8(int32(s)+i) == 10 {
				if i < srcp {
					a = 1 + i
				} else {
					b = i
				}
			}
		}
		Write(0, 0, int32(s)+a, b-a)
		if srcp > a {
			write(Cat(Kc(10), ntake(srcp-a-1, Kc(32))))
		}
	}
	write(Ku(2654)) // ^\n
	panic(srcp)
}

type rdf = func(int32, T, int32) K
type scf = func(K, int32, T, int32) K

func ech(x K) K { return ti(df, int32(l2(x, 0))) } // '
func rdc(x K) K { return ti(df, int32(l2(x, 2))) } // /
func scn(x K) K { return ti(df, int32(l2(x, 4))) } // \

func Ech(f, x K) K {
	var r K
	t := tp(f)
	if isfunc(t) == 0 {
		return Bin(f, Fst(x))
	}
	if nn(x) == 1 {
		x = Fst(x)
	} else {
		return ecn(f, x)
	}
	if tp(x) < 16 {
		trap() //type
	}
	xt := tp(x)
	if xt == Dt {
		r = x0(x)
		return Key(r, Ech(f, l1(r1(x))))
	}
	if xt == Tt {
		x = explode(x)
	}
	xn := nn(x)
	r = mk(Lt, xn)
	rp := int32(r)
	for i := int32(0); i < xn; i++ {
		SetI64(rp, int64(Atx(rx(f), ati(rx(x), i))))
		rp += 8
	}
	dxy(f, x)
	return uf(r)
}
func ecn(f, x K) K {
	if nn(x) == 2 {
		r := x0(x)
		x = r1(x)
		//if r == 0 { project?
		//	return Ech(f, l1(x))
		//}
		if tp(f) == 0 && int32(f) == 13 {
			if tp(r) == Tt && tp(x) == Tt { // T,'T (horcat)
				if nn(r) != nn(x) {
					trap() //length
				}
				f = Cat(x0(r), x0(x))
				return key(f, Cat(r1(r), r1(x)), Tt)
			}
		}
		return ec2(f, r, x)
	}
	return Ech(20, l2(f, Flp(x)))
}
func ec2(f, x, y K) K {
	var r K
	t := dtypes(x, y)
	if t > Lt {
		r = dkeys(x, y)
		return key(r, ec2(f, dvals(x), dvals(y)), t)
	}
	n := conform(x, y)
	switch n {
	case 0: // a-a
		return Cal(f, l2(x, y))
	case 1: // a-v
		n = nn(y)
	default: // v-a, v-v
		n = nn(x)
	}
	r = mk(Lt, n)
	rp := int32(r)
	for i := int32(0); i < n; i++ {
		SetI64(rp, int64(Cal(rx(f), l2(ati(rx(x), i), ati(rx(y), i)))))
		rp += 8
	}
	dx(f)
	dxy(x, y)
	return uf(r)
}

func Rdc(f, x K) K { // x f/y   (x=0):f/y
	t := tp(f)
	if isfunc(t) == 0 {
		if nn(x) == 2 {
			trap() //nyi state machine
		}
		x = Fst(x)
		if t&15 == ct {
			return join(f, x)
		} else {
			return Dec(f, x)
		}
	}
	a := arity(f)
	if a != 2 {
		if a > 2 {
			return rdn(f, x, 0)
		} else {
			return fix(f, Fst(x), 0)
		}
	}

	if nn(x) == 2 {
		return Ecr(f, x)
	}
	x = Fst(x)
	xt := tp(x)
	if xt == Dt {
		x = Val(x)
		xt = tp(x)
	}
	if xt < 16 {
		dx(f)
		return x
	}
	xn := nn(x)
	if t == 0 {
		fp := int32(f)
		if fp > 1 && fp < 8 { // sum,prd,min,max (reduce.go)
			if xt == Tt {
				return Ech(rdc(f), l1(Flp(x)))
			}
			r := Func[256+fp].(rdf)(int32(x), xt, ep(x)) //365
			if r != 0 {
				dx(x)
				return r
			}
		}
		if fp == 13 { // ,/
			if xt < Lt {
				return x
			}
		}
	}
	if xn == 0 {
		dxy(f, x)
		return missing(xt)
	}
	i := int32(1)
	x0 := ati(rx(x), 0)
	for i < xn {
		x0 = cal(rx(f), l2(x0, ati(rx(x), i)))
		i++
	}
	dxy(x, f)
	return x0
}
func rdn(f, x, l K) K { // {x+y*z}/x  {x+y*z}\x
	r := Fst(rx(x))
	x = Flp(ndrop(1, x))
	n := nn(x)
	for i := int32(0); i < n; i++ {
		r = Cal(rx(f), Cat(l1(r), ati(rx(x), i)))
		if l != 0 {
			l = cat1(l, rx(r))
		}
	}
	dxy(f, x)
	if l != 0 {
		dx(r)
		return uf(l)
	}
	return r
}

func Ecr(f, x K) K { //x f/y
	var r K
	y := x1(x)
	x = r0(x)
	yt := tp(y)
	if yt < 16 {
		return cal(f, l2(x, y))
	}
	if yt > Lt {
		t := dtypes(x, y)
		r = dkeys(x, y)
		return key(r, Ecr(f, l2(dvals(x), dvals(y))), t)
	}
	yn := nn(y)
	r = mk(Lt, yn)
	rp := int32(r)
	for i := int32(0); i < yn; i++ {
		SetI64(rp, int64(cal(rx(f), l2(rx(x), ati(rx(y), i)))))
		rp += 8
	}
	dx(f)
	dxy(x, y)
	return uf(r)
}
func fix(f, x, l K) K {
	r := K(0)
	y := rx(x)
	for {
		r = Atx(rx(f), rx(x))
		if match(r, x) != 0 {
			break
		}
		if match(r, y) != 0 {
			break
		}
		dx(x)
		x = r
		if l != 0 {
			l = cat1(l, rx(x))
		}
	}
	dx(f)
	dxy(r, y)
	if l != 0 {
		dx(x)
		return l
	}
	return x
}
func Scn(f, x K) K {
	var r K
	t := tp(f)
	if isfunc(t) == 0 {
		if nn(x) != 1 {
			trap() //rank
		}
		x = Fst(x)
		if t&15 == ct {
			return split(f, x)
		} else {
			return Enc(f, x)
		}
	}
	a := arity(f)
	if a != 2 {
		if a > 2 {
			return rdn(f, x, mk(Lt, 0))
		} else {
			x = rx(Fst(x))
			return fix(f, x, Enl(x))
		}
	}
	if nn(x) == 2 {
		return Ecl(f, x)
	}
	x = Fst(x)
	xt := tp(x)
	if xt < 16 {
		dx(f)
		return x
	}
	xn := nn(x)
	if xn == 0 {
		dx(f)
		return x
	}
	if xt == Dt {
		r = x0(x)
		return Key(r, Scn(f, l1(r1(x))))
	}
	r = mk(Lt, xn)
	rp := int32(r)
	i := int32(1)
	z := ati(rx(x), 0)
	SetI64(rp, int64(rx(z)))
	rp += 8
	for i < xn {
		z = cal(rx(f), l2(z, ati(rx(x), i)))
		SetI64(rp, int64(rx(z)))
		rp += 8
		i++
	}
	dx(z)
	dxy(x, f)
	return uf(r)
}
func Ecl(f, x K) K { // x f\y
	y := x1(x)
	x = r0(x)
	if tp(x) < 16 {
		return cal(f, l2(x, y))
	}
	xn := nn(x)
	r := mk(Lt, xn)
	rp := int32(r)
	for i := int32(0); i < xn; i++ {
		SetI64(rp, int64(cal(rx(f), l2(ati(rx(x), i), rx(y)))))
		rp += 8
	}
	dx(f)
	dxy(x, y)
	return uf(r)
}

func uf(x K) K {
	rt := T(0)
	xn := nn(x)
	xp := int32(x)
	for i := int32(0); i < xn; i++ {
		t := tp(K(I64(xp)))
		if i == 0 {
			rt = t
		} else if t != rt {
			return x
		}
		xp += 8
	}
	if rt == Dt {
		r := Til(x0(x))
		if tp(r) != St {
			dx(r)
			return x
		}
		xp = int32(x)
		for xn > 0 {
			xn--
			if match(r, K(I64(int32(I64(xp))))) == 0 {
				dx(r)
				return x
			}
			xp += 8
		}
		return key(r, Flp(Ech(20, l1(x))), Tt)
	}
	if rt == 0 || rt > zt {
		return x
	}
	r := mk(rt+16, xn)
	for xn > 0 {
		xn--
		r = sti(r, xn, ati(rx(x), xn))
	}
	dx(x)
	return r
}

func minit(a, b int32) {
	p := int32(1 << a)
	for a < b {
		SetI32(4*a, p)
		SetI32(p, 0)
		p *= 2
		a++
	}
	SetI32(128, b)
}
func alloc(n, s int32) int32 {
	size := n * s
	t := bucket(size)
	if int64(n)*int64(s) > 2147483647 /*|| t > 31*/ {
		trap() //grow (oom)
	}
	i := 4 * t
	m := 4 * I32(128)
	for I32(i) == 0 {
		if i >= m {
			m = 4 * grow(i)
		} else {
			i += 4
		}
	}
	a := I32(i)
	SetI32(i, I32(a))
	for j := i - 4; j >= 4*t; j -= 4 {
		u := a + int32(1)<<(j>>2)
		SetI32(u, I32(j))
		SetI32(j, u)
	}
	if a&31 != 0 {
		trap() //memory corruption
	}
	return a
}
func grow(p int32) int32 {
	n := 1 + (p >> 2)                   // required total mem (log2)
	g := (1 << (n - 16)) - Memorysize() // grow by 64k blocks
	if g > 0 {
		if Memorygrow(g) < 0 {
			trap() //grow
		}
	}
	minit(I32(128), n)
	return n
}
func mfree(x, bs int32) {
	if x&31 != 0 {
		trap() //memory corruption
	}
	t := 4 * bs
	SetI32(x, I32(t))
	SetI32(t, x)
}
func bucket(size int32) int32 { return maxi(5, int32(32)-I32clz(15+size)) }
func mk(t T, n int32) K {
	if t < 17 {
		trap() //type
	}
	x := alloc(n, sz(t))
	SetI32(x+12, 1) //rc
	SetI32(x+4, n)
	return ti(t, x+16)
}
func tp(x K) T     { return T(uint64(x) >> 59) }
func nn(x K) int32 { return I32(int32(x) - 12) }
func ep(x K) int32 { return int32(x) + sz(tp(x))*nn(x) }
func sz(t T) int32 {
	if t < 16 {
		return 8
	} else if t < 19 {
		return 1
	} else if t < 21 {
		return 4
	} else if t == Zt {
		return 16
	}
	return 8
}
func rx(x K) K {
	if tp(x) > 4 {
		p := int32(x) - 4
		SetI32(p, 1+I32(p))
	}
	return x
}
func dx(x K) {
	t := tp(x)
	if t < 5 {
		return
	}
	p := int32(x) - 16
	rc := I32(p + 12)
	SetI32(p+12, rc-1)
	if rc == 0 {
		trap() //unref
	}
	if rc == 1 {
		n := nn(x)
		if t&15 > 6 {
			if t == 14 || t == 24 || t == 25 {
				n = 2 // nat | D | T
			} else if t == 12 || t == 13 {
				n = 3 // prj | lam
			}
			p := int32(x)
			e := p + 8*n
			for p < e {
				dx(K(I64(p)))
				p += 8
			}
		}
		mfree(p, bucket(sz(t)*n))
	}
}
func dxy(x, y K) { dx(x); dx(y) }
func rl(x K) { // ref list elements
	e := ep(x)
	p := int32(x)
	for e > p {
		e -= 8
		rx(K(I64(e)))
	}
}

func Cal(x, y K) K {
	y = explode(y)
	if isfunc(tp(x)) != 0 {
		return cal(x, y)
	}
	return atdepth(x, y)
}
func isfunc(t T) int32 { return I32B(t == 0 || uint32(t-10) < 5) }

func cal(f, x K) K {
	r := K(0)
	z := K(0)
	y := K(0)
	t := tp(f)
	fp := int32(f)
	xn := nn(x)
	if t < df {
		switch xn - 1 {
		case 0:
			x = Fst(x)
		case 1:
			r = x1(x)
			x = r0(x)
		default:
			r = x1(x)
			z = x2(x)
			if xn == 4 {
				y = x0(x + 24)
			}
			x = r0(x)
		}
	}
	if t != 0 {
		t -= 9
	}
	switch t {
	case 0: // basic
		switch xn - 1 {
		case 0:
			r = Func[int32(f)].(f1)(x)
		case 1:
			r = Func[fp+64].(f2)(x, r)
		case 2:
			r = Func[fp+192].(f4)(x, r, 1, z)
		case 3:
			r = Func[fp+192].(f4)(x, r, z, y)
		default:
			trap() //rank
			r = 0
		}
		r = r
	case 1: // cf
		switch xn - 1 {
		case 0:
			r = calltrain(f, l1(x))
		case 1:
			r = calltrain(f, l2(x, r))
		default:
			trap() //rank
			r = 0
		}
		r = r
	case 2: // df
		d := x0(f)
		a := 85 + int32(I64(fp+8))
		r = Func[a].(f2)(d, x)
	case 3: // pf
		r = callprj(f, x)
	case 4: // lf
		r = lambda(f, x)
	case 5: // xf
		r = native(f, x)
	default:
		trap() //type
		r = 0
	}
	dx(f)
	return r
}
func calltrain(f, x K) K { return cal(x0(f+8), l1(cal(x0(f), x))) }
func callprj(f, x K) K {
	n := nn(x)
	fn := nn(f)
	if fn != n {
		if n < fn {
			rx(f)
			return prj(f, x)
		}
		trap() //rank
	}
	return Cal(x0(f), stv(x1(f), x2(f), x))
}
func native(f K, x K) K {
	fn := nn(f)
	xn := nn(x)
	if xn != fn {
		if xn < fn {
			rx(f)
			return prj(f, x)
		}
		trap() //rank
	}
	return K(Native(int64(x0(f)), int64(x))) // +/api: KR
}
func lambda(f, x K) K {
	fn := nn(f)
	xn := nn(x)
	if xn < fn {
		rx(f)
		return prj(f, x)
	}
	if xn != fn {
		trap() //rank
	}
	//store vars
	lo := K(I64(int32(f) + 16))
	n := nn(lo)
	a := nn(f)
	z := mk(Zt, n) //use a complex vector to store symbols+values w/o refcounting
	zp := int32(z)
	xp := ep(x)
	vp := I32(8)
	for n > 0 {
		n -= 1
		p := I32(int32(lo) + 4*n)
		SetI32(zp, p)
		p += vp
		SetI64(zp+8, I64(p))
		if n < a { //args
			xp -= 8
			SetI64(p, I64(xp))
		} else { //locals
			SetI64(p, 0)
		}
		zp += 16
	}
	rl(x)
	dx(x)
	x = exec(x0(f)) //execute lambda code
	//restore vars
	zp = int32(z)
	e := ep(z)
	for zp < e {
		p := I32(8) + I32(zp)
		dx(K(I64(p)))
		SetI64(p, I64(zp+8))
		zp += 16
	}
	dx(z)
	return x
}
func com(x, y K) K { return ti(cf, int32(l2(y, x))) } // compose
func prj(f, x K) K { // project
	var r K
	if isfunc(tp(f)) == 0 {
		return atdepth(f, x)
	}
	xn := nn(x)
	xp := int32(x)
	a := mk(It, 0)
	for i := int32(0); i < xn; i++ {
		if I64(xp) == 0 {
			a = cat1(a, Ki(i))
		}
		xp += 8
	}
	ar := arity(f)
	for i := xn; i < ar; i++ {
		a = cat1(a, Ki(i))
		x = cat1(x, 0)
	}
	an := nn(a)
	if tp(f) == pf { // collapse
		r = x1(f)
		y := x2(f)
		f = r0(f)
		x = stv(r, rx(y), x)
		a = Drp(a, y)
	}
	r = l3(f, x, a)
	SetI32(int32(r)-12, an)
	return ti(pf, int32(r))
}
func arity(f K) int32 {
	if tp(f) > df {
		return nn(f)
	}
	return 2
}

func Cat(x, y K) K {
	xt, yt := tp(x), tp(y)
	if xt == Tt && yt == Dt {
		return dcat(x, y)
	}
	if xt&15 == yt&15 {
		if xt < 16 {
			x = Enl(x)
		}
		if yt < 16 {
			return cat1(x, y)
		} else {
			return ucat(x, y)
		}
	} else if xt == Lt && yt < 16 {
		if nn(x) > 0 {
			return cat1(x, y)
		}
	}
	x = uf(Cat(explode(x), explode(y)))
	if nn(x) == 0 {
		dx(x)
		return mk(xt|16, 0)
	}
	return x
}
func Enl(x K) K { return uf(l1(x)) }
func explode(x K) K {
	var r K
	xt := tp(x)
	if xt < 16 || xt == Dt {
		return l1(x)
	} else if xt < Lt {
		xn := nn(x)
		r = mk(Lt, nn(x))
		rp := int32(r)
		for i := int32(0); i < xn; i++ {
			SetI64(rp+8*i, int64(ati(rx(x), i)))
		}
		dx(x)
		return r
	} else if xt == Tt { // Tt
		xn := nn(x)
		k := x0(x)
		x = Flp(r1(x))
		r = mk(Lt, 0)
		for i := int32(0); i < xn; i++ {
			r = cat1(r, Key(rx(k), ati(rx(x), i)))
		}
		dxy(x, k)
		return r
	}
	return x
}
func ucat(x, y K) K { // Bt,Bt .. Tt,Tt
	xt := tp(x)
	if xt > Lt {
		return dcat(x, y)
	}
	xn := nn(x)
	yn := nn(y)
	r := uspc(x, xt, yn)
	s := sz(xt)
	if xt == Lt {
		rl(y)
	}
	Memorycopy(int32(r)+s*xn, int32(y), s*yn)
	dx(y)
	return r
}
func dcat(x, y K) K { // d,d  t,t
	t := tp(x)
	if t == Tt {
		if match(K(I64(int32(x))), K(I64(int32(y)))) == 0 {
			return ucat(explode(x), explode(y))
		}
	}
	r := x0(x)
	x = r1(x)
	q := x0(y)
	y = r1(y)
	if t == Dt {
		r = Cat(r, q)
		return Key(r, Cat(x, y))
	} else {
		dx(q)
		x = Ech(13, l2(x, y))
		return key(r, x, t)
	}
}
func ucat1(x, y, z K) K { return cat1(ucat(x, y), z) }
func cat1(x, y K) K {
	t := tp(x)
	x = uspc(x, t, 1)
	if t == Lt {
		y = l1(rx(y))
		x = ti(Ft, int32(x))
	}
	return ti(t, int32(sti(x, nn(x)-1, y)))
}
func uspc(x K, xt T, ny int32) K {
	r := K(0)
	nx := nn(x)
	s := sz(xt)
	if I32(int32(x)-4) == 1 && bucket(s*nx) == bucket(s*(nx+ny)) {
		r = x
	} else {
		r = mk(xt, nx+ny)
		Memorycopy(int32(r), int32(x), s*nx)
		if xt == Lt {
			rl(x)
		}
		dx(x)
	}
	SetI32(int32(r)-12, nx+ny)
	return r
}

type f1 = func(K) K
type f2 = func(K, K) K
type f3 = func(K, K, K) K
type f4 = func(K, K, K, K) K

func quoted(x K) int32 { return I32B(int32(x) >= 448 && tp(x) == 0) }
func quote(x K) K      { return x + 448 }
func unquote(x K) K    { return x - 448 }

func exec(x K) K {
	var b, c K
	srcp = 0
	a := K(0) // accumulator
	xn := nn(x)
	if xn == 0 {
		dx(x)
		return 0
	}
	p := int32(x)
	e := p + 8*xn
	for p < e {
		u := K(I64(p))
		if tp(u) != 0 {
			push(a)
			a = rx(u)
		} else {
			switch int32(u) >> 6 {
			case 0: //   0..63   monadic
				a = Func[marksrc(u)].(f1)(a)
			case 1: //  64..127  dyadic
				a = Func[marksrc(u)].(f2)(a, pop())
			case 2: // 128       dyadic indirect
				marksrc(a)
				b = pop()
				a = Cal(a, l2(b, pop()))
			case 3: // 192..255  tetradic
				b = pop()
				c = pop()
				a = Func[marksrc(u)].(f4)(a, b, c, pop())
			case 4: // 256       drop
				dx(a)
				a = pop()
			case 5: // 320       jump
				p = p + int32(a)
				a = pop()
			case 6: // 384       jump if not
				u = pop()
				p += int32(a) * I32B(int32(u) == 0)
				dx(u)
				a = pop()
			default: //448..     quoted verb
				push(a)
				a = rx(u - 448)
			}
		}
		p += 8
		continue
	}
	pop() //0
	dx(x)
	return a
}
func marksrc(x K) int32 {
	if p := h48(x); p != 0 {
		srcp = p
	}
	return int32(x)
}
func push(x K) {
	SetI64(sp, int64(x))
	sp += 8
	if sp == 4096 { //512 {
		trap() //stack overflow
	}
}
func pop() K {
	sp -= 8
	if sp < 2048 {
		trap() //stack underflow
	}
	return K(I64(sp))
}
func lst(n K) K {
	r := mk(Lt, int32(n))
	rp := int32(r)
	e := ep(r)
	for rp < e {
		SetI64(rp, int64(pop()))
		rp += 8
	}
	return uf(r)
}
func nul(x K) K { push(x); return 0 }
func lup(x K) K {
	vp := I32(8) + int32(x)
	return x0(K(vp))
}
func Asn(x, y K) K {
	if tp(x) != st {
		trap() //type
	}
	vp := I32(8) + int32(x)
	dx(K(I64(vp)))
	SetI64(vp, int64(rx(y)))
	return y
}
func Amd(x, i, v, y K) K {
	xt := tp(x)
	if xt == st {
		return Asn(x, Amd(lup(x), i, v, y))
	}
	if xt < 16 {
		trap() //type
	}
	if tp(i) == Lt { // @[;;v;]/[x;y;i]
		n := nn(i)
		for j := int32(0); j < n; j++ {
			x = Amd(x, ati(rx(i), j), rx(v), ati(rx(y), j))
		}
		dx(i)
		dxy(v, y)
		return x
	}
	if xt > Lt {
		r := x0(x)
		x = r1(x)
		if xt == Tt && tp(i)&15 == it { // table-assign-rows
			if tp(y) > Lt {
				y = Val(y)
			}
			return key(r, Dmd(x, l2(0, i), v, y), xt)
		}
		r = Unq(Cat(r, rx(i)))
		return key(r, Amd(ntake(nn(r), x), Fnd(rx(r), i), v, y), xt)
	}
	if i == 0 {
		if v == 1 {
			if tp(y) < 16 {
				y = ntake(nn(x), y)
			}
			dx(x)
			return y
		}
		return Cal(v, l2(x, y))
	}
	if tp(v) != 0 || v != 1 {
		y = cal(v, l2(Atx(rx(x), rx(i)), y))
	}
	ti, yt := tp(i), tp(y)
	if xt&15 != yt&15 {
		x, xt = explode(x), Lt
	}
	if ti == it {
		if xt != yt+16 {
			x = explode(x)
		}
		return sti(use(x), int32(i), y)
	}
	if yt < 16 {
		y = ntake(nn(i), y)
		yt = tp(y)
	}
	if xt == Lt {
		y = explode(y)
	}
	return stv(x, i, y)
}
func Dmd(x, i, v, y K) K {
	if tp(x) == st {
		return Asn(x, Dmd(lup(x), i, v, y))
	}
	i = explode(i)
	f := Fst(rx(i))
	if nn(i) == 1 {
		dx(i)
		return Amd(x, f, v, y)
	}
	if f == 0 {
		f = seq(nn(x))
	}
	i = ndrop(1, i)
	if tp(f) > 16 { // matrix-assign
		n := nn(f)
		if nn(i) != 1 {
			trap() //rank
		}
		i = Fst(i)
		if tp(f) == It && tp(x) == Tt {
			t := rx(x0(x))
			return key(t, Dmd(r1(x), l2(Fnd(t, i), f), v, y), Tt)
		}
		if tp(f) != It || tp(x) != Lt {
			trap() // nyi Dt
		}
		x = use(x)
		for j := int32(0); j < n; j++ {
			rj := int32(x) + 8*I32(int32(f)+4*j)
			SetI64(rj, int64(Amd(K(I64(rj)), rx(i), rx(v), ati(rx(y), j))))
		}
		dxy(f, i)
		dxy(v, y)
		return x
	}
	return Amd(x, f, 1, Dmd(Atx(rx(x), f), i, v, y))
}

type f3i = func(int32, int32, int32) int32

func Fnd(x, y K) K { // x?y
	var r K
	xt, yt := tp(x), tp(y)
	if xt < 16 {
		if yt == Tt { // s?T
			r = Drp(rx(x), rx(y))
			return Atx(r, Grp(Atx(y, x)))
		} else {
			return deal(x, y)
		}
	}
	if xt > Lt {
		if xt == Tt {
			trap() //nyi t?..
		}
		r = x0(x)
		return Atx(r, Fnd(r1(x), y))
	} else if xt == yt {
		return Ecr(18+16*K(I32B(yt == Lt)), l2(x, y))
	} else if xt == yt+16 {
		r = Ki(fnd(x, y, yt))
	} else if xt == Lt {
		return fdl(x, y)
	} else if yt == Lt {
		return Ecr(18, l2(x, y))
	} else {
		trap() //type
	}
	dxy(x, y)
	return r
}
func fnd(x, y K, t T) int32 {
	if nn(x) == 0 {
		return nai
	}
	xp := int32(x)
	r := Func[268+t].(f3i)(int32(y), xp, ep(x))
	if r == 0 {
		return nai
	}
	return (r - xp) >> (31 - I32clz(sz(16+t)))
}
func fdl(x, y K) K {
	xp := int32(x)
	dxy(x, y)
	e := ep(x)
	for xp < e {
		if match(K(I64(xp)), y) != 0 {
			return Ki((xp - int32(x)) >> 3)
		}
		xp += 8
	}
	return Ki(nai)
}
func idx(x, a, b int32) int32 {
	for i := a; i < b; i++ {
		if x == I8(i) {
			return i - a
		}
	}
	return -1
}

func Find(x, y K) K { // find[pattern;string] returns all matches (It)
	if tp(x) != Ct || tp(y) != Ct {
		trap()
	}
	xn, yn := nn(x), nn(y)
	if xn*yn == 0 {
		dxy(x, y)
		return mk(It, 0)
	}
	r := mk(It, 0)
	yp := int32(y)
	e := yp + yn + 1 - xn
	for yp < e { // todo rabin-karp / knuth-morris / boyes-moore..
		if findat(int32(x), yp, xn) != 0 {
			r = cat1(r, Ki(yp-int32(y)))
			yp += xn
		} else {
			yp++
		}
		continue
	}
	dxy(x, y)
	return r
}
func findat(xp, yp, n int32) int32 {
	for i := int32(0); i < n; i++ {
		if I8(xp+i) != I8(yp+i) {
			return 0
		}
	}
	return 1
}

func Mtc(x, y K) K {
	dxy(x, y)
	return Ki(match(x, y))
}
func match(x, y K) int32 {
	if x == y {
		return 1
	}
	xt := tp(x)
	if xt != tp(y) {
		return 0
	}
	if xt > 16 {
		n := nn(x)
		if n != nn(y) {
			return 0
		}
		if n == 0 {
			return 1
		}
		xp, yp := int32(x), int32(y)
		if xt < Dt {
			return Func[246+xt].(f3i)(xp, yp, ep(y))
		} else {
			if match(K(I64(xp)), K(I64(yp))) != 0 {
				return match(K(I64(xp+8)), K(I64(yp+8)))
			}
			return 0
		}
	}
	yn := int32(0)
	xp, yp := int32(x), int32(y)
	if xt < ft {
		return I32B(xp == yp)
	}
	switch int32(xt-ft) - 3*I32B(xt > 9) {
	case 0: // ft
		return I32B(0 == cmF(xp, yp))
	case 1: // zt
		return I32B(0 == cmZ(xp, yp))
	case 2: // composition
		yn = 8 * nn(y)
	case 3: // derived
		yn = 16
	case 4: // projection
		yn = 24
	case 5: // lambda
		return match(K(I64(xp+8)), K(I64(yp+8))) // compare strings
	default: // xf
		return I32B(I64(xp) == I64(yp))
	}
	for yn > 0 { // composition, derived, projection
		yn -= 8
		if match(K(I64(xp+yn)), K(I64(yp+yn))) == 0 {
			return 0
		}
	}
	return 1
}
func mtC(xp, yp, e int32) int32 {
	ve := e &^ 7
	for yp < ve {
		if I64(xp) != I64(yp) {
			return 0
		}
		xp += 8
		yp += 8
	}
	for yp < e {
		if I8(xp) != I8(yp) {
			return 0
		}
		xp++
		yp++
	}
	return 1
}
func mtF(xp, yp, e int32) int32 {
	for yp < e {
		if cmF(xp, yp) != 0 {
			return 0
		}
		xp += 8
		yp += 8
		continue
	}
	return 1
}
func mtL(xp, yp, e int32) int32 {
	for yp < e {
		if match(K(I64(xp)), K(I64(yp))) == 0 {
			return 0
		}
		xp += 8
		yp += 8
		continue
	}
	return 1
}
func In(x, y K) K {
	xt, yt := tp(x), tp(y)
	if xt == yt && xt > 16 {
		return Ecl(30, l2(x, y))
	} else if xt+16 != yt {
		trap() //type
	}
	dxy(x, y)
	return Ki(I32B(Func[268+xt].(f3i)(int32(x), int32(y), ep(y)) != 0))
}
func inC(x, yp, e int32) int32 {
	for yp < e { // maybe splat x to int64
		if x == I8(yp) {
			return yp
		}
		yp++
	}
	return 0
}
func inI(x, yp, e int32) int32 {
	for yp < e {
		if x == I32(yp) {
			return yp
		}
		yp += 4
	}
	return 0
}
func inF(xp, yp, e int32) int32 {
	for yp < e {
		if cmF(xp, yp) == 0 {
			return yp
		}
		yp += 8
	}
	return 0
}
func inZ(xp, yp, e int32) int32 {
	for yp < e {
		if cmZ(xp, yp) == 0 {
			return yp
		}
		yp += 16
	}
	return 0
}

func Atx(x, y K) K { // x@y
	var r K
	xt, yt := tp(x), tp(y)
	xp := int32(x)
	if xt < 16 {
		if xt == 0 || xt > 9 {
			return cal(x, l1(y))
		}
		if xt == st {
			if xp == 0 {
				if yt == it { // `123 (quoted verb)
					return K(int32(y))
				}
			}
			xt = ts(x) + 16
			if uint32(xt-18) < 5 { // `c@ .. `z@
				return rtp(xt, y)
			} else {
				return cal(Val(sc(cat1(cs(x), Kc('.')))), l1(y))
			}
		}
	}
	if xt > Lt && yt < Lt {
		r = x0(x)
		x = r1(x)
		if xt == Tt {
			if yt&15 == it {
				return key(r, Ecl(19, l2(x, y)), Dt+T(I32B(yt == It)))
			}
		}
		return Atx(x, Fnd(r, y))
	}
	if yt&15 == ft {
		return Rot(x, y)
	}
	if yt < It {
		y = uptype(y, it)
		yt = tp(y)
	}
	if yt == It {
		return atv(x, y)
	}
	if yt == it {
		return ati(x, int32(y))
	}
	if yt == Lt {
		return Ecr(19, l2(x, y))
	}
	if yt == Dt {
		r = x0(y)
		return Key(r, Atx(x, r1(y)))
	}
	trap() //type f@
	return 0
}
func ati(x K, i int32) K { // x CT..LT
	r := K(0)
	t := tp(x)
	if t < 16 {
		return x
	}
	if t > Lt {
		return Atx(x, Ki(i))
	}
	if i < 0 || i >= nn(x) {
		dx(x)
		return missing(t - 16)
	}
	s := sz(t)
	p := int32(x) + i*s
	switch s >> 2 {
	case 0:
		r = K(uint32(I8(p)))
	case 1:
		r = K(uint32(I32(p)))
	case 2:
		r = K(uint64(I64(p)))
	default:
		dx(x)
		return Kz(F64(p), F64(p+8))
	}
	if t == Ft {
		r = Kf(F64reinterpret_i64(uint64(r)))
	} else if t == Lt {
		r = rx(r)
		dx(x)
		return r
	}
	dx(x)
	return ti(t-16, int32(r))
}
func atv(x, y K) K { // x CT..LT
	t := tp(x)
	if t == Tt {
		return Atx(x, y)
	}
	yn := nn(y)
	if t < 16 {
		dx(y)
		return ntake(yn, x)
	}
	xn := nn(x)
	r := mk(t, yn)
	s := sz(t)
	rp := int32(r)
	xp := int32(x)
	yp := int32(y)
	e := ep(y)

	na := missing(t - 16)
	switch s >> 2 {
	case 0:
		for yp < e {
			xi := I32(yp)
			if uint32(xi) >= uint32(xn) {
				SetI8(rp, int32(na))
			} else {
				SetI8(rp, I8(xp+xi))
			}
			rp++
			yp += 4
		}
	case 1:
		for yp < e {
			xi := I32(yp)
			if uint32(xi) >= uint32(xn) {
				SetI32(rp, int32(na))
			} else {
				SetI32(rp, I32(xp+4*xi))
			}
			rp += 4
			yp += 4
		}
	case 2:
		for yp < e {
			xi := I32(yp)
			if uint32(xi) >= uint32(xn) {
				if t == Lt {
					SetI64(rp, int64(na))
				} else {
					SetI64(rp, I64(int32(na)))
				}
			} else {
				SetI64(rp, I64(xp+8*xi))
			}
			rp += 8
			yp += 4
		}
	default:
		for yp < e {
			xi := I32(yp)
			if uint32(xi) >= uint32(xn) {
				SetI64(rp, I64(int32(na)))
				SetI64(rp+8, I64(int32(na)))
			} else {
				xi *= 16
				SetI64(rp, I64(xp+xi))
				SetI64(rp+8, I64(8+xp+xi))
			}
			rp += 16
			yp += 4
		}
	}
	if t == Lt {
		rl(r)
		r = uf(r)
	}
	dx(na)
	dxy(x, y)
	return r
}
func stv(x, i, y K) K {
	if It != tp(i) {
		trap() //type
	}
	n := nn(i)
	if n == 0 {
		dxy(y, i)
		return x
	}
	if n != nn(y) {
		trap() //length
	}
	x = use(x)
	xt := tp(x)
	xn := nn(x)
	s := sz(xt)
	xp := int32(x)
	yp := int32(y)
	ip := int32(i)
	e := ep(y)
	for j := int32(0); j < n; j++ {
		xi := uint32(I32(ip + 4*j))
		if xi >= uint32(xn) {
			trap() //index
		}
	}
	switch s >> 2 {
	case 0:
		for yp < e {
			SetI8(xp+I32(ip), I8(yp))
			ip += 4
			yp++
		}
	case 1:
		for yp < e {
			SetI32(xp+4*I32(ip), I32(yp))
			ip += 4
			yp += 4
		}
	case 2:
		if xt == Lt {
			rl(y)
			for j := int32(0); j < n; j++ {
				dx(K(I64(xp + 8*I32(ip))))
				ip += 4
			}
			ip = int32(i)
		}
		for yp < e {
			SetI64(xp+8*I32(ip), I64(yp))
			ip += 4
			yp += 8
		}
		if xt == Lt {
			x = uf(x)
		}
	default:
		for yp < e {
			xp = int32(x) + 16*I32(ip)
			SetI64(xp, I64(yp))
			SetI64(xp+8, I64(yp+8))
			ip += 4
			yp += 16
		}
	}
	dxy(i, y)
	return x
}
func sti(x K, i int32, y K) K {
	xt := tp(x)
	if uint32(i) >= uint32(nn(x)) {
		trap() //index
	}
	s := sz(xt)
	xp := int32(x)
	yp := int32(y)
	switch s >> 2 {
	case 0:
		SetI8(xp+i, yp)
	case 1:
		SetI32(xp+4*i, yp)
	case 2:
		xp += 8 * i
		if xt == Lt {
			dx(K(I64(xp)))
			SetI64(xp, int64(rx(y)))
			x = uf(x)
		} else {
			SetI64(xp, I64(yp))
		}
	default:
		xp += 16 * i
		SetI64(xp, I64(yp))
		SetI64(xp+8, I64(yp+8))
	}
	dx(y)
	return x
}

func atdepth(x, y K) K {
	xt := tp(x)
	if xt < 16 {
		trap() //type
	}
	f := Fst(rx(y))
	if f == 0 {
		f = seq(nn(x))
	}
	x = Atx(x, f)
	if nn(y) == 1 {
		dx(y)
		return x
	}
	y = ndrop(1, y)
	if tp(f) > 16 {
		if nn(y) == 1 && xt == Tt {
			return Atx(x, Fst(y))
		}
		return Ecl(20, l2(x, y))
	}
	return atdepth(x, y)
}

const nai int32 = -2147483648 // 0N
var loc, xyz K
var na, inf float64
var pp, pe, sp, srcp, rand_ int32 //parse position/end, stack position, src pointer

//   0....7  key
//   8...15  val
//  16...19  src(int32)
//  20..127  free list
// 128..131  memsize log2
// 132..226  char map (starts at 100)    -+
// 227..252  :+-*%!&|<>=~,^#_$?@.':/:\:   | text
// 253..279  vbcisfzldtcdpl000BCISFZLDT   | section
// 280.....  z.k                         -+
// 2k....4k  stack

func kinit() {
	minit(12, 16) //4k..64k
	sp = 2048
	SetI32(16, int32(mk(Ct, 0))) //SetI64(512, int64(mk(Ct, 0))) //src
	na = F64reinterpret_i64(uint64(0x7FF8000000000001))
	inf = F64reinterpret_i64(uint64(0x7FF0000000000000))
	rand_ = 1592653589
	SetI64(0, int64(mk(Lt, 0)))
	SetI64(8, int64(mk(Lt, 0)))
	xyz = sc(Ku(0))
	xyz = Ech(17, l2(xyz, Ku(8026488))) //`$'"xyz": `x`y`z -> 8 16 24
	zk()
}

type K uint64
type T int32

// typeof(x K): t=x>>59
// isatom:      t<16
// isvector:    t>16
// isflat:      t<22
// basetype:    t&15  0..9
// istagged:    t<5
// haspointers: t>5   (recursive unref)
// elementsize: $[t<19;1;t<21;4;8]
const ( //base t&15          bytes  atom  vector
	ct T = 2 // char    1      2     18
	it T = 3 // int     4      3     19
	st T = 4 // symbol  4      4     20
	ft T = 5 // float   8      5     21
	zt T = 6 // complex(8)     6     22

	cf T = 10 // comp   (8)    10
	df T = 11 // derived(8)    11
	pf T = 12 // proj   (8)    12
	lf T = 13 // lambda (8)    13
	xf T = 14 // native (8)    14
	Ct T = 18
	It T = 19
	St T = 20
	Ft T = 21
	Zt T = 22
	Lt T = 23 // list
	Dt T = 24 // dict
	Tt T = 25 // table
)

// func t=0
// basic x < 64 (triadic/tetradic)
// composition .. f2 f1 f0
// derived     func    symb
// projection  func    arglist  emptylist
// lambda      code    string	locals
// native      ptr(Ct) string

// ptr: int32(x)
//  p-12    p-4 p
// [length][rc][data]

func ti(t T, i int32) K { return K(t)<<59 | K(uint32(i)) }
func Kc(x int32) K      { return ti(ct, x) }
func Ki(x int32) K      { return ti(it, x) }
func Ks(x int32) K      { return ti(st, x) }
func Kf(x float64) K {
	r := mk(Ft, 1)
	SetF64(int32(r), x)
	return ti(ft, int32(r))
}
func Kz(x, y float64) K {
	r := mk(Zt, 1)
	rp := int32(r)
	SetF64(rp, x)
	SetF64(rp+8, y)
	return ti(zt, rp)
}
func l1(x K) K {
	r := mk(Lt, 1)
	SetI64(int32(r), int64(x))
	return r
}
func l2(x, y K) K {
	r := mk(Lt, 2)
	rp := int32(r)
	SetI64(rp, int64(x))
	SetI64(8+rp, int64(y))
	return r
}
func l3(x, y, z K) K { return cat1(l2(x, y), z) }
func r0(x K) K       { r := x0(x); dx(x); return r }
func r1(x K) K       { r := x1(x); dx(x); return r }
func x0(x K) K       { return rx(K(I64(int32(x)))) }
func x1(x K) K       { return x0(x + 8) }
func x2(x K) K       { return x0(x + 16) }
func Ku(x uint64) K { // Ct
	r := mk(Ct, 8)
	p := int32(r)
	SetI64(p, int64(x))
	SetI32(p-12, idx(0, p, p+8)) //assume <8
	return r
}

/* encode bytes for Ku(..) with: https://play.golang.org/p/4ethx6OEVCR
func enc(x []byte) uint64 {
	r := uint32(0)
	var o uint64 = 1
	for _, b := range x {
		r += o * uint64(b)
		o <<= 8
	}
	return r
}
*/

func kx(u int32, x K) K { return cal(Val(Ks(u)), l1(x)) } //call k func from z.k
func sc(c K) K { //symbol from chars
	s := K(I64(0))
	sp := int32(s)
	sn := nn(s)
	for i := int32(0); i < sn; i++ {
		if match(c, K(I64(sp))) != 0 {
			dx(c)
			return ti(st, sp-int32(s))
		}
		sp += 8
	}
	SetI64(0, int64(cat1(s, c)))
	SetI64(8, int64(cat1(K(I64(8)), 0)))
	return ti(st, 8*sn)
}
func cs(x K) K { return x0(K(I32(0)) + x) } //chars from symbol
func missing(t T) K {
	switch t - 2 {
	case 0: // ct
		return Kc(32)
	case 1: // it
		return Ki(nai)
	case 2: // st
		return Ks(0)
	case 3: // ft
		return Kf(na)
	case 4: // zt
		return Kz(na, na)
	default: // lt
		return mk(Ct, 0) //Kb(0)
	}
}

// softfloat implementation of cosin_ atan2 log exp pow frexp is 2464 b

const pi float64 = 3.141592653589793
const maxfloat float64 = 1.797693134862315708145274237317043567981e+308

func hypot(p, q float64) float64 {
	p, q = F64abs(p), F64abs(q)
	if p < q {
		t := p
		p = q
		q = t
	}
	if p == 0.0 {
		return 0.0
	}
	q = q / p
	return p * F64sqrt(1+q*q)
}
func cosin(deg float64, rp int32) {
	c, s := 0.0, 0.0
	if deg == 0 {
		c = 1.0
	} else if deg == 90 {
		s = 1.0
	} else if deg == 180 {
		c = -1.0
	} else if deg == 270 {
		s = -1.0
	} else {
		cosin_(deg*0.017453292519943295, rp, 0)
		return
	}
	SetF64(rp, c)
	SetF64(rp+8, s)
}
func ang2(y, x float64) float64 {
	if y == 0 {
		if x < 0 {
			return 180.0
		}
		return 0.
	}
	if x == 0 {
		if y < 0 {
			return 270.0
		}
		return 90.0
	}
	deg := 57.29577951308232 * atan2(y, x)
	if deg < 0 {
		deg += 360.0
	}
	return deg
}
func exp1(xp, yp, rp int32) { SetF64(rp, exp(F64(xp))) }
func log1(xp, yp, rp int32) { SetF64(rp, log(F64(xp))) }
func pow2(xp, yp, rp int32) { SetF64(rp, pow(F64(xp), F64(yp))) }
func sin1(xp, yp, rp int32) { cosin_(F64(xp), rp, 1) }
func cos1(xp, yp, rp int32) { cosin_(F64(xp), rp, 2) }
func cosin_(x float64, rp int32, csonly int32) {
	c, s, ss, cs := 0.0, 0.0, int32(0), int32(0)
	if x < 0 {
		x = -x
		ss = 1
	}
	j := int64(x * 1.2732395447351628) // *4/pi
	y := float64(j)
	if j&1 == 1 {
		j++
		y++
	}
	j &= 7
	z := ((x - y*7.85398125648498535156e-1) - y*3.77489470793079817668e-8) - y*2.69515142907905952645e-15
	if j > 3 {
		j -= 4
		//ss, cs = !ss, !cs
		ss, cs = 1-ss, 1-cs
	}
	if j > 1 {
		cs = 1 - cs
	}
	zz := z * z
	c = 1.0 - 0.5*zz + zz*zz*((((((-1.13585365213876817300e-11*zz)+2.08757008419747316778e-9)*zz+-2.75573141792967388112e-7)*zz+2.48015872888517045348e-5)*zz+-1.38888888888730564116e-3)*zz+4.16666666666665929218e-2)
	s = z + z*zz*((((((1.58962301576546568060e-10*zz)+-2.50507477628578072866e-8)*zz+2.75573136213857245213e-6)*zz+-1.98412698295895385996e-4)*zz+8.33333333332211858878e-3)*zz+-1.66666666666666307295e-1)
	if j == 1 || j == 2 {
		x = c
		c = s
		s = x
	}
	if cs != 0 {
		c = -c
	}
	if ss != 0 {
		s = -s
	}
	SetF64(rp, c)
	if csonly == 0 {
		SetF64(rp+8, s)
	} else if csonly == 1 {
		SetF64(rp, s)
	}
}
func atan2(y, x float64) float64 {
	// todo nan/inf
	q := atan(y / x)
	if x < 0 {
		if q <= 0 {
			return q + pi
		}
		return q - pi
	}
	return q
}
func atan(x float64) float64 {
	if x > 0 {
		return satan(x)
	} else {
		return -satan(-x)
	}
}
func satan(x float64) float64 {
	if x <= 0.66 {
		return xatan(x)
	}
	if x > 2.41421356237309504880 {
		return 1.5707963267948966 - xatan(1.0/x) + 6.123233995736765886130e-17
	}
	return 0.7853981633974483 + xatan((x-1)/(x+1)) + 0.5*6.123233995736765886130e-17
}
func xatan(x float64) float64 {
	z := x * x
	z = z * ((((-8.750608600031904122785e-01*z+-1.615753718733365076637e+01)*z+-7.500855792314704667340e+01)*z+-1.228866684490136173410e+02)*z + -6.485021904942025371773e+01) / (((((z+2.485846490142306297962e+01)*z+1.650270098316988542046e+02)*z+4.328810604912902668951e+02)*z+4.853903996359136964868e+02)*z + 1.945506571482613964425e+02)
	z = x*z + x
	return z
}
func exp(x float64) float64 {
	var k int64
	if x != x {
		return x
	}
	if x > 7.09782712893383973096e+02 {
		return inf
	}
	if x < -7.45133219101941108420e+02 {
		return 0.0
	}
	if -3.725290298461914e-09 < x && x < 3.725290298461914e-09 {
		return 1.0 + x
	}
	if x < 0 {
		k = int64(1.44269504088896338700*x - 0.5)
	} else {
		k = int64(1.44269504088896338700*x + 0.5)
	}
	hi := x - float64(k)*6.93147180369123816490e-01
	lo := float64(k) * 1.90821492927058770002e-10
	return expmulti(hi, lo, k)
}
func expmulti(hi, lo float64, k int64) float64 {
	r := hi - lo
	t := r * r
	c := r - t*(1.66666666666666657415e-01+t*(-2.77777777770155933842e-03+t*(6.61375632143793436117e-05+t*(-1.65339022054652515390e-06+t*4.13813679705723846039e-08))))
	y := 1 - ((lo - (r*c)/(2-c)) - hi)
	return ldexp(y, k)
}
func ldexp(frac float64, exp int64) float64 {
	if frac == 0 || frac > maxfloat || frac < -maxfloat || (frac != frac) {
		return frac
	}
	nf := normalize(frac)
	if nf != frac {
		exp -= 52
		frac = nf
	}
	x := uint64(I64reinterpret_f64(frac))
	exp += int64(x>>52)&2047 - 1023
	if exp < int64(-1075) {
		return F64copysign(0, frac)
	}
	if exp > int64(1023) {
		if frac < 0 {
			return -inf
		}
		return inf
	}
	m := 1.0
	if exp < int64(-1022) {
		exp += 53
		m = 1.1102230246251565e-16
	}
	x &^= 9218868437227405312
	x |= uint64(exp+1023) << 52
	return m * F64reinterpret_i64(uint64(x))
}
func frexp1(f float64) int32 {
	if f == 0.0 {
		return 0
	}
	if f < -maxfloat || f > maxfloat || (f != f) {
		return 0
	}
	return 1
}
func frexp2(f float64) float64 {
	f = normalize(f)
	x := I64reinterpret_f64(f)
	x &^= 9218868437227405312
	x |= 4602678819172646912
	return F64reinterpret_i64(x)
}
func frexp3(f float64) int64 {
	exp := int64(0)
	nf := normalize(f)
	if nf != f {
		exp = int64(-52)
		f = nf
	}
	x := I64reinterpret_f64(f)
	return exp + int64((x>>52)&2047) - 1022
}
func normalize(x float64) float64 {
	if F64abs(x) < 2.2250738585072014e-308 {
		return x * 4.503599627370496e+15
	}
	return x
}
func log(x float64) float64 {
	if (x != x) || x > maxfloat {
		return x
	}
	if x < 0 {
		return na
	}
	if x == 0 {
		return -inf
	}
	f1 := x
	ki := int64(0)
	if frexp1(x) != 0 {
		f1 = frexp2(x)
		ki = frexp3(x)
	}
	if f1 < 0.7071067811865476 {
		f1 *= 2
		ki--
	}
	f := f1 - 1
	k := float64(ki)
	s := f / (2 + f)
	s2 := s * s
	s4 := s2 * s2
	t1 := s2 * (6.666666666666735130e-01 + s4*(2.857142874366239149e-01+s4*(1.818357216161805012e-01+s4*1.479819860511658591e-01)))
	t2 := s4 * (3.999999999940941908e-01 + s4*(2.222219843214978396e-01+s4*1.531383769920937332e-01))
	R := t1 + t2
	hfsq := 0.5 * f * f
	return k*6.93147180369123816490e-01 - ((hfsq - (s*(hfsq+R) + k*1.90821492927058770002e-10)) - f)
}
func modabsfi(f float64) float64 {
	if f < 1.0 { // simplified for f > 0
		return 0
	}
	x := I64reinterpret_f64(f)
	e := (x>>52)&2047 - 1023
	if e < 52 {
		x &^= uint64(1)<<(52-e) - uint64(1)
	}
	return F64reinterpret_i64(x)
}
func pow(x, y float64) float64 {
	if y == 0.0 || x == 1.0 {
		return 1.0
	}
	if y == 1.0 {
		return x
	}
	if (x != x) || (y != y) || y > maxfloat || y < -maxfloat { // simplified
		return na
	}
	if x == 0 { // simplified
		if y < 0 {
			return inf
		} else {
			return 0.0
		}
	}
	if y == 0.5 {
		return F64sqrt(x)
	}
	if y == -0.5 {
		return 1.0 / F64sqrt(x)
	}

	yf := F64abs(y)
	yi := modabsfi(yf)
	yf -= yi
	if yf != 0.0 && x < 0.0 {
		return na
	}
	if yi >= 9.223372036854776e+18 {
		if x == -1.0 {
			return 1.0
		} else if (F64abs(x) < 1.0) == (y > 0.0) {
			return 0.0
		} else {
			return inf
		}
	}
	a1 := 1.0
	ae := int64(0)
	if yf != 0 {
		if yf > 0.5 {
			yf -= 1.0
			yi += 1.0
		}
		a1 = exp(yf * log(x))
	}
	x1 := x
	xe := int64(0)
	if frexp1(x) != 0 {
		x1 = frexp2(x)
		xe = frexp3(x)
	}
	for i := int64(yi); i != 0; i >>= int64(1) {
		if xe < int64(-4096) || 4096 < xe {
			ae += xe
			break
		}
		if i&1 == 1 {
			a1 *= x1
			ae += xe
		}
		x1 *= x1
		xe <<= int64(1)
		if x1 < 0.5 {
			x1 += x1
			xe--
		}
	}
	if y < 0.0 {
		a1 = 1.0 / a1
		ae = -ae
	}
	return ldexp(a1, ae)
}
func ipow(x K, y int32) K {
	if tp(x) == It {
		return Ecr(42, l2(Ki(y), x))
	} else {
		return Ki(iipow(int32(x), y))
	}
}
func iipow(x, y int32) int32 {
	r := int32(1)
	for {
		if y&1 == 1 {
			r *= x
		}
		y >>= 1
		if y == 0 {
			break
		}
		x *= x
	}
	return r
}

var ps int32

func Prs(x K) K { return parse(Tok(x)) } // `p"src"  `p(token list)
func parse(x K) K {
	if tp(x) != Lt {
		trap() //type
	}
	pp = int32(x)
	n := 8 * nn(x)
	pe = n + pp
	r := es()
	if pp != pe {
		trap() //parse
	}
	mfree(int32(x)-16, bucket(n)) //free non-recursive
	return r
}
func es() K {
	r := mk(Lt, 0)
	for {
		n := next()
		if n == 0 {
			break
		}
		if n == 59 {
			continue
		}
		pp -= 8
		x := e(t()) &^ 1
		if x == 0 {
			break
		}
		if nn(r) != 0 {
			r = cat1(r, 256) // drop
		}
		r = Cat(r, x)
	}
	return r
}
func e(x K) K { // Lt
	var r K
	xv := x & 1
	x &^= 1
	if x == 0 {
		return 0
	}
	xs := ps
	y := t()
	yv := y & 1
	y &^= 1
	if y == 0 {
		return x + xv
	}
	if yv != 0 && xv == 0 {
		r = e(t())
		ev := r & 1
		r &^= 1
		a := pasn(x, y, r)
		if a != 0 {
			return a
		}
		if r == 0 || ev == 1 { // 1+ (projection)
			x = ucat1(cat1(ucat1(l1(0), x, Ki(2)), 27), y, 92)
			if ev == 1 { // 1+-
				return ucat1(r, x, 91) + 1
			}
			return x + 1
		}
		return dyadic(ucat(r, x), y) // dyadic
	}
	r = e(rx(y) + yv)
	ev := r & 1
	r &^= 1
	dx(y)
	if xv == 0 {
		return ucat1(r, x, 83|K(xs)<<32) // juxtaposition
	} else if (r == y && xv+yv == 2) || ev == 1 {
		return ucat1(r, x, 91) + 1 // composition
	}
	return idiom(monadic(ucat(r, x))) // monadic
}
func t() K { // Lt
	r := next()
	if r == 0 {
		return 0
	}
	rt := tp(r)
	if rt == 0 && int32(r) < 127 {
		if is(int32(r), 32) != 0 {
			pp -= 8
			return 0
		}
	}
	verb := K(0)
	if r == K('(') {
		r = rlist(plist(41)&^1, 0)
	} else if r == K('{') {
		r = plam(ps)
	} else if r == K('[') {
		r = es()
		if next() != K(']') {
			trap() //parse
		}
		return r
	} else if rt == st {
		r = l2(r, 20|(K(ps)<<32)) // .`x (lookup)
	} else {
		if rt == 0 {
			r, verb = quote(r)|K(ps)<<32, 1
		} else if rt == St {
			if nn(r) == 1 {
				r = Fst(r)
			}
		}
		r = l1(r)
	}
f:
	for {
		n := next()
		if n == 0 {
			break f
		}
		ks := K(ps) << 32
		a := int32(n)
		if tp(n) == 0 && a > 20 && a < 27 { // +/
			r, verb = cat1(r, n), 1
		} else if n == 91 { // [
			verb = 0
			n = plist(93)
			p := K(84) + 8*(n&1) // 92(project) or call(84)
			n &^= 1
			s := pspec(r, n)
			if s != 0 {
				return s
			}
			if nn(n) == 1 {
				r = ucat1(Fst(n), r, 83|ks)
			} else {
				r = cat1(Cat(rlist(n, 2), r), p|ks)
			}
		} else {
			if tp(n) == 4 && rt < 6 {
				if is(I8(I32(16)+ps-2), 4) != 0 {
					r = cat1(l3(n, 20, Fst(r)), 68|ks)
					continue f
				}
			}
			pp -= 8
			break f // within else-if
		}
	}
	return r + verb
}
func pasn(x, y, r K) K {
	l := K(I64(int32(y)))
	v := int32(l)
	sp := h48(l)
	if nn(y) == 1 && tp(l) == 0 && v == 449 || (v > 544 && v < 565) {
		dx(y)
		xn := nn(x)
		if xn > 2 { // indexed amd/dmd
			if v > 544 { // indexed-modified
				l -= 96
			}
			s := ati(rx(x), xn-3)
			lp := 0xff000000ffffffff & lastp(x)
			// (+;.i.;`x;.;@) -> x:@[x;.i.;+;rhs] which is (+;.i.;`x;.;211 or 212)
			// lp+128 is @[amd..] or .[dmd..]
			if lp == 92 {
				lp = 84 // x[i;]:.. no projection
			}
			x = cat1(ucat1(l1(l), ldrop(-2, x), 20), (K(sp)<<32)|(lp+128))
			y = l2(s, 448) // s:..
		} else if v == 449 || v == 545 {
			if xn == 1 { // `x: is (,`x) but type Lt replace with `"x." to use with `x@
				x = sc(cat1(cs(Fst(Fst(x))), Kc(46))) // `x: -> `"x."
			} else {
				x = Fst(x) // (`x;.)
			}
			if loc != 0 && v == 449 {
				loc = Cat(loc, rx(x))
			}
			x = l1(x)
			y = l1(448) // asn
		} else { // modified
			y = cat1(l2(unquote(l-32), Fst(rx(x))), 448)
		}
		return dyadic(ucat(r, x), y)
	}
	return 0
}
func plam(s0 int32) K {
	r := K(0)
	slo := loc
	loc = 0
	ar := int32(-1)
	n := next()
	if n == 91 { // argnames
		n := plist(93) &^ 1
		ln := nn(n)
		loc = Ech(4, l1(n)) // [a]->,(`a;.)  [a;b]->((`a;.);(`b;.))
		if ln > 0 && tp(loc) != St {
			trap() //parse
		}
		ar = nn(loc)
		if ar == 0 {
			dx(loc)
			loc = mk(St, 0)
		}
	} else {
		pp -= 8
		loc = mk(St, 0)
	}
	//c := cat1(es(), 30) //rst
	c := es()
	n = next()
	if n != 125 {
		trap() //parse
	}
	cn := nn(c)
	cp := int32(c)
	if ar < 0 {
		ar = 0
		for cn > 0 {
			cn--
			r = K(I64(cp))
			if tp(r) == 0 && int32(r) == 20 {
				r = K(I64(cp - 8))
				y := int32(r) >> 3
				if tp(r) == st && y > 0 && y < 4 {
					ar = maxi(ar, y)
				}
			}
			cp += 8
		}
		loc = Cat(ntake(ar, rx(xyz)), loc)
	}
	i := Add(seq(1+ps-s0), Ki(s0-1))
	s := atv(rx(src()), i)
	r = l3(c, s, Unq(loc))
	loc = slo
	cp = int32(r)
	SetI32(cp-12, ar)
	return l1(ti(lf, cp) | K(s0)<<32)
}
func pspec(r, n K) K {
	ln := nn(n)
	v := K(I64(int32(r)))
	if nn(r) == 1 && ln > 2 { // $[..] cond
		if tp(v) == 0 && int32(v) == 465 {
			dx(r)
			return cond(n, ln)
		}
	}
	if nn(r) == 2 && ln > 1 && int32(v) == 64 { // while[..]
		dx(r)
		return whl(n, ln-1)
	}
	return 0
}
func whl(x K, xn int32) K {
	r := cat1(Fst(rx(x)), 0)
	p := nn(r) - 1
	r = ucat(r, l2(384, 256)) //jif drop
	xp := int32(x)
	sum := int32(2)
	for i := int32(0); i < xn; i++ {
		if i != 0 {
			r = cat1(r, 256)
		}
		xp += 8
		y := x0(K(xp))
		sum += 1 + nn(y)
		r = ucat(r, y)
	}
	r = cat1(cat1(r, Ki(-8*(2+nn(r)))), 320) // jmp back
	SetI64(int32(r)+8*p, int64(Ki(8*sum)))   // jif
	dx(x)
	return ucat(l1(0), r) // null for empty while
}
func cond(x K, xn int32) K {
	nxt := int32(0)
	sum := int32(0)
	xp := int32(x) + 8*xn
	state := int32(1)
	for xp != int32(x) {
		xp -= 8
		r := K(I64(xp))
		if sum > 0 {
			state = 1 - state
			if state != 0 {
				r = cat1(cat1(r, Ki(nxt)), 384) // jif
			} else {
				r = cat1(cat1(r, Ki(sum)), 320) // j
			}
			SetI64(xp, int64(r))
		}
		nxt = 8 * nn(r)
		sum += nxt
	}
	return Rdc(13, l1(x))
}
func plist(c K) K {
	p := K(0)
	r := mk(Lt, 0)
	for {
		b := next()
		if b == 0 || b == c {
			break
		}
		if nn(r) == 0 {
			pp -= 8
		}
		x := e(t()) &^ 1
		if x == 0 {
			p = 1
		}
		r = cat1(r, x)
	}
	return r + p
}
func rlist(x, p K) K {
	n := nn(x)
	if n == 0 {
		return l1(x)
	}
	if n == 1 {
		return Fst(x)
	}
	if p != 2 {
		p = clist(x)
		if p != 0 {
			return l1(p)
		}
	}
	return cat1(cat1(Rdc(13, l1(Rev(x))), Ki(n)), 27)
}
func clist(x K) K { //constant-fold list
	p := int32(x)
	e := ep(x)
	for p < e {
		xi := K(I64(p))
		t := tp(xi)
		if t != Lt {
			return 0
		}
		if nn(xi) != 1 {
			return 0
		}
		if tp(K(I64(int32(xi)))) == 0 {
			return 0
		}
		p += 8
	}
	return uf(Rdc(13, l1(x)))
}

func next() K {
	if pp == pe {
		return 0
	}
	r := K(I64(pp))
	ps = h48(r)
	pp += 8
	return r & 0xff000000ffffffff
}
func lastp(x K) K   { return K(I64(ep(x) - 8)) }
func h48(x K) int32 { return 0xffffff & int32(x>>32) }
func dyadic(x, y K) K {
	l := lastp(y)
	if quoted(l) != 0 {
		return ucat1(x, ldrop(-1, y), 64+unquote(l))
	}
	return ucat1(x, y, 128)
}
func monadic(x K) K {
	l := lastp(x)
	if quoted(l) != 0 {
		x = ldrop(-1, x)
		if int32(l) == 449 { // :x return lambda
			return cat1(cat1(x, Ki(1048576)), 320) //identity+long jump
		} else {
			return cat1(x, unquote(l))
		}
	}
	return cat1(x, 83) // dyadic-@
}
func ldrop(n int32, x K) K { return explode(ndrop(n, x)) }
func svrb(p int32) int32 {
	x := K(I64(p))
	return I32B(int32(x) < 64 && tp(x) == 0) * int32(x)
}
func idiom(x K) K {
	l := int32(x) + 8*(nn(x)-2)
	i := svrb(l) + svrb(l+8)<<6
	if i == 262 || i == 263 { // *& 6 4 -> 40
		i = 34 // 6->40(Fwh) 7->41(Las)
	} else {
		return x
	}
	SetI64(l, I64(l)+int64(i))
	return ndrop(-1, x)
}

func rnd() int32 {
	r := rand_
	r ^= (r << 13)
	r ^= (r >> 17)
	r ^= (r << 5)
	rand_ = r
	return r
}
func roll(x K) K { // ?x (atom) ?n(uniform 0..1) ?-n(normal) ?z(binormal)
	xt := tp(x)
	xp := int32(x)
	if xt == it {
		if xp > 0 {
			return kx(72, x) // .rf uniform
		} else {
			r := kx(80, Ki((1+-xp)/2))
			SetI32(int32(r)-12, -xp)
			return ti(Ft, int32(r)) // normal
		}
	}
	if xt == zt {
		dx(x)
		return kx(80, Ki(int32(F64floor(F64(xp))))) //.rz binormal
	}
	trap() //type
	return 0
}
func deal(x, y K) K { // x?y (x atom) n?n(with replacement) -n?n(without) n?L (-#L)?L shuffle
	yt := tp(y)
	if yt > 16 {
		return In(x, y)
	}
	if tp(x) != it {
		trap() //type
	}
	xp := int32(x)
	if yt == ct {
		return Add(Kc(97), Flr(deal(x, Ki(int32(y)-96))))
	}
	if yt == st {
		return Ech(17, l2(Ks(0), deal(x, Fst(cs(y))))) // `$'x?*$y
	}
	if yt != it {
		trap() //type
	}
	yp := int32(y)
	if xp > 0 {
		return randI(yp, xp) // n?m
	}
	// todo n<<m
	return ntake(-xp, shuffle(seq(yp), -xp)) //-n?m (no duplicates)
}
func randi(n int32) int32 {
	v := uint32(rnd())
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = uint32(rnd())
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
		}
	}
	return int32(prod >> 32)
}
func randI(i, n int32) K {
	r := mk(It, n)
	rp := int32(r)
	e := ep(r)
	if i == 0 {
		for rp < e {
			SetI32(rp, rnd())
			rp += 4
		}
	} else {
		for rp < e {
			SetI32(rp, randi(i))
			rp += 4
		}
	}
	return r
}
func shuffle(r K, m int32) K { // I, inplace
	rp := int32(r)
	n := nn(r)
	m = mini(n-1, m)
	for i := int32(0); i < m; i++ {
		j := rp + 4*randi(n-i)
		t := I32(rp)
		SetI32(rp, I32(j))
		SetI32(j, t)
		rp += 4
	}
	return r
}
func rd0(yp int32, t T, n int32) K { return 0 }
func min(yp int32, t T, e int32) K { // &/x
	var xp int32
	switch t - 18 {
	case 0: // Ct
		xp = 127
		for yp < e {
			xp = mini(xp, I8(yp))
			yp++
		}
		return Kc(xp)
	case 1: // It
		xp = 2147483647
		for yp < e {
			xp = mini(xp, I32(yp))
			yp += 4
		}
		return Ki(xp)
	case 2: // St
		xp = (nn(K(I64(8))) << 3) - 8
		for yp < e {
			xp = mini(xp, I32(yp))
			yp += 4
		}
		return Ks(xp)
	case 3: // Ft
		f := inf
		for yp < e {
			f = F64min(f, F64(yp))
			yp += 8
		}
		return Kf(f)
	default:
		return 0
	}
}
func max(yp int32, t T, e int32) K { // |/x
	var xp int32
	switch t - 18 {
	case 0: // Ct
		xp = -128
		for yp < e {
			xp = maxi(xp, I8(yp))
			yp++
		}
		return Kc(xp)
	case 1: // It
		xp = nai
		for yp < e {
			xp = maxi(xp, I32(yp))
			yp += 4
		}
		return Ki(xp)
	case 2: // St
		xp = 0
		for yp < e {
			xp = maxi(xp, I32(yp))
			yp += 4
		}
		return Ks(xp)
	case 3: // Ft
		f := -inf
		for yp < e {
			f = F64max(f, F64(yp))
			yp += 8
		}
		return Kf(f)
	default:
		return 0
	}
}
func sum(yp int32, t T, e int32) K { // +/x
	xp := int32(0)
	switch t - 18 {
	case 0: // Ct
		for yp < e {
			xp += I8(yp)
			yp++
		}
		return Kc(xp)
	case 1: // It
		return Ki(xp + sumi(yp, e))
	case 2: // St
		return 0
	case 3: // Ft
		f := 0.0
		return Kf(f + sumf(yp, e, 8))
	case 4: // Zt
		re := 0.0
		im := 0.0
		return Kz(re+sumf(yp, e, 16), im+sumf(yp+8, e, 16))
	default:
		return 0
	}
}
func sumi(xp, e int32) int32 {
	r := int32(0)
	for xp < e {
		r += I32(xp)
		xp += 4
	}
	return r
}
func sumf(xp, e, s int32) float64 {
	r := 0.0
	for xp < e {
		r += F64(xp)
		xp += s
	}
	return r
}
func prd(yp int32, t T, e int32) K { // */x
	xp := int32(1)
	switch t - 18 {
	case 0: // Ct
		for yp < e {
			xp *= I8(yp)
			yp++
		}
		return Kc(xp)
	case 1: // It
		for yp < e {
			xp *= I32(yp)
			yp += 4
		}
		return Ki(xp)
	case 2: // St
		return 0
	case 3: // Ft
		f := 1.0
		for yp < e {
			f *= F64(yp)
			yp += 8
		}
		return Kf(f)
	default:
		return 0
	}
}
func Srt(x K) K { // ^x
	var r K
	xt := tp(x)
	if xt < 16 {
		trap() //type
	}
	if xt == Dt {
		r = x0(x)
		x = r1(x)
		i := rx(Asc(rx(x)))
		return Key(atv(r, i), atv(x, i))
	}
	if nn(x) < 2 {
		return x
	}
	return atv(x, Asc(rx(x)))
}
func Asc(x K) K { // <x  <`file
	if tp(x) == st {
		return readfile(cs(x))
	}
	return grade(x, 1)
}
func Dsc(x K) K { return grade(x, -1) } //254 // >x
func grade(x K, f int32) K { // <x >x
	var r K
	xt := tp(x)
	if xt < 16 {
		trap() //type
	}
	if xt == Dt {
		r = x0(x)
		return Atx(r, grade(r1(x), f))
	}
	n := nn(x)
	if xt == Tt {
		return cal(lup(Ks(88)), l2(x, Ki(I32B(f == -1)))) //gdt ngn:{(!#x){x@<y x}/|.+x}
	}
	if n < 2 {
		dx(x)
		return seq(n)
	}
	r = seq(n)
	rp := int32(r)
	xp := int32(x)
	w := mk(It, n)
	wp := int32(w)
	Memorycopy(wp, rp, 4*n)
	msrt(wp, rp, 0, n, xp, int32(xt), f)
	dxy(w, x)
	return r
}

func msrt(x, r, a, b, p, t, f int32) {
	if b-a < 2 {
		return
	}
	c := (a + b) >> 1
	msrt(r, x, a, c, p, t, f)
	msrt(r, x, c, b, p, t, f)
	mrge(x, r, 4*a, 4*b, 4*c, p, t, f)
}
func mrge(x, r, a, b, c, p, t, f int32) {
	var q int32
	i, j := a, c
	s := sz(T(t))
	for k := a; k < b; k += 4 {
		if i < c && j < b {
			q = I32B(f == Func[234+t].(f2i)(p+s*I32(x+i), p+s*I32(x+j)))
		} else {
			q = 0
		}
		if i >= c || q != 0 {
			SetI32(r+k, I32(x+j))
			j += 4
		} else {
			SetI32(r+k, I32(x+i))
			i += 4
		}
	}
}
func cmL(xp, yp int32) int32 { // compare lists lexically
	var r int32
	x, y := K(I64(xp)), K(I64(yp))
	xt, yt := tp(x), tp(y)
	if xt != yt {
		return I32B(xt > yt) - I32B(xt < yt)
	}
	if xt < 16 { // 11(derived), 12(proj), 13(lambda), 14(native)?
		xp, yp := int32(x), int32(y)
		return Func[245+xt].(f2i)(xp, yp)
	}
	if xt > Lt {
		xp, yp := int32(x), int32(y)
		r = cmL(xp, yp)
		if r == 0 {
			r = cmL(xp+8, yp+8)
		}
		return r
	}
	xn, yn := nn(x), nn(y)
	xp = int32(x)
	yp = int32(y) - xp
	n := mini(xn, yn)
	s := sz(xt)
	e := xp + n*s
	for xp < e {
		r = Func[234+xt].(f2i)(xp, xp+yp)
		if r != 0 {
			return r
		}
		xp += s
	}
	return I32B(xn > yn) - I32B(xn < yn)
}
func Kst(x K) K { return Atx(Ks(32), x) } // `k@
func Lst(x K) K { return Atx(Ks(40), x) } // `l@
func Str(x K) K {
	var r K
	xt := tp(x)
	if xt > 16 {
		return Ech(17, l1(x))
	}
	xp := int32(x)
	if xt > 8 {
		switch xt - cf {
		case 0: // cf
			rx(x)
			r = Rdc(13, l1(Rev(Str(ti(Lt, xp)))))
		case 1: // df
			r = ucat(Str(x0(x)), Str(21+x1(x)))
		case 2: //pf
			f := x0(x)
			l := x1(x)
			i := x2(x)
			ft := tp(f)
			f = Str(f)
			dx(i)
			if nn(i) == 1 && I32(int32(i)) == 1 && (ft == 0 || ft == df) {
				r = ucat(Kst(Fst(l)), f)
			} else {
				r = ucat(f, emb('[', ']', ndrop(-1, ndrop(1, Kst(l)))))
			}
		default: //lf, native
			r = x1(x)
		}
		dx(x)
		return r
	} else {
		switch xt {
		case 0:
			if xp > 448 {
				return Str(K(xp) - 448)
			}
			ip := xp
			switch xp >> 6 {
			case 0: //  0..63  monadic
				if xp == 0 {
					return mk(Ct, 0)
				}
			case 1: // 64..127 dyadic
				ip -= 64
			case 2: // 128     dyadic indirect
				ip -= 128
			default: // 192     tetradic
				ip -= 192
			}
			if ip > 25 || ip == 0 {
				return ucat(Ku('`'), si(xp))
			}
			r = Ku(uint64(I8(226 + ip)))
		case 1: //not reached
			r = 0
		case ct:
			r = Ku(uint64(xp))
		case it:
			r = si(xp)
		case st:
			r = cs(x)
		case ft:
			r = sf(F64(xp))
		default:
			r = sfz(F64(xp), F64(xp+8))
		}
	}
	dx(x)
	return r
}
func emb(a, b int32, x K) K { return cat1(Cat(Kc(a), x), Kc(b)) }
func si(x int32) K {
	if x == 0 {
		return Ku(uint64('0'))
	} else if x == nai {
		return Ku(20016) // 0N
	} else if x < 0 {
		return ucat(Ku(uint64('-')), si(-x))
	}
	r := mk(Ct, 0)
	for x != 0 {
		r = cat1(r, Kc('0'+x%10))
		x /= 10
	}
	return Rev(r)
}
func sf(x float64) K {
	c := int32(0)
	if x != x {
		return Ku(28208) // 0n
	}
	u := uint64(I64reinterpret_f64(x))
	if u == uint64(I64reinterpret_f64(inf)) {
		return Ku(30512) // 0w
	} else if u == uint64(I64reinterpret_f64(-inf)) {
		return Ku(7811117) // -0w
	}
	if x < 0 {
		return ucat(Ku(uint64('-')), sf(-x))
	}
	if x > 0 && (x >= 1e6 || x <= 1e-6) {
		return se(x)
	}
	r := mk(Ct, 0)
	i := int64(x)
	if i == 0 {
		r = cat1(r, Kc('0'))
	}
	for i != 0 {
		r = cat1(r, Kc(int32('0'+i%10)))
		i /= 10
	}

	r = Rev(r)
	r = cat1(r, Kc('.'))
	x -= F64floor(x)
	for i := int32(0); i < 6; i++ {
		x *= 10
		r = cat1(r, Kc('0'+(int32(x)%10)))
		continue
	}
	n := nn(r)
	rp := int32(r)
	for n > 0 {
		n--
		if I8(rp) == '0' {
			c++
		} else {
			c = 0
		}
		rp++
	}
	return ndrop(-c, r)
}
func se(x float64) K {
	f := x
	e := int64(0)
	if frexp1(x) != 0 {
		f = frexp2(x)
		e = frexp3(x)
	}
	x = 0.3010299956639812 * float64(e) // log10(2)*
	ei := int32(F64floor(x))
	x = x - float64(ei)
	return ucat(cat1(sf(f*pow(10.0, x)), Kc('e')), si(ei))
}
func sfz(re, im float64) K {
	if (re != re) || (im != im) {
		return Ku(6385200) // 0na
	}
	z := hypot(re, im)
	a := ang2(im, re)
	r := cat1(trdot(sf(z)), Kc('a'))
	if a != 0.0 {
		r = ucat(r, trdot(sf(a)))
	}
	return r
}
func trdot(x K) K {
	n := nn(x)
	if I8(int32(x)+n-1) == '.' {
		return ndrop(-1, x)
	}
	return x
}

func Cst(x, y K) K { // x$y
	yt := tp(y)
	if yt > Zt {
		return Ecr(17, l2(x, y))
	}
	if yt == ct {
		y, yt = Enl(y), Ct
	}
	if tp(x) != st || yt != Ct {
		trap() //type
	}
	if int32(x) == 0 { // `$"sym"
		return sc(y)
	}
	t := ts(x)
	y = val(y)
	yt = tp(y)
	if t == yt {
		return y
	}
	if y == 0 && t > 16 {
		return mk(t, 0)
	}
	if t-yt > 15 {
		y = Enl(y)
	}
	if t&15 > yt&15 {
		y = uptype(y, t&15)
	}
	return y
}
func ts(x K) T {
	c := inC(int32(Rdc(2, l1(cs(x)))), 254, 279)
	if c > 0 {
		return T(c - 253)
	}
	return 0
}
func rtp(t T, x K) K { // `c@ `i@ `s@ `f@ `z@ (reinterpret data)
	xt := tp(x)
	if uint32(xt-18) > 5 {
		trap()
	}
	n := nn(x) * sz(xt)
	m := n / sz(t)
	if n != m*sz(t) {
		trap() //length
	}
	x = use(x)
	SetI32(int32(x)-12, m)
	return ti(t, int32(x))
}
func repl(x K) {
	c := I8(int32(x))
	x = val(x)
	if x != 0 {
		if c == 32 {
			dx(Out(x))
		} else {
			write(cat1(join(Kc(10), Lst(x)), Kc(10)))
		}
	}
}

func Out(x K) K {
	write(cat1(Kst(rx(x)), Kc(10)))
	return x
}
func Otu(x, y K) K {
	write(cat1(Kst(x), Kc(':')))
	return Out(y)
}
func write(x K) {
	Write(0, 0, int32(x), nn(x))
	dx(x)
}
func readfile(x K) K { // x C
	var r K
	if nn(x) == 0 {
		dx(x)
		r = mk(Ct, 496)
		r = ntake(ReadIn(int32(r), 496), r)
		return r
	}
	n := Read(int32(x), nn(x), 0)
	if n < 0 {
		dx(x)
		return mk(Ct, 0)
	}
	r = mk(Ct, n)
	Read(int32(x), nn(x), int32(r))
	dx(x)
	return r
}
func writefile(x, y K) K { // x, y C
	r := Write(int32(x), nn(x), int32(y), nn(y))
	if r != 0 {
		trap() //io
	}
	dx(x)
	return y
}

type ftok = func() K

func tok(x K) K {
	s := cat1(src(), Kc(10))
	pp = nn(s)
	s = Cat(s, x)  // src contains all src
	pp += int32(s) // pp is the parser position within src
	pe = pp + nn(x)
	r := mk(Lt, 0)
	for {
		ws()
		if pp == pe {
			break
		}
		for i := int32(193); i < 200; i++ { // tchr, tnms, tvrb, tpct, tvar, tsym, trap
			y := Func[i].(ftok)()
			if y != 0 {
				y |= K(int64(pp-int32(s)) << 32)
				r = cat1(r, y)
				break
			}
		}
	}
	SetI32(16, int32(s)) //SetI64(512, int64(s))
	return r
}
func src() K { return ti(Ct, I32(16)) }
func tchr() K {
	if I8(pp) == '0' && pp < pe { // 0x01ab (lower case only)
		if I8(1+pp) == 'x' {
			pp += 2
			return thex()
		}
	}
	if I8(pp) != 34 {
		return 0
	}
	pp++
	r := mk(Ct, 0)
	q := uint32(0)
	for {
		if pp == pe {
			trap() //parse
		}
		c := I8(pp)
		pp++
		if c == 34 && q == 0 {
			break
		}
		if c == '\\' && q == 0 {
			q = 1
			continue
		}
		if q != 0 {
			c = cq(c)
			q = 0
		}
		r = cat1(r, Kc(c))
	}
	if nn(r) == 1 {
		return Fst(r)
	}
	return r
}
func cq(c int32) int32 { // \t \n \r \" \\   -> 9 10 13 34 92
	if c == 116 {
		return 9
	}
	if c == 110 {
		return 10
	}
	if c == 114 {
		return 13
	}
	return c
}
func thex() K {
	r := mk(Ct, 0)
	for pp < pe-1 {
		c := I8(pp)
		if is(c, 128) == 0 {
			break
		}
		r = cat1(r, Kc((hx(c)<<4)+hx(I8(1+pp))))
		pp += 2
	}
	if nn(r) == 1 {
		return Fst(r)
	}
	return r
}
func hx(c int32) int32 {
	if is(c, 4) != 0 {
		return c - '0'
	} else {
		return c - 'W'
	}
}
func tnms() K {
	r := tnum()
	for pp < pe-1 && I8(pp) == ' ' {
		pp++
		x := tnum()
		if x == 0 {
			break
		}
		t := tp(r)
		if t < 16 {
			r = Enl(r)
		}
		t = maxtype(r, x)
		r = uptype(r, t)
		r = cat1(r, uptype(x, t))
	}
	return r
}
func tnum() K {
	c := I8(pp)
	if c == '-' || c == '.' {
		if is(I8(pp-1), 64) != 0 {
			return 0 // e.g. x-1 is (x - 1) not (x -1)
		}
	}
	if c == '-' && pp < 1+pe {
		pp++
		r := tunm()
		if r == 0 {
			pp--
			return 0
		}
		return Neg(r)
	}
	return tunm()
}
func tunm() K {
	p := pp
	r := pu()
	if r == 0 && p == pp {
		if I8(p) == '.' {
			if is(I8(1+p), 4) != 0 {
				return pflt(r)
			}
		}
		return 0
	}
	if pp < pe {
		c := I8(pp)
		if c == '.' {
			return pflt(r)
		}
		if c == 'p' {
			return ppi(float64(r))
		}
		if c == 'a' {
			return pflz(float64(r))
		}
		if c == 'e' || c == 'E' {
			return Kf(pexp(float64(r)))
		}
		if r == 0 {
			if c == 'N' {
				pp++
				return missing(it)
			}
			if c == 'n' || c == 'w' {
				q := Kf(0)
				SetI64(int32(q), int64(0x7FF8000000000001)) // 0n
				if c == 'w' {
					SetF64(int32(q), inf) // 0w
				}
				pp++
				if pp < pe && I8(pp) == 'a' {
					dx(q)
					return pflz(F64(int32(q)))
				}
				return q
			}
		}
	}
	return Ki(int32(r))
}
func pu() int64 {
	r := int64(0)
	for pp < pe {
		c := I8(pp)
		if is(c, 4) == 0 {
			break
		}
		r = 10*r + int64(c-'0')
		pp++
	}
	return r
}
func pexp(f float64) float64 {
	pp++
	e := int64(1)
	if pp < pe {
		c := I8(pp)
		if c == '-' || c == '+' {
			if c == '-' {
				e = int64(-1)
			}
			pp++
		}
	}
	e *= pu()
	return f * pow(10.0, float64(e))
}
func pflt(i int64) K {
	var c int32
	d := 1.0
	f := float64(i)
	pp++ // .
	for pp < pe {
		c = I8(pp)
		if is(c, 4) == 0 {
			break
		}
		d /= 10.0
		f += d * float64(c-'0')
		pp++
	}
	if pp < pe {
		c = I8(pp)
		if c == 'e' || c == 'E' {
			f = pexp(f)
		}
	}
	if pp < pe {
		c = I8(pp)
		if c == 'a' {
			return pflz(f)
		}
		if c == 'p' {
			return ppi(f)
		}
	}
	return Kf(f)
}
func pflz(f float64) K {
	r := K(0)
	pp++
	if pp < pe {
		r = tunm()
	}
	return Rot(Kf(f), r)
}
func ppi(f float64) K {
	pp++
	return Kf(pi * f)
}

func tvrb() K {
	c := I8(pp)
	if is(c, 1) == 0 {
		return 0
	}
	pp++
	if c == 92 && I8(pp-2) == 32 { // \out
		return K(29)
	}
	o := int32(1)
	if pp < pe {
		if I8(pp) == 58 { // :
			pp++
			if is(c, 8) != 0 {
				trap() //parse
			}
			o = 97
			/*
				if is(c, 8) != 0 {
					o = 2 // ':
				} else {
					o = 97 // +:
				}
			*/
		}
	}
	return K(o + idx(c, 227, 253))
}
func tpct() K {
	c := I8(pp)
	if is(c, 48) != 0 { // ([{}]); \n
		pp++
		return K(c)
	}
	if c == 10 {
		pp++
		return K(';')
	}
	return 0
}
func tvar() K {
	c := I8(pp)
	if is(c, 2) == 0 {
		return 0
	}
	pp++
	r := Ku(uint64(c))
	for pp < pe {
		c = I8(pp)
		if is(c, 6) == 0 {
			break
		}
		r = cat1(r, ti(ct, c))
		pp++
	}
	return sc(r)
}
func tsym() K {
	r := K(0)
	for I8(pp) == 96 {
		pp++
		if r == 0 {
			r = mk(St, 0)
		}
		s := K(0)
		if pp < pe {
			s = tchr()
			if tp(s) == ct {
				s = sc(Enl(s))
			} else if s != 0 {
				s = sc(s)
			} else {
				s = tvar()
			}
		}
		if s == 0 {
			s = K(st) << 59
		}
		r = cat1(r, s)
		if pp == pe {
			break
		}
	}
	return r
}
func ws() {
	var c int32
	for pp < pe {
		c = I8(pp)
		if c == 10 || c > 32 {
			break
		}
		pp++
	}
	for pp < pe {
		c = I8(pp)
		if c == 47 && I8(pp-1) < 33 {
			pp++
			for pp < pe {
				c = I8(pp)
				if c == 10 {
					break
				}
				pp++
			}
		} else {
			return
		}
	}
}
func is(x, m int32) int32 { return m & I8(100+x) }
func nyi(x K) K           { trap(); return 0 }
func Idy(x K) K           { return x } // :x
func Dex(x, y K) K { // x:y
	dx(x)
	return y
}
func Flp(x K) K { // +x
	xt := tp(x)
	if xt == Lt {
		n := nn(x)
		xp := int32(x)
		m := Ki(maxcount(xp, n))
		x = Atx(Rdc(13, l1(Ecr(15, l2(m, x)))), Ecl(2, l2(Til(m), Mul(m, Til(Ki(n))))))
	} else if xt > Lt {
		r := x0(x)
		x = r1(x)
		if xt == Tt {
			x = Key(r, x)
		} else {
			if tp(r) != St || tp(x) != Lt {
				trap() //type
			}
			m := maxcount(int32(x), nn(x))
			x = Ech(15, l2(Ki(m), x)) // (|/#'x)#'x
			r = l2(r, x)
			SetI32(int32(r)-12, m)
			x = ti(Tt, int32(r))
		}
	}
	return x
}
func maxcount(xp int32, n int32) int32 { // |/#l
	r := int32(0)
	for n > 0 {
		n--
		x := K(I64(xp))
		xp += 8
		if tp(x) < 16 {
			r = maxi(1, r)
		} else {
			r = maxi(nn(x), r)
		}
	}
	return r
}
func Fst(x K) K { // *x
	t := tp(x)
	if t < 16 {
		return x
	}
	if t == Dt {
		return Fst(Val(x))
	}
	return ati(x, 0)
}
func Las(x K) K { // *|x
	t := tp(x)
	if t < 16 {
		return x
	}
	if t == Dt {
		x = Val(x)
	}
	n := nn(x)
	if n == 0 {
		return Fst(x)
	}
	return ati(x, n-1)
}

func Cnt(x K) K { // #x
	t := tp(x)
	dx(x)
	if t < 16 {
		return Ki(1)
	}
	return Ki(nn(x))
}
func Not(x K) K { // ~x
	if tp(x)&15 == st {
		x = Eql(Ks(0), x)
	} else {
		x = Eql(Ki(0), x)
	}
	return x
}
func Til(x K) K {
	xt := tp(x)
	if xt > Lt {
		t := x0(x)
		dx(x)
		return t
	}
	if xt == it {
		return seq(int32(x))
	}
	if xt == It {
		return Enc(x, Til(Rdc(4, l1(rx(x))))) //{x\!*/x}
	}
	trap() //type
	return 0
}
func seq(n int32) K {
	n = maxi(n, 0)
	r := mk(It, n)
	rp := int32(r)
	seqI(rp, rp + ev(4*n))
	/*
	for n > 0 {
		n--
		SetI32(int32(r)+4*n, n)
	}
	*/
	return r
}
func Unq(x K) K { // ?x
	var r K
	xt := tp(x)
	if xt < 16 {
		return roll(x)
	}
	if xt >= Lt {
		if xt == Dt {
			trap() //type
		}
		if xt == Tt {
			r = x0(x)
			x = r1(x)
			return key(r, Flp(Unq(Flp(x))), xt)
		}
	}
	rx(rx(x))
	return atv(x, Wer(Eql(seq(nn(x)), Fnd(x, x)))) // x@&(!#x)==x?x
}
func Grp(x K) K    { return kx(96, x) }     // =x grp.
func Key(x, y K) K { return key(x, y, Dt) } // x!y
func key(x, y K, t T) K { // Dt or Tt
	xt := tp(x)
	yt := tp(y)
	if xt < 16 {
		if xt == it {
			return Mod(y, x)
		}
		if xt == st {
			if yt == Tt { // s!t (key table)
				x = rx(x)
				y = rx(y)
				return Key(Tak(x, y), Drp(x, y))
			}
		}
		x = Enl(x) //allow `a!,1 2 3 short for (`a)!,1 2 3
	}
	xn := nn(x)
	if t == Tt {
		if xn > 0 {
			xn = nn(K(I64(int32(y))))
		}
	} else if yt < 16 {
		trap() //type
	} else if xn != nn(y) {
		trap() //length
	}
	x = l2(x, y)
	SetI32(int32(x)-12, xn)
	return ti(t, int32(x))
}
func Tak(x, y K) K { // x#y
	xt := tp(x)
	yt := tp(y)
	if yt == Dt {
		x = rx(x)
		if xt == it {
			r := x0(y)
			y = r1(y)
			r = Tak(x, r)
			y = Tak(x, y)
			return Key(r, y)
		} else {
			return Key(x, Atx(y, x))
		}
	} else if yt == Tt {
		if xt&15 == st {
			if xt == st {
				x = Enl(x)
			}
			x = rx(x)
			return key(x, Atx(y, x), yt)
		} else {
			return Ecr(15, l2(x, y))
		}
	}
	if xt == it {
		return ntake(int32(x), y)
	}
	y = rx(y)
	if xt > 16 && xt == yt {
		return atv(y, Wer(In(y, x))) // set take
	}
	return Atx(y, Wer(Cal(x, l1(y)))) // f#
}
func ntake(n int32, y K) K {
	var r K
	t := tp(y)
	if n == nai {
		if t < 16 {
			n = 1
		} else {
			n = nn(y)
		}
	}
	if n < 0 {
		if tp(y) < 16 {
			return ntake(-n, y)
		}
		n += nn(y)
		if n < 0 {
			return ucat(ntake(-n, missing(t-16)), y)
		}
		return ndrop(n, y)
	}
	if t < 16 {
		return atv(Enl(y), Wer(Ki(n)))
	}
	yn := nn(y)
	s := sz(t)
	yp := int32(y)
	if I32(yp-4) == 1 && bucket(s*yn) == bucket(s*n) && n <= yn && t < Lt {
		SetI32(yp-12, n)
		return y
	}
	r = seq(n)
	if n > yn && yn > 0 {
		r = Mod(r, Ki(yn))
	}
	return atv(y, r)
}
func Drp(x, y K) K { // x_y
	xt := tp(x)
	yt := tp(y)
	if yt > Lt {
		if yt == Dt || (yt == Tt && xt&15 == st) {
			r := x0(y)
			y = r1(y)
			if xt < 16 {
				x = Enl(x)
			}
			x = rx(Wer(Not(In(rx(r), x))))
			return key(Atx(r, x), Atx(y, x), yt)
		} else {
			return Ecr(16, l2(x, y))
		}
	}
	if xt == it {
		return ndrop(int32(x), y)
	}
	if xt > 16 && xt == yt {
		return atv(y, Wer(Not(In(rx(y), x)))) // set drop
	}
	if yt == it {
		return atv(x, Wer(Not(Eql(y, seq(nn(x))))))
	}
	return Atx(y, Wer(Not(Cal(x, l1(rx(y)))))) // f#
}
func ndrop(n int32, y K) K {
	var r K
	yt := tp(y)
	if yt < 16 || yt > Lt {
		trap() //type
	}
	yn := nn(y)
	if n < 0 {
		return ntake(maxi(0, yn+n), y)
	}
	rn := yn - n
	if rn < 0 {
		dx(y)
		return mk(yt, 0)
	}
	s := sz(yt)
	yp := int32(y)
	if I32(yp-4) == 1 && bucket(s*yn) == bucket(s*rn) && yt < Lt {
		r = rx(y)
		SetI32(yp-12, rn)
	} else {
		r = mk(yt, rn)
	}
	rp := int32(r)
	Memorycopy(rp, yp+s*n, s*rn)
	if yt == Lt {
		rl(r)
		r = uf(r)
	}
	dx(y)
	return r
}

func Cut(x, y K) K { // x^y
	yt := tp(y)
	if yt == it || yt == ft {
		return Pow(y, x)
	}
	xt := tp(x)
	if xt == It {
		return cuts(x, y)
	}
	if xt == Ct && yt == Ct { // "set"^"abc"
		x = rx(Wer(In(rx(y), x)))
		return rcut(y, Cat(Ki(0), Add(Ki(1), x)), Cat(x, Ki(nn(y))))
	}
	if xt != it || yt < 16 {
		trap() //type
	}
	xp := int32(x)
	if xp <= 0 {
		xp = nn(y) / -xp
	}
	r := mk(Lt, xp)
	rp := int32(r)
	e := ep(r)
	n := nn(y) / xp
	x = seq(n)
	for rp < e {
		SetI64(rp, int64(atv(rx(y), rx(x))))
		x = Add(Ki(n), x)
		rp += 8
		continue
	}
	dxy(x, y)
	return r
}
func cuts(x, y K) K { return rcut(y, x, cat1(ndrop(1, rx(x)), Ki(nn(y)))) }
func rcut(x, a, b K) K { // a, b start-stop ranges
	n := nn(a)
	ap, bp := int32(a), int32(b)
	r := mk(Lt, n)
	rp := int32(r)
	for n > 0 {
		n--
		o := I32(ap)
		m := I32(bp) - o
		if m < 0 {
			trap() //value
		}
		SetI64(rp, int64(atv(rx(x), Add(Ki(o), seq(m)))))
		rp += 8
		ap += 4
		bp += 4
	}
	dxy(a, b)
	dx(x)
	return r
}
func split(x, y K) K {
	xt, yt := tp(x), tp(y)
	xn := int32(1)
	if yt == xt+16 {
		x = Wer(Eql(x, rx(y)))
	} else {
		if xt == yt && xt == Ct {
			xn = nn(x)
			x = Find(x, rx(y))
		} else {
			trap() //type
		}
	}
	x = rx(x)
	return rcut(y, Cat(Ki(0), Add(Ki(xn), x)), cat1(x, Ki(nn(y))))
}
func join(x, y K) K { // {(-#x)_,/y,\x}
	n := -int32(Cnt(rx(x)))
	return ndrop(n, Rdc(13, l1(Ecl(13, l2(y, x)))))
}
func Bin(x, y K) K { // x'y
	var r K
	xt := tp(x)
	yt := tp(y)
	if xt == yt || yt == Lt {
		return Ecr(40, l2(x, y))
	} else if xt == yt+16 {
		r = Ki(ibin(x, y, xt))
	} else {
		trap() //type
	}
	dxy(x, y)
	return r
}
func ibin(x, y K, t T) int32 {
	var h int32
	k := int32(0)
	n := nn(x)
	xp := int32(x)
	yp := int32(y)
	j := n - 1
	s := sz(t)
	switch s >> 2 {
	case 0:
		for {
			if k > j {
				return k - 1
			}
			h = (k + j) >> 1
			if I8(xp+h) > yp {
				j = h - 1
			} else {
				k = h + 1
			}
		}
	case 1:
		for {
			if k > j {
				return k - 1
			}
			h = (k + j) >> 1
			if I32(xp+4*h) > yp {
				j = h - 1
			} else {
				k = h + 1
			}
		}
	default:
		f := F64(yp)
		for {
			if k > j {
				return k - 1
			}
			h = (k + j) >> 1
			if F64(xp+8*h) > f {
				j = h - 1
			} else {
				k = h + 1
			}
		}
	}
	return 0 // not reached
}
func Flr(x K) K { // _x
	var r K
	rp := int32(0)
	xt := tp(x)
	xp := int32(x)

	if xt < 16 {
		switch xt - 2 {
		case 0: // c
			return Kc(lc(xp))
		case 1: // i
			return Kc(xp)
		case 2: // s
			return Ki(int32(xp))
		case 3: // f
			dx(x)
			return Ki(int32(F64floor(F64(xp))))
		case 4: // z
			dx(x)
			return Kf(F64(xp))
		default:
			return x
		}
	}
	xn := nn(x)
	switch xt - 18 {
	case 0: //C
		return lower(x)
	case 1: //I
		r = mk(Ct, xn)
		rp = int32(r)
		e := rp + xn
		for rp < e {
			SetI8(rp, I32(xp))
			xp += 4
			rp++
		}
	case 2: //S
		x = use(x)
		return ti(It, int32(x))
	case 3: //F
		r = mk(It, xn)
		rp = int32(r)
		for xn > 0 {
			xn--
			SetI32(rp, int32(F64floor(F64(xp))))
			xp += 8
			rp += 4
		}
	case 4: // Z
		r = atv(rtp(Ft, rx(x)), Mul(Ki(2), seq(xn)))
	default: // L/D/T
		return Ech(16, l1(x))
	}
	dx(x)
	return r
}
func lower(x K) K {
	x = use(x)
	p := int32(x)
	e := p + nn(x)
	for p < e {
		SetI8(p, lc(I8(p)))
		p++
	}
	return x
}
func lc(x int32) int32 { return x + 32*I32B(uint32(x-65) < 26) }

func Rev(x K) K { // |x
	var r K
	t := tp(x)
	if t < 16 {
		return x
	}
	if t == Dt {
		r = x0(x)
		return Key(Rev(r), Rev(r1(x)))
	}
	xn := nn(x)
	if xn < 2 {
		return x
	}
	r = mk(It, xn)
	rp := int32(r)
	for xn > 0 {
		xn--
		SetI32(rp, xn)
		rp += 4
	}
	return atv(x, r)
}

func Wer(x K) K { // &x
	r := K(0)
	t := tp(x)
	if t < 16 {
		x = Enl(x)
		t = tp(x)
	}
	if t == Dt {
		r = x0(x)
		return Atx(r, Wer(r1(x)))
	}
	xn := nn(x)
	xp := int32(x)
	if t == It {
		n := sumi(xp, ep(x))
		r = mk(It, n)
		rp := int32(r)
		for i := int32(0); i < xn; i++ {
			j := I32(xp)
			for j > 0 {
				j--
				SetI32(rp, i)
				rp += 4
			}
			xp += 4
		}
	} else if xn == 0 {
		r = mk(It, 0)
	} else {
		trap() //type
	}
	dx(x)
	return r
}
func Fwh(x K) K { // *&x
	t := tp(x)
	if t == It {
		dx(x)
		p := int32(x)
		e := ep(x)
		for p < e {
			if I32(p) != 0 {
				return Ki((p - int32(x)) >> 2)
			}
			p += 4
		}
		return Ki(nai)
	}
	return Fst(Wer(x))
}
func Typ(x K) K { // @x
	dx(x)
	return sc(Ku(uint64(I8(253 + int32(tp(x))))))
}
func Tok(x K) K { // `t@"src"
	if tp(x) == Ct {
		return tok(x)
	} else {
		return x
	}
}
func Val(x K) K {
	xt := tp(x)
	if xt == st {
		return lup(x)
	}
	if xt == Ct {
		return val(x)
	}
	if xt == lf || xt == xf { // lambda: (code;string;locals;arity)
		//xp := int32(x)  // native: (ptr;string;arity)
		r := l2(x0(x), x1(x))
		if xt == lf {
			r = cat1(r, x2(x))
		}
		dx(x)
		return cat1(r, Ki(nn(x)))
	}
	if xt == Lt {
		return exec(x) // .L e.g. 1+2 is (1;2;`66)
	}
	if xt > Lt {
		return r1(x)
	} else {
		trap() //type
		return 0
	}
}
func val(x K) K {
	x = parse(tok(x))
	xn := nn(x)
	xp := int32(x) + 8*(xn-1)
	a := int32(0)
	if xn > 2 && I64(xp) == 64 {
		a = 1
	}
	x = exec(x)
	if a != 0 {
		dx(x)
		return 0
	}
	return x
}
func Enc(x, y K) K {
	yt := tp(y)
	if yt == It {
		return cal(lup(Ks(104)), l2(x, y))
	}
	if yt != it {
		trap()
	}
	yi := int32(y)
	n := int32(0)
	if tp(x) == It {
		n = nn(x)
	}
	r := mk(It, 0)
	for {
		n--
		xi := int32(ati(rx(x), n))
		r = cat1(r, Ki(modi(yi, xi)))
		yi = divi(yi, xi)
		if n == 0 {
			break
		}
		if n < 0 && uint32(yi+1) < 2 {
			if yi == -1 {
				r = cat1(r, Ki(-1))
			}
			break
		}
	}
	dx(x)
	return Rev(r)
}
func Dec(x, y K) K { // x//y   {z+x*y}/[0;x;y]
	if tp(y) < 16 {
		trap() //type
	}
	r := Fst(rx(y))
	n := nn(y)
	for i := int32(1); i < n; i++ {
		r = Add(ati(rx(y), i), Mul(ati(rx(x), i), r))
	}
	dxy(x, y)
	return r
}
func zk() {
	Data(280, "`k`l`a`b`while`\"rf.\"`\"rz.\"`\"gdt.\"`\"grp.\"`\"enc.\"\n`x:,/+\"0123456789abcdef\"@16 16\\256!\n`t:`39\n`p:`46\n`enc:{$[#y;+(&'(|/c)-c:#'r),'r:{x\\y}/[x;y];(#x)#,!0]}\n`gdt:{[t;g]($[g;{x@>y x};{x@<y x}])/(,!#t),|.t}\n`grp:{(x@*'g)!g:(&~a~'a@-1+!#a:x i)^i:<x}\nabs:`32;sin:`44;cos:`45;find:`31;imag:`33;conj:`34;angle:`35;exp:`42;log:`43\n`pad:{x@\\!|/#'x}\n`lxy:{\nkt:{[x;y;k;T]x:$[`T~@x;T[x;k];`pad(\"\";\"-\"),$x];(x,'\"|\"),'T[y;k]}\nd:{[x;k;kt;T]r:!x;x:.x;$[`T~@x;kt[r;x;k;T];,'[,'[`pad(k'r);\"|\"];k'x]]}\nT:{[x;k]$[`L?@'.x;,k x;(,*x),(,(#*x)#\"-\"),1_x:\" \"/'+`pad@'$(!x),'.x]}\nt:@y;k:`kxy@*x;h:*|x\ndd:(\"\";,\"..\")h<#y:$[(@y)?`L`D`T;y;y~*y;y;[t:`L;,y]]\ny:$[y~*y;y;(h&#y)#y]\n$[`D~t;d[y;k;kt;T];`T~t;T[y;k];y~*y;,k y;k'y],dd}\n`l:`lxy 70 20\n`str:{q:{c,(\"\\\\\"/(0,i)^@[x;i;(qs!\"tnr\\\"\\\\\")x i:&x?\\qs:\"\\t\\n\\r\\\"\\\\\"]),c:_34}\n$[|/x?\\\"\\t\\n\\r\"__!31;\"0x\",`x@x;q x]}\n`kxy:{\na:{t:@x;x:$x;$[`c~t;`str x;`s~t;\"`\",x;x]}\nd:{[x;k]r:\"!\",k@.x;n:#!x;x:k@!x;$[(n<2)|(@.x)?`D`T;\"(\",x,\")\";x],r}\nv:{[x;k;m]t:@x;x:(m&n:#x)#x\nx:$[`L~t;k'x;`C~t;x;$x]\nx:$[`C~t;`str x;`S~t;c,(c:\"`\")/x;`L~t;$[1~n;*x;\"(\",(\";\"/x),\")\"];\" \"/x]\n$[m<#x:((\"\";\",\")(1~n)),x;((m-2)#x),\"..\";x]}\nt:@y;k:`kxy x\n$[`T~t;\"+\",d[+y;k];`D~t;d[y;k];0~#y;(`C`I`S`F`Z`L!(\"\\\"\\\"\";\"!0\";\"0#`\";\"0#0.\";\"0@0a\";\"()\"))t;y~*y;a y;v[y;k;x]]}\n`k:`kxy 1000000\n`d:{x-(*x),-1_x}\n`rf: {.5+(x?0)%4294967295.}\n`rf1:{.5+(1.+x?0)%4294967295.}        \n`rz: {(%-2*log `rf1 x)@360.*`rf x}\n")
	zn := int32(1363) // should end before 2k
	x := mk(Ct, zn)
	Memorycopy(int32(x), 280, zn)
	dx(Val(x))
}

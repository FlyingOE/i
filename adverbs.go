package main

import (
	. "github.com/ktye/wg/module"
)

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
	if t == 0 && xt < Lt && int32(f) < 11 { // :+-*%&|<>=
		return prior(int32(f), x)
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
func prior(f int32, x K) K {
	n := nn(x)
	if n == 0 {
		if f < 8 {
			dx(x)
			return mk(It, 0)
		}
		return x
	}
	a := Ki(0) // <=> first is always 0
	if f < 8 { // :' rotates
		a = ati(rx(x), (n-1)*I32B(f == 1)) // +-*%&| keep 1st element
	}
	return sti(Func[f+64].(f2)(x, atv(rx(x), Add(Ki(-1), seq(n)))), 0, a)
}

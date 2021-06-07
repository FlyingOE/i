package main

import (
	"fmt"

	. "github.com/ktye/wg/module"
)

func parse(x K) (r K) {
	l := tok(x)
	pp = int32(l)
	pe = pp + 8*nn(l)
	r = es()
	if pp != pe {
		trap(Parse)
	}
	lfree(l)
	return r
}
func es() (r K) {
	r = mk(Lt, 0)
	for {
		n, _ := next()
		if n == 0 {
			break
		}
		if n == 59 {
			continue
		}
		pp -= 8
		x := e(t())
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
func e(x K, xv int32) (r K) { // Lt
	if x == 0 {
		return 0
	}
	y, yv := t()
	if y == 0 {
		return x
	}
	if yv != 0 && xv == 0 {
		r = e(t())
		if r == 0 { // 1+ (projection)
			return cat1(ucat(ucat(l1(0), x), y), 128)
		}
		x, y = pasn(x, y)
		r = ucat(r, x)
		return dyadic(r, y) // dyadic
	}
	r = e(y, yv)
	if xv == 0 {
		return cat1(ucat(r, x), 83) // juxtaposition
	} else if r == y && xv+yv == 2 {
		return cat1(ucat(r, x), 91) // composition
	}
	r = ucat(r, x)
	return monadic(r) // monadic
}
func t() (r K, verb int32) { // Lt
	var ln, s int32
	r, s = next()
	if r == 0 {
		return 0, 0
	}
	if r < 127 && is(int32(r), 32) {
		pp -= 8
		return 0, 0
	}
	if r == K('(') {
		r = rlist(plist(41))
	} else if r == K('{') {
		r = plam(s)
	} else if tp(r) == st {
		r = l2(r, 20) // .`x (lookup)
	} else {
		rt := tp(r)
		if rt == 0 {
			r, verb = quote(r)|K(s)<<32, 1
		} else if rt == St {
			if nn(r) == 1 {
				r = Fst(r)
			}
		}
		r = l1(r)
	}
f:
	for {
		n, _ := next()
		if n == 0 {
			break f
		}
		a := int32(n)
		tn := tp(n)
		if tn == 0 && a > 20 && a < 27 { // +/
			r, verb = cat1(r, n), 1
		} else if n == 91 { // [
			verb = 0
			n, ln = plist(93)
			n, ln = pspec(r, n, ln)
			if ln < 0 {
				return n, 0
			}
			if ln == 1 {
				r = cat1(ucat(Fst(n), r), 83)
			} else {
				n = rlist(n, ln)
				r = cat1(Cat(n, r), 84)
			}
		} else {
			pp -= 8
			break f // within else-if
		}
	}
	return r, verb
}
func pasn(x, y K) (K, K) {
	l := K(I64(int32(y)))
	v := int32(l)
	if nn(y) == 1 && tp(l) == 0 && v == 449 || (v > 544 && v < 565) {
		dx(y)
		xn := nn(x)
		if xn > 2 { // indexed
			if v > 544 { // indexed-modified
				l -= 96
			}
			s := ati(rx(x), xn-3)
			if lastp(x) == 83 { // (+;.i.;`x;.@) -> x:@[x;.i.;+;rhs] which is (+;.i.;`x;0;211)
				x = cat1(cat1(ucat(l1(l), ndrop(-2, x)), 20), 211)
			} else {
				trap(Nyi) // indexed-modified-in-depth
			}
			y = l2(s, 448) // s:..
		} else if v == 449 || v == 545 {
			s := Fst(x) // (`x;.)
			if loc != 0 && v == 494 {
				loc = cat1(loc, s)
			}
			x = l1(s)
			y = l1(448) // asn
		} else { // modified
			y = cat1(l2(unquote(l-32), Fst(rx(x))), 448)
		}
	}
	return x, y
}
func plam(s0 int32) (r K) {
	fmt.Println("plam")
	loc = mk(St, 0)
	c := es() // todo: translate srcp
	n, s1 := next()
	if n != 125 {
		trap(Parse)
	}
	cn := nn(c)
	cp := int32(c)
	ar := int32(0)
	for i := int32(0); i < cn; i++ {
		if I64(cp) == 20 {
			if y := I32(cp-8) >> 3; y > 1 && y < 4 {
				ar = maxi(ar, y)
			}
		}
		cp += 8
	}
	i := Add(seq(1+s1-s0), Ki(s0-1))
	s := atv(rx(src), i)
	loc = Cat(ntake(ar, xyz), Unq(loc))
	cn = nn(loc)
	r = cat1(l3(c, loc, mk(Lt, cn)), s)
	loc = 0
	rp := int32(r)
	SetI32(rp-12, ar)
	return l1(K(rp) | K(lf)<<59)
}
func pspec(r, n K, ln int32) (K, int32) {
	v := K(I64(int32(r)))
	if nn(r) == 1 && ln > 2 { // $[..] cond
		if tp(v) == 0 && int32(v) == 465 {
			dx(r)
			return cond(n, ln), -1
		}
	}
	if nn(r) == 2 && ln > 1 && int32(v) == 40 { // while[..]
		dx(r)
		return whl(n, ln-1), -1
	}
	return n, ln
}
func whl(x K, xn int32) (r K) {
	r = cat1(Fst(rx(x)), 0)
	p := nn(r) - 1
	r = cat1(r, 384) // jif
	r = cat1(r, 256) // drop
	xp := int32(x)
	sum := int32(16)
	for i := int32(0); i < xn; i++ {
		if i != 0 {
			r = cat1(r, 256)
		}
		xp += 8
		y := x0(xp)
		sum += 1 + nn(y)
		r = ucat(r, y)
	}
	r = cat1(cat1(r, Ki(-8*(2+nn(r)))), 320) // jmp back
	SetI64(int32(r)+8*p, int64(Ki(8*sum)))   // jif
	dx(x)
	return ucat(cat1(l1(K(st)<<59), 0), r) // <null> for empty while
}
func cond(x K, xn int32) (r K) {
	xp := int32(x) + 8*xn
	var next, sum int32
	state := int32(1)
	for xp != int32(x) {
		xp -= 8
		r = K(I64(xp))
		if sum > 0 {
			state = 1 - state
			if state != 0 {
				r = cat1(cat1(r, Ki(next)), 384) // jif
			} else {
				r = cat1(cat1(r, Ki(sum)), 320) // j
			}
			SetI64(xp, int64(r))
		}
		next = 8 * nn(r)
		sum += next
	}
	return flat(x)
}
func plist(c K) (r K, n int32) {
	r = mk(Lt, 0)
	for {
		b, _ := next()
		if b == 0 || b == c {
			break
		}
		if n == 0 {
			pp -= 8
		}
		if n != 0 && b != 59 {
			trap(Parse)
		}
		n++
		x := e(t())
		r = cat1(r, x)
		if x == 0 {
			// r = cat1(r, 448) // in reverse: 448 0 return null
		}
	}
	return r, n
	/*
		if n == 1 {
			return Fst(r), 1
		}
		return r, n // cat3(flat(Rev(r)), Ki(n), 27, 1), 1
	*/
}
func rlist(x K, n int32) K {
	if n == 1 {
		return Fst(x)
	}
	return cat1(cat1(flat(Rev(x)), Ki(n)), 27)
}

func next() (r K, s int32) {
	if pp == pe {
		return 0, 0
	}
	r = K(I64(pp))
	s = 0xffffff & int32(r>>32)
	r = r &^ (K(0xffffff) << 32)
	pp += 8
	return r, s
}
func lastp(x K) K { return K(I64(int32(x) + 8*(nn(x)-1))) }
func dyadic(x, y K) K {
	l := lastp(y)
	if quoted(l) {
		return cat1(ucat(x, ndrop(-1, y)), 64+unquote(l))
	}
	return cat1(ucat(x, y), 128)
}
func monadic(x K) K {
	l := lastp(x)
	if quoted(l) {
		return cat1(ndrop(-1, x), unquote(l))
	}
	return cat1(x, 83) // dyadic-@
}

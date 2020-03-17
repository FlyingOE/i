// +build ignore

// reference implementation for k.w
// go run k.go t
// go run k.go 5 mki til rev

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

const trace = false

type c = byte
type s = string
type i = uint32
type j = uint64
type f = float64

var MC []c // MC, MI, MJ, MF share array (see msl)
var MI []i
var MJ []j
var MF []f
var WT []i
var MT [256]interface{}

type vt1 func(i) i
type vt2 func(i, i) i
type slice struct {
	p uintptr
	l int
	c int
}

const naI i = 2147483648
const naJ j = 9221120237041090561

func main() {
	if len(os.Args) == 2 && os.Args[1] == "t" {
		runtest()
	} else {
		fmt.Println(run(strings.Join(os.Args[1:], " ")))
	}
}
func runtest() {
	b, e := ioutil.ReadFile("t")
	if e != nil {
		panic(e)
	}
	v := strings.Split(strings.TrimSpace(string(b)), "\n")
	for i := range v {
		if len(v[i]) == 0 || v[i][0] == '/' {
			continue
		}
		vv := strings.Split(v[i], " /")
		if len(vv) != 2 {
			panic("test file")
		}
		in := strings.TrimSpace(vv[0])
		exp := strings.TrimSpace(vv[1])
		got := run(in)
		fmt.Println(in, "/", got)
		if exp != got {
			fmt.Printf("!")
			os.Exit(1)
		}
	}
}
func parseVector(s string) i {
	fc := ":+-*%&|<>=!~,^#_$?@."
	if len(s) > 1 && s[1] == ':' && strings.Index(fc, s[:1]) != -1 {
		return i(s[0])
	}
	if len(s) > 0 && strings.Index(fc, s[:1]) != -1 {
		return i(128 + s[0])
	}
	if len(s) > 0 && s[0] == '"' { // "char"
		s = strings.Trim(s, `"`)
		b := []c(s)
		r := mk(1, i(len(b)))
		for i := 0; i < len(b); i++ {
			MC[8+int(r)+i] = b[i]
		}
		return r
	}
	if len(s) > 0 && s[0] == '`' { // `symbols`b`c
		v := strings.Split(s[1:], "`")
		sn := i(len(v))
		sv := mk(5, sn)
		for i := i(0); i < sn; i++ {
			b := v[i]
			rn := uint32(len(b))
			r := mk(1, rn)
			for k := uint32(0); k < rn; k++ {
				MC[8+r+k] = b[k]
			}
			sI(sv+8+4*i, r)
		}
		return sv
	}
	if len(s) > 0 && s[0] == '(' { // (`list;1;2)
		return parseList(s[1:])
	}
	f := strings.Index(s, ".") // 1.23,2.34 (float)
	v := strings.Split(s, ",") // 1,2,3 (int vector)
	n := uint32(len(v))
	iv := make([]int64, n)
	fv := make([]float64, n)
	var e error
	for i, s := range v {
		if f == -1 {
			iv[i], e = strconv.ParseInt(s, 10, 32)
		} else {
			fv[i], e = strconv.ParseFloat(s, 64)
		}
		if e != nil {
			panic("parse number: " + s)
		}
	}
	if f == -1 {
		x := mk(2, n)
		for i := uint32(0); i < n; i++ {
			MI[(x+8+i*4)>>2] = uint32(iv[i])
		}
		return x
	} else {
		x := mk(3, n)
		for i := uint32(0); i < n; i++ {
			MF[(x+8+i*8)>>3] = fv[i]
		}
		return x
	}
}
func parseList(s string) i {
	if len(s) == 0 || s[len(s)-1] != ')' {
		panic("parse list")
	} else if len(s) == 1 {
		return mk(6, 0)
	}
	r := make([]i, 0)
	s = s[:len(s)-1]
	l, a := 0, 0
	for i, c := range s {
		if c == '(' {
			l++
		} else if c == ')' {
			l--
			if l < 0 {
				panic(")")
			}
		} else if l == 0 && c == ';' {
			r = append(r, parseVector(s[a:i]))
			a = i + 1
		}
	}
	r = append(r, parseVector(s[a:]))
	x := mk(6, i(len(r)))
	for k := range r {
		MI[2+(x>>2)+i(k)] = r[k]
	}
	return x
}
func run(s string) string {
	m0 := 16
	MJ = make([]j, (1<<m0)>>3)
	msl()
	ini(16)
	x := parseVector(s)
	return kst(evl(x))
}
func ini(x i) i {
	sJ(0, 289360742959022336) // uint64(0x0404041008040100)
	sI(128, x)
	p := i(256)
	for i := i(8); i < x; i++ {
		sI(4*i, p)
		p *= 2
	}
	copy(MT[0:], []interface{}{
		//   1    2    3    4    5    6    7    8    9    10   11   12   13   14   15
		nil, gtC, gtI, gtF, gtL, gtL, nil, nil, nil, eqC, eqI, eqF, eqZ, eqL, eqL, nil, // 000..015
		abc, abi, abf, abz, nec, nei, nef, nez, nil, nil, nil, nil, sqc, sqi, sqf, sqz, // 016..031
		nil, til, nil, cnt, str, sqr, wer, nil, nil, nil, fst, abs, enl, neg, val, nil, // 032..047
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, grd, eql, gdn, unq, // 048..063
		typ, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, // 064..079
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, srt, flr, // 080..095
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, // 096..111
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, rev, nil, not, nil, // 112..127
		nag, nac, nai, naf, naz, nas, nal, nil, nil, nil, nil, nil, nil, nil, nil, nil, // 128..143
		adc, adi, adf, adz, suc, sui, suf, suz, muc, mui, muf, muz, dic, dii, dif, diz, // 144..159
		nil, mkd, nil, rsh, cst, diw, min, nil, nil, nil, mul, add, cat, sub, cal, nil, // 160..175
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, les, eql, mor, fnd, // 176..191
		atx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, // 192..207
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, exc, cut, // 208..223
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, // 224..239
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, max, nil, mtc, nil, // 240..255
	})
	return x
}
func msl() { // update slice headers after set/inc MJ
	cp := *(*slice)(unsafe.Pointer(&MC))
	ip := *(*slice)(unsafe.Pointer(&MI))
	jp := *(*slice)(unsafe.Pointer(&MJ))
	fp := *(*slice)(unsafe.Pointer(&MF))
	fp.l, fp.c, fp.p = jp.l, jp.c, jp.p
	ip.l, ip.c, ip.p = jp.l*2, jp.c*2, jp.p
	cp.l, cp.c, cp.p = ip.l*4, ip.c*4, ip.p
	MF = *(*[]f)(unsafe.Pointer(&fp))
	MI = *(*[]i)(unsafe.Pointer(&ip))
	MC = *(*[]c)(unsafe.Pointer(&cp))
	// todo Z
}
func bk(t, n i) (r i) {
	r = i(32 - bits.LeadingZeros32(7+n*i(C(t))))
	if r < 4 {
		return 4
	}
	return r
}
func mk(x, y i) i {
	t := bk(x, y)
	i := 4 * t
	for I(i) == 0 {
		i += 4
	}
	if i == 128 {
		panic("oom")
	}
	a := I(i)
	sI(i, I(a))
	for j := i - 4; j >= 4*t; j -= 4 {
		u := a + 1<<(j>>2)
		sI(u, I(j))
		sI(j, u)
	}
	sI(a, y|x<<29)
	sI(a+4, 1)
	return a
}
func fr(x i) {
	xt, xn, _ := v1(x)
	t := 4 * bk(xt, xn)
	sI(x, I(t))
	sI(t, x)
}
func dx(x i) {
	if x > 255 {
		xr := I(x + 4)
		sI(x+4, xr-1)
		if xr == 1 {
			xt, xn, xp := v1(x)
			if xt > 4 {
				for i := i(0); i < xn; i++ {
					dx(I(xp + 4*i))
				}
			}
			fr(x)
		}
	}
}
func rx(x i) {
	if x > 255 {
		MI[1+x>>2]++
	}
}
func rl(x i) {
	_, xn, xp := v1(x)
	for i := i(0); i < xn; i++ {
		rx(I(xp))
		xp += 4
	}
}
func dxr(x, r i) i     { dx(x); return r }
func dxyr(x, y, r i) i { dx(x); dx(y); return r }
func mki(i i) (r i)    { r = mk(2, 1); sI(r+8, i); return r }
func mkd(x, y i) (r i) {
	x, y = ext(x, y)
	r = mk(7, 2)
	MI[2+r>>2] = x
	MI[3+r>>2] = y
	return r
}
func v1(x i) (xt, xn, xp i) { u := I(x); return u >> 29, u & 536870911, 8 + x }
func v2(x, y i) (xt, yt, xn, yn, xp, yp i) {
	xt, xn, xp = v1(x)
	yt, yn, yp = v1(y)
	return
}
func use(x i) (r i) {
	if I(x+4) == 1 {
		return x
	}
	xt, xn, xp := v1(x)
	r = mk(xt, xn)
	mv(r+8, xp, xn*i(C(xt)))
	dx(x)
	return r
}
func lrc(x, f i) (r i) { // list recurse
	_, xn, xp := v1(x)
	rl(x)
	g := MT[f].(func(i) i)
	r = mk(6, xn)
	rp := r + 8
	for i := i(0); i < xn; i++ {
		sI(rp, g(xp))
		rp += 4
		xp += 4
	}
	return dxr(x, r)
}
func drc(x, f i) (r i) { // dict recurse
	g := MT[f].(func(i) i)
	dx(x + 8)
	dx(x + 12)
	dx(x)
	return mkd(x+8, g(12+x))
}
func drx(x, f i) (r i) { // dict recurse on values
	rx(12 + x)
	dx(x)
	g := MT[f].(func(i) i)
	return g(12 + x)
}
func mv(dst, src, n i) { copy(MC[dst:dst+n], MC[src:src+n]) }
func ext(x, y i) (rx, ry i) {
	xt, yt, xn, yn, _, _ := v2(x, y)
	if xt != yt {
		trap()
	}
	if xn == yn {
		return x, y
	}
	if xn == 1 && yn > 1 {
		return take(x, yn), y
	}
	if xn > 1 && yn == 1 {
		return x, take(y, xn)
	}
	panic("length")
}
func upx(x, y i) (i, i) {
	xt, yt, xn, yn, _, _ := v2(x, y)
	if xt == yt {
		return x, y
	}
	if xt >= 5 || yt >= 5 {
		trap()
	}
	for xt < yt {
		x = up(x, xt, xn)
		xt++
	}
	for yt < xt {
		y = up(y, yt, yn)
		yt++
	}
	return x, y
}
func up(x, t, n i) (r i) {
	r = mk(t+1, n)
	xp, rp := x+8, r+8
	switch t {
	case 1:
		for i := i(0); i < n; i++ {
			sI(rp, uint32(C(xp+i)))
			rp += 4
		}
	case 2:
		for i := i(0); i < n; i++ {
			sF(rp, float64(I(xp)))
			xp += 4
			rp += 8
		}
	default:
		trap()
	}
	return dxr(x, r)
}
func expl(x i) (r i) {
	_, n, _ := v1(x)
	r = mk(6, n)
	rp := r + 8
	for i := i(0); i < n; i++ {
		rx(x)
		sI(rp, atx(x, mki(i)))
		rp += 4
	}
	return dxr(x, r)
}
func til(x i) (r i) {
	xt, _, xp := v1(x)
	if xt != 2 {
		trap()
	}
	n := I(xp)
	dx(x)
	if ii := int32(n); ii < 0 {
		return tir(i(-ii))
	}
	return seq(0, n, 1)
}
func seq(a, n, s i) (r i) {
	r = mk(2, n)
	rp := r + 8
	for i := i(0); i < n; i++ {
		sI(rp, s*(a+i))
		rp += 4
	}
	return r
}
func tir(n i) (r i) {
	r = mk(2, n)
	rp := 8 + r + 4*(n-1)
	for i := i(0); i < n; i++ {
		sI(rp, i)
		rp -= 4
	}
	return r
}
func rev(x i) (r i) {
	_, n, _ := v1(x)
	if n == 0 {
		return x
	}
	return atx(x, tir(n))
}
func fst(x i) (r i) {
	xt, _, _ := v1(x)
	if xt == 7 {
		return drx(x, '*')
	}
	return atx(x, mki(0))
}
func drop(x, n i) (r i) {
	_, xn, _ := v1(x)
	if n > xn {
		n = xn
	}
	return atx(x, seq(n, xn-n, 1))
}
func cut(x, y i) (r i) {
	xt, _, xn, yn, xp, _ := v2(x, y)
	if xt != 2 {
		panic("type")
	}
	if xn == 1 {
		n := I(xp)
		return dxr(x, drop(y, n))
	}
	r = mk(6, xn)
	rp := r + 8
	for i := i(0); i < xn; i++ {
		a := I(xp)
		b := I(xp + 4)
		if i == xn-1 {
			b = yn
		}
		if b < a {
			panic("domain")
		}
		rx(y)
		sI(rp, atx(y, seq(a, b-a, 1)))
		xp += 4
		rp += 4
	}
	return dxyr(x, y, r)
}
func rsh(x, y i) (r i) {
	xt, _, xn, _, xp, _ := v2(x, y)
	if xt != 2 {
		panic("type")
	}
	n := prod(xp, xn)
	r = take(y, n)
	if xn == 1 {
		return dxr(x, r)
	}
	xn--
	xe := xp + 4*xn
	for i := i(0); i < xn; i++ {
		m := I(xe)
		n /= m
		n = prod(xp, xn-i)
		r = cut(seq(0, n, m), r)
		xe -= 4
	}
	return dxr(x, r)
}
func prod(xp, n i) (r i) {
	r = 1
	for i := i(0); i < n; i++ {
		r *= I(xp)
		xp += 4
	}
	return r
}
func take(x, n i) (r i) {
	_, xn, _ := v1(x)
	r = seq(0, n, 1)
	if xn < n {
		rp := 8 + r
		for i := i(0); i < n; i++ {
			sI(rp, I(rp)%xn)
			rp += 4
		}
	}
	return atx(x, r)
}
func atx(x, y i) (r i) {
	//fmt.Printf("atx x=%x y=%x\n", x, y)
	//fmt.Printf("atx:x=%s y=%s\n", kst(x), kst(y))
	xt, yt, xn, yn, xp, yp := v2(x, y)
	if xt == 0 {
		return cal(x, enl(y))
	}
	if yt != 2 {
		panic("atx yt~I")
	} // todo dict
	r = mk(xt, yn)
	rp := r + 8
	w := i(C(xt))
	f := MT[xt+128].(func(i))
	for i := i(0); i < yn; i++ {
		yi := I(yp)
		if yi < xn {
			mv(rp, xp+w*yi, w)
		} else {
			f(rp)
		}
		rp += w
		yp += 4
	}
	if xt > 4 {
		rl(r)
	}
	if xt == 6 && yn == 1 {
		rx(I(r + 8))
		dx(r)
		r = I(r + 8)
	}
	return dxyr(x, y, r)
}
func cal(x, y i) (r i) {
	yt, yn, yp := v1(y)
	if yt != 6 {
		panic("type")
	}
	if x < 128 {
		if yn != 1 {
			panic("arity")
		}
		f := MT[x].(func(i) i)
		return f(fst(y))
	} else if x < 256 {
		if yn != 2 {
			panic("arity")
		}
		rl(y)
		dx(y)
		f := MT[x].(func(i, i) i)
		return f(I(yp), I(yp+4))
	}
	panic("nyi")
}
func cat(x, y i) (r i) {
	xt, yt, _, _, _, _ := v2(x, y)
	if xt == 0 || yt == 0 {
		trap()
	}
	if xt == yt {
		return ucat(x, y)
	}
	if xt == 6 {
		return lcat(x, y)
	}
	panic("nyi cat")
}
func ucat(x, y i) (r i) {
	xt, _, xn, yn, xp, yp := v2(x, y)
	if xt > 4 {
		rl(x)
	}
	if xt > 5 {
		r = mkd(x+8, x+12)
		return dxyr(x, y, r)
	}
	r = mk(xt, xn+yn)
	w := i(C(xt))
	mv(r+8, xp, w*xn)
	mv(r+8+w*xn, yp, w*yn)
	return dxyr(x, y, r)
}
func lcat(x, y i) (r i) { // list append
	x = use(x)
	xt, xn, xp := v1(x)
	if bk(xt, xn) < bk(xt, xn+1) {
		r = mk(xt, xn+1)
		mv(r+8, xp, 4*xn)
		dx(x)
		x, xp = r, r+8
	}
	sI(x, (xn+1)|6<<29)
	sI(xp+4*xn, y)
	return x
}
func enl(x i) (r i) { return lcat(mk(6, 0), x) }
func cnt(x i) (r i) {
	_, xn, _ := v1(x)
	dx(x)
	return mki(xn)
}
func typ(x i) (r i) {
	xt, _, _ := v1(x)
	r = mk(2, 1)
	sI(8+r, xt)
	return dxr(x, r)
}
func wer(x i) (r i) {
	xt, xn, xp := v1(x)
	if xt != 2 {
		panic("type")
	}
	n := i(0)
	for i := i(0); i < xn; i++ {
		n += I(xp + 4*i)
	}
	r = mk(2, n)
	rp := 8 + r
	for i := i(0); i < xn; i++ {
		nj := I(xp)
		for j := uint32(0); j < nj; j++ {
			sI(rp, i)
			rp += 4
		}
		xp += 4
	}
	return dxr(x, r)
}
func mtc(x, y i) (r i) { // x~y
	r = mk(2, 1)
	sI(r+8, match(x, y))
	return dxyr(x, y, r)
}
func match(x, y i) (r i) { // x~y
	if x == y {
		return 1
	}
	if I(x) != I(y) {
		return 0
	}
	xt, xn, xp := v1(x)
	yp, nn := y+8, i(0)
	switch xt {
	case 0:
		return 1 // todo
	case 1:
		nn = xn
	case 2:
		nn = xn << 2
	case 3:
		nn = xn << 3
	case 4:
		nn = xn << 4
	default:
		for i := i(0); i < xn; i++ {
			if match(I(xp), I(yp)) == 0 {
				return 0
			}
			xp += 4
			yp += 4
		}
		return 1
	}
	for i := i(0); i < nn; i++ {
		if C(xp+i) != C(yp+i) {
			return 0
		}
	}
	return 1
}
func not(x i) (r i) { return eql(mki(0), x) }
func cmp(x, y, eq i) (r i) {
	x, y = upx(x, y)
	x, y = ext(x, y)
	t, _, n, _, xp, yp := v2(x, y)
	cm := MT[t].(func(i, i) i)
	if eq == 1 {
		cm = MT[t+8].(func(i, i) i)
	}
	w := uint32(C(t))
	r = mk(2, n)
	rp := r + 8
	for i := i(0); i < n; i++ {
		sI(rp, cm(xp, yp))
		xp += w
		yp += w
		rp += 4
	}
	return dxyr(x, y, r)
}
func eql(x, y i) (r i) { return cmp(x, y, 1) }
func mor(x, y i) (r i) { return cmp(x, y, 0) }
func les(x, y i) (r i) { return cmp(y, x, 0) }
func fnd(x, y i) (r i) { // x?y
	xt, yt, _, yn, _, yp := v2(x, y)
	if xt != yt {
		trap()
	}
	r = mk(2, yn)
	rp := r + 8
	w := i(C(yt))
	for i := i(0); i < yn; i++ {
		sI(rp, fnx(x, yp))
		rp += 4
		yp += w
	}
	return dxyr(x, y, r)
}
func fnx(x, yp i) (r i) {
	xt, xn, xp := v1(x)
	eq := MT[8+xt].(func(i, i) i)
	w := uint32(C(xt))
	for i := i(0); i < xn; i++ {
		if eq(xp, yp) == 1 {
			return i
		}
		xp += w
	}
	return xn
}
func exc(x, y i) (r i) { // x^y
	_, yn, _ := v1(y)
	r = mk(2, 1)
	sI(r+8, yn)
	rx(x)
	return atx(x, wer(eql(r, fnd(y, x)))) // x@&xn=y?x
}
func grd(x i) (r i) { // <x
	xt, xn, xp := v1(x)
	r = seq(0, xn, 1)
	y := seq(0, xn, 1)
	msrt(y+8, r+8, 0, xn, xp, xt) // xt:1,2,3,4,5
	return dxyr(x, y, r)
}
func gdn(x i) (r i) { return rev(grd(x)) }           // >x
func srt(x i) (r i) { rx(x); return atx(x, grd(x)) } // ^x
func msrt(x, y, z, x3, x4, x5 i) { // merge sort
	if x3-z < 2 {
		return
	}
	c := (z + x3) / 2
	msrt(y, x, z, c, x4, x5)
	msrt(y, x, c, x3, x4, x5)
	mrge(x, y, z, x3, c, x4, x5)
}
func mrge(x, y, z, x3, x4, x5, x6 i) {
	k, j, a := z, x4, i(0)
	gt := MT[x6].(func(i, i) i)
	w := uint32(C(x6))
	for i := z; i < x3; i++ {
		if k >= x4 || (j < x3 && gt(x5+w*I(x+k<<2), x5+w*I(x+j<<2)) == 1) {
			a = j
			j++
		} else {
			a = k
			k++
		}
		sI(y+i<<2, I(x+a<<2))
	}
}
func str(x i) (r i)    { panic("nyi") }
func unq(x i) (r i)    { panic("nyi") }
func flr(x i) (r i)    { panic("nyi") }
func cst(x, y i) (r i) { panic("nyi") }
func min(x, y i) (r i) { panic("nyi") }
func max(x, y i) (r i) { panic("nyi") }
func nm(x, f i) (r i) {
	r = use(x)
	t, n, rp := v1(r)
	xp := x + 8
	w := uint32(C(t))
	g := MT[f+t].(func(i, i))
	for i := i(0); i < n; i++ {
		g(xp, rp)
		xp += w
		rp += w
	}
	if t == 4 && f == 19 { // +z
		return zre(r)
	}
	return r
}
func nd(x, y, f i) i {
	x, y = upx(x, y)
	x, y = ext(x, y)
	t, _, n, _, xp, yp := v2(x, y)
	w := uint32(C(t))
	g := MT[f+t].(func(i, i, i))
	r := mk(t, n)
	rp := r + 8
	for i := i(0); i < n; i++ {
		g(xp, yp, rp)
		xp += w
		yp += w
		rp += w
	}
	return dxyr(x, y, r)
}
func gtC(x, y i) i { return boolvar(C(x) > C(y)) }
func eqC(x, y i) i { return boolvar(C(x) == C(y)) }
func gtI(x, y i) i { return boolvar(int32(I(x)) > int32(I(y))) }
func eqI(x, y i) i { return boolvar(int32(I(x)) == int32(I(y))) }
func gtF(x, y i) i { return boolvar(F(x) > F(y)) }
func eqF(x, y i) i { return boolvar(F(x) == F(y)) }
func eqZ(x, y i) i { return boolvar(F(x) == F(y) && F(x+8) == F(y+8)) }
func gtL(x, y i) i {
	x, y = I(x), I(y)
	xt, yt, xn, yn, xp, yp := v2(x, y)
	if xt != yt {
		return boolvar(xt > yt)
	}
	n := xn
	if yn < xn {
		n = yn
	}
	gt := MT[xt].(func(i, i) i)
	w := uint32(C(xt))
	for i := i(0); i < n; i++ {
		a, b := xp+w*i, yp+w*i
		if gt(a, b) == 1 {
			return 1
		}
		if gt(b, a) == 1 {
			return 0
		}
	}
	return boolvar(xn > yn)
}
func eqL(x, y i) i  { return match(I(x), I(y)) }
func nag(r i)       { sI(r, 0) }
func nac(r i)       { sC(r, 32) }
func nai(r i)       { sI(r, naI) }
func naf(r i)       { sJ(r, naJ) }
func naz(r i)       { sJ(r, naJ); sJ(r+8, naJ) }
func nas(r i)       { sI(r, mk(1, 0)); sI(4+I(r), 0) }
func nal(r i)       { sI(r, 0) }
func adc(x, y, r i) { sC(r, C(x)+C(y)) }
func adi(x, y, r i) { sI(r, I(x)+I(y)) }
func adf(x, y, r i) { sF(r, F(x)+F(y)) }
func adz(x, y, r i) { sZ(r, Z(x)+Z(y)) }
func suc(x, y, r i) { sC(r, C(x)-C(y)) }
func sui(x, y, r i) { sI(r, I(x)-I(y)) }
func suf(x, y, r i) { sF(r, F(x)-F(y)) }
func suz(x, y, r i) { sZ(r, Z(x)-Z(y)) }
func muc(x, y, r i) { sC(r, C(x)*C(y)) }
func mui(x, y, r i) { sI(r, I(x)*I(y)) }
func muf(x, y, r i) { sF(r, F(x)*F(y)) }
func muz(x, y, r i) { sZ(r, Z(x)*Z(y)) }
func dic(x, y, r i) { sC(r, C(x)/C(y)) }
func dii(x, y, r i) { sI(r, I(x)/I(y)) }
func dif(x, y, r i) { sF(r, F(x)/F(y)) }
func diz(x, y, r i) { sZ(r, Z(x)/Z(y)) }
func add(x, y i) i  { return nd(x, y, 15+128) }
func sub(x, y i) i  { return nd(x, y, 19+128) }
func mul(x, y i) i  { return nd(x, y, 23+128) }
func diw(x, y i) i  { return nd(x, y, 27+128) }
func abs(x i) i     { return nm(x, 15) }
func neg(x i) i     { return nm(x, 19) }
func sqr(x i) i     { return nm(x, 27) }
func abc(x, r i) { // +c (toupper)
	fmt.Println("abc / toupper")
	if c := C(x); craz(c) {
		sC(r, c-32)
	} else {
		sC(r, c)
	}
}
func abi(x, r i) {
	if c := int32(I(x)); c < 0 {
		sI(r, i(-c))
	} else {
		sI(r, i(c))
	}
}
func abf(x, r i) { sF(r, math.Abs(F(x))) }
func abz(x, r i) { sF(r, cmplx.Abs(Z(x))) }
func nec(x, r i) { // -c (tolower)
	if c := C(x); crAZ(c) {
		sC(r, c+32)
	} else {
		sC(r, c)
	}
}
func nei(x, r i) { sI(r, i(-int32(I(x)))) }
func nef(x, r i) { sF(r, -F(x)) }
func nez(x, r i) { sZ(r, -Z(x)) }
func sqc(x, r i) { panic("%c") } // %c ?
func sqi(x, r i) { panic("%i") } // %i ?
func sqf(x, r i) { sF(r, math.Sqrt(F(x))) }
func sqz(x, r i) { sZ(r, cmplx.Conj(Z(x))) } // %z complex conjugate
func zri(x i, o i) (r i) {
	t, n, xp := v1(x)
	if t != 4 {
		panic("type")
	}
	r = mk(3, n)
	rp := r + 8
	xp += o
	for i := i(0); i < n; i++ {
		sF(rp, F(xp))
		rp += 8
		xp += 16
	}
	return dxr(x, r)
}
func zre(x i) (r i) { return zri(x, 0) }
func zim(x i) (r i) { return zri(x, 8) }

func val(x i) (r i) {
	xt, _, _ := v1(x)
	switch xt {
	case 6:
		return evl(x)
	case 7:
		rx(x + 12)
		return dxr(x, x+12)
	default:
		fmt.Printf("val xt=%d\n", xt)
		panic("nyi")
	}
}
func evl(x i) (r i) {
	xt, xn, xp := v1(x)
	if xt != 6 || xn == 0 {
		return x
	} else if xn == 1 {
		return fst(x)
	}
	rl(x)
	r = mk(6, xn)
	rp := r + 8
	for i := i(0); i < xn; i++ {
		sI(rp, evl(I(xp)))
		rp += 4
		xp += 4
	}
	rp = r + 8
	dx(x)
	if xn == 2 {
		rl(r)
		return dxr(r, atx(I(rp), I(rp+4)))
	} else if xn == 3 {
		rx(I(rp))
		return cal(I(rp), drop(r, 1))
	} else {
		panic("args")
	}
}

func craz(x c) bool {
	if x < 'a' || x > 'z' {
		return false
	}
	return true
}
func crAZ(x c) bool {
	if x < 'A' || x > 'Z' {
		return false
	}
	return true
}
func boolvar(b bool) i {
	if b {
		return 1
	}
	return 0
}
func trap() { panic("trap") }
func dump(a, n i) i {
	p := a >> 2
	fmt.Printf("%.8x ", a)
	for i := i(0); i < n; i++ {
		x := MI[p+i]
		fmt.Printf(" %.8x", x)
		if i > 0 && (i+1)%8 == 0 {
			fmt.Printf("\n%.8x ", a+4*i+4)
		} else if i > 0 && (i+1)%4 == 0 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
	return 0
}
func C(a i) c              { return MC[a] } // global get, e.g. I i
func I(a i) i              { return MI[a>>2] }
func J(a i) j              { return MJ[a>>3] }
func F(a i) f              { return MF[a>>3] }
func Z(a i) complex128     { return complex(MF[a>>3], MF[1+a>>3]) }
func sC(a i, v c)          { MC[a] = v } // global set, e.g. i::v
func sI(a i, v i)          { MI[a>>2] = v }
func sJ(a i, v j)          { MJ[a>>3] = v }
func sF(a i, v f)          { MF[a>>3] = v }
func sZ(a i, v complex128) { MF[a>>3] = real(v); MF[1+a>>3] = imag(v) }
func atoi(s string) i {
	if x, e := strconv.Atoi(s); e == nil {
		return i(x)
	}
	panic("atoi")
}
func mark() { // mark bucket type within free blocks
	for t := i(4); t < 32; t++ {
		for p := MI[t] >> 2; p != 0; p = MI[p] >> 2 {
			MI[1+p] = 0
			MI[2+p] = t
		}
	}
}
func leak() {
	mark()
	//dump(0, 200)
	p := i(64)
	for p < i(len(MI)) {
		if MI[p+1] != 0 {
			panic(fmt.Errorf("non-free block at %d(%x)", p<<2, p<<2))
		}
		t := MI[p+2]
		if t < 4 || t > 31 {
			panic(fmt.Errorf("illegal bucket type %d at %d(%x)", t, p<<2, p<<2))
		}
		dp := i(1) << t
		p += dp >> 2
	}
}
func kst(x i) s {
	t, n, _ := v1(x)
	var f func(i i) s
	var tof func(s) s = func(s s) s { return s }
	istr := func(i i) s {
		if n := int32(MI[i+2+x>>2]); n == -2147483648 {
			return "0N"
		} else {
			return strconv.Itoa(int(n))
		}
	}
	fstr := func(i i) s {
		if f := MF[i+1+x>>3]; math.IsNaN(f) {
			return "0n"
		} else {
			return strconv.FormatFloat(f, 'g', -1, 64)
		}
	}
	zstr := func(i i) s {
		if z := Z(x + 8 + 16*i); cmplx.IsNaN(z) {
			return "0ni0n"
		} else {
			return strconv.FormatFloat(real(z), 'g', -1, 64) + "i" + strconv.FormatFloat(imag(z), 'g', -1, 64)
		}
	}
	sstr := func(i i) s {
		r := I(x + 8 + 4*i)
		rn := I(r) & 536870911
		return string(MC[r+8 : r+8+rn])
	}
	sep := " "
	switch t {
	case 0:
		fc := []byte(":+-*%&|<>=!~,^#_$?@.")
		if x < 128 && bytes.Index(fc, []byte{byte(x)}) != -1 {
			return string(byte(x)) + ":"
		} else if x < 256 && bytes.Index(fc, []byte{byte(x - 128)}) != -1 {
			return string(byte(x - 128))
		} else {
			panic(fmt.Errorf("nyi: kst func %x\n", x))
		}
	case 1:
		return `"` + string(MC[x+8:x+8+n]) + `"`
	case 2:
		f = istr
	case 3:
		f = fstr
		tof = func(s s) s {
			if strings.Index(s, ".") == -1 {
				return s + "f"
			}
			return s
		}
	case 4:
		f = zstr
	case 5:
		f = sstr
		sep = "`"
		tof = func(s s) s { return "`" + s }
	case 6:
		f = func(i i) s { return kst(MI[2+i+x>>2]) }
		sep = ";"
		tof = func(s s) s { return "(" + s + ")" }
	case 7:
		return kst(x+8) + "!" + kst(x+12)
	default:
		panic(fmt.Sprintf("nyi: kst: t=%d", t))
	}
	r := make([]s, n)
	for k := range r {
		r[k] = f(i(k))
	}
	return tof(strings.Join(r, sep))
}

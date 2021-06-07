package main

import (
	"fmt"

	"github.com/ktye/wg/wasi_unstable"

	. "github.com/ktye/wg/module"
)

//func LK(x K) K         { return K(wasi_unstable.L64(uint64(x))) }
//func LI(x int32) int32 { return wasi_unstable.L32(x) }
func main() { // _start
	kinit()
	write(Ku(117851310093419)) // "ktye/k"
	for {
		write(Ku(8202)) // "\n "
		x := read()
		repl(x)
	}
}
func repl(x K) {
	m := mcount()
	n := nn(x)
	xp := int32(x)
	if n > 0 && I8(xp) == 92 { // \
		if n == 1 {
			help()
		} else if I8(1+xp) == 92 { // \\
			wasi_unstable.Proc_exit(0)
		} else if I8(1+xp) == 116 { // \t
			bench(ndrop(2, x))
		}
		return
	}

	x = val(x)
	if x != 0 {
		dx(Out(x))
	}
	fmt.Println("diff", int32(m)-int32(mcount()))
}
func bench(x K) {
	i := fndc(x, 32)
	if i < 0 {
		trap(Parse)
	}
	n := maxi(1, int32(prs(it, ntake(i, rx(x)))))
	x = parse(ndrop(i, x))
	t := time()
	for n > 0 {
		dx(exec(rx(x)))
		n--
		continue
	}
	t = time() - t
	dx(Out(Kf(float64(t))))
}

func Out(x K) K {
	write(cat1(Kst(rx(x)), Kc(10)))
	return x
}
func Otu(x, y K) K {
	write(cat1(Kst(x), Kc(':')))
	return Out(y)
}
func time() int64 {
	wasi_unstable.Clock_time_get(1, wasi_unstable.Timestamp(0), 512)
	return I64(512)
}
func read() (r K) {
	r = mk(Ct, 504)
	rp := int32(r)
	SetI32(512, rp)
	SetI32(516, 504)
	if wasi_unstable.Fd_read(0, 512, 1, 512) != 0 {
		trap(Io)
	}
	return ntake(maxi(0, I32(512)-1), r)
}
func write(x K) {
	SetI32(512, int32(x))
	SetI32(516, nn(x))
	if wasi_unstable.Fd_write(1, 512, 1, 512) != 0 {
		trap(Io)
	}
	dx(x)
}
func help() {
	trap(Nyi)
}

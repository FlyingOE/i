package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/wasm"
)

const trace = false

func TestWagon(t *testing.T) {
	if broken {
		t.Skip()
	}
	b, e := ioutil.ReadFile("t")
	if e != nil {
		t.Fatal(e)
	}
	mb, e := ioutil.ReadFile("../../k.w")
	if e != nil {
		t.Fatal(e)
	}
	v := strings.Split(strings.TrimSpace(string(b)), "\n")
	for i := range v {
		if len(v[i]) == 0 {
			fmt.Println("skip")
			return
		} else if v[i][0] == '/' {
			continue
		}
		vv := strings.Split(v[i], " /")
		if len(vv) != 2 {
			panic("test file")
		}
		in := strings.TrimRight(vv[0], " \t\r")
		exp := strings.TrimSpace(vv[1])
		m, tab, data := run(bytes.NewReader(mb))
		if e = runWagon(tab, m.wasm(tab, data), in, exp); e != nil {
			t.Fatalf("%d: %s", i+1, e)
		}
	}
}

type K struct {
	m  *wasm.Module
	vm *exec.VM
}

func runWagon(tab []segment, b []byte, s string, exp string) error {
	fmt.Printf("%s / ", s)
	m, e := wasm.ReadModule(bytes.NewReader(b), hostFuncs)
	if e != nil {
		return e
	}
	//if e := validate.VerifyModule(m); e != nil { // fails with hostFuncs
	//	return e
	//}
	vm, e := exec.NewVM(m)
	if e != nil {
		return e
	}
	if trace {
		fmt.Println("memory", len(vm.Memory()))
	}
	K := K{m: m, vm: vm}
	K.call("ini", 16)
	//r := K.call("prs", K.mks(s))
	r := K.call("val", K.mks(s))
	//r := K.call("evl", K.parseVector(s), 0)
	got := K.kst(r)
	fmt.Println(got)
	if got != exp {
		return fmt.Errorf("expected/got:\n%s\n%s", exp, got)
	}
	// free result and check for memory leaks
	mem := K.vm.Memory()
	K.call("dx", r)
	K.call("dx", get(mem, 132)) // kkey
	K.call("dx", get(mem, 136)) // kval
	K.call("dx", get(mem, 148)) // xyz
	if e := leak(mem); e != nil {
		return e
	}
	return nil
}
func hostFuncs(name string) (*wasm.Module, error) { // imported as module "ext"
	sin := func(proc *exec.Process, x float64) float64 { return math.Sin(x) }
	cos := func(proc *exec.Process, x float64) float64 { return math.Cos(x) }
	atan2 := func(proc *exec.Process, x, y float64) float64 { return math.Atan2(x, y) }

	m := wasm.NewModule()
	m.Types = &wasm.SectionTypes{
		Entries: []wasm.FunctionSig{
			{Form: 0, ParamTypes: []wasm.ValueType{wasm.ValueTypeF64}, ReturnTypes: []wasm.ValueType{wasm.ValueTypeF64}},
			{Form: 0, ParamTypes: []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, ReturnTypes: []wasm.ValueType{wasm.ValueTypeF64}},
		},
	}
	m.FunctionIndexSpace = []wasm.Function{
		{Sig: &m.Types.Entries[0], Host: reflect.ValueOf(sin), Body: &wasm.FunctionBody{}},
		{Sig: &m.Types.Entries[0], Host: reflect.ValueOf(cos), Body: &wasm.FunctionBody{}},
		{Sig: &m.Types.Entries[1], Host: reflect.ValueOf(atan2), Body: &wasm.FunctionBody{}},
	}
	m.Export = &wasm.SectionExports{
		Entries: map[string]wasm.ExportEntry{
			"sin":   {FieldStr: "sin", Kind: wasm.ExternalFunction, Index: 0},
			"cos":   {FieldStr: "cos", Kind: wasm.ExternalFunction, Index: 1},
			"atan2": {FieldStr: "atan2", Kind: wasm.ExternalFunction, Index: 2},
		},
	}
	return m, nil
}
func (K *K) call(s string, argv ...uint32) uint32 {
	m, vm := K.m, K.vm
	x, ok := m.Export.Entries[s]
	if !ok {
		panic(fmt.Errorf("function does not exist: %s", s))
	}
	fidx := m.Function.Types[x.Index]
	ftyp := m.Types.Entries[fidx]
	n := len(ftyp.ParamTypes)
	var e error
	var res interface{}
	if n != len(argv) {
		panic(fmt.Errorf("%s expects %d arguments (got %d)", s, n, len(argv)))
	}
	if n == 1 {
		res, e = vm.ExecCode(int64(x.Index), uint64(argv[0]))
	} else if n == 2 {
		res, e = vm.ExecCode(int64(x.Index), uint64(argv[0]), uint64(argv[1]))
	} else if n == 3 {
		res, e = vm.ExecCode(int64(x.Index), uint64(argv[0]), uint64(argv[1]), uint64(argv[2]))
	} else {
		panic(fmt.Errorf("%s expects %d arguments", s, n))
	}
	if e != nil {
		panic(e)
	}
	switch v := res.(type) {
	case nil:
		return 0
	case uint32:
		return v
	default:
		panic(fmt.Errorf("%s returns %T", s, res))
	}
}
func (K *K) mk(t, n uint32) uint32 { return K.call("mk", t, n) }
func (K *K) mks(s string) (r uint32) {
	m := K.vm.Memory()
	r = K.mk(1, uint32(len(s)))
	copy(m[8+r:], s)
	return r
}
func (K *K) parseVector(s string) uint32 {
	m := K.vm.Memory()
	fc := ":+-*%&|<>=!~,^#_$?@.+'/\\"
	if len(s) == 0 {
		return 0
	}
	if len(s) > 1 && s[1] == ':' && strings.Index(fc, s[:1]) != -1 {
		return uint32(s[0]) + 128
	}
	if len(s) > 0 && strings.Index(fc, s[:1]) != -1 {
		return uint32(s[0])
	}
	if len(s) > 0 && s[0] == '"' {
		s = strings.Trim(s, `"`)
		b := []c(s)
		p := K.call("mk", 1, uint32(len(b)))
		for i := 0; i < len(b); i++ {
			m[8+int(p)+i] = b[i]
		}
		return p
	}
	if len(s) > 0 && s[0] == '`' {
		v := strings.Split(s[1:], "`")
		sn := uint32(len(v))
		sv := K.call("mk", 5, sn)
		for i := uint32(0); i < sn; i++ {
			b := v[i]
			rn := uint32(len(b))
			r := K.call("mk", 1, rn)
			for k := uint32(0); k < rn; k++ {
				m[8+r+k] = b[k]
			}
			binary.LittleEndian.PutUint32(m[sv+8+4*i:], uint32(r)) // sI(sv+8+4*i, r)
		}
		return sv
	}
	if len(s) > 0 && s[0] == '(' {
		return K.parseList(s[1:])
	}
	f := strings.Index(s, ".")
	v := strings.Split(s, " ")
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
			panic(fmt.Errorf("parse: %s", s))
		}
	}
	if f == -1 {
		x := K.call("mk", 2, n)
		for i := uint32(0); i < n; i++ {
			binary.LittleEndian.PutUint32(m[x+8+i*4:], uint32(iv[i]))
		}
		return x
	} else {
		x := K.call("mk", 3, n)
		for i := uint32(0); i < n; i++ {
			binary.LittleEndian.PutUint64(m[x+8+i*8:], math.Float64bits(fv[i]))
		}
		return x
	}
}
func (K *K) parseList(s string) uint32 {
	if len(s) == 0 || s[len(s)-1] != ')' {
		panic("parse list")
	} else if len(s) == 1 {
		return K.call("mk", 6, 0)
	}
	r := make([]uint32, 0)
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
			r = append(r, K.parseVector(s[a:i]))
			a = i + 1
		}
	}
	r = append(r, K.parseVector(s[a:]))
	x := K.call("mk", 6, uint32(len(r)))
	m := K.vm.Memory()
	for k := range r {
		binary.LittleEndian.PutUint32(m[8+x+4*uint32(k):], r[k])
	}
	return x
}
func u64(v interface{}) uint64 {
	switch x := v.(type) {
	case uint32:
		return uint64(x)
	case uint64:
		return x
	case float64:
		return math.Float64bits(x)
	default:
		panic(x)
	}
}
func (K *K) dump(a, n k) { dump(K.vm.Memory(), a, n) }
func dump(m []byte, a, n k) {
	fmt.Printf("%.8x ", a)
	for i := k(0); i < n; i++ {
		p := a + 4*i
		x := get(m, p)
		fmt.Printf(" %.8x", x)
		if i > 0 && (i+1)%8 == 0 {
			fmt.Printf("\n%.8x ", p+4)
		} else if i > 0 && (i+1)%4 == 0 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}

type k = uint32

func (K *K) kst(a k) s {
	if a == 0 {
		return ""
	}
	m := K.vm.Memory()
	x, t, n := k(0), k(0), k(0)
	if a > 256 {
		x = get(m, a)
		t, n = x>>29, x&536870911
	}
	var f func(i int) s
	var tof func(s) s = func(s s) s { return s }
	istr := func(i int) s {
		n := int32(get(m, 8+4*k(i)+a))
		return strconv.Itoa(int(n))
	}
	fstr := func(i int) s {
		if f := getf(m, a+8+8*k(i)); math.IsNaN(f) {
			return "0n"
		} else {
			return strconv.FormatFloat(f, 'g', -1, 64)
		}
	}
	zstr := func(i int) s {
		if z := complex(getf(m, x+8+16*k(i)), getf(m, x+16+16*k(i))); cmplx.IsNaN(z) {
			return "0ni0n"
		} else {
			return strconv.FormatFloat(real(z), 'g', -1, 64) + "i" + strconv.FormatFloat(imag(z), 'g', -1, 64)
		}
	}
	sstr := func(i int) s {
		r := get(m, a+8+4*k(i))
		rn := get(m, r) & 536870911
		return string(m[r+8 : r+8+rn])
	}
	sep := " "
	switch t {
	case 0:
		if a < 128 {
			return string([]byte{byte(a)})
		} else if a < 256 {
			return string([]byte{byte(a) - 128}) + ":"
		} else if n == 4 {
			return sstr(0)
		} else {
			fmt.Printf("x=%x a=%x\n", x, a)
			dump(m, 0, 200)
			panic("kst t=0 nyi")
		}
	case 1:
		return `"` + string(m[a+8:a+8+n]) + `"`
	case 2:
		f = istr
	case 3:
		f = fstr
		tof = func(s s) s {
			if strings.Index(s, ".") == -1 {
				return s + ".0"
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
		f = func(i int) s { return K.kst(get(m, 8+4*uint32(i)+a)) }
		sep = ";"
		tof = func(s s) s { return "(" + s + ")" }
	case 7:
		return K.kst(x+8) + "!" + K.kst(x+12)
	default:
		K.dump(0, 200)
		panic(fmt.Sprintf("nyi: kst: t=%d a=%x x=%x", t, a, x))
	}
	r := make([]s, n)
	for i := range r {
		r[i] = f(i)
	}
	return tof(strings.Join(r, sep))
}
func get(m []byte, a k) k        { return binary.LittleEndian.Uint32(m[a:]) }
func getf(m []byte, a k) float64 { return math.Float64frombits(binary.LittleEndian.Uint64(m[a:])) }
func mark(m []byte) { // mark bucket type within free blocks
	for t := k(4); t < 32; t++ {
		p := get(m, 4*t) // free pointer of type t
		for p != 0 {
			m[4+p] = 0
			m[8+p] = c(t)
			p = get(m, p) // pointer to next free
		}
	}
}
func leak(m []byte) error {
	//dump(m, 0, 200)
	mark(m)
	p := k(256) // first data block
	for p < k(len(m)) {
		// a free block has refcount==0 at p+4 and bucket type at p+8 (after marking)
		if get(m, p+4) != 0 {
			dump(m, 0, 200)
			return fmt.Errorf("non-free block at %d(%x)", p, p)
		}
		t := m[p+8]
		if t < 4 || t > 31 {
			return fmt.Errorf("illegal bucket type %d at %d(%x)", t, p, p)
		}
		dp := 1 << t
		p += k(dp)
	}
	return nil
}

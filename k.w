ini:I:I{0::1130366807310592j;128::x;p:256;i:8;(i<x)?/((4*i)::p;p*:2;i+:1);x} /x:16(64k)
bk:I:II{r:32-*7+y*I?C x;(r<4)? :4;r}
mk:I:II{t:x bk y;i:4*t;(~I i)?/i+:4;(128~i)?!;a:I i;i::I a;j:i-4;(j>=4*t)?/(u:a+1<<j>>2;u::I j;j::u;j-:4);a::y|x<<29;(a+4)::1;a}
mki:I:I{r:2 mk 1;(r+8)::x;r}mkd:I:II{v2;(xt~5)?!;(xn~yn)?!;r:7 mk 2;(r+8)::x;(r+12)::y;r}
v1:{xt:(I x)>>29;xn:(I x)&536870911;xp:8+x}v2:{v1;yt:(I y)>>29;yn:(I y)&536870911;yp:8+y}
fr:0:I{v1;t:4*xt bk xn;x::I t;t::x}dx:0:I{(x>255)?(xr:I x+4;(x+4)::xr-1;(1~xr)?(v1;(xt>5)?xn/(dx I xp+4*i);fr x))}dxr:{dx x;r}dxyr:{dx x;dx y;r}rx:0:I{(x>255)?(x+:4;x::1+I x)}rl:0:I{v1;x+:8;xn/(rx x;x+:4)}
til:I:I{v1;(~2~xt)?!;n:I xp;dx x;(n<'0)? :tir -n;r:xt mk n;rp:8+r;n/(rp::i;rp+:4);r}tir:I:I{r:2 mk x;rp:4+r+4*x;x/(rp::i;rp-:4);r}
rev:I:I{v1;(~xn)? :x;x atx tir xn}
fst:I:I{v1;(xt~7)?(rx 12+x;dx x; :fst 12+x);x atx mki 0}
atx:I:II{v2;(xt~7)?!;(~yt~2)?!;r:xt mk yn;rp:r+8;xt?[!;atC;atI;atF;!];dxyr}
atC:{(yn/((rp+i)::C?32;yi:I yp;(yi<xn)?(rp+i)::C xp+yi;yp+:4))}
atI:{(yn/(rp::naI;yi:I yp;(yi<xn)?rp::I xp+4*yi;rp+:4;yp+:4))}naI:{-2147483648}
atF:{(yn/(rp::naF;yi:I yp;(yi<xn)?rp::F xp+8*yi;rp+:8;yp+:4))}naF:{9221120237041090561f}
rsh:I:II{v2;(~xt~2)? :x;(xn~1)? :x tak y;y} /nyi rsh
tak:I:II{v2;n:I xp;r:til x;rp:r+8;(yn<n)?(n/(rp+4*i)::i\yn);y atx r}
use:{(~1~I v+4)?(r:xt mk xn;rp+r+8;mv(rp;xp;xn*I?C xt);dx x;x:r;xp:x+8)} mv:0:III{(7+z%8)/(x::J y;x+:8;y+:8)}
lcat:I:II{v1;((xt bk xn)<(xt bk xn+1))?(r:xt mk xn+1;mv(r+8;xp;4*xn);dx x;x:r;xp:x+8);(xp+4*xn)::y;x::(xn+1)|6<<29;x}
enl:I:I{(6 mk 0) lcat x}

\
/atL:{(yn/(rp::0;(i<xn)?(rp::I xp;rx I rp;)rp+:4;xp+:4))}

\
lrc:I:II{!;1}drc:I:II{!;1} /nyi
til:I:I{v1;(~2~xt)?!;r:xt mk n:I xp;rp:8+r;n/(rp::i;rp+:4);dxr}
fst:I:I{v1;(~xt)? :x;(7~xt)? :x;r:xt mk 1;xt?[;(r+8)::C xp;;(r+8)::J xp;((r+8)::J xp;(r+16)::J xp+8);;;!;(r+8)::I xp];(xt~6)?xn lnc x+8;dxr}
rev:I:I{v1;xt?[!;;;;;;r:((inc xp)mkd inc xp+4);r:xt mk xn];xt?[!;(rp:r+7+xn;xn/((rp-i)::C xp+i));;(rp:r+8*xn;xn/(rp::J xp;rp-:8;xp+:8));(rp:r-8+16*xn;xn/(rp::J xp;(8+rp)::J 1+xp;rp-:16;xp+:16));;((r+8)::rev r+8;(r+12)::rev r+12);;;(rp:r+4+4*xn;xn/(rp::I xp;xp+:4;rp-:4))];dxr}
cnt:I:I{v1;decr x;mki xn}
tip:I:I{v1;r:2 mk 2;(8+r)::xt;(12+r)::xn;dxr}
sumi:I:II{y/r+:I x+1;r}
wer:I:I{v1;(~2~xt)?!;xn/(0>'I xp+i)?!;r:2 mk xn sumi 8+xp;rp:8+r;xn/(I xp)/(rp::i;rp+:4);dxr}
enl:I:I{v1;r:1 mk 6;(8+r)::x;r}
neg:I:I{v1;xt?[!;!;;;;!; :45 lrc x; :45 drc x;r:xt mk xn];decr x;(2~xt)?(rp:r+8;xn/(rp::0-I xp;rp+:4;xp+:4); :r);(4~xt)?(xt:3;xn*:2);rp:8+r;xn/(rp::-F xp;rp+:4;xp+:4);r}
unq:I:I{v1;xt?[!;;;;;;;!;r:xt mk 0];rn:0;xn/(~rn~r fnd1 xi:x ati i)?(rn+:1;r:r cat xi);(xt~6)?lnc r;dxr}
ati:I:I{v1;r:xt mk 1;rp:8+r;xt?[!;rp::C xp;;rp::J xp;(rp::J xp;(8+rp):J xp+8);;!;rp::I xp];(6~xt)?(rp::inc rp);r}
fnd:I:II{v2;r:2 mk yn];rp:r+8;yn/(yi:y ati i;rp::x fnd1 yi);dxyr}fnd1:I:II{v2;decr y;r:xn; xn/(xi:x ati i)mc y}
mc:I:II{(x~y)? :0; xt?[ :1; (xn/((~(C xp)~(C yp) :0);xp+:1;yp+:1); :1); ;xn*:2;xn*:4; ; :x mtchl y; :(xp mtl yp)&&(4+xp)mtl 4+yp) ]; xn/((~(I xp)~(I yp) :0);xp+:4;yp+:4); :1}
mtc:I:II{v2;r:2 mk 1;rp::x mc y;dxyr}mtl:I:II{v2;(xn/(xi:x ati;yi y ati i;r:(xi mc yi);(~r)?dxyr);dxyr)}

r8:{rp:r+8}
add:I:II{v2; ..conform: r:xn mk xt; r8; xt?[adddC;adddI;adddF;adddZ!]; dxyr }
addd:T{ nd1+nd2 }subbb:T{ nd1-nd2 }mulll:T{ nd1*nd2 }
nd1:{vx:xn>1;vy:yn>1; xn/(rp,i)::(T xp,i*vx) }nd2:{ T yp,i*vy; dxyr }  /vx vy: 0(atom) 1(vector)

p,i → (addr,offset,width)
e.g. (p,i)::I 144,0

01234567   xt:x>>29       xn:x&536870911 (-1+1<<29)
Fcifzsld   xt~0(function) x<256(basic) 
0148x440   ft~xn&0xff00  (derived, proj, lambda, native)  composition==lambda?
	   fn~xn&0xff    (argn)

+ add abs                 abs:+z         memory
- sub neg                                0..  7   type sizes   0 1 4 8 16 4 4 0
* mul fst                                8.. 11   k-tree/key   pointer     todo
% div sqr                               12.. 15   k-tree/value pointer     todo
& min wer                               16..127   free pointers (4*i) for bt i, i:4..31
| max rev                              128..131   memsize log2
< les gup                 
> mor gdn                 
= eql grp                 
~ mtc not   mc mtl
! key til                 z:re!im  im:!z
, cat enl                 
^ exc asc                 
$ str cst   
# tak cnt
_ drp flr                 ang:_z
? fnd unq   fnd1              
@ atx typ                 z:abs@ang  z@ang
. cal val                 re:. z

package main

import . "github.com/ktye/wg/module"

func zk() {
	Data(600, "/               64     72    80    88     96     104    112    120                          \n``x`y`z`k`l`a`b`while`\"rf.\"`\"rz.\"`\"uqs.\"`\"uqf.\"`\"gdt.\"`\"lin.\"`\"odo.\"\n\n`\"x.\":{,/+\"0123456789abcdef\"@(x%16;16/x:256/256+x)} /`x@ (hex)\n`\"t.\":`45         /`t@ token\n`\"p.\":`46         /`p@ parse\n`\"b.\":(`46)[`b;]  /`b@ reinterpret\n`\"c.\":(`46)[`c;]\n`\"i.\":(`46)[`i;]\n`\"s.\":(`46)[`s;]\n`\"f.\":(`46)[`f;]\n`\"z.\":(`46)[`z;]\n\n`\"l.\":{t:@x;k:`kxy 1\n kt:{x:$[`T~@x;T x;`pad(\"\";\"-\"),$x];(x,'\"|\"),'T y}\n d:{r:!x;x:.x;$[`T~@x;kt[r;x];,'[,'[`pad(k'r);\"|\"];k'x]]}\n T:{$[`L':@'.x;,k x;(,*x),(,(#*x)#\"-\"),1_x:\" \"/:'+`pad@'$(!x),'.x]}\n dd:(\"\";,\"..\")20<#x\n x:$[x~*x;x;(20&#x)#x]\n $[`L~t;k'x;`D~t;d x;`T~t;T x;,k x],dd} /352\n\n`\"pad.\":{(|/#'x)#'x}\n`\"split.\":{$[`L~@x;`split@'x;\" \"\\:$[\" \"=x@-1+#x:x@&~i&~':i:\" \"=x;-1_x;x]]} /multiple whitespace/leading/trailing\n`\"edit.\":{et:{\"+\",(`k@!x),\"!\",el@.x}\n el:{t:(`S`B`C`I`F`Z!``b`c`i`f`z)@@'x; (`k@t),\"$'+`split@'\\\"\\\\n\\\"\\\\:-1_1_\\\"\\n\",(\"\\n\"/:\" \"/:'+`pad@'$x),\"\\n\\\"\"}\n $[`T~t:@x;et x;`L~t;el x;`C~t;\"-1_1_\\\"\\n\",x,\"\\n\\\"\";\"*\",el@,x]}\n\nl:(`a`b`c;1 22 3;1.5 6 7)\nt:+`alpha`beta`gamma!l\n\n`\"kxy.\":{t:@y;n:#y;k:`kxy x;m:x\n q:{c,(\"\\\\\"/:(0,i)^@[x;i;(qs!\"tnr\\\"\\\\\")x i:&x':qs:\"\\t\\n\\r\\\"\\\\\"]),c:_34}\n s:{$[|/x':\"\\t\\n\\r\"__!31;\"0x\",`x@x;q x]}\n a:{x:$x;$[`c~t;s x;`s~t;\"`\",x;x]}\n d:{r:\"!\",k@.x;n:#!x;x:k@!x;$[(1~n)|(@.x)':`D`T;\"(\",x,\")\";x],r}\n v:{m*:(.`\".kstm\")t;  dd:(\"\";\"..\")m<#x;x:(m&#x)#x\n  x:$[`L~t;k'x;`C~t;x;$x]\n  x:$[`B~t;(*'x),\"b\";`C~t;s x;`S~t;c,(c:\"`\")/:x;`L~t;$[1~n;*x;\"(\",(\";\"/:x),\")\"];\" \"/:x]\n  ((\"\";\",\")(1~n)),x,dd}\n $[n~0;(.`\".kst0\")@t;`T~t;\"+\",d@+y;`D~t;d y;y~*y;a y;v y]} /344\n \n`\"k.\":`kxy 1000000\n\n`\"uqs.\":{x@&~0b~':x:^x} \n`\"uqf.\":{x@&(!#x)=x?x}\n`\"gdt.\":{[t;g](!#t){x g y x}/|.t}\n\n`\".kst0\":`B`C`I`S`F`Z`L!(\"0#0b\";c,c:_34;\"!0\";\"0#`\";\"0#0.\";\"0#0a\";\"()\")\n`\".kstm\":`B`C`I`S`F`Z`L!100 100 30 30 20 10 20\n\n`\"lin.\":{$[`L~@z;(.`\"lin.\")[x;y]'z;[dx:0.+1_-':x;dy:0.+1_-':y;b:(-2+#x)&0|x'z;(y b)+(dy b)*(z-x b)%dx b]]}\n`\"odo.\":{{x/y}'[x;(!*/x)%/:{(*/x)%\\ x}x]}\n\ndot:{dotmv:{{+/x*y}\\:[x;y]};dotmv[x;y]}\n\nsolve:{qslv:{H:x 0;r:x 1;n:x 2;m:x 3;j:0;K:!m\n while[j<n;y[K]-:(+/(conj H[j;K])*y K)*H[j;K];K:1_K;j+:1]\n i:n-1;J:!n;y[i]%:r@i\n while[i;j:i_J;i-:1;y[i]:(y[i]-+/H[j;i]*y@j)%r@i]\n n#y}\n q:$[`i~@*|x;x;qr x];$[`L~@y;qslv/:[q;y];qslv[q;y]]}\n\nqr:{K:!m:#*x;I:!n:#x;j:0;r:n#0a;turn:$[`Z~@*x;{(-x)@angle y};{x*1. -1@y>0}]\n while[j<n;I:1_I\n  r[j]:turn[s:abs/j_x j;xx:x[j;j]]\n  x[j;j]-:r[j]\n  x[j;K]%:%s*(s+abs xx)\n  x[I;K]-:{+/x*y}/:[(conj x[j;K]);x[I;K]]*\\:x[j;K]\n  K:1_K;j+:1];(x;r;n;m)}\n\nany:`30;abs:`32;sin:`44;cos:`39;find:`31;fill:`38;imag:`33;conj:`34;angle:`35;exp:`42;log:`43\n\nej:{(y j),'x_z i j:&~0N=i:(z x)?y x} /sym t1 t2\navg:{(+/x)%0.+#x}\nvar:{(+/x*x:(x-avg x))%-1+#x}\nstd:{%var x}\n\n`\"rf.\": {.5+(x?0)%4294967295.}\n`\"rf1.\":{.5+(1.+x?0)%4294967295.}\n`\"rz.\": {(%-2*log `rf1 x)@360.*`rf x}\n`\".html\":{$[`L~@x;\"<div style='display:flex;flex-direction:column'>\",(,/(.`\".html\")'x),\"</div>\";`S~@x;\"<div style='display:flex;flex-direction:row'>\",(,/(.`\".html\")'x),\"</div>\";~`s~@x;\"\";{{\"<\",y,\" id='\",x,\"'></\",y,\">\"}[x;$[y':`i`c`f`z`s`I`B`C`F`Z`S;\"input\";`b~y;\"input type='checkbox'\";`T~y;\"table\";`l~t;\"button\";\"pre\"]]}[$x;@.x]]}\n\n`\"pack.\":{w:{(`c@,#x),x};($t),$[`s~t:@x;`pack@$x;x~*x;w `c@,x;`L~@x;(`c@,#x),,/`pack@'x;(@x)':`D`T;(`pack@.x),`pack@!x;`S~t;,/`pack@$x;w `c x]}\n`\"unpack.\":{s:x;g:{[n]r:n#s;s::n_s;r};n:{*`i@g 4};u:{x;$[(t:*g 1)':\"bcifz\";*(`$t)g n[];t~\"s\";`$u 0;t~\"S\";`$u 0;t~\"L\";u'!n[];t~\"D\";(u 0)!u 0;t~\"T\";+(u 0)!u 0;(`$_t+32)g n[]]};u 0}\n\ncsv:{c:{s:`$'x@i:&x':\"ifzs\";n:`i$\" \"\\:-1_@[x;i;\" \"];y[a]:(y[a],''\"a\"),''y[1+a:&s=`z];s$'y n};s:$[\" \"~(*x);`split@;(*x)\\:];x:1_x;y:+s'$[`L~@y;y;\"\\n\"\\:y];$[#x;c[x;y];y]} / csv[\";1z3i2f\"; \"input..\"]\n\nPW:800;PH:600\n\nplot:{plt::`plot x;`Plot@`plt;s:`c`k`z`d!\"plt.\",/:\"ckzd\"\n ui::\"<canvas width=\",($PW),\" height=\",($PH),\" id='\",(s`c),\"' data-click='\",(s`k),\"' data-zoom='\",(s`z),\"'></canvas><div id='\",(s`d),\"'></div>\"}\n\n`\"plot.\":{[d];l:$!d;v:.d; t:$[2~#d;`xy;`polar];\n y:$[t~`xy; $[`L~@y:v 1;y;,y];          $[`L~@y:_*v;y;,y]]\n x:$[t~`xy; $[`L~@x:v 0;x;(,x)@(#y)#0]; $[`L~@x:imag@*v;x;,x]]\n xt:`tics(&/&/x;|/|/x);yt:`tics(&/&/y;|/|/y)\n a:$[t~`xy;(xt 0;*-1#xt;yt 0;*-1#yt);(-a;a;-a;a:*|`tics@0.,|/|/abs@*v)]\n c:c@(#c:\",\"\\:\"#1f77b4,#ff7f0e,#2ca02c,#d62728,#9467bd,#8c564b,#e377c2,#7f7f7f,#bcbd22,#17becf\")/!#x\n style:$[t~`polar;\"..\";`i~@**y;\"||\";\"--\"] / -.| line points bar\n size: $[t~`polar;2;style~\"||\";(--/((**x),-1#*x))%-1+#*x ;2]\n lines:{`style`size`color`x`y!(style;size;z;x;0.+y)}'[x;y;c]\n `L`t`l`a`f`fw`fh`pw`ph!(lines;t;l;a;\"20px monospace\";FW;FH;$[t~`xy;PW;PH];PH)}\n\n`\"Plot.\":{[sym];x:.sym; w:x`fw; h:x`fh; W:x`pw; H:x`ph; a:x`a\n dst:(w;W-w;H-h;h);rdst:(w;h;W-2*w;H-2*h)\n xs:(a 0 1)(dst 0 1)' /transform axis to canvas\n ys:(a 2 3)(dst 2 3)'\n d:{(x i)!x 1+i:2*!2\\#x}\n bars:{[l]$[\"|\"':l`style;(`fillStyle;l`color),,/{(`fillRect;((-dx%2)+xs x;ys y;dx:-/xs(l`size;0.);(ys a 2)-ys y))}'[l`x;l`y];()]}\n line:{[l]$[\"-\"':l`style;(`lineWidth;l`size;`strokeStyle;l`color;`beginPath;()),(,/(`moveTo,(-1+#l`x)#`lineTo),','+(xs l`x;ys l`y)),(`stroke;());()]}\n dots:{[l]$[\".\"':l`style;(`fillStyle;l`color),,/{(`beginPath;();`arc;(xs x;ys y;1.5*l`size;0;2p);`fill;())}'[l`x;l`y];()]}\n c:`font`textBaseline!(x`f;\"bottom\")\n xy:{[]c,:`fillText!(($a 0;w;H);(s;2\\W-w*#s:(x`l)0;PH);(s;W-w*1+#s:$a 1;H))\n       c,:`fillText!(((\"[\",(\", \"/:$a 2 3),\"]\");w;h);(s;2\\W-w*#s:(x`l)1;h))\n       c,:`rect`strokeStyle`lineWidth`stroke`clip!(rdst;\"black\";2;();())\n       c,:`strokeStyle`lineWidth!(\"lightgrey\";1)\n       c,:d@,/{(`moveTo;(x;dst 2);`lineTo;(x;dst 3);`stroke;())}'xs`tics x[`a;0 1]\n       c,:d@,/{(`moveTo;(dst 0;x);`lineTo;(dst 1;x);`stroke;())}'ys`tics x[`a;2 3]}\n po:{[]c,:`fillText!((s;2\\W-w*#s:(x`l)0;h);($(x`a)3;W%2;H))\n       c,:`strokeStyle`lineWidth`beginPath`arc`stroke`clip!(\"black\";2;();(W%2;H%2;(H%2)-h;0;2p);();())\n       c,:`strokeStyle`lineWidth`moveTo`lineTo`stroke`moveTo`lineTo`stroke!(\"lightgrey\";1;(h;H%2);(W-h;H%2);();(W%2;h);(W%2;H-h);())\n       c,:d@,/{(`beginPath;();`arc;(W%2;H%2;x;0;2p);`stroke;())}'(xs`tics 0.,x[`a;3])-W%2\n };$[`xy~x`t;xy[];po[]]\n c,:d@,/bars'x`L\n c,:d@,/line'x`L\n c,:d@,/dots'x`L\n back:{[a;dst;x;y]((dst 0 1)(a 0 1)'x;(dst 2 3)(a 2 3)'y)}[a;dst]  /transform canvas to axis\n click:{[b;s;t;x;y];(`0)[s;$[t~`xy;\"\\n\"/:(\"x:\";\"y:\"),'$b[x;y]; $imag/|b[x;y]]]}[back;`$($sym),\".d\";x`t]\n zoom:{[b;s;x]a:b[x 0;x 1],b[x 2;x 3];.[s;(`a;0 2 1 3);@[(^a 0 2),^a 1 3;0 2 1 3]];.[s;`t;`xy];`Plot s}[back;sym]\n (`0)[`$($sym),\".c\";c]     /assign to dynamic symbol\n (`0)[`$($sym),\".k\";click]\n (`0)[`$($sym),\".d\";\"\"]\n (`0)[`$($sym),\".z\";zoom]}\n\n `\"tics.\":{[minmax]nice:{[x;r]f:x%0.+10^ex:_log[10;x];(1 2 5 10.@*&(~f>1 2 5 0w;f<1.5 3 7 0w)[r])*10^ex};e:nice[-/|minmax;0];s:nice[e%4.;1];n:_1.5+e%s;$[~(minmax 1)>*-2#r:s*(_(*minmax)%s)+!n;-1_r;r]}\n\nhist:{$[`i~@x;hist[(x;&/y;|/y);y];(Y;(`38)[0;(#'=1+((d%2.0)+-1_Y:(x 1)+(d:(--/1_x)%-1.+n)*!n)'y)@!n:_0.+*x])]} /[xn;y] [xn xmin xmax;y] e.g. plot `x`y!hist[10 -4 4.;?-1000]\n")
	zn := int32(6872) // should end before 8k
	x := mk(Ct, zn)
	Memorycopy(int32(x), 600, zn)
	dx(Val(x))
}

package main

import . "github.com/ktye/wg/module"

func zk() {
	Data(600, "/               64     72    80    88     96     104    112    120    128                 \n``x`y`z`k`l`a`b`while`\"rf.\"`\"rz.\"`\"uqs.\"`\"uqf.\"`\"gdt.\"`\"lin.\"`\"odo.\"`\"grp.\"\n\n`\"x.\":{,/+\"0123456789abcdef\"@(x%16;16/x:256/256+x)} /`x@ (hex)\n`\"t.\":`45         /`t@ token\n`\"p.\":`46         /`p@ parse\n`\"b.\":(`46)[`b;]  /`b@ reinterpret\n`\"c.\":(`46)[`c;]\n`\"i.\":(`46)[`i;]\n`\"s.\":(`46)[`s;]\n`\"f.\":(`46)[`f;]\n`\"z.\":(`46)[`z;]\n\n`\"l.\":{t:@x;k:`kxy 1\n kt:{x:$[`T~@x;T x;`pad(\"\";\"-\"),$x];(x,'\"|\"),'T y}\n d:{r:!x;x:.x;$[`T~@x;kt[r;x];,'[,'[`pad(k'r);\"|\"];k'x]]}\n T:{$[`L':@'.x;,k x;(,*x),(,(#*x)#\"-\"),1_x:\" \"/:'+`pad@'$(!x),'.x]}\n dd:(\"\";,\"..\")20<#x:$[(@x)':`L`D`T;x;x~*x;x;[t:`L;,x]]\n x:$[x~*x;x;(20&#x)#x]\n $[`D~t;d x;`T~t;T x;x~*x;,k x;k'x],dd} /352\n\n`\"pad.\":{(|/#'x)#'x}\n`\"split.\":{$[`L~@x;`split@'x;\" \"\\:$[\" \"=x@-1+#x:x@&~i&~':i:\" \"=x;-1_x;x]]} /multiple whitespace/leading/trailing\n`\"edit.\":{et:{\"+\",(`k@!x),\"!\",el@.x}\n el:{t:(`S`B`C`I`F`Z!``b`c`i`f`z)@@'x; (`k@t),\"$'+`split@'\\\"\\\\n\\\"\\\\:-1_1_\\\"\\n\",(\"\\n\"/:\" \"/:'+`pad@'$x),\"\\n\\\"\"}\n $[`T~t:@x;et x;`L~t;el x;`C~t;\"-1_1_\\\"\\n\",x,\"\\n\\\"\";\"*\",el@,x]}\n\n`\"kxy.\":{t:@y;n:#y;k:`kxy x;m:x\n q:{c,(\"\\\\\"/:(0,i)^@[x;i;(qs!\"tnr\\\"\\\\\")x i:&x':qs:\"\\t\\n\\r\\\"\\\\\"]),c:_34}\n s:{$[|/x':\"\\t\\n\\r\"__!31;\"0x\",`x@x;q x]}\n a:{x:$x;$[`c~t;s x;`s~t;\"`\",x;x]}\n d:{r:\"!\",k@.x;n:#!x;x:k@!x;$[(1~n)|(@.x)':`D`T;\"(\",x,\")\";x],r}\n v:{m*:(.`\".kstm\")t; dd:(\"\";\"..\")m<#x;x:(m&#x)#x\n  x:$[`L~t;k'x;`C~t;x;$x]\n  x:$[`B~t;(*'x),\"b\";`C~t;s x;`S~t;c,(c:\"`\")/:x;`L~t;$[1~n;*x;\"(\",(\";\"/:x),\")\"];\" \"/:x]\n  ((\"\";\",\")(1~n)),x,dd}\n $[n~0;(.`\".kst0\")@t;`T~t;\"+\",d@+y;`D~t;d y;y~*y;a y;v y]} /344\n \n`\"k.\":`kxy 1000000\n\n`\"uqs.\":{x@&~0b~':x:^x} \n`\"uqf.\":{x@&(!#x)=x?x}\n`\"gdt.\":{[t;g](!#t){x g y x}/|.t}\n\n`\".kst0\":`B`C`I`S`F`Z`L!(\"0#0b\";c,c:_34;\"!0\";\"0#`\";\"0#0.\";\"0#0a\";\"()\")\n`\".kstm\":`B`C`I`S`F`Z`L!100 100 30 30 20 10 20\n\n`\"lin.\":{$[`L~@z;(.`\"lin.\")[x;y]'z;[dx:0.+1_-':x;dy:0.+1_-':y;b:(-2+#x)&0|x'z;(y b)+(dy b)*(z-x b)%dx b]]}\n`\"odo.\":{{x/y}'[x;(!*/x)%/:{(*/x)%\\ x}x]}\n`\"grp.\":{(x@*'g)!g:(&~x~':x i)^i:<x}\n\ndot:{[xt;y]{+/x*y}\\:[xt;y]} /col-major\n\nsolve:{qslv:{H:x 0;r:x 1;n:x 2;m:x 3;j:0;K:!m\n while[j<n;y[K]-:(+/(conj H[j;K])*y K)*H[j;K];K:1_K;j+:1]\n i:n-1;J:!n;y[i]%:r@i\n while[i;j:i_J;i-:1;y[i]:(y[i]-+/H[j;i]*y@j)%r@i]\n n#y}\n q:$[`i~@*|x;x;qr x];$[`L~@y;qslv/:[q;y];qslv[q;y]]}\n\nqr:{K:!m:#*x;I:!n:#x;j:0;r:n#0a;turn:$[`Z~@*x;{(-x)@angle y};{x*1. -1@y>0}]\n while[j<n;I:1_I\n  r[j]:turn[s:abs@abs/j_x j;xx:x[j;j]]\n  x[j;j]-:r[j]\n  x[j;K]%:%s*(s+abs xx)\n  x[I;K]-:{+/x*y}/:[(conj x[j;K]);x[I;K]]*\\:x[j;K]\n  K:1_K;j+:1];(x;r;n;m)}\n\nany:`30;abs:`32;sin:`44;cos:`39;find:`31;fill:`38;imag:`33;conj:`34;angle:`35;exp:`42;log:`43\n\nej:{(y j),'x_z i j:&~0N=i:(z x)?y x} /sym t1 t2\navg:{(+/x)%0.+#x}\nvar:{(+/x*x:(x-avg x))%-1+#x}\nstd:{%var x}\nrem:{x/x+x/y} /like apl for negative arguments\n\n`\"rf.\": {.5+(x?0)%4294967295.}\n`\"rf1.\":{.5+(1.+x?0)%4294967295.}\n`\"rz.\": {(%-2*log `rf1 x)@360.*`rf x}\n`\".html\":{$[`L~@x;\"<div style='display:flex;flex-direction:column'>\",(,/(.`\".html\")'x),\"</div>\";`S~@x;\"<div style='display:flex;flex-direction:row'>\",(,/(.`\".html\")'x),\"</div>\";~`s~@x;\"\";{{\"<\",y,\" id='\",x,\"'></\",y,\">\"}[x;$[y':`i`c`f`z`s`I`B`C`F`Z`S;\"input\";`b~y;\"input type='checkbox'\";`T~y;\"table\";`l~t;\"button\";\"pre\"]]}[$x;@.x]]}\n\n`\"pack.\":{w:{(`c@,#x),x};($t),$[`s~t:@x;`pack@$x;x~*x;w `c@,x;`L~@x;(`c@,#x),,/`pack@'x;(@x)':`D`T;(`pack@.x),`pack@!x;`S~t;,/`pack@$x;w `c x]}\n`\"unpack.\":{s:x;g:{[n]r:n#s;s::n_s;r};n:{*`i@g 4};u:{x;$[(t:*g 1)':\"bcifz\";*(`$t)g n[];t~\"s\";`$u 0;t~\"S\";`$u 0;t~\"L\";u'!n[];t~\"D\";(u 0)!u 0;t~\"T\";+(u 0)!u 0;(`$_t+32)g n[]]};u 0}\n\ncsv:{c:{s:`$'x@i:&x':\"ifzs\";n:`i$\" \"\\:-1_@[x;i;\" \"];y[a]:(y[a],''\"a\"),''y[1+a:&s=`z];s$'y n};s:$[\" \"~(*x);`split@;(*x)\\:];x:1_x;y:+s'$[`L~@y;y;\"\\n\"\\:y];$[#x;c[x;y];y]} / csv[\";1z3i2f\"; \"input..\"]\n\nucal:{[s;u;r](#u)#+solve[u,0a+r=/:?r;s]} /k\npcal:{[s;u]+solve[+u;+s]}                /k  s:s0-s1, u:u0\nuslv:{[qrk;s]qrsolve[qrk;s]}             /u\nuidx:{[u;a]solve[((#u)#1a;1@a);u] }      /(u0;uR)   rref:uidx[u*1@a;a]\n\nPW:800;PH:600;FH:20;FONT:\"monospace\"\n\n`\"pltnn.\":{wi:</</:;$[#i:&wi[*x;+\\plt`px`pw]&wi[x 1;+\\plt`py`ph];*i;0]}\n`\"pltco.\":{[p;x;y]w:p`fh;h:p`fh;X:p`px;Y:p`py;W:p`pw;H:p`ph;C:(X+W%2;Y+H%2);R:(W%2)&(H%2)-h;d:$[`xy~p`t;(X+w;X+W-w;Y+H-h;Y+h);((C-R),C+R)0 2 3 1];a:p`a; ((d 0 1)(a 0 1)'x;(d 2 3)(a 2 3)'y)}\n`\"pltcl.\":{[x;y]p:.`\"pltco.\";n:`pltnn(x;y); xy:p[plt n;x;y]; `<$[`polar~plt[n;`t];$imag/|xy;\"x:\",($xy 0), \" y:\",($xy 1)],_10 32}\n`\"pltzo.\":{[x;y;w;h]n:`pltnn(x;y);p:.`\"pltco.\";xy0:p[plt n;x;y+h]; xy1:p[plt n;x+w;y]; plt[`a;n]:(xy0,xy1)0 2 1 3; plt[`t;n]:`xy; draw[`plts@`plt;(PW;PH)]}\n\nplot:{n:#x;multi:{x[`pw]%:n;x[`px]:(!n)*x`pw;x};plt::$[(@x)':`L`T;multi@(`plot@)'x;,`plot x];Show[draw[`plts@`plt;(PW;PH)];.`\"pltcl.\";.`\"pltzo.\"]}\n\n`\"plot.\":{[d]l:$!d;v:.d; t:$[2~#d;`xy;`polar];\n y:$[t~`xy; $[`L~@y:v 1;y;,y];          $[`L~@y:_*v;y;,y]]\n x:$[t~`xy; $[`L~@x:v 0;x;(,x)@(#y)#0]; $[`L~@x:imag@*v;x;,x]]\n xt:`tics(&/&/x;|/|/x);yt:`tics(&/&/y;|/|/y)\n a:$[t~`xy;(xt 0;*-1#xt;yt 0;*-1#yt);(-a;a;-a;a:*|`tics@0.,|/|/abs@*v)]\n c:c@(#c:11826975 950271 2924588 2631638 12412820 4937356 12744675 8355711 2276796 13614615)/!#x\n style:$[t~`polar;\"..\";`i~@**y;\"||\";\"--\"] / -.| line points bar\n size: $[t~`polar;2;style~\"||\";(--/((**x),-1#*x))%-1+#*x ;2]\n lines:{`style`size`color`x`y!(style;size;z;x;0.+y)}'[x;y;c]\n pw:PW;ph:PH;`L`T`t`l`a`f`fh`px`py`pw`ph!(lines;\"\";t;l;a;FONT;FH;0;0;pw;ph)}\n \n`\"plts.\":{[sym];x:.sym;$[`D~@x;`Plot x;,/(`Plot@)'x]}\n\n`\"Plot.\":{[x];w:x`fh; h:x`fh; X:x`px; Y:x`py; W:x`pw; H:x`ph; a:x`a;T:x`T;grey:13882323\n C:(X+W%2;Y+H%2);R:(W%2)&(H%2)-h\n dst:$[`xy~x`t;(X+w;X+W-w;Y+H-h;Y+h);((C-R),C+R)0 2 3 1];rdst:(X+w;Y+h;W-2*w;H-2*h)\n xs:(a 0 1)(dst 0 1)' /transform axis to canvas\n ys:(a 2 3)(dst 2 3)'\n bars:{[l]$[\"|\"':l`style;(`color;l`color),,/{(`Rect;((-dx%2)+xs x;ys y;dx:-/xs(l`size;0.);(ys a 2)-ys y))}'[l`x;l`y];()]}\n line:{[l]$[\"-\"':l`style;(`linewidth;l`size;`color;l`color;`poly;(xs l`x;ys l`y));()]}\n dots:{[l]$[\".\"':l`style;(`color;l`color),,/{(`Circle;(xs x;ys y;1.5*l`size))}'[l`x;l`y];()]}\n c:(`clip;(X;Y;W;H);`font;(x`f;x`fh);`color;0;`text;((X+W%2;Y+h);1;T))\n xy:{[]c,:(`text;((X+w;Y+H);0;$a 0);`text;((X+W%2;Y+H);1;(x`l)0);`text;((X+W-w;Y+H);2;$a 1))\n       c,:(`Text;((X+w;Y+H-h);0;$a 2);`Text;((X+w;Y+H%2);2;(x`l)1);`Text;((X+w;Y+h);2;$a 3))\n       c,:(`color;0;`linewidth;2;`rect;rdst)      /todo: clip rdst\n       c,:(`linewidth;1;`color;grey)\n       c,:(`clip;rdst)\n       c,:,/{(`line;0.+(x;dst 2;x;dst 3))}'xs`tics x[`a;0 1]\n       c,:,/{(`line;0.+(dst 0;x;dst 1;x))}'ys`tics x[`a;2 3]}\n po:{[]c,:(`text;((C 0;Y+H);1;(x`l)0);`text;(C+.75*R;6;$(x`a)3))\n       c,:(`font;(x`f;_h*.8)),,/{(`text;(C+R*(_;imag)@'x;y;z))}'[1@270.+a;0 0 6 6 4 4 2 2;$a:30 60 120 150 210 240 300 330]\n       c,:(`color;0),/{(`line;,/+C+(R-w%2;R)*/:(_;imag)@'x)}'1@30.*!12\n       /c,:(`clip;C,R) /bug in cairo?\n       c,:(`color;grey;`linewidth;1;`line;((-R)+*C;C 1;R+*C;C 1);`line;(*C;(-R)+C 1;*C;R+C 1))\n       c,:,/{(`circle;0.+C,x)}'r:(xs@`tics 0.,x[`a;3])-*C\n       c,:(`color;0;`linewidth;2;`circle;C,R)}\n $[`xy~x`t;xy[];po[]]\n c,:,/bars'x`L\n c,:,/line'x`L\n c,:,/dots'x`L}\n \n`\"tics.\":{[minmax]nice:{[x;r]f:x%0.+10^ex:_log[10;x];(1 2 5 10.@*&(~f>1 2 5 0w;f<1.5 3 7 0w)[r])*10^ex};e:nice[-/|minmax;0];s:nice[e%4.;1];n:_1.5+e%s;$[~(minmax 1)>*-2#r:s*(_(*minmax)%s)+!n;-1_r;r]}\n\n`\"ceg.\":{(x i)!0-':1+i:&(1_~~':x),1b} /#'= (^x)\nhist:{$[`i~@x;hist[(x;&/y;|/y);y];(Y;(`38)[0;(`ceg@^1+((d%2.0)+-1_Y:(x 1)+(d:(--/1_x)%-1.+n)*!n)'y)@!n:_0.+*x])]} /[xn;y] [xn xmin xmax;y] e.g. plot `x`y!hist[10 -4 4.;?-1000]\n")
	zn := int32(7451) // should end before 8k
	x := mk(Ct, zn)
	Memorycopy(int32(x), 600, zn)
	dx(Val(x))
}

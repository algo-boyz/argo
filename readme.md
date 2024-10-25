

todo: 
    - run argo which results in asm.s output
    - use [LinGo- micro framework for building DSLs](https://about.gitlab.com/blog/2022/05/26/a-go-micro-language-framework-for-building-dsls/)

arGO - arm assembly like GO by Go

Overview of arGO's 

	1.	Tokenization: chunk source code into tokens (the smallest elements of meaning).
	2.	Parse: convert tokens into an AST by following arGO's grammar rules.
	3.	Generate Code: walk through our AST and emit those juicy ARM assembly instructions.

| Token Type | Value|
|----------|--------|
|Identifier|      x |
|Operator  |	 := |
|Literal   |	 42 |
|Identifier|	  y |
|Operator  |	 := |
|Identifier|	  x |
|Operator  |	  + |

```Go
➜  asm (wellsync-us-dev-core-aks:careconnect) go run cmd/argo/argo.go 
Token: {Type:IDENT Value:x}
Token: {Type:ASSIGN Value::=}
Token: {Type:NUMBER Value:42}
Token: {Type:OP Value:+}
Token: {Type:NUMBER Value:5}
Token: {Type:EOF Value:}
AST: &{Name:x Value:0x1400007c300}
```
```arm
MOV X0, #42
MOV X1, #5
ADD X2, X0, X1
STR X0, [x]
```

<!-- <program> ::= <statement>*

<statement> ::= <declaration> | <assignment> | <expression> | <if-statement> | <loop>

<declaration> ::= "var" <identifier> "=" <expression>
<assignment>  ::= <identifier> "=" <expression>

<expression>  ::= <literal> | <binary-op> | <function-call>

<binary-op>   ::= <expression> <operator> <expression>
<operator>    ::= "+" | "-" | "*" | "/"

<function-call> ::= <identifier> "(" <expression>* ")"

<if-statement> ::= "if" <expression> <block> ["else" <block>]
<loop> ::= "for" <expression> <block>

<block> ::= "{" <statement>* "}"

<identifier> ::= [a-zA-Z_][a-zA-Z0-9_]*

<literal> ::= [0-9]+ -->


[Whirlwind Tour of ARM Assembly](https://www.coranac.com/tonc/text/asm.htm)
[Azeria - intro to ARM ASM](https://azeria-labs.com/writing-arm-assembly-part-1/)

```
    $ as hello.s -o hello.o
    $ ld hello.o -o hello -l System -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _main -arch arm64
```
Build and link hello world program 


SVC 

[arm simulator](https://cpulator.01xz.net/?sys=arm)
[arm sysCall list](https://arm.syscall.sh/)
[linux x86 64 sysCall list](https://blog.rchapman.org/posts/Linux_System_Call_Table_for_x86_64/)
[macOS sysCall list](https://github.com/rewired-gh/macos-system-call-table/tree/main)

Tutorials

[Laurie Wired](https://www.youtube.com/watch?v=kKtWsuuJEDs&list=PLn_It163He32Ujm-l_czgEBhbJjOUgFhg&index=1)
[HelloSilicon](https://github.com/below/HelloSilicon?tab=readme-ov-file)
[Azeria student notes](https://github.com/azeria-labs/ARM-assembly-examples/blob/master/write.s)

[Assembly Language Primer](https://github.com/pkivolowitz/asm_book?tab=readme-ov-file)
Sources

[Linux kernel cheat sheet](https://github.com/cirosantilli/linux-kernel-module-cheat/tree/master/userland/arch/arm)
[ARM asm exercises](https://github.com/chibby0ne/ARM_Assembly)
[ARM RPI programs](https://github.com/hadefuwa/ARM-Assembly)
[ARM practice](https://github.com/memoriasIT/Assembly-Practice/tree/master/ARM%20Projects)
[ARM asm: Fundamentals and Techniques](https://github.com/jincongho/ARM-Assembly-Language-Fundamentals-and-Techniques)
[Professional Assembly](https://github.com/bobblestiltskin/professional_assembly_language/blob/master/arm/chap05/Makefile)
[RPI asm](https://github.com/Apress/Raspberry-Pi-Assembly-Language-Programming)
[Qfplib](https://github.com/mysterywolf/Qfplib-M0-full)
[macOS ARM Shellcode](https://github.com/daem0nc0re/macOS_ARM64_Shellcode)
[Rverse eng](https://github.com/mytechnotalent/Reverse-Engineering)
[Reversing for everyone](https://0xinfection.github.io/reversing/reversing-for-everyone.pdf)
[EggHunter](https://github.com/kmkz/Assembly-language/blob/master/EggHunter.asm)
[Snake](https://github.com/haperofi/asm_snake)
[Assembly Lab](https://github.com/Nitesh8998/arm_assembly_lab/blob/master/lab3.s)
[Defcon31](https://github.com/eigentourist/defcon31?tab=readme-ov-file)

[Solving the Mystery of ARM7TDMI Multiply Carry Flag](https://bmchtech.github.io/post/multiply/)

[Arm by Example](https://armasm.com/)
https://hackaday.com/2022/10/11/arm-programming-by-example/
.global _start
.align 2

_start:
    b _printf
    b _terminate

_printf:
    mov X0, #1 // stdout
    adr X1, helloworld // ptr `hello world`
    mov X2, #12 // stdout
    mov X16, #4 // write
    
_reboot:
    mov X0, #1 // instant reboot
    mov X16, #55 // reboot
    svc 0   // syscall
    
_terminate:
    mov X0, #0 // return 0
    mov X16, #93 // exit
    svc 0   // syscall

helloworld: .ascii "hello world\n"
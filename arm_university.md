ARM assembly language provides a wide range of instructions to perform various operations, including arithmetic, logical operations, data movement, and flow control. Here’s a list of 30 ARM assembly instructions, along with a brief description of what each does. These instructions follow the standard ARM architecture (32-bit, ARMv7-A as an example).

1. MOV – Move a value into a register.

	•	Example: MOV R0, #5 – Move the value 5 into register R0.

2. ADD – Add two values.

	•	Example: ADD R0, R1, R2 – Add the values in R1 and R2, store the result in R0.

3. SUB – Subtract one value from another.

	•	Example: SUB R0, R1, R2 – Subtract the value in R2 from R1, store the result in R0.

4. RSB – Reverse subtract.

	•	Example: RSB R0, R1, #0 – Subtract R1 from 0 and store the result in R0.

5. MUL – Multiply two values.

	•	Example: MUL R0, R1, R2 – Multiply R1 by R2, store the result in R0.

6. DIV – Divide two values (on some ARM architectures, but not ARMv7).

	•	Example: SDIV R0, R1, R2 – Divide R1 by R2, store the result in R0.

7. AND – Bitwise AND operation.

	•	Example: AND R0, R1, R2 – Perform a bitwise AND of R1 and R2, store the result in R0.

8. ORR – Bitwise OR operation.

	•	Example: ORR R0, R1, R2 – Perform a bitwise OR of R1 and R2, store the result in R0.

9. EOR – Bitwise Exclusive OR (XOR).

	•	Example: EOR R0, R1, R2 – Perform a bitwise XOR of R1 and R2, store the result in R0.

10. BIC – Bitwise AND with complement.

	•	Example: BIC R0, R1, R2 – Perform a bitwise AND of R1 with the complement of R2.

11. CMP – Compare two values (sets flags based on the result).

	•	Example: CMP R0, R1 – Compare R0 and R1 by subtracting R1 from R0 (result affects flags).

12. CMN – Compare Negative (sets flags as if adding two values).

	•	Example: CMN R0, R1 – Compare by adding R0 and R1 (affects flags).

13. TST – Test bits (bitwise AND, sets flags based on result).

	•	Example: TST R0, R1 – Perform a bitwise AND between R0 and R1, set flags.

14. TEQ – Test Equivalence (bitwise XOR, sets flags based on result).

	•	Example: TEQ R0, R1 – Perform a bitwise XOR between R0 and R1, set flags.

15. LDR – Load a value from memory into a register.

	•	Example: LDR R0, [R1] – Load the value at the address in R1 into R0.

16. STR – Store a register value into memory.

	•	Example: STR R0, [R1] – Store the value in R0 into the memory address in R1.

17. LDRB – Load a byte from memory.

	•	Example: LDRB R0, [R1] – Load a byte at the address in R1 into R0.

18. STRB – Store a byte into memory.

	•	Example: STRB R0, [R1] – Store the byte in R0 into the address in R1.

19. LDM – Load multiple registers from memory.

	•	Example: LDM R0, {R1, R2} – Load the values at addresses starting at R0 into R1 and R2.

20. STM – Store multiple registers into memory.

	•	Example: STM R0, {R1, R2} – Store the values in R1 and R2 starting at the address in R0.

21. B – Unconditional branch (jump to a label).

	•	Example: B loop – Jump to the label “loop”.

22. BL – Branch with link (call a subroutine).

	•	Example: BL function – Call the subroutine at the label “function” and save the return address.

23. BX – Branch and exchange instruction set (switch to Thumb or ARM mode).

	•	Example: BX LR – Branch to the address in the link register (LR) and switch mode.

24. BEQ – Branch if equal (based on condition flags).

	•	Example: BEQ label – Branch to “label” if the zero flag (Z) is set (equality).

25. BNE – Branch if not equal.

	•	Example: BNE label – Branch to “label” if the zero flag is clear (inequality).

26. BGT – Branch if greater than.

	•	Example: BGT label – Branch if greater than (N and V flags match, and Z flag is clear).

27. BLT – Branch if less than.

	•	Example: BLT label – Branch if less than (N flag differs from V flag).

28. LSL – Logical shift left.

	•	Example: LSL R0, R1, #2 – Shift the bits in R1 left by 2 places, store the result in R0.

29. LSR – Logical shift right.

	•	Example: LSR R0, R1, #2 – Shift the bits in R1 right by 2 places, store the result in R0.

30. ROR – Rotate bits right.

	•	Example: ROR R0, R1, #2 – Rotate the bits in R1 right by 2 places, store the result in R0.
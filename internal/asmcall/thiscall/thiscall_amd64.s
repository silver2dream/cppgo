TEXT ·Call0(SB),0,$32-24
  MOVQ a+8(FP), CX

  MOVQ addr+0(FP), AX
  CALL AX
  MOVQ AX, ret+16(FP)
  RET

TEXT ·Call1(SB),0,$32-32
  MOVQ a+8(FP), CX
  MOVQ b+16(FP), DX

  MOVQ addr+0(FP), AX
  CALL AX
  MOVQ AX, ret+24(FP)
  RET

TEXT ·Call2(SB),0,$32-40
  MOVQ a+8(FP), CX
  MOVQ b+16(FP), DX
  MOVQ c+24(FP), R8

  MOVQ addr+0(FP), AX
  CALL AX
  MOVQ AX, ret+32(FP)
  RET

TEXT ·Call3(SB),0,$32-48
  MOVQ a+8(FP), CX
  MOVQ b+16(FP), DX
  MOVQ c+24(FP), R8
  MOVQ d+32(FP), R9

  MOVQ addr+0(FP), AX
  CALL AX
  MOVQ AX, ret+40(FP)
  RET

TEXT ·Call4(SB),0,$40-56
  MOVQ a+8(FP), CX
  MOVQ b+16(FP), DX
  MOVQ c+24(FP), R8
  MOVQ d+32(FP), R9
  MOVQ e+40(FP), BX
  MOVQ BX, e-8(SP)

  MOVQ addr+0(FP), AX
  CALL AX
  MOVQ AX, ret+48(FP)
  RET

TEXT ·Call5(SB),0,$48-64
  MOVQ a+8(FP), CX
  MOVQ b+16(FP), DX
  MOVQ c+24(FP), R8
  MOVQ d+32(FP), R9
  MOVQ e+40(FP), BX
  MOVQ BX, e-16(SP)
  MOVQ f+48(FP), BX
  MOVQ BX, f-8(SP)

  MOVQ addr+0(FP), AX
  CALL AX
  MOVQ AX, ret+56(FP)
  RET
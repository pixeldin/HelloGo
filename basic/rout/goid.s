#include "textflag.h"
#include "funcdata.h"


// func getGptr() unsafe.Pointer
TEXT ·getGptr(SB), NOSPLIT, $0-8
	MOVQ (TLS), BX
	MOVQ BX, ret+0(FP)
	RET

TEXT ·GoId(SB),NOSPLIT,$0-8
    NO_LOCAL_POINTERS
    MOVQ ·goid_offset(SB),AX
    // get runtime.g
    MOVQ (TLS),BX
    ADDQ BX,AX
    MOVQ (AX),BX
    MOVQ BX,ret+0(FP)
    RET

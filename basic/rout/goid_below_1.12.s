//+build !go1.12

#include "textflag.h"
#include "funcdata.h"

// func getG() interface{}
TEXT ·getG(SB),NOSPLIT,$32-16
    NO_LOCAL_POINTERS
    MOVQ $0, ret_type+0(FP)
    MOVQ $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED
    // get runtime.g
    MOVQ (TLS),BX
    // get runtime.g type
    MOVQ $type·runtime·g(SB),AX
    // convert (*g) to interface{}
    MOVQ AX,  0(SP)
    MOVQ BX,  8(SP)
    CALL runtime·convT2E(SB)
    MOVQ  16(SP),AX
    MOVQ  24(SP),BX
    // return interface{}
    MOVQ AX,ret_type+0(FP)
    MOVQ BX,ret_data+8(FP)
    RET


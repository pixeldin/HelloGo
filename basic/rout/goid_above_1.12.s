//+build go1.12

#include "textflag.h"
#include "funcdata.h"

// func getG() interface{}
TEXT ·getG(SB),NOSPLIT,$32-16
    NO_LOCAL_POINTERS
    MOVQ $0, ret_type+0(FP)
    MOVQ $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED
    RET


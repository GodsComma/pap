package constants

// MASKS FOR MOV OP_CODE
const OP_MASK_R2R byte = 0b11111100
const OP_MASK_I2R byte = 0b11110000
const OP_MASK_DI2R byte = 0b11111110
const OP_MASK_M2A byte = 0b11111110

// MASK ADD OP
const OP_MASK_ADD_2B byte = 0b11111100
const OP_MASK_ADD_1B byte = 0b11111110

// MASK MOVE ARGUMENTS 1st Byte
const D_MASK byte = 0b00000010
const W_MASK byte = 0b00000001
const S_MASK byte = 0b00000010

// MASK MOVE ARGUMENTS 2nd Byte
const MOD_MASK byte = 0b11000000
const REG_MASK byte = 0b00111000
const RM_MASK byte = 0b00000111

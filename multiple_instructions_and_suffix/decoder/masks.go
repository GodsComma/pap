package decoder

// MASKS FOR THE OP_CODE
const OP_MASK_R2R byte = 0b11111100
const OP_MASK_I2R byte = 0b11110000

// MASK MOVE ARGUMENTS 1st Byte
const D_MASK byte = 0b00000010
const W_MASK byte = 0b00000001

// MASK MOVE ARGUMENTS 2nd Byte
const MOD_MASK byte = 0b11000000
const REG_MASK byte = 0b00111000
const RM_MASK byte = 0b00000111

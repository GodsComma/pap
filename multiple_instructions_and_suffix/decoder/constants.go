package decoder

// OP_CODES
const MOV_R2R byte = 0b10001000
const MOV_I2R byte = 0b10110000
const MOV_DI2R byte = 0b11000110

// Direct Address
const D_ADDR = 0b00000110

// MOD_MODES
const MM_ND0 = 0b00000000
const MM_08B = 0b01000000
const MM_16B = 0b10000000
const MM_RM0 = 0b11000000

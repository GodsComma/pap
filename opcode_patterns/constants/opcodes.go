package constants

// MOV
const MOV_R2R byte = 0b10001000
const MOV_I2R byte = 0b10110000
const MOV_DI2R byte = 0b11000110
const MOV_M2A byte = 0b10100000
const MOV_A2M byte = 0b10100010

// ADD
const ADD_R2RM byte = 0b00000000
const ADD_I2RM byte = 0b10000000
const ADD_I2A byte = 0b00000100

// Direct Address
const D_ADDR byte = 0b00000110

// MOD_MODES
const MM_ND0 byte = 0b00000000
const MM_08B byte = 0b01000000
const MM_16B byte = 0b10000000
const MM_RM0 byte = 0b11000000

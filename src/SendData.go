// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SendDataABI is the input ABI used to generate the binding from.
const SendDataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ORACLE_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_index\",\"type\":\"string\"},{\"name\":\"_calldata\",\"type\":\"string\"}],\"name\":\"send\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reqId\",\"type\":\"bytes32\"},{\"name\":\"_stateCode\",\"type\":\"uint64\"},{\"name\":\"_response\",\"type\":\"bytes\"}],\"name\":\"getOracelResponse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resp\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_stateCode\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_randomNum\",\"type\":\"bytes\"}],\"name\":\"OracleResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_winner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_funds\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_roundTimes\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playId\",\"type\":\"uint256\"}],\"name\":\"DistributeFunds\",\"type\":\"event\"}]"

// SendDataBin is the compiled bytecode used for deploying new contracts.
var SendDataBin = "0xPUSH1 0x80 PUSH1 0x40 MSTORE PUSH2 0x7530 PUSH1 0x0 SSTORE CALLER PUSH1 0x1 PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP PUSH1 0x0 SLOAD CALLVALUE LT ISZERO ISZERO ISZERO PUSH2 0xC5 JUMPI PUSH1 0x40 MLOAD PUSH32 0x8C379A000000000000000000000000000000000000000000000000000000000 DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE PUSH1 0x13 DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x6D696E20626574206973203120737A61626F2100000000000000000000000000 DUP2 MSTORE POP PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST PUSH2 0xD44 DUP1 PUSH2 0xD4 PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN INVALID PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x83 JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x31D8A985 EQ PUSH2 0x88 JUMPI DUP1 PUSH4 0x7ADBF973 EQ PUSH2 0xB3 JUMPI DUP1 PUSH4 0x7DC0D1D0 EQ PUSH2 0x104 JUMPI DUP1 PUSH4 0x8DA5CB5B EQ PUSH2 0x15B JUMPI DUP1 PUSH4 0xBD6DE11C EQ PUSH2 0x1B2 JUMPI DUP1 PUSH4 0xC265028D EQ PUSH2 0x311 JUMPI DUP1 PUSH4 0xE9F5E299 EQ PUSH2 0x3CD JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x94 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x9D PUSH2 0x45D JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0xBF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x102 PUSH1 0x4 DUP1 CALLDATASIZE SUB PUSH1 0x20 DUP2 LT ISZERO PUSH2 0xD6 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH2 0x463 JUMP JUMPDEST STOP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x110 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x119 PUSH2 0x56C JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x167 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x170 PUSH2 0x592 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1BE JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x30F PUSH1 0x4 DUP1 CALLDATASIZE SUB PUSH1 0x40 DUP2 LT ISZERO PUSH2 0x1D5 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 PUSH5 0x100000000 DUP2 GT ISZERO PUSH2 0x1F2 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP3 ADD DUP4 PUSH1 0x20 DUP3 ADD GT ISZERO PUSH2 0x204 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP2 DUP5 PUSH1 0x1 DUP4 MUL DUP5 ADD GT PUSH5 0x100000000 DUP4 GT OR ISZERO PUSH2 0x226 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST SWAP2 SWAP1 DUP1 DUP1 PUSH1 0x1F ADD PUSH1 0x20 DUP1 SWAP2 DIV MUL PUSH1 0x20 ADD PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 SWAP4 SWAP3 SWAP2 SWAP1 DUP2 DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP4 DUP1 DUP3 DUP5 CALLDATACOPY PUSH1 0x0 DUP2 DUP5 ADD MSTORE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND SWAP1 POP DUP1 DUP4 ADD SWAP3 POP POP POP POP POP POP POP SWAP2 SWAP3 SWAP2 SWAP3 SWAP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 PUSH5 0x100000000 DUP2 GT ISZERO PUSH2 0x289 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP3 ADD DUP4 PUSH1 0x20 DUP3 ADD GT ISZERO PUSH2 0x29B JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP2 DUP5 PUSH1 0x1 DUP4 MUL DUP5 ADD GT PUSH5 0x100000000 DUP4 GT OR ISZERO PUSH2 0x2BD JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST SWAP2 SWAP1 DUP1 DUP1 PUSH1 0x1F ADD PUSH1 0x20 DUP1 SWAP2 DIV MUL PUSH1 0x20 ADD PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 SWAP4 SWAP3 SWAP2 SWAP1 DUP2 DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP4 DUP1 DUP3 DUP5 CALLDATACOPY PUSH1 0x0 DUP2 DUP5 ADD MSTORE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND SWAP1 POP DUP1 DUP4 ADD SWAP3 POP POP POP POP POP POP POP SWAP2 SWAP3 SWAP2 SWAP3 SWAP1 POP POP POP PUSH2 0x5B8 JUMP JUMPDEST STOP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x31D JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x3B3 PUSH1 0x4 DUP1 CALLDATASIZE SUB PUSH1 0x60 DUP2 LT ISZERO PUSH2 0x334 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 DUP1 CALLDATALOAD PUSH8 0xFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 PUSH5 0x100000000 DUP2 GT ISZERO PUSH2 0x36F JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP3 ADD DUP4 PUSH1 0x20 DUP3 ADD GT ISZERO PUSH2 0x381 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP2 DUP5 PUSH1 0x1 DUP4 MUL DUP5 ADD GT PUSH5 0x100000000 DUP4 GT OR ISZERO PUSH2 0x3A3 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST SWAP1 SWAP2 SWAP3 SWAP4 SWAP2 SWAP3 SWAP4 SWAP1 POP POP POP PUSH2 0x89A JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x3D9 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x3E2 PUSH2 0xA02 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE DUP4 DUP2 DUP2 MLOAD DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP1 DUP4 DUP4 PUSH1 0x0 JUMPDEST DUP4 DUP2 LT ISZERO PUSH2 0x422 JUMPI DUP1 DUP3 ADD MLOAD DUP2 DUP5 ADD MSTORE PUSH1 0x20 DUP2 ADD SWAP1 POP PUSH2 0x407 JUMP JUMPDEST POP POP POP POP SWAP1 POP SWAP1 DUP2 ADD SWAP1 PUSH1 0x1F AND DUP1 ISZERO PUSH2 0x44F JUMPI DUP1 DUP3 SUB DUP1 MLOAD PUSH1 0x1 DUP4 PUSH1 0x20 SUB PUSH2 0x100 EXP SUB NOT AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP JUMPDEST POP SWAP3 POP POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST PUSH1 0x0 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x528 JUMPI PUSH1 0x40 MLOAD PUSH32 0x8C379A000000000000000000000000000000000000000000000000000000000 DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE PUSH1 0xB DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x6F6E6C79206F776E657221000000000000000000000000000000000000000000 DUP2 MSTORE POP PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST DUP1 PUSH1 0x3 PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP POP JUMP JUMPDEST PUSH1 0x3 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 JUMP JUMPDEST PUSH1 0x60 DUP1 PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 PUSH1 0x32 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x7B2275726C223A22687474703A2F2F31302E3231342E3234322E3232383A3830 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x39302F73656E64446174613F696E6465783D0000000000000000000000000000 DUP2 MSTORE POP SWAP1 POP PUSH1 0x60 PUSH2 0x624 DUP3 DUP6 PUSH2 0xAA0 JUMP JUMPDEST SWAP1 POP PUSH1 0x60 PUSH2 0x667 DUP3 PUSH1 0x40 DUP1 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 PUSH1 0x6 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x26646174613D0000000000000000000000000000000000000000000000000000 DUP2 MSTORE POP PUSH2 0xAA0 JUMP JUMPDEST SWAP1 POP PUSH1 0x60 PUSH2 0x675 DUP3 DUP7 PUSH2 0xAA0 JUMP JUMPDEST SWAP1 POP PUSH1 0x60 PUSH2 0x6B8 DUP3 PUSH1 0x40 DUP1 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 PUSH1 0x16 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x222C22726573706F6E7365506172616D73223A5B5D7D00000000000000000000 DUP2 MSTORE POP PUSH2 0xAA0 JUMP JUMPDEST SWAP1 POP PUSH1 0x60 DUP2 SWAP1 POP PUSH1 0x3 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH4 0x937A865E PUSH1 0x0 SLOAD PUSH1 0x0 ADDRESS DUP6 PUSH1 0x40 MLOAD DUP6 PUSH4 0xFFFFFFFF AND PUSH29 0x100000000000000000000000000000000000000000000000000000000 MUL DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP5 DUP2 MSTORE PUSH1 0x20 ADD DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH1 0x20 ADD DUP1 PUSH1 0x20 ADD DUP4 DUP2 SUB DUP4 MSTORE PUSH1 0x27 DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x6765744F726163656C526573706F6E736528627974657333322C75696E743634 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x2C62797465732900000000000000000000000000000000000000000000000000 DUP2 MSTORE POP PUSH1 0x40 ADD DUP4 DUP2 SUB DUP3 MSTORE DUP5 DUP2 DUP2 MLOAD DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP1 DUP4 DUP4 PUSH1 0x0 JUMPDEST DUP4 DUP2 LT ISZERO PUSH2 0x805 JUMPI DUP1 DUP3 ADD MLOAD DUP2 DUP5 ADD MSTORE PUSH1 0x20 DUP2 ADD SWAP1 POP PUSH2 0x7EA JUMP JUMPDEST POP POP POP POP SWAP1 POP SWAP1 DUP2 ADD SWAP1 PUSH1 0x1F AND DUP1 ISZERO PUSH2 0x832 JUMPI DUP1 DUP3 SUB DUP1 MLOAD PUSH1 0x1 DUP4 PUSH1 0x20 SUB PUSH2 0x100 EXP SUB NOT AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP JUMPDEST POP SWAP6 POP POP POP POP POP POP PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP9 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0x853 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0x867 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x20 DUP2 LT ISZERO PUSH2 0x87E JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP2 ADD SWAP1 DUP1 DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP POP POP POP POP POP POP POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH1 0x3 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x961 JUMPI PUSH1 0x40 MLOAD PUSH32 0x8C379A000000000000000000000000000000000000000000000000000000000 DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE PUSH1 0xC DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x6F6E6C79206F7261636C65210000000000000000000000000000000000000000 DUP2 MSTORE POP PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST PUSH32 0x4F367325E187D31D7E593B3A97CFBA9B4419B5B46D5A9C4623E911CE3C32EB6B DUP6 DUP6 DUP6 DUP6 PUSH1 0x40 MLOAD DUP1 DUP6 DUP2 MSTORE PUSH1 0x20 ADD DUP5 PUSH8 0xFFFFFFFFFFFFFFFF AND PUSH8 0xFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE DUP5 DUP5 DUP3 DUP2 DUP2 MSTORE PUSH1 0x20 ADD SWAP3 POP DUP1 DUP3 DUP5 CALLDATACOPY PUSH1 0x0 DUP2 DUP5 ADD MSTORE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND SWAP1 POP DUP1 DUP4 ADD SWAP3 POP POP POP SWAP6 POP POP POP POP POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG1 DUP3 DUP3 PUSH1 0x2 SWAP2 SWAP1 PUSH2 0x9F9 SWAP3 SWAP2 SWAP1 PUSH2 0xC73 JUMP JUMPDEST POP SWAP5 SWAP4 POP POP POP POP JUMP JUMPDEST PUSH1 0x2 DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV DUP1 PUSH1 0x1F ADD PUSH1 0x20 DUP1 SWAP2 DIV MUL PUSH1 0x20 ADD PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 SWAP3 SWAP2 SWAP1 DUP2 DUP2 MSTORE PUSH1 0x20 ADD DUP3 DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV DUP1 ISZERO PUSH2 0xA98 JUMPI DUP1 PUSH1 0x1F LT PUSH2 0xA6D JUMPI PUSH2 0x100 DUP1 DUP4 SLOAD DIV MUL DUP4 MSTORE SWAP2 PUSH1 0x20 ADD SWAP2 PUSH2 0xA98 JUMP JUMPDEST DUP3 ADD SWAP2 SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 JUMPDEST DUP2 SLOAD DUP2 MSTORE SWAP1 PUSH1 0x1 ADD SWAP1 PUSH1 0x20 ADD DUP1 DUP4 GT PUSH2 0xA7B JUMPI DUP3 SWAP1 SUB PUSH1 0x1F AND DUP3 ADD SWAP2 JUMPDEST POP POP POP POP POP DUP2 JUMP JUMPDEST PUSH1 0x60 DUP1 DUP4 SWAP1 POP PUSH1 0x60 DUP4 SWAP1 POP PUSH1 0x60 DUP2 MLOAD DUP4 MLOAD ADD PUSH1 0x40 MLOAD SWAP1 DUP1 DUP3 MSTORE DUP1 PUSH1 0x1F ADD PUSH1 0x1F NOT AND PUSH1 0x20 ADD DUP3 ADD PUSH1 0x40 MSTORE DUP1 ISZERO PUSH2 0xAE4 JUMPI DUP2 PUSH1 0x20 ADD PUSH1 0x1 DUP3 MUL DUP1 CODESIZE DUP4 CODECOPY DUP1 DUP3 ADD SWAP2 POP POP SWAP1 POP JUMPDEST POP SWAP1 POP PUSH1 0x60 DUP2 SWAP1 POP PUSH1 0x0 DUP1 SWAP1 POP PUSH1 0x0 DUP1 SWAP1 POP JUMPDEST DUP6 MLOAD DUP2 LT ISZERO PUSH2 0xBAA JUMPI DUP6 DUP2 DUP2 MLOAD DUP2 LT ISZERO ISZERO PUSH2 0xB0E JUMPI INVALID JUMPDEST SWAP1 PUSH1 0x20 ADD ADD MLOAD PUSH32 0x100000000000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH32 0x100000000000000000000000000000000000000000000000000000000000000 MUL DUP4 DUP4 DUP1 PUSH1 0x1 ADD SWAP5 POP DUP2 MLOAD DUP2 LT ISZERO ISZERO PUSH2 0xB6D JUMPI INVALID JUMPDEST SWAP1 PUSH1 0x20 ADD ADD SWAP1 PUSH31 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF NOT AND SWAP1 DUP2 PUSH1 0x0 BYTE SWAP1 MSTORE8 POP DUP1 DUP1 PUSH1 0x1 ADD SWAP2 POP POP PUSH2 0xAF7 JUMP JUMPDEST POP PUSH1 0x0 DUP1 SWAP1 POP JUMPDEST DUP5 MLOAD DUP2 LT ISZERO PUSH2 0xC64 JUMPI DUP5 DUP2 DUP2 MLOAD DUP2 LT ISZERO ISZERO PUSH2 0xBC8 JUMPI INVALID JUMPDEST SWAP1 PUSH1 0x20 ADD ADD MLOAD PUSH32 0x100000000000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH32 0x100000000000000000000000000000000000000000000000000000000000000 MUL DUP4 DUP4 DUP1 PUSH1 0x1 ADD SWAP5 POP DUP2 MLOAD DUP2 LT ISZERO ISZERO PUSH2 0xC27 JUMPI INVALID JUMPDEST SWAP1 PUSH1 0x20 ADD ADD SWAP1 PUSH31 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF NOT AND SWAP1 DUP2 PUSH1 0x0 BYTE SWAP1 MSTORE8 POP DUP1 DUP1 PUSH1 0x1 ADD SWAP2 POP POP PUSH2 0xBB1 JUMP JUMPDEST POP DUP3 SWAP6 POP POP POP POP POP POP SWAP3 SWAP2 POP POP JUMP JUMPDEST DUP3 DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x1F ADD PUSH1 0x20 SWAP1 DIV DUP2 ADD SWAP3 DUP3 PUSH1 0x1F LT PUSH2 0xCB4 JUMPI DUP1 CALLDATALOAD PUSH1 0xFF NOT AND DUP4 DUP1 ADD OR DUP6 SSTORE PUSH2 0xCE2 JUMP JUMPDEST DUP3 DUP1 ADD PUSH1 0x1 ADD DUP6 SSTORE DUP3 ISZERO PUSH2 0xCE2 JUMPI SWAP2 DUP3 ADD JUMPDEST DUP3 DUP2 GT ISZERO PUSH2 0xCE1 JUMPI DUP3 CALLDATALOAD DUP3 SSTORE SWAP2 PUSH1 0x20 ADD SWAP2 SWAP1 PUSH1 0x1 ADD SWAP1 PUSH2 0xCC6 JUMP JUMPDEST JUMPDEST POP SWAP1 POP PUSH2 0xCEF SWAP2 SWAP1 PUSH2 0xCF3 JUMP JUMPDEST POP SWAP1 JUMP JUMPDEST PUSH2 0xD15 SWAP2 SWAP1 JUMPDEST DUP1 DUP3 GT ISZERO PUSH2 0xD11 JUMPI PUSH1 0x0 DUP2 PUSH1 0x0 SWAP1 SSTORE POP PUSH1 0x1 ADD PUSH2 0xCF9 JUMP JUMPDEST POP SWAP1 JUMP JUMPDEST SWAP1 JUMP INVALID LOG1 PUSH6 0x627A7A723058 KECCAK256 JUMPDEST SIGNEXTEND 0xe7 DUP11 DUP15 JUMPI PUSH17 0x9006249EFCA1EAD840E9C9297A034B353 PUSH27 0xEDC76EAFED6B8B0029000000000000000000000000000000000000"

// DeploySendData deploys a new Ethereum contract, binding an instance of SendData to it.
func DeploySendData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SendData, error) {
	parsed, err := abi.JSON(strings.NewReader(SendDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SendDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SendData{SendDataCaller: SendDataCaller{contract: contract}, SendDataTransactor: SendDataTransactor{contract: contract}, SendDataFilterer: SendDataFilterer{contract: contract}}, nil
}

// SendData is an auto generated Go binding around an Ethereum contract.
type SendData struct {
	SendDataCaller     // Read-only binding to the contract
	SendDataTransactor // Write-only binding to the contract
	SendDataFilterer   // Log filterer for contract events
}

// SendDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type SendDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SendDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SendDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SendDataSession struct {
	Contract     *SendData         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SendDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SendDataCallerSession struct {
	Contract *SendDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SendDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SendDataTransactorSession struct {
	Contract     *SendDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SendDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type SendDataRaw struct {
	Contract *SendData // Generic contract binding to access the raw methods on
}

// SendDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SendDataCallerRaw struct {
	Contract *SendDataCaller // Generic read-only contract binding to access the raw methods on
}

// SendDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SendDataTransactorRaw struct {
	Contract *SendDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSendData creates a new instance of SendData, bound to a specific deployed contract.
func NewSendData(address common.Address, backend bind.ContractBackend) (*SendData, error) {
	contract, err := bindSendData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SendData{SendDataCaller: SendDataCaller{contract: contract}, SendDataTransactor: SendDataTransactor{contract: contract}, SendDataFilterer: SendDataFilterer{contract: contract}}, nil
}

// NewSendDataCaller creates a new read-only instance of SendData, bound to a specific deployed contract.
func NewSendDataCaller(address common.Address, caller bind.ContractCaller) (*SendDataCaller, error) {
	contract, err := bindSendData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SendDataCaller{contract: contract}, nil
}

// NewSendDataTransactor creates a new write-only instance of SendData, bound to a specific deployed contract.
func NewSendDataTransactor(address common.Address, transactor bind.ContractTransactor) (*SendDataTransactor, error) {
	contract, err := bindSendData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SendDataTransactor{contract: contract}, nil
}

// NewSendDataFilterer creates a new log filterer instance of SendData, bound to a specific deployed contract.
func NewSendDataFilterer(address common.Address, filterer bind.ContractFilterer) (*SendDataFilterer, error) {
	contract, err := bindSendData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SendDataFilterer{contract: contract}, nil
}

// bindSendData binds a generic wrapper to an already deployed contract.
func bindSendData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SendDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendData *SendDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SendData.Contract.SendDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendData *SendDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendData.Contract.SendDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendData *SendDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendData.Contract.SendDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendData *SendDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SendData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendData *SendDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendData *SendDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendData.Contract.contract.Transact(opts, method, params...)
}

// ORACLEFEE is a free data retrieval call binding the contract method 0x31d8a985.
//
// Solidity: function ORACLE_FEE() view returns(uint256)
func (_SendData *SendDataCaller) ORACLEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SendData.contract.Call(opts, &out, "ORACLE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ORACLEFEE is a free data retrieval call binding the contract method 0x31d8a985.
//
// Solidity: function ORACLE_FEE() view returns(uint256)
func (_SendData *SendDataSession) ORACLEFEE() (*big.Int, error) {
	return _SendData.Contract.ORACLEFEE(&_SendData.CallOpts)
}

// ORACLEFEE is a free data retrieval call binding the contract method 0x31d8a985.
//
// Solidity: function ORACLE_FEE() view returns(uint256)
func (_SendData *SendDataCallerSession) ORACLEFEE() (*big.Int, error) {
	return _SendData.Contract.ORACLEFEE(&_SendData.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_SendData *SendDataCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SendData.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_SendData *SendDataSession) Oracle() (common.Address, error) {
	return _SendData.Contract.Oracle(&_SendData.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_SendData *SendDataCallerSession) Oracle() (common.Address, error) {
	return _SendData.Contract.Oracle(&_SendData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SendData *SendDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SendData.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SendData *SendDataSession) Owner() (common.Address, error) {
	return _SendData.Contract.Owner(&_SendData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SendData *SendDataCallerSession) Owner() (common.Address, error) {
	return _SendData.Contract.Owner(&_SendData.CallOpts)
}

// Resp is a free data retrieval call binding the contract method 0xe9f5e299.
//
// Solidity: function resp() view returns(bytes)
func (_SendData *SendDataCaller) Resp(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SendData.contract.Call(opts, &out, "resp")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Resp is a free data retrieval call binding the contract method 0xe9f5e299.
//
// Solidity: function resp() view returns(bytes)
func (_SendData *SendDataSession) Resp() ([]byte, error) {
	return _SendData.Contract.Resp(&_SendData.CallOpts)
}

// Resp is a free data retrieval call binding the contract method 0xe9f5e299.
//
// Solidity: function resp() view returns(bytes)
func (_SendData *SendDataCallerSession) Resp() ([]byte, error) {
	return _SendData.Contract.Resp(&_SendData.CallOpts)
}

// GetOracelResponse is a paid mutator transaction binding the contract method 0xc265028d.
//
// Solidity: function getOracelResponse(bytes32 _reqId, uint64 _stateCode, bytes _response) returns(bool)
func (_SendData *SendDataTransactor) GetOracelResponse(opts *bind.TransactOpts, _reqId [32]byte, _stateCode uint64, _response []byte) (*types.Transaction, error) {
	return _SendData.contract.Transact(opts, "getOracelResponse", _reqId, _stateCode, _response)
}

// GetOracelResponse is a paid mutator transaction binding the contract method 0xc265028d.
//
// Solidity: function getOracelResponse(bytes32 _reqId, uint64 _stateCode, bytes _response) returns(bool)
func (_SendData *SendDataSession) GetOracelResponse(_reqId [32]byte, _stateCode uint64, _response []byte) (*types.Transaction, error) {
	return _SendData.Contract.GetOracelResponse(&_SendData.TransactOpts, _reqId, _stateCode, _response)
}

// GetOracelResponse is a paid mutator transaction binding the contract method 0xc265028d.
//
// Solidity: function getOracelResponse(bytes32 _reqId, uint64 _stateCode, bytes _response) returns(bool)
func (_SendData *SendDataTransactorSession) GetOracelResponse(_reqId [32]byte, _stateCode uint64, _response []byte) (*types.Transaction, error) {
	return _SendData.Contract.GetOracelResponse(&_SendData.TransactOpts, _reqId, _stateCode, _response)
}

// Send is a paid mutator transaction binding the contract method 0xbd6de11c.
//
// Solidity: function send(string _index, string _calldata) returns()
func (_SendData *SendDataTransactor) Send(opts *bind.TransactOpts, _index string, _calldata string) (*types.Transaction, error) {
	return _SendData.contract.Transact(opts, "send", _index, _calldata)
}

// Send is a paid mutator transaction binding the contract method 0xbd6de11c.
//
// Solidity: function send(string _index, string _calldata) returns()
func (_SendData *SendDataSession) Send(_index string, _calldata string) (*types.Transaction, error) {
	return _SendData.Contract.Send(&_SendData.TransactOpts, _index, _calldata)
}

// Send is a paid mutator transaction binding the contract method 0xbd6de11c.
//
// Solidity: function send(string _index, string _calldata) returns()
func (_SendData *SendDataTransactorSession) Send(_index string, _calldata string) (*types.Transaction, error) {
	return _SendData.Contract.Send(&_SendData.TransactOpts, _index, _calldata)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_SendData *SendDataTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _SendData.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_SendData *SendDataSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _SendData.Contract.SetOracle(&_SendData.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_SendData *SendDataTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _SendData.Contract.SetOracle(&_SendData.TransactOpts, _oracle)
}

// SendDataDistributeFundsIterator is returned from FilterDistributeFunds and is used to iterate over the raw logs and unpacked data for DistributeFunds events raised by the SendData contract.
type SendDataDistributeFundsIterator struct {
	Event *SendDataDistributeFunds // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SendDataDistributeFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SendDataDistributeFunds)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SendDataDistributeFunds)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SendDataDistributeFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SendDataDistributeFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SendDataDistributeFunds represents a DistributeFunds event raised by the SendData contract.
type SendDataDistributeFunds struct {
	Winner     common.Address
	Funds      *big.Int
	RoundTimes *big.Int
	PlayId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDistributeFunds is a free log retrieval operation binding the contract event 0x28a1f6b9098b70e9b325c9ad9e2bb1fd2289c6facabdc9a3f8ddade40a815aef.
//
// Solidity: event DistributeFunds(address _winner, uint256 _funds, uint256 _roundTimes, uint256 playId)
func (_SendData *SendDataFilterer) FilterDistributeFunds(opts *bind.FilterOpts) (*SendDataDistributeFundsIterator, error) {

	logs, sub, err := _SendData.contract.FilterLogs(opts, "DistributeFunds")
	if err != nil {
		return nil, err
	}
	return &SendDataDistributeFundsIterator{contract: _SendData.contract, event: "DistributeFunds", logs: logs, sub: sub}, nil
}

// WatchDistributeFunds is a free log subscription operation binding the contract event 0x28a1f6b9098b70e9b325c9ad9e2bb1fd2289c6facabdc9a3f8ddade40a815aef.
//
// Solidity: event DistributeFunds(address _winner, uint256 _funds, uint256 _roundTimes, uint256 playId)
func (_SendData *SendDataFilterer) WatchDistributeFunds(opts *bind.WatchOpts, sink chan<- *SendDataDistributeFunds) (event.Subscription, error) {

	logs, sub, err := _SendData.contract.WatchLogs(opts, "DistributeFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SendDataDistributeFunds)
				if err := _SendData.contract.UnpackLog(event, "DistributeFunds", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributeFunds is a log parse operation binding the contract event 0x28a1f6b9098b70e9b325c9ad9e2bb1fd2289c6facabdc9a3f8ddade40a815aef.
//
// Solidity: event DistributeFunds(address _winner, uint256 _funds, uint256 _roundTimes, uint256 playId)
func (_SendData *SendDataFilterer) ParseDistributeFunds(log types.Log) (*SendDataDistributeFunds, error) {
	event := new(SendDataDistributeFunds)
	if err := _SendData.contract.UnpackLog(event, "DistributeFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SendDataOracleResponseIterator is returned from FilterOracleResponse and is used to iterate over the raw logs and unpacked data for OracleResponse events raised by the SendData contract.
type SendDataOracleResponseIterator struct {
	Event *SendDataOracleResponse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SendDataOracleResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SendDataOracleResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SendDataOracleResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SendDataOracleResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SendDataOracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SendDataOracleResponse represents a OracleResponse event raised by the SendData contract.
type SendDataOracleResponse struct {
	ReqId     [32]byte
	StateCode uint64
	RandomNum []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleResponse is a free log retrieval operation binding the contract event 0x4f367325e187d31d7e593b3a97cfba9b4419b5b46d5a9c4623e911ce3c32eb6b.
//
// Solidity: event OracleResponse(bytes32 _reqId, uint64 _stateCode, bytes _randomNum)
func (_SendData *SendDataFilterer) FilterOracleResponse(opts *bind.FilterOpts) (*SendDataOracleResponseIterator, error) {

	logs, sub, err := _SendData.contract.FilterLogs(opts, "OracleResponse")
	if err != nil {
		return nil, err
	}
	return &SendDataOracleResponseIterator{contract: _SendData.contract, event: "OracleResponse", logs: logs, sub: sub}, nil
}

// WatchOracleResponse is a free log subscription operation binding the contract event 0x4f367325e187d31d7e593b3a97cfba9b4419b5b46d5a9c4623e911ce3c32eb6b.
//
// Solidity: event OracleResponse(bytes32 _reqId, uint64 _stateCode, bytes _randomNum)
func (_SendData *SendDataFilterer) WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *SendDataOracleResponse) (event.Subscription, error) {

	logs, sub, err := _SendData.contract.WatchLogs(opts, "OracleResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SendDataOracleResponse)
				if err := _SendData.contract.UnpackLog(event, "OracleResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleResponse is a log parse operation binding the contract event 0x4f367325e187d31d7e593b3a97cfba9b4419b5b46d5a9c4623e911ce3c32eb6b.
//
// Solidity: event OracleResponse(bytes32 _reqId, uint64 _stateCode, bytes _randomNum)
func (_SendData *SendDataFilterer) ParseOracleResponse(log types.Log) (*SendDataOracleResponse, error) {
	event := new(SendDataOracleResponse)
	if err := _SendData.contract.UnpackLog(event, "OracleResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
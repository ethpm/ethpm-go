package ethpm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestAddDependency(t *testing.T) {
	pm := &PackageManifest{}
	name := `ethereum-libraries-basic-math`
	uri := `https://github.com/modular-network/ethereum-libraries-basic-math/commit/618d432d34cb294ff0ec8c9825348e592c0a9cad`

	pm.AddDependency(name, uri)
	got := pm.BuildDependencies["ethereum-libraries-basic-math"]

	if got != uri {
		t.Fatalf("Got '%v', expected '%v'", got, uri)
	}
}

func ExamplePackageManifest_AddDependency() {
	pm := &PackageManifest{}
	name := `ethereum-libraries-basic-math`
	uri := `https://github.com/modular-network/ethereum-libraries-basic-math/commit/618d432d34cb294ff0ec8c9825348e592c0a9cad`

	pm.AddDependency(name, uri)
	fmt.Println(pm.BuildDependencies["ethereum-libraries-basic-math"])
	// Output: https://github.com/modular-network/ethereum-libraries-basic-math/commit/618d432d34cb294ff0ec8c9825348e592c0a9cad
}

func TestAddContractType(t *testing.T) {
	pm := &PackageManifest{}
	settingsjson := `{
    "optimizer": {},
    "outputSelection": {
      "*": {
        "*": ["evm.bytecode", "evm.deployedBytecode"]
      }
    }
  }`
	outputjson := `{
	  "contracts": {
	    "BasicMathLib.sol": {
	      "BasicMathLib": {
	        "abi": [{
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "times",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "res",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }, {
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "plus",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "res",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }, {
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "dividedBy",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "i",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }, {
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "minus",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "res",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }],
	        "devdoc": {
	          "author": "Modular, Inc * version 1.2.7 Copyright (c) 2017 Modular, Inc The MIT License (MIT) https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE * The Basic Math Library is inspired by the Safe Math library written by OpenZeppelin at https://github.com/OpenZeppelin/zeppelin-solidity/ . Modular provides smart contract services and security reviews for contract deployments in addition to working on open source projects in the Ethereum community. Our purpose is to test, document, and deploy reusable code onto the blockchain and improve both security and usability. We also educate non-profits, schools, and other community members about the application of blockchain technology. For further information: modular.network, openzeppelin.org * THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.",
	          "methods": {
	            "dividedBy(uint256,uint256)": {
	              "details": "Divides two numbers but checks for 0 in the divisor first. Does not throw.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if 'b' is 0res The quotient of a and b, or 0 if 'b' is 0"
	            },
	            "minus(uint256,uint256)": {
	              "details": "Subtracts two numbers and checks for underflow before returning. Does not throw but rather logs an Err event if there is underflow.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if there is underflowres The difference between a and b, or 0 if there is underflow"
	            },
	            "plus(uint256,uint256)": {
	              "details": "Adds two numbers and checks for overflow before returning. Does not throw.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if there is overflowres The sum of a and b, or 0 if there is overflow"
	            },
	            "times(uint256,uint256)": {
	              "details": "Multiplies two numbers and checks for overflow before returning. Does not throw.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if there is overflowres The product of a and b, or 0 if there is overflow"
	            }
	          },
	          "title": "Basic Math Library"
	        },
	        "evm": {
	          "bytecode": {
	            "linkReferences": {},
	            "object": "610198610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029",
	            "opcodes": "PUSH2 0x198 PUSH2 0x30 PUSH1 0xB DUP3 DUP3 DUP3 CODECOPY DUP1 MLOAD PUSH1 0x0 BYTE PUSH1 0x73 EQ PUSH1 0x0 DUP2 EQ PUSH2 0x20 JUMPI PUSH2 0x22 JUMP JUMPDEST INVALID JUMPDEST POP ADDRESS PUSH1 0x0 MSTORE PUSH1 0x73 DUP2 MSTORE8 DUP3 DUP2 RETURN STOP PUSH20 0x0 ADDRESS EQ PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x78 JUMPI PUSH4 0xFFFFFFFF PUSH29 0x100000000000000000000000000000000000000000000000000000000 PUSH1 0x0 CALLDATALOAD DIV AND PUSH4 0x1D3B9EDF DUP2 EQ PUSH2 0x7D JUMPI DUP1 PUSH4 0x66098D4F EQ PUSH2 0xA6 JUMPI DUP1 PUSH4 0xE39BBF68 EQ PUSH2 0xB4 JUMPI DUP1 PUSH4 0xF4F3BDC1 EQ PUSH2 0xC2 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xD0 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD SWAP3 ISZERO ISZERO DUP4 MSTORE PUSH1 0x20 DUP4 ADD SWAP2 SWAP1 SWAP2 MSTORE DUP1 MLOAD SWAP2 DUP3 SWAP1 SUB ADD SWAP1 RETURN JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xF9 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x116 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x14C JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 MUL DUP3 ISZERO DUP4 DUP3 DIV DUP6 EQ OR DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x1 SWAP3 POP PUSH1 0x0 SWAP2 POP JUMPDEST POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 ADD DUP3 DUP2 SUB DUP5 EQ DUP4 DUP3 GT DUP3 DUP6 EQ OR AND DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x0 DUP1 DUP1 DUP4 ISZERO DUP1 ISZERO PUSH2 0x12F JUMPI PUSH1 0x1 SWAP4 POP PUSH1 0x0 SWAP3 POP PUSH2 0x143 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP6 DUP8 DIV PUSH1 0x20 SWAP1 SWAP2 ADD DUP2 SWAP1 MSTORE SWAP3 POP DUP3 SWAP2 POP JUMPDEST POP POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 DUP4 SUB DUP1 DUP4 ADD DUP5 EQ DUP5 DUP3 LT DUP3 DUP7 EQ OR AND PUSH1 0x1 EQ DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0x26 LOG0 DIV 0xc4 DELEGATECALL SMOD SIGNEXTEND SLT MSTORE8 0xb6 0xd1 SSTORE NOT PUSH21 0x7659F677073157D481638DB922F9ADBC6EC002900 ",
	            "sourceMap": "1353:2288:0:-;;132:2:-1;166:7;155:9;146:7;137:37;252:7;246:14;243:1;238:23;232:4;229:33;270:1;265:20;;;;222:63;;265:20;274:9;222:63;;298:9;295:1;288:20;328:4;319:7;311:22;352:7;343;336:24"
	          },
	          "deployedBytecode": {
	            "linkReferences": {},
	            "object": "73000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029",
	            "opcodes": "PUSH20 0x0 ADDRESS EQ PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x78 JUMPI PUSH4 0xFFFFFFFF PUSH29 0x100000000000000000000000000000000000000000000000000000000 PUSH1 0x0 CALLDATALOAD DIV AND PUSH4 0x1D3B9EDF DUP2 EQ PUSH2 0x7D JUMPI DUP1 PUSH4 0x66098D4F EQ PUSH2 0xA6 JUMPI DUP1 PUSH4 0xE39BBF68 EQ PUSH2 0xB4 JUMPI DUP1 PUSH4 0xF4F3BDC1 EQ PUSH2 0xC2 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xD0 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD SWAP3 ISZERO ISZERO DUP4 MSTORE PUSH1 0x20 DUP4 ADD SWAP2 SWAP1 SWAP2 MSTORE DUP1 MLOAD SWAP2 DUP3 SWAP1 SUB ADD SWAP1 RETURN JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xF9 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x116 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x14C JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 MUL DUP3 ISZERO DUP4 DUP3 DIV DUP6 EQ OR DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x1 SWAP3 POP PUSH1 0x0 SWAP2 POP JUMPDEST POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 ADD DUP3 DUP2 SUB DUP5 EQ DUP4 DUP3 GT DUP3 DUP6 EQ OR AND DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x0 DUP1 DUP1 DUP4 ISZERO DUP1 ISZERO PUSH2 0x12F JUMPI PUSH1 0x1 SWAP4 POP PUSH1 0x0 SWAP3 POP PUSH2 0x143 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP6 DUP8 DIV PUSH1 0x20 SWAP1 SWAP2 ADD DUP2 SWAP1 MSTORE SWAP3 POP DUP3 SWAP2 POP JUMPDEST POP POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 DUP4 SUB DUP1 DUP4 ADD DUP5 EQ DUP5 DUP3 LT DUP3 DUP7 EQ OR AND PUSH1 0x1 EQ DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0x26 LOG0 DIV 0xc4 DELEGATECALL SMOD SIGNEXTEND SLT MSTORE8 0xb6 0xd1 SSTORE NOT PUSH21 0x7659F677073157D481638DB922F9ADBC6EC002900 ",
	            "sourceMap": "1353:2288:0:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1664:230;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;2790:245;;;;;;;;2161:349;;;;;;;;3386:253;;;;;;;;1664:230;1722:8;1773;;;1798:9;;1812:10;;;1809:17;;1795:32;1834:50;;;;1788:96;;1834:50;1858:1;1851:8;;1875:1;1868:8;;1788:96;;1758:132;;;;;:::o;2790:245::-;2847:8;2899;;;2928:10;;;2925:17;;2947:9;;;2957;;;2944:23;2921:47;2975:50;;;;2914:111;;2161:349;2223:8;;;2289:9;;2305:139;;;;2476:1;2469:8;;2491:1;2486:6;;2282:218;;2305:139;2363:4;2357:11;2329:8;;;2392:4;2384:13;;;2377:25;;;2329:8;-1:-1:-1;2329:8:0;;-1:-1:-1;2282:218:0;;2274:232;;;;;;:::o;3386:253::-;3444:8;3495;;;3527:10;;;3524:17;;3546:9;;;3557;;;3543:24;3520:48;3570:1;3517:55;3579:50;;;;3510:119;"
	          }
	        },
	        "userdoc": {
	          "methods": {}
	        }
	      }
	    }
	  },
	  "sources": {
	    "BasicMathLib.sol": {
	      "id": 0
	    }
	  }
	}`
	err := pm.AddContractType("solc", settingsjson, outputjson, "BasicMathLib")
	if err != nil {
		t.Fatalf("Got '%v', expected '<nil>'", err)
	}
}

func ExamplePackageManifest_AddContractType() {
	pm := &PackageManifest{}
	settingsjson := `{
    "optimizer": {},
    "outputSelection": {
      "*": {
        "*": ["evm.bytecode", "evm.deployedBytecode"]
      }
    }
  }`
	outputjson := `{
	  "contracts": {
	    "BasicMathLib.sol": {
	      "BasicMathLib": {
	        "abi": [{
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "times",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "res",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }, {
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "plus",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "res",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }, {
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "dividedBy",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "i",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }, {
	          "constant": true,
	          "inputs": [{
	            "name": "a",
	            "type": "uint256"
	          }, {
	            "name": "b",
	            "type": "uint256"
	          }],
	          "name": "minus",
	          "outputs": [{
	            "name": "err",
	            "type": "bool"
	          }, {
	            "name": "res",
	            "type": "uint256"
	          }],
	          "payable": false,
	          "stateMutability": "pure",
	          "type": "function"
	        }],
	        "devdoc": {
	          "author": "Modular, Inc * version 1.2.7 Copyright (c) 2017 Modular, Inc The MIT License (MIT) https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE * The Basic Math Library is inspired by the Safe Math library written by OpenZeppelin at https://github.com/OpenZeppelin/zeppelin-solidity/ . Modular provides smart contract services and security reviews for contract deployments in addition to working on open source projects in the Ethereum community. Our purpose is to test, document, and deploy reusable code onto the blockchain and improve both security and usability. We also educate non-profits, schools, and other community members about the application of blockchain technology. For further information: modular.network, openzeppelin.org * THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.",
	          "methods": {
	            "dividedBy(uint256,uint256)": {
	              "details": "Divides two numbers but checks for 0 in the divisor first. Does not throw.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if 'b' is 0res The quotient of a and b, or 0 if 'b' is 0"
	            },
	            "minus(uint256,uint256)": {
	              "details": "Subtracts two numbers and checks for underflow before returning. Does not throw but rather logs an Err event if there is underflow.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if there is underflowres The difference between a and b, or 0 if there is underflow"
	            },
	            "plus(uint256,uint256)": {
	              "details": "Adds two numbers and checks for overflow before returning. Does not throw.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if there is overflowres The sum of a and b, or 0 if there is overflow"
	            },
	            "times(uint256,uint256)": {
	              "details": "Multiplies two numbers and checks for overflow before returning. Does not throw.",
	              "params": {
	                "a": "First number",
	                "b": "Second number"
	              },
	              "return": "err False normally, or true if there is overflowres The product of a and b, or 0 if there is overflow"
	            }
	          },
	          "title": "Basic Math Library"
	        },
	        "evm": {
	          "bytecode": {
	            "linkReferences": {},
	            "object": "610198610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029",
	            "opcodes": "PUSH2 0x198 PUSH2 0x30 PUSH1 0xB DUP3 DUP3 DUP3 CODECOPY DUP1 MLOAD PUSH1 0x0 BYTE PUSH1 0x73 EQ PUSH1 0x0 DUP2 EQ PUSH2 0x20 JUMPI PUSH2 0x22 JUMP JUMPDEST INVALID JUMPDEST POP ADDRESS PUSH1 0x0 MSTORE PUSH1 0x73 DUP2 MSTORE8 DUP3 DUP2 RETURN STOP PUSH20 0x0 ADDRESS EQ PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x78 JUMPI PUSH4 0xFFFFFFFF PUSH29 0x100000000000000000000000000000000000000000000000000000000 PUSH1 0x0 CALLDATALOAD DIV AND PUSH4 0x1D3B9EDF DUP2 EQ PUSH2 0x7D JUMPI DUP1 PUSH4 0x66098D4F EQ PUSH2 0xA6 JUMPI DUP1 PUSH4 0xE39BBF68 EQ PUSH2 0xB4 JUMPI DUP1 PUSH4 0xF4F3BDC1 EQ PUSH2 0xC2 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xD0 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD SWAP3 ISZERO ISZERO DUP4 MSTORE PUSH1 0x20 DUP4 ADD SWAP2 SWAP1 SWAP2 MSTORE DUP1 MLOAD SWAP2 DUP3 SWAP1 SUB ADD SWAP1 RETURN JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xF9 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x116 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x14C JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 MUL DUP3 ISZERO DUP4 DUP3 DIV DUP6 EQ OR DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x1 SWAP3 POP PUSH1 0x0 SWAP2 POP JUMPDEST POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 ADD DUP3 DUP2 SUB DUP5 EQ DUP4 DUP3 GT DUP3 DUP6 EQ OR AND DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x0 DUP1 DUP1 DUP4 ISZERO DUP1 ISZERO PUSH2 0x12F JUMPI PUSH1 0x1 SWAP4 POP PUSH1 0x0 SWAP3 POP PUSH2 0x143 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP6 DUP8 DIV PUSH1 0x20 SWAP1 SWAP2 ADD DUP2 SWAP1 MSTORE SWAP3 POP DUP3 SWAP2 POP JUMPDEST POP POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 DUP4 SUB DUP1 DUP4 ADD DUP5 EQ DUP5 DUP3 LT DUP3 DUP7 EQ OR AND PUSH1 0x1 EQ DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0x26 LOG0 DIV 0xc4 DELEGATECALL SMOD SIGNEXTEND SLT MSTORE8 0xb6 0xd1 SSTORE NOT PUSH21 0x7659F677073157D481638DB922F9ADBC6EC002900 ",
	            "sourceMap": "1353:2288:0:-;;132:2:-1;166:7;155:9;146:7;137:37;252:7;246:14;243:1;238:23;232:4;229:33;270:1;265:20;;;;222:63;;265:20;274:9;222:63;;298:9;295:1;288:20;328:4;319:7;311:22;352:7;343;336:24"
	          },
	          "deployedBytecode": {
	            "linkReferences": {},
	            "object": "73000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029",
	            "opcodes": "PUSH20 0x0 ADDRESS EQ PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x78 JUMPI PUSH4 0xFFFFFFFF PUSH29 0x100000000000000000000000000000000000000000000000000000000 PUSH1 0x0 CALLDATALOAD DIV AND PUSH4 0x1D3B9EDF DUP2 EQ PUSH2 0x7D JUMPI DUP1 PUSH4 0x66098D4F EQ PUSH2 0xA6 JUMPI DUP1 PUSH4 0xE39BBF68 EQ PUSH2 0xB4 JUMPI DUP1 PUSH4 0xF4F3BDC1 EQ PUSH2 0xC2 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xD0 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD SWAP3 ISZERO ISZERO DUP4 MSTORE PUSH1 0x20 DUP4 ADD SWAP2 SWAP1 SWAP2 MSTORE DUP1 MLOAD SWAP2 DUP3 SWAP1 SUB ADD SWAP1 RETURN JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0xF9 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x116 JUMP JUMPDEST PUSH2 0x8B PUSH1 0x4 CALLDATALOAD PUSH1 0x24 CALLDATALOAD PUSH2 0x14C JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 MUL DUP3 ISZERO DUP4 DUP3 DIV DUP6 EQ OR DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x1 SWAP3 POP PUSH1 0x0 SWAP2 POP JUMPDEST POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 ADD DUP3 DUP2 SUB DUP5 EQ DUP4 DUP3 GT DUP3 DUP6 EQ OR AND DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP JUMPDEST PUSH1 0x0 DUP1 DUP1 DUP4 ISZERO DUP1 ISZERO PUSH2 0x12F JUMPI PUSH1 0x1 SWAP4 POP PUSH1 0x0 SWAP3 POP PUSH2 0x143 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP6 DUP8 DIV PUSH1 0x20 SWAP1 SWAP2 ADD DUP2 SWAP1 MSTORE SWAP3 POP DUP3 SWAP2 POP JUMPDEST POP POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 DUP4 SUB DUP1 DUP4 ADD DUP5 EQ DUP5 DUP3 LT DUP3 DUP7 EQ OR AND PUSH1 0x1 EQ DUP1 ISZERO PUSH2 0xE8 JUMPI PUSH2 0xF1 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0x26 LOG0 DIV 0xc4 DELEGATECALL SMOD SIGNEXTEND SLT MSTORE8 0xb6 0xd1 SSTORE NOT PUSH21 0x7659F677073157D481638DB922F9ADBC6EC002900 ",
	            "sourceMap": "1353:2288:0:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1664:230;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;2790:245;;;;;;;;2161:349;;;;;;;;3386:253;;;;;;;;1664:230;1722:8;1773;;;1798:9;;1812:10;;;1809:17;;1795:32;1834:50;;;;1788:96;;1834:50;1858:1;1851:8;;1875:1;1868:8;;1788:96;;1758:132;;;;;:::o;2790:245::-;2847:8;2899;;;2928:10;;;2925:17;;2947:9;;;2957;;;2944:23;2921:47;2975:50;;;;2914:111;;2161:349;2223:8;;;2289:9;;2305:139;;;;2476:1;2469:8;;2491:1;2486:6;;2282:218;;2305:139;2363:4;2357:11;2329:8;;;2392:4;2384:13;;;2377:25;;;2329:8;-1:-1:-1;2329:8:0;;-1:-1:-1;2282:218:0;;2274:232;;;;;;:::o;3386:253::-;3444:8;3495;;;3527:10;;;3524:17;;3546:9;;;3557;;;3543:24;3520:48;3570:1;3517:55;3579:50;;;;3510:119;"
	          }
	        },
	        "userdoc": {
	          "methods": {}
	        }
	      }
	    }
	  },
	  "sources": {
	    "BasicMathLib.sol": {
	      "id": 0
	    }
	  }
	}`
	if err := pm.AddContractType("solc", settingsjson, outputjson, "BasicMathLib"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(pm.ContractTypes["BasicMathLib"].Compiler.Name)
	// Output: solc
}

func TestSourceInliner(t *testing.T) {
	p := PackageManifest{}
	if err := p.SourceInliner("../../test/testdata", "", "sol"); err != nil {
		t.Fatal(err)
	}

	got := p.Sources["./BasicMathLib.sol"]
	want, _ := ioutil.ReadFile("../../test/testdata/BasicMathLib.sol")
	if got != string(want) {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func ExamplePackageManifest_SourceInliner() {
	p := PackageManifest{}
	if err := p.SourceInliner("path/to/sol/files", "", "sol"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Sources["./SomeContract.sol"])

	if err := p.SourceInliner("path/to/sol/files", "./contracts/", "sol"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Sources["./contracts/SomeContract.sol"])
}

func TestAddLocalPathForSource(t *testing.T) {
	p := PackageManifest{}
	if err := p.AddLocalPathForSource("../../test/testdata", "./contracts/", "sol"); err != nil {
		t.Fatal(err)
	}

	got := p.Sources["./contracts/BasicMathLib.sol"]
	want := "./contracts/BasicMathLib.sol"
	if got != string(want) {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func ExamplePackageManifest_AddLocalPathForSource() {
	p := PackageManifest{}
	if err := p.AddLocalPathForSource("../../test/testdata", "./contracts/", "sol"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Sources["./contracts/SomeContract.sol"])
}

func TestManifestValidate(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.ManifestVersion = "2"
	got = p.Validate()
	want = errors.New("PackageManifest:package_name returned error 'must provide a package name'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.PackageName = "alexs-package"
	p.ManifestVersion = " 2"
	got = p.Validate()
	want = errors.New("PackageManifest:manifest_version returned error 'manifest_version " +
		"should be 2, manifest_version is showing  2. Ensure there are no extra spaces or characters'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.ManifestVersion = "2 "
	got = p.Validate()
	want = errors.New("PackageManifest:manifest_version returned error 'manifest_version " +
		"should be 2, manifest_version is showing 2 . Ensure there are no extra spaces or characters'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.ManifestVersion = "2"
	got = p.Validate()
	want = errors.New("PackageManifest:version returned error 'must provide a version number'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.Version = "1.0.0"
	got = p.Validate()
	if got != nil {
		t.Fatalf("Got '%v', expected '<nil>'", got)
	}
}

func TestCheckSources(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.Sources = make(map[string]string)
	p.Sources["./hello"] = "/Nick"
	got = checkSources(p.Sources)
	want = errors.New("Source with key './hello' and location value '/Nick' does not exist or is unreachable. " +
		"Please check the url or filepath and fix or consider contacting the maintainer.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	delete(p.Sources, "./hello")
	p.Sources["/not-valid-other-chris"] = "https://github.com/modular-network"
	got = checkSources(p.Sources)
	want = errors.New("Invalid path for source key '/not-valid-other-chris'. Please make this a relative path in accordance " +
		"with the spec found here https://ethpm.github.io/ethpm-spec/package-spec.html#sources-sources.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	delete(p.Sources, "/not-valid-other-chris")
	p.Sources["./contracts"] = "https://github.com/modular-network/ethereum-libraries"
	got = checkSources(p.Sources)
	if got != nil {
		t.Fatalf("Got '%v', expected '<nil>'", got)
	}
}

package ethcontract

import (
	"errors"
	"log"
	"testing"

	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
)

func TestBuild(t *testing.T) {
	ct := ContractType{}
	js := `{"abi":[{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"times","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"plus","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"dividedBy","outputs":[{"name":"err","type":"bool"},{"name":"i","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"minus","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"}],"devdoc":{"author":"Modular, Inc * version 1.2.7 Copyright (c) 2017 Modular, Inc The MIT License (MIT) https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE * The Basic Math Library is inspired by the Safe Math library written by OpenZeppelin at https://github.com/OpenZeppelin/zeppelin-solidity/ . Modular provides smart contract services and security reviews for contract deployments in addition to working on open source projects in the Ethereum community. Our purpose is to test, document, and deploy reusable code onto the blockchain and improve both security and usability. We also educate non-profits, schools, and other community members about the application of blockchain technology. For further information: modular.network, openzeppelin.org * THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.","methods":{"dividedBy(uint256,uint256)":{"details":"Divides two numbers but checks for 0 in the divisor first. Does not throw.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if b is 0res The quotient of a and b, or 0 if b is 0"},"minus(uint256,uint256)":{"details":"Subtracts two numbers and checks for underflow before returning. Does not throw but rather logs an Err event if there is underflow.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if there is underflowres The difference between a and b, or 0 if there is underflow"},"plus(uint256,uint256)":{"details":"Adds two numbers and checks for overflow before returning. Does not throw.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if there is overflowres The sum of a and b, or 0 if there is overflow"},"times(uint256,uint256)":{"details":"Multiplies two numbers and checks for overflow before returning. Does not throw.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if there is overflowres The product of a and b, or 0 if there is overflow"}},"title":"Basic Math Library"},"evm":{"bytecode":{"linkReferences":{},"object":"610198610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029"},"deployedBytecode":{"linkReferences":{},"object":"73000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029"}},"userdoc":{"methods":{}}}`
	s := `{ "optimizer": { "enabled": true, "runs": 200 }, "outputSelection": { "*": { "*": ["abi", "evm.bytecode", "evm.deployedBytecode", "devdoc", "userdoc"] } } }`
	got := ct.Build("solc", s, js)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}

func ExampleContractType() {
	ct := ContractType{}
	js := `{"abi":[{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"times","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"plus","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"dividedBy","outputs":[{"name":"err","type":"bool"},{"name":"i","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"minus","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"}],"devdoc":{"author":"Modular, Inc * version 1.2.7 Copyright (c) 2017 Modular, Inc The MIT License (MIT) https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE * The Basic Math Library is inspired by the Safe Math library written by OpenZeppelin at https://github.com/OpenZeppelin/zeppelin-solidity/ . Modular provides smart contract services and security reviews for contract deployments in addition to working on open source projects in the Ethereum community. Our purpose is to test, document, and deploy reusable code onto the blockchain and improve both security and usability. We also educate non-profits, schools, and other community members about the application of blockchain technology. For further information: modular.network, openzeppelin.org * THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.","methods":{"dividedBy(uint256,uint256)":{"details":"Divides two numbers but checks for 0 in the divisor first. Does not throw.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if b is 0res The quotient of a and b, or 0 if b is 0"},"minus(uint256,uint256)":{"details":"Subtracts two numbers and checks for underflow before returning. Does not throw but rather logs an Err event if there is underflow.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if there is underflowres The difference between a and b, or 0 if there is underflow"},"plus(uint256,uint256)":{"details":"Adds two numbers and checks for overflow before returning. Does not throw.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if there is overflowres The sum of a and b, or 0 if there is overflow"},"times(uint256,uint256)":{"details":"Multiplies two numbers and checks for overflow before returning. Does not throw.","params":{"a":"First number","b":"Second number"},"return":"err False normally, or true if there is overflowres The product of a and b, or 0 if there is overflow"}},"title":"Basic Math Library"},"evm":{"bytecode":{"linkReferences":{},"object":"610198610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029"},"deployedBytecode":{"linkReferences":{},"object":"73000000000000000000000000000000000000000030146080604052600436106100785763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d3b9edf811461007d57806366098d4f146100a6578063e39bbf68146100b4578063f4f3bdc1146100c2575b600080fd5b61008b6004356024356100d0565b60408051921515835260208301919091528051918290030190f35b61008b6004356024356100f9565b61008b600435602435610116565b61008b60043560243561014c565b6000828202821583820485141780156100e8576100f1565b60019250600091505b509250929050565b60008282018281038414838211828514171680156100e8576100f1565b600080808315801561012f576001935060009250610143565b604051858704602090910181905292508291505b50509250929050565b60008183038083018414848210828614171660011480156100e8576100f15600a165627a7a7230582026a004c4f4070b1253b6d155197407659f677073157d481638db922f9adbc6ec0029"}},"userdoc":{"methods":{}}}`
	s := `{ "optimizer": { "enabled": true, "runs": 200 }, "outputSelection": { "*": { "*": ["abi", "evm.bytecode", "evm.deployedBytecode", "devdoc", "userdoc"] } } }`
	if err := ct.Build("solc", s, js); err != nil {
		log.Fatal(err)
	}
}

func TestValidate(t *testing.T) {
	var want error
	var got error
	ct := ContractType{}

	got = ct.Validate("PauloPiresToken")
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}

	ct.ContractName = " token-contract"
	got = ct.Validate("BenVeraToken")
	want = errors.New("contract_type[BenVeraToken]:contract_name error 'Name ' token-contract' " +
		"does not conform to the standard. Please check for extra whitespace and see " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier for the requirement.'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	ct.ContractName = ""
	dub := &bc.UnlinkedBytecode{}
	dub.Bytecode = "0x3g"
	ct.DeploymentBytecode = dub
	got = ct.Validate("NickGToken")
	want = errors.New("deployment_bytecode for contract_type[NickGToken] returned the " +
		"following error: unlinked_bytecode:bytecode error 'Does not conform to a hexadecimal string'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
	dub.Bytecode = ""

	rub := &bc.UnlinkedBytecode{}
	rub.Bytecode = "0x324"
	ct.RuntimeBytecode = rub
	got = ct.Validate("DaviToken")
	want = errors.New("runtime_bytecode for contract_type[DaviToken] returned the " +
		"following error: unlinked_bytecode:bytecode error 'The string does not contain 2 characters per byte, length is showing '5''")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	dub.Bytecode = "0x431101"
	rub.Bytecode = "0x3071d1"

	got = ct.Validate("NoMoreTokens")
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}

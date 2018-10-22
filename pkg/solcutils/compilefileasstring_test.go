package solcutils

import (
	"fmt"
	"log"
	"testing"
)

func TestCompileFileAsString(t *testing.T) {
	s := `pragma solidity ^0.4.21;

	/**
	 * @title TokenLib
	 * @author Modular Inc, https://modular.network
	 *
	 * version 1.3.3
	 * Copyright (c) 2017 Modular, Inc
	 * The MIT License (MIT)
	 * https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE
	 *
	 * The Token Library provides functionality to create a variety of ERC20 tokens.
	 * See https://github.com/Modular-Network/ethereum-contracts for an example of how to
	 * create a basic ERC20 token.
	 *
	 * Modular works on open source projects in the Ethereum community with the
	 * purpose of testing, documenting, and deploying reusable code onto the
	 * blockchain to improve security and usability of smart contracts. Modular
	 * also strives to educate non-profits, schools, and other community members
	 * about the application of blockchain technology.
	 * For further information: modular.network
	 *
	 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
	 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
	 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
	 * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
	 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
	 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
	 */

	import "ethereum-libraries-basic-math/contracts/BasicMathLib.sol";

	library TokenLib {
	  using BasicMathLib for uint256;

	  struct TokenStorage {
	    bool initialized;
	    mapping (address => uint256) balances;
	    mapping (address => mapping (address => uint256)) allowed;

	    string name;
	    string symbol;
	    uint256 totalSupply;
	    uint256 initialSupply;
	    address owner;
	    uint8 decimals;
	    bool stillMinting;
	  }

	  event Transfer(address indexed from, address indexed to, uint256 value);
	  event Approval(address indexed owner, address indexed spender, uint256 value);
	  event OwnerChange(address from, address to);
	  event Burn(address indexed burner, uint256 value);
	  event MintingClosed(bool mintingClosed);

	  /// @dev Called by the Standard Token upon creation.
	  /// @param self Stored token from token contract
	  /// @param _name Name of the new token
	  /// @param _symbol Symbol of the new token
	  /// @param _decimals Decimal places for the token represented
	  /// @param _initial_supply The initial token supply
	  /// @param _allowMinting True if additional tokens can be created, false otherwise
	  function init(TokenStorage storage self,
	                address _owner,
	                string _name,
	                string _symbol,
	                uint8 _decimals,
	                uint256 _initial_supply,
	                bool _allowMinting)
	                public
	  {
	    require(!self.initialized);
	    self.initialized = true;
	    self.name = _name;
	    self.symbol = _symbol;
	    self.totalSupply = _initial_supply;
	    self.initialSupply = _initial_supply;
	    self.decimals = _decimals;
	    self.owner = _owner;
	    self.stillMinting = _allowMinting;
	    self.balances[_owner] = _initial_supply;
	  }

	  /// @dev Transfer tokens from caller's account to another account.
	  /// @param self Stored token from token contract
	  /// @param _to Address to send tokens
	  /// @param _value Number of tokens to send
	  /// @return True if completed
	  function transfer(TokenStorage storage self, address _to, uint256 _value) public returns (bool) {
	    require(_to != address(0));
	    bool err;
	    uint256 balance;

	    (err,balance) = self.balances[msg.sender].minus(_value);
	    require(!err);
	    self.balances[msg.sender] = balance;
	    //It's not possible to overflow token supply
	    self.balances[_to] = self.balances[_to] + _value;
	    emit Transfer(msg.sender, _to, _value);
	    return true;
	  }

	  /// @dev Authorized caller transfers tokens from one account to another
	  /// @param self Stored token from token contract
	  /// @param _from Address to send tokens from
	  /// @param _to Address to send tokens to
	  /// @param _value Number of tokens to send
	  /// @return True if completed
	  function transferFrom(TokenStorage storage self,
	                        address _from,
	                        address _to,
	                        uint256 _value)
	                        public
	                        returns (bool)
	  {
	    uint256 _allowance = self.allowed[_from][msg.sender];
	    bool err;
	    uint256 balanceOwner;
	    uint256 balanceSpender;

	    (err,balanceOwner) = self.balances[_from].minus(_value);
	    require(!err);

	    (err,balanceSpender) = _allowance.minus(_value);
	    require(!err);

	    self.balances[_from] = balanceOwner;
	    self.allowed[_from][msg.sender] = balanceSpender;
	    self.balances[_to] = self.balances[_to] + _value;

	    emit Transfer(_from, _to, _value);
	    return true;
	  }

	  /// @dev Retrieve token balance for an account
	  /// @param self Stored token from token contract
	  /// @param _owner Address to retrieve balance of
	  /// @return balance The number of tokens in the subject account
	  function balanceOf(TokenStorage storage self, address _owner) public view returns (uint256 balance) {
	    return self.balances[_owner];
	  }

	  /// @dev Authorize an account to send tokens on caller's behalf
	  /// @param self Stored token from token contract
	  /// @param _spender Address to authorize
	  /// @param _value Number of tokens authorized account may send
	  /// @return True if completed
	  function approve(TokenStorage storage self, address _spender, uint256 _value) public returns (bool) {
	    // must set to zero before changing approval amount in accordance with spec
	    require((_value == 0) || (self.allowed[msg.sender][_spender] == 0));

	    self.allowed[msg.sender][_spender] = _value;
	    emit Approval(msg.sender, _spender, _value);
	    return true;
	  }

	  /// @dev Remaining tokens third party spender has to send
	  /// @param self Stored token from token contract
	  /// @param _owner Address of token holder
	  /// @param _spender Address of authorized spender
	  /// @return remaining Number of tokens spender has left in owner's account
	  function allowance(TokenStorage storage self, address _owner, address _spender)
	                     public
	                     view
	                     returns (uint256 remaining) {
	    return self.allowed[_owner][_spender];
	  }

	  /// @dev Authorize third party transfer by increasing/decreasing allowed rather than setting it
	  /// @param self Stored token from token contract
	  /// @param _spender Address to authorize
	  /// @param _valueChange Increase or decrease in number of tokens authorized account may send
	  /// @param _increase True if increasing allowance, false if decreasing allowance
	  /// @return True if completed
	  function approveChange (TokenStorage storage self, address _spender, uint256 _valueChange, bool _increase)
	                          public returns (bool)
	  {
	    uint256 _newAllowed;
	    bool err;

	    if(_increase) {
	      (err, _newAllowed) = self.allowed[msg.sender][_spender].plus(_valueChange);
	      require(!err);

	      self.allowed[msg.sender][_spender] = _newAllowed;
	    } else {
	      if (_valueChange > self.allowed[msg.sender][_spender]) {
	        self.allowed[msg.sender][_spender] = 0;
	      } else {
	        _newAllowed = self.allowed[msg.sender][_spender] - _valueChange;
	        self.allowed[msg.sender][_spender] = _newAllowed;
	      }
	    }

	    emit Approval(msg.sender, _spender, _newAllowed);
	    return true;
	  }

	  /// @dev Change owning address of the token contract, specifically for minting
	  /// @param self Stored token from token contract
	  /// @param _newOwner Address for the new owner
	  /// @return True if completed
	  function changeOwner(TokenStorage storage self, address _newOwner) public returns (bool) {
	    require((self.owner == msg.sender) && (_newOwner > 0));

	    self.owner = _newOwner;
	    emit OwnerChange(msg.sender, _newOwner);
	    return true;
	  }

	  /// @dev Mints additional tokens, new tokens go to owner
	  /// @param self Stored token from token contract
	  /// @param _amount Number of tokens to mint
	  /// @return True if completed
	  function mintToken(TokenStorage storage self, uint256 _amount) public returns (bool) {
	    require((self.owner == msg.sender) && self.stillMinting);
	    uint256 _newAmount;
	    bool err;

	    (err, _newAmount) = self.totalSupply.plus(_amount);
	    require(!err);

	    self.totalSupply =  _newAmount;
	    self.balances[self.owner] = self.balances[self.owner] + _amount;
	    emit Transfer(0x0, self.owner, _amount);
	    return true;
	  }

	  /// @dev Permanent stops minting
	  /// @param self Stored token from token contract
	  /// @return True if completed
	  function closeMint(TokenStorage storage self) public returns (bool) {
	    require(self.owner == msg.sender);

	    self.stillMinting = false;
	    emit MintingClosed(true);
	    return true;
	  }

	  /// @dev Permanently burn tokens
	  /// @param self Stored token from token contract
	  /// @param _amount Amount of tokens to burn
	  /// @return True if completed
	  function burnToken(TokenStorage storage self, uint256 _amount) public returns (bool) {
	      uint256 _newBalance;
	      bool err;

	      (err, _newBalance) = self.balances[msg.sender].minus(_amount);
	      require(!err);

	      self.balances[msg.sender] = _newBalance;
	      self.totalSupply = self.totalSupply - _amount;
	      emit Burn(msg.sender, _amount);
	      emit Transfer(msg.sender, 0x0, _amount);
	      return true;
	  }
	}
`
	d := []string{"ethereum-libraries-basic-math"}
	got, _, err := CompileFileAsString("solc", "", "BasicMathLib", true, "", s, d, false, 0)
	want := `{"language":"Solidity","sources":{"BasicMathLib":{"content":"pragma solidity ^0.4.21;\n\n\t/**\n\t * @title TokenLib\n\t * @author Modular Inc, https://modular.network\n\t *\n\t * version 1.3.3\n\t * Copyright (c) 2017 Modular, Inc\n\t * The MIT License (MIT)\n\t * https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE\n\t *\n\t * The Token Library provides functionality to create a variety of ERC20 tokens.\n\t * See https://github.com/Modular-Network/ethereum-contracts for an example of how to\n\t * create a basic ERC20 token.\n\t *\n\t * Modular works on open source projects in the Ethereum community with the\n\t * purpose of testing, documenting, and deploying reusable code onto the\n\t * blockchain to improve security and usability of smart contracts. Modular\n\t * also strives to educate non-profits, schools, and other community members\n\t * about the application of blockchain technology.\n\t * For further information: modular.network\n\t *\n\t * THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS\n\t * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF\n\t * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.\n\t * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY\n\t * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,\n\t * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE\n\t * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.\n\t */\n\n\timport \"ethereum-libraries-basic-math/contracts/BasicMathLib.sol\";\n\n\tlibrary TokenLib {\n\t  using BasicMathLib for uint256;\n\n\t  struct TokenStorage {\n\t    bool initialized;\n\t    mapping (address =\u003e uint256) balances;\n\t    mapping (address =\u003e mapping (address =\u003e uint256)) allowed;\n\n\t    string name;\n\t    string symbol;\n\t    uint256 totalSupply;\n\t    uint256 initialSupply;\n\t    address owner;\n\t    uint8 decimals;\n\t    bool stillMinting;\n\t  }\n\n\t  event Transfer(address indexed from, address indexed to, uint256 value);\n\t  event Approval(address indexed owner, address indexed spender, uint256 value);\n\t  event OwnerChange(address from, address to);\n\t  event Burn(address indexed burner, uint256 value);\n\t  event MintingClosed(bool mintingClosed);\n\n\t  /// @dev Called by the Standard Token upon creation.\n\t  /// @param self Stored token from token contract\n\t  /// @param _name Name of the new token\n\t  /// @param _symbol Symbol of the new token\n\t  /// @param _decimals Decimal places for the token represented\n\t  /// @param _initial_supply The initial token supply\n\t  /// @param _allowMinting True if additional tokens can be created, false otherwise\n\t  function init(TokenStorage storage self,\n\t                address _owner,\n\t                string _name,\n\t                string _symbol,\n\t                uint8 _decimals,\n\t                uint256 _initial_supply,\n\t                bool _allowMinting)\n\t                public\n\t  {\n\t    require(!self.initialized);\n\t    self.initialized = true;\n\t    self.name = _name;\n\t    self.symbol = _symbol;\n\t    self.totalSupply = _initial_supply;\n\t    self.initialSupply = _initial_supply;\n\t    self.decimals = _decimals;\n\t    self.owner = _owner;\n\t    self.stillMinting = _allowMinting;\n\t    self.balances[_owner] = _initial_supply;\n\t  }\n\n\t  /// @dev Transfer tokens from caller's account to another account.\n\t  /// @param self Stored token from token contract\n\t  /// @param _to Address to send tokens\n\t  /// @param _value Number of tokens to send\n\t  /// @return True if completed\n\t  function transfer(TokenStorage storage self, address _to, uint256 _value) public returns (bool) {\n\t    require(_to != address(0));\n\t    bool err;\n\t    uint256 balance;\n\n\t    (err,balance) = self.balances[msg.sender].minus(_value);\n\t    require(!err);\n\t    self.balances[msg.sender] = balance;\n\t    //It's not possible to overflow token supply\n\t    self.balances[_to] = self.balances[_to] + _value;\n\t    emit Transfer(msg.sender, _to, _value);\n\t    return true;\n\t  }\n\n\t  /// @dev Authorized caller transfers tokens from one account to another\n\t  /// @param self Stored token from token contract\n\t  /// @param _from Address to send tokens from\n\t  /// @param _to Address to send tokens to\n\t  /// @param _value Number of tokens to send\n\t  /// @return True if completed\n\t  function transferFrom(TokenStorage storage self,\n\t                        address _from,\n\t                        address _to,\n\t                        uint256 _value)\n\t                        public\n\t                        returns (bool)\n\t  {\n\t    uint256 _allowance = self.allowed[_from][msg.sender];\n\t    bool err;\n\t    uint256 balanceOwner;\n\t    uint256 balanceSpender;\n\n\t    (err,balanceOwner) = self.balances[_from].minus(_value);\n\t    require(!err);\n\n\t    (err,balanceSpender) = _allowance.minus(_value);\n\t    require(!err);\n\n\t    self.balances[_from] = balanceOwner;\n\t    self.allowed[_from][msg.sender] = balanceSpender;\n\t    self.balances[_to] = self.balances[_to] + _value;\n\n\t    emit Transfer(_from, _to, _value);\n\t    return true;\n\t  }\n\n\t  /// @dev Retrieve token balance for an account\n\t  /// @param self Stored token from token contract\n\t  /// @param _owner Address to retrieve balance of\n\t  /// @return balance The number of tokens in the subject account\n\t  function balanceOf(TokenStorage storage self, address _owner) public view returns (uint256 balance) {\n\t    return self.balances[_owner];\n\t  }\n\n\t  /// @dev Authorize an account to send tokens on caller's behalf\n\t  /// @param self Stored token from token contract\n\t  /// @param _spender Address to authorize\n\t  /// @param _value Number of tokens authorized account may send\n\t  /// @return True if completed\n\t  function approve(TokenStorage storage self, address _spender, uint256 _value) public returns (bool) {\n\t    // must set to zero before changing approval amount in accordance with spec\n\t    require((_value == 0) || (self.allowed[msg.sender][_spender] == 0));\n\n\t    self.allowed[msg.sender][_spender] = _value;\n\t    emit Approval(msg.sender, _spender, _value);\n\t    return true;\n\t  }\n\n\t  /// @dev Remaining tokens third party spender has to send\n\t  /// @param self Stored token from token contract\n\t  /// @param _owner Address of token holder\n\t  /// @param _spender Address of authorized spender\n\t  /// @return remaining Number of tokens spender has left in owner's account\n\t  function allowance(TokenStorage storage self, address _owner, address _spender)\n\t                     public\n\t                     view\n\t                     returns (uint256 remaining) {\n\t    return self.allowed[_owner][_spender];\n\t  }\n\n\t  /// @dev Authorize third party transfer by increasing/decreasing allowed rather than setting it\n\t  /// @param self Stored token from token contract\n\t  /// @param _spender Address to authorize\n\t  /// @param _valueChange Increase or decrease in number of tokens authorized account may send\n\t  /// @param _increase True if increasing allowance, false if decreasing allowance\n\t  /// @return True if completed\n\t  function approveChange (TokenStorage storage self, address _spender, uint256 _valueChange, bool _increase)\n\t                          public returns (bool)\n\t  {\n\t    uint256 _newAllowed;\n\t    bool err;\n\n\t    if(_increase) {\n\t      (err, _newAllowed) = self.allowed[msg.sender][_spender].plus(_valueChange);\n\t      require(!err);\n\n\t      self.allowed[msg.sender][_spender] = _newAllowed;\n\t    } else {\n\t      if (_valueChange \u003e self.allowed[msg.sender][_spender]) {\n\t        self.allowed[msg.sender][_spender] = 0;\n\t      } else {\n\t        _newAllowed = self.allowed[msg.sender][_spender] - _valueChange;\n\t        self.allowed[msg.sender][_spender] = _newAllowed;\n\t      }\n\t    }\n\n\t    emit Approval(msg.sender, _spender, _newAllowed);\n\t    return true;\n\t  }\n\n\t  /// @dev Change owning address of the token contract, specifically for minting\n\t  /// @param self Stored token from token contract\n\t  /// @param _newOwner Address for the new owner\n\t  /// @return True if completed\n\t  function changeOwner(TokenStorage storage self, address _newOwner) public returns (bool) {\n\t    require((self.owner == msg.sender) \u0026\u0026 (_newOwner \u003e 0));\n\n\t    self.owner = _newOwner;\n\t    emit OwnerChange(msg.sender, _newOwner);\n\t    return true;\n\t  }\n\n\t  /// @dev Mints additional tokens, new tokens go to owner\n\t  /// @param self Stored token from token contract\n\t  /// @param _amount Number of tokens to mint\n\t  /// @return True if completed\n\t  function mintToken(TokenStorage storage self, uint256 _amount) public returns (bool) {\n\t    require((self.owner == msg.sender) \u0026\u0026 self.stillMinting);\n\t    uint256 _newAmount;\n\t    bool err;\n\n\t    (err, _newAmount) = self.totalSupply.plus(_amount);\n\t    require(!err);\n\n\t    self.totalSupply =  _newAmount;\n\t    self.balances[self.owner] = self.balances[self.owner] + _amount;\n\t    emit Transfer(0x0, self.owner, _amount);\n\t    return true;\n\t  }\n\n\t  /// @dev Permanent stops minting\n\t  /// @param self Stored token from token contract\n\t  /// @return True if completed\n\t  function closeMint(TokenStorage storage self) public returns (bool) {\n\t    require(self.owner == msg.sender);\n\n\t    self.stillMinting = false;\n\t    emit MintingClosed(true);\n\t    return true;\n\t  }\n\n\t  /// @dev Permanently burn tokens\n\t  /// @param self Stored token from token contract\n\t  /// @param _amount Amount of tokens to burn\n\t  /// @return True if completed\n\t  function burnToken(TokenStorage storage self, uint256 _amount) public returns (bool) {\n\t      uint256 _newBalance;\n\t      bool err;\n\n\t      (err, _newBalance) = self.balances[msg.sender].minus(_amount);\n\t      require(!err);\n\n\t      self.balances[msg.sender] = _newBalance;\n\t      self.totalSupply = self.totalSupply - _amount;\n\t      emit Burn(msg.sender, _amount);\n\t      emit Transfer(msg.sender, 0x0, _amount);\n\t      return true;\n\t  }\n\t}\n"}},"settings":{"optimizer":{},"outputSelection":{"*":{"*":["evm.bytecode","evm.deployedBytecode"]}},"remappings":["ethereum-libraries-basic-math/=./ethereum-libraries-basic-math/"]}}`
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if got != want {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func ExampleCompileFileAsString() {
	s := `pragma solidity ^0.4.21;

	/**
	 * @title TokenLib
	 * @author Modular Inc, https://modular.network
	 *
	 * version 1.3.3
	 * Copyright (c) 2017 Modular, Inc
	 * The MIT License (MIT)
	 * https://github.com/Modular-Network/ethereum-libraries/blob/master/LICENSE
	 *
	 * The Token Library provides functionality to create a variety of ERC20 tokens.
	 * See https://github.com/Modular-Network/ethereum-contracts for an example of how to
	 * create a basic ERC20 token.
	 *
	 * Modular works on open source projects in the Ethereum community with the
	 * purpose of testing, documenting, and deploying reusable code onto the
	 * blockchain to improve security and usability of smart contracts. Modular
	 * also strives to educate non-profits, schools, and other community members
	 * about the application of blockchain technology.
	 * For further information: modular.network
	 *
	 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
	 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
	 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
	 * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
	 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
	 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
	 */

	import "ethereum-libraries-basic-math/contracts/BasicMathLib.sol";

	library TokenLib {
	  using BasicMathLib for uint256;

	  struct TokenStorage {
	    bool initialized;
	    mapping (address => uint256) balances;
	    mapping (address => mapping (address => uint256)) allowed;

	    string name;
	    string symbol;
	    uint256 totalSupply;
	    uint256 initialSupply;
	    address owner;
	    uint8 decimals;
	    bool stillMinting;
	  }

	  event Transfer(address indexed from, address indexed to, uint256 value);
	  event Approval(address indexed owner, address indexed spender, uint256 value);
	  event OwnerChange(address from, address to);
	  event Burn(address indexed burner, uint256 value);
	  event MintingClosed(bool mintingClosed);

	  /// @dev Called by the Standard Token upon creation.
	  /// @param self Stored token from token contract
	  /// @param _name Name of the new token
	  /// @param _symbol Symbol of the new token
	  /// @param _decimals Decimal places for the token represented
	  /// @param _initial_supply The initial token supply
	  /// @param _allowMinting True if additional tokens can be created, false otherwise
	  function init(TokenStorage storage self,
	                address _owner,
	                string _name,
	                string _symbol,
	                uint8 _decimals,
	                uint256 _initial_supply,
	                bool _allowMinting)
	                public
	  {
	    require(!self.initialized);
	    self.initialized = true;
	    self.name = _name;
	    self.symbol = _symbol;
	    self.totalSupply = _initial_supply;
	    self.initialSupply = _initial_supply;
	    self.decimals = _decimals;
	    self.owner = _owner;
	    self.stillMinting = _allowMinting;
	    self.balances[_owner] = _initial_supply;
	  }

	  /// @dev Transfer tokens from caller's account to another account.
	  /// @param self Stored token from token contract
	  /// @param _to Address to send tokens
	  /// @param _value Number of tokens to send
	  /// @return True if completed
	  function transfer(TokenStorage storage self, address _to, uint256 _value) public returns (bool) {
	    require(_to != address(0));
	    bool err;
	    uint256 balance;

	    (err,balance) = self.balances[msg.sender].minus(_value);
	    require(!err);
	    self.balances[msg.sender] = balance;
	    //It's not possible to overflow token supply
	    self.balances[_to] = self.balances[_to] + _value;
	    emit Transfer(msg.sender, _to, _value);
	    return true;
	  }

	  /// @dev Authorized caller transfers tokens from one account to another
	  /// @param self Stored token from token contract
	  /// @param _from Address to send tokens from
	  /// @param _to Address to send tokens to
	  /// @param _value Number of tokens to send
	  /// @return True if completed
	  function transferFrom(TokenStorage storage self,
	                        address _from,
	                        address _to,
	                        uint256 _value)
	                        public
	                        returns (bool)
	  {
	    uint256 _allowance = self.allowed[_from][msg.sender];
	    bool err;
	    uint256 balanceOwner;
	    uint256 balanceSpender;

	    (err,balanceOwner) = self.balances[_from].minus(_value);
	    require(!err);

	    (err,balanceSpender) = _allowance.minus(_value);
	    require(!err);

	    self.balances[_from] = balanceOwner;
	    self.allowed[_from][msg.sender] = balanceSpender;
	    self.balances[_to] = self.balances[_to] + _value;

	    emit Transfer(_from, _to, _value);
	    return true;
	  }

	  /// @dev Retrieve token balance for an account
	  /// @param self Stored token from token contract
	  /// @param _owner Address to retrieve balance of
	  /// @return balance The number of tokens in the subject account
	  function balanceOf(TokenStorage storage self, address _owner) public view returns (uint256 balance) {
	    return self.balances[_owner];
	  }

	  /// @dev Authorize an account to send tokens on caller's behalf
	  /// @param self Stored token from token contract
	  /// @param _spender Address to authorize
	  /// @param _value Number of tokens authorized account may send
	  /// @return True if completed
	  function approve(TokenStorage storage self, address _spender, uint256 _value) public returns (bool) {
	    // must set to zero before changing approval amount in accordance with spec
	    require((_value == 0) || (self.allowed[msg.sender][_spender] == 0));

	    self.allowed[msg.sender][_spender] = _value;
	    emit Approval(msg.sender, _spender, _value);
	    return true;
	  }

	  /// @dev Remaining tokens third party spender has to send
	  /// @param self Stored token from token contract
	  /// @param _owner Address of token holder
	  /// @param _spender Address of authorized spender
	  /// @return remaining Number of tokens spender has left in owner's account
	  function allowance(TokenStorage storage self, address _owner, address _spender)
	                     public
	                     view
	                     returns (uint256 remaining) {
	    return self.allowed[_owner][_spender];
	  }

	  /// @dev Authorize third party transfer by increasing/decreasing allowed rather than setting it
	  /// @param self Stored token from token contract
	  /// @param _spender Address to authorize
	  /// @param _valueChange Increase or decrease in number of tokens authorized account may send
	  /// @param _increase True if increasing allowance, false if decreasing allowance
	  /// @return True if completed
	  function approveChange (TokenStorage storage self, address _spender, uint256 _valueChange, bool _increase)
	                          public returns (bool)
	  {
	    uint256 _newAllowed;
	    bool err;

	    if(_increase) {
	      (err, _newAllowed) = self.allowed[msg.sender][_spender].plus(_valueChange);
	      require(!err);

	      self.allowed[msg.sender][_spender] = _newAllowed;
	    } else {
	      if (_valueChange > self.allowed[msg.sender][_spender]) {
	        self.allowed[msg.sender][_spender] = 0;
	      } else {
	        _newAllowed = self.allowed[msg.sender][_spender] - _valueChange;
	        self.allowed[msg.sender][_spender] = _newAllowed;
	      }
	    }

	    emit Approval(msg.sender, _spender, _newAllowed);
	    return true;
	  }

	  /// @dev Change owning address of the token contract, specifically for minting
	  /// @param self Stored token from token contract
	  /// @param _newOwner Address for the new owner
	  /// @return True if completed
	  function changeOwner(TokenStorage storage self, address _newOwner) public returns (bool) {
	    require((self.owner == msg.sender) && (_newOwner > 0));

	    self.owner = _newOwner;
	    emit OwnerChange(msg.sender, _newOwner);
	    return true;
	  }

	  /// @dev Mints additional tokens, new tokens go to owner
	  /// @param self Stored token from token contract
	  /// @param _amount Number of tokens to mint
	  /// @return True if completed
	  function mintToken(TokenStorage storage self, uint256 _amount) public returns (bool) {
	    require((self.owner == msg.sender) && self.stillMinting);
	    uint256 _newAmount;
	    bool err;

	    (err, _newAmount) = self.totalSupply.plus(_amount);
	    require(!err);

	    self.totalSupply =  _newAmount;
	    self.balances[self.owner] = self.balances[self.owner] + _amount;
	    emit Transfer(0x0, self.owner, _amount);
	    return true;
	  }

	  /// @dev Permanent stops minting
	  /// @param self Stored token from token contract
	  /// @return True if completed
	  function closeMint(TokenStorage storage self) public returns (bool) {
	    require(self.owner == msg.sender);

	    self.stillMinting = false;
	    emit MintingClosed(true);
	    return true;
	  }

	  /// @dev Permanently burn tokens
	  /// @param self Stored token from token contract
	  /// @param _amount Amount of tokens to burn
	  /// @return True if completed
	  function burnToken(TokenStorage storage self, uint256 _amount) public returns (bool) {
	      uint256 _newBalance;
	      bool err;

	      (err, _newBalance) = self.balances[msg.sender].minus(_amount);
	      require(!err);

	      self.balances[msg.sender] = _newBalance;
	      self.totalSupply = self.totalSupply - _amount;
	      emit Burn(msg.sender, _amount);
	      emit Transfer(msg.sender, 0x0, _amount);
	      return true;
	  }
	}
`
	d := []string{"ethereum-libraries-basic-math"}
	i, o, err := CompileFileAsString("solc", "", "BasicMathLib", true, "", s, d, false, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
	fmt.Println(o)
}

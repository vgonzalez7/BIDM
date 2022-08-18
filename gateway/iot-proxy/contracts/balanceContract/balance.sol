pragma solidity >=0.4.0;
import "./ERC20.sol";
//SPDX-License-Identifier: UNLICENSED"

contract dataLedgerContract {
    function topicExists(string memory topic) public view returns (bool){}
    function getIoTAddress(bytes32 hash) public view returns (address){}
    function getIoTAddress(string memory topic) public view returns (address){}
}

contract balanceContract is ERC20Basic
{
    dataLedgerContract dataContractDef;

    struct retainedMoneyStruct {
        address measurementOwner;
        uint256 tokens;
    }

    // Hash => Price
    mapping(bytes32 => uint256) private prices;
    
    //mapping(string => uint256) public pricesSubs;

    // Hash => (address => Held money)
    mapping(bytes32 => mapping(address => retainedMoneyStruct)) private retentions;
    
    // Cada topic tiene unos suscriptores
    mapping(string => address[]) private subscribers;

    event PriceSet(bytes32 indexed _hash, uint256 _price);
    event RequestPurchase(bytes32 indexed _hash, address indexed _from, address indexed _to);
    event CompletePurchase(bytes32 indexed _hash, address indexed _from, address indexed _to, bytes32 _txHash);
    event PurchaseRevoked(bytes32 indexed _hash, address indexed _from, address indexed _to);
    event AdquireTokens(address indexed _to, uint256 tokens);
    event Subscribe(string indexed _topic, address indexed _from, address indexed _to);


    function setPriceToMeasurement(bytes32 hash, uint256 price) public  {
        // Check that the account that is setting the price is the one who inserted
        // the measurement
        address iotAddr = dataContractDef.getIoTAddress(hash);
        require(msg.sender == iotAddr, "Only the account who inserted the data can set its price");
        prices[hash] = price;
        emit PriceSet(hash, price);
    }

    function getPriceMeasurement(bytes32 hash) public view returns(uint256) {
        return prices[hash];
    }
    
    // function setPriceToSubscription(string memory topic, uint256 price) public {}
    
    function getPriceSubscription(string memory topic) public view returns(uint256) {
        // TEMPORAL
        // return pricesSubs[topic];
        if(dataContractDef.topicExists(topic)) { return 10; }
        else { return 0; }
    }

    function purchaseMeasurement(bytes32 hash) public {
        // get the price of the measurement
        uint256 price = getPriceMeasurement(hash);
        
        // Get the address of the owner
        address owner = dataContractDef.getIoTAddress(hash);
        
        // Check the price is not 0
        require(price > 0, "The data does not exist");

        // Check the user has enough tokens
        require(balanceOf(msg.sender) >= price, "You do not have enough tokens to complete the purchase");

        // retain the money of the purchase from the customer until is delivered
        retentions[hash][msg.sender].measurementOwner = owner;
        retentions[hash][msg.sender].tokens = price;

        transfer(dummyAccount, price);

        emit RequestPurchase(hash, msg.sender, owner);
    }


    // This function is used by the admin to deliver the purchase
    function completePurchase(bytes32 hash, address buyer, bytes32 txHash) public {
        // Only the admin user can complete the purchase
        require(msg.sender == admin, "You do not have enough privileges to do this action");

        // Check whether there is retained money associated to
        // the purchase of "hash" by "buyer"
        require(retentions[hash][buyer].tokens > 0, "There is not transaction associate to this hash by this buyer");

        // Get the address of the owner of the measurement
        address addrTo = retentions[hash][buyer].measurementOwner;

        // Send the withheld moeney to the owner of the measurement
        transferFrom(dummyAccount, addrTo, retentions[hash][buyer].tokens);

        // Clean the struct that held the withheld money
        delete retentions[hash][buyer];

        emit CompletePurchase(hash, buyer, addrTo, txHash);
    }

    function purchaseSubscription(string memory topic) public {
        // get the price of the sub
        uint256 price = getPriceSubscription(topic);
        
        // Check the price is not 0
        require(price > 0, "The topic does not exist");

        // Check the user has enough tokens
        require(balanceOf(msg.sender) >= price, "You do not have enough tokens to complete the purchase");
        
        address owner = dataContractDef.getIoTAddress(topic);
        
        transfer(owner, price);

        subscribers[topic].push(msg.sender);
        emit Subscribe(topic, msg.sender, owner);
    }
    
    function isSubscribed(address client, string memory topic) public view returns (bool)
    {
        address[] memory subs = subscribers[topic];
        uint len = subs.length;
        bool flag = false;
        for(uint i = 0; i < len; i++) {
            if(subs[i] == client) {
                flag = true;
                break;
            }
        }
        
        return flag;
    }

    // Used by the admin to revoke a transaction
    function revokeTransaction(bytes32 hash, address buyer) public {
        require(msg.sender == admin, "You do not have enough privileges to do this action");
        uint256 retainedTokens = retentions[hash][buyer].tokens;
        address iot = retentions[hash][buyer].measurementOwner;
        transferFrom(dummyAccount, buyer, retainedTokens);

        delete retentions[hash][buyer];
        emit PurchaseRevoked(hash, buyer, iot);
    }

    // Sends tokens to clients. Only the administrator can use this function
    function sendTokenToClient(address to, uint256 total) public
    {
        require(msg.sender == admin, "You do not have enough privileges");
        require(total <= balanceOf(admin), "There are not enough tokens");
        transfer(to, total);
        emit AdquireTokens(to, total);
    }


    // Set data address in this contract to read the variable ledger
    function setAddress(address _address) public {
        require(msg.sender == admin, "You do not have privileges to do this action");
        dataContractDef = dataLedgerContract(_address);
    }


}
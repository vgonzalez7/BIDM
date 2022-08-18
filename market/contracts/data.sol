pragma solidity >=0.4.0;
pragma experimental ABIEncoderV2;
//SPDX-License-Identifier: UNLICENSED"


contract accessControlContract
{
    mapping(address => bool) public allowedAccounts;
}


contract dataLedgerContract 
{
    struct dataStruct 
    {
        string uri;
        string[] topic;
        address addr;
    }
    
    accessControlContract accessContract;
    
    event evtStoreInfo(bytes32 indexed _hash, string _uri, string[] _topics);
    event deleteInfo(bytes32 indexed _hash);
    event newTopic(string indexed _topic);
    address admin = 0x7fFF1F93d22E9b2AA4C6a536531950ED386bd95e;
    
    mapping(bytes32  => dataStruct) public ledger;
    
    string[] topics;
    mapping(string => address) private topicOwners;
    mapping(string => bytes32) private topicKeys;
    
    // Stores information in the blockchain
    function storeInfo(bytes32  hash, string memory uri, string[] memory topic) public
    {
        require(checkAccess(msg.sender) == true, "The ID that you are using is not registered");
        dataStruct memory dataToStore;
    
        uint len = topic.length;
        for(uint i = 0; i < len; i++) {
            require(topicExists(topic[i]) == true, "The topic does not exist");
        }
        
        dataToStore.uri = uri;
        dataToStore.topic = topic;
        dataToStore.addr = msg.sender;
        
        ledger[hash] = dataToStore;
        
        // Emit an event once the data has been stored in the blockchain
        emit evtStoreInfo(hash, uri, topic);
    }
    
    // Deletes a measurement from the blockchain
    function deleteMeasurement(bytes32 hash) public 
    {
        // Only the admin user can delete a measurement from the blockchain
        require(msg.sender == admin
        , "You do not have enough privileges to do this action");
        
        // Delete the measurement associated to the hash indicated by the admin
        delete ledger[hash];
        
        // Emit an event that indicates the time when the element was remove
        emit deleteInfo(hash);
    }
    
    // Crear nuevo topic, solo productores
    function addTopic(string memory topic, bytes32 key) public
    {
        require(checkAccess(msg.sender) == true, "The ID that you are using is not registered");
        topics.push(topic);
        topicOwners[topic] = msg.sender;
        topicKeys[topic] = key;
        emit newTopic(topic);
    }
    
    //Ver topics
    function getTopics() public view returns (string[] memory)
    {
        return topics;
    }
    
    function getTopicKey(string memory topic) public view returns (bytes32)
    {
        bool requirement = (msg.sender == admin);
        if (!requirement) { requirement = (msg.sender == getIoTAddress(topic)); }
        
        require(requirement, "You do not have privileges to do this action");   
        return topicKeys[topic];
    }
    
    function topicExists(string memory topic) public view returns (bool)
    {
        uint len = topics.length;
        bool flag = false;
        for(uint i = 0; i < len; i++) {
            if(keccak256(abi.encodePacked(topics[i])) == keccak256(abi.encodePacked(topic))) {
                flag = true;
                break;
            }
        }
        return flag;
    }
    
    //Ver topics de una medida
    function getMeasurementTopics(bytes32 hash) public view returns (string[] memory)
    {
        return ledger[hash].topic;
    }
    
    /*** Access the mapping stored in the contract accessControl ***/
    
    // Set the address where the access control contract is stored 
    function setAddress(address _address) public
    {
        require(msg.sender == admin, "You do not have privileges to do this action");
        accessContract = accessControlContract(_address);
    }
    
    // Get the value of the mapping
    function checkAccess(address producer) public view returns (bool)
    {
        if(msg.sender == admin) { return true; }
        return accessContract.allowedAccounts(producer);
    }
    
    
    // Get the IoT address from a stored dataStruct
    function getIoTAddress(bytes32 hash) public view returns (address)
    {
        return ledger[hash].addr;
    }
    
    // Get the IoT address from a topic
    function getIoTAddress(string memory topic) public view returns (address)
    {
        return topicOwners[topic];
    }
    
}

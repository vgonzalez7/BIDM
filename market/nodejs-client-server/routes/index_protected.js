const common = require("../libs/common.js");
const ecies = require('eciesjs');
const fetch = require('node-fetch');

/**
 * Destroys session cookie and redirects user to /
 */
function logout() {
  return function(req, res) {
    if (req.session) {
      req.session.destroy((err) => {
        if (err) {
          res.status(401);
          res.end();
          return ;
        }
        res.status(200);
        res.end();
      });
    }
  }
}

/**
 * Renders the user's wallet webpage.
 * From this webpage users can check their purchases
 * and purchase new measurements
 * @param {Client configuration} ethClient 
 */
function renderWallet(ethClient) {
  return async function (req, res) {
    // Get the client's address form the session cookie
    let address = req.session.userID;

    // Get the top navigation bar parameters
    let navParams = await common.getNavParams(req, ethClient);

    // Filter the purchases made by the user
    let csAddress = ethClient.web3.utils.toChecksumAddress(address);
    let filter = {
      _from: csAddress
    };

    const fs = require('fs');
    var tstamp1 = Date.now();

    let purchases = [];

    // Read the events, associated to the insertion of new measurements, 
    let lastBlock = await ethClient.web3.eth.getBlockNumber();
    let n = 0;
    let times = 1;
    let fromB = 1;
    while ((n < 2000) && (fromB != 0))  {
      fromB = lastBlock - 500*times;
      let toB = lastBlock - 500*(times-1);
      if (fromB < 0) {fromB = 0}
      let purchasesArray = await ethClient.balanceSC.getPastEvents('CompletePurchase', {filter, fromBlock: fromB, toBlock: toB});
      times++;

      let i = purchasesArray.length - 1;
      for (; i >= 0; i--) {
        // Extract the fields of interest from the bundle
        let hash = purchasesArray[i].returnValues._hash;
        let blockNumber = purchasesArray[i].blockNumber;
        let txHash = purchasesArray[i].returnValues._txHash;
        let price = await ethClient.balanceSC.methods.getPriceMeasurement(hash).call();

        // Append the measurement to the response array
        purchases.push({
          blockNumber: blockNumber,
          hash: hash,
          txHash: txHash,
          price: price
        });
        n++;
        if (n >= 2000) { break; }
      }
      purchasesArray = [];
    }

    var tstamp2 = Date.now();
    var total = tstamp2 - tstamp1;
    fs.appendFile("/home/administrator/actual_blockchain/timestampsV.log", total + " Compradas\n", function (err) {
        if (err) throw err;
    } );

    // Topics the client is subscribed to
    let topics = await ethClient.dataSC.methods.getTopics().call();
    let topicsSub = [];
    for (i = 0; i < topics.length; i++) {
    	let isSubbed = await ethClient.balanceSC.methods.isSubscribed(csAddress,topics[i]).call();
    	if(isSubbed){
    		topicsSub.push(topics[i]);
    	}
    }
    //console.log(topicsSub);

    // Instant when the client subscribed to those topics
    let subBlock = [];
    let subBlocks = await ethClient.balanceSC.getPastEvents('Subscribe', {filter, fromBlock: 0})
    .catch((err) => {console.log(err);});
    for (i = 0; i < topicsSub.length; i++) {
      let nowTopic = ethClient.web3.utils.soliditySha3(topicsSub[i]);
      let last = 0;
      for (j = 0; j < subBlocks.length; j++) {
        if (subBlocks[j].returnValues._topic == nowTopic) {
          last = subBlocks[j].blockNumber;
        }
      }
      subBlock.push(last);
    }

    let measurementsSubbed = [];

    // Registered measurements since that instant
    let activeTopics = topicsSub.slice();
    let nextSub = Math.max(...subBlock);

    n = 0;
    times = 1;
    fromB = 1;
    while ((n < 2000) && (fromB != 0))  {
      fromB = lastBlock - 500*times;
      let toB = lastBlock - 500*(times-1);
      if (fromB < 0) {fromB = 0}
      let medidasArray = await ethClient.dataSC.getPastEvents('evtStoreInfo', {fromBlock: fromB, toBlock: toB});
      times++;

      i = medidasArray.length - 1;
      for (; i >= 0; i--) {
          let mTopics = medidasArray[i].returnValues._topics;
          let bNumber = medidasArray[i].blockNumber;
          
          if (bNumber < nextSub) {
            index = activeTopics.indexOf(topicsSub[subBlock.indexOf(nextSub)]);
            if (index > -1) { activeTopics.splice(index, 1); }
            subBlock[subBlock.indexOf(nextSub)] = -Infinity;
            nextSub = Math.max(...subBlock);
          }

          let sharedTopic = mTopics.some(r => activeTopics.includes(r));
          if (sharedTopic) {
            let hash = medidasArray[i].returnValues._hash;
            let txHash = medidasArray[i].transactionHash;
            measurementsSubbed.push({
              blockNumber: bNumber,
              hash: hash,
              txHash: txHash,
              topics: mTopics
            });
            n++;
            if (n >= 2000) { break; }
          }
      }
      medidasArray = [];
    }

    var tstamp3 = Date.now();
    total = tstamp3 - tstamp2;
    fs.appendFile("/home/administrator/actual_blockchain/timestampsV.log", total + " Suscritas\n", function (err) {
        if (err) throw err;
    } );

    res.render('private/wallet', {navParams: navParams, purchases: purchases, topics: topics, topicsSub: topicsSub, subs: measurementsSubbed});

  }
}

/**
 * Purchases a measurement
 * @param {Client Config} ehtClient 
 */
function purchaseMeasurement(ethClient) {
  return function(req, res) {
	//TIMESTAMP 1
	const fs = require('fs');
    var tstamp = Date.now();
    fs.appendFile("/home/administrator/actual_blockchain/timestampsC.log", tstamp + " 1 Comprada\n", function (err) {
        if (err) throw err;
	} );

    // Get the body of the message
    let body = "";
    req.on('data', chunk => 
    {
      body += chunk;
    });
    
    req.on('end', async function() {
      let jsonBody = JSON.parse(body);
      let hash = jsonBody.hash;
      
      // Recover the user's address and private key from
      // the cookie session
      let address = ethClient.web3.utils.toChecksumAddress(req.session.userID);
      let privKey = req.session.privKey;

      // Get the public key of the client
      let pubKey = req.session.pubKey;

      // Store the client's public key in the blockchain
      await common.sendTransactionContract(ethClient.web3
        , ethClient.accessSC.methods.addPubKey(pubKey)
        , privKey);
      
      // Check if the client has already bought the measurement
      let filter = {
          _from: ethClient.web3.utils.toChecksumAddress(address),
          _hash: hash
      };
      
      // Get all the events that matches the previous filter
      let events = await ethClient.balanceSC.getPastEvents('RequestPurchase', { filter, fromBlock: 0 });

      if (events.length != 0) {
        res.status(500);
        res.end();
        return;
      }

      // Get the price of the measurement
      let price = await ethClient.balanceSC.methods.getPriceMeasurement(hash).call();

      // Check user's balance
      let balance = await ethClient.balanceSC.methods.balanceOf(address).call();

      if (balance < price ) {
        res.status(421);
        res.end();
        return;
      }

    // Purchase the measurement
	// TIMESTAMP 2
	tstamp = Date.now();
	fs.appendFile("/home/administrator/actual_blockchain/timestampsC.log", tstamp + " 2 Enviada\n", function (err) {
	  if (err) throw err;
	} );
      await common.sendTransactionContract(ethClient.web3
        , ethClient.balanceSC.methods.purchaseMeasurement(hash)
        , privKey)
	// TIMESTAMP 3
	tstamp = Date.now();
	fs.appendFile("/home/administrator/actual_blockchain/timestampsC.log", tstamp + " 3 Registrada\n", function (err) {
	          if (err) throw err;
	        } );
	console.log("Compra realizada en timestamp: " + tstamp);

      res.status(200);
      res.end();
    });
  }
}

/* Gets the value of the measurement purchased by the client
 * @param {Client} ethClient 
 */
function valueMeasurement(ethClient) {
  return async function(req, res) {

    // Get the hash of the measurement
    let hash = req.params.hash;
    let clientAddr = ethClient.web3.utils.toChecksumAddress(req.session.userID);

    // Get client's private key
    let privKey = req.session.privKey;

    // Get the transaction hash that contains the 
    // encrypted URL
    let filter = {
      _hash: hash,
      _from: clientAddr
    };

    let filterS = {
      _hash: hash
    };

    // Topics of the measurement
    let topics = await ethClient.dataSC.methods.getMeasurementTopics(hash).call();
    let isAllowed = false;
    for (i = 0; i < topics.length; i++) {
      let isSubbed = await ethClient.balanceSC.methods.isSubscribed(clientAddr,topics[i]).call();
      if(isSubbed){
        isAllowed = true;
        break;
      }
    }

    if (!isAllowed) {
        let purchasesArray = await ethClient.balanceSC.getPastEvents('CompletePurchase', {filter, fromBlock: 0})
          .catch((err) => {console.log(err);});
        if (purchasesArray.length != 0) { isAllowed = true; }
    }


    if (!isAllowed) {
        res.status(400).send("You don't have the right to claim this measurement");
        res.end();
        return;
    }

    // Fetch measurement from the marketplace
	  // Slice: to remove "0x"
    url = ethClient.config.marketAddr + "/admin";
    let peticion = {
      "action": "medida",
      "hash" : hash.slice(-64),
      "addr" : clientAddr.slice(-40)
    };
    let respuesta;

    await fetch(url, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(peticion)
      })
      .then(res => res.buffer())
      .then(res => {
        respuesta = res; 
      });

      //let events = await ethClient.dataSC.getPastEvents('evtStoreInfo', {filterS, fromBlock: 0})
        //.catch((err) => {console.log(err);});
      //let txHash = events[events.length - 1].transactionHash;

      // Get IoT address
      let iotAddr = (await ethClient.dataSC.methods.ledger(hash).call()).addr;

/*  #!old stuff
    // Get the encrypted url
    let secret = (await ethClient.web3.eth.getTransaction(txHash)).input;

    // Decrypt the URL and the symmetric key using the clients private key
    let plainData = ecies.decrypt(Buffer.from(privKey.substring(2), 'hex'), Buffer.from(secret.substring(2), 'hex'));
    let symmetricKey = plainData.slice(0, 32);
    let cid = String(plainData.slice(32));

    // Query the url to obtain the data
    let encryptedData = await common.fetchIPFSData(ethClient.config.IPFSaddr, cid);

    // Decipher the measurement using the symmetric key
    let plain = common.decryptAES(symmetricKey, encryptedData);
*/

    let measurement = respuesta.slice(0, respuesta.length-65 );
    let signature = respuesta.slice(respuesta.length-64);

    let r = ethClient.web3.utils.bytesToHex(signature.slice(0, 32));
    let s = ethClient.web3.utils.bytesToHex(signature.slice(32));

    // Get the top navigation bar parameters  
    let navParams = await common.getNavParams(req, ethClient);

    res.render('private/measurement', {
      navParams: navParams, 
      hash: hash, 
      measurement: measurement,
      //txHash: txHash,
      clientAddr: clientAddr,
      iotAddr: iotAddr,
      r: r,
      s: s
    });
  }
}

/**
 * Purchases a sub
 * @param {Client Config} ehtClient 
 */
function purchaseSubscription(ethClient) {
  return function(req, res) {
    const fs = require('fs');
    var tstamp = Date.now();
    fs.appendFile("/home/administrator/actual_blockchain/timestampsS.log", tstamp + " Comprada\n", function (err) {
        if (err) throw err;
    } );

    // Get the body of the message
    let body = "";
    req.on('data', chunk => 
    {
      body += chunk;
    });
    
    req.on('end', async function() {
      let jsonBody = JSON.parse(body);
      let topic = jsonBody.topic;
      
      // Recover the user's address and private key from
      // the cookie session
      let address = ethClient.web3.utils.toChecksumAddress(req.session.userID);
      let privKey = req.session.privKey;

      // Get the public key of the client
      let pubKey = req.session.pubKey;

      // Check if topic exists
      let tExists = await ethClient.dataSC.methods.topicExists(topic).call();

      if (!tExists) {
        res.status(500);
        res.end();
        return;
      }

      // Store the client's public key in the blockchain
      await common.sendTransactionContract(ethClient.web3
        , ethClient.accessSC.methods.addPubKey(pubKey)
        , privKey);

      // Check if subscription is already active
      let isSub = await ethClient.balanceSC.methods.isSubscribed(address,topic).call();

      if (isSub) {
        res.status(500);
        res.end();
        return;
      }

      // Get the price of the sub
      let price = await ethClient.balanceSC.methods.getPriceSubscription(topic).call();

      // Check user's balance
      let balance = await ethClient.balanceSC.methods.balanceOf(address).call();

      if (balance < price ) {
        res.status(421);
        res.end();
        return;
      }

      // Purchase the sub
      await common.sendTransactionContract(ethClient.web3
        , ethClient.balanceSC.methods.purchaseSubscription(topic)
        , privKey)
      tstamp = Date.now();
      console.log("Suscripción realizada en timestamp: " + tstamp);
      fs.appendFile("/home/administrator/actual_blockchain/timestampsS.log", tstamp + " Finalizada\n", function (err) {
        if (err) throw err;
      } );

      res.status(200);
      res.end();
    });
  }
}

module.exports = {
  renderWallet,
  purchaseMeasurement,
  purchaseSubscription,
  logout,
  valueMeasurement,
  purchaseSubscription
}

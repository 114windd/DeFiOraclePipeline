// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import "../src/Oracle.sol";
import "../src/ConsumerDemo.sol";

contract DeployScript is Script {
    function run() external {
        // Get the deployer's private key from environment
        uint256 deployerPrivateKey = vm.envUint("ANVIL_PRIVATE_KEY");
        
        // Start broadcasting transactions
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy Oracle contract
        // Use the deployer as the initial updater
        address deployer = vm.addr(deployerPrivateKey);
        Oracle oracle = new Oracle(deployer);
        
        // Deploy ConsumerDemo contract
        ConsumerDemo consumerDemo = new ConsumerDemo(payable(address(oracle)));
        
        // Stop broadcasting
        vm.stopBroadcast();
        
        // Log deployment information
        console.log("Oracle deployed at:", address(oracle));
        console.log("ConsumerDemo deployed at:", address(consumerDemo));
        console.log("Deployer address:", deployer);
    }
}

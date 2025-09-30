// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "../src/Oracle.sol";

contract OracleTest is Test {
    Oracle public oracle;
    address public updater;
    address public owner;
    
    uint256 public constant INITIAL_PRICE = 2000 * 10**8; // $2000 with 8 decimals
    uint256 public constant PRICE_DECIMALS = 8;
    
    event PriceUpdated(uint256 indexed price, uint256 indexed timestamp, uint256 indexed roundId);
    event UpdaterChanged(address indexed oldUpdater, address indexed newUpdater);
    
    function setUp() public {
        owner = address(this);
        updater = makeAddr("updater");
        oracle = new Oracle(updater);
    }
    
    function testInitialState() public {
        assertEq(oracle.owner(), owner);
        assertEq(oracle.updater(), updater);
        assertFalse(oracle.paused());
        
        (uint256 price, uint256 timestamp, uint256 roundId) = oracle.latestPrice();
        assertEq(price, 0);
        assertEq(timestamp, 0);
        assertEq(roundId, 0);
    }
    
    function testUpdatePrice() public {
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE);
        
        (uint256 price, uint256 timestamp, uint256 roundId) = oracle.getLatestPrice();
        assertEq(price, INITIAL_PRICE);
        assertEq(timestamp, block.timestamp);
        assertEq(roundId, 1);
    }
    
    function testUpdatePriceOnlyUpdater() public {
        vm.expectRevert("Oracle: caller is not the updater");
        oracle.updatePrice(INITIAL_PRICE);
    }
    
    function testUpdatePriceWhenPaused() public {
        oracle.pause();
        
        vm.prank(updater);
        vm.expectRevert("Pausable: paused");
        oracle.updatePrice(INITIAL_PRICE);
    }
    
    function testUpdatePriceTooLow() public {
        uint256 lowPrice = 5 * 10**(PRICE_DECIMALS - 1); // $0.50
        
        vm.prank(updater);
        vm.expectRevert("Oracle: price too low");
        oracle.updatePrice(lowPrice);
    }
    
    function testUpdatePriceTooHigh() public {
        uint256 highPrice = 2000000 * 10**PRICE_DECIMALS; // $2M
        
        vm.prank(updater);
        vm.expectRevert("Oracle: price too high");
        oracle.updatePrice(highPrice);
    }
    
    function testSetUpdater() public {
        address newUpdater = makeAddr("newUpdater");
        
        vm.expectEmit(true, true, false, true);
        emit UpdaterChanged(updater, newUpdater);
        oracle.setUpdater(newUpdater);
        
        assertEq(oracle.updater(), newUpdater);
    }
    
    function testSetUpdaterOnlyOwner() public {
        address newUpdater = makeAddr("newUpdater");
        
        vm.prank(updater);
        vm.expectRevert("Ownable: caller is not the owner");
        oracle.setUpdater(newUpdater);
    }
    
    function testSetUpdaterZeroAddress() public {
        vm.expectRevert("Oracle: new updater cannot be zero address");
        oracle.setUpdater(address(0));
    }
    
    function testPauseUnpause() public {
        assertFalse(oracle.paused());
        
        oracle.pause();
        assertTrue(oracle.paused());
        
        oracle.unpause();
        assertFalse(oracle.paused());
    }
    
    function testPauseOnlyOwner() public {
        vm.prank(updater);
        vm.expectRevert("Ownable: caller is not the owner");
        oracle.pause();
    }
    
    function testIsStale() public {
        // Initially stale (no price set, timestamp = 0)
        assertTrue(oracle.isStale());
        
        // Set a price
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE);
        assertFalse(oracle.isStale());
        
        // Fast forward past max age
        vm.warp(block.timestamp + 1 hours + 1);
        assertTrue(oracle.isStale());
    }
    
    function testGetPriceAge() public {
        // Initially no age (no price set)
        assertEq(oracle.getPriceAge(), block.timestamp);
        
        // Set a price
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE);
        assertEq(oracle.getPriceAge(), 0);
        
        // Fast forward 1 hour
        vm.warp(block.timestamp + 1 hours);
        assertEq(oracle.getPriceAge(), 1 hours);
    }
    
    function testGetCurrentRoundId() public {
        assertEq(oracle.getCurrentRoundId(), 0);
        
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE);
        assertEq(oracle.getCurrentRoundId(), 1);
        
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE + 100);
        assertEq(oracle.getCurrentRoundId(), 2);
    }
    
    function testGetContractInfo() public {
        (
            address updaterAddress,
            bool isPaused,
            uint256 maxAge,
            uint256 minPrice,
            uint256 maxPrice
        ) = oracle.getContractInfo();
        
        assertEq(updaterAddress, updater);
        assertFalse(isPaused);
        assertEq(maxAge, 1 hours);
        assertEq(minPrice, 1 * 10**PRICE_DECIMALS);
        assertEq(maxPrice, 1000000 * 10**PRICE_DECIMALS);
    }
    
    function testEmergencyWithdraw() public {
        // Send some ETH to the contract
        vm.deal(address(oracle), 1 ether);
        assertEq(address(oracle).balance, 1 ether);
        
        uint256 ownerBalanceBefore = owner.balance;
        oracle.emergencyWithdraw();
        
        assertEq(address(oracle).balance, 0);
        assertEq(owner.balance, ownerBalanceBefore + 1 ether);
    }
    
    function testReceive() public {
        // Send ETH to the contract
        vm.deal(address(this), 1 ether);
        (bool success,) = address(oracle).call{value: 1 ether}("");
        assertTrue(success);
        assertEq(address(oracle).balance, 1 ether);
    }
    
    function testPriceUpdatedEvent() public {
        vm.expectEmit(true, true, true, true);
        emit PriceUpdated(INITIAL_PRICE, block.timestamp, 1);
        
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE);
    }
    
    function testMultiplePriceUpdates() public {
        uint256[] memory prices = new uint256[](5);
        prices[0] = 2000 * 10**PRICE_DECIMALS;
        prices[1] = 2100 * 10**PRICE_DECIMALS;
        prices[2] = 1950 * 10**PRICE_DECIMALS;
        prices[3] = 2200 * 10**PRICE_DECIMALS;
        prices[4] = 2050 * 10**PRICE_DECIMALS;
        
        for (uint256 i = 0; i < prices.length; i++) {
            vm.prank(updater);
            oracle.updatePrice(prices[i]);
            
            (uint256 price, uint256 timestamp, uint256 roundId) = oracle.getLatestPrice();
            assertEq(price, prices[i]);
            assertEq(timestamp, block.timestamp);
            assertEq(roundId, i + 1);
        }
    }
    
    // Allow the test contract to receive ETH
    receive() external payable {}
}


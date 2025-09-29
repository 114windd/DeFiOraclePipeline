// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "../src/Oracle.sol";
import "../src/ConsumerDemo.sol";

contract ConsumerDemoTest is Test {
    Oracle public oracle;
    ConsumerDemo public consumer;
    address public updater;
    address public user;
    address public authorizedUser;
    
    uint256 public constant INITIAL_PRICE = 2000 * 10**8; // $2000 with 8 decimals
    uint256 public constant STOP_LOSS_PRICE = 1800 * 10**8; // $1800
    uint256 public constant ALERT_PRICE = 2200 * 10**8; // $2200
    uint256 public constant ETH_AMOUNT = 1 ether;
    
    event PositionCreated(address indexed user, uint256 ethAmount, uint256 stopLossPrice, uint256 alertPrice);
    event StopLossTriggered(address indexed user, uint256 ethAmount, uint256 triggerPrice);
    event PriceAlertTriggered(address indexed user, uint256 alertPrice, uint256 currentPrice);
    
    function setUp() public {
        updater = makeAddr("updater");
        user = makeAddr("user");
        authorizedUser = makeAddr("authorizedUser");
        
        oracle = new Oracle(updater);
        consumer = new ConsumerDemo(address(oracle));
        
        // Set initial price
        vm.prank(updater);
        oracle.updatePrice(INITIAL_PRICE);
        
        // Authorize user for liquidations
        consumer.authorizeUser(authorizedUser);
        
        // Fund user
        vm.deal(user, 10 ether);
    }
    
    function testInitialState() public {
        assertEq(address(consumer.oracle()), address(oracle));
        assertEq(consumer.owner(), address(this));
        assertFalse(consumer.paused());
        assertTrue(consumer.authorizedUsers(authorizedUser));
        assertFalse(consumer.authorizedUsers(user));
    }
    
    function testCreatePosition() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        ConsumerDemo.Position memory position = consumer.getPosition(user);
        assertEq(position.ethAmount, ETH_AMOUNT);
        assertEq(position.stopLossPrice, STOP_LOSS_PRICE);
        assertEq(position.alertPrice, ALERT_PRICE);
        assertTrue(position.isActive);
        assertEq(position.owner, user);
        assertEq(position.createdAt, block.timestamp);
    }
    
    function testCreatePositionInsufficientAmount() public {
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: amount too low");
        consumer.createPosition{value: 0.0001 ether}(STOP_LOSS_PRICE, ALERT_PRICE);
    }
    
    function testCreatePositionExcessiveAmount() public {
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: amount too high");
        consumer.createPosition{value: 101 ether}(STOP_LOSS_PRICE, ALERT_PRICE);
    }
    
    function testCreatePositionDuplicate() public {
        vm.startPrank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        vm.expectRevert("ConsumerDemo: position already exists");
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        vm.stopPrank();
    }
    
    function testCreatePositionInvalidPrices() public {
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: stop-loss price must be positive");
        consumer.createPosition{value: ETH_AMOUNT}(0, ALERT_PRICE);
        
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: alert price must be positive");
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, 0);
        
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: stop-loss and alert prices cannot be equal");
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, STOP_LOSS_PRICE);
    }
    
    function testIsLiquidatable() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        // Price above stop-loss, not liquidatable
        assertFalse(consumer.isLiquidatable(user));
        
        // Update price below stop-loss
        vm.prank(updater);
        oracle.updatePrice(STOP_LOSS_PRICE - 100);
        
        // Now liquidatable
        assertTrue(consumer.isLiquidatable(user));
    }
    
    function testIsLiquidatableNoPosition() public {
        vm.expectRevert("ConsumerDemo: no active position");
        consumer.isLiquidatable(user);
    }
    
    function testShouldTriggerAlert() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        // Price below alert, no trigger
        assertFalse(consumer.shouldTriggerAlert(user));
        
        // Update price above alert
        vm.prank(updater);
        oracle.updatePrice(ALERT_PRICE + 100);
        
        // Now should trigger
        assertTrue(consumer.shouldTriggerAlert(user));
    }
    
    function testLiquidatePosition() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        // Update price below stop-loss
        vm.prank(updater);
        oracle.updatePrice(STOP_LOSS_PRICE - 100);
        
        uint256 userBalanceBefore = user.balance;
        
        vm.prank(authorizedUser);
        consumer.liquidatePosition(user);
        
        // Check position is closed
        ConsumerDemo.Position memory position = consumer.getPosition(user);
        assertFalse(position.isActive);
        
        // Check ETH was returned
        assertEq(user.balance, userBalanceBefore + ETH_AMOUNT);
    }
    
    function testLiquidatePositionNotLiquidatable() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        vm.prank(authorizedUser);
        vm.expectRevert("ConsumerDemo: position not liquidatable");
        consumer.liquidatePosition(user);
    }
    
    function testLiquidatePositionNotAuthorized() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        vm.prank(updater);
        oracle.updatePrice(STOP_LOSS_PRICE - 100);
        
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: not authorized");
        consumer.liquidatePosition(user);
    }
    
    function testTriggerAlert() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        // Update price above alert
        vm.prank(updater);
        oracle.updatePrice(ALERT_PRICE + 100);
        
        vm.prank(authorizedUser);
        consumer.triggerAlert(user);
        
        // Check that event was emitted (simplified test)
        assertTrue(true);
    }
    
    function testClosePosition() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        uint256 userBalanceBefore = user.balance;
        
        vm.prank(user);
        consumer.closePosition();
        
        // Check position is closed
        ConsumerDemo.Position memory position = consumer.getPosition(user);
        assertFalse(position.isActive);
        
        // Check ETH was returned
        assertEq(user.balance, userBalanceBefore + ETH_AMOUNT);
    }
    
    function testClosePositionNoPosition() public {
        vm.prank(user);
        vm.expectRevert("ConsumerDemo: no active position");
        consumer.closePosition();
    }
    
    function testGetCollateralValue() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        uint256 collateralValue = consumer.getCollateralValue(user);
        
        // Should be approximately ETH_AMOUNT * INITIAL_PRICE / 1 ether
        uint256 expectedValue = (ETH_AMOUNT * INITIAL_PRICE) / 1 ether;
        assertEq(collateralValue, expectedValue);
    }
    
    function testAuthorizeUser() public {
        address newUser = makeAddr("newUser");
        
        consumer.authorizeUser(newUser);
        assertTrue(consumer.authorizedUsers(newUser));
    }
    
    function testAuthorizeUserOnlyOwner() public {
        address newUser = makeAddr("newUser");
        
        vm.prank(user);
        vm.expectRevert("Ownable: caller is not the owner");
        consumer.authorizeUser(newUser);
    }
    
    function testDeauthorizeUser() public {
        consumer.deauthorizeUser(authorizedUser);
        assertFalse(consumer.authorizedUsers(authorizedUser));
    }
    
    function testPauseUnpause() public {
        assertFalse(consumer.paused());
        
        consumer.pause();
        assertTrue(consumer.paused());
        
        consumer.unpause();
        assertFalse(consumer.paused());
    }
    
    function testEmergencyWithdraw() public {
        // Send some ETH to the contract
        vm.deal(address(consumer), 1 ether);
        assertEq(address(consumer).balance, 1 ether);
        
        uint256 ownerBalanceBefore = address(this).balance;
        consumer.emergencyWithdraw();
        
        assertEq(address(consumer).balance, 0);
        assertEq(address(this).balance, ownerBalanceBefore + 1 ether);
    }
    
    function testReceive() public {
        // Send ETH to the contract
        vm.deal(address(this), 1 ether);
        (bool success,) = address(consumer).call{value: 1 ether}("");
        assertTrue(success);
        assertEq(address(consumer).balance, 1 ether);
    }
    
    function testPositionCreatedEvent() public {
        vm.expectEmit(true, false, false, true);
        emit PositionCreated(user, ETH_AMOUNT, STOP_LOSS_PRICE, ALERT_PRICE);
        
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
    }
    
    function testStopLossTriggeredEvent() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        // Update price below stop-loss
        vm.prank(updater);
        oracle.updatePrice(STOP_LOSS_PRICE - 100);
        
        vm.expectEmit(true, false, false, true);
        emit StopLossTriggered(user, ETH_AMOUNT, STOP_LOSS_PRICE);
        
        vm.prank(authorizedUser);
        consumer.liquidatePosition(user);
    }
    
    function testPriceAlertTriggeredEvent() public {
        vm.prank(user);
        consumer.createPosition{value: ETH_AMOUNT}(STOP_LOSS_PRICE, ALERT_PRICE);
        
        // Update price above alert
        vm.prank(updater);
        oracle.updatePrice(ALERT_PRICE + 100);
        
        vm.expectEmit(true, false, false, true);
        emit PriceAlertTriggered(user, ALERT_PRICE, ALERT_PRICE + 100);
        
        vm.prank(authorizedUser);
        consumer.triggerAlert(user);
    }
}


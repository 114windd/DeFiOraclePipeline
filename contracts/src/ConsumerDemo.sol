// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./Oracle.sol";

/**
 * @title Consumer Demo Contract
 * @dev Demonstrates usage of the ETH/USD Oracle for stop-loss and price alerts
 * @author DeFi Oracle Pipeline
 */
contract ConsumerDemo is Ownable, Pausable, ReentrancyGuard {
    // Oracle contract interface
    Oracle public immutable oracle;
    
    // User position structure
    struct Position {
        uint256 ethAmount;           // Amount of ETH deposited
        uint256 stopLossPrice;       // Stop-loss price (8 decimals)
        uint256 alertPrice;          // Price alert threshold (8 decimals)
        bool isActive;               // Whether position is active
        uint256 createdAt;           // Creation timestamp
        address owner;               // Position owner
    }
    
    // State variables
    mapping(address => Position) public positions;
    mapping(address => bool) public authorizedUsers;
    
    uint256 public constant PRICE_DECIMALS = 8;
    uint256 public constant MIN_ETH_AMOUNT = 0.001 ether; // Minimum 0.001 ETH
    uint256 public constant MAX_ETH_AMOUNT = 100 ether;   // Maximum 100 ETH
    
    // Events
    event PositionCreated(address indexed user, uint256 ethAmount, uint256 stopLossPrice, uint256 alertPrice);
    event StopLossTriggered(address indexed user, uint256 ethAmount, uint256 triggerPrice);
    event PriceAlertTriggered(address indexed user, uint256 alertPrice, uint256 currentPrice);
    event PositionClosed(address indexed user, uint256 ethAmount);
    event UserAuthorized(address indexed user);
    event UserDeauthorized(address indexed user);
    
    // Modifiers
    modifier onlyAuthorized() {
        require(authorizedUsers[msg.sender] || msg.sender == owner(), "ConsumerDemo: not authorized");
        _;
    }
    
    modifier validPosition(address user) {
        require(positions[user].isActive, "ConsumerDemo: no active position");
        _;
    }
    
    modifier validAmount(uint256 amount) {
        require(amount >= MIN_ETH_AMOUNT, "ConsumerDemo: amount too low");
        require(amount <= MAX_ETH_AMOUNT, "ConsumerDemo: amount too high");
        _;
    }
    
    /**
     * @dev Constructor sets the Oracle contract address
     * @param _oracle Address of the Oracle contract
     */
    constructor(address _oracle) {
        require(_oracle != address(0), "ConsumerDemo: oracle cannot be zero address");
        oracle = Oracle(_oracle);
    }
    
    /**
     * @dev Creates a new position with stop-loss and price alert
     * @param _stopLossPrice Stop-loss price (8 decimals)
     * @param _alertPrice Price alert threshold (8 decimals)
     */
    function createPosition(uint256 _stopLossPrice, uint256 _alertPrice) 
        external 
        payable 
        whenNotPaused 
        nonReentrant 
        validAmount(msg.value) 
    {
        require(_stopLossPrice > 0, "ConsumerDemo: stop-loss price must be positive");
        require(_alertPrice > 0, "ConsumerDemo: alert price must be positive");
        require(_stopLossPrice != _alertPrice, "ConsumerDemo: stop-loss and alert prices cannot be equal");
        
        // Check if user already has an active position
        require(!positions[msg.sender].isActive, "ConsumerDemo: position already exists");
        
        // Create new position
        positions[msg.sender] = Position({
            ethAmount: msg.value,
            stopLossPrice: _stopLossPrice,
            alertPrice: _alertPrice,
            isActive: true,
            createdAt: block.timestamp,
            owner: msg.sender
        });
        
        emit PositionCreated(msg.sender, msg.value, _stopLossPrice, _alertPrice);
    }
    
    /**
     * @dev Checks if a position should be liquidated (stop-loss triggered)
     * @param user Address of the position owner
     * @return True if position should be liquidated
     */
    function isLiquidatable(address user) 
        public 
        view 
        validPosition(user) 
        returns (bool) 
    {
        Position memory position = positions[user];
        
        // Get current price from oracle
        (uint256 currentPrice,,) = oracle.getLatestPrice();
        
        // Check if price is below stop-loss
        return currentPrice <= position.stopLossPrice;
    }
    
    /**
     * @dev Checks if price alert should be triggered
     * @param user Address of the position owner
     * @return True if alert should be triggered
     */
    function shouldTriggerAlert(address user) 
        public 
        view 
        validPosition(user) 
        returns (bool) 
    {
        Position memory position = positions[user];
        
        // Get current price from oracle
        (uint256 currentPrice,,) = oracle.getLatestPrice();
        
        // Check if price has reached alert threshold
        return currentPrice >= position.alertPrice;
    }
    
    /**
     * @dev Liquidates a position (stop-loss execution)
     * @param user Address of the position owner
     */
    function liquidatePosition(address user) 
        external 
        onlyAuthorized 
        whenNotPaused 
        nonReentrant 
        validPosition(user) 
    {
        require(isLiquidatable(user), "ConsumerDemo: position not liquidatable");
        
        Position memory position = positions[user];
        
        // Mark position as inactive
        positions[user].isActive = false;
        
        // In a real implementation, you would:
        // 1. Execute the stop-loss logic (e.g., swap ETH to USDC)
        // 2. Send the resulting tokens to the user
        // 3. Keep a small fee for the service
        
        // For demo purposes, we'll just refund the ETH
        (bool success, ) = payable(user).call{value: position.ethAmount}("");
        require(success, "ConsumerDemo: ETH transfer failed");
        
        emit StopLossTriggered(user, position.ethAmount, position.stopLossPrice);
    }
    
    /**
     * @dev Triggers a price alert
     * @param user Address of the position owner
     */
    function triggerAlert(address user) 
        external 
        onlyAuthorized 
        whenNotPaused 
        validPosition(user) 
    {
        require(shouldTriggerAlert(user), "ConsumerDemo: alert not triggered");
        
        Position memory position = positions[user];
        (uint256 currentPrice,,) = oracle.getLatestPrice();
        
        emit PriceAlertTriggered(user, position.alertPrice, currentPrice);
    }
    
    /**
     * @dev Closes a position manually
     */
    function closePosition() 
        external 
        whenNotPaused 
        nonReentrant 
        validPosition(msg.sender) 
    {
        Position memory position = positions[msg.sender];
        
        // Mark position as inactive
        positions[msg.sender].isActive = false;
        
        // Refund ETH to user
        (bool success, ) = payable(msg.sender).call{value: position.ethAmount}("");
        require(success, "ConsumerDemo: ETH transfer failed");
        
        emit PositionClosed(msg.sender, position.ethAmount);
    }
    
    /**
     * @dev Gets the collateral value of a position in USD
     * @param user Address of the position owner
     * @return Value in USD (8 decimals)
     */
    function getCollateralValue(address user) 
        external 
        view 
        validPosition(user) 
        returns (uint256) 
    {
        Position memory position = positions[user];
        (uint256 currentPrice,,) = oracle.getLatestPrice();
        
        // Calculate value: ETH amount * current price
        // Note: This is a simplified calculation
        return (position.ethAmount * currentPrice) / 1 ether;
    }
    
    /**
     * @dev Gets position information
     * @param user Address of the position owner
     * @return Position data
     */
    function getPosition(address user) 
        external 
        view 
        returns (Position memory) 
    {
        return positions[user];
    }
    
    /**
     * @dev Authorizes a user to trigger liquidations and alerts
     * @param user Address to authorize
     */
    function authorizeUser(address user) external onlyOwner {
        require(user != address(0), "ConsumerDemo: user cannot be zero address");
        authorizedUsers[user] = true;
        emit UserAuthorized(user);
    }
    
    /**
     * @dev Deauthorizes a user
     * @param user Address to deauthorize
     */
    function deauthorizeUser(address user) external onlyOwner {
        authorizedUsers[user] = false;
        emit UserDeauthorized(user);
    }
    
    /**
     * @dev Pauses the contract
     */
    function pause() external onlyOwner {
        _pause();
    }
    
    /**
     * @dev Unpauses the contract
     */
    function unpause() external onlyOwner {
        _unpause();
    }
    
    /**
     * @dev Emergency function to withdraw any ETH sent to the contract
     */
    function emergencyWithdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "ConsumerDemo: no ETH to withdraw");
        
        (bool success, ) = payable(owner()).call{value: balance}("");
        require(success, "ConsumerDemo: ETH withdrawal failed");
    }
    
    /**
     * @dev Fallback function to receive ETH
     */
    receive() external payable {
        // Contract can receive ETH for positions
    }
}


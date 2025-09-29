// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * @title ETH/USD Oracle Contract
 * @dev This contract stores ETH/USD price data and provides access control
 * @author DeFi Oracle Pipeline
 */
contract Oracle is Ownable, Pausable, ReentrancyGuard {
    // Price data structure
    struct PriceData {
        uint256 price;      // Price with 8 decimal precision
        uint256 timestamp;  // Unix timestamp
        uint256 roundId;    // Round ID for tracking updates
    }

    // State variables
    PriceData public latestPrice;
    address public updater;
    uint256 public constant PRICE_DECIMALS = 8;
    uint256 public constant MAX_PRICE = 1000000 * 10**PRICE_DECIMALS; // $1M max
    uint256 public constant MIN_PRICE = 1 * 10**PRICE_DECIMALS; // $1 min
    uint256 public constant MAX_AGE = 1 hours; // Max age for price data
    
    // Events
    event PriceUpdated(uint256 indexed price, uint256 indexed timestamp, uint256 indexed roundId);
    event UpdaterChanged(address indexed oldUpdater, address indexed newUpdater);
    event EmergencyWithdraw(address indexed token, uint256 amount);

    // Modifiers
    modifier onlyUpdater() {
        require(msg.sender == updater, "Oracle: caller is not the updater");
        _;
    }

    modifier validPrice(uint256 _price) {
        require(_price >= MIN_PRICE, "Oracle: price too low");
        require(_price <= MAX_PRICE, "Oracle: price too high");
        _;
    }

    modifier notStale() {
        require(block.timestamp - latestPrice.timestamp <= MAX_AGE, "Oracle: price data is stale");
        _;
    }

    /**
     * @dev Constructor sets the initial updater address
     * @param _updater Address that can update prices
     */
    constructor(address _updater) {
        require(_updater != address(0), "Oracle: updater cannot be zero address");
        updater = _updater;
        
        // Initialize with zero price data
        latestPrice = PriceData({
            price: 0,
            timestamp: 0,
            roundId: 0
        });
    }

    /**
     * @dev Updates the ETH/USD price. Only callable by authorized updater.
     * @param _newPrice New price with 8 decimal precision
     */
    function updatePrice(uint256 _newPrice) 
        external 
        onlyUpdater 
        whenNotPaused 
        nonReentrant 
        validPrice(_newPrice) 
    {
        // Increment round ID
        latestPrice.roundId++;
        
        // Update price data
        latestPrice.price = _newPrice;
        latestPrice.timestamp = block.timestamp;
        
        emit PriceUpdated(_newPrice, block.timestamp, latestPrice.roundId);
    }

    /**
     * @dev Returns the most recent ETH/USD price and timestamp
     * @return price The latest price with 8 decimal precision
     * @return timestamp The timestamp of the latest price
     * @return roundId The round ID of the latest price
     */
    function getLatestPrice() 
        external 
        view 
        returns (uint256 price, uint256 timestamp, uint256 roundId) 
    {
        return (latestPrice.price, latestPrice.timestamp, latestPrice.roundId);
    }

    /**
     * @dev Returns the latest price with staleness check
     * @return price The latest price with 8 decimal precision
     * @return timestamp The timestamp of the latest price
     * @return roundId The round ID of the latest price
     */
    function getLatestPriceSafe() 
        external 
        view 
        returns (uint256 price, uint256 timestamp, uint256 roundId) 
    {
        require(block.timestamp - latestPrice.timestamp <= MAX_AGE, "Oracle: price data is stale");
        return (latestPrice.price, latestPrice.timestamp, latestPrice.roundId);
    }

    /**
     * @dev Allows the contract owner to change the authorized updater address
     * @param _newUpdater New updater address
     */
    function setUpdater(address _newUpdater) external onlyOwner {
        require(_newUpdater != address(0), "Oracle: new updater cannot be zero address");
        address oldUpdater = updater;
        updater = _newUpdater;
        emit UpdaterChanged(oldUpdater, _newUpdater);
    }

    /**
     * @dev Pauses the contract in case of emergency
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
     * @dev Checks if the price data is stale
     * @return True if price is older than MAX_AGE
     */
    function isStale() external view returns (bool) {
        return block.timestamp - latestPrice.timestamp > MAX_AGE;
    }

    /**
     * @dev Returns the age of the latest price in seconds
     * @return Age in seconds
     */
    function getPriceAge() external view returns (uint256) {
        return block.timestamp - latestPrice.timestamp;
    }

    /**
     * @dev Returns the current round ID
     * @return Current round ID
     */
    function getCurrentRoundId() external view returns (uint256) {
        return latestPrice.roundId;
    }

    /**
     * @dev Emergency function to withdraw any ETH sent to the contract
     */
    function emergencyWithdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "Oracle: no ETH to withdraw");
        
        (bool success, ) = payable(owner()).call{value: balance}("");
        require(success, "Oracle: ETH withdrawal failed");
        
        emit EmergencyWithdraw(address(0), balance);
    }

    /**
     * @dev Emergency function to withdraw any ERC20 tokens sent to the contract
     * @param _token Token contract address
     */
    function emergencyWithdrawToken(address _token) external onlyOwner {
        require(_token != address(0), "Oracle: token address cannot be zero");
        
        // This would require IERC20 interface
        // For now, this is a placeholder
        emit EmergencyWithdraw(_token, 0);
    }

    /**
     * @dev Returns contract information
     * @return updaterAddress Current updater address
     * @return isPaused Whether contract is paused
     * @return maxAge Maximum age for price data
     * @return minPrice Minimum allowed price
     * @return maxPrice Maximum allowed price
     */
    function getContractInfo() 
        external 
        view 
        returns (
            address updaterAddress,
            bool isPaused,
            uint256 maxAge,
            uint256 minPrice,
            uint256 maxPrice
        ) 
    {
        return (
            updater,
            paused(),
            MAX_AGE,
            MIN_PRICE,
            MAX_PRICE
        );
    }

    /**
     * @dev Fallback function to receive ETH
     */
    receive() external payable {
        // Contract can receive ETH but doesn't do anything with it
        // Owner can withdraw using emergencyWithdraw()
    }
}

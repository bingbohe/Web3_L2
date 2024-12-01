// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
// ERC20代币库
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
// 权限库
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyToken is ERC20, Ownable {
    uint256 public constant RATE = 100000000; // 100000000 MyToken per 1 ETH
    uint256 public constant MIN_ETH = 0.00000001 ether;

    // 定义事件
    event MintSuccess(address indexed  user, uint256 ethAmount, uint256 tokenAmount);
    event MintFailed(address indexed user, string reason);
    event ETHWithdraw(address indexed owner, uint256 amount);


    constructor(address initialOwner) ERC20("RCCDemoToken", "RDT") Ownable(msg.sender) {
    }
    // 铸币函数，把派生的token发送到调用地址中
    function mint() public payable {
        if (msg.value < MIN_ETH) {
            emit MintFailed(msg.sender, "Not enough ETH sent"); // 触发失败事件
            revert("Not enough ETH sent");
        }
        
        uint256 tokensToMint = (msg.value * RATE);
        _mint(msg.sender, tokensToMint);

        emit MintSuccess(msg.sender, msg.value, tokensToMint); // 触发成功事件
    }
    // 合约所有者从合约中取钱
    function withdrawETH() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No ETH to withdraw");
        payable(owner()).transfer(balance);
        // 触发事件
        emit ETHWithdraw(owner(), balance);

    }
    // 用户向合约发送ETH的时候，调用receive函数
    receive() external payable {
        mint();
    }
}
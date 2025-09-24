// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Storage {
    mapping(uint256 => string) public data;

    function setData(uint256 key, string calldata value) external {
        data[key] = value;
    }

    function getData(uint256 key) external view returns (string memory) {
        return data[key];
    }
}

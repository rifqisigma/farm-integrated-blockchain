// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract MessageStore {
    struct Message {
        uint256 id;
        string status;
        string text;
        address from;
        address to;
        uint256 createdAt;
    }

    Message[] public messages;

    event MessageCreated(uint256 id, address from, address to);

    function createMessage(
        uint256 _id,
        string memory _status,
        string memory _text,
        address _to
    ) public {
        messages.push(
            Message({
                id: _id,
                status: _status,
                text: _text,
                from: msg.sender,
                to: _to,
                createdAt: block.timestamp
            })
        );
        emit MessageCreated(_id, msg.sender, _to);
    }

    function getMessage(uint256 index)
        public
        view
        returns (uint256, string memory, string memory, address, address, uint256)
    {
        Message memory m = messages[index];
        return (m.id, m.status, m.text, m.from, m.to, m.createdAt);
    }

    function getCount() public view returns (uint256) {
        return messages.length;
    }
}

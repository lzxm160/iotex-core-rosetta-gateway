pragma solidity >=0.4.22 <0.6.0;
contract Multisend {
    event Transfer(address recipient,uint amount);
    event Refund(uint refund);
    event Payload(string payload);
    function multiSend(address payable[] memory recipients, uint[] memory amounts,string memory payload) public payable{
        require(recipients.length <= 300, "number of recipients is larger than 300");
        require(recipients.length == amounts.length, "parameters not match");
        uint totalAmount = 0;
        for(uint i = 0; i < recipients.length; i++) {
            totalAmount+= amounts[i];
        }
        require(msg.value >= totalAmount, "not enough token");
        uint refund = msg.value - totalAmount;
        for(uint i = 0; i < recipients.length; i++) {
            recipients[i].transfer(amounts[i]);
            emit Transfer(recipients[i],amounts[i]);
        }
        if (refund>0) {
            msg.sender.transfer(refund);
            emit Refund(refund);
        }
        emit Payload(payload);
    }
    function () external payable {
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

contract NFTWhiteList {
    bytes32 public merkleRoot;
    address public owner;

    constructor() {
        owner = msg.sender;   
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Sorry, Only Admin can change the Merkle Root");
        _;
    }

    function SetMerkleRoot(bytes32 _merkleRoot) public onlyOwner{
        merkleRoot = _merkleRoot;
    }

    function GetMerkleRoot() public view returns (bytes32) {
        return merkleRoot;
    }
    
    function VerifyWhitelist(bytes32[] memory _proof) private view returns(bool){
        bytes32 check;

        check = keccak256(abi.encodePacked(msg.sender));
        uint i=0;
        for (i; i<_proof.length; i++){
            if (check < _proof[i]){
                check = keccak256(abi.encodePacked(check, _proof[i]));
            }
            else{
                check = keccak256(abi.encodePacked(_proof[i], check));
            }
        }
        return check==merkleRoot;
    }

    function SayHelloToWhiteListed(bytes32[] memory _proof) public view returns(string memory){
        require(merkleRoot!= 0x0000000000000000000000000000000000000000000000000000000000000000, "It's not started yet.");
        require(VerifyWhitelist(_proof), "You aren't white listed");
        return("Congradulations!!! You've been white listed!");
    }
    
}


// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

contract NFTWhiteList {
    bytes32 public merkleRoot;
    address public owner;
    bytes32[] private addList;
    bytes32 private nullBytes = 0x0000000000000000000000000000000000000000000000000000000000000000;
    
    constructor() {
        owner = msg.sender;   
    }

    struct merkleNode {
        bytes32 leftHash;
        bytes32 rightHash;
        bytes32 parentHash;
        bytes32 nodeHash;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Sorry, Only Admin can change the Merkle Root");
        _;
    }

    function ConstructMerkleTree(merkleNode[] memory leaves) private view returns(merkleNode[] memory){
        uint len = leaves.length / 2;
        merkleNode[] memory tree = new merkleNode[](len);
        if (leaves.length == 1){
            return leaves;
        }

        for (uint i=0; i< leaves.length; i+2){
            bytes32 parentHash;
            if (leaves[i].nodeHash < leaves[i+1].nodeHash){
                parentHash =  keccak256(abi.encodePacked(leaves[i].nodeHash, leaves[i+1].nodeHash));
                tree[i/2]=(merkleNode(leaves[i].nodeHash, leaves[i+1].nodeHash, nullBytes, parentHash));
            }
            else{
                parentHash =  keccak256(abi.encodePacked(leaves[i+1].nodeHash, leaves[i].nodeHash));
                tree[i/2]=(merkleNode(leaves[i+1].nodeHash, leaves[i].nodeHash, nullBytes, parentHash));
            }
            
        }
        return ConstructMerkleTree(tree);
    }

    function BuildTree() public onlyOwner view returns(merkleNode[] memory){
        merkleNode[] memory tree;
        merkleNode[] memory leaves;
        for (uint i=0; i<addList.length; i++) {
            leaves[i] = merkleNode(nullBytes, nullBytes, nullBytes, keccak256(abi.encodePacked(addList[i])));
        }
        tree = ConstructMerkleTree(leaves);
        return tree;
    }

    function AddAddressToWhiteList(address newAdd) public onlyOwner{
        addList.push(keccak256(abi.encodePacked(newAdd)));
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
        require(merkleRoot!= nullBytes, "It's not started yet.");
        require(VerifyWhitelist(_proof), "You aren't white listed");
        return("Congradulations!!! You've been white listed!");
    }
    
}


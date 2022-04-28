// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// import "hardhat/console.sol";

contract NFTWhiteList {
    bytes32[] public listingRequests;
    bytes32 public merkleRoot;
    address public owner;
    uint public addedItems = 0;

    mapping(bytes32 => bytes32[]) private proofs;

    constructor(uint _max_allowed){
        owner = msg.sender;

        if (_max_allowed%2 != 0) _max_allowed++;
        listingRequests = new bytes32[](_max_allowed);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Sorry, Only Admin can change the Merkle Root");
        _;
    }

    function SetMerkleRoot(bytes32 _root) private {
        merkleRoot = _root;
    }

    function GetMerkleRoot() public view onlyOwner returns (bytes32){
        return merkleRoot;
    }

    function getProof() public view returns(bytes32[] memory){
        bytes32 convertedAdd = keccak256(abi.encodePacked(msg.sender));
        return proofs[convertedAdd];
    }

    function calculateArrayLength(uint initial_leaves) private pure returns(uint, uint){
        uint len = 0;
        uint iterations = 0;
        while ( initial_leaves > 0) {
            if (initial_leaves == 1){
                iterations ++;
                len ++;
                return (len, iterations);
            }
            if (initial_leaves%2 != 0){
                initial_leaves ++;
            }
            iterations ++;
            len += initial_leaves;
            initial_leaves /=2;
        }
        return (len, iterations);
    }

    function GetParentIndex(bytes32[] memory tree, bytes32 parent, uint currentIndex) private pure returns(uint){
        for (uint i=currentIndex; i<tree.length; i++){
            if (parent == tree[i]) return i;
        }
        return tree.length +1;
    }


    function GenerateProof(bytes32[] memory tree, uint _iterations) private {
        uint i = 0;
        bytes32[] memory _proof;
        while (i < addedItems){
            uint iter = 0;
            uint index = i;
            _proof = new bytes32[](_iterations-1);
            while (iter<(_iterations-1)){
                bytes32 parent;
                if (index%2 == 0) {
                    _proof[iter] = tree[index+1];
                    parent = keccak256(abi.encodePacked(tree[index], tree[index+1]));
                    if (tree[index] > tree[index+1]) parent = keccak256(abi.encodePacked(tree[index+1], tree[index]));
                }
                else {
                    _proof[iter] = tree[index-1];
                    parent = keccak256(abi.encodePacked(tree[index], tree[index-1]));
                    if (tree[index] > tree[index-1]) parent = keccak256(abi.encodePacked(tree[index-1], tree[index]));
                }
                index = GetParentIndex(tree, parent, index);
                iter ++;
            }
            proofs[tree[i]] = _proof;
            i++;
        }
    }

    function BuildMerkleTree() public onlyOwner{
        uint len = 0;
        uint iterations = 0;
        uint currentIndex = 0;
        bytes32 nullBytes = 0x0000000000000000000000000000000000000000000000000000000000000000;
        if (addedItems%2 != 0) {
            addedItems++;
            // handle adding null bytes32 to be even numbers of leaves;
            listingRequests.push(keccak256(abi.encodePacked(nullBytes)));
        }

        (len, iterations) = calculateArrayLength(addedItems);
        bytes32[] memory merkleTree = new bytes32[](len);
        for (uint i = 0; i < addedItems; i++){
            merkleTree[i] = listingRequests[i];
            currentIndex ++;
        }
        
        uint j = 0;
        uint start_index = 0;
        uint end_index = currentIndex;
        while(currentIndex < len){
            j=start_index;
            while(j < end_index){
                bytes32 parent;
                if (merkleTree[j] < merkleTree[j+1]) parent = keccak256(abi.encodePacked(merkleTree[j], merkleTree[j+1]));
                else parent = keccak256(abi.encodePacked(merkleTree[j+1], merkleTree[j]));
                merkleTree[currentIndex] = parent;
                currentIndex++;
                j += 2;
            }
            start_index = end_index;
            if (currentIndex%2 != 0 && currentIndex != len) {
                merkleTree[currentIndex] = keccak256(abi.encodePacked(nullBytes));
                currentIndex++;
            }
            end_index = currentIndex;
            }
        SetMerkleRoot(merkleTree[len-1]);
        GenerateProof(merkleTree, iterations);
    }

    function AddAddressToWhiteList(bytes memory addr) public onlyOwner{
        require (addedItems < listingRequests.length, "Can't add more");

        listingRequests[addedItems] = (keccak256(abi.encodePacked(addr)));
        addedItems ++;
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
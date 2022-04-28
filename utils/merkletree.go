package utils

import (
	"encoding/hex"
	"fmt"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)
type merkleNode struct {
	Left	*merkleNode	`json:"left"`
	Right	*merkleNode	`json:"right"`
	Parent	string	`json:"parent"`
	Hash	string		`json:"hash"`
}

func (m *merkleNode) UpdateParent(parent string) merkleNode{
	return merkleNode{m.Left,m.Right,parent,m.Hash}
}

func (m *merkleNode) Modify(parent string) {
	m.Parent = parent
}

type merkleTree struct {
	Root	*merkleNode	`json:"root"`
}

func calculateMerkleTree(leaves []merkleNode, tree *[]merkleNode) merkleTree{
	// This methods accept two list variables, the first one is the leaves that is used
	// to calculate the Merkle Tree and the second, tree, is a pointer to store the whole structure of three.
	// Size of leaves will decrease during recursive calls, whereas tree will grow.
	if len(leaves) == 1 {
		merkleRoot := merkleTree{&leaves[0]}
		*tree = append(*tree, *merkleRoot.Root)
		return merkleRoot
	}
	var parents []merkleNode

	// if len(leaves) == 2 {
	// }
	
	for i:=0; i < len(leaves); i+=2{
		v1, v2 := SortValues(leaves[i].Hash,leaves[i+1].Hash)
		hash := solsha3.SoliditySHA3(
			// types
			[]string{"bytes32", "bytes32"},
			// values
			[]interface{}{
				v1,v2,
			},
		)
		hashString := "0x" + hex.EncodeToString(hash)
		merkleParent := merkleNode{&leaves[i], &leaves[i+1], "", hashString}
		leaves[i].Modify(hashString)
		leaves[i+1].Modify(hashString)
		// leaves[i] = leaves[i].UpdateParent(merkleParent)
		// leaves[i+1] = leaves[i+1].UpdateParent(merkleParent)
		parents = append(parents, merkleParent)
	}
	
	mt := calculateMerkleTree(parents, tree)
	// Adding to the Tree
	*tree = append(*tree, leaves...)
	fmt.Println(parents)
	return mt
}

func Find(tree []merkleNode, hash string) int {
    for i, item := range tree {
        if hash == item.Hash {
            return i
        }
    }
    return len(tree)
}

func (node *merkleNode) getNodeSibling(tree []merkleNode) (*merkleNode, merkleNode, bool){
	// Get a merkle node and return its sibling through parent node
	// If curent node is the merkle root, it return the same node with false
	isRoot := false
	if node.Parent == "" {
		isRoot = true
		return node, *node, isRoot
	}
	parent := tree[Find(tree, node.Parent)]
	leftSibling := parent.Left
	if leftSibling.Hash == node.Hash {
		return parent.Right, parent, isRoot
	}
	return leftSibling, parent, isRoot
}

func (leaf *merkleNode) generateProof(proof *[]string, tree []merkleNode){
	// It's a recursive method to get the Inclusive Proof of a leaf node
	// and return a string of needed hash to calculate the merkle root
	sibling, parent, isRoot := leaf.getNodeSibling(tree)
	if !isRoot{
		*proof = append(*proof, sibling.Hash)
		parent.generateProof(proof, tree)
	}
}


func hashAddresses(addresses []string) []merkleNode {
	// This Method Calculate hash for all addresses based on solidity keccak256 methods.
	var nodes []merkleNode
	for _, add := range addresses {
		solAddress := solsha3.Address(add)
		hash := solsha3.SoliditySHA3(solAddress)
		hashSting:= hex.EncodeToString(hash)
		nodes = append(nodes, merkleNode{nil, nil, "", "0x" + hashSting})
	}
	return nodes
}

func printTree(node merkleNode, indent int){
	indentation := func() string{
		ind := ""
		for i:=0; i< indent; i++{
			ind += "----"
		}
		return ind
	}
	fmt.Println("|"+indentation(), node)
	newIndent := indent + 1
	if node.Left != nil{
		printTree(*node.Left, newIndent)
	}
	if node.Right!= nil{
		printTree(*node.Right, newIndent)
	}
}

func convertStringToBytes32(st string) ([]byte){
	bytesValue := [32]byte{}
	copy(bytesValue[:], []byte(st))
	return bytesValue[:]
}

func SortValues(st1 string, st2 string) (string, string){
	// This method is used to sort pairs for constructing the tree
	// and makes things easy when verifying the proof.
	// and return (the lower value, bigger value) as implemented in proof function in solidity. 
	b1 := convertStringToBytes32(st1)
	b2 := convertStringToBytes32(st2)
	for i, _ := range b1 {
		if b1[i] > b2 [i]{
			return st2, st1
		} 
		if b2[i] > b1[i]{
			return st1, st2
		}
	}
	return st1, st2
}

// // var address []string
// // 	address = append(address, "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
// // 	address = append(address, "0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2")
// // 	address = append(address, "0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db")
// // 	address = append(address, "0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB")
// // 	address = append(address, "0x617F2E2fD72FD9D5503197092aC168c91465E7f2")
// // 	address = append(address, "0x17F6AD8Ef982297579C203069C1DbfFE4348c372")
// // 	address = append(address, "0x5c6B0f7Bf3E7ce046039Bd8FABdfD3f9F5021678")
// // 	address = append(address, "0x03C6FcED478cBbC9a4FAB34eF9f40767739D1Ff7")
// ["0x5B38Da6a701c568545dCfcB03FcB875f56beddC4","0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2","0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db","0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB","0x617F2E2fD72FD9D5503197092aC168c91465E7f2","0x17F6AD8Ef982297579C203069C1DbfFE4348c372","0x5c6B0f7Bf3E7ce046039Bd8FABdfD3f9F5021678","0x03C6FcED478cBbC9a4FAB34eF9f40767739D1Ff7"]

// ["0x999bf57501565dbd2fdcea36efa2b9aef8340a8901e3459f4a4c926275d36cdb","0x4726e4102af77216b09ccd94f40daa10531c87c4d60bba7f3b3faf5ff9f19b3c","0xc7348b15532ca2d26bca10e3d3b985a564813cd9a0f710f7acb7ab94ad00da56"]
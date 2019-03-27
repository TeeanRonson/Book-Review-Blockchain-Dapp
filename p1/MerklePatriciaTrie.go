package p1

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"reflect"
)

type Flag_value struct {
	encoded_prefix []uint8
	value string
}

type Node struct {
	node_type int // 0: Null, 1: Branch, 2: Ext or Leaf
	branch_value [17]string
	flag_value Flag_value
}

type MerklePatriciaTrie struct {
	db     map[string]Node
	Inputs map[string]string
	Root   string
}


func (mpt *MerklePatriciaTrie) GetHelper2(node string, path []uint8, position int) string {

	currNode := mpt.db[node]
	nodeType := currNode.node_type
	switch nodeType {
	case 0:
		return ""
	case 1:
		if len(path) == 0 {
			return currNode.branch_value[16]
		}
		nextHash := currNode.branch_value[path[0]]
		return mpt.GetHelper2(nextHash, path[1:], 0)
	case 2:
		isLeaf := isLeaf(currNode)
		nibbles := Compact_decode(currNode.flag_value.encoded_prefix)
		if reflect.DeepEqual(nibbles, path) && isLeaf {
			return currNode.flag_value.value
		} else if !reflect.DeepEqual(nibbles, path) && isLeaf {
			return ""
		} else {
			length := len(nibbles)
			nextHash := currNode.flag_value.value
			return mpt.GetHelper2(nextHash, path[length:], 0)
		}
	}
	return ""
}

/**
Traverses the MPT to find the value associated with the key
 */
func (mpt *MerklePatriciaTrie) GetHelper1(path []uint8) string {

	if path == nil || mpt.Root == "" {
		fmt.Println("Nothing")
		return ""
	}
	value := mpt.GetHelper2(mpt.Root, path, 0)
	return value
}

/*
Takes a key as the argument, traverses down the MPT to find the value
if the key doesnt exist, return an empty string
 */
func (mpt *MerklePatriciaTrie) Get(key string) string {
	path := EncodeToHex(key)
	fmt.Println("\nNew Get\nPath:", path)
	return mpt.GetHelper1(path)
}


/**
Insert Helper method
 */
func (mpt *MerklePatriciaTrie) insertHelp(parent string, currHash string, encodedKey []uint8, newValue string) (newHash string) {

	currNode := mpt.db[currHash]
	nodeType := currNode.node_type
	switch nodeType {
	case 0:
		encodedKey = append(encodedKey, 16)
		leaf := createNode(2, [17]string{}, encodedKey, newValue)
		mpt.addToMap(leaf)
		return leaf.hash_node()
	case 1:
		if len(encodedKey) == 0 {
			delete(mpt.db, currHash)
			currNode.branch_value[16] = newValue
			mpt.addToMap(currNode)
			return currNode.hash_node()
		}

		nextHash := currNode.branch_value[encodedKey[0]]
		if	nextHash == "" {
			delete(mpt.db, currHash)
			encodedKey = append(encodedKey, 16)
			newLeaf := createNode(2, [17]string{}, encodedKey[1:], newValue)
			mpt.addLeavesToBranch(newLeaf, &currNode, encodedKey[0])
			mpt.addToMap(newLeaf)
			mpt.addToMap(currNode)
			return currNode.hash_node()
		} else {
			newHash := mpt.insertHelp(currHash, nextHash, encodedKey[1:], newValue)
			if newHash != nextHash {
				delete(mpt.db, currHash)
				currNode.branch_value[encodedKey[0]] = newHash
				mpt.addToMap(currNode)
				return currNode.hash_node()
			}
			return currHash
		}
	case 2: //leaf/extension
		nibbles := Compact_decode(currNode.flag_value.encoded_prefix)
		match := findMatch(0, nibbles, encodedKey)
		if isLeaf(currNode) {
			if reflect.DeepEqual(encodedKey, nibbles) { //exact match
				delete(mpt.db, currHash)
				currNode.flag_value.value = newValue
				mpt.addToMap(currNode)
				return currNode.hash_node()
			} else if match == 0 { //no match
				return mpt.breakNodeNoMatch(currNode, nibbles, encodedKey, newValue, true)
			} else if match > 0 && int(match) <= len(nibbles) { //partial matches - 3 types
				if len(encodedKey[match:]) != 0 && len(nibbles[match:]) != 0 { //excess path and excess nibbles
					return mpt.breakNodeDoubleExcess(currNode, match, nibbles, encodedKey, newValue, true)
				} else if len(encodedKey[match:]) != 0 && len(nibbles[match:]) == 0 { //partial match with excess path only
					return mpt.breakLeafSingleExcess(currNode, match, nibbles, encodedKey, newValue, true)
				} else if len(encodedKey[match:]) == 0 && len(nibbles[match:]) != 0 { //partial match with excess nibbles
					return mpt.breakLeafSingleExcess(currNode, match, nibbles, encodedKey, newValue, false)
				}
			}
		} else { //is Extension
			if reflect.DeepEqual(nibbles, encodedKey) { //exact match
				nextHash := currNode.flag_value.value
				newHash := mpt.insertHelp(currHash, nextHash, encodedKey[match:], newValue)
				if newHash != nextHash {
					delete(mpt.db, currHash)
					currNode.flag_value.value = newHash
					mpt.addToMap(currNode)
					return currNode.hash_node()
				}
			} else if match == 0 { //no match
				return mpt.breakNodeNoMatch(currNode, nibbles, encodedKey, newValue, false)
			} else if match > 0 && int(match) <= len(nibbles) {
				if len(encodedKey[match:]) != 0 && len(nibbles[match:]) != 0 {
					return mpt.breakNodeDoubleExcess(currNode, match, nibbles, encodedKey, newValue, false)
				} else if len(encodedKey[match:]) != 0 && len(nibbles[match:]) == 0 {
					nextHash := currNode.flag_value.value
					newHash = mpt.insertHelp(currHash, nextHash, encodedKey[match:], newValue)
					if newHash != nextHash {
						delete(mpt.db, currHash)
						currNode.flag_value.value = newHash
						mpt.addToMap(currNode)
						return currNode.hash_node()
					}
				} else if len(encodedKey[match:]) == 0 && len(nibbles[match:]) != 0 {
					return mpt.breakExtSingleExcess(currNode, match, nibbles, encodedKey, newValue, false)
				}
			}
		}
	}
	return currHash
}

/*
Takes a pair of <key, value> as arguments. It will traverse down the MPT and find the right place to insert the value
 */
func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {

	if len(key) == 0 || len(new_value) == 0 {
		fmt.Println("Nil Key and Value given as input")
		return
	}
	//fmt.Println("\nNewInsertion")
	//Insert into extra database
	mpt.Inputs[key] = new_value

	encodedKey := EncodeToHex(key)
	fmt.Println("Insert Path in MPT:", encodedKey, new_value)
	newHash := mpt.insertHelp("", mpt.Root, encodedKey, new_value)
	if newHash != mpt.Root {
		mpt.Root = newHash
		//fmt.Println("Newhash:", newHash)
		//fmt.Println("DB final:", mpt.db)
	}
}

/**
Delete helper method
If we find the path, delete the value, make new adjustments and send up a new hashvalue with nil error
If we dont find the path, then we send up the same hashvalue with an error value
 */
func (mpt *MerklePatriciaTrie) deleteHelper(parent string, currHash string, path []uint8) (string, Node, error) {

	currNode := mpt.db[currHash]
	nodeType := currNode.node_type
	switch nodeType {
	case 0:
		return "", currNode, errors.New("Path Not Found")
	case 1:
		if len(path) == 0 {
			delete(mpt.db, currHash)
			currNode.branch_value[16] = ""
			branchCount := branchItems(currNode)
			return mpt.checkBranch(branchCount, currNode)

		} else {
			nextHash := currNode.branch_value[path[0]]
			if nextHash == "" {
				return currHash, currNode, errors.New("Path Not Found")
			} else { //recurse down
				fmt.Println("2 & 3", path)
				newHash, child, err := mpt.deleteHelper(currHash, nextHash, path[1:])
				if err != nil || newHash == nextHash {
					return currHash, currNode, err
				} else {
					if newHash != nextHash {
						//update at the newHash
						fmt.Println("Child:", child)
						delete(mpt.db, currHash)
						currNode.branch_value[path[0]] = newHash
						branchCount := branchItems(currNode)
						return mpt.checkBranch(branchCount, currNode)
					}
				}
			}
		}
	case 2:
		nibbles := Compact_decode(currNode.flag_value.encoded_prefix)
		match := findMatch(0, nibbles, path)
		if isLeaf(currNode) {
			if reflect.DeepEqual(nibbles, path) { //exact match
				delete(mpt.db, currHash)
				return "", currNode, nil
			} else if len(path[match:]) > 0 || match == 0 {
				return currHash, currNode, errors.New("Path Not Found")
			}
		} else { //if extension
			fmt.Println("Extension")
			//If there are still more nibbles, return error
			if len(path[match:]) == 0 && len(nibbles[match:]) != 0 {
				return currHash, currNode, errors.New("Path Not Found")
			} else if reflect.DeepEqual(nibbles, path) || (len(nibbles[match:]) == 0 && len(path[match:]) != 0) {
				//we have an exact match or we have some extra path but no more nibbles
				//so we recurse further down
				nextHash := currNode.flag_value.value
				newHash, child, err := mpt.deleteHelper(currHash, nextHash, path[match:])
				if err != nil {
					return currHash, currNode, err
				}
				if newHash == "" {
					return "", currNode, nil
				}
				if newHash != nextHash {
					delete(mpt.db, currHash)
					switch child.node_type {
					//converting myself into a leaf node
					//	delete(mpt.db, currHash)
					//	nibbles = append(nibbles, 16)
					//	newLeaf := createNode(2, [17]string{}, nibbles, child.branch_value[16])
					//	mpt.addToMap(newLeaf)
					//	return newLeaf.hash_node(), newLeaf, nil
					case 1:
						currNode.flag_value.value = newHash
						mpt.addToMap(currNode)
						return currNode.hash_node(), currNode, nil
					case 2:
						//if my child is a leaf, merge myself with my child
						if isLeaf(child) {
							childNibbles := Compact_decode(child.flag_value.encoded_prefix)
							extNibbles := Compact_decode(currNode.flag_value.encoded_prefix)
							newNibbles := append(mergeArrays(extNibbles, childNibbles), 16)
							newLeaf := createNode(2, [17]string{}, newNibbles, child.flag_value.value)
							delete(mpt.db, child.hash_node())
							mpt.addToMap(newLeaf)
							return newLeaf.hash_node(), newLeaf, nil
						} else {
							//merge myself with my child
							//We might never need to get in here
							fmt.Println("Forbidden lands")
							delete(mpt.db, currHash)
							childNibbles := Compact_decode(child.flag_value.encoded_prefix)
							currNibbles := Compact_decode(currNode.flag_value.encoded_prefix)
							newNibbles := mergeArrays(currNibbles, childNibbles)
							newExtension := createNode(2, [17]string{}, newNibbles, child.flag_value.value)
							mpt.addToMap(newExtension)
							return newExtension.hash_node(), newExtension, nil
						}
					}
				}
			}
		}
	}
	return currHash, currNode, nil
}
/*
Function takes a key as the argument, traverses down the MPT and finds the Key.
If the key exists, delete the corresponding value and re-balance the Trie, if necessary.
if the key doesn't exist, return 'path_not_found'
 */
func (mpt *MerklePatriciaTrie) Delete(key string) (string, error) {
	fmt.Println("\nNewDeletion")
	if len(key) == 0 {
		return "", errors.New("No Input key")
	}
	encodedKey := EncodeToHex(key)
	fmt.Println("Delete Path:", encodedKey)
	newHash, child, err := mpt.deleteHelper("", mpt.Root, encodedKey)

	fmt.Println("Child:", child)
	if err != nil {
		return "", err
	}
	if newHash == "" {
		return "", errors.New("Tree is Empty")
	}
	if newHash != mpt.Root {
		mpt.Root = newHash
		fmt.Println("Newhash:", newHash)
		fmt.Println("DB final:", mpt.db)
		return "Successful Deletion", errors.New("")
	}
	return "", errors.New("Path Not found")
}
/*
Function takes an array of HEX value as the input, mark the Node type(Branch, Leaf, Extension),
makes sure the length is even, and converts it into an array of ASCII numbers as the output.
//If the last value is 16, it is a leaf node
 */
func Compact_encode(hex_array []uint8) []uint8 {
	term := 0
	if hex_array[len(hex_array)-1] == 16 {
		hex_array = hex_array[0: len(hex_array) - 1]
		term = 1
	}
	flags := make([]uint8, 0)
	oddlen := len(hex_array) % 2
	flags = append(flags, uint8(2*term+oddlen))
	if oddlen == 1 {
		hex_array = append(flags, hex_array...)
	} else {
		flags = append(flags, 0)
		hex_array = append(flags, hex_array...)
	}
	result := make([]uint8, 0)
	for i:= 0; i < len(hex_array); i += 2 {
		result = append(result, 16*hex_array[i]+hex_array[i+1])
	}
	return result
}


/*
Reverse the compact_encode function
 */
func Compact_decode(encoded_arr []uint8) []uint8 {
	unpack := ConvertToHex(encoded_arr)
	checkPrefix := 2 - unpack[0]&1
	return unpack[checkPrefix:]
}

func Test_compact_encode() {
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	fmt.Println(reflect.DeepEqual(Compact_decode(Compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
}

/*
 */
func (node *Node) hash_node() string {
	var str string
	switch node.node_type {
	case 0:
		str = ""
	case 1:
		str = "branch_"
		for _, v := range node.branch_value {
			str += v
		}
	case 2:
		str = node.flag_value.value
	}

	sum := sha3.Sum256([]byte(str))
	return "HashStart_" + hex.EncodeToString(sum[:]) + "_HashEnd"
}
package p1

/**
Checks the number if items in the branch node
 */
func branchItems(currNode Node) int {

    count := 0
    for _, value := range currNode.branch_value {
        if value != "" {
            count++
        }
    }
    return count
}

/**
Finds the non empty space in branch node
 */
 func findNonEmptySpace(currNode Node) uint8 {

     for i, value := range currNode.branch_value {
         if value != "" {
             return uint8(i)
         }
     }
     return 17
 }

 /**
 Merge two []uint8 arrays
  */
  func mergeArrays(one []uint8, two[]uint8) []uint8 {
      var result []uint8
      for _, value := range one {
          result = append(result, value)
      }
      for _, value := range two {
          result = append(result, value)
      }
      return result
  }

  func (mpt *MerklePatriciaTrie) mergeBranchAndLeaf(currNode Node, childNode Node, branchNibble []uint8) Node {
      //currNibble := []uint8{findNonEmptySpace(currNode)}
      //childNode := mpt.db[currNode.branch_value[currNibble[0]]]
      childNibbles  := Compact_decode(childNode.flag_value.encoded_prefix)
      newLeafNibbles := append(mergeArrays(branchNibble, childNibbles), 16)
      mergedLeaf := createNode(2, [17]string{}, newLeafNibbles, childNode.flag_value.value)
      delete(mpt.db, currNode.hash_node())
      delete(mpt.db, childNode.hash_node())
      mpt.addToMap(mergedLeaf)
      return mergedLeaf
  }

  func (mpt *MerklePatriciaTrie) mergeBranchAndExtension(currNode Node, childNode Node, branchNibble[]uint8) Node {
      childNibbles := Compact_decode(childNode.flag_value.encoded_prefix)
      newNibbles := mergeArrays(branchNibble, childNibbles)
      newExtension := createNode(2, [17]string{}, newNibbles, childNode.flag_value.value)
      delete(mpt.db, currNode.hash_node())
      delete(mpt.db, childNode.hash_node())
      mpt.addToMap(newExtension)
      return newExtension
  }

  func (mpt *MerklePatriciaTrie) checkBranch(branchCount int, currNode Node)  (string, Node, error) {

      if branchCount == 0 {
          return "", currNode, nil
      } else if branchCount == 1 {
          currNibble := []uint8{findNonEmptySpace(currNode)}
          if currNibble[0] == 16 {
              newLeaf := createNode(2, [17]string{}, []uint8{16}, currNode.branch_value[currNibble[0]])
              mpt.addToMap(newLeaf)
              return newLeaf.hash_node(), newLeaf, nil
          } else {
              childNode := mpt.db[currNode.branch_value[currNibble[0]]]
              switch childNode.node_type {
              case 1:
                  delete(mpt.db, currNode.hash_node())
                  newExtension := createNode(2, [17]string{}, currNibble, childNode.hash_node())
                  mpt.addToMap(newExtension)
                  return newExtension.hash_node(), newExtension, nil
              case 2:
                  if isLeaf(childNode) {
                      mergedLeaf := mpt.mergeBranchAndLeaf(currNode, childNode, currNibble)
                      return mergedLeaf.hash_node(), mergedLeaf, nil
                  } else {
                      mergedExtension := mpt.mergeBranchAndExtension(currNode, childNode, currNibble)
                      return mergedExtension.hash_node(), mergedExtension, nil
                  }
              }
          }
      }
      mpt.addToMap(currNode)
      return currNode.hash_node(), currNode, nil
  }

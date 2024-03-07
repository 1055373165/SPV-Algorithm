package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
)

var res [][][]byte

func main() {
	var root MerkleTree
	data := [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6"), []byte("7"), []byte("8")}
	root = *NewMerkleTree(data)
	fmt.Println("层序遍历结果(非递归):")
	fmt.Println()
	fmt.Println("_________________________________________")
	fmt.Println()
	travel1(root.RootNode)
	fmt.Println("层序遍历结果(递归):")
	fmt.Println()
	fmt.Println("_________________________________________")
	fmt.Println()
	travel_layer_recur(root.RootNode, 0)
	for i, layers := range res {
		fmt.Printf("第%d层的hash值为:\n", i)
		for _, val := range layers {
			fmt.Println(hex.EncodeToString(val))
		}
		fmt.Println()
		fmt.Println("_________________________________________")
		fmt.Println()
	}

	fmt.Println("获取Merkle Branch:")
	hash, _ := hex.DecodeString("4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a")
	Merkle_Branch := getMerkleBranch(hash)
	for i, v := range Merkle_Branch {
		fmt.Printf("Merkle Branch中的第%d个hash值:\n", i)
		fmt.Println(hex.EncodeToString(v))
	}
	fmt.Println()
	fmt.Println("_________________________________________")
	fmt.Println()
	fmt.Println("开始验证该交易是否存在于该区块内:")
	//参数为tx的hash值, tx所在深度, tx所位于的索引, 主节点发来的根hash值和Merkle Branch
	flag := Verify("4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a", 3, 3, "8f454ce466216a6b194e492727c49f68955bb174d2dc229b36cc3ed403099572", Merkle_Branch)
	fmt.Println(flag)
}

func travel1(root *MerkleNode) {
	q := []*MerkleNode{}
	if root != nil {
		q = append(q, root)
	}
	var node *MerkleNode
	layer := 0
	for len(q) > 0 {
		size := len(q)
		fmt.Printf("第%d层的hash值为:\n", layer)
		for size > 0 {
			node = q[0]
			q = q[1:]
			fmt.Println(hex.EncodeToString(node.Data))
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}
		layer++
		fmt.Println()
		fmt.Println("_________________________________________")
		fmt.Println()
	}
}

func travel_layer_recur(node *MerkleNode, depth int) {
	if node == nil {
		return
	}
	if len(res) == depth {
		res = append(res, [][]byte{})
	}
	res[depth] = append(res[depth], node.Data)
	travel_layer_recur(node.Left, depth+1)
	travel_layer_recur(node.Right, depth+1)
}

// MerkleTree represent a Merkle tree
// 默克尔树结构存储默克尔树根节点
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode represent a Merkle tree node
// 默克尔数节点结构 左右孩子、节点上存储的数据值，最底层叶子节点存储的是序列化交易的哈希
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// NewMerkleTree从一系列数据中创建新的Merkle树
func NewMerkleTree(data [][]byte) *MerkleTree {
	//默克尔树节点集
	var nodes []MerkleNode

	//构造默克尔树叶子节点集 填入交易信息 所有这些节点均为叶子节点
	for _, datum := range data {
		//叶子节点无左右孩子节点 最终存储的数据是序列化的交易的哈希值
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	//构造默克尔树
	for i := 0; i < len(data)/2-1; i++ {
		//每一层的默克尔节点集
		var newLevel []MerkleNode
		//如果当前层的nodes个数不是偶数个,就用最后一个节点填充
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		//一次处理两个节点 构造出一个上层节点
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		//更新为上一层默克尔节点集 不需要保存每一层的节点信息 我们只需要求MerkleTree的根节点就可以代表这颗merkle tree
		nodes = newLevel
	}

	//最终最上层的一个默克尔树节点即为默克尔树根节点
	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

// NewMerkleNode creates a new Merkle tree node
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}
	//如果是叶子节点 data存入该笔交易的hash值
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		//如果不是叶子节点 需要存入左右孩子节点交易级联的哈希值
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}
	//更新该节点的左右孩子节点
	mNode.Left = left
	mNode.Right = right

	return &mNode
}

func getMerkleBranch(hash []byte) [][]byte { //从full node那里获取getMerkleBranch
	ans := [][]byte{} //存储各层的内部hash组成Merkle Branch
	index := -1       // 用来记录匹配给定hash的位置，-1表示没找到
	/* 首先处理最后一层 */
	layer := res[len(res)-1]
	//找给定hash的位置
	for i, val := range layer {
		if reflect.DeepEqual(val, hash) {
			index = i
		}
	}
	//如果没找到，返回nil
	if index == -1 {
		return nil
	}
	//如果找到了，且索引位置能整除2，（比如index=2，其实是第三个节点，位于一对中的左侧）
	if index%2 == 0 { //位于左边,取右边节点的哈希
		ans = append(ans, layer[index+1])
	} else {
		ans = append(ans, layer[index-1])
	}

	//继续往上处理，获取其他层的内部hash值,因为上一层的个数是本层的一半，那么可以推得index处对应的上层父节点的下标为index/2.
	for i := len(res) - 2; i > 0; i-- {
		layer = res[i]
		index /= 2
		if index%2 == 0 {
			ans = append(ans, layer[index+1])
		} else {
			ans = append(ans, layer[index-1])
		}
	}
	return ans
}

// 参数为tx的hash值, tx所在深度, tx所位于的索引, 主节点发来的根hash值和Merkle Branch
func Verify(hash string, depth, index int, rootHash string, merkle_Branch [][]byte) bool {
	//求节点所在位置编号,从1开始编号
	number := int(math.Pow(float64(2), float64(depth))) - 1
	number += index + 1
	number_binary := intToBinary(number)
	hash_byte, _ := hex.DecodeString(hash)

	for i, hash_require := range merkle_Branch {
		bits := number_binary[0]
		number_binary = number_binary[1:]
		if bits == 0 {
			v := append(hash_byte, hash_require...)
			temp := sha256.Sum256(v)
			hash_byte = temp[:]
		} else {
			v := append(hash_require, hash_byte...)
			temp := sha256.Sum256(v)
			hash_byte = temp[:]
		}
		fmt.Printf("第%d次级联计算出的结果为:\n", i+1)
		fmt.Println(hex.EncodeToString(hash_byte))
		fmt.Println()
	}
	if hex.EncodeToString(hash_byte) == rootHash {
		return true
	} else {
		return false
	}
}

func intToBinary(n int) []int {
	ans := []int{}
	for ; n > 0; n = n / 2 {
		ans = append(ans, n%2)
	}
	return ans
}

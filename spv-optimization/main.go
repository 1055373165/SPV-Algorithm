package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"

	"github.com/charmbracelet/log"
)

var res [][][]byte

type MerkleProof struct {
	MerkleRoot []byte
	MerklePath [][]byte
}

func main() {
	var root MerkleTree
	data := [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6"), []byte("7"), []byte("8"), []byte("9"), []byte("10"), []byte("11"), []byte("12"), []byte("13"), []byte("14"), []byte("15"), []byte("16"), []byte("17"), []byte("18"), []byte("19"), []byte("20"), []byte("21"), []byte("22"), []byte("23"), []byte("24"), []byte("25"), []byte("26"), []byte("27"), []byte("28"), []byte("29"), []byte("30"), []byte("31"), []byte("32"), []byte("33"), []byte("34"), []byte("35"), []byte("36"), []byte("37"), []byte("38"), []byte("39"), []byte("40"), []byte("41"), []byte("42"), []byte("43"), []byte("44"), []byte("45"), []byte("46"), []byte("47"), []byte("48"), []byte("49"), []byte("50"), []byte("51"), []byte("52"), []byte("53"), []byte("54"), []byte("55"), []byte("56"), []byte("57"), []byte("58"), []byte("59"), []byte("60"), []byte("61"), []byte("62"), []byte("63"), []byte("64"), []byte("65"), []byte("66"), []byte("67"), []byte("68"), []byte("69"), []byte("70"), []byte("71"), []byte("72"), []byte("73"), []byte("74"), []byte("75"), []byte("76"), []byte("77"), []byte("78"), []byte("79"), []byte("80"), []byte("81"), []byte("82"), []byte("83"), []byte("84"), []byte("85"), []byte("86"), []byte("87"), []byte("88"), []byte("89"), []byte("90"), []byte("91"), []byte("92"), []byte("93"), []byte("94"), []byte("95"), []byte("96"), []byte("97"), []byte("98"), []byte("99"), []byte("100"), []byte("101"), []byte("102"), []byte("103"), []byte("104"), []byte("105"), []byte("106"), []byte("107"), []byte("108"), []byte("109"), []byte("110"), []byte("111"), []byte("112"), []byte("113"), []byte("114"), []byte("115"), []byte("116"), []byte("117"), []byte("118"), []byte("119"), []byte("120"), []byte("121"), []byte("122"), []byte("123"), []byte("124"), []byte("125"), []byte("126"), []byte("127"), []byte("128"), []byte("129"), []byte("130"), []byte("131"), []byte("132"), []byte("133"), []byte("134"), []byte("135"), []byte("136"), []byte("137"), []byte("138"), []byte("139"), []byte("140"), []byte("141"), []byte("142"), []byte("143"), []byte("144"), []byte("145"), []byte("146"), []byte("147"), []byte("148"), []byte("149"), []byte("150"), []byte("151"), []byte("152"), []byte("153"), []byte("154"), []byte("155"), []byte("156"), []byte("157"), []byte("158"), []byte("159"), []byte("160"), []byte("161"), []byte("162"), []byte("163"), []byte("164"), []byte("165"), []byte("166"), []byte("167"), []byte("168"), []byte("169"), []byte("170"), []byte("171"), []byte("172"), []byte("173"), []byte("174"), []byte("175"), []byte("176"), []byte("177"), []byte("178"), []byte("179"), []byte("180"), []byte("181"), []byte("182"), []byte("183"), []byte("184"), []byte("185"), []byte("186"), []byte("187"), []byte("188"), []byte("189"), []byte("190"), []byte("191"), []byte("192"), []byte("193"), []byte("194"), []byte("195"), []byte("196"), []byte("197"), []byte("198"), []byte("199"), []byte("200"), []byte("201"), []byte("202"), []byte("203"), []byte("204"), []byte("205"), []byte("206"), []byte("207"), []byte("208"), []byte("209"), []byte("210"), []byte("211"), []byte("212"), []byte("213"), []byte("214"), []byte("215"), []byte("216"), []byte("217"), []byte("218"), []byte("219"), []byte("220"), []byte("221"), []byte("222"), []byte("223"), []byte("224"), []byte("225"), []byte("226"), []byte("227"), []byte("228"), []byte("229"), []byte("230"), []byte("231"), []byte("232"), []byte("233"), []byte("234"), []byte("235"), []byte("236"), []byte("237"), []byte("238"), []byte("239"), []byte("240"), []byte("241"), []byte("242"), []byte("243"), []byte("244"), []byte("245"), []byte("246"), []byte("247"), []byte("248"), []byte("249"), []byte("250"), []byte("251"), []byte("252"), []byte("253"), []byte("254"), []byte("255"), []byte("256"), []byte("257"), []byte("258"), []byte("259"), []byte("260"), []byte("261"), []byte("262"), []byte("263"), []byte("264"), []byte("265"), []byte("266"), []byte("267"), []byte("268"), []byte("269"), []byte("270"), []byte("271"), []byte("272"), []byte("273"), []byte("274"), []byte("275"), []byte("276"), []byte("277"), []byte("278"), []byte("279"), []byte("280"), []byte("281"), []byte("282"), []byte("283"), []byte("284"), []byte("285"), []byte("286"), []byte("287"), []byte("288"), []byte("289"), []byte("290"), []byte("291"), []byte("292"), []byte("293"), []byte("294"), []byte("295"), []byte("296"), []byte("297"), []byte("298"), []byte("299"), []byte("300"), []byte("301"), []byte("302"), []byte("303"), []byte("304"), []byte("305"), []byte("306"), []byte("307"), []byte("308"), []byte("309"), []byte("310"), []byte("311"), []byte("312"), []byte("313"), []byte("314"), []byte("315"), []byte("316"), []byte("317"), []byte("318"), []byte("319"), []byte("320"), []byte("321"), []byte("322"), []byte("323"), []byte("324"), []byte("325"), []byte("326"), []byte("327"), []byte("328"), []byte("329"), []byte("330"), []byte("331"), []byte("332"), []byte("333"), []byte("334"), []byte("335"), []byte("336"), []byte("337"), []byte("338"), []byte("339"), []byte("340"), []byte("341"), []byte("342"), []byte("343"), []byte("344"), []byte("345"), []byte("346"), []byte("347"), []byte("348"), []byte("349"), []byte("350"), []byte("351"), []byte("352"), []byte("353"), []byte("354"), []byte("355"), []byte("356"), []byte("357"), []byte("358"), []byte("359"), []byte("360"), []byte("361"), []byte("362"), []byte("363"), []byte("364"), []byte("365"), []byte("366"), []byte("367"), []byte("368"), []byte("369"), []byte("370"), []byte("371"), []byte("372"), []byte("373"), []byte("374"), []byte("375"), []byte("376"), []byte("377"), []byte("378"), []byte("379"), []byte("380"), []byte("381"), []byte("382"), []byte("383"), []byte("384"), []byte("385"), []byte("386"), []byte("387"), []byte("388"), []byte("389"), []byte("390"), []byte("391"), []byte("392"), []byte("393"), []byte("394"), []byte("395"), []byte("396"), []byte("397"), []byte("398"), []byte("399"), []byte("400"), []byte("401"), []byte("402"), []byte("403"), []byte("404"), []byte("405"), []byte("406"), []byte("407"), []byte("408"), []byte("409"), []byte("410"), []byte("411"), []byte("412"), []byte("413"), []byte("414"), []byte("415"), []byte("416"), []byte("417"), []byte("418"), []byte("419"), []byte("420"), []byte("421"), []byte("422"), []byte("423"), []byte("424"), []byte("425"), []byte("426"), []byte("427"), []byte("428"), []byte("429"), []byte("430"), []byte("431"), []byte("432"), []byte("433"), []byte("434"), []byte("435"), []byte("436"), []byte("437"), []byte("438"), []byte("439"), []byte("440"), []byte("441"), []byte("442"), []byte("443"), []byte("444"), []byte("445"), []byte("446"), []byte("447"), []byte("448"), []byte("449"), []byte("450"), []byte("451"), []byte("452"), []byte("453"), []byte("454"), []byte("455"), []byte("456"), []byte("457"), []byte("458"), []byte("459"), []byte("460"), []byte("461"), []byte("462"), []byte("463"), []byte("464"), []byte("465"), []byte("466"), []byte("467"), []byte("468"), []byte("469"), []byte("470"), []byte("471"), []byte("472"), []byte("473"), []byte("474"), []byte("475"), []byte("476"), []byte("477"), []byte("478"), []byte("479"), []byte("480"), []byte("481"), []byte("482"), []byte("483"), []byte("484"), []byte("485"), []byte("486"), []byte("487"), []byte("488"), []byte("489"), []byte("490"), []byte("491"), []byte("492"), []byte("493"), []byte("494"), []byte("495"), []byte("496"), []byte("497"), []byte("498"), []byte("499"), []byte("500"), []byte("501"), []byte("502"), []byte("503"), []byte("504"), []byte("505"), []byte("506"), []byte("507"), []byte("508"), []byte("509"), []byte("510"), []byte("511"), []byte("512"), []byte("513"), []byte("514"), []byte("515"), []byte("516"), []byte("517"), []byte("518"), []byte("519"), []byte("520"), []byte("521"), []byte("522"), []byte("523"), []byte("524"), []byte("525"), []byte("526"), []byte("527"), []byte("528"), []byte("529"), []byte("530"), []byte("531"), []byte("532"), []byte("533"), []byte("534"), []byte("535"), []byte("536"), []byte("537"), []byte("538"), []byte("539"), []byte("540"), []byte("541"), []byte("542"), []byte("543"), []byte("544"), []byte("545"), []byte("546"), []byte("547"), []byte("548"), []byte("549"), []byte("550"), []byte("551"), []byte("552"), []byte("553"), []byte("554"), []byte("555"), []byte("556"), []byte("557"), []byte("558"), []byte("559"), []byte("560"), []byte("561"), []byte("562"), []byte("563"), []byte("564"), []byte("565"), []byte("566"), []byte("567"), []byte("568"), []byte("569"), []byte("570"), []byte("571"), []byte("572"), []byte("573"), []byte("574"), []byte("575"), []byte("576"), []byte("577"), []byte("578"), []byte("579"), []byte("580"), []byte("581"), []byte("582"), []byte("583"), []byte("584"), []byte("585"), []byte("586"), []byte("587"), []byte("588"), []byte("589"), []byte("590"), []byte("591"), []byte("592"), []byte("593"), []byte("594"), []byte("595"), []byte("596"), []byte("597"), []byte("598"), []byte("599"), []byte("600"), []byte("601"), []byte("602"), []byte("603"), []byte("604"), []byte("605"), []byte("606"), []byte("607"), []byte("608"), []byte("609"), []byte("610"), []byte("611"), []byte("612"), []byte("613"), []byte("614"), []byte("615"), []byte("616"), []byte("617"), []byte("618"), []byte("619"), []byte("620"), []byte("621"), []byte("622"), []byte("623"), []byte("624"), []byte("625"), []byte("626"), []byte("627"), []byte("628"), []byte("629"), []byte("630"), []byte("631"), []byte("632"), []byte("633"), []byte("634"), []byte("635"), []byte("636"), []byte("637"), []byte("638"), []byte("639"), []byte("640"), []byte("641"), []byte("642"), []byte("643"), []byte("644"), []byte("645"), []byte("646"), []byte("647"), []byte("648"), []byte("649"), []byte("650"), []byte("651"), []byte("652"), []byte("653"), []byte("654"), []byte("655"), []byte("656"), []byte("657"), []byte("658"), []byte("659"), []byte("660"), []byte("661"), []byte("662"), []byte("663"), []byte("664"), []byte("665"), []byte("666"), []byte("667"), []byte("668"), []byte("669"), []byte("670"), []byte("671"), []byte("672"), []byte("673"), []byte("674"), []byte("675"), []byte("676"), []byte("677"), []byte("678"), []byte("679"), []byte("680"), []byte("681"), []byte("682"), []byte("683"), []byte("684"), []byte("685"), []byte("686"), []byte("687"), []byte("688"), []byte("689"), []byte("690"), []byte("691"), []byte("692"), []byte("693"), []byte("694"), []byte("695"), []byte("696"), []byte("697"), []byte("698"), []byte("699"), []byte("700"), []byte("701"), []byte("702"), []byte("703"), []byte("704"), []byte("705"), []byte("706"), []byte("707"), []byte("708"), []byte("709"), []byte("710"), []byte("711"), []byte("712"), []byte("713"), []byte("714"), []byte("715"), []byte("716"), []byte("717"), []byte("718"), []byte("719"), []byte("720"), []byte("721"), []byte("722"), []byte("723"), []byte("724"), []byte("725"), []byte("726"), []byte("727"), []byte("728"), []byte("729"), []byte("730"), []byte("731"), []byte("732"), []byte("733"), []byte("734"), []byte("735"), []byte("736"), []byte("737"), []byte("738"), []byte("739"), []byte("740"), []byte("741"), []byte("742"), []byte("743"), []byte("744"), []byte("745"), []byte("746"), []byte("747"), []byte("748"), []byte("749"), []byte("750"), []byte("751"), []byte("752"), []byte("753"), []byte("754"), []byte("755"), []byte("756"), []byte("757"), []byte("758"), []byte("759"), []byte("760"), []byte("761"), []byte("762"), []byte("763"), []byte("764"), []byte("765"), []byte("766"), []byte("767"), []byte("768"), []byte("769"), []byte("770"), []byte("771"), []byte("772"), []byte("773"), []byte("774"), []byte("775"), []byte("776"), []byte("777"), []byte("778"), []byte("779"), []byte("780"), []byte("781"), []byte("782"), []byte("783"), []byte("784"), []byte("785"), []byte("786"), []byte("787"), []byte("788"), []byte("789"), []byte("790"), []byte("791"), []byte("792"), []byte("793"), []byte("794"), []byte("795"), []byte("796"), []byte("797"), []byte("798"), []byte("799"), []byte("800"), []byte("801"), []byte("802"), []byte("803"), []byte("804"), []byte("805"), []byte("806"), []byte("807"), []byte("808"), []byte("809"), []byte("810"), []byte("811"), []byte("812"), []byte("813"), []byte("814"), []byte("815"), []byte("816"), []byte("817"), []byte("818"), []byte("819"), []byte("820"), []byte("821"), []byte("822"), []byte("823"), []byte("824"), []byte("825"), []byte("826"), []byte("827"), []byte("828"), []byte("829"), []byte("830"), []byte("831"), []byte("832"), []byte("833"), []byte("834"), []byte("835"), []byte("836"), []byte("837"), []byte("838"), []byte("839"), []byte("840"), []byte("841"), []byte("842"), []byte("843"), []byte("844"), []byte("845"), []byte("846"), []byte("847"), []byte("848"), []byte("849"), []byte("850"), []byte("851"), []byte("852"), []byte("853"), []byte("854"), []byte("855"), []byte("856"), []byte("857"), []byte("858"), []byte("859"), []byte("860"), []byte("861"), []byte("862"), []byte("863"), []byte("864"), []byte("865"), []byte("866"), []byte("867"), []byte("868"), []byte("869"), []byte("870"), []byte("871"), []byte("872"), []byte("873"), []byte("874"), []byte("875"), []byte("876"), []byte("877"), []byte("878"), []byte("879"), []byte("880"), []byte("881"), []byte("882"), []byte("883"), []byte("884"), []byte("885"), []byte("886"), []byte("887"), []byte("888"), []byte("889"), []byte("890"), []byte("891"), []byte("892"), []byte("893"), []byte("894"), []byte("895"), []byte("896"), []byte("897"), []byte("898"), []byte("899"), []byte("900"), []byte("901"), []byte("902"), []byte("903"), []byte("904"), []byte("905"), []byte("906"), []byte("907"), []byte("908"), []byte("909"), []byte("910"), []byte("911"), []byte("912"), []byte("913"), []byte("914"), []byte("915"), []byte("916"), []byte("917"), []byte("918"), []byte("919"), []byte("920"), []byte("921"), []byte("922"), []byte("923"), []byte("924"), []byte("925"), []byte("926"), []byte("927"), []byte("928"), []byte("929"), []byte("930"), []byte("931"), []byte("932"), []byte("933"), []byte("934"), []byte("935"), []byte("936"), []byte("937"), []byte("938"), []byte("939"), []byte("940"), []byte("941"), []byte("942"), []byte("943"), []byte("944"), []byte("945"), []byte("946"), []byte("947"), []byte("948"), []byte("949"), []byte("950"), []byte("951"), []byte("952"), []byte("953"), []byte("954"), []byte("955"), []byte("956"), []byte("957"), []byte("958"), []byte("959"), []byte("960"), []byte("961"), []byte("962"), []byte("963"), []byte("964"), []byte("965"), []byte("966"), []byte("967"), []byte("968"), []byte("969"), []byte("970"), []byte("971"), []byte("972"), []byte("973"), []byte("974"), []byte("975"), []byte("976"), []byte("977"), []byte("978"), []byte("979"), []byte("980"), []byte("981"), []byte("982"), []byte("983"), []byte("984"), []byte("985"), []byte("986"), []byte("987"), []byte("988"), []byte("989"), []byte("990"), []byte("991"), []byte("992"), []byte("993"), []byte("994"), []byte("995"), []byte("996"), []byte("997"), []byte("998"), []byte("999")}
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
		if i == 0 || i == 1 {
			fmt.Printf("第%d层的hash值为:\n", i)
			for _, val := range layers {
				fmt.Println(hex.EncodeToString(val))
			}
		}
	}

	fmt.Println()
	fmt.Println("_________________________________________")
	fmt.Println()
	txs := computeHash(data)
	fmt.Println(hex.EncodeToString(getMerkleRoot(txs)))

	fmt.Println("使用优化算法获取 Merkle Proof:")
	hash, _ := hex.DecodeString("2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3")
	merkleProof, err := getMerkleProof_New(txs, hash, 7)
	if err != nil {
		log.Errorf("get merkler proof(new) failed, err = %v", err)
		return
	}
	fmt.Println("MerkleRoot: ", hex.EncodeToString(merkleProof.MerkleRoot))
	fmt.Println("MerkleBranch: ")
	merklesibling := merkleProof.MerklePath
	for i, v := range merklesibling {
		fmt.Println(i, ":", hex.EncodeToString(v))
	}
	fmt.Println()
	fmt.Println("_________________________________________")
	fmt.Println()
	fmt.Println("使用 BTCRelay 算法获取 MerkleProof:")
	merkleProof = getMerkleProof_BTCRelay(txs, 7)
	fmt.Println("MerkleRoot: ", hex.EncodeToString(merkleProof.MerkleRoot))
	fmt.Println("MerkleBranch: ")
	merklesibling = merkleProof.MerklePath
	for i, v := range merklesibling {
		fmt.Println(i, ":", hex.EncodeToString(v))
	}
	fmt.Println()
	fmt.Println("_________________________________________")
	fmt.Println()
	fmt.Println("开始验证该交易是否存在于该区块内:")
	//参数为tx的hash值, tx所在深度, tx所位于的索引, 主节点发来的根hash值和Merkle Branch
	flag := Verify_New(hash, 7, merkleProof.MerkleRoot, merkleProof.MerklePath)
	fmt.Println(flag)

	flag = Verify_BTCRelay(hash, 7, merkleProof.MerkleRoot, merkleProof.MerklePath)
	fmt.Println(flag)

	getMerkleRoot(txs)
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
	for len(nodes) != 1 {
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

func getMerkleProof_BTCRelay(data [][]byte, index int) MerkleProof {
	target := data[index]
	merkle_sibling := [][]byte{}

	for len(data) != 1 {
		newTree := make([][]byte, (len(data)+1)/2)
		for j := 0; j < len(data); j += 2 {
			hash1, hash2 := data[j], data[int(math.Min(float64(j+1), float64(len(data)-1)))]
			temp := sha256.Sum256(append(hash1, hash2...))
			newTree[j/2] = temp[:]

			if reflect.DeepEqual(target, hash1) {
				merkle_sibling = append(merkle_sibling, hash2)
				target = newTree[j/2]
			} else if reflect.DeepEqual(target, hash2) {
				merkle_sibling = append(merkle_sibling, hash1)
				target = newTree[j/2]
			}
		}

		data = newTree
	}

	return MerkleProof{data[0], merkle_sibling}
}

func getMerkleProof_New(data [][]byte, hash []byte, index int) (MerkleProof, error) { //从full node那里获取getMerkleBranch
	// 1. 先比较是否相等，不相等直接退出
	if !reflect.DeepEqual(data[index], hash) {
		return MerkleProof{}, errors.New("DeepEqual 验证失败")
	}

	merkle_sibling := [][]byte{}

	for len(data) != 1 {
		newTree := make([][]byte, (len(data)+1)/2)
		for j := 0; j < len(data); j += 2 {
			hash1, hash2 := data[j], data[int(math.Min(float64(j+1), float64(len(data)-1)))]
			temp := sha256.Sum256(append(hash1, hash2...))
			newTree[j/2] = temp[:]
			// 2. 无需一个个比较，直接判断是否已经计算到 index 处即可
			if j == index {
				merkle_sibling = append(merkle_sibling, hash2)
			} else if j+1 == index {
				merkle_sibling = append(merkle_sibling, hash1)
			}
		}
		data = newTree
		index /= 2
	}

	return MerkleProof{data[0], merkle_sibling}, nil
}

func Verify_BTCRelay(hash []byte, index int, MerkleRoot []byte, merkle_sibling [][]byte) bool {
	for i := 0; i < len(merkle_sibling); i++ {
		if index%2 == 0 {
			temp := sha256.Sum256(append(hash, merkle_sibling[i]...))
			hash = temp[:]
			index /= 2
		} else {
			temp := sha256.Sum256(append(merkle_sibling[i], hash...))
			hash = temp[:]
			index /= 2
		}
	}

	return reflect.DeepEqual(hash, MerkleRoot)
}

// 参数为tx的hash值, tx所在深度, tx所位于的索引, 主节点发来的根hash值和Merkle Branch
func Verify_New(hash []byte, index int, MerkleRoot []byte, merkle_sibling [][]byte) bool {
	// 求节点所在位置编号,从1开始编号
	depth := len(merkle_sibling)
	number := int(math.Pow(float64(2), float64(depth))) + index
	number_binary := intToBinary(number)

	for i, another := range merkle_sibling {
		bits := number_binary[i]
		if bits == 0 {
			v := append(hash, another...)
			temp := sha256.Sum256(v)
			hash = temp[:]
		} else {
			v := append(another, hash...)
			temp := sha256.Sum256(v)
			hash = temp[:]
		}
	}

	return reflect.DeepEqual(hash, MerkleRoot)
}

func intToBinary(n int) []int {
	ans := []int{}
	for ; n > 0; n = n / 2 {
		ans = append(ans, n%2)
	}
	return ans
}

func computeHash(data [][]byte) [][]byte {
	ans := [][]byte{}
	for i := 0; i < len(data); i++ {
		temp := sha256.Sum256(data[i])
		ans = append(ans, temp[:])
	}

	return ans
}

func getMerkleRoot(data [][]byte) []byte {
	for len(data) != 1 {
		newTree := make([][]byte, (len(data)+1)/2)
		for j := 0; j < len(data); j += 2 {
			hash1, hash2 := data[j], data[int(math.Min(float64(j+1), float64(len(data)-1)))]
			temp := sha256.Sum256(append(hash1, hash2...))
			newTree[j/2] = temp[:]
		}

		data = newTree
	}
	return data[0]
}

func generateDataSet(txNums int) [][]byte {
	// read data template
	file, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Errorf("read data from file failed, err: %v", err)
		return [][]byte{}
	}
	var dataSet [][]byte
	// construct data set (2000 tx)
	// data1 -> data2000
	for i := 0; i < txNums; i++ {
		dataSet = append(dataSet, []byte(string(data)+strconv.Itoa(i+1)))
	}
	return dataSet
}

func getTxHash(index int) string {
	file, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	buf.Write(data)
	buf.WriteString(strconv.Itoa(index + 1))
	temp := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(temp[:])
}

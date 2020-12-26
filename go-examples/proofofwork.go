package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"time"
)

// 0000 0000 0000 0000 1001 0001 0000  .... 0001
//256位Hash里面前面至少有16个零
const TargetBit = 16 // 20

//step2：创建pow结构体
type ProofOfWork struct {
	//要验证的区块
	Block *Block

	//大整数存储,目标哈希
	Target *big.Int
}

//step3: 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1.创建一个big对象 0000000.....00001
	/*
	   0000 0001
	   0010 0000
	*/
	target := big.NewInt(1)

	//2.左移256-bits位
	target = target.Lsh(target, 256-TargetBit)

	return &ProofOfWork{block, target}
}

//step4：返回有效的哈希和nonce值
func (pow *ProofOfWork) Run() ([]byte, int64) {
	//1.将Block的属性拼接成字节数组
	//2.生成Hash
	//3.循环判断Hash的有效性，满足条件，跳出循环结束验证
	nonce := 0
	//var hashInt big.Int //用于存储新生成的hash
	hashInt := new(big.Int)
	var hash [32]byte
	for {
		//获取字节数组
		dataBytes := pow.prepareData(nonce)
		//生成hash
		hash = sha256.Sum256(dataBytes)
		//fmt.Printf("%d: %x\n",nonce,hash)
		fmt.Printf("\r%d: %x", nonce, hash)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Block里的target
		/*
		   Com compares x and y and returns:
		   -1 if x < y
		   0 if x == y
		   1 if x > y
		*/
		if pow.Target.Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	return hash[:], int64(nonce)
}

//step5：根据block生成一个byte数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(int64(TargetBit)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) IsValid() bool {
	hashInt := new(big.Int)
	hashInt.SetBytes(pow.Block.Hash)
	return pow.Target.Cmp(hashInt) == 1
}

/******************************* Block *******************************/
//step1:创建Block结构体
type Block struct {
	//字段：
	//高度Height：其实就是区块的编号，第一个区块叫创世区块，高度为0
	Height int64
	//上一个区块的哈希值ProvHash：
	PrevBlockHash []byte
	//交易数据Data：目前先设计为[]byte,后期是Transaction
	Data []byte
	//时间戳TimeStamp：
	TimeStamp int64
	//哈希值Hash：32个的字节，64个16进制数
	Hash []byte

	Nonce int64
}

//step2：创建新的区块
func NewBlock(data string, provBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, provBlockHash, []byte(data), time.Now().Unix(), nil, 0}
	//step5：设置block的hash和nonce
	//设置哈希
	//block.SetHash()
	//调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

//step4:创建创世区块：
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}

/******************************* Utils *******************************/
/*
将一个int64的整数：转为二进制后，每8bit一个byte。转为[]byte
*/
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	//将二进制数据写入w
	//func Write(w io.Writer, order ByteOrder, data interface{}) error
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	//转为[]byte并返回
	return buff.Bytes()
}

func main() {
	//1.测试Block
	//block:=BLC.NewBlock("I am a block",make([]byte,32,32),1)
	//fmt.Println(block)
	//2.测试创世区块
	//genesisBlock :=BLC.CreateGenesisBlock("Genesis Block..")
	//fmt.Println(genesisBlock)

	//3.测试区块链
	//genesisBlockChain := BLC.CreateBlockChainWithGenesisBlock()
	//fmt.Println(genesisBlockChain)
	//fmt.Println(genesisBlockChain.Blocks)
	//fmt.Println(genesisBlockChain.Blocks[0])

	//4.测试添加新区块
	blockChain := CreateGenesisBlock("Genesis Block..")
	fmt.Println(blockChain)

	pow := NewProofOfWork(blockChain)
	fmt.Printf("%v\n", pow.IsValid())

	/*
	   // 5.检测pow
	   //1.创建一个big对象 0000000.....00001
	   target := big.NewInt(1)
	   fmt.Printf("0x%x\n",target) //0x1

	   //2.左移256-bits位
	   target = target.Lsh(target, 256-BLC.TargetBit)

	   fmt.Printf("0x%x\n",target) //61
	   //61位：0x1000000000000000000000000000000000000000000000000000000000000
	   //64位：0x0001000000000000000000000000000000000000000000000000000000000000

	   s1:="HelloWorld"
	   hash:=sha256.Sum256([]byte(s1))
	   fmt.Printf("0x%x\n",hash)
	*/
}

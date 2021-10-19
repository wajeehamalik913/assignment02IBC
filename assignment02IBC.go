package assignment02IBC

import (
	"crypto/sha256"
	"fmt"
)

const miningReward = 100
const rootUser = "Satoshi"
var aliceBal = 0
var bobBal = 0
var satoshiBal = 0
type BlockData struct {
	Title    string
	Sender   string
	Receiver string
	Amount   int
}
type Block struct {
	Data        []BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {
	//var recieved int = 0
	//var spent int = 0
	var bal int =0
	for chainHead != nil {
		for i, _ := range chainHead.Data {
			if chainHead.Data[i].Receiver == userName {
				bal = bal + chainHead.Data[i].Amount
			}
			if chainHead.Data[i].Sender == userName {
				bal = bal - chainHead.Data[i].Amount
			}
		}
		chainHead = chainHead.PrevPointer
	}
	//bal := recieved - spent
	return bal
}
func CalculateHash(inputBlock *Block) string {
	var c1 string
	//fmt.Println("calc hash entered")
	if inputBlock.PrevPointer == nil {
		v := fmt.Sprintf("%v", inputBlock.Data)
		c1 = fmt.Sprintf("%x", sha256.Sum256([]byte(v)))
		return c1
	} else {
		n := fmt.Sprintf("%v", inputBlock.Data)
		str := n + inputBlock.PrevHash
		//v:=fmt.Sprintf("%v",str)
		//fmt.Printf(v,"hiiiii\n")
		c1 = fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
		//fmt.Printf(c1)
		return c1
	}
}
func VerifyTransaction(transaction *BlockData, chainHead *Block) bool {
	//fmt.Print(transaction.Sender,"****BALANCE****",CalculateBalance(transaction.Sender, chainHead),"*************")
	//fmt.Println()
	if CalculateBalance(transaction.Sender, chainHead) > transaction.Amount {
		return true
	} else {
		return false
	}
}
func InsertBlock(blockData []BlockData, chainHead *Block) *Block {
	var verify bool = true
	var data *BlockData
	var temp *Block
	//var temp2 *Block
	//temp2=chainHead
	temp=chainHead
	if temp == nil {
		//fmt.Printf(newBlock.CurrentHash,"\n")
		for i, _ := range blockData {
			//var newBlock Block
			/*c:=blockData
			x:=append(c, blockData[i])
			newBlock.Data = x
			newBlock.CurrentHash = CalculateHash(&newBlock)
			fmt.Print("*************")
			fmt.Print(newBlock.Data)
			fmt.Print("*************")*/
			data = &blockData[i]
			if VerifyTransaction(data, chainHead) == false {
				fmt.Println()
				fmt.Printf("transaction not valid\n")
				fmt.Println()
				return chainHead
			}
		}
	} else {
		for temp != nil {

			//fmt.Printf(newBlock.CurrentHash)
			for i, _ := range blockData {
				data = &blockData[i]
				if VerifyTransaction(data, chainHead) == false {
					fmt.Println()
					fmt.Printf("transaction not valid\n")
					fmt.Println()
					return chainHead
				}
			}
			temp = temp.PrevPointer
		}
	}
	if verify == true {
		if chainHead == nil {
			var newBlock Block
			blockData=append(blockData, []BlockData{{Title: "Coinbase", Sender: "system", Receiver: rootUser, Amount: 100}}...)
			newBlock.Data = blockData
			newBlock.CurrentHash = CalculateHash(&newBlock)
			//fmt.Printf(newBlock.CurrentHash,"\n")
			chainHead = &newBlock
			for i, _ := range blockData {
				data = &blockData[i]
				if CalculateBalance(data.Sender, chainHead) <=0 {
					fmt.Println()
					fmt.Printf("transaction not valid\n")
					fmt.Println()
					chainHead=chainHead.PrevPointer
					return chainHead
				}
			return chainHead

		}
		}else {
			var newBlock Block
			blockData=append(blockData, []BlockData{{Title: "Coinbase", Sender: "system", Receiver: rootUser, Amount: 100}}...)
			newBlock.Data = blockData
			newBlock.PrevPointer = chainHead
			newBlock.PrevHash = chainHead.CurrentHash
			newBlock.CurrentHash = CalculateHash(&newBlock)
			//fmt.Printf(newBlock.CurrentHash)
			chainHead = &newBlock
			for i, _ := range blockData {
				data = &blockData[i]
				if CalculateBalance(data.Sender, chainHead) <=0 {
					fmt.Println()
					fmt.Printf("transaction not valid\n")
					fmt.Println()
					chainHead=chainHead.PrevPointer
					return chainHead
			}
			return chainHead
		}
	}
}
	return chainHead
}
func ListBlocks(chainHead *Block) {
		fmt.Println()
	currentAddress := chainHead
	for {
		fmt.Print(currentAddress.Data)
		if currentAddress.PrevPointer != nil {
			fmt.Print(" --> ")
			currentAddress = currentAddress.PrevPointer
		} else {
			fmt.Println()
			break
		}

	}
	fmt.Println()
/*	for chainHead != nil {
		for i, _ := range chainHead.Data {
			fmt.Printf("%v", chainHead.Data[i])
			fmt.Printf("<-")
		}

		chainHead = chainHead.PrevPointer
	}
	fmt.Printf("\n")*/
}
func VerifyChain(chainHead *Block) {

}
func PremineChain(chainHead *Block, numBlocks int) *Block {
	i := 0
	if chainHead == nil {
		for i < numBlocks {
			transactions := []BlockData{{Title: "Premined", Sender: "nil", Receiver: "nil", Amount: 0}, {Title: "Premined", Sender: "System", Receiver: "Satoshi", Amount: 100}}
			var newBlock Block
			newBlock.Data = transactions
			newBlock.PrevPointer = chainHead
			newBlock.CurrentHash = CalculateHash(&newBlock)
			//fmt.Printf(newBlock.CurrentHash)
			chainHead = &newBlock
			i++
		}
	}
	return chainHead
}

package main

import (
	hdWallet "github.com/miguelmota/go-ethereum-hdwallet"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type EthBalance struct {
	gorm.Model
	Address    string `gorm:"size:255;uniqueIndex"`
	PublicKey  string `gorm:"size:255"`
	PrivateKey string `gorm:"size:255"`
	Mnemonic   string `gorm:"size:255"`
	Type       string `gorm:"size:255;index"`
}

var err error
var db *gorm.DB
var dsn = "root:123456@tcp(127.0.0.1:3306)/wallet?charset=utf8mb4&parseTime=True&loc=Local"
var wg sync.WaitGroup

func main() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&EthBalance{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().CreateTable(&EthBalance{})
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go InsertBtcBalance128()
		go InsertBtcBalance256()
	}
	wg.Wait()
}

func InsertBtcBalance128() {
	defer wg.Done()

	for i := 0; i < 1000000; i++ {
		btcBalance, err := createMnemonic(128, "english-128")
		if err != nil {
			log.Fatal(err)
		}

		err = db.Create(btcBalance).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InsertBtcBalance256() {
	defer wg.Done()

	for i := 0; i < 1000000; i++ {
		btcBalance, err := createMnemonic(256, "english-256")
		if err != nil {
			log.Fatal(err)
		}

		err = db.Create(btcBalance).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createMnemonic(bitSize int, addressType string) (*EthBalance, error) {
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	wallet, err := hdWallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdWallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return nil, err
	}

	publicKey, err := wallet.PublicKeyHex(account)
	if err != nil {
		return nil, err
	}

	return &EthBalance{
		Address:    account.Address.Hex(),
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Mnemonic:   mnemonic,
		Type:       addressType,
	}, nil
}

func DingDingNotice() {
	url := "https://oapi.dingtalk.com/robot/send?access_token=b8dd74418dc292f6de7b4c1426dae4a3ec483abdbbff77deb8b77d7151504eab"
	method := "POST"

	payload := strings.NewReader(`{"msgtype": "text","text": {"content": "我就是我, 是不一样的烟火"}}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	//{"errcode":0,"errmsg":"ok"}
}

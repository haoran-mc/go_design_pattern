package facade

import (
	"fmt"
	"log"
	"testing"
)

func TestFacade(t *testing.T) {
	walletFacade := newWalletFacade("abc", 1234)          // 创建一个钱包
	err := walletFacade.addMoneyToWallet("abc", 1234, 10) // 往钱包里存入 10 元
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()

	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5) // 花费 5 元
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

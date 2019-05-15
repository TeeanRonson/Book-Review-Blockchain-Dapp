package nodeData

import (
    "fmt"
)

type Wallet struct {
    transactions []float32
    fees float32
}

/**
Return a new Wallet for the Miner
 */
func NewWallet() Wallet{
    txList := make([]float32, 0)
    return Wallet{txList, 0}
}

/**
Add new fees for the miner
 */
func (wallet *Wallet) AddFee(fee float32) {
    wallet.transactions = append(wallet.transactions, fee)
    wallet.fees += fee
}

/**
Show the total fees accumulated
 */
 func (wallet *Wallet) ShowWallet() string {

     myTransactions := ""
     for i, fee := range wallet.transactions {
         s := fmt.Sprintf("%f", fee)
         myTransactions += "Transaction " + string(i) + ": " + s + "\n"
     }
     s := fmt.Sprintf("%f", wallet.fees)
     myTransactions += "Total fees collected: " + s
    return myTransactions
 }
package main

import (
	"avs-operator-go/bindings"
	"avs-operator-go/config"
	"context"
	"fmt"
	"log"
	"time"

	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// generateRandomName creates a random name by combining a random adjective, a random noun, and a random number.
func generateRandomName() string {
	adjectives := []string{"Quick", "Lazy", "Sleepy", "Noisy", "Hungry"}
	nouns := []string{"Fox", "Dog", "Cat", "Mouse", "Bear"}

	// Seed the random number generator to ensure different results on each run
	rand.Seed(time.Now().UnixNano())

	// Select random adjective and noun
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	// Generate a random number between 0 and 999
	number := rand.Intn(1000)

	// Create and return the random name
	return fmt.Sprintf("%s%s%d", adjective, noun, number)
}

func createNewTask(taskName string) {
	cfg := config.LoadConfig()

	privKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		log.Fatalf("failed to connect to the Ethereum network: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chain ID: %v", err)
	}

	contract, err := bindings.NewContract(common.HexToAddress(cfg.ContractAddress), client)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatalf("failed to create transactor: %v", err)
	}

	tx, _ := contract.CreateNewTask(opts, taskName)
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil || receipt.Status != 1 {
		log.Fatalf("Error sending transaction: %v", err)
	}
	log.Println("Transaction successful with hash:", tx.Hash().Hex())
}

func startCreatingTasks() {
	for {
		taskName := generateRandomName()
		createNewTask(taskName)
		time.Sleep(15 * time.Second)
	}
}

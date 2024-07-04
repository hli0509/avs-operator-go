package operator

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"avs-operator-go/bindings"
	"avs-operator-go/config"
)

type Operator struct {
	delegationManager *bindings.Delegation
	avsDirectory      *bindings.AvsDirectoryCaller
	registry          *bindings.Registry
	contract          *bindings.Contract
	chainID           *big.Int
	client            *ethclient.Client

	contractAddress common.Address
	privKey         *ecdsa.PrivateKey
}

func NewOperator(cfg config.Config) *Operator {
	privKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		log.Fatalf("failed to connect to rpc endpoint: %v", err)
	}
	if err != nil {
		log.Fatalf("failed to connect to ws endpoint: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chain ID: %v", err)
	}

	delegationManager, err := bindings.NewDelegation(common.HexToAddress(cfg.DelegationManagerAddress), client)
	if err != nil {
		log.Fatalf("failed to instantiate delegation manager contract: %v", err)
	}

	avsDirectory, err := bindings.NewAvsDirectoryCaller(common.HexToAddress(cfg.AvsDirectoryAddress), client)
	if err != nil {
		log.Fatalf("failed to instantiate AVS directory contract: %v", err)
	}

	registry, err := bindings.NewRegistry(common.HexToAddress(cfg.StakeRegistryAddress), client)
	if err != nil {
		log.Fatalf("failed to instantiate registry contract: %v", err)
	}

	contract, err := bindings.NewContract(common.HexToAddress(cfg.ContractAddress), client)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}

	return &Operator{
		delegationManager: delegationManager,
		avsDirectory:      avsDirectory,
		registry:          registry,
		contract:          contract,
		chainID:           chainID,
		client:            client,
		contractAddress:   common.HexToAddress(cfg.ContractAddress),
		privKey:           privKey,
	}

}

func (o *Operator) RegisterOperator() error {
	log.Println("check")
	address := crypto.PubkeyToAddress(o.privKey.PublicKey)
	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, o.chainID)
	if err != nil {
		return err
	}
	tx1, err := o.delegationManager.RegisterAsOperator(
		opts,
		bindings.IDelegationManagerOperatorDetails{
			EarningsReceiver:         address, // public address
			DelegationApprover:       common.HexToAddress("0x0000000000000000000000000000000000000000"),
			StakerOptOutWindowBlocks: 0,
		}, "",
	)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), o.client, tx1)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return fmt.Errorf("tx failed: %v", receipt.Status)
	}
	log.Println("Operator registered on EL successfully")

	salt, err := generateRandomBytes()
	if err != nil {
		return err
	}
	expiry := big.NewInt(time.Now().Add(time.Hour).Unix()) // 1 hour from now

	digestHash, err := o.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		nil,
		address,
		o.contractAddress,
		salt,
		expiry,
	)
	if err != nil {
		return err
	}
	// sign the hash
	sig, err := crypto.Sign(digestHash[:], o.privKey)
	if err != nil {
		return err
	}

	tx2, err := o.registry.RegisterOperatorWithSignature(
		opts,
		address,
		bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
			Salt:      salt,
			Expiry:    expiry,
			Signature: sig,
		},
	)
	if err != nil {
		return err
	}
	receipt, err = bind.WaitMined(context.Background(), o.client, tx2)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return fmt.Errorf("tx failed: %v", receipt.Status)
	}
	log.Println("Operator registered on AVS successfully")
	return nil
}

func (o *Operator) SignAndRespondToTask(taskIndex, taskCreatedBlock uint32, taskName string) error {
	message := fmt.Sprintf("Hello, %s", taskName)
	// keccak256 hash of the task name
	messageHash := crypto.Keccak256([]byte(message))
	sig, err := crypto.Sign(messageHash, o.privKey)
	if err != nil {
		return err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, o.chainID)
	if err != nil {
		return err
	}
	tx, err := o.contract.RespondToTask(
		opts,
		bindings.IHelloWorldServiceManagerTask{
			Name:             taskName,
			TaskCreatedBlock: taskCreatedBlock,
		},
		taskIndex,
		sig,
	)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(context.Background(), o.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return fmt.Errorf("tx failed: %v", receipt.Status)
	}
	log.Println("Responded to task.")
	return nil
}

func (o *Operator) MonitorNewTasks() error {

	newTaskCreated := make(chan *bindings.ContractNewTaskCreated)
	var start uint32 = 0
	go func() {
		ticker := time.NewTicker(3 * time.Second)

		for _ = range ticker.C {
			filterOpt := &bind.FilterOpts{
				Start: uint64(start),
				End:   nil,
			}
			iter, err := o.contract.FilterNewTaskCreated(filterOpt, nil)
			if err != nil {
				return
			}
			for iter.Next() {
				start = iter.Event.Task.TaskCreatedBlock + 1
				newTaskCreated <- iter.Event
			}
		}
	}()

	log.Println("Monitoring for new tasks...")
	for {
		select {
		case event := <-newTaskCreated:
			log.Printf("New task created: %s\n", event.Task.Name)
			o.SignAndRespondToTask(event.TaskIndex, event.Task.TaskCreatedBlock, event.Task.Name)
		}
	}

	return nil
}

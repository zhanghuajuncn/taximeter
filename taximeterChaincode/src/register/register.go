package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	var A, B string    // Entities
	var time, chaincodeID string

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	time = args[2]
	chaincodeID = args[3]

	err = stub.PutState(A+B+time, []byte(chaincodeID))
	if err != nil {
		return nil, err
	}
	return nil, nil
}


// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "invoke" {
		// set chaincodeID
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "invoke" {
		// set chaincodeID
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	if function != "query" {
		fmt.Printf("Function is query")
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A,B string // Entities
	var time string //Time
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]
	B = args[1]
	time = args[2]

	// Get the blockchainID from the ledger
	blockchainID, err := stub.GetState(A+B+time)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + B + time + "\"}"
		return nil, errors.New(jsonResp)
	}

	if blockchainID == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + B + time + "\"}"
		return nil, errors.New(jsonResp)
	}

	fmt.Printf("Query Response:%s\n", blockchainID)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
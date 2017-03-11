
package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	var A, B string    // Entities
	var timePrice, milePrice string // Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	B = args[1]
	timePrice = args[2]
	milePrice = args[3]


	// Write the state to the ledger
	err = stub.PutState("drv", []byte(A))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("psg", []byte(B))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("timePrice", []byte(timePrice))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("milePrice", []byte(milePrice))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	var role, time, jsonGPS string  
	//var lng, lat string
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	role = args[0]
	time = args[1]
	//lat = args[2]
	//lng = args[3]
	
	//jsonGPS = "{\"lat\":"+ lat + ",\"lng\":" + lng + "}"
	jsonGPS = "lalal"


	err = stub.PutState(role+time, []byte(jsonGPS))
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
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} 
	return nil, errors.New("Received unknown function invocation")
}

func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) queryGPS(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var role, time, jsonResp string // Entities
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	role = args[0]
	time = args[1]
	

	// Get the state from the ledger
	posval, err :=  stub.GetState(role+time)
	if err != nil {
		return nil, errors.New("Failed to get position")
	}
	if posval == nil {
		return nil, errors.New("Empty position")
	}

	jsonResp = string(posval)
	fmt.Printf("Query Response:%s\n", jsonResp)
	return posval, nil
}

func (t *SimpleChaincode) queryPrice(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var whichPrice string 
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	whichPrice = args[0]
	priceval, err :=  stub.GetState(whichPrice)
	if err != nil {
		return nil, errors.New("Failed to get position")
	}

	price := string(priceval)
	fmt.Printf("Query Response:%s\n", price)
	return priceval, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	if function == "queryGPS" {
		fmt.Printf("Query GPS")
		return t.queryGPS(stub, args)
	} else if function == "queryPrice" {
		fmt.Printf("Query price")
		return t.queryPrice(stub, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
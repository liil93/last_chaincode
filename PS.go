package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// var _CCstr string

type PS struct { // Petsitting chaincode
}

type TradeRec struct { // Trade record (KEY: PSID#CSID#Trade complete time)
	PSID       string // Petsitter ID
	PSNickname string
	CSID       string // Consumer ID
	TS         string // Transaction start time
	TE         string // Transaction end time
	TC         string // Transaction complete time
	TA         string // Transaction amount
	TH         string // Transaction history
}

// type CityCode struct {
// 	R103 string // Region code 103 (nowongu)
// 	R104 string // Region code 104 (gangnamgu)
// 	R105 string // Region code 105 (zongrogu)
// } // if string is not good, using map... ex Region[int]string / Region[103] = "key1"...

type Petsitter struct { // User information (KEY: User email)
	Nickname string
	CostL    string
	CostM    string
	CostS    string
	Start    string
	End      string
	Except   string
	TotalNum string
	NumL     string
	NumM     string
	NumS     string
	Home     string
	HomeInfo string
	SaveTime string
}

type HomeAsset struct { // Information about home (KEY: User email#home)
	State    string
	City     string
	Street   string
	Adt      string
	Code     string
	Type     string
	Room     string
	Elevator string
	Parking  string
}

// type PetAsset struct { // Information about pet (KEY: User email#pet)
// 	Name   string // Pet name
// 	Birth  string // Pet birth
// 	Gender string // Pet gender
// 	Kind   string // Pet kind
// 	Size   string // Pet size (S: ~5kg, M: 5~10, L: 10~)
// 	NS     string // Neutralization surgery
// 	Vac    string // Vaccine check
// }

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(PS))
	if err != nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                         <<<< Start chaincode >>>>")
		fmt.Printf("                     Error starting PS chaincode: %s", err)
		fmt.Println("=======================================================================")
		fmt.Println()
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *PS) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 0 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                              <<<< Init >>>>")
		fmt.Println("             Incorrect number of arguments. Expecting 0")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[INIT] Incorrect number of arguments. Expecting 0")
	}
	// _CCstr = "_CityCodeStruct"
	// cityCode := CityCode{}
	// cityCode.R103 = "/"
	// cityCode.R104 = "/"
	// cityCode.R105 = "/"
	// jsonAsBytes, _ := json.Marshal(cityCode)
	// stub.PutState(_CCstr, jsonAsBytes)
	//
	// fmt.Println()
	// fmt.Println("=======================================================================")
	// fmt.Println("                         <<<< Init >>>>")
	// fmt.Println("Init CityCode R013 : " + cityCode.R103)
	// fmt.Println("Init CityCode R014 : " + cityCode.R104)
	// fmt.Println("Init CityCode R015 : " + cityCode.R105)
	// fmt.Println("=======================================================================")
	// fmt.Println()

	return nil, nil
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *PS) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "save_petsitter" {
		return t.UInsert(stub, args)
	} else if function == "modify_petsitter" {
		return t.UChange(stub, args)
	} else if function == "delete_petsitter" {
		return t.UDelete(stub, args)
	} else if function == "save_home_address" {
		return t.SHomeAddress(stub, args)
	} else if function == "save_home_room" {
		return t.SHomeRoom(stub, args)
	} else if function == "save_home_car_elevator" {
		return t.SHomeCarElevator(stub, args)
	} else if function == "modify_home_address" {
		return t.HChangeAddress(stub, args)
	} else if function == "modify_home_room" {
		return t.HChangeRoom(stub, args)
	} else if function == "modify_home_car_elevator" {
		return t.HChangeCarElevator(stub, args)
	} else if function == "save_tran" {
		return t.TInsert(stub, args)
	}

	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Println("                              <<<< Invoke >>>>")
	fmt.Println("               Invoke did not find func: " + function)
	fmt.Println("=======================================================================")
	fmt.Println()

	return nil, errors.New("[INVOKE] Received unknown function invocation: " + function)
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *PS) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "read_petsitter" {
		return t.URead(stub, args)
	} else if function == "read_house" {
		return t.HRead(stub, args)
	} else if function == "search_tran" {
		return t.TSearch(stub, args)
	}
	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Println("                              <<<< Query >>>>")
	fmt.Println("               Query did not find func: " + function)
	fmt.Println("=======================================================================")
	fmt.Println()

	return nil, errors.New("[QUERY] Received unknown function query: " + function)
}

// ============================================================================================================================
// TSearch - search trade information
// ============================================================================================================================
func (t *PS) TSearch(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Trade Search >>>>")
		fmt.Println("                Incorrect number of arguments. Expecting 3")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[TRADE SEARCH] Incorrect number of arguments. Expecting 3")
	}
	psid := args[0]
	csid := args[1]
	tc := args[2]
	valAsbytes, _ := stub.GetState(psid + "#" + csid + "#" + tc)
	if valAsbytes == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Trade Search >>>>")
		fmt.Println("                           Not exist transaction")
		fmt.Println("=======================================================================")
		fmt.Println()
		return []byte("[TRADE SEARCH] Not exist transaction"), errors.New("[TRADE SEARCH] Not exist transaction")
	}
	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Println("                           <<<< Trade Search >>>>")
	fmt.Println("                           Trade reading success")
	fmt.Println("=======================================================================")
	fmt.Println()
	return valAsbytes, nil
}

// ============================================================================================================================
// TInsert - insert transaction information
// ============================================================================================================================
func (t *PS) TInsert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 8 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Trade Insert >>>>")
		fmt.Println("                Incorrect number of arguments. Expecting 8")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[TRADE INSSERT] Incorrect number of arguments. Expecting 8")
	}
	psid := args[0]
	psnick := args[1]
	csid := args[2]
	ts := args[3]
	te := args[4]
	tc := args[5]
	ta := args[6]
	th := args[7]

	tradeRec := TradeRec{}
	tradeRec.PSID = psid
	tradeRec.PSNickname = psnick
	tradeRec.CSID = csid
	tradeRec.TS = ts
	tradeRec.TE = te
	tradeRec.TC = tc
	tradeRec.TA = ta
	tradeRec.TH = th
	jsonAsBytes, _ := json.Marshal(tradeRec)
	stub.PutState(psid+"#"+csid+"#"+tc, jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// URead - read user information
// ============================================================================================================================
func (t *PS) HRead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Read >>>>")
		fmt.Println("                Incorrect number of arguments. Expecting 1")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home QUERY] Incorrect number of arguments. Expecting 1")
	}
	key := args[0] + "#home"
	valAsbytes, _ := stub.GetState(key) //get the pet information from chaincode state
	if valAsbytes == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Read >>>>")
		fmt.Println("                              Not exist Home")
		fmt.Println("=======================================================================")
		fmt.Println()
		return []byte("[Home QUERY] Not exist Home"), errors.New("[Home QUERY] Not exist Home")
	}
	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Println("                           <<<< Home Read >>>>")
	fmt.Println("                      Reading success, ID: " + key)
	fmt.Println("=======================================================================")
	fmt.Println()
	return valAsbytes, nil
}

// ============================================================================================================================
// UChange - change user information (PW, AP)
// ============================================================================================================================
func (t *PS) HChangeCarElevator(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Change >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 3")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home CHANGE] Incorrect number of arguments. Expecting 3")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Change >>>>")
		fmt.Println("                               Not exist Home")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home CHANGE] Not exist Home")
	}
	homeAsset := HomeAsset{}
	json.Unmarshal(confUser, &homeAsset)
	if args[1] != "none" {
		homeAsset.Elevator = args[1]
	}
	if args[2] != "none" {
		homeAsset.Parking = args[2]
	}

	jsonAsBytes, _ := json.Marshal(homeAsset)
	stub.PutState(args[0]+"#home", jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UChange - change user information (PW, AP)
// ============================================================================================================================
func (t *PS) HChangeRoom(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Change >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 3")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home CHANGE] Incorrect number of arguments. Expecting 3")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Change >>>>")
		fmt.Println("                               Not exist Home")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home CHANGE] Not exist Home")
	}
	homeAsset := HomeAsset{}
	json.Unmarshal(confUser, &homeAsset)
	if args[1] != "none" {
		homeAsset.Type = args[1]
	}
	if args[2] != "none" {
		homeAsset.Room = args[2]
	}

	jsonAsBytes, _ := json.Marshal(homeAsset)
	stub.PutState(args[0]+"#home", jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UChange - change user information (PW, AP)
// ============================================================================================================================
func (t *PS) HChangeAddress(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 6 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Change >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 6")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home CHANGE] Incorrect number of arguments. Expecting 6")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Home Change >>>>")
		fmt.Println("                               Not exist Home")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home CHANGE] Not exist Home")
	}
	homeAsset := HomeAsset{}
	json.Unmarshal(confUser, &homeAsset)
	if args[1] != "none" {
		homeAsset.State = args[1]
	}
	if args[2] != "none" {
		homeAsset.City = args[2]
	}
	if args[3] != "none" {
		homeAsset.Street = args[3]
	}
	if args[4] != "none" {
		homeAsset.Adt = args[4]
	}
	if args[5] != "none" {
		homeAsset.Code = args[5]
	}

	jsonAsBytes, _ := json.Marshal(homeAsset)
	stub.PutState(args[0]+"#home", jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UInsert - insert user information
// ============================================================================================================================
func (t *PS) SHomeAddress(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 6 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                          <<<< Home Insert >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 6")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home INSSERT] Incorrect number of arguments. Expecting 6")
	}
	conf, _ := stub.GetState(args[0])
	homeAsset := HomeAsset{}
	json.Unmarshal(conf, &homeAsset)
	homeAsset.State = args[1]
	homeAsset.City = args[2]
	homeAsset.Street = args[3]
	homeAsset.Adt = args[4]
	homeAsset.Code = args[5]
	jsonAsBytes, _ := json.Marshal(homeAsset)
	stub.PutState(args[0]+"#home", jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UInsert - insert user information
// ============================================================================================================================
func (t *PS) SHomeRoom(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                          <<<< Home Insert >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 3")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home INSSERT] Incorrect number of arguments. Expecting 3")
	}
	conf, _ := stub.GetState(args[0])
	homeAsset := HomeAsset{}
	json.Unmarshal(conf, &homeAsset)
	homeAsset.Type = args[1]
	homeAsset.Room = args[2]
	jsonAsBytes, _ := json.Marshal(homeAsset)
	stub.PutState(args[0]+"#home", jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UInsert - insert user information
// ============================================================================================================================
func (t *PS) SHomeCarElevator(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                          <<<< Home Insert >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 3")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Home INSSERT] Incorrect number of arguments. Expecting 3")
	}
	conf, _ := stub.GetState(args[0])
	homeAsset := HomeAsset{}
	json.Unmarshal(conf, &homeAsset)
	homeAsset.Elevator = args[1]
	homeAsset.Parking = args[2]
	jsonAsBytes, _ := json.Marshal(homeAsset)
	stub.PutState(args[0]+"#home", jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UInsert - insert user information
// ============================================================================================================================
func (t *PS) UInsert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 15 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                          <<<< Petsitter Insert >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 15")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter INSSERT] Incorrect number of arguments. Expecting 15")
	}
	conf, _ := stub.GetState(args[0])
	if conf != nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                          <<<< Petsitter Insert >>>>")
		fmt.Println("                            Already exist Petsitter")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter INSSERT] Already exist Petsitter")
	}
	email := args[0]
	time := time.Now()
	petsitter := Petsitter{args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], time.String()}
	jsonAsBytes, _ := json.Marshal(petsitter)
	stub.PutState(email, jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// UChange - change user information (PW, AP)
// ============================================================================================================================
func (t *PS) UChange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 15 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Petsitter Change >>>>")
		fmt.Println("               Incorrect number of arguments. Expecting 15")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter CHANGE] Incorrect number of arguments. Expecting 15")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Petsitter Change >>>>")
		fmt.Println("                               Not exist Petsitter")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter CHANGE] Not exist Petsitter")
	}
	petsitter := Petsitter{}
	json.Unmarshal(confUser, &petsitter)
	if args[1] != "none" {
		petsitter.Nickname = args[1]
	}
	if args[2] != "none" {
		petsitter.CostL = args[2]
	}
	if args[3] != "none" {
		petsitter.CostM = args[3]
	}
	if args[4] != "none" {
		petsitter.CostS = args[4]
	}
	if args[5] != "none" {
		petsitter.Start = args[5]
	}
	if args[6] != "none" {
		petsitter.End = args[6]
	}
	if args[7] != "none" {
		petsitter.Except = args[7]
	}
	if args[8] != "none" {
		petsitter.TotalNum = args[8]
	}
	if args[9] != "none" {
		petsitter.NumL = args[9]
	}
	if args[10] != "none" {
		petsitter.NumM = args[10]
	}
	if args[11] != "none" {
		petsitter.NumS = args[11]
	}
	if args[12] != "none" {
		petsitter.Home = args[12]
	}
	if args[13] != "none" {
		petsitter.HomeInfo = args[13]
	}
	petsitter.SaveTime = time.Now().String()

	jsonAsBytes, _ := json.Marshal(petsitter)
	stub.PutState(args[0], jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// URead - read user information
// ============================================================================================================================
func (t *PS) URead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Petsitter Read >>>>")
		fmt.Println("                Incorrect number of arguments. Expecting 1")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter QUERY] Incorrect number of arguments. Expecting 1")
	}
	key := args[0]
	valAsbytes, _ := stub.GetState(key) //get the pet information from chaincode state
	if valAsbytes == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Petsitter Read >>>>")
		fmt.Println("                              Not exist Petsitter")
		fmt.Println("=======================================================================")
		fmt.Println()
		return []byte("[Petsitter QUERY] Not exist Petsitter"), errors.New("[Petsitter QUERY] Not exist Petsitter")
	}
	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Println("                           <<<< Petsitter Read >>>>")
	fmt.Println("                      Reading success, ID: " + key)
	fmt.Println("=======================================================================")
	fmt.Println()
	return valAsbytes, nil
}

// ============================================================================================================================
// UDelete - delete Petsitter information
// ============================================================================================================================
func (t *PS) UDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                        <<<< Petsitter Delete >>>>")
		fmt.Println("                Incorrect number of arguments. Expecting 1")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter DELETE] Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]
	conf, _ := stub.GetState(userID)
	if conf == nil {
		fmt.Println()
		fmt.Println("=======================================================================")
		fmt.Println("                           <<<< Petsitter Delete >>>>")
		fmt.Println("                              Not exist Petsitter")
		fmt.Println("=======================================================================")
		fmt.Println()
		return nil, errors.New("[Petsitter DELETE] Not exist Petsitter")
	}
	stub.DelState(args[0])

	return nil, nil
}
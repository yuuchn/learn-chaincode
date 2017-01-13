/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

// AssetChaincode example simple Chaincode implementation
type AssetChaincode struct {
}

// ============================================================================================================================
// 定数Key
// ============================================================================================================================
/* PC_A */
const PC_A string = "pc_a"
/* PC_B */
const PC_B string = "pc_b"
/* PC_B */
const PC_C string = "pc_c"

/* Moblile Wifi A */
const WIFI_A string = "wifi_a"
/* Moblile Wifi B */
const WIFI_B string = "wifi_b"
/* Moblile Wifi C */
const WIFI_C string = "wifi_c"
/* Moblile Wifi D */
const WIFI_D string = "wifi_d"

// ============================================================================================================================
// Main
// ============================================================================================================================
// Validating Peerに接続し、チェーンコードを実行
func main() {
	err := shim.Start(new(AssetChaincode))
		if err != nil {
		fmt.Printf("Error starting Asset chaincode: %s", err)
	}
}

// 資産の所有者情報を初期値
func (t *AssetChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// TODO 今回は引数のチェックは行わない。
	//if len(args) != 1 {
	//	return nil, errors.New("Incorrect number of arguments. Expecting 1")
	//}

	var timeStmp string
	timeStmp = time.Now().Format("2006/01/02-15:04:05")
	// ワールドステート（台帳）に追加
	// PC
	stub.PutState(PC_A, []byte("ict" + "," + timeStmp))
	stub.PutState(PC_B, []byte("ict" + "," + timeStmp))
	stub.PutState(PC_C, []byte("ict" + "," + timeStmp))
	// モバイルWifi
	stub.PutState(WIFI_A, []byte("ict" + "," + timeStmp))
	stub.PutState(WIFI_B, []byte("ict" + "," + timeStmp))
	stub.PutState(WIFI_C, []byte("ict" + "," + timeStmp))
	stub.PutState(WIFI_D, []byte("ict" + "," + timeStmp))

	return nil, nil
}

// 資産の所有者情報を更新
func (t *AssetChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// function名でハンドリング
	if function == "init" {
		// 所有者情報を初期化
		return t.Init(stub, "init", args)
	} else if function == "write" {
		// 所有者情報の更新
		return t.write(stub, args)
	}
	// 定義外の関数
	fmt.Println("invoke did not find func: " + function)
	// エラー
	return nil, errors.New("Received unknown function invocation: " + function)
}

// 所有者情報を参照
func (t *AssetChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// function名でハンドリング
	if function == "read" {
		// 所有者情報の取得
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// 所有者情報の更新
func (t *AssetChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error

	fmt.Println("running write()")

	// 引数にKey/Valueのペアがない場合はエラーを返却
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0]   // Key
	value = args[1] // Value

	// Keyチェック
	if strings.Contains(key, PC_A) {
		// 何もしない
	} else if strings.Contains(key, PC_B) {
		// 何もしない
	} else if strings.Contains(key, PC_B) {
		// 何もしない
	} else if strings.Contains(key, PC_C) {
		// 何もしない
	} else if strings.Contains(key, WIFI_A) {
		// 何もしない
	} else if strings.Contains(key, WIFI_B) {
		// 何もしない
	} else if strings.Contains(key, WIFI_C) {
		// 何もしない
	} else if strings.Contains(key, WIFI_D) {
		// 何もしない
	} else {
		// 定義外のKeyの場合はエラーを返却
		return nil, errors.New("Incorrect Key of arguments: " + key)
	}

	// TODO 対象の履歴を取得
	valAsbytes, err := stub.GetState(key)

	var timeStmp string
	var strByte string

	strByte = string([]byte(valAsbytes))
	timeStmp = time.Now().Format("2006/01/02-15:04:05")

	if valAsbytes != nil {
		// 既に登録情報が存在する場合
		// 追記(既存値, 値 + タイムスタンプ)
		strByte = strByte + "|" + value + "," + timeStmp
	} else {
		strByte = value + "," + timeStmp
	}

	// 所有者情報を更新
	valAsbytes = []byte(strByte)
	err = stub.PutState(key, valAsbytes)

	// 更新時にエラーが発生した場合
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 所有者情報の取得
func (t *AssetChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	// 引数チェック
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

package data

import (
	"encoding/json"
	"fmt"
)

type RegisterData struct {
	AssignedId int32 `json:"assignedId"`
	PeerMapJson string `json:"peerMapJson"`
}

/**
Return a New Register Data struct
 */
func NewRegisterData(id int32, peerMapJson string) RegisterData {

	return RegisterData{id, peerMapJson}
}

/**
Encode the Register data into Json Format
 */
func (data *RegisterData) EncodeToJson() (string, error) {

	jsonFormatted, err := json.MarshalIndent(data, "", "")

	if err != nil {
		fmt.Println("Error in EncodeToJson")
		return "", err
	}

	return string(jsonFormatted), nil
}
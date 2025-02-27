package main

import (
	"encoding/json"
	"fmt"
)

type ShipmentsInputModel struct {
	LotID
}

type LotID struct {
	LotId  string `json:"lotId,omitempty"`
	Lot_ID string `json:"lot_id,omitempty"`
}

func (s *ShipmentsInputModel) setLodID(id string) {
	s.LotId = id
	s.Lot_ID = id
}

func main() {
	shipment := ShipmentsInputModel{}
	shipment.setLodID("someLotID")

	// Convert struct to JSON
	jsonData, err := json.Marshal(shipment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// prints: {"lotId":"someLotID","lot_id":"someLotID"}
	fmt.Println(string(jsonData))
}

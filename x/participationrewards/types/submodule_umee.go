package types

import (
	"encoding/json"
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
)

type UmeeProtocolData struct {
	Denom       string
	LastUpdated time.Time
	Data        json.RawMessage
}

func (upd *UmeeProtocolData) ValidateBasic() error {
	if upd.Denom == "" {
		return ErrUndefinedAttribute
	}
	return nil
}

func (upd *UmeeProtocolData) GenerateKey() []byte {
	return []byte(upd.Denom)
}

func GetUnderlyingData[V sdkmath.Int | sdkmath.LegacyDec](upd *UmeeProtocolData) (V, error) {
	var data V
	err := json.Unmarshal(upd.Data, &data)
	if err != nil {
		return data, fmt.Errorf("1: unable to unmarshal concrete reservedata: %w", err)
	}
	return data, nil
}

// UmeeReservesProtocolData defines protocol state to track qAssets in
// Umee reserves.
type UmeeReservesProtocolData struct {
	UmeeProtocolData
}

func (upd *UmeeReservesProtocolData) GetReserveAmount() (math.Int, error) {
	return GetUnderlyingData[math.Int](&upd.UmeeProtocolData)
}

type UmeeTotalBorrowsProtocolData struct {
	UmeeProtocolData
}

func (upd *UmeeTotalBorrowsProtocolData) GetTotalBorrows() (sdkmath.LegacyDec, error) {
	return GetUnderlyingData[sdkmath.LegacyDec](&upd.UmeeProtocolData)
}

type UmeeInterestScalarProtocolData struct {
	UmeeProtocolData
}

func (upd *UmeeInterestScalarProtocolData) GetInterestScalar() (sdkmath.LegacyDec, error) {
	return GetUnderlyingData[sdkmath.LegacyDec](&upd.UmeeProtocolData)
}

type UmeeUTokenSupplyProtocolData struct {
	UmeeProtocolData
}

func (upd *UmeeUTokenSupplyProtocolData) GetUTokenSupply() (math.Int, error) {
	return GetUnderlyingData[sdkmath.Int](&upd.UmeeProtocolData)
}

type UmeeLeverageModuleBalanceProtocolData struct {
	UmeeProtocolData
}

func (upd *UmeeLeverageModuleBalanceProtocolData) GetModuleBalance() (sdkmath.Int, error) {
	return GetUnderlyingData[sdkmath.Int](&upd.UmeeProtocolData)
}

// -----------------------------------------------------

type UmeeParamsProtocolData struct {
	ChainID string
}

func (uppd UmeeParamsProtocolData) ValidateBasic() error {
	if uppd.ChainID == "" {
		return ErrUndefinedAttribute
	}

	return nil
}

func (UmeeParamsProtocolData) GenerateKey() []byte {
	return []byte(UmeeParamsKey)
}

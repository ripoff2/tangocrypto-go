package tangocrypto_go

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	resourceEpochs     = "epochs"
	resourceParameters = "parameters"
	resourceCurrent    = "current"
)

type EpochParameters struct {
	EpochNo               int       `json:"epoch_no"`
	MinFeeA               int       `json:"min_fee_a"`
	MinFeeB               int       `json:"min_fee_b"`
	MaxBlockSize          int       `json:"max_block_size"`
	MaxTxSize             int       `json:"max_tx_size"`
	MaxBlockHeaderSize    int       `json:"max_block_header_size"`
	KeyDeposit            int       `json:"key_deposit"`
	PoolDeposit           int       `json:"pool_deposit"`
	MaxEpoch              int       `json:"max_epoch"`
	OptimalPoolCount      int       `json:"optimal_pool_count"`
	InfluenceA0           float64   `json:"influence_a0"`
	MonetaryExpandRateRho float64   `json:"monetary_expand_rate_rho"`
	TreasuryGrowthRateTau float64   `json:"treasury_growth_rate_tau"`
	Decentralisation      int       `json:"decentralisation"`
	ExtraEntropy          string    `json:"extra_entropy"`
	ProtocolMajor         int       `json:"protocol_major"`
	ProtocolMinor         int       `json:"protocol_minor"`
	MinUtxo               int       `json:"min_utxo"`
	MinPoolCost           int       `json:"min_pool_cost"`
	Nonce                 string    `json:"nonce"`
	CoinsPerUtxoSize      int       `json:"coins_per_utxo_size"`
	PriceMem              float64   `json:"price_mem"`
	PriceStep             float64   `json:"price_step"`
	MaxTxExMem            int       `json:"max_tx_ex_mem"`
	MaxTxExSteps          int64     `json:"max_tx_ex_steps"`
	MaxBlockExMem         int       `json:"max_block_ex_mem"`
	MaxBlockExSteps       int64     `json:"max_block_ex_steps"`
	MaxValSize            int       `json:"max_val_size"`
	CollateralPercent     int       `json:"collateral_percent"`
	MaxCollateralInputs   int       `json:"max_collateral_inputs"`
	CostModel             CostModel `json:"cost_model"`
}

type PlutusV1 struct {
	BDataCPUArguments                                   int `json:"bData-cpu-arguments"`
	IDataCPUArguments                                   int `json:"iData-cpu-arguments"`
	TraceCPUArguments                                   int `json:"trace-cpu-arguments"`
	MkConsCPUArguments                                  int `json:"mkCons-cpu-arguments"`
	FstPairCPUArguments                                 int `json:"fstPair-cpu-arguments"`
	MapDataCPUArguments                                 int `json:"mapData-cpu-arguments"`
	SndPairCPUArguments                                 int `json:"sndPair-cpu-arguments"`
	UnBDataCPUArguments                                 int `json:"unBData-cpu-arguments"`
	UnIDataCPUArguments                                 int `json:"unIData-cpu-arguments"`
	BDataMemoryArguments                                int `json:"bData-memory-arguments"`
	CekLamCostExBudgetCPU                               int `json:"cekLamCost-exBudgetCPU"`
	CekVarCostExBudgetCPU                               int `json:"cekVarCost-exBudgetCPU"`
	HeadListCPUArguments                                int `json:"headList-cpu-arguments"`
	IDataMemoryArguments                                int `json:"iData-memory-arguments"`
	ListDataCPUArguments                                int `json:"listData-cpu-arguments"`
	NullListCPUArguments                                int `json:"nullList-cpu-arguments"`
	TailListCPUArguments                                int `json:"tailList-cpu-arguments"`
	TraceMemoryArguments                                int `json:"trace-memory-arguments"`
	MkConsMemoryArguments                               int `json:"mkCons-memory-arguments"`
	MkNilDataCPUArguments                               int `json:"mkNilData-cpu-arguments"`
	UnMapDataCPUArguments                               int `json:"unMapData-cpu-arguments"`
	CekApplyCostExBudgetCPU                             int `json:"cekApplyCost-exBudgetCPU"`
	CekConstCostExBudgetCPU                             int `json:"cekConstCost-exBudgetCPU"`
	CekDelayCostExBudgetCPU                             int `json:"cekDelayCost-exBudgetCPU"`
	CekForceCostExBudgetCPU                             int `json:"cekForceCost-exBudgetCPU"`
	ChooseDataCPUArguments                              int `json:"chooseData-cpu-arguments"`
	ChooseListCPUArguments                              int `json:"chooseList-cpu-arguments"`
	ChooseUnitCPUArguments                              int `json:"chooseUnit-cpu-arguments"`
	ConstrDataCPUArguments                              int `json:"constrData-cpu-arguments"`
	FstPairMemoryArguments                              int `json:"fstPair-memory-arguments"`
	IfThenElseCPUArguments                              int `json:"ifThenElse-cpu-arguments"`
	MapDataMemoryArguments                              int `json:"mapData-memory-arguments"`
	MkPairDataCPUArguments                              int `json:"mkPairData-cpu-arguments"`
	SndPairMemoryArguments                              int `json:"sndPair-memory-arguments"`
	UnBDataMemoryArguments                              int `json:"unBData-memory-arguments"`
	UnIDataMemoryArguments                              int `json:"unIData-memory-arguments"`
	UnListDataCPUArguments                              int `json:"unListData-cpu-arguments"`
	CekLamCostExBudgetMemory                            int `json:"cekLamCost-exBudgetMemory"`
	CekVarCostExBudgetMemory                            int `json:"cekVarCost-exBudgetMemory"`
	HeadListMemoryArguments                             int `json:"headList-memory-arguments"`
	ListDataMemoryArguments                             int `json:"listData-memory-arguments"`
	NullListMemoryArguments                             int `json:"nullList-memory-arguments"`
	Sha2256MemoryArguments                              int `json:"sha2_256-memory-arguments"`
	Sha3256MemoryArguments                              int `json:"sha3_256-memory-arguments"`
	TailListMemoryArguments                             int `json:"tailList-memory-arguments"`
	CekBuiltinCostExBudgetCPU                           int `json:"cekBuiltinCost-exBudgetCPU"`
	CekStartupCostExBudgetCPU                           int `json:"cekStartupCost-exBudgetCPU"`
	MkNilDataMemoryArguments                            int `json:"mkNilData-memory-arguments"`
	UnConstrDataCPUArguments                            int `json:"unConstrData-cpu-arguments"`
	UnMapDataMemoryArguments                            int `json:"unMapData-memory-arguments"`
	CekApplyCostExBudgetMemory                          int `json:"cekApplyCost-exBudgetMemory"`
	CekConstCostExBudgetMemory                          int `json:"cekConstCost-exBudgetMemory"`
	CekDelayCostExBudgetMemory                          int `json:"cekDelayCost-exBudgetMemory"`
	CekForceCostExBudgetMemory                          int `json:"cekForceCost-exBudgetMemory"`
	ChooseDataMemoryArguments                           int `json:"chooseData-memory-arguments"`
	ChooseListMemoryArguments                           int `json:"chooseList-memory-arguments"`
	ChooseUnitMemoryArguments                           int `json:"chooseUnit-memory-arguments"`
	ConstrDataMemoryArguments                           int `json:"constrData-memory-arguments"`
	EqualsDataMemoryArguments                           int `json:"equalsData-memory-arguments"`
	IfThenElseMemoryArguments                           int `json:"ifThenElse-memory-arguments"`
	MkNilPairDataCPUArguments                           int `json:"mkNilPairData-cpu-arguments"`
	MkPairDataMemoryArguments                           int `json:"mkPairData-memory-arguments"`
	UnListDataMemoryArguments                           int `json:"unListData-memory-arguments"`
	Blake2B256MemoryArguments                           int `json:"blake2b_256-memory-arguments"`
	Sha2256CPUArgumentsSlope                            int `json:"sha2_256-cpu-arguments-slope"`
	Sha3256CPUArgumentsSlope                            int `json:"sha3_256-cpu-arguments-slope"`
	CekBuiltinCostExBudgetMemory                        int `json:"cekBuiltinCost-exBudgetMemory"`
	CekStartupCostExBudgetMemory                        int `json:"cekStartupCost-exBudgetMemory"`
	EqualsStringMemoryArguments                         int `json:"equalsString-memory-arguments"`
	IndexByteStringCPUArguments                         int `json:"indexByteString-cpu-arguments"`
	UnConstrDataMemoryArguments                         int `json:"unConstrData-memory-arguments"`
	AddIntegerCPUArgumentsSlope                         int `json:"addInteger-cpu-arguments-slope"`
	DecodeUtf8CPUArgumentsSlope                         int `json:"decodeUtf8-cpu-arguments-slope"`
	EncodeUtf8CPUArgumentsSlope                         int `json:"encodeUtf8-cpu-arguments-slope"`
	EqualsDataCPUArgumentsSlope                         int `json:"equalsData-cpu-arguments-slope"`
	EqualsIntegerMemoryArguments                        int `json:"equalsInteger-memory-arguments"`
	MkNilPairDataMemoryArguments                        int `json:"mkNilPairData-memory-arguments"`
	Blake2B256CPUArgumentsSlope                         int `json:"blake2b_256-cpu-arguments-slope"`
	AppendStringCPUArgumentsSlope                       int `json:"appendString-cpu-arguments-slope"`
	EqualsStringCPUArgumentsSlope                       int `json:"equalsString-cpu-arguments-slope"`
	IndexByteStringMemoryArguments                      int `json:"indexByteString-memory-arguments"`
	LengthOfByteStringCPUArguments                      int `json:"lengthOfByteString-cpu-arguments"`
	LessThanIntegerMemoryArguments                      int `json:"lessThanInteger-memory-arguments"`
	Sha2256CPUArgumentsIntercept                        int `json:"sha2_256-cpu-arguments-intercept"`
	Sha3256CPUArgumentsIntercept                        int `json:"sha3_256-cpu-arguments-intercept"`
	AddIntegerMemoryArgumentsSlope                      int `json:"addInteger-memory-arguments-slope"`
	DecodeUtf8MemoryArgumentsSlope                      int `json:"decodeUtf8-memory-arguments-slope"`
	EncodeUtf8MemoryArgumentsSlope                      int `json:"encodeUtf8-memory-arguments-slope"`
	EqualsByteStringMemoryArguments                     int `json:"equalsByteString-memory-arguments"`
	EqualsIntegerCPUArgumentsSlope                      int `json:"equalsInteger-cpu-arguments-slope"`
	ModIntegerCPUArgumentsConstant                      int `json:"modInteger-cpu-arguments-constant"`
	ModIntegerMemoryArgumentsSlope                      int `json:"modInteger-memory-arguments-slope"`
	AddIntegerCPUArgumentsIntercept                     int `json:"addInteger-cpu-arguments-intercept"`
	ConsByteStringCPUArgumentsSlope                     int `json:"consByteString-cpu-arguments-slope"`
	DecodeUtf8CPUArgumentsIntercept                     int `json:"decodeUtf8-cpu-arguments-intercept"`
	EncodeUtf8CPUArgumentsIntercept                     int `json:"encodeUtf8-cpu-arguments-intercept"`
	EqualsDataCPUArgumentsIntercept                     int `json:"equalsData-cpu-arguments-intercept"`
	AppendStringMemoryArgumentsSlope                    int `json:"appendString-memory-arguments-slope"`
	Blake2B256CPUArgumentsIntercept                     int `json:"blake2b_256-cpu-arguments-intercept"`
	EqualsStringCPUArgumentsConstant                    int `json:"equalsString-cpu-arguments-constant"`
	LengthOfByteStringMemoryArguments                   int `json:"lengthOfByteString-memory-arguments"`
	LessThanByteStringMemoryArguments                   int `json:"lessThanByteString-memory-arguments"`
	LessThanIntegerCPUArgumentsSlope                    int `json:"lessThanInteger-cpu-arguments-slope"`
	ModIntegerMemoryArgumentsMinimum                    int `json:"modInteger-memory-arguments-minimum"`
	MultiplyIntegerCPUArgumentsSlope                    int `json:"multiplyInteger-cpu-arguments-slope"`
	SliceByteStringCPUArgumentsSlope                    int `json:"sliceByteString-cpu-arguments-slope"`
	SubtractIntegerCPUArgumentsSlope                    int `json:"subtractInteger-cpu-arguments-slope"`
	AppendByteStringCPUArgumentsSlope                   int `json:"appendByteString-cpu-arguments-slope"`
	AppendStringCPUArgumentsIntercept                   int `json:"appendString-cpu-arguments-intercept"`
	DivideIntegerCPUArgumentsConstant                   int `json:"divideInteger-cpu-arguments-constant"`
	DivideIntegerMemoryArgumentsSlope                   int `json:"divideInteger-memory-arguments-slope"`
	EqualsByteStringCPUArgumentsSlope                   int `json:"equalsByteString-cpu-arguments-slope"`
	EqualsStringCPUArgumentsIntercept                   int `json:"equalsString-cpu-arguments-intercept"`
	AddIntegerMemoryArgumentsIntercept                  int `json:"addInteger-memory-arguments-intercept"`
	ConsByteStringMemoryArgumentsSlope                  int `json:"consByteString-memory-arguments-slope"`
	DecodeUtf8MemoryArgumentsIntercept                  int `json:"decodeUtf8-memory-arguments-intercept"`
	EncodeUtf8MemoryArgumentsIntercept                  int `json:"encodeUtf8-memory-arguments-intercept"`
	EqualsIntegerCPUArgumentsIntercept                  int `json:"equalsInteger-cpu-arguments-intercept"`
	ModIntegerMemoryArgumentsIntercept                  int `json:"modInteger-memory-arguments-intercept"`
	ConsByteStringCPUArgumentsIntercept                 int `json:"consByteString-cpu-arguments-intercept"`
	DivideIntegerMemoryArgumentsMinimum                 int `json:"divideInteger-memory-arguments-minimum"`
	LessThanByteStringCPUArgumentsSlope                 int `json:"lessThanByteString-cpu-arguments-slope"`
	LessThanEqualsIntegerMemoryArguments                int `json:"lessThanEqualsInteger-memory-arguments"`
	MultiplyIntegerMemoryArgumentsSlope                 int `json:"multiplyInteger-memory-arguments-slope"`
	QuotientIntegerCPUArgumentsConstant                 int `json:"quotientInteger-cpu-arguments-constant"`
	QuotientIntegerMemoryArgumentsSlope                 int `json:"quotientInteger-memory-arguments-slope"`
	SliceByteStringMemoryArgumentsSlope                 int `json:"sliceByteString-memory-arguments-slope"`
	SubtractIntegerMemoryArgumentsSlope                 int `json:"subtractInteger-memory-arguments-slope"`
	AppendByteStringMemoryArgumentsSlope                int `json:"appendByteString-memory-arguments-slope"`
	AppendStringMemoryArgumentsIntercept                int `json:"appendString-memory-arguments-intercept"`
	EqualsByteStringCPUArgumentsConstant                int `json:"equalsByteString-cpu-arguments-constant"`
	LessThanIntegerCPUArgumentsIntercept                int `json:"lessThanInteger-cpu-arguments-intercept"`
	MultiplyIntegerCPUArgumentsIntercept                int `json:"multiplyInteger-cpu-arguments-intercept"`
	RemainderIntegerCPUArgumentsConstant                int `json:"remainderInteger-cpu-arguments-constant"`
	RemainderIntegerMemoryArgumentsSlope                int `json:"remainderInteger-memory-arguments-slope"`
	SliceByteStringCPUArgumentsIntercept                int `json:"sliceByteString-cpu-arguments-intercept"`
	SubtractIntegerCPUArgumentsIntercept                int `json:"subtractInteger-cpu-arguments-intercept"`
	VerifyEd25519SignatureMemoryArguments               int `json:"verifyEd25519Signature-memory-arguments"`
	AppendByteStringCPUArgumentsIntercept               int `json:"appendByteString-cpu-arguments-intercept"`
	DivideIntegerMemoryArgumentsIntercept               int `json:"divideInteger-memory-arguments-intercept"`
	EqualsByteStringCPUArgumentsIntercept               int `json:"equalsByteString-cpu-arguments-intercept"`
	QuotientIntegerMemoryArgumentsMinimum               int `json:"quotientInteger-memory-arguments-minimum"`
	ConsByteStringMemoryArgumentsIntercept              int `json:"consByteString-memory-arguments-intercept"`
	LessThanEqualsByteStringMemoryArguments             int `json:"lessThanEqualsByteString-memory-arguments"`
	LessThanEqualsIntegerCPUArgumentsSlope              int `json:"lessThanEqualsInteger-cpu-arguments-slope"`
	RemainderIntegerMemoryArgumentsMinimum              int `json:"remainderInteger-memory-arguments-minimum"`
	LessThanByteStringCPUArgumentsIntercept             int `json:"lessThanByteString-cpu-arguments-intercept"`
	MultiplyIntegerMemoryArgumentsIntercept             int `json:"multiplyInteger-memory-arguments-intercept"`
	QuotientIntegerMemoryArgumentsIntercept             int `json:"quotientInteger-memory-arguments-intercept"`
	SliceByteStringMemoryArgumentsIntercept             int `json:"sliceByteString-memory-arguments-intercept"`
	SubtractIntegerMemoryArgumentsIntercept             int `json:"subtractInteger-memory-arguments-intercept"`
	VerifyEd25519SignatureCPUArgumentsSlope             int `json:"verifyEd25519Signature-cpu-arguments-slope"`
	AppendByteStringMemoryArgumentsIntercept            int `json:"appendByteString-memory-arguments-intercept"`
	RemainderIntegerMemoryArgumentsIntercept            int `json:"remainderInteger-memory-arguments-intercept"`
	LessThanEqualsByteStringCPUArgumentsSlope           int `json:"lessThanEqualsByteString-cpu-arguments-slope"`
	LessThanEqualsIntegerCPUArgumentsIntercept          int `json:"lessThanEqualsInteger-cpu-arguments-intercept"`
	ModIntegerCPUArgumentsModelArgumentsSlope           int `json:"modInteger-cpu-arguments-model-arguments-slope"`
	VerifyEd25519SignatureCPUArgumentsIntercept         int `json:"verifyEd25519Signature-cpu-arguments-intercept"`
	LessThanEqualsByteStringCPUArgumentsIntercept       int `json:"lessThanEqualsByteString-cpu-arguments-intercept"`
	DivideIntegerCPUArgumentsModelArgumentsSlope        int `json:"divideInteger-cpu-arguments-model-arguments-slope"`
	ModIntegerCPUArgumentsModelArgumentsIntercept       int `json:"modInteger-cpu-arguments-model-arguments-intercept"`
	QuotientIntegerCPUArgumentsModelArgumentsSlope      int `json:"quotientInteger-cpu-arguments-model-arguments-slope"`
	RemainderIntegerCPUArgumentsModelArgumentsSlope     int `json:"remainderInteger-cpu-arguments-model-arguments-slope"`
	DivideIntegerCPUArgumentsModelArgumentsIntercept    int `json:"divideInteger-cpu-arguments-model-arguments-intercept"`
	QuotientIntegerCPUArgumentsModelArgumentsIntercept  int `json:"quotientInteger-cpu-arguments-model-arguments-intercept"`
	RemainderIntegerCPUArgumentsModelArgumentsIntercept int `json:"remainderInteger-cpu-arguments-model-arguments-intercept"`
}

type PlutusV2 struct {
	BDataCPUArguments                                    int   `json:"bData-cpu-arguments"`
	IDataCPUArguments                                    int   `json:"iData-cpu-arguments"`
	TraceCPUArguments                                    int   `json:"trace-cpu-arguments"`
	MkConsCPUArguments                                   int   `json:"mkCons-cpu-arguments"`
	FstPairCPUArguments                                  int   `json:"fstPair-cpu-arguments"`
	MapDataCPUArguments                                  int   `json:"mapData-cpu-arguments"`
	SndPairCPUArguments                                  int   `json:"sndPair-cpu-arguments"`
	UnBDataCPUArguments                                  int   `json:"unBData-cpu-arguments"`
	UnIDataCPUArguments                                  int   `json:"unIData-cpu-arguments"`
	BDataMemoryArguments                                 int   `json:"bData-memory-arguments"`
	CekLamCostExBudgetCPU                                int   `json:"cekLamCost-exBudgetCPU"`
	CekVarCostExBudgetCPU                                int   `json:"cekVarCost-exBudgetCPU"`
	HeadListCPUArguments                                 int   `json:"headList-cpu-arguments"`
	IDataMemoryArguments                                 int   `json:"iData-memory-arguments"`
	ListDataCPUArguments                                 int   `json:"listData-cpu-arguments"`
	NullListCPUArguments                                 int   `json:"nullList-cpu-arguments"`
	TailListCPUArguments                                 int   `json:"tailList-cpu-arguments"`
	TraceMemoryArguments                                 int   `json:"trace-memory-arguments"`
	MkConsMemoryArguments                                int   `json:"mkCons-memory-arguments"`
	MkNilDataCPUArguments                                int   `json:"mkNilData-cpu-arguments"`
	UnMapDataCPUArguments                                int   `json:"unMapData-cpu-arguments"`
	CekApplyCostExBudgetCPU                              int   `json:"cekApplyCost-exBudgetCPU"`
	CekConstCostExBudgetCPU                              int   `json:"cekConstCost-exBudgetCPU"`
	CekDelayCostExBudgetCPU                              int   `json:"cekDelayCost-exBudgetCPU"`
	CekForceCostExBudgetCPU                              int   `json:"cekForceCost-exBudgetCPU"`
	ChooseDataCPUArguments                               int   `json:"chooseData-cpu-arguments"`
	ChooseListCPUArguments                               int   `json:"chooseList-cpu-arguments"`
	ChooseUnitCPUArguments                               int   `json:"chooseUnit-cpu-arguments"`
	ConstrDataCPUArguments                               int   `json:"constrData-cpu-arguments"`
	FstPairMemoryArguments                               int   `json:"fstPair-memory-arguments"`
	IfThenElseCPUArguments                               int   `json:"ifThenElse-cpu-arguments"`
	MapDataMemoryArguments                               int   `json:"mapData-memory-arguments"`
	MkPairDataCPUArguments                               int   `json:"mkPairData-cpu-arguments"`
	SndPairMemoryArguments                               int   `json:"sndPair-memory-arguments"`
	UnBDataMemoryArguments                               int   `json:"unBData-memory-arguments"`
	UnIDataMemoryArguments                               int   `json:"unIData-memory-arguments"`
	UnListDataCPUArguments                               int   `json:"unListData-cpu-arguments"`
	CekLamCostExBudgetMemory                             int   `json:"cekLamCost-exBudgetMemory"`
	CekVarCostExBudgetMemory                             int   `json:"cekVarCost-exBudgetMemory"`
	HeadListMemoryArguments                              int   `json:"headList-memory-arguments"`
	ListDataMemoryArguments                              int   `json:"listData-memory-arguments"`
	NullListMemoryArguments                              int   `json:"nullList-memory-arguments"`
	Sha2256MemoryArguments                               int   `json:"sha2_256-memory-arguments"`
	Sha3256MemoryArguments                               int   `json:"sha3_256-memory-arguments"`
	TailListMemoryArguments                              int   `json:"tailList-memory-arguments"`
	CekBuiltinCostExBudgetCPU                            int   `json:"cekBuiltinCost-exBudgetCPU"`
	CekStartupCostExBudgetCPU                            int   `json:"cekStartupCost-exBudgetCPU"`
	MkNilDataMemoryArguments                             int   `json:"mkNilData-memory-arguments"`
	UnConstrDataCPUArguments                             int   `json:"unConstrData-cpu-arguments"`
	UnMapDataMemoryArguments                             int   `json:"unMapData-memory-arguments"`
	CekApplyCostExBudgetMemory                           int   `json:"cekApplyCost-exBudgetMemory"`
	CekConstCostExBudgetMemory                           int   `json:"cekConstCost-exBudgetMemory"`
	CekDelayCostExBudgetMemory                           int   `json:"cekDelayCost-exBudgetMemory"`
	CekForceCostExBudgetMemory                           int   `json:"cekForceCost-exBudgetMemory"`
	ChooseDataMemoryArguments                            int   `json:"chooseData-memory-arguments"`
	ChooseListMemoryArguments                            int   `json:"chooseList-memory-arguments"`
	ChooseUnitMemoryArguments                            int   `json:"chooseUnit-memory-arguments"`
	ConstrDataMemoryArguments                            int   `json:"constrData-memory-arguments"`
	EqualsDataMemoryArguments                            int   `json:"equalsData-memory-arguments"`
	IfThenElseMemoryArguments                            int   `json:"ifThenElse-memory-arguments"`
	MkNilPairDataCPUArguments                            int   `json:"mkNilPairData-cpu-arguments"`
	MkPairDataMemoryArguments                            int   `json:"mkPairData-memory-arguments"`
	UnListDataMemoryArguments                            int   `json:"unListData-memory-arguments"`
	Blake2B256MemoryArguments                            int   `json:"blake2b_256-memory-arguments"`
	Sha2256CPUArgumentsSlope                             int   `json:"sha2_256-cpu-arguments-slope"`
	Sha3256CPUArgumentsSlope                             int   `json:"sha3_256-cpu-arguments-slope"`
	CekBuiltinCostExBudgetMemory                         int   `json:"cekBuiltinCost-exBudgetMemory"`
	CekStartupCostExBudgetMemory                         int   `json:"cekStartupCost-exBudgetMemory"`
	EqualsStringMemoryArguments                          int   `json:"equalsString-memory-arguments"`
	IndexByteStringCPUArguments                          int   `json:"indexByteString-cpu-arguments"`
	UnConstrDataMemoryArguments                          int   `json:"unConstrData-memory-arguments"`
	AddIntegerCPUArgumentsSlope                          int   `json:"addInteger-cpu-arguments-slope"`
	DecodeUtf8CPUArgumentsSlope                          int   `json:"decodeUtf8-cpu-arguments-slope"`
	EncodeUtf8CPUArgumentsSlope                          int   `json:"encodeUtf8-cpu-arguments-slope"`
	EqualsDataCPUArgumentsSlope                          int   `json:"equalsData-cpu-arguments-slope"`
	EqualsIntegerMemoryArguments                         int   `json:"equalsInteger-memory-arguments"`
	MkNilPairDataMemoryArguments                         int   `json:"mkNilPairData-memory-arguments"`
	Blake2B256CPUArgumentsSlope                          int   `json:"blake2b_256-cpu-arguments-slope"`
	AppendStringCPUArgumentsSlope                        int   `json:"appendString-cpu-arguments-slope"`
	EqualsStringCPUArgumentsSlope                        int   `json:"equalsString-cpu-arguments-slope"`
	IndexByteStringMemoryArguments                       int   `json:"indexByteString-memory-arguments"`
	LengthOfByteStringCPUArguments                       int   `json:"lengthOfByteString-cpu-arguments"`
	LessThanIntegerMemoryArguments                       int   `json:"lessThanInteger-memory-arguments"`
	Sha2256CPUArgumentsIntercept                         int   `json:"sha2_256-cpu-arguments-intercept"`
	Sha3256CPUArgumentsIntercept                         int   `json:"sha3_256-cpu-arguments-intercept"`
	AddIntegerMemoryArgumentsSlope                       int   `json:"addInteger-memory-arguments-slope"`
	DecodeUtf8MemoryArgumentsSlope                       int   `json:"decodeUtf8-memory-arguments-slope"`
	EncodeUtf8MemoryArgumentsSlope                       int   `json:"encodeUtf8-memory-arguments-slope"`
	EqualsByteStringMemoryArguments                      int   `json:"equalsByteString-memory-arguments"`
	EqualsIntegerCPUArgumentsSlope                       int   `json:"equalsInteger-cpu-arguments-slope"`
	ModIntegerCPUArgumentsConstant                       int   `json:"modInteger-cpu-arguments-constant"`
	ModIntegerMemoryArgumentsSlope                       int   `json:"modInteger-memory-arguments-slope"`
	SerialiseDataCPUArgumentsSlope                       int   `json:"serialiseData-cpu-arguments-slope"`
	AddIntegerCPUArgumentsIntercept                      int   `json:"addInteger-cpu-arguments-intercept"`
	ConsByteStringCPUArgumentsSlope                      int   `json:"consByteString-cpu-arguments-slope"`
	DecodeUtf8CPUArgumentsIntercept                      int   `json:"decodeUtf8-cpu-arguments-intercept"`
	EncodeUtf8CPUArgumentsIntercept                      int   `json:"encodeUtf8-cpu-arguments-intercept"`
	EqualsDataCPUArgumentsIntercept                      int   `json:"equalsData-cpu-arguments-intercept"`
	AppendStringMemoryArgumentsSlope                     int   `json:"appendString-memory-arguments-slope"`
	Blake2B256CPUArgumentsIntercept                      int   `json:"blake2b_256-cpu-arguments-intercept"`
	EqualsStringCPUArgumentsConstant                     int   `json:"equalsString-cpu-arguments-constant"`
	LengthOfByteStringMemoryArguments                    int   `json:"lengthOfByteString-memory-arguments"`
	LessThanByteStringMemoryArguments                    int   `json:"lessThanByteString-memory-arguments"`
	LessThanIntegerCPUArgumentsSlope                     int   `json:"lessThanInteger-cpu-arguments-slope"`
	ModIntegerMemoryArgumentsMinimum                     int   `json:"modInteger-memory-arguments-minimum"`
	MultiplyIntegerCPUArgumentsSlope                     int   `json:"multiplyInteger-cpu-arguments-slope"`
	SliceByteStringCPUArgumentsSlope                     int   `json:"sliceByteString-cpu-arguments-slope"`
	SubtractIntegerCPUArgumentsSlope                     int   `json:"subtractInteger-cpu-arguments-slope"`
	AppendByteStringCPUArgumentsSlope                    int   `json:"appendByteString-cpu-arguments-slope"`
	AppendStringCPUArgumentsIntercept                    int   `json:"appendString-cpu-arguments-intercept"`
	DivideIntegerCPUArgumentsConstant                    int   `json:"divideInteger-cpu-arguments-constant"`
	DivideIntegerMemoryArgumentsSlope                    int   `json:"divideInteger-memory-arguments-slope"`
	EqualsByteStringCPUArgumentsSlope                    int   `json:"equalsByteString-cpu-arguments-slope"`
	EqualsStringCPUArgumentsIntercept                    int   `json:"equalsString-cpu-arguments-intercept"`
	SerialiseDataMemoryArgumentsSlope                    int   `json:"serialiseData-memory-arguments-slope"`
	AddIntegerMemoryArgumentsIntercept                   int   `json:"addInteger-memory-arguments-intercept"`
	ConsByteStringMemoryArgumentsSlope                   int   `json:"consByteString-memory-arguments-slope"`
	DecodeUtf8MemoryArgumentsIntercept                   int   `json:"decodeUtf8-memory-arguments-intercept"`
	EncodeUtf8MemoryArgumentsIntercept                   int   `json:"encodeUtf8-memory-arguments-intercept"`
	EqualsIntegerCPUArgumentsIntercept                   int   `json:"equalsInteger-cpu-arguments-intercept"`
	ModIntegerMemoryArgumentsIntercept                   int   `json:"modInteger-memory-arguments-intercept"`
	SerialiseDataCPUArgumentsIntercept                   int   `json:"serialiseData-cpu-arguments-intercept"`
	ConsByteStringCPUArgumentsIntercept                  int   `json:"consByteString-cpu-arguments-intercept"`
	DivideIntegerMemoryArgumentsMinimum                  int   `json:"divideInteger-memory-arguments-minimum"`
	LessThanByteStringCPUArgumentsSlope                  int   `json:"lessThanByteString-cpu-arguments-slope"`
	LessThanEqualsIntegerMemoryArguments                 int   `json:"lessThanEqualsInteger-memory-arguments"`
	MultiplyIntegerMemoryArgumentsSlope                  int   `json:"multiplyInteger-memory-arguments-slope"`
	QuotientIntegerCPUArgumentsConstant                  int   `json:"quotientInteger-cpu-arguments-constant"`
	QuotientIntegerMemoryArgumentsSlope                  int   `json:"quotientInteger-memory-arguments-slope"`
	SliceByteStringMemoryArgumentsSlope                  int   `json:"sliceByteString-memory-arguments-slope"`
	SubtractIntegerMemoryArgumentsSlope                  int   `json:"subtractInteger-memory-arguments-slope"`
	AppendByteStringMemoryArgumentsSlope                 int   `json:"appendByteString-memory-arguments-slope"`
	AppendStringMemoryArgumentsIntercept                 int   `json:"appendString-memory-arguments-intercept"`
	EqualsByteStringCPUArgumentsConstant                 int   `json:"equalsByteString-cpu-arguments-constant"`
	LessThanIntegerCPUArgumentsIntercept                 int   `json:"lessThanInteger-cpu-arguments-intercept"`
	MultiplyIntegerCPUArgumentsIntercept                 int   `json:"multiplyInteger-cpu-arguments-intercept"`
	RemainderIntegerCPUArgumentsConstant                 int   `json:"remainderInteger-cpu-arguments-constant"`
	RemainderIntegerMemoryArgumentsSlope                 int   `json:"remainderInteger-memory-arguments-slope"`
	SliceByteStringCPUArgumentsIntercept                 int   `json:"sliceByteString-cpu-arguments-intercept"`
	SubtractIntegerCPUArgumentsIntercept                 int   `json:"subtractInteger-cpu-arguments-intercept"`
	VerifyEd25519SignatureMemoryArguments                int   `json:"verifyEd25519Signature-memory-arguments"`
	AppendByteStringCPUArgumentsIntercept                int   `json:"appendByteString-cpu-arguments-intercept"`
	DivideIntegerMemoryArgumentsIntercept                int   `json:"divideInteger-memory-arguments-intercept"`
	EqualsByteStringCPUArgumentsIntercept                int   `json:"equalsByteString-cpu-arguments-intercept"`
	QuotientIntegerMemoryArgumentsMinimum                int   `json:"quotientInteger-memory-arguments-minimum"`
	SerialiseDataMemoryArgumentsIntercept                int   `json:"serialiseData-memory-arguments-intercept"`
	ConsByteStringMemoryArgumentsIntercept               int   `json:"consByteString-memory-arguments-intercept"`
	LessThanEqualsByteStringMemoryArguments              int   `json:"lessThanEqualsByteString-memory-arguments"`
	LessThanEqualsIntegerCPUArgumentsSlope               int   `json:"lessThanEqualsInteger-cpu-arguments-slope"`
	RemainderIntegerMemoryArgumentsMinimum               int   `json:"remainderInteger-memory-arguments-minimum"`
	LessThanByteStringCPUArgumentsIntercept              int   `json:"lessThanByteString-cpu-arguments-intercept"`
	MultiplyIntegerMemoryArgumentsIntercept              int   `json:"multiplyInteger-memory-arguments-intercept"`
	QuotientIntegerMemoryArgumentsIntercept              int   `json:"quotientInteger-memory-arguments-intercept"`
	SliceByteStringMemoryArgumentsIntercept              int   `json:"sliceByteString-memory-arguments-intercept"`
	SubtractIntegerMemoryArgumentsIntercept              int   `json:"subtractInteger-memory-arguments-intercept"`
	VerifyEd25519SignatureCPUArgumentsSlope              int   `json:"verifyEd25519Signature-cpu-arguments-slope"`
	AppendByteStringMemoryArgumentsIntercept             int   `json:"appendByteString-memory-arguments-intercept"`
	RemainderIntegerMemoryArgumentsIntercept             int   `json:"remainderInteger-memory-arguments-intercept"`
	VerifyEcdsaSecp256K1SignatureCPUArguments            int64 `json:"verifyEcdsaSecp256k1Signature-cpu-arguments"`
	LessThanEqualsByteStringCPUArgumentsSlope            int   `json:"lessThanEqualsByteString-cpu-arguments-slope"`
	LessThanEqualsIntegerCPUArgumentsIntercept           int   `json:"lessThanEqualsInteger-cpu-arguments-intercept"`
	ModIntegerCPUArgumentsModelArgumentsSlope            int   `json:"modInteger-cpu-arguments-model-arguments-slope"`
	VerifyEcdsaSecp256K1SignatureMemoryArguments         int64 `json:"verifyEcdsaSecp256k1Signature-memory-arguments"`
	VerifyEd25519SignatureCPUArgumentsIntercept          int   `json:"verifyEd25519Signature-cpu-arguments-intercept"`
	LessThanEqualsByteStringCPUArgumentsIntercept        int   `json:"lessThanEqualsByteString-cpu-arguments-intercept"`
	VerifySchnorrSecp256K1SignatureMemoryArguments       int64 `json:"verifySchnorrSecp256k1Signature-memory-arguments"`
	DivideIntegerCPUArgumentsModelArgumentsSlope         int   `json:"divideInteger-cpu-arguments-model-arguments-slope"`
	ModIntegerCPUArgumentsModelArgumentsIntercept        int   `json:"modInteger-cpu-arguments-model-arguments-intercept"`
	QuotientIntegerCPUArgumentsModelArgumentsSlope       int   `json:"quotientInteger-cpu-arguments-model-arguments-slope"`
	VerifySchnorrSecp256K1SignatureCPUArgumentsSlope     int   `json:"verifySchnorrSecp256k1Signature-cpu-arguments-slope"`
	RemainderIntegerCPUArgumentsModelArgumentsSlope      int   `json:"remainderInteger-cpu-arguments-model-arguments-slope"`
	DivideIntegerCPUArgumentsModelArgumentsIntercept     int   `json:"divideInteger-cpu-arguments-model-arguments-intercept"`
	QuotientIntegerCPUArgumentsModelArgumentsIntercept   int   `json:"quotientInteger-cpu-arguments-model-arguments-intercept"`
	VerifySchnorrSecp256K1SignatureCPUArgumentsIntercept int64 `json:"verifySchnorrSecp256k1Signature-cpu-arguments-intercept"`
	RemainderIntegerCPUArgumentsModelArgumentsIntercept  int   `json:"remainderInteger-cpu-arguments-model-arguments-intercept"`
}

type Costs struct {
	PlutusV1 PlutusV1 `json:"PlutusV1"`
	PlutusV2 PlutusV2 `json:"PlutusV2"`
}

type CostModel struct {
	Hash  string `json:"hash"`
	Costs Costs  `json:"costs"`
}

type CurrentEpoch struct {
	OutSum    string    `json:"out_sum"`
	Fees      string    `json:"fees"`
	TxCount   int       `json:"tx_count"`
	BlkCount  int       `json:"blk_count"`
	No        int       `json:"no"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// ProtocolParameters Retrieves the protocol parameters for a given epoch.
func (c *apiClient) ProtocolParameters(ctx context.Context, epochNumber string) (eParams EpochParameters, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s/%s", c.server, c.appID, resourceEpochs, epochNumber, resourceParameters))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return
	}

	resp, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&eParams); err != nil {
		return
	}

	return eParams, nil
}

func (c *apiClient) CurrentEpoch(ctx context.Context) (cEpoch CurrentEpoch, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s", c.server, c.appID, resourceEpochs, resourceCurrent))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return
	}

	resp, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&cEpoch); err != nil {
		return
	}

	return cEpoch, nil
}

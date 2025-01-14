// Code generated by "stringer -type=ComputationKind -trimprefix=ComputationKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ComputationKindUnknown-0]
	_ = x[ComputationKindStatement-1001]
	_ = x[ComputationKindLoop-1002]
	_ = x[ComputationKindFunctionInvocation-1003]
	_ = x[ComputationKindCreateCompositeValue-1010]
	_ = x[ComputationKindTransferCompositeValue-1011]
	_ = x[ComputationKindDestroyCompositeValue-1012]
	_ = x[ComputationKindCreateArrayValue-1025]
	_ = x[ComputationKindTransferArrayValue-1026]
	_ = x[ComputationKindDestroyArrayValue-1027]
	_ = x[ComputationKindCreateDictionaryValue-1040]
	_ = x[ComputationKindTransferDictionaryValue-1041]
	_ = x[ComputationKindDestroyDictionaryValue-1042]
	_ = x[ComputationKindEncodeValue-1080]
	_ = x[ComputationKindSTDLIBPanic-1100]
	_ = x[ComputationKindSTDLIBAssert-1101]
	_ = x[ComputationKindSTDLIBUnsafeRandom-1102]
	_ = x[ComputationKindSTDLIBRLPDecodeString-1108]
	_ = x[ComputationKindSTDLIBRLPDecodeList-1109]
}

const (
	_ComputationKind_name_0 = "Unknown"
	_ComputationKind_name_1 = "StatementLoopFunctionInvocation"
	_ComputationKind_name_2 = "CreateCompositeValueTransferCompositeValueDestroyCompositeValue"
	_ComputationKind_name_3 = "CreateArrayValueTransferArrayValueDestroyArrayValue"
	_ComputationKind_name_4 = "CreateDictionaryValueTransferDictionaryValueDestroyDictionaryValue"
	_ComputationKind_name_5 = "EncodeValue"
	_ComputationKind_name_6 = "STDLIBPanicSTDLIBAssertSTDLIBUnsafeRandom"
	_ComputationKind_name_7 = "STDLIBRLPDecodeStringSTDLIBRLPDecodeList"
)

var (
	_ComputationKind_index_1 = [...]uint8{0, 9, 13, 31}
	_ComputationKind_index_2 = [...]uint8{0, 20, 42, 63}
	_ComputationKind_index_3 = [...]uint8{0, 16, 34, 51}
	_ComputationKind_index_4 = [...]uint8{0, 21, 44, 66}
	_ComputationKind_index_6 = [...]uint8{0, 11, 23, 41}
	_ComputationKind_index_7 = [...]uint8{0, 21, 40}
)

func (i ComputationKind) String() string {
	switch {
	case i == 0:
		return _ComputationKind_name_0
	case 1001 <= i && i <= 1003:
		i -= 1001
		return _ComputationKind_name_1[_ComputationKind_index_1[i]:_ComputationKind_index_1[i+1]]
	case 1010 <= i && i <= 1012:
		i -= 1010
		return _ComputationKind_name_2[_ComputationKind_index_2[i]:_ComputationKind_index_2[i+1]]
	case 1025 <= i && i <= 1027:
		i -= 1025
		return _ComputationKind_name_3[_ComputationKind_index_3[i]:_ComputationKind_index_3[i+1]]
	case 1040 <= i && i <= 1042:
		i -= 1040
		return _ComputationKind_name_4[_ComputationKind_index_4[i]:_ComputationKind_index_4[i+1]]
	case i == 1080:
		return _ComputationKind_name_5
	case 1100 <= i && i <= 1102:
		i -= 1100
		return _ComputationKind_name_6[_ComputationKind_index_6[i]:_ComputationKind_index_6[i+1]]
	case 1108 <= i && i <= 1109:
		i -= 1108
		return _ComputationKind_name_7[_ComputationKind_index_7[i]:_ComputationKind_index_7[i+1]]
	default:
		return "ComputationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

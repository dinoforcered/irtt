// Code generated by "stringer -type=Code"; DO NOT EDIT.

package irtt

import "fmt"

const (
	_Code_name_0 = "OpenTimeoutTooShortInvalidReceivedStatsStringInvalidReceivedStatsIntParamsChangedInvalidServerRestrictionOpenTimeoutServerClosedConnTokenZeroDurationNonPositiveIntervalNonPositiveNoSuchWaiterNoSuchTimerNoSuchFillerNoSuchAveragerInvalidWaitDurationInvalidWaitFactorInvalidWaitStringInvalidSleepFactorUnexpectedSequenceNumberClockMismatchStampAtMismatchShortReplyExpectedReplyFlagTTLErrorDFErrorUnexpectedOpenFlagAllocateResultsPanicInvalidExpAvgAlphaInvalidWinAvgWindow"
	_Code_name_1 = "AddressMismatchLargeRequestShortIntervalInvalidConnTokenNoSuitableAddressFoundUnexpectedReplyFlagInvalidGCModeStringUnspecifiedWithSpecifiedAddressesNoMatchingInterfacesUpNoMatchingInterfaces"
	_Code_name_2 = "InvalidParamValueUnknownParamParamOverflowShortParamBufferInvalidFlagBitsSetDFNotSupportedInconsistentClocksNonexclusiveMidpointTStampUnexpectedHMACBadHMACNoHMACBadMagicInvalidClockIntInvalidClockStringInvalidAllowStampStringInvalidStampAtIntInvalidStampAtStringFieldsCapacityTooLargeFieldsLengthTooLargeInvalidDFStringShortWrite"
	_Code_name_3 = "MultipleAddressesServerStartServerStopListenerStartListenerStopListenerErrorDropNewConnOpenCloseCloseConnNoDSCPSupportExceededDurationNoReceiveDstAddrSupportRemoveNoConn"
	_Code_name_4 = "ConnectingConnectedWaitForPacketsServerRestrictionNoTest"
)

var (
	_Code_index_0 = [...]uint16{0, 19, 45, 68, 81, 105, 116, 128, 141, 160, 179, 191, 202, 214, 228, 247, 264, 281, 299, 323, 336, 351, 361, 378, 386, 393, 411, 431, 449, 468}
	_Code_index_1 = [...]uint8{0, 15, 27, 40, 56, 78, 97, 116, 149, 171, 191}
	_Code_index_2 = [...]uint16{0, 17, 29, 42, 58, 76, 90, 108, 134, 148, 155, 161, 169, 184, 202, 225, 242, 262, 284, 304, 319, 329}
	_Code_index_3 = [...]uint8{0, 17, 28, 38, 51, 63, 76, 80, 87, 96, 105, 118, 134, 157, 169}
	_Code_index_4 = [...]uint8{0, 10, 19, 33, 50, 56}
)

func (i Code) String() string {
	switch {
	case -2076 <= i && i <= -2048:
		i -= -2076
		return _Code_name_0[_Code_index_0[i]:_Code_index_0[i+1]]
	case -1033 <= i && i <= -1024:
		i -= -1033
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	case -21 <= i && i <= -1:
		i -= -21
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	case 1024 <= i && i <= 1037:
		i -= 1024
		return _Code_name_3[_Code_index_3[i]:_Code_index_3[i+1]]
	case 2048 <= i && i <= 2052:
		i -= 2048
		return _Code_name_4[_Code_index_4[i]:_Code_index_4[i+1]]
	default:
		return fmt.Sprintf("Code(%d)", i)
	}
}

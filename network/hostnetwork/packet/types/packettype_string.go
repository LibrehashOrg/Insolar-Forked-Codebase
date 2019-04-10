// Code generated by "stringer -type=PacketType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ping-1]
	_ = x[RPC-2]
	_ = x[Cascade-3]
	_ = x[Pulse-4]
	_ = x[Bootstrap-5]
	_ = x[Authorize-6]
	_ = x[Register-7]
	_ = x[Genesis-8]
	_ = x[Challenge1-9]
	_ = x[Challenge2-10]
	_ = x[Disconnect-11]
}

const _PacketType_name = "PingRPCCascadePulseBootstrapAuthorizeRegisterGenesisChallenge1Challenge2Disconnect"

var _PacketType_index = [...]uint8{0, 4, 7, 14, 19, 28, 37, 45, 52, 62, 72, 82}

func (i PacketType) String() string {
	i -= 1
	if i < 0 || i >= PacketType(len(_PacketType_index)-1) {
		return "PacketType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _PacketType_name[_PacketType_index[i]:_PacketType_index[i+1]]
}
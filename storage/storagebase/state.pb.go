// Code generated by protoc-gen-gogo.
// source: cockroach/storage/storagebase/state.proto
// DO NOT EDIT!

/*
	Package storagebase is a generated protocol buffer package.

	It is generated from these files:
		cockroach/storage/storagebase/state.proto

	It has these top-level messages:
		ReplicaState
		RangeInfo
*/
package storagebase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_storage_engine_enginepb "github.com/cockroachdb/cockroach/storage/engine/enginepb"
import cockroach_roachpb2 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/util/hlc"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// ReplicaState is the part of the Range Raft state machine which is cached
// in memory and which is manipulated exclusively through consensus.
type ReplicaState struct {
	// The highest (and last) index applied to the state machine.
	RaftAppliedIndex uint64 `protobuf:"varint,1,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	// The highest (and last) lease index applied to the state machine.
	LeaseAppliedIndex uint64 `protobuf:"varint,2,opt,name=lease_applied_index,json=leaseAppliedIndex,proto3" json:"lease_applied_index,omitempty"`
	// The Range descriptor.
	// The pointer may change, but the referenced RangeDescriptor struct itself
	// must be treated as immutable; it is leaked out of the lock.
	//
	// Changes of the descriptor should normally go through one of the
	// (*Replica).setDesc* methods.
	Desc *cockroach_roachpb.RangeDescriptor `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	// The latest lease holder, if any.
	Lease *cockroach_roachpb1.Lease `protobuf:"bytes,4,opt,name=lease" json:"lease,omitempty"`
	// The truncation state of the Raft log.
	TruncatedState *cockroach_roachpb2.RaftTruncatedState `protobuf:"bytes,5,opt,name=truncated_state,json=truncatedState" json:"truncated_state,omitempty"`
	// gcThreshold is the GC threshold of the Range, typically updated when keys
	// are garbage collected. Reads and writes at timestamps <= this time will
	// not be served.
	GCThreshold cockroach_util_hlc.Timestamp                `protobuf:"bytes,6,opt,name=gc_threshold,json=gcThreshold" json:"gc_threshold"`
	Stats       cockroach_storage_engine_enginepb.MVCCStats `protobuf:"bytes,7,opt,name=stats" json:"stats"`
	Frozen      bool                                        `protobuf:"varint,8,opt,name=frozen,proto3" json:"frozen,omitempty"`
}

func (m *ReplicaState) Reset()                    { *m = ReplicaState{} }
func (m *ReplicaState) String() string            { return proto.CompactTextString(m) }
func (*ReplicaState) ProtoMessage()               {}
func (*ReplicaState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

type RangeInfo struct {
	ReplicaState `protobuf:"bytes,1,opt,name=state,embedded=state" json:"state"`
	// The highest (and last) index in the Raft log.
	LastIndex  uint64 `protobuf:"varint,2,opt,name=lastIndex,proto3" json:"lastIndex,omitempty"`
	NumPending uint64 `protobuf:"varint,3,opt,name=num_pending,json=numPending,proto3" json:"num_pending,omitempty"`
}

func (m *RangeInfo) Reset()                    { *m = RangeInfo{} }
func (m *RangeInfo) String() string            { return proto.CompactTextString(m) }
func (*RangeInfo) ProtoMessage()               {}
func (*RangeInfo) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{1} }

func init() {
	proto.RegisterType((*ReplicaState)(nil), "cockroach.storage.storagebase.ReplicaState")
	proto.RegisterType((*RangeInfo)(nil), "cockroach.storage.storagebase.RangeInfo")
}
func (m *ReplicaState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReplicaState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintState(data, i, uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintState(data, i, uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintState(data, i, uint64(m.Desc.Size()))
		n1, err := m.Desc.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Lease != nil {
		data[i] = 0x22
		i++
		i = encodeVarintState(data, i, uint64(m.Lease.Size()))
		n2, err := m.Lease.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.TruncatedState != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintState(data, i, uint64(m.TruncatedState.Size()))
		n3, err := m.TruncatedState.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	data[i] = 0x32
	i++
	i = encodeVarintState(data, i, uint64(m.GCThreshold.Size()))
	n4, err := m.GCThreshold.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	data[i] = 0x3a
	i++
	i = encodeVarintState(data, i, uint64(m.Stats.Size()))
	n5, err := m.Stats.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.Frozen {
		data[i] = 0x40
		i++
		if m.Frozen {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *RangeInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RangeInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintState(data, i, uint64(m.ReplicaState.Size()))
	n6, err := m.ReplicaState.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.LastIndex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintState(data, i, uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintState(data, i, uint64(m.NumPending))
	}
	return i, nil
}

func encodeFixed64State(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32State(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintState(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Lease != nil {
		l = m.Lease.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.TruncatedState != nil {
		l = m.TruncatedState.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = m.GCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovState(uint64(l))
	if m.Frozen {
		n += 2
	}
	return n
}

func (m *RangeInfo) Size() (n int) {
	var l int
	_ = l
	l = m.ReplicaState.Size()
	n += 1 + l + sovState(uint64(l))
	if m.LastIndex != 0 {
		n += 1 + sovState(uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		n += 1 + sovState(uint64(m.NumPending))
	}
	return n
}

func sovState(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplicaState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplicaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseAppliedIndex", wireType)
			}
			m.LeaseAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LeaseAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.Desc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lease == nil {
				m.Lease = &cockroach_roachpb1.Lease{}
			}
			if err := m.Lease.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TruncatedState == nil {
				m.TruncatedState = &cockroach_roachpb2.RaftTruncatedState{}
			}
			if err := m.TruncatedState.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GCThreshold.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Frozen = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipState(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RangeInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RangeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReplicaState.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LastIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPending", wireType)
			}
			m.NumPending = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NumPending |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipState(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthState
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowState
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipState(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthState = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorState = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0x9b, 0x67, 0x69, 0x9f, 0xce, 0x99, 0x78, 0xf1, 0x10, 0x8a, 0xa6, 0xad, 0x9d, 0x2a,
	0x55, 0x02, 0x51, 0x39, 0xd2, 0x90, 0xb8, 0xa7, 0x45, 0x82, 0x8a, 0x17, 0x21, 0x53, 0x10, 0xe2,
	0x26, 0x72, 0x1d, 0x37, 0x8d, 0x48, 0xec, 0x28, 0x71, 0x11, 0xe2, 0x53, 0x70, 0xc7, 0x57, 0xea,
	0xe5, 0x2e, 0xb9, 0xa1, 0x82, 0xf1, 0x45, 0xf0, 0x4b, 0xb2, 0x66, 0x74, 0xe2, 0xc2, 0x89, 0x7d,
	0xce, 0xcf, 0xff, 0x63, 0xff, 0x7d, 0xc0, 0x7d, 0x2a, 0xe8, 0xc7, 0x42, 0x10, 0xba, 0x0c, 0x4a,
	0x29, 0x0a, 0x12, 0xb3, 0xfa, 0x3f, 0x27, 0xa5, 0x9e, 0x13, 0xc9, 0x50, 0x5e, 0x08, 0x29, 0xe0,
	0xc9, 0x25, 0x8a, 0x2a, 0x04, 0x35, 0xd0, 0xa3, 0xd1, 0xae, 0x12, 0xe3, 0x71, 0xc2, 0xeb, 0x5f,
	0x3e, 0x0f, 0xb2, 0x4f, 0x94, 0x5a, 0xb1, 0xa3, 0xe1, 0x96, 0x36, 0x5f, 0x95, 0x4d, 0xb8, 0x64,
	0x05, 0x27, 0x69, 0x58, 0x90, 0x85, 0xac, 0xb0, 0xd3, 0x5d, 0x2c, 0x63, 0x92, 0x44, 0x44, 0x92,
	0x8a, 0x38, 0xde, 0x25, 0x1a, 0xd9, 0xc1, 0x36, 0xbb, 0x92, 0x49, 0x1a, 0x2c, 0x53, 0x1a, 0xc8,
	0x24, 0x63, 0xea, 0x5e, 0x59, 0x5e, 0x31, 0x77, 0x62, 0x11, 0x0b, 0x33, 0x0d, 0xf4, 0xcc, 0x46,
	0x07, 0x3f, 0xf6, 0xc0, 0x01, 0x66, 0x79, 0x9a, 0x50, 0xf2, 0x46, 0x9b, 0x00, 0x47, 0x00, 0xea,
	0x83, 0x85, 0x24, 0x57, 0x41, 0x16, 0x85, 0x09, 0x8f, 0xd8, 0x67, 0xdf, 0x39, 0x75, 0xee, 0xb9,
	0xf8, 0x96, 0xce, 0x3c, 0xb6, 0x89, 0xa9, 0x8e, 0x43, 0x04, 0x0e, 0x53, 0xa6, 0x6c, 0xf9, 0x0b,
	0xff, 0xcf, 0xe0, 0xb7, 0x4d, 0xea, 0x0a, 0xff, 0x08, 0xb8, 0x11, 0x2b, 0xa9, 0xbf, 0xa7, 0x00,
	0xef, 0x6c, 0x80, 0xb6, 0x5e, 0x57, 0xb7, 0x42, 0x98, 0xf0, 0x98, 0x3d, 0x51, 0x4c, 0x91, 0xe4,
	0xca, 0x5f, 0x6c, 0x78, 0x55, 0xa7, 0x6d, 0xc4, 0x7c, 0xd7, 0x6c, 0xf4, 0xaf, 0xd9, 0xf8, 0x42,
	0xe7, 0xb1, 0xc5, 0xe0, 0x2b, 0x70, 0x53, 0x16, 0x2b, 0x4e, 0xd5, 0x8d, 0xa2, 0xd0, 0xbc, 0xae,
	0xdf, 0x36, 0x3b, 0x87, 0xd7, 0x96, 0x5c, 0xc8, 0x59, 0x4d, 0x1b, 0x17, 0xf0, 0x0d, 0x79, 0x65,
	0x0d, 0xdf, 0x82, 0x83, 0x98, 0x86, 0x72, 0x59, 0xb0, 0x72, 0x29, 0xd2, 0xc8, 0xef, 0x18, 0xb1,
	0x93, 0x86, 0x98, 0xf6, 0x1d, 0x29, 0xdf, 0xd1, 0xac, 0xf6, 0x7d, 0x7c, 0xb8, 0xde, 0xf4, 0x5b,
	0x17, 0x9b, 0xbe, 0xf7, 0x74, 0x32, 0xab, 0x77, 0x62, 0x2f, 0xa6, 0x97, 0x0b, 0xf8, 0x0c, 0xb4,
	0xf5, 0xe1, 0x4a, 0xff, 0x7f, 0xa3, 0x37, 0x42, 0xbb, 0xbd, 0x67, 0xbb, 0x0a, 0xd5, 0xcd, 0x85,
	0x5e, 0xbe, 0x9b, 0x4c, 0xf4, 0x99, 0xca, 0xb1, 0xab, 0xe5, 0xb1, 0x15, 0x80, 0x77, 0x41, 0x67,
	0x51, 0x88, 0x2f, 0x8c, 0xfb, 0x5d, 0x25, 0xd5, 0xc5, 0xd5, 0x6a, 0xf0, 0xcd, 0x01, 0xfb, 0xc6,
	0xd2, 0x29, 0x5f, 0x08, 0xf8, 0xdc, 0xd6, 0x63, 0xe6, 0x3d, 0xbd, 0xb3, 0x07, 0xe8, 0x9f, 0xbd,
	0x8e, 0x9a, 0x8d, 0x31, 0xee, 0xea, 0x72, 0xe7, 0x9b, 0xbe, 0x63, 0x4b, 0x32, 0x78, 0x0c, 0xf6,
	0x53, 0x52, 0xca, 0x69, 0xe3, 0xc5, 0xb7, 0x01, 0xd8, 0x07, 0x1e, 0x5f, 0x65, 0x61, 0xce, 0x78,
	0x94, 0xf0, 0xd8, 0x3c, 0xb8, 0x8b, 0x81, 0x0a, 0xbd, 0xb6, 0x91, 0xf1, 0x70, 0xfd, 0xab, 0xd7,
	0x5a, 0x5f, 0xf4, 0x9c, 0x73, 0x35, 0xbe, 0xab, 0xf1, 0x53, 0x8d, 0xaf, 0xbf, 0x7b, 0xad, 0x0f,
	0x5e, 0xe3, 0x0c, 0xef, 0xdb, 0xf3, 0x8e, 0xe9, 0xd4, 0x87, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x21, 0x28, 0xa9, 0xc4, 0x03, 0x00, 0x00,
}

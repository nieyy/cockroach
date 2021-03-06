// Code generated by protoc-gen-gogo.
// source: cockroach/roachpb/internal_raft.proto
// DO NOT EDIT!

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import raftpb "github.com/coreos/etcd/raft/raftpb"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A RaftCommand is a command which can be serialized and sent via
// raft.
type RaftCommand struct {
	RangeID       RangeID           `protobuf:"varint,1,opt,name=range_id,json=rangeId,casttype=RangeID" json:"range_id"`
	OriginReplica ReplicaDescriptor `protobuf:"bytes,2,opt,name=origin_replica,json=originReplica" json:"origin_replica"`
	Cmd           BatchRequest      `protobuf:"bytes,3,opt,name=cmd" json:"cmd"`
}

func (m *RaftCommand) Reset()                    { *m = RaftCommand{} }
func (m *RaftCommand) String() string            { return proto.CompactTextString(m) }
func (*RaftCommand) ProtoMessage()               {}
func (*RaftCommand) Descriptor() ([]byte, []int) { return fileDescriptorInternalRaft, []int{0} }

// RaftTruncatedState contains metadata about the truncated portion of the raft log.
// Raft requires access to the term of the last truncated log entry even after the
// rest of the entry has been discarded.
type RaftTruncatedState struct {
	// The highest index that has been removed from the log.
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index"`
	// The term corresponding to 'index'.
	Term uint64 `protobuf:"varint,2,opt,name=term" json:"term"`
}

func (m *RaftTruncatedState) Reset()                    { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string            { return proto.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()               {}
func (*RaftTruncatedState) Descriptor() ([]byte, []int) { return fileDescriptorInternalRaft, []int{1} }

// RaftTombstone contains information about a replica that has been deleted.
type RaftTombstone struct {
	NextReplicaID ReplicaID `protobuf:"varint,1,opt,name=next_replica_id,json=nextReplicaId,casttype=ReplicaID" json:"next_replica_id"`
}

func (m *RaftTombstone) Reset()                    { *m = RaftTombstone{} }
func (m *RaftTombstone) String() string            { return proto.CompactTextString(m) }
func (*RaftTombstone) ProtoMessage()               {}
func (*RaftTombstone) Descriptor() ([]byte, []int) { return fileDescriptorInternalRaft, []int{2} }

// RaftSnapshotData is the payload of a raftpb.Snapshot. It contains a raw copy of
// all of the range's data and metadata, including the raft log, sequence cache, etc.
type RaftSnapshotData struct {
	// The latest RangeDescriptor
	RangeDescriptor RangeDescriptor             `protobuf:"bytes,1,opt,name=range_descriptor,json=rangeDescriptor" json:"range_descriptor"`
	KV              []RaftSnapshotData_KeyValue `protobuf:"bytes,2,rep,name=KV,json=kV" json:"KV"`
	LogEntries      []raftpb.Entry              `protobuf:"bytes,3,rep,name=log_entries,json=logEntries" json:"log_entries"`
}

func (m *RaftSnapshotData) Reset()                    { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string            { return proto.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()               {}
func (*RaftSnapshotData) Descriptor() ([]byte, []int) { return fileDescriptorInternalRaft, []int{3} }

type RaftSnapshotData_KeyValue struct {
	Key       []byte    `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value     []byte    `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Timestamp Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp"`
}

func (m *RaftSnapshotData_KeyValue) Reset()         { *m = RaftSnapshotData_KeyValue{} }
func (m *RaftSnapshotData_KeyValue) String() string { return proto.CompactTextString(m) }
func (*RaftSnapshotData_KeyValue) ProtoMessage()    {}
func (*RaftSnapshotData_KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptorInternalRaft, []int{3, 0}
}

func init() {
	proto.RegisterType((*RaftCommand)(nil), "cockroach.roachpb.RaftCommand")
	proto.RegisterType((*RaftTruncatedState)(nil), "cockroach.roachpb.RaftTruncatedState")
	proto.RegisterType((*RaftTombstone)(nil), "cockroach.roachpb.RaftTombstone")
	proto.RegisterType((*RaftSnapshotData)(nil), "cockroach.roachpb.RaftSnapshotData")
	proto.RegisterType((*RaftSnapshotData_KeyValue)(nil), "cockroach.roachpb.RaftSnapshotData.KeyValue")
}
func (m *RaftCommand) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftCommand) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.RangeID))
	data[i] = 0x12
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.OriginReplica.Size()))
	n1, err := m.OriginReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.Cmd.Size()))
	n2, err := m.Cmd.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *RaftTruncatedState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftTruncatedState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.Index))
	data[i] = 0x10
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.Term))
	return i, nil
}

func (m *RaftTombstone) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftTombstone) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.NextReplicaID))
	return i, nil
}

func (m *RaftSnapshotData) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftSnapshotData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.RangeDescriptor.Size()))
	n3, err := m.RangeDescriptor.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.KV) > 0 {
		for _, msg := range m.KV {
			data[i] = 0x12
			i++
			i = encodeVarintInternalRaft(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.LogEntries) > 0 {
		for _, msg := range m.LogEntries {
			data[i] = 0x1a
			i++
			i = encodeVarintInternalRaft(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RaftSnapshotData_KeyValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftSnapshotData_KeyValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		data[i] = 0xa
		i++
		i = encodeVarintInternalRaft(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if m.Value != nil {
		data[i] = 0x12
		i++
		i = encodeVarintInternalRaft(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintInternalRaft(data, i, uint64(m.Timestamp.Size()))
	n4, err := m.Timestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64InternalRaft(data []byte, offset int, v uint64) int {
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
func encodeFixed32InternalRaft(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInternalRaft(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RaftCommand) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternalRaft(uint64(m.RangeID))
	l = m.OriginReplica.Size()
	n += 1 + l + sovInternalRaft(uint64(l))
	l = m.Cmd.Size()
	n += 1 + l + sovInternalRaft(uint64(l))
	return n
}

func (m *RaftTruncatedState) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternalRaft(uint64(m.Index))
	n += 1 + sovInternalRaft(uint64(m.Term))
	return n
}

func (m *RaftTombstone) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternalRaft(uint64(m.NextReplicaID))
	return n
}

func (m *RaftSnapshotData) Size() (n int) {
	var l int
	_ = l
	l = m.RangeDescriptor.Size()
	n += 1 + l + sovInternalRaft(uint64(l))
	if len(m.KV) > 0 {
		for _, e := range m.KV {
			l = e.Size()
			n += 1 + l + sovInternalRaft(uint64(l))
		}
	}
	if len(m.LogEntries) > 0 {
		for _, e := range m.LogEntries {
			l = e.Size()
			n += 1 + l + sovInternalRaft(uint64(l))
		}
	}
	return n
}

func (m *RaftSnapshotData_KeyValue) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(m.Key)
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovInternalRaft(uint64(l))
	return n
}

func sovInternalRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternalRaft(x uint64) (n int) {
	return sovInternalRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftCommand) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftCommand: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftCommand: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeID |= (RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OriginReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cmd.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftTruncatedState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftTruncatedState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftTruncatedState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftTombstone) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftTombstone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftTombstone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextReplicaID", wireType)
			}
			m.NextReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextReplicaID |= (ReplicaID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftSnapshotData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftSnapshotData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftSnapshotData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeDescriptor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RangeDescriptor.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KV", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KV = append(m.KV, RaftSnapshotData_KeyValue{})
			if err := m.KV[len(m.KV)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogEntries = append(m.LogEntries, raftpb.Entry{})
			if err := m.LogEntries[len(m.LogEntries)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftSnapshotData_KeyValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: KeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], data[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], data[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func skipInternalRaft(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternalRaft
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
					return 0, ErrIntOverflowInternalRaft
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
					return 0, ErrIntOverflowInternalRaft
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
				return 0, ErrInvalidLengthInternalRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternalRaft
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
				next, err := skipInternalRaft(data[start:])
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
	ErrInvalidLengthInternalRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternalRaft   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorInternalRaft = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xcd, 0xa7, 0xd2, 0x4e, 0x1a, 0x1a, 0xac, 0x1e, 0xa2, 0x6d, 0x95, 0x94, 0x15, 0x48, 0x1c,
	0xd0, 0x46, 0xaa, 0x40, 0x5c, 0xd1, 0x92, 0x4a, 0xb4, 0x95, 0x90, 0xd8, 0x54, 0x11, 0x82, 0x43,
	0xe4, 0xae, 0x4d, 0x62, 0x25, 0x6b, 0x2f, 0x8e, 0x83, 0x9a, 0x7f, 0xc1, 0xcf, 0xca, 0xb1, 0x47,
	0x2e, 0x54, 0x50, 0xfe, 0x05, 0xa7, 0xda, 0x5e, 0x27, 0x69, 0xb5, 0x39, 0xd8, 0x9a, 0x9d, 0xf7,
	0xe6, 0x79, 0xe6, 0xcd, 0xc2, 0x8b, 0x58, 0xc4, 0x13, 0x29, 0x70, 0x3c, 0xee, 0xda, 0x3b, 0xbd,
	0xea, 0x32, 0xae, 0xa8, 0xe4, 0x78, 0x3a, 0x94, 0xf8, 0x9b, 0x0a, 0x52, 0x29, 0x94, 0x40, 0x4f,
	0xd7, 0xb4, 0xc0, 0xd1, 0xbc, 0xc3, 0x7c, 0x25, 0x4e, 0x59, 0xc6, 0xf7, 0x8e, 0xf2, 0x20, 0xc1,
	0x0a, 0x3b, 0xf4, 0x38, 0x8f, 0x26, 0x54, 0xe1, 0x07, 0x8c, 0x43, 0xaa, 0x62, 0xd2, 0x35, 0x0d,
	0xd8, 0x4b, 0x13, 0x36, 0xcd, 0x78, 0x07, 0x23, 0x31, 0x12, 0x36, 0xec, 0x9a, 0x28, 0xcb, 0xfa,
	0x37, 0x45, 0xa8, 0x47, 0x9a, 0xf4, 0x5e, 0x24, 0x09, 0xe6, 0x04, 0xbd, 0x81, 0x1d, 0x89, 0xf9,
	0x88, 0x0e, 0x19, 0x69, 0x15, 0x8f, 0x8b, 0x2f, 0xcb, 0xa1, 0xb7, 0xbc, 0xed, 0x14, 0xee, 0x6e,
	0x3b, 0xb5, 0xc8, 0xe4, 0xcf, 0x7a, 0xff, 0x37, 0x61, 0x54, 0xb3, 0xdc, 0x33, 0x82, 0x3e, 0xc1,
	0x13, 0x21, 0xd9, 0x88, 0xf1, 0xa1, 0xa4, 0xe9, 0x94, 0xc5, 0xb8, 0x55, 0xd2, 0xc5, 0xf5, 0x93,
	0xe7, 0x41, 0xce, 0x82, 0x20, 0xca, 0x18, 0x3d, 0x3a, 0x8b, 0x25, 0x4b, 0x95, 0x90, 0x61, 0xc5,
	0x3c, 0x11, 0x35, 0x32, 0x05, 0x07, 0xa3, 0xb7, 0x50, 0x8e, 0x13, 0xd2, 0x2a, 0x5b, 0x9d, 0xce,
	0x16, 0x9d, 0x10, 0xab, 0x78, 0x1c, 0xd1, 0xef, 0x73, 0x3a, 0x53, 0x4e, 0xc2, 0x54, 0xf8, 0xe7,
	0x80, 0xcc, 0x44, 0x97, 0x72, 0xce, 0x63, 0xac, 0x28, 0xe9, 0x2b, 0x7d, 0x23, 0x0f, 0xaa, 0x8c,
	0x13, 0x7a, 0x6d, 0xa7, 0xaa, 0x38, 0x7e, 0x96, 0x42, 0x2d, 0xa8, 0xe8, 0xe5, 0x25, 0xb6, 0xe7,
	0x15, 0x64, 0x33, 0xfe, 0x57, 0x68, 0x58, 0x2d, 0x91, 0x5c, 0xcd, 0x94, 0xe0, 0x14, 0x9d, 0xc3,
	0x3e, 0xa7, 0xd7, 0x6a, 0x35, 0xe6, 0xca, 0xa6, 0x6a, 0xe8, 0x3b, 0x9b, 0x1a, 0x1f, 0x35, 0xec,
	0x66, 0xb0, 0x66, 0xed, 0xae, 0x3f, 0xa2, 0x06, 0x7f, 0x80, 0x11, 0xff, 0x77, 0x09, 0x9a, 0x46,
	0xbd, 0xcf, 0x71, 0x3a, 0x1b, 0x0b, 0xd5, 0xd3, 0x9b, 0x44, 0x7d, 0x68, 0x66, 0x0b, 0x20, 0x6b,
	0x7f, 0xec, 0x0b, 0xf5, 0x13, 0x7f, 0x9b, 0x97, 0x86, 0x9a, 0x73, 0x72, 0x5f, 0x3e, 0x4e, 0xa3,
	0x0f, 0x50, 0xba, 0x18, 0xe8, 0xf1, 0xca, 0x5a, 0xe6, 0xd5, 0x56, 0x99, 0xc7, 0x5d, 0x04, 0x17,
	0x74, 0x31, 0xc0, 0xd3, 0x39, 0x0d, 0xc1, 0x8d, 0xa5, 0xeb, 0xa3, 0xd2, 0x64, 0x80, 0x5e, 0x43,
	0x7d, 0x2a, 0x46, 0x43, 0xca, 0x95, 0x64, 0x74, 0xa6, 0xb7, 0x63, 0x24, 0x1b, 0x41, 0xf6, 0xbb,
	0x05, 0xa7, 0x3a, 0xbd, 0x70, 0x4d, 0x80, 0xe6, 0x9d, 0x66, 0x34, 0x4f, 0xc1, 0xce, 0x4a, 0x11,
	0x35, 0xa1, 0x3c, 0xa1, 0x0b, 0x3b, 0xd3, 0x5e, 0x64, 0x42, 0x74, 0x00, 0xd5, 0x1f, 0x06, 0xb2,
	0xfe, 0xef, 0x45, 0xd9, 0x07, 0x7a, 0x07, 0xbb, 0x8a, 0x25, 0x7a, 0xb7, 0x38, 0x49, 0xdd, 0x5f,
	0x70, 0xb4, 0xa5, 0xf5, 0xcb, 0x15, 0xc7, 0x3d, 0xbb, 0x29, 0x0a, 0x9f, 0x2d, 0xff, 0xb6, 0x0b,
	0xcb, 0xbb, 0x76, 0xf1, 0x46, 0x9f, 0x5f, 0xfa, 0xfc, 0xd1, 0xe7, 0xe7, 0xbf, 0x76, 0xe1, 0x4b,
	0xcd, 0x55, 0x7f, 0xae, 0xdc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x35, 0x71, 0xb0, 0xfa, 0xcb, 0x03,
	0x00, 0x00,
}

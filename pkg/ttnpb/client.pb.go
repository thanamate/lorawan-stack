// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/client.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import strconv "strconv"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// ClientState enum defines all the possible tenant admin reviewing states
// that a third-party client request can be at.
type ClientState int32

const (
	// State that denotes that the client request is pending to review by the
	// tenant admin.
	STATE_PENDING ClientState = 0
	// Denotes that the client request has been approved by the tenant admin
	// and therefore the client can be used.
	STATE_APPROVED ClientState = 1
	// Denotes that the client request has beenr rejected by the tenant admin
	// and therefore it cannot be used.
	STATE_REJECTED ClientState = 2
)

var ClientState_name = map[int32]string{
	0: "STATE_PENDING",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
}
var ClientState_value = map[string]int32{
	"STATE_PENDING":  0,
	"STATE_APPROVED": 1,
	"STATE_REJECTED": 2,
}

func (ClientState) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

// Scope enum defines the different scopes a third-party client can have access to.
type ClientScope int32

const (
	// Denotes whether if the client has access to manage user's applications.
	SCOPE_APPLICATION ClientScope = 0
	// Denotes wheter if the client has r-w access to user's profile.
	SCOPE_PROFILE ClientScope = 1
)

var ClientScope_name = map[int32]string{
	0: "SCOPE_APPLICATION",
	1: "SCOPE_PROFILE",
}
var ClientScope_value = map[string]int32{
	"SCOPE_APPLICATION": 0,
	"SCOPE_PROFILE":     1,
}

func (ClientScope) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{1} }

// Grant enum defines the OAuth2 flows a third-party client can use to get
// access to a token.
type ClientGrant int32

const (
	// Grant type used to exchange an authorization code for an access token.
	GRANT_AUTHORIZATION_CODE ClientGrant = 0
	// Grant type used to exchange an username and password for an access token.
	GRANT_PASSWORD ClientGrant = 1
	// Grant type used to exchange a refresh token for an access token.
	GRANT_REFRESH_TOKEN ClientGrant = 2
)

var ClientGrant_name = map[int32]string{
	0: "GRANT_AUTHORIZATION_CODE",
	1: "GRANT_PASSWORD",
	2: "GRANT_REFRESH_TOKEN",
}
var ClientGrant_value = map[string]int32{
	"GRANT_AUTHORIZATION_CODE": 0,
	"GRANT_PASSWORD":           1,
	"GRANT_REFRESH_TOKEN":      2,
}

func (ClientGrant) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{2} }

// Client is the message that defines a third-party client on the network.
type Client struct {
	ClientIdentifiers `protobuf:"bytes,1,opt,name=client,embedded=client" json:"client"`
	// description is the description of the client.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// secret is the secret used to prove the client identity.
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	// callback_uri is the callback URI of the client.
	CallbackURI string `protobuf:"bytes,4,opt,name=callback_uri,json=callbackUri,proto3" json:"callback_uri,omitempty" db:"callback_uri"`
	// state denotes the reviewing state of the client by the tenant admin.
	State ClientState `protobuf:"varint,5,opt,name=state,proto3,enum=ttn.v3.ClientState" json:"state,omitempty"`
	// official denotes that a client has been labelled as an official third-party
	// client by the tenant admins and therefore this fact will be shown to
	Official bool `protobuf:"varint,6,opt,name=official,proto3" json:"official,omitempty"`
	// grants denotes which OAuth2 flows can the client use to get a token.
	Grants []ClientGrant `protobuf:"varint,7,rep,packed,name=grants,enum=ttn.v3.ClientGrant" json:"grants,omitempty"`
	// scope denotes what scopes the client will have access to.
	Scope []ClientScope `protobuf:"varint,8,rep,packed,name=scope,enum=ttn.v3.ClientScope" json:"scope,omitempty"`
	// created_at denotes when the client was created.
	CreatedAt time.Time `protobuf:"bytes,9,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	// archived_at denotes when the client was disabled.
	ArchivedAt *time.Time `protobuf:"bytes,10,opt,name=archived_at,json=archivedAt,stdtime" json:"archived_at,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

func (m *Client) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Client) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Client) GetCallbackURI() string {
	if m != nil {
		return m.CallbackURI
	}
	return ""
}

func (m *Client) GetState() ClientState {
	if m != nil {
		return m.State
	}
	return STATE_PENDING
}

func (m *Client) GetOfficial() bool {
	if m != nil {
		return m.Official
	}
	return false
}

func (m *Client) GetGrants() []ClientGrant {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *Client) GetScope() []ClientScope {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *Client) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Client) GetArchivedAt() *time.Time {
	if m != nil {
		return m.ArchivedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "ttn.v3.Client")
	proto.RegisterEnum("ttn.v3.ClientState", ClientState_name, ClientState_value)
	proto.RegisterEnum("ttn.v3.ClientScope", ClientScope_name, ClientScope_value)
	proto.RegisterEnum("ttn.v3.ClientGrant", ClientGrant_name, ClientGrant_value)
}
func (x ClientState) String() string {
	s, ok := ClientState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ClientScope) String() string {
	s, ok := ClientScope_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ClientGrant) String() string {
	s, ok := ClientGrant_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Client) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Client) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintClient(dAtA, i, uint64(m.ClientIdentifiers.Size()))
	n1, err := m.ClientIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Secret) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Secret)))
		i += copy(dAtA[i:], m.Secret)
	}
	if len(m.CallbackURI) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.CallbackURI)))
		i += copy(dAtA[i:], m.CallbackURI)
	}
	if m.State != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.State))
	}
	if m.Official {
		dAtA[i] = 0x30
		i++
		if m.Official {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Grants) > 0 {
		dAtA3 := make([]byte, len(m.Grants)*10)
		var j2 int
		for _, num := range m.Grants {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintClient(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.Scope) > 0 {
		dAtA5 := make([]byte, len(m.Scope)*10)
		var j4 int
		for _, num := range m.Scope {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x42
		i++
		i = encodeVarintClient(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintClient(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n6, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.ArchivedAt != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintClient(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.ArchivedAt)))
		n7, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ArchivedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeFixed64Client(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Client(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Client) Size() (n int) {
	var l int
	_ = l
	l = m.ClientIdentifiers.Size()
	n += 1 + l + sovClient(uint64(l))
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.CallbackURI)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovClient(uint64(m.State))
	}
	if m.Official {
		n += 2
	}
	if len(m.Grants) > 0 {
		l = 0
		for _, e := range m.Grants {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	if len(m.Scope) > 0 {
		l = 0
		for _, e := range m.Scope {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovClient(uint64(l))
	if m.ArchivedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ArchivedAt)
		n += 1 + l + sovClient(uint64(l))
	}
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Client) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Client{`,
		`ClientIdentifiers:` + strings.Replace(strings.Replace(this.ClientIdentifiers.String(), "ClientIdentifiers", "ClientIdentifiers", 1), `&`, ``, 1) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Secret:` + fmt.Sprintf("%v", this.Secret) + `,`,
		`CallbackURI:` + fmt.Sprintf("%v", this.CallbackURI) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`Official:` + fmt.Sprintf("%v", this.Official) + `,`,
		`Grants:` + fmt.Sprintf("%v", this.Grants) + `,`,
		`Scope:` + fmt.Sprintf("%v", this.Scope) + `,`,
		`CreatedAt:` + strings.Replace(strings.Replace(this.CreatedAt.String(), "Timestamp", "google_protobuf1.Timestamp", 1), `&`, ``, 1) + `,`,
		`ArchivedAt:` + strings.Replace(fmt.Sprintf("%v", this.ArchivedAt), "Timestamp", "google_protobuf1.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClient(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Client) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Client: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Client: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ClientState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Official", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Official = bool(v != 0)
		case 7:
			if wireType == 0 {
				var v ClientGrant
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (ClientGrant(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Grants = append(m.Grants, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v ClientGrant
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (ClientGrant(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Grants = append(m.Grants, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
			}
		case 8:
			if wireType == 0 {
				var v ClientScope
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (ClientScope(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Scope = append(m.Scope, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v ClientScope
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (ClientScope(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Scope = append(m.Scope, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArchivedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ArchivedAt == nil {
				m.ArchivedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ArchivedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowClient
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowClient
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClient
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipClient(dAtA[start:])
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
	ErrInvalidLengthClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/client.proto", fileDescriptorClient)
}

var fileDescriptorClient = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3d, 0x6f, 0xd3, 0x4c,
	0x1c, 0xcf, 0xf5, 0xc5, 0x4f, 0x7a, 0x79, 0xa8, 0xd2, 0xab, 0x00, 0x13, 0x21, 0x27, 0xaa, 0x18,
	0xd2, 0x22, 0x1c, 0xd4, 0x8a, 0x05, 0xc4, 0xe0, 0x3a, 0x6e, 0x6a, 0xa8, 0xec, 0xe8, 0xec, 0x82,
	0xd4, 0x81, 0xc8, 0x71, 0x2e, 0xce, 0xa9, 0xa9, 0x6d, 0xd9, 0xd7, 0xb2, 0x32, 0x32, 0xf2, 0x1d,
	0xf8, 0x32, 0x1d, 0x3b, 0x32, 0x05, 0xf0, 0xc4, 0x06, 0xe2, 0x13, 0x20, 0xdf, 0xd9, 0x4d, 0x10,
	0x48, 0x65, 0xb2, 0xff, 0xbf, 0xb7, 0xfb, 0xf9, 0x7f, 0x32, 0x7c, 0x1c, 0x50, 0x36, 0x39, 0x1f,
	0xaa, 0x7e, 0x74, 0xd6, 0x71, 0x27, 0xc4, 0x9d, 0xd0, 0x30, 0x48, 0x2d, 0xc2, 0xde, 0x46, 0xc9,
	0x69, 0x87, 0xb1, 0xb0, 0xe3, 0xc5, 0xb4, 0xe3, 0x4f, 0x29, 0x09, 0x99, 0x1a, 0x27, 0x11, 0x8b,
	0x90, 0xc4, 0x58, 0xa8, 0x5e, 0xec, 0x35, 0x1e, 0x2d, 0x38, 0x83, 0x28, 0x88, 0x3a, 0x9c, 0x1e,
	0x9e, 0x8f, 0xf9, 0xc4, 0x07, 0xfe, 0x26, 0x6c, 0x8d, 0x27, 0xff, 0x72, 0x10, 0x1d, 0x91, 0x90,
	0xd1, 0x31, 0x25, 0x49, 0x5a, 0xd8, 0x9a, 0x41, 0x14, 0x05, 0x53, 0x32, 0x0f, 0x67, 0xf4, 0x8c,
	0xa4, 0xcc, 0x3b, 0x8b, 0x85, 0x60, 0xeb, 0xfb, 0x32, 0x94, 0x74, 0xde, 0x0f, 0x3d, 0x83, 0x92,
	0x68, 0x2a, 0x83, 0x16, 0x68, 0xd7, 0x76, 0xef, 0xa9, 0xa2, 0xaa, 0x2a, 0x78, 0x73, 0x1e, 0xbe,
	0x5f, 0xbd, 0x9c, 0x35, 0x2b, 0x57, 0xb3, 0x26, 0xc0, 0x85, 0x05, 0xb5, 0x60, 0x6d, 0x44, 0x52,
	0x3f, 0xa1, 0x31, 0xa3, 0x51, 0x28, 0x2f, 0xb5, 0x40, 0x7b, 0x0d, 0x2f, 0x42, 0xe8, 0x0e, 0x94,
	0x52, 0xe2, 0x27, 0x84, 0xc9, 0xcb, 0x9c, 0x2c, 0x26, 0xd4, 0x83, 0xff, 0xfb, 0xde, 0x74, 0x3a,
	0xf4, 0xfc, 0xd3, 0xc1, 0x79, 0x42, 0xe5, 0x95, 0x9c, 0xdd, 0x7f, 0x90, 0xcd, 0x9a, 0x35, 0xbd,
	0xc0, 0x8f, 0xb1, 0xf9, 0x73, 0xd6, 0xdc, 0x18, 0x0d, 0x9f, 0x6e, 0x2d, 0x4a, 0xb7, 0x70, 0xad,
	0x1c, 0x8f, 0x13, 0x8a, 0xb6, 0xe1, 0x6a, 0xca, 0x3c, 0x46, 0xe4, 0xd5, 0x16, 0x68, 0xaf, 0xef,
	0x6e, 0xfe, 0x5e, 0xdf, 0xc9, 0x29, 0x2c, 0x14, 0xa8, 0x01, 0xab, 0xd1, 0x78, 0x4c, 0x7d, 0xea,
	0x4d, 0x65, 0xa9, 0x05, 0xda, 0x55, 0x7c, 0x3d, 0xa3, 0x87, 0x50, 0x0a, 0x12, 0x2f, 0x64, 0xa9,
	0xfc, 0x5f, 0x6b, 0xf9, 0xcf, 0x9c, 0x5e, 0xce, 0xe1, 0x42, 0xc2, 0xcf, 0xf4, 0xa3, 0x98, 0xc8,
	0xd5, 0xbf, 0x69, 0x9d, 0x9c, 0xc2, 0x42, 0x81, 0x74, 0x08, 0xfd, 0x84, 0x78, 0x8c, 0x8c, 0x06,
	0x1e, 0x93, 0xd7, 0xf8, 0x8a, 0x1b, 0xaa, 0xb8, 0x1f, 0xb5, 0xbc, 0x1f, 0xd5, 0x2d, 0xef, 0x47,
	0xec, 0xf8, 0xc3, 0xe7, 0x26, 0xc0, 0x6b, 0x85, 0x4f, 0x63, 0x48, 0x83, 0x35, 0x2f, 0xf1, 0x27,
	0xf4, 0x42, 0xa4, 0xc0, 0x1b, 0x53, 0x56, 0x78, 0x02, 0x2c, 0x4d, 0x1a, 0xdb, 0xb1, 0x60, 0x6d,
	0x61, 0x23, 0x68, 0x03, 0xde, 0x72, 0x5c, 0xcd, 0x35, 0x06, 0x7d, 0xc3, 0xea, 0x9a, 0x56, 0xaf,
	0x5e, 0x41, 0x08, 0xae, 0x0b, 0x48, 0xeb, 0xf7, 0xb1, 0xfd, 0xca, 0xe8, 0xd6, 0xc1, 0x1c, 0xc3,
	0xc6, 0x0b, 0x43, 0x77, 0x8d, 0x6e, 0x7d, 0xa9, 0xb1, 0xf2, 0xfe, 0xa3, 0x52, 0xd9, 0x79, 0x7e,
	0x9d, 0xc7, 0x3f, 0xf3, 0x36, 0xdc, 0x70, 0x74, 0xbb, 0xcf, 0xcd, 0x47, 0xa6, 0xae, 0xb9, 0xa6,
	0x6d, 0xd5, 0x2b, 0xfc, 0x18, 0x0e, 0xf7, 0xb1, 0x7d, 0x60, 0x1e, 0x19, 0x75, 0x50, 0xd8, 0xdf,
	0x94, 0x76, 0xbe, 0x58, 0x74, 0x1f, 0xca, 0x3d, 0xac, 0x59, 0xee, 0x40, 0x3b, 0x76, 0x0f, 0x6d,
	0x6c, 0x9e, 0xf0, 0x80, 0x81, 0x6e, 0x77, 0x0d, 0xd1, 0x4c, 0xb0, 0x7d, 0xcd, 0x71, 0x5e, 0xdb,
	0x38, 0x6f, 0x76, 0x17, 0x6e, 0x0a, 0x0c, 0x1b, 0x07, 0xd8, 0x70, 0x0e, 0x07, 0xae, 0xfd, 0xd2,
	0xb0, 0xca, 0x7a, 0xfb, 0xbd, 0x4f, 0x5f, 0x95, 0xca, 0xbb, 0x4c, 0x01, 0x97, 0x99, 0x02, 0xae,
	0x32, 0x05, 0x7c, 0xc9, 0x14, 0xf0, 0x2d, 0x53, 0x2a, 0x3f, 0x32, 0x05, 0x9c, 0x6c, 0xdf, 0xf4,
	0x5b, 0xc5, 0xa7, 0x41, 0xfe, 0x8c, 0x87, 0x43, 0x89, 0x6f, 0x77, 0xef, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x9d, 0x80, 0x72, 0xf3, 0x03, 0x00, 0x00,
}

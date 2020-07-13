// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/ca.proto

package v1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Certificate request message.
// The authentication should be based on
// 1. bearer tokens carried in the side channel;
// 2. Client-side certificate via Mutual TLS handshake.
// Note: the server side may overwrite any requested certificate field based on its policies.
type CertificateRequest struct {
	// PEM-encoded certificate request.
	// The public key in the CSR is used to generate the certificate,
	// and other fields in the generated certificate may be overwritten by the CA.
	Csr string `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Optional: requested certificate validity period.
	Validity             *types.Duration `protobuf:"bytes,2,opt,name=validity,proto3" json:"validity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd8562bb175eac3, []int{0}
}
func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return m.Size()
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCsr() string {
	if m != nil {
		return m.Csr
	}
	return ""
}

func (m *CertificateRequest) GetValidity() *types.Duration {
	if m != nil {
		return m.Validity
	}
	return nil
}

// Certificate response message.
type CertificateResponse struct {
	// PEM-encoded certificate chain.
	// The leaf cert is the first element, and the root cert is the last element.
	CertificateChain     []string `protobuf:"bytes,1,rep,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd8562bb175eac3, []int{1}
}
func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificateChain() []string {
	if m != nil {
		return m.CertificateChain
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "istio.security.v1beta1.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "istio.security.v1beta1.CertificateResponse")
}

func init() { proto.RegisterFile("security/v1beta1/ca.proto", fileDescriptor_bdd8562bb175eac3) }

var fileDescriptor_bdd8562bb175eac3 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x8d, 0x03, 0x71, 0xf1, 0x66, 0x8b, 0x20, 0xdd, 0xc0, 0x52, 0x7a, 0x55, 0x2c, 0x24,
	0x6c, 0xe2, 0x0b, 0xac, 0x3e, 0x41, 0xbd, 0x13, 0x44, 0xd2, 0xec, 0x6c, 0x1e, 0x18, 0x4d, 0x4d,
	0xd2, 0xc2, 0xee, 0x7d, 0x38, 0x2f, 0x7d, 0x04, 0xe9, 0x93, 0x48, 0x97, 0xcd, 0x39, 0xf5, 0x62,
	0x77, 0xe1, 0x9c, 0xef, 0x7c, 0x9c, 0xfc, 0x87, 0x8e, 0x2c, 0xa8, 0xda, 0xa0, 0x5b, 0x8b, 0x66,
	0x52, 0x80, 0x93, 0x13, 0xa1, 0x24, 0xaf, 0x8c, 0x76, 0x9a, 0x5d, 0xa1, 0x75, 0xa8, 0xf9, 0x0e,
	0xe0, 0x5b, 0x60, 0x1c, 0x2e, 0xb5, 0x5e, 0xae, 0x40, 0x6c, 0xa8, 0xa2, 0x5e, 0x88, 0x79, 0x6d,
	0xa4, 0x43, 0x5d, 0xfa, 0xb9, 0xf8, 0x89, 0xb2, 0x0c, 0x8c, 0xc3, 0x05, 0x2a, 0xe9, 0x20, 0x87,
	0xd7, 0x1a, 0xac, 0x63, 0x03, 0xda, 0x53, 0xd6, 0x04, 0x24, 0x22, 0x49, 0x3f, 0xef, 0x9e, 0xec,
	0x8e, 0x9e, 0x37, 0x72, 0x85, 0x73, 0x74, 0xeb, 0xe0, 0x34, 0x22, 0xc9, 0xc5, 0x74, 0xc4, 0xbd,
	0x9a, 0xef, 0xd4, 0xfc, 0x7e, 0xab, 0xce, 0xbf, 0xd1, 0x78, 0x46, 0x2f, 0x0f, 0xf4, 0xb6, 0xd2,
	0xa5, 0x05, 0x96, 0xd2, 0xa1, 0xda, 0x97, 0x9f, 0xd5, 0x8b, 0xc4, 0x32, 0x20, 0x51, 0x2f, 0xe9,
	0xe7, 0x83, 0x1f, 0x8d, 0xac, 0xab, 0x4f, 0xdf, 0xc8, 0xc1, 0x8e, 0x0f, 0x60, 0x1a, 0x54, 0xc0,
	0x4a, 0x3a, 0xcc, 0x0c, 0x74, 0xd4, 0xbe, 0xc7, 0x6e, 0xf8, 0xff, 0x39, 0xf0, 0xbf, 0x9f, 0x1c,
	0xa7, 0x47, 0xb1, 0x7e, 0xe3, 0xf8, 0x64, 0x96, 0xbe, 0xb7, 0x21, 0xf9, 0x68, 0x43, 0xf2, 0xd9,
	0x86, 0xe4, 0xf1, 0xda, 0xcf, 0xa2, 0x16, 0xb2, 0x42, 0xf1, 0xfb, 0x2e, 0xc5, 0xd9, 0x26, 0x94,
	0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x53, 0x01, 0xa5, 0xb2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateServiceClient interface {
	// Returns a signed certificate or error based on the CertificateRequest, CA configurations and policies.
	CreateCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateServiceClient(cc *grpc.ClientConn) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) CreateCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/istio.security.v1beta1.CertificateService/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
type CertificateServiceServer interface {
	// Returns a signed certificate or error based on the CertificateRequest, CA configurations and policies.
	CreateCertificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (*UnimplementedCertificateServiceServer) CreateCertificate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}

func RegisterCertificateServiceServer(s *grpc.Server, srv CertificateServiceServer) {
	s.RegisterService(&_CertificateService_serviceDesc, srv)
}

func _CertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.security.v1beta1.CertificateService/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.security.v1beta1.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificateService_CreateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/v1beta1/ca.proto",
}

func (m *CertificateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Validity != nil {
		{
			size, err := m.Validity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCa(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Csr) > 0 {
		i -= len(m.Csr)
		copy(dAtA[i:], m.Csr)
		i = encodeVarintCa(dAtA, i, uint64(len(m.Csr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CertificateChain) > 0 {
		for iNdEx := len(m.CertificateChain) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CertificateChain[iNdEx])
			copy(dAtA[i:], m.CertificateChain[iNdEx])
			i = encodeVarintCa(dAtA, i, uint64(len(m.CertificateChain[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCa(dAtA []byte, offset int, v uint64) int {
	offset -= sovCa(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CertificateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Csr)
	if l > 0 {
		n += 1 + l + sovCa(uint64(l))
	}
	if m.Validity != nil {
		l = m.Validity.Size()
		n += 1 + l + sovCa(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CertificateChain) > 0 {
		for _, s := range m.CertificateChain {
			l = len(s)
			n += 1 + l + sovCa(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCa(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCa(x uint64) (n int) {
	return sovCa(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CertificateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCa
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCa
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validity == nil {
				m.Validity = &types.Duration{}
			}
			if err := m.Validity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCa
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCa
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificateChain = append(m.CertificateChain, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCa
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCa(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCa
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
					return 0, ErrIntOverflowCa
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
					return 0, ErrIntOverflowCa
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
			if length < 0 {
				return 0, ErrInvalidLengthCa
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCa
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCa
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
				next, err := skipCa(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCa
				}
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
	ErrInvalidLengthCa = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCa   = fmt.Errorf("proto: integer overflow")
)

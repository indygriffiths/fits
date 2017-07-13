// Code generated by protoc-gen-go.
// source: fits.proto
// DO NOT EDIT!

/*
Package fits is a generated protocol buffer package.

It is generated from these files:
	fits.proto

It has these top-level messages:
	Site
	Observation
	Response
	SiteID
	ObservationRequest
	ObservationResult
	DeleteObservationsRequest
*/
package fits

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Site struct {
	// siteID - a short code uniquely identifying Site e.g., GISB. Required.
	SiteID string `protobuf:"bytes,1,opt,name=siteID" json:"siteID,omitempty"`
	// name of the site. Required.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// latitude (WGS84) of the site. Required.
	Latitude float64 `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	// longitude (WGS84) of the site.  Required.
	Longitude float64 `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	// height of the site in meters.  Required.
	Height float64 `protobuf:"fixed64,5,opt,name=height" json:"height,omitempty"`
	// groundRelationship - the distance from the site to the ground in meters.
	// e.g., a site above ground has a negative ground relationship.  Required.
	GroundRelationship float64 `protobuf:"fixed64,6,opt,name=ground_relationship,json=groundRelationship" json:"ground_relationship,omitempty"`
}

func (m *Site) Reset()                    { *m = Site{} }
func (m *Site) String() string            { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()               {}
func (*Site) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Site) GetSiteID() string {
	if m != nil {
		return m.SiteID
	}
	return ""
}

func (m *Site) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Site) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Site) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Site) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Site) GetGroundRelationship() float64 {
	if m != nil {
		return m.GroundRelationship
	}
	return 0
}

// Observation is uniquely identified by (siteID, typeID, methodID, sampleID, seconds, nanoSeconds).
type Observation struct {
	// siteID for the observation e.g., GISB.  Required.
	SiteID string `protobuf:"bytes,1,opt,name=siteID" json:"siteID,omitempty"`
	// typeID for the observation e.g., e.  Required.
	TypeID string `protobuf:"bytes,2,opt,name=typeID" json:"typeID,omitempty"`
	// methodID for the observation e.g., bernese.  Required.
	MethodID string `protobuf:"bytes,3,opt,name=methodID" json:"methodID,omitempty"`
	// sampleID for the observation.  Defaults to "none",
	SampleID string `protobuf:"bytes,4,opt,name=sampleID" json:"sampleID,omitempty"`
	// systemID for the sampleID for the observation.  Defaults to "none",
	SystemID string `protobuf:"bytes,5,opt,name=systemID" json:"systemID,omitempty"`
	// seconds - Unix time of the observation in seconds.
	Seconds int64 `protobuf:"varint,6,opt,name=seconds" json:"seconds,omitempty"`
	// nanoSeconds - fractional part of observation time in nanoseconds.
	NanoSeconds int64 `protobuf:"varint,7,opt,name=nano_seconds,json=nanoSeconds" json:"nano_seconds,omitempty"`
	// observation value.
	Value float64 `protobuf:"fixed64,8,opt,name=value" json:"value,omitempty"`
	// error of the observation value.
	Error float64 `protobuf:"fixed64,9,opt,name=error" json:"error,omitempty"`
}

func (m *Observation) Reset()                    { *m = Observation{} }
func (m *Observation) String() string            { return proto.CompactTextString(m) }
func (*Observation) ProtoMessage()               {}
func (*Observation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Observation) GetSiteID() string {
	if m != nil {
		return m.SiteID
	}
	return ""
}

func (m *Observation) GetTypeID() string {
	if m != nil {
		return m.TypeID
	}
	return ""
}

func (m *Observation) GetMethodID() string {
	if m != nil {
		return m.MethodID
	}
	return ""
}

func (m *Observation) GetSampleID() string {
	if m != nil {
		return m.SampleID
	}
	return ""
}

func (m *Observation) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *Observation) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Observation) GetNanoSeconds() int64 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

func (m *Observation) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Observation) GetError() float64 {
	if m != nil {
		return m.Error
	}
	return 0
}

type Response struct {
	// affected - the number of records affected.
	Affected int64 `protobuf:"varint,1,opt,name=affected" json:"affected,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetAffected() int64 {
	if m != nil {
		return m.Affected
	}
	return 0
}

type SiteID struct {
	// siteID - a short code uniquely identifying Site e.g., GISB.
	SiteID string `protobuf:"bytes,1,opt,name=siteID" json:"siteID,omitempty"`
}

func (m *SiteID) Reset()                    { *m = SiteID{} }
func (m *SiteID) String() string            { return proto.CompactTextString(m) }
func (*SiteID) ProtoMessage()               {}
func (*SiteID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SiteID) GetSiteID() string {
	if m != nil {
		return m.SiteID
	}
	return ""
}

type ObservationRequest struct {
	// siteID for the observation e.g., GISB.  Required.
	SiteID string `protobuf:"bytes,1,opt,name=siteID" json:"siteID,omitempty"`
	// typeID for the observation e.g., e.  Required.
	TypeID string `protobuf:"bytes,2,opt,name=typeID" json:"typeID,omitempty"`
}

func (m *ObservationRequest) Reset()                    { *m = ObservationRequest{} }
func (m *ObservationRequest) String() string            { return proto.CompactTextString(m) }
func (*ObservationRequest) ProtoMessage()               {}
func (*ObservationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ObservationRequest) GetSiteID() string {
	if m != nil {
		return m.SiteID
	}
	return ""
}

func (m *ObservationRequest) GetTypeID() string {
	if m != nil {
		return m.TypeID
	}
	return ""
}

type ObservationResult struct {
	// seconds - Unix time of the observation in seconds.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// nanoSeconds - fractional part of observation time in nanoseconds.
	NanoSeconds int64 `protobuf:"varint,2,opt,name=nano_seconds,json=nanoSeconds" json:"nano_seconds,omitempty"`
	// observation value.
	Value float64 `protobuf:"fixed64,3,opt,name=value" json:"value,omitempty"`
	// error of the observation value.
	Error float64 `protobuf:"fixed64,4,opt,name=error" json:"error,omitempty"`
}

func (m *ObservationResult) Reset()                    { *m = ObservationResult{} }
func (m *ObservationResult) String() string            { return proto.CompactTextString(m) }
func (*ObservationResult) ProtoMessage()               {}
func (*ObservationResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ObservationResult) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *ObservationResult) GetNanoSeconds() int64 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

func (m *ObservationResult) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ObservationResult) GetError() float64 {
	if m != nil {
		return m.Error
	}
	return 0
}

type DeleteObservationsRequest struct {
	// siteID for the observation e.g., GISB.  Required.
	SiteID string `protobuf:"bytes,1,opt,name=siteID" json:"siteID,omitempty"`
	// typeID for the observation e.g., e.  Required.
	TypeID string `protobuf:"bytes,2,opt,name=typeID" json:"typeID,omitempty"`
}

func (m *DeleteObservationsRequest) Reset()                    { *m = DeleteObservationsRequest{} }
func (m *DeleteObservationsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteObservationsRequest) ProtoMessage()               {}
func (*DeleteObservationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteObservationsRequest) GetSiteID() string {
	if m != nil {
		return m.SiteID
	}
	return ""
}

func (m *DeleteObservationsRequest) GetTypeID() string {
	if m != nil {
		return m.TypeID
	}
	return ""
}

func init() {
	proto.RegisterType((*Site)(nil), "fits.Site")
	proto.RegisterType((*Observation)(nil), "fits.Observation")
	proto.RegisterType((*Response)(nil), "fits.Response")
	proto.RegisterType((*SiteID)(nil), "fits.SiteID")
	proto.RegisterType((*ObservationRequest)(nil), "fits.ObservationRequest")
	proto.RegisterType((*ObservationResult)(nil), "fits.ObservationResult")
	proto.RegisterType((*DeleteObservationsRequest)(nil), "fits.DeleteObservationsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Fits service

type FitsClient interface {
	// GetSite returns the Site referred to by SiteID.
	GetSite(ctx context.Context, in *SiteID, opts ...grpc.CallOption) (*Site, error)
	//
	GetObservations(ctx context.Context, in *ObservationRequest, opts ...grpc.CallOption) (Fits_GetObservationsClient, error)
	// SaveSite saves or updates Site.
	SaveSite(ctx context.Context, in *Site, opts ...grpc.CallOption) (*Response, error)
	// DeleteSite deletes the site referred to by SiteID.
	// Any observations for the site are also deleted.
	DeleteSite(ctx context.Context, in *SiteID, opts ...grpc.CallOption) (*Response, error)
	// SaveObservations saves or updates the Observation.
	SaveObservation(ctx context.Context, in *Observation, opts ...grpc.CallOption) (*Response, error)
	// SaveObservations saves or updates the stream of Observation.
	SaveObservations(ctx context.Context, opts ...grpc.CallOption) (Fits_SaveObservationsClient, error)
	// DeleteObservations deletes all the Observations for the request.
	DeleteObservations(ctx context.Context, in *DeleteObservationsRequest, opts ...grpc.CallOption) (*Response, error)
}

type fitsClient struct {
	cc *grpc.ClientConn
}

func NewFitsClient(cc *grpc.ClientConn) FitsClient {
	return &fitsClient{cc}
}

func (c *fitsClient) GetSite(ctx context.Context, in *SiteID, opts ...grpc.CallOption) (*Site, error) {
	out := new(Site)
	err := grpc.Invoke(ctx, "/fits.Fits/GetSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fitsClient) GetObservations(ctx context.Context, in *ObservationRequest, opts ...grpc.CallOption) (Fits_GetObservationsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Fits_serviceDesc.Streams[0], c.cc, "/fits.Fits/GetObservations", opts...)
	if err != nil {
		return nil, err
	}
	x := &fitsGetObservationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fits_GetObservationsClient interface {
	Recv() (*ObservationResult, error)
	grpc.ClientStream
}

type fitsGetObservationsClient struct {
	grpc.ClientStream
}

func (x *fitsGetObservationsClient) Recv() (*ObservationResult, error) {
	m := new(ObservationResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fitsClient) SaveSite(ctx context.Context, in *Site, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/fits.Fits/SaveSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fitsClient) DeleteSite(ctx context.Context, in *SiteID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/fits.Fits/DeleteSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fitsClient) SaveObservation(ctx context.Context, in *Observation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/fits.Fits/SaveObservation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fitsClient) SaveObservations(ctx context.Context, opts ...grpc.CallOption) (Fits_SaveObservationsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Fits_serviceDesc.Streams[1], c.cc, "/fits.Fits/SaveObservations", opts...)
	if err != nil {
		return nil, err
	}
	x := &fitsSaveObservationsClient{stream}
	return x, nil
}

type Fits_SaveObservationsClient interface {
	Send(*Observation) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type fitsSaveObservationsClient struct {
	grpc.ClientStream
}

func (x *fitsSaveObservationsClient) Send(m *Observation) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fitsSaveObservationsClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fitsClient) DeleteObservations(ctx context.Context, in *DeleteObservationsRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/fits.Fits/DeleteObservations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fits service

type FitsServer interface {
	// GetSite returns the Site referred to by SiteID.
	GetSite(context.Context, *SiteID) (*Site, error)
	//
	GetObservations(*ObservationRequest, Fits_GetObservationsServer) error
	// SaveSite saves or updates Site.
	SaveSite(context.Context, *Site) (*Response, error)
	// DeleteSite deletes the site referred to by SiteID.
	// Any observations for the site are also deleted.
	DeleteSite(context.Context, *SiteID) (*Response, error)
	// SaveObservations saves or updates the Observation.
	SaveObservation(context.Context, *Observation) (*Response, error)
	// SaveObservations saves or updates the stream of Observation.
	SaveObservations(Fits_SaveObservationsServer) error
	// DeleteObservations deletes all the Observations for the request.
	DeleteObservations(context.Context, *DeleteObservationsRequest) (*Response, error)
}

func RegisterFitsServer(s *grpc.Server, srv FitsServer) {
	s.RegisterService(&_Fits_serviceDesc, srv)
}

func _Fits_GetSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FitsServer).GetSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fits.Fits/GetSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FitsServer).GetSite(ctx, req.(*SiteID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fits_GetObservations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObservationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FitsServer).GetObservations(m, &fitsGetObservationsServer{stream})
}

type Fits_GetObservationsServer interface {
	Send(*ObservationResult) error
	grpc.ServerStream
}

type fitsGetObservationsServer struct {
	grpc.ServerStream
}

func (x *fitsGetObservationsServer) Send(m *ObservationResult) error {
	return x.ServerStream.SendMsg(m)
}

func _Fits_SaveSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Site)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FitsServer).SaveSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fits.Fits/SaveSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FitsServer).SaveSite(ctx, req.(*Site))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fits_DeleteSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FitsServer).DeleteSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fits.Fits/DeleteSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FitsServer).DeleteSite(ctx, req.(*SiteID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fits_SaveObservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Observation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FitsServer).SaveObservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fits.Fits/SaveObservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FitsServer).SaveObservation(ctx, req.(*Observation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fits_SaveObservations_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FitsServer).SaveObservations(&fitsSaveObservationsServer{stream})
}

type Fits_SaveObservationsServer interface {
	SendAndClose(*Response) error
	Recv() (*Observation, error)
	grpc.ServerStream
}

type fitsSaveObservationsServer struct {
	grpc.ServerStream
}

func (x *fitsSaveObservationsServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fitsSaveObservationsServer) Recv() (*Observation, error) {
	m := new(Observation)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fits_DeleteObservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FitsServer).DeleteObservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fits.Fits/DeleteObservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FitsServer).DeleteObservations(ctx, req.(*DeleteObservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fits_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fits.Fits",
	HandlerType: (*FitsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSite",
			Handler:    _Fits_GetSite_Handler,
		},
		{
			MethodName: "SaveSite",
			Handler:    _Fits_SaveSite_Handler,
		},
		{
			MethodName: "DeleteSite",
			Handler:    _Fits_DeleteSite_Handler,
		},
		{
			MethodName: "SaveObservation",
			Handler:    _Fits_SaveObservation_Handler,
		},
		{
			MethodName: "DeleteObservations",
			Handler:    _Fits_DeleteObservations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetObservations",
			Handler:       _Fits_GetObservations_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SaveObservations",
			Handler:       _Fits_SaveObservations_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "fits.proto",
}

func init() { proto.RegisterFile("fits.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0x66, 0x93, 0xeb, 0x25, 0x99, 0x16, 0x6b, 0x47, 0xa9, 0x6b, 0x10, 0x8c, 0x27, 0x94, 0xe0,
	0x43, 0x15, 0x15, 0x7c, 0x97, 0xb3, 0x25, 0xf8, 0x20, 0x5c, 0xde, 0x7c, 0x29, 0xd7, 0xde, 0x24,
	0x39, 0xb8, 0xdc, 0x9e, 0xb7, 0x7b, 0x81, 0x82, 0x7f, 0xcb, 0x3f, 0x27, 0x08, 0xb2, 0xb3, 0x97,
	0xeb, 0x99, 0xa4, 0x41, 0x7d, 0xdb, 0x6f, 0xbe, 0x99, 0xd9, 0xf9, 0xbe, 0x61, 0x17, 0x60, 0x96,
	0x1a, 0x7d, 0x5e, 0x94, 0xca, 0x28, 0xf4, 0xec, 0x39, 0xf8, 0x21, 0xc0, 0x9b, 0xa6, 0x86, 0xf0,
	0x14, 0x7c, 0x9d, 0x1a, 0x9a, 0x84, 0x52, 0x8c, 0xc4, 0x78, 0x10, 0xd5, 0x08, 0x11, 0xbc, 0x3c,
	0x5e, 0x92, 0xec, 0x70, 0x94, 0xcf, 0x38, 0x84, 0x7e, 0x16, 0x9b, 0xd4, 0x54, 0x09, 0xc9, 0xee,
	0x48, 0x8c, 0x45, 0xd4, 0x60, 0x7c, 0x06, 0x83, 0x4c, 0xe5, 0x73, 0x47, 0x7a, 0x4c, 0xde, 0x05,
	0xec, 0x2d, 0x0b, 0x4a, 0xe7, 0x0b, 0x23, 0x0f, 0x98, 0xaa, 0x11, 0xbe, 0x86, 0x47, 0xf3, 0x52,
	0x55, 0x79, 0x72, 0x55, 0x92, 0x6d, 0xa5, 0x72, 0xbd, 0x48, 0x0b, 0xe9, 0x73, 0x12, 0x3a, 0x2a,
	0x6a, 0x31, 0xc1, 0x4f, 0x01, 0x87, 0x5f, 0xae, 0x35, 0x95, 0x2b, 0x8e, 0xdd, 0x3b, 0xfe, 0x29,
	0xf8, 0xe6, 0xb6, 0xb0, 0x71, 0x27, 0xa0, 0x46, 0x56, 0xc2, 0x92, 0xcc, 0x42, 0x25, 0x93, 0x90,
	0x25, 0x0c, 0xa2, 0x06, 0x5b, 0x4e, 0xc7, 0xcb, 0x22, 0xb3, 0x55, 0x9e, 0xe3, 0xd6, 0x98, 0xb9,
	0x5b, 0x6d, 0x68, 0x39, 0x09, 0x59, 0x82, 0xe5, 0x6a, 0x8c, 0x12, 0x7a, 0x9a, 0x6e, 0x54, 0x9e,
	0x68, 0x1e, 0xbc, 0x1b, 0xad, 0x21, 0xbe, 0x80, 0xa3, 0x3c, 0xce, 0xd5, 0xd5, 0x9a, 0xee, 0x31,
	0x7d, 0x68, 0x63, 0xd3, 0x3a, 0xe5, 0x31, 0x1c, 0xac, 0xe2, 0xac, 0x22, 0xd9, 0x67, 0xcd, 0x0e,
	0xd8, 0x28, 0x95, 0xa5, 0x2a, 0xe5, 0xc0, 0x45, 0x19, 0x04, 0x67, 0xd0, 0x8f, 0x48, 0x17, 0x2a,
	0xd7, 0xbc, 0x8b, 0x78, 0x36, 0xa3, 0x1b, 0x43, 0x09, 0x4b, 0xef, 0x46, 0x0d, 0x0e, 0x46, 0xe0,
	0x4f, 0x1b, 0x1b, 0x76, 0xd9, 0x13, 0x84, 0x80, 0x2d, 0x17, 0x23, 0xfa, 0x56, 0x91, 0x36, 0xff,
	0x6a, 0x66, 0xf0, 0x1d, 0x4e, 0xfe, 0xe8, 0xa2, 0xab, 0xcc, 0xb4, 0xdd, 0x10, 0xfb, 0xdd, 0xe8,
	0xec, 0x71, 0xa3, 0xbb, 0xd3, 0x0d, 0xaf, 0xed, 0xc6, 0x67, 0x78, 0x1a, 0x52, 0x46, 0x86, 0x5a,
	0x33, 0xe8, 0xff, 0x94, 0xf2, 0xf6, 0x57, 0x07, 0xbc, 0x8b, 0xd4, 0x68, 0x7c, 0x09, 0xbd, 0x4b,
	0x32, 0xfc, 0x34, 0x8e, 0xce, 0xf9, 0xd9, 0x38, 0x2b, 0x87, 0x70, 0x87, 0xf0, 0x02, 0x8e, 0x2f,
	0xc9, 0xb4, 0xef, 0x45, 0xe9, 0xe8, 0x6d, 0x57, 0x87, 0x4f, 0x76, 0x30, 0xd6, 0xa9, 0x37, 0x02,
	0xcf, 0xa0, 0x3f, 0x8d, 0x57, 0xc4, 0x3d, 0x5b, 0xfd, 0x87, 0x0f, 0xdc, 0xb9, 0x59, 0xf6, 0x2b,
	0x00, 0x27, 0x75, 0xc7, 0x5c, 0x9b, 0xb9, 0xef, 0xe1, 0xd8, 0xf6, 0x6c, 0x3f, 0x92, 0x93, 0xad,
	0x09, 0xb6, 0xaa, 0x3e, 0xc0, 0xc3, 0x8d, 0x2a, 0xfd, 0x17, 0x65, 0x63, 0x81, 0x9f, 0x00, 0xb7,
	0xb7, 0x80, 0xcf, 0x5d, 0xde, 0xbd, 0xfb, 0xd9, 0x6c, 0xf4, 0xd1, 0xff, 0xca, 0xff, 0xd2, 0xb5,
	0xcf, 0x9f, 0xd4, 0xbb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xd0, 0x83, 0x58, 0xb2, 0x04,
	0x00, 0x00,
}

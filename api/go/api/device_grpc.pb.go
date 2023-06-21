// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/device.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// Create the given device.
	Create(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get returns the device for the given DevEUI.
	Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	// Update the given device.
	Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the device with the given DevEUI.
	Delete(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the list of devices.
	List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	// Create the given device-keys.
	CreateKeys(ctx context.Context, in *CreateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the device-keys for the given DevEUI.
	GetKeys(ctx context.Context, in *GetDeviceKeysRequest, opts ...grpc.CallOption) (*GetDeviceKeysResponse, error)
	// Update the given device-keys.
	UpdateKeys(ctx context.Context, in *UpdateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the device-keys for the given DevEUI.
	DeleteKeys(ctx context.Context, in *DeleteDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// FlushDevNonces flushes the OTAA device nonces.
	FlushDevNonces(ctx context.Context, in *FlushDevNoncesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Activate (re)activates the device with the given parameters (for ABP or for
	// importing OTAA activations).
	Activate(ctx context.Context, in *ActivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deactivate de-activates the device.
	Deactivate(ctx context.Context, in *DeactivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetActivation returns the current activation details of the device (OTAA or
	// ABP).
	GetActivation(ctx context.Context, in *GetDeviceActivationRequest, opts ...grpc.CallOption) (*GetDeviceActivationResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into
	// account.
	GetRandomDevAddr(ctx context.Context, in *GetRandomDevAddrRequest, opts ...grpc.CallOption) (*GetRandomDevAddrResponse, error)
	// GetMetrics returns the device metrics.
	// Note that this requires a device-profile with codec and measurements
	// configured.
	GetMetrics(ctx context.Context, in *GetDeviceMetricsRequest, opts ...grpc.CallOption) (*GetDeviceMetricsResponse, error)
	// GetLinkMetrics returns the device link metrics.
	// This includes uplinks, downlinks, RSSI, SNR, etc...
	GetLinkMetrics(ctx context.Context, in *GetDeviceLinkMetricsRequest, opts ...grpc.CallOption) (*GetDeviceLinkMetricsResponse, error)
	// Enqueue adds the given item to the downlink queue.
	Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error)
	// FlushQueue flushes the downlink device-queue.
	FlushQueue(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetQueue returns the downlink device-queue.
	GetQueue(ctx context.Context, in *GetDeviceQueueItemsRequest, opts ...grpc.CallOption) (*GetDeviceQueueItemsResponse, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) Create(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Delete(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	out := new(ListDevicesResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) CreateKeys(ctx context.Context, in *CreateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/CreateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetKeys(ctx context.Context, in *GetDeviceKeysRequest, opts ...grpc.CallOption) (*GetDeviceKeysResponse, error) {
	out := new(GetDeviceKeysResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) UpdateKeys(ctx context.Context, in *UpdateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/UpdateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) DeleteKeys(ctx context.Context, in *DeleteDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/DeleteKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) FlushDevNonces(ctx context.Context, in *FlushDevNoncesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/FlushDevNonces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Activate(ctx context.Context, in *ActivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Deactivate(ctx context.Context, in *DeactivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetActivation(ctx context.Context, in *GetDeviceActivationRequest, opts ...grpc.CallOption) (*GetDeviceActivationResponse, error) {
	out := new(GetDeviceActivationResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetActivation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetRandomDevAddr(ctx context.Context, in *GetRandomDevAddrRequest, opts ...grpc.CallOption) (*GetRandomDevAddrResponse, error) {
	out := new(GetRandomDevAddrResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetRandomDevAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetMetrics(ctx context.Context, in *GetDeviceMetricsRequest, opts ...grpc.CallOption) (*GetDeviceMetricsResponse, error) {
	out := new(GetDeviceMetricsResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetLinkMetrics(ctx context.Context, in *GetDeviceLinkMetricsRequest, opts ...grpc.CallOption) (*GetDeviceLinkMetricsResponse, error) {
	out := new(GetDeviceLinkMetricsResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetLinkMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error) {
	out := new(EnqueueDeviceQueueItemResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) FlushQueue(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/FlushQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetQueue(ctx context.Context, in *GetDeviceQueueItemsRequest, opts ...grpc.CallOption) (*GetDeviceQueueItemsResponse, error) {
	out := new(GetDeviceQueueItemsResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	// Create the given device.
	Create(context.Context, *CreateDeviceRequest) (*emptypb.Empty, error)
	// Get returns the device for the given DevEUI.
	Get(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	// Update the given device.
	Update(context.Context, *UpdateDeviceRequest) (*emptypb.Empty, error)
	// Delete the device with the given DevEUI.
	Delete(context.Context, *DeleteDeviceRequest) (*emptypb.Empty, error)
	// Get the list of devices.
	List(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	// Create the given device-keys.
	CreateKeys(context.Context, *CreateDeviceKeysRequest) (*emptypb.Empty, error)
	// Get the device-keys for the given DevEUI.
	GetKeys(context.Context, *GetDeviceKeysRequest) (*GetDeviceKeysResponse, error)
	// Update the given device-keys.
	UpdateKeys(context.Context, *UpdateDeviceKeysRequest) (*emptypb.Empty, error)
	// Delete the device-keys for the given DevEUI.
	DeleteKeys(context.Context, *DeleteDeviceKeysRequest) (*emptypb.Empty, error)
	// FlushDevNonces flushes the OTAA device nonces.
	FlushDevNonces(context.Context, *FlushDevNoncesRequest) (*emptypb.Empty, error)
	// Activate (re)activates the device with the given parameters (for ABP or for
	// importing OTAA activations).
	Activate(context.Context, *ActivateDeviceRequest) (*emptypb.Empty, error)
	// Deactivate de-activates the device.
	Deactivate(context.Context, *DeactivateDeviceRequest) (*emptypb.Empty, error)
	// GetActivation returns the current activation details of the device (OTAA or
	// ABP).
	GetActivation(context.Context, *GetDeviceActivationRequest) (*GetDeviceActivationResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into
	// account.
	GetRandomDevAddr(context.Context, *GetRandomDevAddrRequest) (*GetRandomDevAddrResponse, error)
	// GetMetrics returns the device metrics.
	// Note that this requires a device-profile with codec and measurements
	// configured.
	GetMetrics(context.Context, *GetDeviceMetricsRequest) (*GetDeviceMetricsResponse, error)
	// GetLinkMetrics returns the device link metrics.
	// This includes uplinks, downlinks, RSSI, SNR, etc...
	GetLinkMetrics(context.Context, *GetDeviceLinkMetricsRequest) (*GetDeviceLinkMetricsResponse, error)
	// Enqueue adds the given item to the downlink queue.
	Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error)
	// FlushQueue flushes the downlink device-queue.
	FlushQueue(context.Context, *FlushDeviceQueueRequest) (*emptypb.Empty, error)
	// GetQueue returns the downlink device-queue.
	GetQueue(context.Context, *GetDeviceQueueItemsRequest) (*GetDeviceQueueItemsResponse, error)
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) Create(context.Context, *CreateDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDeviceServiceServer) Get(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDeviceServiceServer) Update(context.Context, *UpdateDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDeviceServiceServer) Delete(context.Context, *DeleteDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDeviceServiceServer) List(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDeviceServiceServer) CreateKeys(context.Context, *CreateDeviceKeysRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeys not implemented")
}
func (UnimplementedDeviceServiceServer) GetKeys(context.Context, *GetDeviceKeysRequest) (*GetDeviceKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeys not implemented")
}
func (UnimplementedDeviceServiceServer) UpdateKeys(context.Context, *UpdateDeviceKeysRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeys not implemented")
}
func (UnimplementedDeviceServiceServer) DeleteKeys(context.Context, *DeleteDeviceKeysRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKeys not implemented")
}
func (UnimplementedDeviceServiceServer) FlushDevNonces(context.Context, *FlushDevNoncesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushDevNonces not implemented")
}
func (UnimplementedDeviceServiceServer) Activate(context.Context, *ActivateDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedDeviceServiceServer) Deactivate(context.Context, *DeactivateDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedDeviceServiceServer) GetActivation(context.Context, *GetDeviceActivationRequest) (*GetDeviceActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivation not implemented")
}
func (UnimplementedDeviceServiceServer) GetRandomDevAddr(context.Context, *GetRandomDevAddrRequest) (*GetRandomDevAddrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomDevAddr not implemented")
}
func (UnimplementedDeviceServiceServer) GetMetrics(context.Context, *GetDeviceMetricsRequest) (*GetDeviceMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedDeviceServiceServer) GetLinkMetrics(context.Context, *GetDeviceLinkMetricsRequest) (*GetDeviceLinkMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkMetrics not implemented")
}
func (UnimplementedDeviceServiceServer) Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (UnimplementedDeviceServiceServer) FlushQueue(context.Context, *FlushDeviceQueueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushQueue not implemented")
}
func (UnimplementedDeviceServiceServer) GetQueue(context.Context, *GetDeviceQueueItemsRequest) (*GetDeviceQueueItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueue not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Create(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Get(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Update(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Delete(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).List(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_CreateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).CreateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/CreateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).CreateKeys(ctx, req.(*CreateDeviceKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetKeys(ctx, req.(*GetDeviceKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_UpdateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).UpdateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/UpdateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).UpdateKeys(ctx, req.(*UpdateDeviceKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_DeleteKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).DeleteKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/DeleteKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).DeleteKeys(ctx, req.(*DeleteDeviceKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_FlushDevNonces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushDevNoncesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).FlushDevNonces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/FlushDevNonces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).FlushDevNonces(ctx, req.(*FlushDevNoncesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Activate(ctx, req.(*ActivateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Deactivate(ctx, req.(*DeactivateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetActivation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetActivation(ctx, req.(*GetDeviceActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetRandomDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomDevAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetRandomDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetRandomDevAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetRandomDevAddr(ctx, req.(*GetRandomDevAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetMetrics(ctx, req.(*GetDeviceMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetLinkMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceLinkMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetLinkMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetLinkMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetLinkMetrics(ctx, req.(*GetDeviceLinkMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDeviceQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Enqueue(ctx, req.(*EnqueueDeviceQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_FlushQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushDeviceQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).FlushQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/FlushQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).FlushQueue(ctx, req.(*FlushDeviceQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetQueue(ctx, req.(*GetDeviceQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeviceService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceService_List_Handler,
		},
		{
			MethodName: "CreateKeys",
			Handler:    _DeviceService_CreateKeys_Handler,
		},
		{
			MethodName: "GetKeys",
			Handler:    _DeviceService_GetKeys_Handler,
		},
		{
			MethodName: "UpdateKeys",
			Handler:    _DeviceService_UpdateKeys_Handler,
		},
		{
			MethodName: "DeleteKeys",
			Handler:    _DeviceService_DeleteKeys_Handler,
		},
		{
			MethodName: "FlushDevNonces",
			Handler:    _DeviceService_FlushDevNonces_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _DeviceService_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _DeviceService_Deactivate_Handler,
		},
		{
			MethodName: "GetActivation",
			Handler:    _DeviceService_GetActivation_Handler,
		},
		{
			MethodName: "GetRandomDevAddr",
			Handler:    _DeviceService_GetRandomDevAddr_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _DeviceService_GetMetrics_Handler,
		},
		{
			MethodName: "GetLinkMetrics",
			Handler:    _DeviceService_GetLinkMetrics_Handler,
		},
		{
			MethodName: "Enqueue",
			Handler:    _DeviceService_Enqueue_Handler,
		},
		{
			MethodName: "FlushQueue",
			Handler:    _DeviceService_FlushQueue_Handler,
		},
		{
			MethodName: "GetQueue",
			Handler:    _DeviceService_GetQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/device.proto",
}

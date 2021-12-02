// Copyright 2021 logmesh.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: income_transaction.proto

package income_transaction_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IncomeTransaction_Status int32

const (
	IncomeTransaction_PADDING  IncomeTransaction_Status = 0
	IncomeTransaction_APPROVED IncomeTransaction_Status = 1
	IncomeTransaction_REJECT   IncomeTransaction_Status = 2
)

// Enum value maps for IncomeTransaction_Status.
var (
	IncomeTransaction_Status_name = map[int32]string{
		0: "PADDING",
		1: "APPROVED",
		2: "REJECT",
	}
	IncomeTransaction_Status_value = map[string]int32{
		"PADDING":  0,
		"APPROVED": 1,
		"REJECT":   2,
	}
)

func (x IncomeTransaction_Status) Enum() *IncomeTransaction_Status {
	p := new(IncomeTransaction_Status)
	*p = x
	return p
}

func (x IncomeTransaction_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncomeTransaction_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_income_transaction_proto_enumTypes[0].Descriptor()
}

func (IncomeTransaction_Status) Type() protoreflect.EnumType {
	return &file_income_transaction_proto_enumTypes[0]
}

func (x IncomeTransaction_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncomeTransaction_Status.Descriptor instead.
func (IncomeTransaction_Status) EnumDescriptor() ([]byte, []int) {
	return file_income_transaction_proto_rawDescGZIP(), []int{0, 0}
}

type IncomeTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId uint32                   `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`                                         // идентификатор пользователя
	Txid      string                   `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`                                                                     // хэш транзакции
	TraceId   string                   `protobuf:"bytes,3,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`                                                // идентификатор операции в пределах транзакции (txid)
	Amount    string                   `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`                                                                 // сумма
	Address   string                   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`                                                               // адрес получателя
	BlockNum  uint64                   `protobuf:"varint,6,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`                                            // порядковый номер блока
	StatusId  IncomeTransaction_Status `protobuf:"varint,7,opt,name=status_id,json=statusId,proto3,enum=logmesh.type.IncomeTransaction_Status" json:"status_id,omitempty"` // статус транзакции
	LogId     uint64                   `protobuf:"varint,8,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`                                                     // идентификатор изменения (актуальным состоянием, является наибольшее значение "log_id" для уникального ключа ["payment_system" + "currency" + "txid"])
}

func (x *IncomeTransaction) Reset() {
	*x = IncomeTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_income_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeTransaction) ProtoMessage() {}

func (x *IncomeTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_income_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeTransaction.ProtoReflect.Descriptor instead.
func (*IncomeTransaction) Descriptor() ([]byte, []int) {
	return file_income_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *IncomeTransaction) GetAccountId() uint32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *IncomeTransaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *IncomeTransaction) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *IncomeTransaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *IncomeTransaction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *IncomeTransaction) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *IncomeTransaction) GetStatusId() IncomeTransaction_Status {
	if x != nil {
		return x.StatusId
	}
	return IncomeTransaction_PADDING
}

func (x *IncomeTransaction) GetLogId() uint64 {
	if x != nil {
		return x.LogId
	}
	return 0
}

var File_income_transaction_proto protoreflect.FileDescriptor

var file_income_transaction_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x6f, 0x67, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x43, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x6c, 0x6f, 0x67, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x67, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3b, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_income_transaction_proto_rawDescOnce sync.Once
	file_income_transaction_proto_rawDescData = file_income_transaction_proto_rawDesc
)

func file_income_transaction_proto_rawDescGZIP() []byte {
	file_income_transaction_proto_rawDescOnce.Do(func() {
		file_income_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_income_transaction_proto_rawDescData)
	})
	return file_income_transaction_proto_rawDescData
}

var file_income_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_income_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_income_transaction_proto_goTypes = []interface{}{
	(IncomeTransaction_Status)(0), // 0: logmesh.type.IncomeTransaction.Status
	(*IncomeTransaction)(nil),     // 1: logmesh.type.IncomeTransaction
}
var file_income_transaction_proto_depIdxs = []int32{
	0, // 0: logmesh.type.IncomeTransaction.status_id:type_name -> logmesh.type.IncomeTransaction.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_income_transaction_proto_init() }
func file_income_transaction_proto_init() {
	if File_income_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_income_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomeTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_income_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_income_transaction_proto_goTypes,
		DependencyIndexes: file_income_transaction_proto_depIdxs,
		EnumInfos:         file_income_transaction_proto_enumTypes,
		MessageInfos:      file_income_transaction_proto_msgTypes,
	}.Build()
	File_income_transaction_proto = out.File
	file_income_transaction_proto_rawDesc = nil
	file_income_transaction_proto_goTypes = nil
	file_income_transaction_proto_depIdxs = nil
}

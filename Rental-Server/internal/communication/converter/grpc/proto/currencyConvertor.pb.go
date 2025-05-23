// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/currencyConvertor.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Currency int32

const (
	Currency_USD Currency = 0
	Currency_JPY Currency = 1
	Currency_BGN Currency = 2
	Currency_CZK Currency = 3
	Currency_DKK Currency = 4
	Currency_GBP Currency = 5
	Currency_HUF Currency = 6
	Currency_PLN Currency = 7
	Currency_RON Currency = 8
	Currency_SEK Currency = 9
	Currency_CHF Currency = 10
	Currency_ISK Currency = 11
	Currency_NOK Currency = 12
	Currency_TRY Currency = 13
	Currency_AUD Currency = 14
	Currency_BRL Currency = 15
	Currency_CAD Currency = 16
	Currency_CNY Currency = 17
	Currency_HKD Currency = 18
	Currency_IDR Currency = 19
	Currency_ILS Currency = 20
	Currency_INR Currency = 21
	Currency_KRW Currency = 22
	Currency_MXN Currency = 23
	Currency_MYR Currency = 24
	Currency_NZD Currency = 25
	Currency_PHP Currency = 26
	Currency_SGD Currency = 27
	Currency_THB Currency = 28
	Currency_ZAR Currency = 29
	Currency_EUR Currency = 30
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0:  "USD",
		1:  "JPY",
		2:  "BGN",
		3:  "CZK",
		4:  "DKK",
		5:  "GBP",
		6:  "HUF",
		7:  "PLN",
		8:  "RON",
		9:  "SEK",
		10: "CHF",
		11: "ISK",
		12: "NOK",
		13: "TRY",
		14: "AUD",
		15: "BRL",
		16: "CAD",
		17: "CNY",
		18: "HKD",
		19: "IDR",
		20: "ILS",
		21: "INR",
		22: "KRW",
		23: "MXN",
		24: "MYR",
		25: "NZD",
		26: "PHP",
		27: "SGD",
		28: "THB",
		29: "ZAR",
		30: "EUR",
	}
	Currency_value = map[string]int32{
		"USD": 0,
		"JPY": 1,
		"BGN": 2,
		"CZK": 3,
		"DKK": 4,
		"GBP": 5,
		"HUF": 6,
		"PLN": 7,
		"RON": 8,
		"SEK": 9,
		"CHF": 10,
		"ISK": 11,
		"NOK": 12,
		"TRY": 13,
		"AUD": 14,
		"BRL": 15,
		"CAD": 16,
		"CNY": 17,
		"HKD": 18,
		"IDR": 19,
		"ILS": 20,
		"INR": 21,
		"KRW": 22,
		"MXN": 23,
		"MYR": 24,
		"NZD": 25,
		"PHP": 26,
		"SGD": 27,
		"THB": 28,
		"ZAR": 29,
		"EUR": 30,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_currencyConvertor_proto_enumTypes[0].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_proto_currencyConvertor_proto_enumTypes[0]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_proto_currencyConvertor_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_currencyConvertor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyConvertor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_currencyConvertor_proto_rawDescGZIP(), []int{0}
}

type CurrencyList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currencies    []Currency             `protobuf:"varint,1,rep,packed,name=currencies,proto3,enum=Currency" json:"currencies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrencyList) Reset() {
	*x = CurrencyList{}
	mi := &file_proto_currencyConvertor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyList) ProtoMessage() {}

func (x *CurrencyList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyConvertor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyList.ProtoReflect.Descriptor instead.
func (*CurrencyList) Descriptor() ([]byte, []int) {
	return file_proto_currencyConvertor_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyList) GetCurrencies() []Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

type ConversionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Given         Currency               `protobuf:"varint,1,opt,name=given,proto3,enum=Currency" json:"given,omitempty"`
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Target        Currency               `protobuf:"varint,3,opt,name=target,proto3,enum=Currency" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionRequest) Reset() {
	*x = ConversionRequest{}
	mi := &file_proto_currencyConvertor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionRequest) ProtoMessage() {}

func (x *ConversionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyConvertor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionRequest.ProtoReflect.Descriptor instead.
func (*ConversionRequest) Descriptor() ([]byte, []int) {
	return file_proto_currencyConvertor_proto_rawDescGZIP(), []int{2}
}

func (x *ConversionRequest) GetGiven() Currency {
	if x != nil {
		return x.Given
	}
	return Currency_USD
}

func (x *ConversionRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ConversionRequest) GetTarget() Currency {
	if x != nil {
		return x.Target
	}
	return Currency_USD
}

type ConversionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Converted     Currency               `protobuf:"varint,1,opt,name=converted,proto3,enum=Currency" json:"converted,omitempty"`
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionResponse) Reset() {
	*x = ConversionResponse{}
	mi := &file_proto_currencyConvertor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionResponse) ProtoMessage() {}

func (x *ConversionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencyConvertor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionResponse.ProtoReflect.Descriptor instead.
func (*ConversionResponse) Descriptor() ([]byte, []int) {
	return file_proto_currencyConvertor_proto_rawDescGZIP(), []int{3}
}

func (x *ConversionResponse) GetConverted() Currency {
	if x != nil {
		return x.Converted
	}
	return Currency_USD
}

func (x *ConversionResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_proto_currencyConvertor_proto protoreflect.FileDescriptor

const file_proto_currencyConvertor_proto_rawDesc = "" +
	"\n" +
	"\x1dproto/currencyConvertor.proto\"\a\n" +
	"\x05empty\"9\n" +
	"\fcurrencyList\x12)\n" +
	"\n" +
	"currencies\x18\x01 \x03(\x0e2\t.currencyR\n" +
	"currencies\"o\n" +
	"\x11conversionRequest\x12\x1f\n" +
	"\x05given\x18\x01 \x01(\x0e2\t.currencyR\x05given\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x01R\x06amount\x12!\n" +
	"\x06target\x18\x03 \x01(\x0e2\t.currencyR\x06target\"U\n" +
	"\x12conversionResponse\x12'\n" +
	"\tconverted\x18\x01 \x01(\x0e2\t.currencyR\tconverted\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x01R\x06amount*\xa1\x02\n" +
	"\bcurrency\x12\a\n" +
	"\x03USD\x10\x00\x12\a\n" +
	"\x03JPY\x10\x01\x12\a\n" +
	"\x03BGN\x10\x02\x12\a\n" +
	"\x03CZK\x10\x03\x12\a\n" +
	"\x03DKK\x10\x04\x12\a\n" +
	"\x03GBP\x10\x05\x12\a\n" +
	"\x03HUF\x10\x06\x12\a\n" +
	"\x03PLN\x10\a\x12\a\n" +
	"\x03RON\x10\b\x12\a\n" +
	"\x03SEK\x10\t\x12\a\n" +
	"\x03CHF\x10\n" +
	"\x12\a\n" +
	"\x03ISK\x10\v\x12\a\n" +
	"\x03NOK\x10\f\x12\a\n" +
	"\x03TRY\x10\r\x12\a\n" +
	"\x03AUD\x10\x0e\x12\a\n" +
	"\x03BRL\x10\x0f\x12\a\n" +
	"\x03CAD\x10\x10\x12\a\n" +
	"\x03CNY\x10\x11\x12\a\n" +
	"\x03HKD\x10\x12\x12\a\n" +
	"\x03IDR\x10\x13\x12\a\n" +
	"\x03ILS\x10\x14\x12\a\n" +
	"\x03INR\x10\x15\x12\a\n" +
	"\x03KRW\x10\x16\x12\a\n" +
	"\x03MXN\x10\x17\x12\a\n" +
	"\x03MYR\x10\x18\x12\a\n" +
	"\x03NZD\x10\x19\x12\a\n" +
	"\x03PHP\x10\x1a\x12\a\n" +
	"\x03SGD\x10\x1b\x12\a\n" +
	"\x03THB\x10\x1c\x12\a\n" +
	"\x03ZAR\x10\x1d\x12\a\n" +
	"\x03EUR\x10\x1e2r\n" +
	"\tConvertor\x124\n" +
	"\aconvert\x12\x12.conversionRequest\x1a\x13.conversionResponse\"\x00\x12/\n" +
	"\x14getAvailableCurrency\x12\x06.empty\x1a\r.currencyList\"\x00B\x80\x01\n" +
	"+com.prestige.wheels.currency.converter.grpcZQgithub.com/megamxl/se-project/Rental-Server/internal/communication/converter/grpcb\x06proto3"

var (
	file_proto_currencyConvertor_proto_rawDescOnce sync.Once
	file_proto_currencyConvertor_proto_rawDescData []byte
)

func file_proto_currencyConvertor_proto_rawDescGZIP() []byte {
	file_proto_currencyConvertor_proto_rawDescOnce.Do(func() {
		file_proto_currencyConvertor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_currencyConvertor_proto_rawDesc), len(file_proto_currencyConvertor_proto_rawDesc)))
	})
	return file_proto_currencyConvertor_proto_rawDescData
}

var file_proto_currencyConvertor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_currencyConvertor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_currencyConvertor_proto_goTypes = []any{
	(Currency)(0),              // 0: currency
	(*Empty)(nil),              // 1: empty
	(*CurrencyList)(nil),       // 2: currencyList
	(*ConversionRequest)(nil),  // 3: conversionRequest
	(*ConversionResponse)(nil), // 4: conversionResponse
}
var file_proto_currencyConvertor_proto_depIdxs = []int32{
	0, // 0: currencyList.currencies:type_name -> currency
	0, // 1: conversionRequest.given:type_name -> currency
	0, // 2: conversionRequest.target:type_name -> currency
	0, // 3: conversionResponse.converted:type_name -> currency
	3, // 4: Convertor.convert:input_type -> conversionRequest
	1, // 5: Convertor.getAvailableCurrency:input_type -> empty
	4, // 6: Convertor.convert:output_type -> conversionResponse
	2, // 7: Convertor.getAvailableCurrency:output_type -> currencyList
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_currencyConvertor_proto_init() }
func file_proto_currencyConvertor_proto_init() {
	if File_proto_currencyConvertor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_currencyConvertor_proto_rawDesc), len(file_proto_currencyConvertor_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currencyConvertor_proto_goTypes,
		DependencyIndexes: file_proto_currencyConvertor_proto_depIdxs,
		EnumInfos:         file_proto_currencyConvertor_proto_enumTypes,
		MessageInfos:      file_proto_currencyConvertor_proto_msgTypes,
	}.Build()
	File_proto_currencyConvertor_proto = out.File
	file_proto_currencyConvertor_proto_goTypes = nil
	file_proto_currencyConvertor_proto_depIdxs = nil
}

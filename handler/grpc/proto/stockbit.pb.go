// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: handler/grpc/proto/stockbit.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetOneMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetOneMovieRequest) Reset() {
	*x = GetOneMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneMovieRequest) ProtoMessage() {}

func (x *GetOneMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneMovieRequest.ProtoReflect.Descriptor instead.
func (*GetOneMovieRequest) Descriptor() ([]byte, []int) {
	return file_handler_grpc_proto_stockbit_proto_rawDescGZIP(), []int{0}
}

func (x *GetOneMovieRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOneMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string    `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string    `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string    `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string    `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string    `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string    `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string    `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string    `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string    `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string    `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string    `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string    `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string    `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Ratings    []*Rating `protobuf:"bytes,15,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Metascore  string    `protobuf:"bytes,16,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating string    `protobuf:"bytes,17,opt,name=ImdbRating,proto3" json:"ImdbRating,omitempty"`
	ImdbID     string    `protobuf:"bytes,18,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type       string    `protobuf:"bytes,19,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string    `protobuf:"bytes,20,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string    `protobuf:"bytes,21,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string    `protobuf:"bytes,22,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string    `protobuf:"bytes,23,opt,name=Website,proto3" json:"Website,omitempty"`
	Response   string    `protobuf:"bytes,24,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *GetOneMovieResponse) Reset() {
	*x = GetOneMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneMovieResponse) ProtoMessage() {}

func (x *GetOneMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneMovieResponse.ProtoReflect.Descriptor instead.
func (*GetOneMovieResponse) Descriptor() ([]byte, []int) {
	return file_handler_grpc_proto_stockbit_proto_rawDescGZIP(), []int{1}
}

func (x *GetOneMovieResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetOneMovieResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetOneMovieResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *GetOneMovieResponse) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *GetOneMovieResponse) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *GetOneMovieResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetOneMovieResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *GetOneMovieResponse) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *GetOneMovieResponse) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *GetOneMovieResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *GetOneMovieResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetOneMovieResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetOneMovieResponse) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *GetOneMovieResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *GetOneMovieResponse) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetOneMovieResponse) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *GetOneMovieResponse) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *GetOneMovieResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetOneMovieResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetOneMovieResponse) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *GetOneMovieResponse) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *GetOneMovieResponse) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *GetOneMovieResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetOneMovieResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_handler_grpc_proto_stockbit_proto_rawDescGZIP(), []int{2}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SearchMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Page string `protobuf:"bytes,2,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (x *SearchMovieRequest) Reset() {
	*x = SearchMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMovieRequest) ProtoMessage() {}

func (x *SearchMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMovieRequest.ProtoReflect.Descriptor instead.
func (*SearchMovieRequest) Descriptor() ([]byte, []int) {
	return file_handler_grpc_proto_stockbit_proto_rawDescGZIP(), []int{3}
}

func (x *SearchMovieRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SearchMovieRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type SearchMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*MovieSummary `protobuf:"bytes,1,rep,name=Search,proto3" json:"Search,omitempty"`
	TotalResults string          `protobuf:"bytes,2,opt,name=TotalResults,proto3" json:"TotalResults,omitempty"`
	Response     string          `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *SearchMovieResponse) Reset() {
	*x = SearchMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMovieResponse) ProtoMessage() {}

func (x *SearchMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMovieResponse.ProtoReflect.Descriptor instead.
func (*SearchMovieResponse) Descriptor() ([]byte, []int) {
	return file_handler_grpc_proto_stockbit_proto_rawDescGZIP(), []int{4}
}

func (x *SearchMovieResponse) GetSearch() []*MovieSummary {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *SearchMovieResponse) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *SearchMovieResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type MovieSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *MovieSummary) Reset() {
	*x = MovieSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieSummary) ProtoMessage() {}

func (x *MovieSummary) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_proto_stockbit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieSummary.ProtoReflect.Descriptor instead.
func (*MovieSummary) Descriptor() ([]byte, []int) {
	return file_handler_grpc_proto_stockbit_proto_rawDescGZIP(), []int{5}
}

func (x *MovieSummary) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieSummary) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieSummary) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieSummary) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieSummary) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

var File_handler_grpc_proto_stockbit_proto protoreflect.FileDescriptor

var file_handler_grpc_proto_stockbit_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x22, 0x83, 0x05, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6d, 0x64, 0x62, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6d, 0x64,
	0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x44, 0x56, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3a, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0x85, 0x01,
	0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7c, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x32, 0xa6, 0x01, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74,
	0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x1c, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x62, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handler_grpc_proto_stockbit_proto_rawDescOnce sync.Once
	file_handler_grpc_proto_stockbit_proto_rawDescData = file_handler_grpc_proto_stockbit_proto_rawDesc
)

func file_handler_grpc_proto_stockbit_proto_rawDescGZIP() []byte {
	file_handler_grpc_proto_stockbit_proto_rawDescOnce.Do(func() {
		file_handler_grpc_proto_stockbit_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_grpc_proto_stockbit_proto_rawDescData)
	})
	return file_handler_grpc_proto_stockbit_proto_rawDescData
}

var file_handler_grpc_proto_stockbit_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_handler_grpc_proto_stockbit_proto_goTypes = []interface{}{
	(*GetOneMovieRequest)(nil),  // 0: stockbit.GetOneMovieRequest
	(*GetOneMovieResponse)(nil), // 1: stockbit.GetOneMovieResponse
	(*Rating)(nil),              // 2: stockbit.Rating
	(*SearchMovieRequest)(nil),  // 3: stockbit.SearchMovieRequest
	(*SearchMovieResponse)(nil), // 4: stockbit.SearchMovieResponse
	(*MovieSummary)(nil),        // 5: stockbit.MovieSummary
}
var file_handler_grpc_proto_stockbit_proto_depIdxs = []int32{
	2, // 0: stockbit.GetOneMovieResponse.Ratings:type_name -> stockbit.Rating
	5, // 1: stockbit.SearchMovieResponse.Search:type_name -> stockbit.MovieSummary
	0, // 2: stockbit.stockbit.GetOneMovie:input_type -> stockbit.GetOneMovieRequest
	3, // 3: stockbit.stockbit.SearchMovie:input_type -> stockbit.SearchMovieRequest
	1, // 4: stockbit.stockbit.GetOneMovie:output_type -> stockbit.GetOneMovieResponse
	4, // 5: stockbit.stockbit.SearchMovie:output_type -> stockbit.SearchMovieResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_handler_grpc_proto_stockbit_proto_init() }
func file_handler_grpc_proto_stockbit_proto_init() {
	if File_handler_grpc_proto_stockbit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_grpc_proto_stockbit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneMovieRequest); i {
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
		file_handler_grpc_proto_stockbit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneMovieResponse); i {
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
		file_handler_grpc_proto_stockbit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_handler_grpc_proto_stockbit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMovieRequest); i {
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
		file_handler_grpc_proto_stockbit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMovieResponse); i {
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
		file_handler_grpc_proto_stockbit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieSummary); i {
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
			RawDescriptor: file_handler_grpc_proto_stockbit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_grpc_proto_stockbit_proto_goTypes,
		DependencyIndexes: file_handler_grpc_proto_stockbit_proto_depIdxs,
		MessageInfos:      file_handler_grpc_proto_stockbit_proto_msgTypes,
	}.Build()
	File_handler_grpc_proto_stockbit_proto = out.File
	file_handler_grpc_proto_stockbit_proto_rawDesc = nil
	file_handler_grpc_proto_stockbit_proto_goTypes = nil
	file_handler_grpc_proto_stockbit_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StockbitClient is the client API for Stockbit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockbitClient interface {
	GetOneMovie(ctx context.Context, in *GetOneMovieRequest, opts ...grpc.CallOption) (*GetOneMovieResponse, error)
	SearchMovie(ctx context.Context, in *SearchMovieRequest, opts ...grpc.CallOption) (*SearchMovieResponse, error)
}

type stockbitClient struct {
	cc grpc.ClientConnInterface
}

func NewStockbitClient(cc grpc.ClientConnInterface) StockbitClient {
	return &stockbitClient{cc}
}

func (c *stockbitClient) GetOneMovie(ctx context.Context, in *GetOneMovieRequest, opts ...grpc.CallOption) (*GetOneMovieResponse, error) {
	out := new(GetOneMovieResponse)
	err := c.cc.Invoke(ctx, "/stockbit.stockbit/GetOneMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockbitClient) SearchMovie(ctx context.Context, in *SearchMovieRequest, opts ...grpc.CallOption) (*SearchMovieResponse, error) {
	out := new(SearchMovieResponse)
	err := c.cc.Invoke(ctx, "/stockbit.stockbit/SearchMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockbitServer is the server API for Stockbit service.
type StockbitServer interface {
	GetOneMovie(context.Context, *GetOneMovieRequest) (*GetOneMovieResponse, error)
	SearchMovie(context.Context, *SearchMovieRequest) (*SearchMovieResponse, error)
}

// UnimplementedStockbitServer can be embedded to have forward compatible implementations.
type UnimplementedStockbitServer struct {
}

func (*UnimplementedStockbitServer) GetOneMovie(context.Context, *GetOneMovieRequest) (*GetOneMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneMovie not implemented")
}
func (*UnimplementedStockbitServer) SearchMovie(context.Context, *SearchMovieRequest) (*SearchMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovie not implemented")
}

func RegisterStockbitServer(s *grpc.Server, srv StockbitServer) {
	s.RegisterService(&_Stockbit_serviceDesc, srv)
}

func _Stockbit_GetOneMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockbitServer).GetOneMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stockbit.stockbit/GetOneMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockbitServer).GetOneMovie(ctx, req.(*GetOneMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stockbit_SearchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockbitServer).SearchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stockbit.stockbit/SearchMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockbitServer).SearchMovie(ctx, req.(*SearchMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stockbit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stockbit.stockbit",
	HandlerType: (*StockbitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneMovie",
			Handler:    _Stockbit_GetOneMovie_Handler,
		},
		{
			MethodName: "SearchMovie",
			Handler:    _Stockbit_SearchMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handler/grpc/proto/stockbit.proto",
}
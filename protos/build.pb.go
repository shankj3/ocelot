// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	build.proto
	common.proto
	commonevententities.proto
	projectrootdir.proto
	respositories.proto
	webhook.proto

It has these top-level messages:
	BuildConfig
	Stage
	WerkerTask
	PushBuildBundle
	PRBuildBundle
	LinkUrl
	LinkAndName
	Links
	Owner
	Repository
	PullRequestEntity
	PRInfo
	Project
	Changeset
	Commit
	RepoSourceFile
	PaginatedRootDirs
	PaginatedRepository
	RepoPush
	PullRequest
	PullRequestApproved
	CreateWebhook
	GetWebhooks
	Webhooks
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protos2 "github.com/shankj3/ocelot/protos/leveler_resources"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// this is a direct translation of the ocelot.yaml file
type BuildConfig struct {
	Image    string            `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Packages []string          `protobuf:"bytes,2,rep,name=packages" json:"packages,omitempty"`
	Env      map[string]string `protobuf:"bytes,3,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Before   *Stage            `protobuf:"bytes,4,opt,name=before" json:"before,omitempty"`
	After    *Stage            `protobuf:"bytes,5,opt,name=after" json:"after,omitempty"`
	Build    *Stage            `protobuf:"bytes,6,opt,name=build" json:"build,omitempty"`
	Test     *Stage            `protobuf:"bytes,7,opt,name=test" json:"test,omitempty"`
	Deploy   *Stage            `protobuf:"bytes,8,opt,name=deploy" json:"deploy,omitempty"`
}

func (m *BuildConfig) Reset()                    { *m = BuildConfig{} }
func (m *BuildConfig) String() string            { return proto.CompactTextString(m) }
func (*BuildConfig) ProtoMessage()               {}
func (*BuildConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildConfig) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *BuildConfig) GetPackages() []string {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *BuildConfig) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *BuildConfig) GetBefore() *Stage {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *BuildConfig) GetAfter() *Stage {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *BuildConfig) GetBuild() *Stage {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *BuildConfig) GetTest() *Stage {
	if m != nil {
		return m.Test
	}
	return nil
}

func (m *BuildConfig) GetDeploy() *Stage {
	if m != nil {
		return m.Deploy
	}
	return nil
}

type Stage struct {
	Env    map[string]string `protobuf:"bytes,1,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Script []string          `protobuf:"bytes,2,rep,name=script" json:"script,omitempty"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Stage) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Stage) GetScript() []string {
	if m != nil {
		return m.Script
	}
	return nil
}

type WerkerTask struct {
	VaultToken   string                  `protobuf:"bytes,1,opt,name=vaultToken" json:"vaultToken,omitempty"`
	CheckoutHash string                  `protobuf:"bytes,2,opt,name=checkoutHash" json:"checkoutHash,omitempty"`
	Pipe         *protos2.PipelineConfig `protobuf:"bytes,3,opt,name=pipe" json:"pipe,omitempty"`
}

func (m *WerkerTask) Reset()                    { *m = WerkerTask{} }
func (m *WerkerTask) String() string            { return proto.CompactTextString(m) }
func (*WerkerTask) ProtoMessage()               {}
func (*WerkerTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WerkerTask) GetVaultToken() string {
	if m != nil {
		return m.VaultToken
	}
	return ""
}

func (m *WerkerTask) GetCheckoutHash() string {
	if m != nil {
		return m.CheckoutHash
	}
	return ""
}

func (m *WerkerTask) GetPipe() *protos2.PipelineConfig {
	if m != nil {
		return m.Pipe
	}
	return nil
}

type PushBuildBundle struct {
	Config       *BuildConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	PushData     *RepoPush    `protobuf:"bytes,2,opt,name=pushData" json:"pushData,omitempty"`
	VaultToken   string       `protobuf:"bytes,3,opt,name=vaultToken" json:"vaultToken,omitempty"`
	CheckoutHash string       `protobuf:"bytes,4,opt,name=checkoutHash" json:"checkoutHash,omitempty"`
}

func (m *PushBuildBundle) Reset()                    { *m = PushBuildBundle{} }
func (m *PushBuildBundle) String() string            { return proto.CompactTextString(m) }
func (*PushBuildBundle) ProtoMessage()               {}
func (*PushBuildBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PushBuildBundle) GetConfig() *BuildConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PushBuildBundle) GetPushData() *RepoPush {
	if m != nil {
		return m.PushData
	}
	return nil
}

func (m *PushBuildBundle) GetVaultToken() string {
	if m != nil {
		return m.VaultToken
	}
	return ""
}

func (m *PushBuildBundle) GetCheckoutHash() string {
	if m != nil {
		return m.CheckoutHash
	}
	return ""
}

type PRBuildBundle struct {
	Config       *BuildConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	PrData       *PullRequest `protobuf:"bytes,2,opt,name=prData" json:"prData,omitempty"`
	VaultToken   string       `protobuf:"bytes,3,opt,name=vaultToken" json:"vaultToken,omitempty"`
	CheckoutHash string       `protobuf:"bytes,4,opt,name=checkoutHash" json:"checkoutHash,omitempty"`
}

func (m *PRBuildBundle) Reset()                    { *m = PRBuildBundle{} }
func (m *PRBuildBundle) String() string            { return proto.CompactTextString(m) }
func (*PRBuildBundle) ProtoMessage()               {}
func (*PRBuildBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PRBuildBundle) GetConfig() *BuildConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PRBuildBundle) GetPrData() *PullRequest {
	if m != nil {
		return m.PrData
	}
	return nil
}

func (m *PRBuildBundle) GetVaultToken() string {
	if m != nil {
		return m.VaultToken
	}
	return ""
}

func (m *PRBuildBundle) GetCheckoutHash() string {
	if m != nil {
		return m.CheckoutHash
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildConfig)(nil), "protos.BuildConfig")
	proto.RegisterType((*Stage)(nil), "protos.Stage")
	proto.RegisterType((*WerkerTask)(nil), "protos.WerkerTask")
	proto.RegisterType((*PushBuildBundle)(nil), "protos.PushBuildBundle")
	proto.RegisterType((*PRBuildBundle)(nil), "protos.PRBuildBundle")
}

func init() { proto.RegisterFile("build.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x49, 0xd3, 0xc6, 0xde, 0xc4, 0xea, 0xb1, 0xca, 0xb1, 0x14, 0x91, 0x1a, 0x11, 0x82,
	0x27, 0x7d, 0xa8, 0x20, 0xe2, 0xe3, 0xe9, 0x81, 0x8f, 0x65, 0x3d, 0xf0, 0x79, 0x9b, 0x4e, 0x9b,
	0x90, 0x98, 0x5d, 0xb3, 0xbb, 0x95, 0x82, 0x2f, 0x7e, 0x20, 0xc1, 0xcf, 0xe4, 0x27, 0x91, 0xec,
	0x6e, 0x43, 0xef, 0x0c, 0x08, 0xc5, 0xa7, 0x64, 0x66, 0x7e, 0x3b, 0xfb, 0x9f, 0x7f, 0x26, 0x10,
	0xaf, 0x4c, 0x51, 0xad, 0xe7, 0xb2, 0x11, 0x5a, 0x90, 0xc8, 0x3e, 0xd4, 0x74, 0xf2, 0x0d, 0x57,
	0xb9, 0x10, 0xa5, 0x4b, 0x4f, 0x1f, 0xc8, 0x42, 0x62, 0x55, 0xd4, 0xe8, 0xe2, 0xe4, 0xf7, 0x00,
	0xe2, 0xab, 0xf6, 0xd8, 0x7b, 0x51, 0x6f, 0x8a, 0x2d, 0x79, 0x0c, 0xa3, 0xe2, 0x0b, 0xdf, 0x22,
	0x0d, 0x66, 0x41, 0x7a, 0xc6, 0x5c, 0x40, 0xa6, 0x30, 0x96, 0x3c, 0x2b, 0xf9, 0x16, 0x15, 0x1d,
	0xcc, 0xc2, 0xf4, 0x8c, 0x75, 0x31, 0x99, 0x43, 0x88, 0xf5, 0x8e, 0x86, 0xb3, 0x30, 0x8d, 0x17,
	0x4f, 0x5c, 0x5b, 0x35, 0x3f, 0xea, 0x39, 0xbf, 0xae, 0x77, 0xd7, 0xb5, 0x6e, 0xf6, 0xac, 0x05,
	0xc9, 0x0b, 0x88, 0x56, 0xb8, 0x11, 0x0d, 0xd2, 0xe1, 0x2c, 0x48, 0xe3, 0xc5, 0xe4, 0x70, 0xe4,
	0x93, 0xe6, 0x5b, 0x64, 0xbe, 0x48, 0x9e, 0xc3, 0x88, 0x6f, 0x34, 0x36, 0x74, 0xd4, 0x47, 0xb9,
	0x5a, 0x0b, 0xd9, 0x99, 0x69, 0xd4, 0x0b, 0xd9, 0x1a, 0x79, 0x06, 0x43, 0x8d, 0x4a, 0xd3, 0x7b,
	0x7d, 0x8c, 0x2d, 0xb5, 0x9a, 0xd6, 0x28, 0x2b, 0xb1, 0xa7, 0xe3, 0x5e, 0x4d, 0xae, 0x38, 0x7d,
	0x03, 0xe3, 0xc3, 0x2c, 0xe4, 0x1c, 0xc2, 0x12, 0xf7, 0xde, 0xa6, 0xf6, 0xb5, 0xb5, 0x6e, 0xc7,
	0x2b, 0x83, 0x74, 0xe0, 0xac, 0xb3, 0xc1, 0xbb, 0xc1, 0xdb, 0x20, 0xf9, 0x11, 0xc0, 0xc8, 0x76,
	0x22, 0xa9, 0x33, 0x2b, 0xb0, 0x66, 0x5d, 0xdc, 0xba, 0xe5, 0x8e, 0x4d, 0x17, 0x10, 0xa9, 0xac,
	0x29, 0xa4, 0xf6, 0x86, 0xfb, 0xe8, 0x64, 0x0d, 0xdf, 0x01, 0x3e, 0x63, 0x53, 0x62, 0x73, 0xc3,
	0x55, 0x49, 0x9e, 0x02, 0xec, 0xb8, 0xa9, 0xf4, 0x8d, 0x28, 0xb1, 0xf6, 0x0d, 0x8e, 0x32, 0x24,
	0x81, 0xfb, 0x59, 0x8e, 0x59, 0x29, 0x8c, 0xfe, 0xc8, 0x55, 0xee, 0xdb, 0xdd, 0xca, 0x91, 0x97,
	0x30, 0x6c, 0x97, 0x89, 0x86, 0xd6, 0xb2, 0x6e, 0x98, 0xa5, 0x5f, 0x30, 0xf7, 0xf1, 0x99, 0x65,
	0x92, 0x5f, 0x01, 0x3c, 0x5c, 0x1a, 0x95, 0xdb, 0xb5, 0xb8, 0x32, 0xf5, 0xba, 0x42, 0x72, 0x09,
	0x51, 0x66, 0x19, 0x7b, 0x7f, 0xbc, 0x78, 0xd4, 0xb3, 0x3b, 0xcc, 0x23, 0xe4, 0x15, 0x8c, 0xa5,
	0x51, 0xf9, 0x07, 0xae, 0xb9, 0x15, 0x13, 0x2f, 0xce, 0x0f, 0x38, 0x43, 0x29, 0xda, 0xde, 0xac,
	0x23, 0xee, 0x8c, 0x17, 0xfe, 0x73, 0xbc, 0xe1, 0xdf, 0xe3, 0x25, 0x3f, 0x03, 0x98, 0x2c, 0xd9,
	0xc9, 0x82, 0x2f, 0x21, 0x92, 0xcd, 0x91, 0xdc, 0x0e, 0x5e, 0x9a, 0xaa, 0x62, 0xf8, 0xd5, 0xa0,
	0xd2, 0xcc, 0x23, 0xff, 0x43, 0xef, 0xca, 0xfd, 0xf0, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x21, 0xa3, 0xb2, 0x06, 0x04, 0x00, 0x00,
}

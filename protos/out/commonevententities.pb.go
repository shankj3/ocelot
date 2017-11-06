// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commonevententities.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	commonevententities.proto
	common.proto
	projectrootdir.proto
	respositories.proto
	webhook.proto

It has these top-level messages:
	Owner
	Repository
	Project
	Changeset
	Commit
	LinkUrl
	LinkAndName
	Links
	PaginatedRootDirs
	PaginatedRepository
	RepoPush
	CreateWebhook
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-entity_userOwner
type Owner struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Links    *Links `protobuf:"bytes,3,opt,name=links" json:"links,omitempty"`
}

func (m *Owner) Reset()                    { *m = Owner{} }
func (m *Owner) String() string            { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()               {}
func (*Owner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Owner) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Owner) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Owner) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-entity_repositoryRepository
type Repository struct {
	Links    *Links   `protobuf:"bytes,1,opt,name=links" json:"links,omitempty"`
	Project  *Project `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	FullName string   `protobuf:"bytes,3,opt,name=fullName,json=full_name" json:"fullName,omitempty"`
	Website  string   `protobuf:"bytes,4,opt,name=website" json:"website,omitempty"`
	Owner    *Owner   `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Repository) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Repository) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *Repository) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Repository) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Repository) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-entity_projectProject
type Project struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Uuid  string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Links *Links `protobuf:"bytes,3,opt,name=links" json:"links,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Project) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

type Changeset struct {
	New     *Changeset_Head `protobuf:"bytes,1,opt,name=new" json:"new,omitempty"`
	Old     *Changeset_Head `protobuf:"bytes,2,opt,name=old" json:"old,omitempty"`
	Links   *Links          `protobuf:"bytes,3,opt,name=links" json:"links,omitempty"`
	Closed  bool            `protobuf:"varint,4,opt,name=closed" json:"closed,omitempty"`
	Created bool            `protobuf:"varint,5,opt,name=created" json:"created,omitempty"`
	Forced  bool            `protobuf:"varint,6,opt,name=forced" json:"forced,omitempty"`
	Commits []*Commit       `protobuf:"bytes,7,rep,name=commits" json:"commits,omitempty"`
}

func (m *Changeset) Reset()                    { *m = Changeset{} }
func (m *Changeset) String() string            { return proto.CompactTextString(m) }
func (*Changeset) ProtoMessage()               {}
func (*Changeset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Changeset) GetNew() *Changeset_Head {
	if m != nil {
		return m.New
	}
	return nil
}

func (m *Changeset) GetOld() *Changeset_Head {
	if m != nil {
		return m.Old
	}
	return nil
}

func (m *Changeset) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Changeset) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func (m *Changeset) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Changeset) GetForced() bool {
	if m != nil {
		return m.Forced
	}
	return false
}

func (m *Changeset) GetCommits() []*Commit {
	if m != nil {
		return m.Commits
	}
	return nil
}

type Changeset_Head struct {
	Type   string  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Target *Commit `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
}

func (m *Changeset_Head) Reset()                    { *m = Changeset_Head{} }
func (m *Changeset_Head) String() string            { return proto.CompactTextString(m) }
func (*Changeset_Head) ProtoMessage()               {}
func (*Changeset_Head) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Changeset_Head) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Changeset_Head) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Changeset_Head) GetTarget() *Commit {
	if m != nil {
		return m.Target
	}
	return nil
}

type Commit struct {
	Hash    string                     `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Author  *Owner                     `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Message string                     `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Date    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
	// ignoring the "parents" field
	Links *Links `protobuf:"bytes,5,opt,name=links" json:"links,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Commit) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Commit) GetAuthor() *Owner {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Commit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Commit) GetDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Commit) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func init() {
	proto.RegisterType((*Owner)(nil), "protos.Owner")
	proto.RegisterType((*Repository)(nil), "protos.Repository")
	proto.RegisterType((*Project)(nil), "protos.Project")
	proto.RegisterType((*Changeset)(nil), "protos.Changeset")
	proto.RegisterType((*Changeset_Head)(nil), "protos.Changeset.Head")
	proto.RegisterType((*Commit)(nil), "protos.Commit")
}

func init() { proto.RegisterFile("commonevententities.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0x34, 0x69, 0xa7, 0xfc, 0x48, 0x3e, 0xac, 0x42, 0x38, 0x50, 0x15, 0x81, 0xc2,
	0x25, 0x2b, 0x95, 0x47, 0xe0, 0xc2, 0x01, 0x01, 0xb2, 0xd0, 0x9e, 0x90, 0x90, 0x9b, 0x4c, 0x93,
	0x40, 0x62, 0x47, 0xb6, 0x43, 0xb5, 0xaf, 0xc5, 0x89, 0xc7, 0xe0, 0x91, 0x90, 0xff, 0xba, 0xb0,
	0xda, 0x4a, 0x3d, 0xc5, 0xdf, 0x7c, 0x33, 0x93, 0xcf, 0x9f, 0x67, 0xe0, 0x59, 0x25, 0x86, 0x41,
	0x70, 0xfc, 0x89, 0x5c, 0x23, 0xd7, 0x9d, 0xee, 0x50, 0x95, 0xa3, 0x14, 0x5a, 0x90, 0xc4, 0x7e,
	0x54, 0xfe, 0xc8, 0xa5, 0xb8, 0x68, 0xfe, 0xa2, 0x11, 0xa2, 0xe9, 0xf1, 0xda, 0xa2, 0xfd, 0x74,
	0xb8, 0xd6, 0xdd, 0x80, 0x4a, 0xb3, 0x61, 0x74, 0x09, 0xdb, 0xaf, 0xb0, 0xf8, 0x74, 0xe4, 0x28,
	0x09, 0x81, 0x58, 0xdf, 0x8e, 0x98, 0x45, 0x9b, 0xa8, 0x58, 0x51, 0x7b, 0x26, 0x39, 0x2c, 0x27,
	0x85, 0x92, 0xb3, 0x01, 0xb3, 0x99, 0x8d, 0x9f, 0x30, 0x79, 0x09, 0x8b, 0xbe, 0xe3, 0x3f, 0x54,
	0x36, 0xdf, 0x44, 0xc5, 0x7a, 0xf7, 0xd8, 0xf5, 0x53, 0xe5, 0x07, 0x13, 0xa4, 0x8e, 0xdb, 0xfe,
	0x8e, 0x00, 0x28, 0x8e, 0x42, 0x75, 0x5a, 0xc8, 0xdb, 0xbb, 0x9a, 0xe8, 0x7c, 0x0d, 0x79, 0x03,
	0xe9, 0x28, 0xc5, 0x77, 0xac, 0xb4, 0xfd, 0xe7, 0x7a, 0xf7, 0x34, 0xa4, 0x7d, 0x76, 0x61, 0x1a,
	0x78, 0xf2, 0x1c, 0x96, 0x87, 0xa9, 0xef, 0x3f, 0x1a, 0x7d, 0x73, 0xab, 0x6f, 0x65, 0xf0, 0x37,
	0x2b, 0x30, 0x83, 0xf4, 0x88, 0x7b, 0xd5, 0x69, 0xcc, 0x62, 0xcb, 0x05, 0x68, 0x64, 0x08, 0x73,
	0xe7, 0x6c, 0xf1, 0xbf, 0x0c, 0x6b, 0x04, 0x75, 0xdc, 0xf6, 0x06, 0x52, 0xff, 0x3f, 0x63, 0x8d,
	0xb5, 0xc0, 0x5b, 0x63, 0xbb, 0x13, 0x88, 0xa7, 0xa9, 0xab, 0xbd, 0x2d, 0xf6, 0x7c, 0x99, 0x25,
	0x7f, 0x66, 0xb0, 0x7a, 0xd7, 0x32, 0xde, 0xa0, 0x42, 0x4d, 0x0a, 0x98, 0x73, 0x3c, 0x7a, 0x3f,
	0xae, 0x42, 0xc1, 0x89, 0x2f, 0xdf, 0x23, 0xab, 0xa9, 0x49, 0x31, 0x99, 0xa2, 0xaf, 0xbd, 0x25,
	0x67, 0x33, 0x45, 0x7f, 0x99, 0x0c, 0x72, 0x05, 0x49, 0xd5, 0x0b, 0x85, 0xb5, 0x35, 0x67, 0x49,
	0x3d, 0x32, 0xae, 0x55, 0x12, 0x99, 0xc6, 0xda, 0xba, 0xb3, 0xa4, 0x01, 0x9a, 0x8a, 0x83, 0x90,
	0x15, 0xd6, 0x59, 0xe2, 0x2a, 0x1c, 0x22, 0x05, 0xa4, 0x66, 0xe4, 0x3a, 0xad, 0xb2, 0x74, 0x33,
	0x2f, 0xd6, 0xbb, 0x27, 0x27, 0x71, 0x36, 0x4c, 0x03, 0x9d, 0xdf, 0x40, 0x6c, 0x54, 0x3e, 0x38,
	0x6a, 0xc1, 0xe3, 0xd9, 0x3f, 0x1e, 0xbf, 0x86, 0x44, 0x33, 0xd9, 0xa0, 0xf6, 0x37, 0xb9, 0xdf,
	0xd8, 0xb3, 0xdb, 0x5f, 0x11, 0x24, 0x2e, 0x64, 0xda, 0xb4, 0x4c, 0xb5, 0xa1, 0xb5, 0x39, 0x93,
	0x57, 0x90, 0xb0, 0x49, 0xb7, 0x42, 0x7a, 0xf3, 0xee, 0xbd, 0xb7, 0x27, 0xcd, 0xcd, 0x07, 0x54,
	0x8a, 0x35, 0x61, 0x96, 0x02, 0x24, 0x25, 0xc4, 0x35, 0xf3, 0x63, 0xb4, 0xde, 0xe5, 0xa5, 0xdb,
	0xa9, 0x32, 0xec, 0x54, 0xf9, 0x25, 0xec, 0x14, 0xb5, 0x79, 0x77, 0x0f, 0xb0, 0x38, 0xff, 0x00,
	0x7b, 0xb7, 0xaf, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x10, 0x1a, 0x9b, 0x36, 0xd3, 0x03,
	0x00, 0x00,
}

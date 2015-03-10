// Code generated by protoc-gen-go.
// source: hologram.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

Copyright 2014 AdRoll, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

It is generated from these files:
	hologram.proto

It has these top-level messages:
	Message
	Ping
	ServerRequest
	AssumeRole
	GetUserCredentials
	AddSSHKey
	SSHChallengeResponse
	MFATokenResponse
	ServerResponse
	SSHChallenge
	SSHVerificationFailure
	STSCredentials
	MFATokenRequest
	AgentRequest
	AgentResponse
	Success
	Failure
*/
package protocol

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Message_Source int32

const (
	Message_OTHER           Message_Source = 0
	Message_HOLOGRAM_SERVER Message_Source = 1
	Message_HOLOGRAM_CLIENT Message_Source = 2
	Message_HOLOGRAM_CLI    Message_Source = 3
)

var Message_Source_name = map[int32]string{
	0: "OTHER",
	1: "HOLOGRAM_SERVER",
	2: "HOLOGRAM_CLIENT",
	3: "HOLOGRAM_CLI",
}
var Message_Source_value = map[string]int32{
	"OTHER":           0,
	"HOLOGRAM_SERVER": 1,
	"HOLOGRAM_CLIENT": 2,
	"HOLOGRAM_CLI":    3,
}

func (x Message_Source) Enum() *Message_Source {
	p := new(Message_Source)
	*p = x
	return p
}
func (x Message_Source) String() string {
	return proto.EnumName(Message_Source_name, int32(x))
}
func (x *Message_Source) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_Source_value, data, "Message_Source")
	if err != nil {
		return err
	}
	*x = Message_Source(value)
	return nil
}

type Ping_RequestResponse int32

const (
	Ping_REQUEST  Ping_RequestResponse = 1
	Ping_RESPONSE Ping_RequestResponse = 2
)

var Ping_RequestResponse_name = map[int32]string{
	1: "REQUEST",
	2: "RESPONSE",
}
var Ping_RequestResponse_value = map[string]int32{
	"REQUEST":  1,
	"RESPONSE": 2,
}

func (x Ping_RequestResponse) Enum() *Ping_RequestResponse {
	p := new(Ping_RequestResponse)
	*p = x
	return p
}
func (x Ping_RequestResponse) String() string {
	return proto.EnumName(Ping_RequestResponse_name, int32(x))
}
func (x *Ping_RequestResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Ping_RequestResponse_value, data, "Ping_RequestResponse")
	if err != nil {
		return err
	}
	*x = Ping_RequestResponse(value)
	return nil
}

type Message struct {
	Error *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// This is useful for statistics and debugging
	Source           *Message_Source `protobuf:"varint,2,opt,name=source,enum=protocol.Message_Source,def=0" json:"source,omitempty"`
	Ping             *Ping           `protobuf:"bytes,5,opt,name=ping" json:"ping,omitempty"`
	ServerRequest    *ServerRequest  `protobuf:"bytes,6,opt,name=serverRequest" json:"serverRequest,omitempty"`
	ServerResponse   *ServerResponse `protobuf:"bytes,7,opt,name=serverResponse" json:"serverResponse,omitempty"`
	AgentRequest     *AgentRequest   `protobuf:"bytes,8,opt,name=agentRequest" json:"agentRequest,omitempty"`
	AgentResponse    *AgentResponse  `protobuf:"bytes,9,opt,name=agentResponse" json:"agentResponse,omitempty"`
	Success          *Success        `protobuf:"bytes,10,opt,name=success" json:"success,omitempty"`
	Failure          *Failure        `protobuf:"bytes,11,opt,name=failure" json:"failure,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

const Default_Message_Source Message_Source = Message_OTHER

func (m *Message) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *Message) GetSource() Message_Source {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return Default_Message_Source
}

func (m *Message) GetPing() *Ping {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *Message) GetServerRequest() *ServerRequest {
	if m != nil {
		return m.ServerRequest
	}
	return nil
}

func (m *Message) GetServerResponse() *ServerResponse {
	if m != nil {
		return m.ServerResponse
	}
	return nil
}

func (m *Message) GetAgentRequest() *AgentRequest {
	if m != nil {
		return m.AgentRequest
	}
	return nil
}

func (m *Message) GetAgentResponse() *AgentResponse {
	if m != nil {
		return m.AgentResponse
	}
	return nil
}

func (m *Message) GetSuccess() *Success {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *Message) GetFailure() *Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

type Ping struct {
	Type             *Ping_RequestResponse `protobuf:"varint,1,opt,name=type,enum=protocol.Ping_RequestResponse,def=1" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}

const Default_Ping_Type Ping_RequestResponse = Ping_REQUEST

func (m *Ping) GetType() Ping_RequestResponse {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Ping_Type
}

type ServerRequest struct {
	AssumeRole         *AssumeRole           `protobuf:"bytes,4,opt,name=assumeRole" json:"assumeRole,omitempty"`
	ChallengeResponse  *SSHChallengeResponse `protobuf:"bytes,5,opt,name=challengeResponse" json:"challengeResponse,omitempty"`
	TokenResponse      *MFATokenResponse     `protobuf:"bytes,6,opt,name=tokenResponse" json:"tokenResponse,omitempty"`
	GetUserCredentials *GetUserCredentials   `protobuf:"bytes,7,opt,name=getUserCredentials" json:"getUserCredentials,omitempty"`
	AddSSHkey          *AddSSHKey            `protobuf:"bytes,8,opt,name=addSSHkey" json:"addSSHkey,omitempty"`
	XXX_unrecognized   []byte                `json:"-"`
}

func (m *ServerRequest) Reset()         { *m = ServerRequest{} }
func (m *ServerRequest) String() string { return proto.CompactTextString(m) }
func (*ServerRequest) ProtoMessage()    {}

func (m *ServerRequest) GetAssumeRole() *AssumeRole {
	if m != nil {
		return m.AssumeRole
	}
	return nil
}

func (m *ServerRequest) GetChallengeResponse() *SSHChallengeResponse {
	if m != nil {
		return m.ChallengeResponse
	}
	return nil
}

func (m *ServerRequest) GetTokenResponse() *MFATokenResponse {
	if m != nil {
		return m.TokenResponse
	}
	return nil
}

func (m *ServerRequest) GetGetUserCredentials() *GetUserCredentials {
	if m != nil {
		return m.GetUserCredentials
	}
	return nil
}

func (m *ServerRequest) GetAddSSHkey() *AddSSHKey {
	if m != nil {
		return m.AddSSHkey
	}
	return nil
}

type AssumeRole struct {
	User             *string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Role             *string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AssumeRole) Reset()         { *m = AssumeRole{} }
func (m *AssumeRole) String() string { return proto.CompactTextString(m) }
func (*AssumeRole) ProtoMessage()    {}

func (m *AssumeRole) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *AssumeRole) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

type GetUserCredentials struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetUserCredentials) Reset()         { *m = GetUserCredentials{} }
func (m *GetUserCredentials) String() string { return proto.CompactTextString(m) }
func (*GetUserCredentials) ProtoMessage()    {}

type AddSSHKey struct {
	Username         *string `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	Passwordhash     *string `protobuf:"bytes,2,req,name=passwordhash" json:"passwordhash,omitempty"`
	Sshkeybytes      *string `protobuf:"bytes,3,req,name=sshkeybytes" json:"sshkeybytes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AddSSHKey) Reset()         { *m = AddSSHKey{} }
func (m *AddSSHKey) String() string { return proto.CompactTextString(m) }
func (*AddSSHKey) ProtoMessage()    {}

func (m *AddSSHKey) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *AddSSHKey) GetPasswordhash() string {
	if m != nil && m.Passwordhash != nil {
		return *m.Passwordhash
	}
	return ""
}

func (m *AddSSHKey) GetSshkeybytes() string {
	if m != nil && m.Sshkeybytes != nil {
		return *m.Sshkeybytes
	}
	return ""
}

type SSHChallengeResponse struct {
	Signature        []byte  `protobuf:"bytes,1,req,name=signature" json:"signature,omitempty"`
	Format           *string `protobuf:"bytes,2,req,name=format" json:"format,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SSHChallengeResponse) Reset()         { *m = SSHChallengeResponse{} }
func (m *SSHChallengeResponse) String() string { return proto.CompactTextString(m) }
func (*SSHChallengeResponse) ProtoMessage()    {}

func (m *SSHChallengeResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SSHChallengeResponse) GetFormat() string {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return ""
}

type MFATokenResponse struct {
	TokenValue       *string `protobuf:"bytes,1,opt,name=tokenValue" json:"tokenValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MFATokenResponse) Reset()         { *m = MFATokenResponse{} }
func (m *MFATokenResponse) String() string { return proto.CompactTextString(m) }
func (*MFATokenResponse) ProtoMessage()    {}

func (m *MFATokenResponse) GetTokenValue() string {
	if m != nil && m.TokenValue != nil {
		return *m.TokenValue
	}
	return ""
}

type ServerResponse struct {
	Challenge           *SSHChallenge           `protobuf:"bytes,4,opt,name=challenge" json:"challenge,omitempty"`
	VerificationFailure *SSHVerificationFailure `protobuf:"bytes,5,opt,name=verificationFailure" json:"verificationFailure,omitempty"`
	Credentials         *STSCredentials         `protobuf:"bytes,6,opt,name=credentials" json:"credentials,omitempty"`
	TokenRequest        *MFATokenRequest        `protobuf:"bytes,7,opt,name=tokenRequest" json:"tokenRequest,omitempty"`
	XXX_unrecognized    []byte                  `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}

func (m *ServerResponse) GetChallenge() *SSHChallenge {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *ServerResponse) GetVerificationFailure() *SSHVerificationFailure {
	if m != nil {
		return m.VerificationFailure
	}
	return nil
}

func (m *ServerResponse) GetCredentials() *STSCredentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *ServerResponse) GetTokenRequest() *MFATokenRequest {
	if m != nil {
		return m.TokenRequest
	}
	return nil
}

type SSHChallenge struct {
	Challenge        []byte `protobuf:"bytes,1,req,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SSHChallenge) Reset()         { *m = SSHChallenge{} }
func (m *SSHChallenge) String() string { return proto.CompactTextString(m) }
func (*SSHChallenge) ProtoMessage()    {}

func (m *SSHChallenge) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

type SSHVerificationFailure struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SSHVerificationFailure) Reset()         { *m = SSHVerificationFailure{} }
func (m *SSHVerificationFailure) String() string { return proto.CompactTextString(m) }
func (*SSHVerificationFailure) ProtoMessage()    {}

type STSCredentials struct {
	AccessKeyId      *string `protobuf:"bytes,1,req,name=accessKeyId" json:"accessKeyId,omitempty"`
	SecretAccessKey  *string `protobuf:"bytes,2,req,name=secretAccessKey" json:"secretAccessKey,omitempty"`
	AccessToken      *string `protobuf:"bytes,3,req,name=accessToken" json:"accessToken,omitempty"`
	Expiration       *int64  `protobuf:"varint,4,req,name=expiration" json:"expiration,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STSCredentials) Reset()         { *m = STSCredentials{} }
func (m *STSCredentials) String() string { return proto.CompactTextString(m) }
func (*STSCredentials) ProtoMessage()    {}

func (m *STSCredentials) GetAccessKeyId() string {
	if m != nil && m.AccessKeyId != nil {
		return *m.AccessKeyId
	}
	return ""
}

func (m *STSCredentials) GetSecretAccessKey() string {
	if m != nil && m.SecretAccessKey != nil {
		return *m.SecretAccessKey
	}
	return ""
}

func (m *STSCredentials) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *STSCredentials) GetExpiration() int64 {
	if m != nil && m.Expiration != nil {
		return *m.Expiration
	}
	return 0
}

type MFATokenRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MFATokenRequest) Reset()         { *m = MFATokenRequest{} }
func (m *MFATokenRequest) String() string { return proto.CompactTextString(m) }
func (*MFATokenRequest) ProtoMessage()    {}

type AgentRequest struct {
	SshAgentSock       *string             `protobuf:"bytes,2,opt,name=sshAgentSock" json:"sshAgentSock,omitempty"`
	AssumeRole         *AssumeRole         `protobuf:"bytes,3,opt,name=assumeRole" json:"assumeRole,omitempty"`
	GetUserCredentials *GetUserCredentials `protobuf:"bytes,4,opt,name=getUserCredentials" json:"getUserCredentials,omitempty"`
	// sshKeyFile should be sent along if the CLI cannot determine
	// how to communicate with the user's SSH agent.
	SshKeyFile       []byte `protobuf:"bytes,5,opt,name=sshKeyFile" json:"sshKeyFile,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AgentRequest) Reset()         { *m = AgentRequest{} }
func (m *AgentRequest) String() string { return proto.CompactTextString(m) }
func (*AgentRequest) ProtoMessage()    {}

func (m *AgentRequest) GetSshAgentSock() string {
	if m != nil && m.SshAgentSock != nil {
		return *m.SshAgentSock
	}
	return ""
}

func (m *AgentRequest) GetAssumeRole() *AssumeRole {
	if m != nil {
		return m.AssumeRole
	}
	return nil
}

func (m *AgentRequest) GetGetUserCredentials() *GetUserCredentials {
	if m != nil {
		return m.GetUserCredentials
	}
	return nil
}

func (m *AgentRequest) GetSshKeyFile() []byte {
	if m != nil {
		return m.SshKeyFile
	}
	return nil
}

type AgentResponse struct {
	Success          *Success `protobuf:"bytes,2,opt,name=success" json:"success,omitempty"`
	Failure          *Failure `protobuf:"bytes,3,opt,name=failure" json:"failure,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AgentResponse) Reset()         { *m = AgentResponse{} }
func (m *AgentResponse) String() string { return proto.CompactTextString(m) }
func (*AgentResponse) ProtoMessage()    {}

func (m *AgentResponse) GetSuccess() *Success {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *AgentResponse) GetFailure() *Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

type Success struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}

type Failure struct {
	ErrorMessage     *string `protobuf:"bytes,1,opt,name=errorMessage" json:"errorMessage,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}

func (m *Failure) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.Message_Source", Message_Source_name, Message_Source_value)
	proto.RegisterEnum("protocol.Ping_RequestResponse", Ping_RequestResponse_name, Ping_RequestResponse_value)
}

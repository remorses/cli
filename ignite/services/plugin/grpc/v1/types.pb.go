// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: ignite/services/plugin/grpc/v1/types.proto

package v1

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

// FlagType represents the flag type.
type FlagType int32

const (
	FlagType_FLAG_TYPE_UNSPECIFIED  FlagType = 0
	FlagType_FLAG_TYPE_STRING       FlagType = 1
	FlagType_FLAG_TYPE_INT          FlagType = 2
	FlagType_FLAG_TYPE_UINT         FlagType = 3
	FlagType_FLAG_TYPE_INT64        FlagType = 4
	FlagType_FLAG_TYPE_UINT64       FlagType = 5
	FlagType_FLAG_TYPE_BOOL         FlagType = 6
	FlagType_FLAG_TYPE_STRING_SLICE FlagType = 7
)

// Enum value maps for FlagType.
var (
	FlagType_name = map[int32]string{
		0: "FLAG_TYPE_UNSPECIFIED",
		1: "FLAG_TYPE_STRING",
		2: "FLAG_TYPE_INT",
		3: "FLAG_TYPE_UINT",
		4: "FLAG_TYPE_INT64",
		5: "FLAG_TYPE_UINT64",
		6: "FLAG_TYPE_BOOL",
		7: "FLAG_TYPE_STRING_SLICE",
	}
	FlagType_value = map[string]int32{
		"FLAG_TYPE_UNSPECIFIED":  0,
		"FLAG_TYPE_STRING":       1,
		"FLAG_TYPE_INT":          2,
		"FLAG_TYPE_UINT":         3,
		"FLAG_TYPE_INT64":        4,
		"FLAG_TYPE_UINT64":       5,
		"FLAG_TYPE_BOOL":         6,
		"FLAG_TYPE_STRING_SLICE": 7,
	}
)

func (x FlagType) Enum() *FlagType {
	p := new(FlagType)
	*p = x
	return p
}

func (x FlagType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlagType) Descriptor() protoreflect.EnumDescriptor {
	return file_ignite_services_plugin_grpc_v1_types_proto_enumTypes[0].Descriptor()
}

func (FlagType) Type() protoreflect.EnumType {
	return &file_ignite_services_plugin_grpc_v1_types_proto_enumTypes[0]
}

func (x FlagType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlagType.Descriptor instead.
func (FlagType) EnumDescriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{0}
}

// ExecutedCommand represents a plugin command under execution.
type ExecutedCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use is the one-line usage message.
	Use string `protobuf:"bytes,1,opt,name=use,proto3" json:"use,omitempty"`
	// Path contains the command path, e.g. `ignite scaffold foo`.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Args are the command arguments.
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Full list of args taken from the command line.
	OsArgs []string `protobuf:"bytes,4,rep,name=os_args,json=osArgs,proto3" json:"os_args,omitempty"`
	// With contains the plugin config parameters.
	With map[string]string `protobuf:"bytes,5,rep,name=with,proto3" json:"with,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExecutedCommand) Reset() {
	*x = ExecutedCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutedCommand) ProtoMessage() {}

func (x *ExecutedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutedCommand.ProtoReflect.Descriptor instead.
func (*ExecutedCommand) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutedCommand) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *ExecutedCommand) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ExecutedCommand) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ExecutedCommand) GetOsArgs() []string {
	if x != nil {
		return x.OsArgs
	}
	return nil
}

func (x *ExecutedCommand) GetWith() map[string]string {
	if x != nil {
		return x.With
	}
	return nil
}

// ExecutedHook represents a plugin hook under execution.
type ExecutedHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hook is a copy of the original Hook defined in the Manifest.
	Hook *Hook `protobuf:"bytes,1,opt,name=hook,proto3" json:"hook,omitempty"`
	// ExecutedCommand gives access to the command attached by the hook.
	ExecutedCommand *ExecutedCommand `protobuf:"bytes,2,opt,name=executed_command,json=executedCommand,proto3" json:"executed_command,omitempty"`
}

func (x *ExecutedHook) Reset() {
	*x = ExecutedHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutedHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutedHook) ProtoMessage() {}

func (x *ExecutedHook) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutedHook.ProtoReflect.Descriptor instead.
func (*ExecutedHook) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutedHook) GetHook() *Hook {
	if x != nil {
		return x.Hook
	}
	return nil
}

func (x *ExecutedHook) GetExecutedCommand() *ExecutedCommand {
	if x != nil {
		return x.ExecutedCommand
	}
	return nil
}

// Manifest represents the plugin behavior.
type Manifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Commands contains the commands that will be added to the list of ignite commands.
	// Each commands are independent, for nested commands use the inner Commands field.
	SharedHost bool `protobuf:"varint,2,opt,name=shared_host,json=sharedHost,proto3" json:"shared_host,omitempty"`
	// Hooks contains the hooks that will be attached to the existing ignite commands.
	Commands []*Command `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
	// Enables sharing a single plugin server across all running instances of a plugin.
	// Useful if a plugin adds or extends long running commands.
	//
	// Example: if a plugin defines a hook on `ignite chain serve`, a plugin server is
	// instanciated when the command is run. Now if you want to interact with that instance
	// from commands defined in that plugin, you need to enable shared host, or else the
	// commands will just instantiate separate plugin servers.
	//
	// When enabled, all plugins of the same path loaded from the same configuration will
	// attach it's RPC client to a an existing RPC server.
	//
	// If a plugin instance has no other running plugin servers, it will create one and it
	// will be the host.
	Hooks []*Hook `protobuf:"bytes,4,rep,name=hooks,proto3" json:"hooks,omitempty"`
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Manifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manifest) GetSharedHost() bool {
	if x != nil {
		return x.SharedHost
	}
	return false
}

func (x *Manifest) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *Manifest) GetHooks() []*Hook {
	if x != nil {
		return x.Hooks
	}
	return nil
}

// Command represents a plugin command.
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use is the one-line usage message.
	//
	// Recommended syntax is as follow:
	//
	//	[ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//	... indicates that you can specify multiple values for the previous argument.
	//	|   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//	    argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//	{ } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//	    optional, they are enclosed in brackets ([ ]).
	//
	// Example: add [-F file | -D dir]... [-f format] profile
	Use string `protobuf:"bytes,1,opt,name=use,proto3" json:"use,omitempty"`
	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string `protobuf:"bytes,2,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// Short is the short description shown in the 'help' output.
	Short string `protobuf:"bytes,3,opt,name=short,proto3" json:"short,omitempty"`
	// Long is the long message shown in the 'help <this-command>' output.
	Long string `protobuf:"bytes,4,opt,name=long,proto3" json:"long,omitempty"`
	// Hidden defines, if this command is hidden and should NOT show up in the list of available commands.
	Hidden bool `protobuf:"varint,5,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Flags holds the list of command flags.
	Flags []*Flag `protobuf:"bytes,6,rep,name=flags,proto3" json:"flags,omitempty"`
	// Indicates where the command should be placed.
	// For instance `ignite scaffold` will place the command at the `scaffold` command.
	// An empty value is interpreted as `ignite` (==root).
	PlaceCommandUnder string `protobuf:"bytes,7,opt,name=place_command_under,json=placeCommandUnder,proto3" json:"place_command_under,omitempty"`
	// List of sub commands.
	Commands []*Command `protobuf:"bytes,8,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *Command) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *Command) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Command) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *Command) GetLong() string {
	if x != nil {
		return x.Long
	}
	return ""
}

func (x *Command) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *Command) GetFlags() []*Flag {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Command) GetPlaceCommandUnder() string {
	if x != nil {
		return x.PlaceCommandUnder
	}
	return ""
}

func (x *Command) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

// Flag represents of a command line flag.
type Flag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name as it appears in the command line.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// One letter abbreviation of the flag.
	Shorthand string `protobuf:"bytes,2,opt,name=shorthand,proto3" json:"shorthand,omitempty"`
	// Help message.
	Usage string `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
	// Default flag value.
	DefaultValue string `protobuf:"bytes,4,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Flag type.
	Type FlagType `protobuf:"varint,5,opt,name=type,proto3,enum=ignite.services.plugin.grpc.v1.FlagType" json:"type,omitempty"`
	// Flag value.
	Value string `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// Indicates wether or not the flag is propagated on children commands.
	Persistent bool `protobuf:"varint,7,opt,name=persistent,proto3" json:"persistent,omitempty"`
}

func (x *Flag) Reset() {
	*x = Flag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag) ProtoMessage() {}

func (x *Flag) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag.ProtoReflect.Descriptor instead.
func (*Flag) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{4}
}

func (x *Flag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flag) GetShorthand() string {
	if x != nil {
		return x.Shorthand
	}
	return ""
}

func (x *Flag) GetUsage() string {
	if x != nil {
		return x.Usage
	}
	return ""
}

func (x *Flag) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

func (x *Flag) GetType() FlagType {
	if x != nil {
		return x.Type
	}
	return FlagType_FLAG_TYPE_UNSPECIFIED
}

func (x *Flag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Flag) GetPersistent() bool {
	if x != nil {
		return x.Persistent
	}
	return false
}

// Hook represents a user defined action within a plugin.
type Hook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the hook for the client to invoke the correct hook.
	// It must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates the command where to register the hooks.
	PlaceHookOn string `protobuf:"bytes,2,opt,name=place_hook_on,json=placeHookOn,proto3" json:"place_hook_on,omitempty"`
}

func (x *Hook) Reset() {
	*x = Hook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hook) ProtoMessage() {}

func (x *Hook) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hook.ProtoReflect.Descriptor instead.
func (*Hook) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP(), []int{5}
}

func (x *Hook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hook) GetPlaceHookOn() string {
	if x != nil {
		return x.PlaceHookOn
	}
	return ""
}

var File_ignite_services_plugin_grpc_v1_types_proto protoreflect.FileDescriptor

var file_ignite_services_plugin_grpc_v1_types_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x69, 0x67,
	0x6e, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x22, 0xec, 0x01, 0x0a,
	0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73,
	0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x4d, 0x0a, 0x04, 0x77, 0x69, 0x74, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x77, 0x69,
	0x74, 0x68, 0x1a, 0x37, 0x0a, 0x09, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa4, 0x01, 0x0a, 0x0c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x38, 0x0a, 0x04,
	0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x67, 0x6e,
	0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x6f, 0x6b,
	0x52, 0x04, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x5a, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x68, 0x6f, 0x6f,
	0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x67, 0x6e, 0x69, 0x74,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x05,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12,
	0x3a, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x75, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x22, 0xe7, 0x01, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x04, 0x48, 0x6f,
	0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f,
	0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x4f, 0x6e, 0x2a, 0xbd, 0x01, 0x0a, 0x08, 0x46,
	0x6c, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x4c, 0x41, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4c, 0x41, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x46,
	0x4c, 0x41, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54,
	0x36, 0x34, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4c,
	0x41, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x06, 0x12, 0x1a,
	0x0a, 0x16, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x4c, 0x49, 0x43, 0x45, 0x10, 0x07, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2f,
	0x63, 0x6c, 0x69, 0x2f, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ignite_services_plugin_grpc_v1_types_proto_rawDescOnce sync.Once
	file_ignite_services_plugin_grpc_v1_types_proto_rawDescData = file_ignite_services_plugin_grpc_v1_types_proto_rawDesc
)

func file_ignite_services_plugin_grpc_v1_types_proto_rawDescGZIP() []byte {
	file_ignite_services_plugin_grpc_v1_types_proto_rawDescOnce.Do(func() {
		file_ignite_services_plugin_grpc_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_ignite_services_plugin_grpc_v1_types_proto_rawDescData)
	})
	return file_ignite_services_plugin_grpc_v1_types_proto_rawDescData
}

var file_ignite_services_plugin_grpc_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ignite_services_plugin_grpc_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ignite_services_plugin_grpc_v1_types_proto_goTypes = []interface{}{
	(FlagType)(0),           // 0: ignite.services.plugin.grpc.v1.FlagType
	(*ExecutedCommand)(nil), // 1: ignite.services.plugin.grpc.v1.ExecutedCommand
	(*ExecutedHook)(nil),    // 2: ignite.services.plugin.grpc.v1.ExecutedHook
	(*Manifest)(nil),        // 3: ignite.services.plugin.grpc.v1.Manifest
	(*Command)(nil),         // 4: ignite.services.plugin.grpc.v1.Command
	(*Flag)(nil),            // 5: ignite.services.plugin.grpc.v1.Flag
	(*Hook)(nil),            // 6: ignite.services.plugin.grpc.v1.Hook
	nil,                     // 7: ignite.services.plugin.grpc.v1.ExecutedCommand.WithEntry
}
var file_ignite_services_plugin_grpc_v1_types_proto_depIdxs = []int32{
	7, // 0: ignite.services.plugin.grpc.v1.ExecutedCommand.with:type_name -> ignite.services.plugin.grpc.v1.ExecutedCommand.WithEntry
	6, // 1: ignite.services.plugin.grpc.v1.ExecutedHook.hook:type_name -> ignite.services.plugin.grpc.v1.Hook
	1, // 2: ignite.services.plugin.grpc.v1.ExecutedHook.executed_command:type_name -> ignite.services.plugin.grpc.v1.ExecutedCommand
	4, // 3: ignite.services.plugin.grpc.v1.Manifest.commands:type_name -> ignite.services.plugin.grpc.v1.Command
	6, // 4: ignite.services.plugin.grpc.v1.Manifest.hooks:type_name -> ignite.services.plugin.grpc.v1.Hook
	5, // 5: ignite.services.plugin.grpc.v1.Command.flags:type_name -> ignite.services.plugin.grpc.v1.Flag
	4, // 6: ignite.services.plugin.grpc.v1.Command.commands:type_name -> ignite.services.plugin.grpc.v1.Command
	0, // 7: ignite.services.plugin.grpc.v1.Flag.type:type_name -> ignite.services.plugin.grpc.v1.FlagType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ignite_services_plugin_grpc_v1_types_proto_init() }
func file_ignite_services_plugin_grpc_v1_types_proto_init() {
	if File_ignite_services_plugin_grpc_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutedCommand); i {
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
		file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutedHook); i {
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
		file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest); i {
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
		file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flag); i {
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
		file_ignite_services_plugin_grpc_v1_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hook); i {
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
			RawDescriptor: file_ignite_services_plugin_grpc_v1_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ignite_services_plugin_grpc_v1_types_proto_goTypes,
		DependencyIndexes: file_ignite_services_plugin_grpc_v1_types_proto_depIdxs,
		EnumInfos:         file_ignite_services_plugin_grpc_v1_types_proto_enumTypes,
		MessageInfos:      file_ignite_services_plugin_grpc_v1_types_proto_msgTypes,
	}.Build()
	File_ignite_services_plugin_grpc_v1_types_proto = out.File
	file_ignite_services_plugin_grpc_v1_types_proto_rawDesc = nil
	file_ignite_services_plugin_grpc_v1_types_proto_goTypes = nil
	file_ignite_services_plugin_grpc_v1_types_proto_depIdxs = nil
}

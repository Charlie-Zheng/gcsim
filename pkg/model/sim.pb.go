// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: protos/model/sim.proto

package model

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

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name,omitempty"`
	Element       string            `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty" bson:"element,omitempty"`
	Level         int32             `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty" bson:"level,omitempty"`
	MaxLevel      int32             `protobuf:"varint,4,opt,name=max_level,proto3" json:"max_level,omitempty" bson:"max_level,omitempty"`
	Cons          int32             `protobuf:"varint,5,opt,name=cons,proto3" json:"cons,omitempty" bson:"cons,omitempty"`
	Weapon        *Weapon           `protobuf:"bytes,6,opt,name=weapon,proto3" json:"weapon,omitempty" bson:"weapon,omitempty"`
	Talents       *CharacterTalents `protobuf:"bytes,7,opt,name=talents,proto3" json:"talents,omitempty" bson:"talents,omitempty"`
	Sets          map[string]int32  `protobuf:"bytes,8,rep,name=sets,proto3" json:"sets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" bson:"sets,omitempty"`
	Stats         []float64         `protobuf:"fixed64,9,rep,packed,name=stats,proto3" json:"stats,omitempty" bson:"stats,omitempty"`
	SnapshotStats []float64         `protobuf:"fixed64,10,rep,packed,name=snapshot_stats,json=snapshot,proto3" json:"snapshot_stats,omitempty" bson:"snapshot,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

func (x *Character) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Character) GetMaxLevel() int32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *Character) GetCons() int32 {
	if x != nil {
		return x.Cons
	}
	return 0
}

func (x *Character) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

func (x *Character) GetTalents() *CharacterTalents {
	if x != nil {
		return x.Talents
	}
	return nil
}

func (x *Character) GetSets() map[string]int32 {
	if x != nil {
		return x.Sets
	}
	return nil
}

func (x *Character) GetStats() []float64 {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Character) GetSnapshotStats() []float64 {
	if x != nil {
		return x.SnapshotStats
	}
	return nil
}

type CharacterTalents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attack int32 `protobuf:"varint,1,opt,name=attack,proto3" json:"attack,omitempty" bson:"attack,omitempty"`
	Skill  int32 `protobuf:"varint,2,opt,name=skill,proto3" json:"skill,omitempty" bson:"skill,omitempty"`
	Burst  int32 `protobuf:"varint,3,opt,name=burst,proto3" json:"burst,omitempty" bson:"burst,omitempty"`
}

func (x *CharacterTalents) Reset() {
	*x = CharacterTalents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterTalents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterTalents) ProtoMessage() {}

func (x *CharacterTalents) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterTalents.ProtoReflect.Descriptor instead.
func (*CharacterTalents) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{1}
}

func (x *CharacterTalents) GetAttack() int32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *CharacterTalents) GetSkill() int32 {
	if x != nil {
		return x.Skill
	}
	return 0
}

func (x *CharacterTalents) GetBurst() int32 {
	if x != nil {
		return x.Burst
	}
	return 0
}

type Weapon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name,omitempty"`
	Refine   int32  `protobuf:"varint,2,opt,name=refine,proto3" json:"refine,omitempty" bson:"refine,omitempty"`
	Level    int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty" bson:"level,omitempty"`
	MaxLevel int32  `protobuf:"varint,4,opt,name=max_level,proto3" json:"max_level,omitempty" bson:"max_level,omitempty"`
}

func (x *Weapon) Reset() {
	*x = Weapon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weapon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weapon) ProtoMessage() {}

func (x *Weapon) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weapon.ProtoReflect.Descriptor instead.
func (*Weapon) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{2}
}

func (x *Weapon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Weapon) GetRefine() int32 {
	if x != nil {
		return x.Refine
	}
	return 0
}

func (x *Weapon) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Weapon) GetMaxLevel() int32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

type Enemy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level                 int32              `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty" bson:"level,omitempty"`
	HP                    float64            `protobuf:"fixed64,2,opt,name=HP,json=hp,proto3" json:"HP,omitempty" bson:"hp,omitempty"`
	Resist                map[string]float64 `protobuf:"bytes,3,rep,name=resist,proto3" json:"resist,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3" bson:"resist,omitempty"`
	Pos                   *Coord             `protobuf:"bytes,4,opt,name=pos,json=position,proto3" json:"pos,omitempty" bson:"position,omitempty"`
	ParticleDropThreshold float64            `protobuf:"fixed64,5,opt,name=particle_drop_threshold,proto3" json:"particle_drop_threshold,omitempty" bson:"particle_drop_threshold,omitempty"`
	ParticleDropCount     float64            `protobuf:"fixed64,6,opt,name=particle_drop_count,proto3" json:"particle_drop_count,omitempty" bson:"particle_drop_count,omitempty"`
	ParticleElement       string             `protobuf:"bytes,7,opt,name=particle_element,proto3" json:"particle_element,omitempty" bson:"particle_element,omitempty"`
}

func (x *Enemy) Reset() {
	*x = Enemy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enemy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enemy) ProtoMessage() {}

func (x *Enemy) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enemy.ProtoReflect.Descriptor instead.
func (*Enemy) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{3}
}

func (x *Enemy) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Enemy) GetHP() float64 {
	if x != nil {
		return x.HP
	}
	return 0
}

func (x *Enemy) GetResist() map[string]float64 {
	if x != nil {
		return x.Resist
	}
	return nil
}

func (x *Enemy) GetPos() *Coord {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *Enemy) GetParticleDropThreshold() float64 {
	if x != nil {
		return x.ParticleDropThreshold
	}
	return 0
}

func (x *Enemy) GetParticleDropCount() float64 {
	if x != nil {
		return x.ParticleDropCount
	}
	return 0
}

func (x *Enemy) GetParticleElement() string {
	if x != nil {
		return x.ParticleElement
	}
	return ""
}

type Coord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty" bson:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty" bson:"y,omitempty"`
	R float64 `protobuf:"fixed64,3,opt,name=r,proto3" json:"r,omitempty" bson:"r,omitempty"`
}

func (x *Coord) Reset() {
	*x = Coord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coord) ProtoMessage() {}

func (x *Coord) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coord.ProtoReflect.Descriptor instead.
func (*Coord) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{4}
}

func (x *Coord) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coord) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Coord) GetR() float64 {
	if x != nil {
		return x.R
	}
	return 0
}

type SimulatorSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration          float64 `protobuf:"fixed64,1,opt,name=duration,proto3" json:"duration,omitempty" bson:"duration,omitempty"`
	DamageMode        bool    `protobuf:"varint,2,opt,name=damage_mode,proto3" json:"damage_mode,omitempty" bson:"damage_mode,omitempty"`
	EnableHitlag      bool    `protobuf:"varint,3,opt,name=enable_hitlag,proto3" json:"enable_hitlag,omitempty" bson:"enable_hitlag,omitempty"`
	DefHalt           bool    `protobuf:"varint,4,opt,name=def_halt,proto3" json:"def_halt,omitempty" bson:"def_halt,omitempty"`
	IgnoreBurstEnergy bool    `protobuf:"varint,5,opt,name=ignore_burst_energy,proto3" json:"ignore_burst_energy,omitempty" bson:"ignore_burst_energy,omitempty"`
	NumberOfWorkers   uint32  `protobuf:"varint,6,opt,name=number_of_workers,proto3" json:"number_of_workers,omitempty" bson:"number_of_workers,omitempty"`
	Iterations        uint32  `protobuf:"varint,7,opt,name=iterations,proto3" json:"iterations,omitempty" bson:"iterations,omitempty"`
	Delays            *Delays `protobuf:"bytes,8,opt,name=delays,proto3" json:"delays,omitempty" bson:"delays,omitempty"`
}

func (x *SimulatorSettings) Reset() {
	*x = SimulatorSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulatorSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorSettings) ProtoMessage() {}

func (x *SimulatorSettings) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorSettings.ProtoReflect.Descriptor instead.
func (*SimulatorSettings) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{5}
}

func (x *SimulatorSettings) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *SimulatorSettings) GetDamageMode() bool {
	if x != nil {
		return x.DamageMode
	}
	return false
}

func (x *SimulatorSettings) GetEnableHitlag() bool {
	if x != nil {
		return x.EnableHitlag
	}
	return false
}

func (x *SimulatorSettings) GetDefHalt() bool {
	if x != nil {
		return x.DefHalt
	}
	return false
}

func (x *SimulatorSettings) GetIgnoreBurstEnergy() bool {
	if x != nil {
		return x.IgnoreBurstEnergy
	}
	return false
}

func (x *SimulatorSettings) GetNumberOfWorkers() uint32 {
	if x != nil {
		return x.NumberOfWorkers
	}
	return 0
}

func (x *SimulatorSettings) GetIterations() uint32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *SimulatorSettings) GetDelays() *Delays {
	if x != nil {
		return x.Delays
	}
	return nil
}

type Delays struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill  int32 `protobuf:"varint,1,opt,name=skill,proto3" json:"skill,omitempty" bson:"skill,omitempty"`
	Burst  int32 `protobuf:"varint,2,opt,name=burst,proto3" json:"burst,omitempty" bson:"burst,omitempty"`
	Attack int32 `protobuf:"varint,3,opt,name=attack,proto3" json:"attack,omitempty" bson:"attack,omitempty"`
	Charge int32 `protobuf:"varint,4,opt,name=charge,proto3" json:"charge,omitempty" bson:"charge,omitempty"`
	Aim    int32 `protobuf:"varint,5,opt,name=aim,proto3" json:"aim,omitempty" bson:"aim,omitempty"`
	Dash   int32 `protobuf:"varint,6,opt,name=dash,proto3" json:"dash,omitempty" bson:"dash,omitempty"`
	Jump   int32 `protobuf:"varint,7,opt,name=jump,proto3" json:"jump,omitempty" bson:"jump,omitempty"`
	Swap   int32 `protobuf:"varint,8,opt,name=swap,proto3" json:"swap,omitempty" bson:"swap,omitempty"`
}

func (x *Delays) Reset() {
	*x = Delays{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delays) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delays) ProtoMessage() {}

func (x *Delays) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delays.ProtoReflect.Descriptor instead.
func (*Delays) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{6}
}

func (x *Delays) GetSkill() int32 {
	if x != nil {
		return x.Skill
	}
	return 0
}

func (x *Delays) GetBurst() int32 {
	if x != nil {
		return x.Burst
	}
	return 0
}

func (x *Delays) GetAttack() int32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *Delays) GetCharge() int32 {
	if x != nil {
		return x.Charge
	}
	return 0
}

func (x *Delays) GetAim() int32 {
	if x != nil {
		return x.Aim
	}
	return 0
}

func (x *Delays) GetDash() int32 {
	if x != nil {
		return x.Dash
	}
	return 0
}

func (x *Delays) GetJump() int32 {
	if x != nil {
		return x.Jump
	}
	return 0
}

func (x *Delays) GetSwap() int32 {
	if x != nil {
		return x.Swap
	}
	return 0
}

type EnergySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active         bool  `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty" bson:"active,omitempty"`
	Once           bool  `protobuf:"varint,2,opt,name=once,proto3" json:"once,omitempty" bson:"once,omitempty"`
	Start          int32 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty" bson:"start,omitempty"`
	End            int32 `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty" bson:"end,omitempty"`
	Amount         int32 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty" bson:"amount,omitempty"`
	LastEnergyDrop int32 `protobuf:"varint,6,opt,name=last_energy_drop,proto3" json:"last_energy_drop,omitempty" bson:"last_energy_drop,omitempty"`
}

func (x *EnergySettings) Reset() {
	*x = EnergySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_sim_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnergySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnergySettings) ProtoMessage() {}

func (x *EnergySettings) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_sim_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnergySettings.ProtoReflect.Descriptor instead.
func (*EnergySettings) Descriptor() ([]byte, []int) {
	return file_protos_model_sim_proto_rawDescGZIP(), []int{7}
}

func (x *EnergySettings) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *EnergySettings) GetOnce() bool {
	if x != nil {
		return x.Once
	}
	return false
}

func (x *EnergySettings) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *EnergySettings) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *EnergySettings) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *EnergySettings) GetLastEnergyDrop() int32 {
	if x != nil {
		return x.LastEnergyDrop
	}
	return 0
}

var File_protos_model_sim_proto protoreflect.FileDescriptor

var file_protos_model_sim_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0xfc, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x57, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x74, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x61, 0x6c,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x07, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x0a,
	0x04, 0x73, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x73, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x01, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56,
	0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x62, 0x75, 0x72, 0x73, 0x74, 0x22, 0x68, 0x0a, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x22, 0xd7, 0x02, 0x0a, 0x05, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x48, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x68, 0x70,
	0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x2e, 0x52,
	0x65, 0x73, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x69,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x39, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x31, 0x0a, 0x05, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12,
	0x0c, 0x0a, 0x01, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x72, 0x22, 0xba, 0x02,
	0x0a, 0x11, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x6c,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x68, 0x69, 0x74, 0x6c, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x5f, 0x68,
	0x61, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x66, 0x5f, 0x68,
	0x61, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x75,
	0x72, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x5f, 0x65,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6f, 0x66, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x73, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x75, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x72, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x69, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x64, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x75, 0x6d, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6a, 0x75, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x77, 0x61, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x77, 0x61, 0x70, 0x22,
	0xa8, 0x01, 0x0a, 0x0e, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x64, 0x72,
	0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6e,
	0x73, 0x69, 0x6d, 0x2f, 0x67, 0x63, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_sim_proto_rawDescOnce sync.Once
	file_protos_model_sim_proto_rawDescData = file_protos_model_sim_proto_rawDesc
)

func file_protos_model_sim_proto_rawDescGZIP() []byte {
	file_protos_model_sim_proto_rawDescOnce.Do(func() {
		file_protos_model_sim_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_sim_proto_rawDescData)
	})
	return file_protos_model_sim_proto_rawDescData
}

var file_protos_model_sim_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_model_sim_proto_goTypes = []interface{}{
	(*Character)(nil),         // 0: model.Character
	(*CharacterTalents)(nil),  // 1: model.CharacterTalents
	(*Weapon)(nil),            // 2: model.Weapon
	(*Enemy)(nil),             // 3: model.Enemy
	(*Coord)(nil),             // 4: model.Coord
	(*SimulatorSettings)(nil), // 5: model.SimulatorSettings
	(*Delays)(nil),            // 6: model.Delays
	(*EnergySettings)(nil),    // 7: model.EnergySettings
	nil,                       // 8: model.Character.SetsEntry
	nil,                       // 9: model.Enemy.ResistEntry
}
var file_protos_model_sim_proto_depIdxs = []int32{
	2, // 0: model.Character.weapon:type_name -> model.Weapon
	1, // 1: model.Character.talents:type_name -> model.CharacterTalents
	8, // 2: model.Character.sets:type_name -> model.Character.SetsEntry
	9, // 3: model.Enemy.resist:type_name -> model.Enemy.ResistEntry
	4, // 4: model.Enemy.pos:type_name -> model.Coord
	6, // 5: model.SimulatorSettings.delays:type_name -> model.Delays
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_model_sim_proto_init() }
func file_protos_model_sim_proto_init() {
	if File_protos_model_sim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_model_sim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_protos_model_sim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterTalents); i {
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
		file_protos_model_sim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weapon); i {
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
		file_protos_model_sim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enemy); i {
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
		file_protos_model_sim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coord); i {
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
		file_protos_model_sim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulatorSettings); i {
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
		file_protos_model_sim_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delays); i {
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
		file_protos_model_sim_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnergySettings); i {
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
			RawDescriptor: file_protos_model_sim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_sim_proto_goTypes,
		DependencyIndexes: file_protos_model_sim_proto_depIdxs,
		MessageInfos:      file_protos_model_sim_proto_msgTypes,
	}.Build()
	File_protos_model_sim_proto = out.File
	file_protos_model_sim_proto_rawDesc = nil
	file_protos_model_sim_proto_goTypes = nil
	file_protos_model_sim_proto_depIdxs = nil
}

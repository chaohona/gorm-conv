// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gorm-db.proto

#include "gorm-db.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
namespace gorm {
class GORM_PB_Table_accountDefaultTypeInternal {
 public:
  ::PROTOBUF_NAMESPACE_ID::internal::ExplicitlyConstructed<GORM_PB_Table_account> _instance;
} _GORM_PB_Table_account_default_instance_;
class GORM_PB_Table_bagDefaultTypeInternal {
 public:
  ::PROTOBUF_NAMESPACE_ID::internal::ExplicitlyConstructed<GORM_PB_Table_bag> _instance;
} _GORM_PB_Table_bag_default_instance_;
}  // namespace gorm
static void InitDefaultsscc_info_GORM_PB_Table_account_gorm_2ddb_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::gorm::_GORM_PB_Table_account_default_instance_;
    new (ptr) ::gorm::GORM_PB_Table_account();
    ::PROTOBUF_NAMESPACE_ID::internal::OnShutdownDestroyMessage(ptr);
  }
  ::gorm::GORM_PB_Table_account::InitAsDefaultInstance();
}

::PROTOBUF_NAMESPACE_ID::internal::SCCInfo<0> scc_info_GORM_PB_Table_account_gorm_2ddb_2eproto =
    {{ATOMIC_VAR_INIT(::PROTOBUF_NAMESPACE_ID::internal::SCCInfoBase::kUninitialized), 0, 0, InitDefaultsscc_info_GORM_PB_Table_account_gorm_2ddb_2eproto}, {}};

static void InitDefaultsscc_info_GORM_PB_Table_bag_gorm_2ddb_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::gorm::_GORM_PB_Table_bag_default_instance_;
    new (ptr) ::gorm::GORM_PB_Table_bag();
    ::PROTOBUF_NAMESPACE_ID::internal::OnShutdownDestroyMessage(ptr);
  }
  ::gorm::GORM_PB_Table_bag::InitAsDefaultInstance();
}

::PROTOBUF_NAMESPACE_ID::internal::SCCInfo<0> scc_info_GORM_PB_Table_bag_gorm_2ddb_2eproto =
    {{ATOMIC_VAR_INIT(::PROTOBUF_NAMESPACE_ID::internal::SCCInfoBase::kUninitialized), 0, 0, InitDefaultsscc_info_GORM_PB_Table_bag_gorm_2ddb_2eproto}, {}};

static ::PROTOBUF_NAMESPACE_ID::Metadata file_level_metadata_gorm_2ddb_2eproto[2];
static constexpr ::PROTOBUF_NAMESPACE_ID::EnumDescriptor const** file_level_enum_descriptors_gorm_2ddb_2eproto = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::ServiceDescriptor const** file_level_service_descriptors_gorm_2ddb_2eproto = nullptr;

const ::PROTOBUF_NAMESPACE_ID::uint32 TableStruct_gorm_2ddb_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_account, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_account, version_),
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_account, id_),
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_account, account_),
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_account, allbinary_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_bag, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_bag, version_),
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_bag, id_),
  PROTOBUF_FIELD_OFFSET(::gorm::GORM_PB_Table_bag, allbinary_),
};
static const ::PROTOBUF_NAMESPACE_ID::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::gorm::GORM_PB_Table_account)},
  { 9, -1, sizeof(::gorm::GORM_PB_Table_bag)},
};

static ::PROTOBUF_NAMESPACE_ID::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::PROTOBUF_NAMESPACE_ID::Message*>(&::gorm::_GORM_PB_Table_account_default_instance_),
  reinterpret_cast<const ::PROTOBUF_NAMESPACE_ID::Message*>(&::gorm::_GORM_PB_Table_bag_default_instance_),
};

const char descriptor_table_protodef_gorm_2ddb_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\rgorm-db.proto\022\004gorm\"X\n\025GORM_PB_Table_a"
  "ccount\022\017\n\007version\030\001 \001(\006\022\n\n\002id\030\002 \001(\017\022\017\n\007a"
  "ccount\030\003 \001(\t\022\021\n\tallbinary\030\004 \001(\014\"C\n\021GORM_"
  "PB_Table_bag\022\017\n\007version\030\001 \001(\006\022\n\n\002id\030\002 \001("
  "\017\022\021\n\tallbinary\030\003 \001(\014B\013Z\tgorm/gormb\006proto"
  "3"
  ;
static const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable*const descriptor_table_gorm_2ddb_2eproto_deps[1] = {
};
static ::PROTOBUF_NAMESPACE_ID::internal::SCCInfoBase*const descriptor_table_gorm_2ddb_2eproto_sccs[2] = {
  &scc_info_GORM_PB_Table_account_gorm_2ddb_2eproto.base,
  &scc_info_GORM_PB_Table_bag_gorm_2ddb_2eproto.base,
};
static ::PROTOBUF_NAMESPACE_ID::internal::once_flag descriptor_table_gorm_2ddb_2eproto_once;
const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_gorm_2ddb_2eproto = {
  false, false, descriptor_table_protodef_gorm_2ddb_2eproto, "gorm-db.proto", 201,
  &descriptor_table_gorm_2ddb_2eproto_once, descriptor_table_gorm_2ddb_2eproto_sccs, descriptor_table_gorm_2ddb_2eproto_deps, 2, 0,
  schemas, file_default_instances, TableStruct_gorm_2ddb_2eproto::offsets,
  file_level_metadata_gorm_2ddb_2eproto, 2, file_level_enum_descriptors_gorm_2ddb_2eproto, file_level_service_descriptors_gorm_2ddb_2eproto,
};

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_gorm_2ddb_2eproto = (static_cast<void>(::PROTOBUF_NAMESPACE_ID::internal::AddDescriptors(&descriptor_table_gorm_2ddb_2eproto)), true);
namespace gorm {

// ===================================================================

void GORM_PB_Table_account::InitAsDefaultInstance() {
}
class GORM_PB_Table_account::_Internal {
 public:
};

GORM_PB_Table_account::GORM_PB_Table_account(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor();
  RegisterArenaDtor(arena);
  // @@protoc_insertion_point(arena_constructor:gorm.GORM_PB_Table_account)
}
GORM_PB_Table_account::GORM_PB_Table_account(const GORM_PB_Table_account& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  account_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  if (!from._internal_account().empty()) {
    account_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), from._internal_account(),
      GetArena());
  }
  allbinary_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  if (!from._internal_allbinary().empty()) {
    allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), from._internal_allbinary(),
      GetArena());
  }
  ::memcpy(&version_, &from.version_,
    static_cast<size_t>(reinterpret_cast<char*>(&id_) -
    reinterpret_cast<char*>(&version_)) + sizeof(id_));
  // @@protoc_insertion_point(copy_constructor:gorm.GORM_PB_Table_account)
}

void GORM_PB_Table_account::SharedCtor() {
  ::PROTOBUF_NAMESPACE_ID::internal::InitSCC(&scc_info_GORM_PB_Table_account_gorm_2ddb_2eproto.base);
  account_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  allbinary_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  ::memset(&version_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&id_) -
      reinterpret_cast<char*>(&version_)) + sizeof(id_));
}

GORM_PB_Table_account::~GORM_PB_Table_account() {
  // @@protoc_insertion_point(destructor:gorm.GORM_PB_Table_account)
  SharedDtor();
  _internal_metadata_.Delete<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

void GORM_PB_Table_account::SharedDtor() {
  GOOGLE_DCHECK(GetArena() == nullptr);
  account_.DestroyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  allbinary_.DestroyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}

void GORM_PB_Table_account::ArenaDtor(void* object) {
  GORM_PB_Table_account* _this = reinterpret_cast< GORM_PB_Table_account* >(object);
  (void)_this;
}
void GORM_PB_Table_account::RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena*) {
}
void GORM_PB_Table_account::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const GORM_PB_Table_account& GORM_PB_Table_account::default_instance() {
  ::PROTOBUF_NAMESPACE_ID::internal::InitSCC(&::scc_info_GORM_PB_Table_account_gorm_2ddb_2eproto.base);
  return *internal_default_instance();
}


void GORM_PB_Table_account::Clear() {
// @@protoc_insertion_point(message_clear_start:gorm.GORM_PB_Table_account)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  account_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  allbinary_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  ::memset(&version_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&id_) -
      reinterpret_cast<char*>(&version_)) + sizeof(id_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* GORM_PB_Table_account::_InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  ::PROTOBUF_NAMESPACE_ID::Arena* arena = GetArena(); (void)arena;
  while (!ctx->Done(&ptr)) {
    ::PROTOBUF_NAMESPACE_ID::uint32 tag;
    ptr = ::PROTOBUF_NAMESPACE_ID::internal::ReadTag(ptr, &tag);
    CHK_(ptr);
    switch (tag >> 3) {
      // fixed64 version = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 9)) {
          version_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<::PROTOBUF_NAMESPACE_ID::uint64>(ptr);
          ptr += sizeof(::PROTOBUF_NAMESPACE_ID::uint64);
        } else goto handle_unusual;
        continue;
      // sfixed32 id = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 21)) {
          id_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<::PROTOBUF_NAMESPACE_ID::int32>(ptr);
          ptr += sizeof(::PROTOBUF_NAMESPACE_ID::int32);
        } else goto handle_unusual;
        continue;
      // string account = 3;
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 26)) {
          auto str = _internal_mutable_account();
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(::PROTOBUF_NAMESPACE_ID::internal::VerifyUTF8(str, "gorm.GORM_PB_Table_account.account"));
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      // bytes allbinary = 4;
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 34)) {
          auto str = _internal_mutable_allbinary();
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      default: {
      handle_unusual:
        if ((tag & 7) == 4 || tag == 0) {
          ctx->SetLastTag(tag);
          goto success;
        }
        ptr = UnknownFieldParse(tag,
            _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
            ptr, ctx);
        CHK_(ptr != nullptr);
        continue;
      }
    }  // switch
  }  // while
success:
  return ptr;
failure:
  ptr = nullptr;
  goto success;
#undef CHK_
}

::PROTOBUF_NAMESPACE_ID::uint8* GORM_PB_Table_account::_InternalSerialize(
    ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:gorm.GORM_PB_Table_account)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // fixed64 version = 1;
  if (this->version() != 0) {
    target = stream->EnsureSpace(target);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::WriteFixed64ToArray(1, this->_internal_version(), target);
  }

  // sfixed32 id = 2;
  if (this->id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::WriteSFixed32ToArray(2, this->_internal_id(), target);
  }

  // string account = 3;
  if (this->account().size() > 0) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_account().data(), static_cast<int>(this->_internal_account().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "gorm.GORM_PB_Table_account.account");
    target = stream->WriteStringMaybeAliased(
        3, this->_internal_account(), target);
  }

  // bytes allbinary = 4;
  if (this->allbinary().size() > 0) {
    target = stream->WriteBytesMaybeAliased(
        4, this->_internal_allbinary(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:gorm.GORM_PB_Table_account)
  return target;
}

size_t GORM_PB_Table_account::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:gorm.GORM_PB_Table_account)
  size_t total_size = 0;

  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string account = 3;
  if (this->account().size() > 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_account());
  }

  // bytes allbinary = 4;
  if (this->allbinary().size() > 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::BytesSize(
        this->_internal_allbinary());
  }

  // fixed64 version = 1;
  if (this->version() != 0) {
    total_size += 1 + 8;
  }

  // sfixed32 id = 2;
  if (this->id() != 0) {
    total_size += 1 + 4;
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    return ::PROTOBUF_NAMESPACE_ID::internal::ComputeUnknownFieldsSize(
        _internal_metadata_, total_size, &_cached_size_);
  }
  int cached_size = ::PROTOBUF_NAMESPACE_ID::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void GORM_PB_Table_account::MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:gorm.GORM_PB_Table_account)
  GOOGLE_DCHECK_NE(&from, this);
  const GORM_PB_Table_account* source =
      ::PROTOBUF_NAMESPACE_ID::DynamicCastToGenerated<GORM_PB_Table_account>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:gorm.GORM_PB_Table_account)
    ::PROTOBUF_NAMESPACE_ID::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:gorm.GORM_PB_Table_account)
    MergeFrom(*source);
  }
}

void GORM_PB_Table_account::MergeFrom(const GORM_PB_Table_account& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:gorm.GORM_PB_Table_account)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.account().size() > 0) {
    _internal_set_account(from._internal_account());
  }
  if (from.allbinary().size() > 0) {
    _internal_set_allbinary(from._internal_allbinary());
  }
  if (from.version() != 0) {
    _internal_set_version(from._internal_version());
  }
  if (from.id() != 0) {
    _internal_set_id(from._internal_id());
  }
}

void GORM_PB_Table_account::CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:gorm.GORM_PB_Table_account)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void GORM_PB_Table_account::CopyFrom(const GORM_PB_Table_account& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:gorm.GORM_PB_Table_account)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool GORM_PB_Table_account::IsInitialized() const {
  return true;
}

void GORM_PB_Table_account::InternalSwap(GORM_PB_Table_account* other) {
  using std::swap;
  _internal_metadata_.Swap<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(&other->_internal_metadata_);
  account_.Swap(&other->account_, &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  allbinary_.Swap(&other->allbinary_, &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(GORM_PB_Table_account, id_)
      + sizeof(GORM_PB_Table_account::id_)
      - PROTOBUF_FIELD_OFFSET(GORM_PB_Table_account, version_)>(
          reinterpret_cast<char*>(&version_),
          reinterpret_cast<char*>(&other->version_));
}

::PROTOBUF_NAMESPACE_ID::Metadata GORM_PB_Table_account::GetMetadata() const {
  return GetMetadataStatic();
}


// ===================================================================

void GORM_PB_Table_bag::InitAsDefaultInstance() {
}
class GORM_PB_Table_bag::_Internal {
 public:
};

GORM_PB_Table_bag::GORM_PB_Table_bag(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor();
  RegisterArenaDtor(arena);
  // @@protoc_insertion_point(arena_constructor:gorm.GORM_PB_Table_bag)
}
GORM_PB_Table_bag::GORM_PB_Table_bag(const GORM_PB_Table_bag& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  allbinary_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  if (!from._internal_allbinary().empty()) {
    allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), from._internal_allbinary(),
      GetArena());
  }
  ::memcpy(&version_, &from.version_,
    static_cast<size_t>(reinterpret_cast<char*>(&id_) -
    reinterpret_cast<char*>(&version_)) + sizeof(id_));
  // @@protoc_insertion_point(copy_constructor:gorm.GORM_PB_Table_bag)
}

void GORM_PB_Table_bag::SharedCtor() {
  ::PROTOBUF_NAMESPACE_ID::internal::InitSCC(&scc_info_GORM_PB_Table_bag_gorm_2ddb_2eproto.base);
  allbinary_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  ::memset(&version_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&id_) -
      reinterpret_cast<char*>(&version_)) + sizeof(id_));
}

GORM_PB_Table_bag::~GORM_PB_Table_bag() {
  // @@protoc_insertion_point(destructor:gorm.GORM_PB_Table_bag)
  SharedDtor();
  _internal_metadata_.Delete<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

void GORM_PB_Table_bag::SharedDtor() {
  GOOGLE_DCHECK(GetArena() == nullptr);
  allbinary_.DestroyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}

void GORM_PB_Table_bag::ArenaDtor(void* object) {
  GORM_PB_Table_bag* _this = reinterpret_cast< GORM_PB_Table_bag* >(object);
  (void)_this;
}
void GORM_PB_Table_bag::RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena*) {
}
void GORM_PB_Table_bag::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const GORM_PB_Table_bag& GORM_PB_Table_bag::default_instance() {
  ::PROTOBUF_NAMESPACE_ID::internal::InitSCC(&::scc_info_GORM_PB_Table_bag_gorm_2ddb_2eproto.base);
  return *internal_default_instance();
}


void GORM_PB_Table_bag::Clear() {
// @@protoc_insertion_point(message_clear_start:gorm.GORM_PB_Table_bag)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  allbinary_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  ::memset(&version_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&id_) -
      reinterpret_cast<char*>(&version_)) + sizeof(id_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* GORM_PB_Table_bag::_InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  ::PROTOBUF_NAMESPACE_ID::Arena* arena = GetArena(); (void)arena;
  while (!ctx->Done(&ptr)) {
    ::PROTOBUF_NAMESPACE_ID::uint32 tag;
    ptr = ::PROTOBUF_NAMESPACE_ID::internal::ReadTag(ptr, &tag);
    CHK_(ptr);
    switch (tag >> 3) {
      // fixed64 version = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 9)) {
          version_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<::PROTOBUF_NAMESPACE_ID::uint64>(ptr);
          ptr += sizeof(::PROTOBUF_NAMESPACE_ID::uint64);
        } else goto handle_unusual;
        continue;
      // sfixed32 id = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 21)) {
          id_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<::PROTOBUF_NAMESPACE_ID::int32>(ptr);
          ptr += sizeof(::PROTOBUF_NAMESPACE_ID::int32);
        } else goto handle_unusual;
        continue;
      // bytes allbinary = 3;
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 26)) {
          auto str = _internal_mutable_allbinary();
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      default: {
      handle_unusual:
        if ((tag & 7) == 4 || tag == 0) {
          ctx->SetLastTag(tag);
          goto success;
        }
        ptr = UnknownFieldParse(tag,
            _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
            ptr, ctx);
        CHK_(ptr != nullptr);
        continue;
      }
    }  // switch
  }  // while
success:
  return ptr;
failure:
  ptr = nullptr;
  goto success;
#undef CHK_
}

::PROTOBUF_NAMESPACE_ID::uint8* GORM_PB_Table_bag::_InternalSerialize(
    ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:gorm.GORM_PB_Table_bag)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // fixed64 version = 1;
  if (this->version() != 0) {
    target = stream->EnsureSpace(target);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::WriteFixed64ToArray(1, this->_internal_version(), target);
  }

  // sfixed32 id = 2;
  if (this->id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::WriteSFixed32ToArray(2, this->_internal_id(), target);
  }

  // bytes allbinary = 3;
  if (this->allbinary().size() > 0) {
    target = stream->WriteBytesMaybeAliased(
        3, this->_internal_allbinary(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:gorm.GORM_PB_Table_bag)
  return target;
}

size_t GORM_PB_Table_bag::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:gorm.GORM_PB_Table_bag)
  size_t total_size = 0;

  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // bytes allbinary = 3;
  if (this->allbinary().size() > 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::BytesSize(
        this->_internal_allbinary());
  }

  // fixed64 version = 1;
  if (this->version() != 0) {
    total_size += 1 + 8;
  }

  // sfixed32 id = 2;
  if (this->id() != 0) {
    total_size += 1 + 4;
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    return ::PROTOBUF_NAMESPACE_ID::internal::ComputeUnknownFieldsSize(
        _internal_metadata_, total_size, &_cached_size_);
  }
  int cached_size = ::PROTOBUF_NAMESPACE_ID::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void GORM_PB_Table_bag::MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:gorm.GORM_PB_Table_bag)
  GOOGLE_DCHECK_NE(&from, this);
  const GORM_PB_Table_bag* source =
      ::PROTOBUF_NAMESPACE_ID::DynamicCastToGenerated<GORM_PB_Table_bag>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:gorm.GORM_PB_Table_bag)
    ::PROTOBUF_NAMESPACE_ID::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:gorm.GORM_PB_Table_bag)
    MergeFrom(*source);
  }
}

void GORM_PB_Table_bag::MergeFrom(const GORM_PB_Table_bag& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:gorm.GORM_PB_Table_bag)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.allbinary().size() > 0) {
    _internal_set_allbinary(from._internal_allbinary());
  }
  if (from.version() != 0) {
    _internal_set_version(from._internal_version());
  }
  if (from.id() != 0) {
    _internal_set_id(from._internal_id());
  }
}

void GORM_PB_Table_bag::CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:gorm.GORM_PB_Table_bag)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void GORM_PB_Table_bag::CopyFrom(const GORM_PB_Table_bag& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:gorm.GORM_PB_Table_bag)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool GORM_PB_Table_bag::IsInitialized() const {
  return true;
}

void GORM_PB_Table_bag::InternalSwap(GORM_PB_Table_bag* other) {
  using std::swap;
  _internal_metadata_.Swap<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(&other->_internal_metadata_);
  allbinary_.Swap(&other->allbinary_, &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(GORM_PB_Table_bag, id_)
      + sizeof(GORM_PB_Table_bag::id_)
      - PROTOBUF_FIELD_OFFSET(GORM_PB_Table_bag, version_)>(
          reinterpret_cast<char*>(&version_),
          reinterpret_cast<char*>(&other->version_));
}

::PROTOBUF_NAMESPACE_ID::Metadata GORM_PB_Table_bag::GetMetadata() const {
  return GetMetadataStatic();
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace gorm
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::gorm::GORM_PB_Table_account* Arena::CreateMaybeMessage< ::gorm::GORM_PB_Table_account >(Arena* arena) {
  return Arena::CreateMessageInternal< ::gorm::GORM_PB_Table_account >(arena);
}
template<> PROTOBUF_NOINLINE ::gorm::GORM_PB_Table_bag* Arena::CreateMaybeMessage< ::gorm::GORM_PB_Table_bag >(Arena* arena) {
  return Arena::CreateMessageInternal< ::gorm::GORM_PB_Table_bag >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>

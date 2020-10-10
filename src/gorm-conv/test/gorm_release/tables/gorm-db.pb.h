// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gorm-db.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_gorm_2ddb_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_gorm_2ddb_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3012000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3012000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_gorm_2ddb_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_gorm_2ddb_2eproto {
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::FieldMetadata field_metadata[];
  static const ::PROTOBUF_NAMESPACE_ID::internal::SerializationTable serialization_table[];
  static const ::PROTOBUF_NAMESPACE_ID::uint32 offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_gorm_2ddb_2eproto;
namespace gorm {
class GORM_PB_Table_account;
class GORM_PB_Table_accountDefaultTypeInternal;
extern GORM_PB_Table_accountDefaultTypeInternal _GORM_PB_Table_account_default_instance_;
class GORM_PB_Table_bag;
class GORM_PB_Table_bagDefaultTypeInternal;
extern GORM_PB_Table_bagDefaultTypeInternal _GORM_PB_Table_bag_default_instance_;
}  // namespace gorm
PROTOBUF_NAMESPACE_OPEN
template<> ::gorm::GORM_PB_Table_account* Arena::CreateMaybeMessage<::gorm::GORM_PB_Table_account>(Arena*);
template<> ::gorm::GORM_PB_Table_bag* Arena::CreateMaybeMessage<::gorm::GORM_PB_Table_bag>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace gorm {

// ===================================================================

class GORM_PB_Table_account PROTOBUF_FINAL :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:gorm.GORM_PB_Table_account) */ {
 public:
  inline GORM_PB_Table_account() : GORM_PB_Table_account(nullptr) {};
  virtual ~GORM_PB_Table_account();

  GORM_PB_Table_account(const GORM_PB_Table_account& from);
  GORM_PB_Table_account(GORM_PB_Table_account&& from) noexcept
    : GORM_PB_Table_account() {
    *this = ::std::move(from);
  }

  inline GORM_PB_Table_account& operator=(const GORM_PB_Table_account& from) {
    CopyFrom(from);
    return *this;
  }
  inline GORM_PB_Table_account& operator=(GORM_PB_Table_account&& from) noexcept {
    if (GetArena() == from.GetArena()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const GORM_PB_Table_account& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const GORM_PB_Table_account* internal_default_instance() {
    return reinterpret_cast<const GORM_PB_Table_account*>(
               &_GORM_PB_Table_account_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(GORM_PB_Table_account& a, GORM_PB_Table_account& b) {
    a.Swap(&b);
  }
  inline void Swap(GORM_PB_Table_account* other) {
    if (other == this) return;
    if (GetArena() == other->GetArena()) {
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(GORM_PB_Table_account* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline GORM_PB_Table_account* New() const final {
    return CreateMaybeMessage<GORM_PB_Table_account>(nullptr);
  }

  GORM_PB_Table_account* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<GORM_PB_Table_account>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const GORM_PB_Table_account& from);
  void MergeFrom(const GORM_PB_Table_account& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  ::PROTOBUF_NAMESPACE_ID::uint8* _InternalSerialize(
      ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(GORM_PB_Table_account* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "gorm.GORM_PB_Table_account";
  }
  protected:
  explicit GORM_PB_Table_account(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_gorm_2ddb_2eproto);
    return ::descriptor_table_gorm_2ddb_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kAccountFieldNumber = 3,
    kAllbinaryFieldNumber = 4,
    kVersionFieldNumber = 1,
    kIdFieldNumber = 2,
  };
  // string account = 3;
  void clear_account();
  const std::string& account() const;
  void set_account(const std::string& value);
  void set_account(std::string&& value);
  void set_account(const char* value);
  void set_account(const char* value, size_t size);
  std::string* mutable_account();
  std::string* release_account();
  void set_allocated_account(std::string* account);
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  std::string* unsafe_arena_release_account();
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  void unsafe_arena_set_allocated_account(
      std::string* account);
  private:
  const std::string& _internal_account() const;
  void _internal_set_account(const std::string& value);
  std::string* _internal_mutable_account();
  public:

  // bytes allbinary = 4;
  void clear_allbinary();
  const std::string& allbinary() const;
  void set_allbinary(const std::string& value);
  void set_allbinary(std::string&& value);
  void set_allbinary(const char* value);
  void set_allbinary(const void* value, size_t size);
  std::string* mutable_allbinary();
  std::string* release_allbinary();
  void set_allocated_allbinary(std::string* allbinary);
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  std::string* unsafe_arena_release_allbinary();
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  void unsafe_arena_set_allocated_allbinary(
      std::string* allbinary);
  private:
  const std::string& _internal_allbinary() const;
  void _internal_set_allbinary(const std::string& value);
  std::string* _internal_mutable_allbinary();
  public:

  // fixed64 version = 1;
  void clear_version();
  ::PROTOBUF_NAMESPACE_ID::uint64 version() const;
  void set_version(::PROTOBUF_NAMESPACE_ID::uint64 value);
  private:
  ::PROTOBUF_NAMESPACE_ID::uint64 _internal_version() const;
  void _internal_set_version(::PROTOBUF_NAMESPACE_ID::uint64 value);
  public:

  // sfixed32 id = 2;
  void clear_id();
  ::PROTOBUF_NAMESPACE_ID::int32 id() const;
  void set_id(::PROTOBUF_NAMESPACE_ID::int32 value);
  private:
  ::PROTOBUF_NAMESPACE_ID::int32 _internal_id() const;
  void _internal_set_id(::PROTOBUF_NAMESPACE_ID::int32 value);
  public:

  // @@protoc_insertion_point(class_scope:gorm.GORM_PB_Table_account)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr account_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr allbinary_;
  ::PROTOBUF_NAMESPACE_ID::uint64 version_;
  ::PROTOBUF_NAMESPACE_ID::int32 id_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_gorm_2ddb_2eproto;
};
// -------------------------------------------------------------------

class GORM_PB_Table_bag PROTOBUF_FINAL :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:gorm.GORM_PB_Table_bag) */ {
 public:
  inline GORM_PB_Table_bag() : GORM_PB_Table_bag(nullptr) {};
  virtual ~GORM_PB_Table_bag();

  GORM_PB_Table_bag(const GORM_PB_Table_bag& from);
  GORM_PB_Table_bag(GORM_PB_Table_bag&& from) noexcept
    : GORM_PB_Table_bag() {
    *this = ::std::move(from);
  }

  inline GORM_PB_Table_bag& operator=(const GORM_PB_Table_bag& from) {
    CopyFrom(from);
    return *this;
  }
  inline GORM_PB_Table_bag& operator=(GORM_PB_Table_bag&& from) noexcept {
    if (GetArena() == from.GetArena()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const GORM_PB_Table_bag& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const GORM_PB_Table_bag* internal_default_instance() {
    return reinterpret_cast<const GORM_PB_Table_bag*>(
               &_GORM_PB_Table_bag_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(GORM_PB_Table_bag& a, GORM_PB_Table_bag& b) {
    a.Swap(&b);
  }
  inline void Swap(GORM_PB_Table_bag* other) {
    if (other == this) return;
    if (GetArena() == other->GetArena()) {
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(GORM_PB_Table_bag* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline GORM_PB_Table_bag* New() const final {
    return CreateMaybeMessage<GORM_PB_Table_bag>(nullptr);
  }

  GORM_PB_Table_bag* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<GORM_PB_Table_bag>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const GORM_PB_Table_bag& from);
  void MergeFrom(const GORM_PB_Table_bag& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  ::PROTOBUF_NAMESPACE_ID::uint8* _InternalSerialize(
      ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(GORM_PB_Table_bag* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "gorm.GORM_PB_Table_bag";
  }
  protected:
  explicit GORM_PB_Table_bag(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_gorm_2ddb_2eproto);
    return ::descriptor_table_gorm_2ddb_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kAllbinaryFieldNumber = 3,
    kVersionFieldNumber = 1,
    kIdFieldNumber = 2,
  };
  // bytes allbinary = 3;
  void clear_allbinary();
  const std::string& allbinary() const;
  void set_allbinary(const std::string& value);
  void set_allbinary(std::string&& value);
  void set_allbinary(const char* value);
  void set_allbinary(const void* value, size_t size);
  std::string* mutable_allbinary();
  std::string* release_allbinary();
  void set_allocated_allbinary(std::string* allbinary);
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  std::string* unsafe_arena_release_allbinary();
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  void unsafe_arena_set_allocated_allbinary(
      std::string* allbinary);
  private:
  const std::string& _internal_allbinary() const;
  void _internal_set_allbinary(const std::string& value);
  std::string* _internal_mutable_allbinary();
  public:

  // fixed64 version = 1;
  void clear_version();
  ::PROTOBUF_NAMESPACE_ID::uint64 version() const;
  void set_version(::PROTOBUF_NAMESPACE_ID::uint64 value);
  private:
  ::PROTOBUF_NAMESPACE_ID::uint64 _internal_version() const;
  void _internal_set_version(::PROTOBUF_NAMESPACE_ID::uint64 value);
  public:

  // sfixed32 id = 2;
  void clear_id();
  ::PROTOBUF_NAMESPACE_ID::int32 id() const;
  void set_id(::PROTOBUF_NAMESPACE_ID::int32 value);
  private:
  ::PROTOBUF_NAMESPACE_ID::int32 _internal_id() const;
  void _internal_set_id(::PROTOBUF_NAMESPACE_ID::int32 value);
  public:

  // @@protoc_insertion_point(class_scope:gorm.GORM_PB_Table_bag)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr allbinary_;
  ::PROTOBUF_NAMESPACE_ID::uint64 version_;
  ::PROTOBUF_NAMESPACE_ID::int32 id_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_gorm_2ddb_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// GORM_PB_Table_account

// fixed64 version = 1;
inline void GORM_PB_Table_account::clear_version() {
  version_ = PROTOBUF_ULONGLONG(0);
}
inline ::PROTOBUF_NAMESPACE_ID::uint64 GORM_PB_Table_account::_internal_version() const {
  return version_;
}
inline ::PROTOBUF_NAMESPACE_ID::uint64 GORM_PB_Table_account::version() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_account.version)
  return _internal_version();
}
inline void GORM_PB_Table_account::_internal_set_version(::PROTOBUF_NAMESPACE_ID::uint64 value) {
  
  version_ = value;
}
inline void GORM_PB_Table_account::set_version(::PROTOBUF_NAMESPACE_ID::uint64 value) {
  _internal_set_version(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_account.version)
}

// sfixed32 id = 2;
inline void GORM_PB_Table_account::clear_id() {
  id_ = 0;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 GORM_PB_Table_account::_internal_id() const {
  return id_;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 GORM_PB_Table_account::id() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_account.id)
  return _internal_id();
}
inline void GORM_PB_Table_account::_internal_set_id(::PROTOBUF_NAMESPACE_ID::int32 value) {
  
  id_ = value;
}
inline void GORM_PB_Table_account::set_id(::PROTOBUF_NAMESPACE_ID::int32 value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_account.id)
}

// string account = 3;
inline void GORM_PB_Table_account::clear_account() {
  account_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline const std::string& GORM_PB_Table_account::account() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_account.account)
  return _internal_account();
}
inline void GORM_PB_Table_account::set_account(const std::string& value) {
  _internal_set_account(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_account.account)
}
inline std::string* GORM_PB_Table_account::mutable_account() {
  // @@protoc_insertion_point(field_mutable:gorm.GORM_PB_Table_account.account)
  return _internal_mutable_account();
}
inline const std::string& GORM_PB_Table_account::_internal_account() const {
  return account_.Get();
}
inline void GORM_PB_Table_account::_internal_set_account(const std::string& value) {
  
  account_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value, GetArena());
}
inline void GORM_PB_Table_account::set_account(std::string&& value) {
  
  account_.Set(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:gorm.GORM_PB_Table_account.account)
}
inline void GORM_PB_Table_account::set_account(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  account_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value),
              GetArena());
  // @@protoc_insertion_point(field_set_char:gorm.GORM_PB_Table_account.account)
}
inline void GORM_PB_Table_account::set_account(const char* value,
    size_t size) {
  
  account_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size), GetArena());
  // @@protoc_insertion_point(field_set_pointer:gorm.GORM_PB_Table_account.account)
}
inline std::string* GORM_PB_Table_account::_internal_mutable_account() {
  
  return account_.Mutable(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline std::string* GORM_PB_Table_account::release_account() {
  // @@protoc_insertion_point(field_release:gorm.GORM_PB_Table_account.account)
  return account_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline void GORM_PB_Table_account::set_allocated_account(std::string* account) {
  if (account != nullptr) {
    
  } else {
    
  }
  account_.SetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), account,
      GetArena());
  // @@protoc_insertion_point(field_set_allocated:gorm.GORM_PB_Table_account.account)
}
inline std::string* GORM_PB_Table_account::unsafe_arena_release_account() {
  // @@protoc_insertion_point(field_unsafe_arena_release:gorm.GORM_PB_Table_account.account)
  GOOGLE_DCHECK(GetArena() != nullptr);
  
  return account_.UnsafeArenaRelease(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      GetArena());
}
inline void GORM_PB_Table_account::unsafe_arena_set_allocated_account(
    std::string* account) {
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (account != nullptr) {
    
  } else {
    
  }
  account_.UnsafeArenaSetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      account, GetArena());
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:gorm.GORM_PB_Table_account.account)
}

// bytes allbinary = 4;
inline void GORM_PB_Table_account::clear_allbinary() {
  allbinary_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline const std::string& GORM_PB_Table_account::allbinary() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_account.allbinary)
  return _internal_allbinary();
}
inline void GORM_PB_Table_account::set_allbinary(const std::string& value) {
  _internal_set_allbinary(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_account.allbinary)
}
inline std::string* GORM_PB_Table_account::mutable_allbinary() {
  // @@protoc_insertion_point(field_mutable:gorm.GORM_PB_Table_account.allbinary)
  return _internal_mutable_allbinary();
}
inline const std::string& GORM_PB_Table_account::_internal_allbinary() const {
  return allbinary_.Get();
}
inline void GORM_PB_Table_account::_internal_set_allbinary(const std::string& value) {
  
  allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value, GetArena());
}
inline void GORM_PB_Table_account::set_allbinary(std::string&& value) {
  
  allbinary_.Set(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:gorm.GORM_PB_Table_account.allbinary)
}
inline void GORM_PB_Table_account::set_allbinary(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value),
              GetArena());
  // @@protoc_insertion_point(field_set_char:gorm.GORM_PB_Table_account.allbinary)
}
inline void GORM_PB_Table_account::set_allbinary(const void* value,
    size_t size) {
  
  allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size), GetArena());
  // @@protoc_insertion_point(field_set_pointer:gorm.GORM_PB_Table_account.allbinary)
}
inline std::string* GORM_PB_Table_account::_internal_mutable_allbinary() {
  
  return allbinary_.Mutable(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline std::string* GORM_PB_Table_account::release_allbinary() {
  // @@protoc_insertion_point(field_release:gorm.GORM_PB_Table_account.allbinary)
  return allbinary_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline void GORM_PB_Table_account::set_allocated_allbinary(std::string* allbinary) {
  if (allbinary != nullptr) {
    
  } else {
    
  }
  allbinary_.SetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), allbinary,
      GetArena());
  // @@protoc_insertion_point(field_set_allocated:gorm.GORM_PB_Table_account.allbinary)
}
inline std::string* GORM_PB_Table_account::unsafe_arena_release_allbinary() {
  // @@protoc_insertion_point(field_unsafe_arena_release:gorm.GORM_PB_Table_account.allbinary)
  GOOGLE_DCHECK(GetArena() != nullptr);
  
  return allbinary_.UnsafeArenaRelease(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      GetArena());
}
inline void GORM_PB_Table_account::unsafe_arena_set_allocated_allbinary(
    std::string* allbinary) {
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (allbinary != nullptr) {
    
  } else {
    
  }
  allbinary_.UnsafeArenaSetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      allbinary, GetArena());
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:gorm.GORM_PB_Table_account.allbinary)
}

// -------------------------------------------------------------------

// GORM_PB_Table_bag

// fixed64 version = 1;
inline void GORM_PB_Table_bag::clear_version() {
  version_ = PROTOBUF_ULONGLONG(0);
}
inline ::PROTOBUF_NAMESPACE_ID::uint64 GORM_PB_Table_bag::_internal_version() const {
  return version_;
}
inline ::PROTOBUF_NAMESPACE_ID::uint64 GORM_PB_Table_bag::version() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_bag.version)
  return _internal_version();
}
inline void GORM_PB_Table_bag::_internal_set_version(::PROTOBUF_NAMESPACE_ID::uint64 value) {
  
  version_ = value;
}
inline void GORM_PB_Table_bag::set_version(::PROTOBUF_NAMESPACE_ID::uint64 value) {
  _internal_set_version(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_bag.version)
}

// sfixed32 id = 2;
inline void GORM_PB_Table_bag::clear_id() {
  id_ = 0;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 GORM_PB_Table_bag::_internal_id() const {
  return id_;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 GORM_PB_Table_bag::id() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_bag.id)
  return _internal_id();
}
inline void GORM_PB_Table_bag::_internal_set_id(::PROTOBUF_NAMESPACE_ID::int32 value) {
  
  id_ = value;
}
inline void GORM_PB_Table_bag::set_id(::PROTOBUF_NAMESPACE_ID::int32 value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_bag.id)
}

// bytes allbinary = 3;
inline void GORM_PB_Table_bag::clear_allbinary() {
  allbinary_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline const std::string& GORM_PB_Table_bag::allbinary() const {
  // @@protoc_insertion_point(field_get:gorm.GORM_PB_Table_bag.allbinary)
  return _internal_allbinary();
}
inline void GORM_PB_Table_bag::set_allbinary(const std::string& value) {
  _internal_set_allbinary(value);
  // @@protoc_insertion_point(field_set:gorm.GORM_PB_Table_bag.allbinary)
}
inline std::string* GORM_PB_Table_bag::mutable_allbinary() {
  // @@protoc_insertion_point(field_mutable:gorm.GORM_PB_Table_bag.allbinary)
  return _internal_mutable_allbinary();
}
inline const std::string& GORM_PB_Table_bag::_internal_allbinary() const {
  return allbinary_.Get();
}
inline void GORM_PB_Table_bag::_internal_set_allbinary(const std::string& value) {
  
  allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value, GetArena());
}
inline void GORM_PB_Table_bag::set_allbinary(std::string&& value) {
  
  allbinary_.Set(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:gorm.GORM_PB_Table_bag.allbinary)
}
inline void GORM_PB_Table_bag::set_allbinary(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value),
              GetArena());
  // @@protoc_insertion_point(field_set_char:gorm.GORM_PB_Table_bag.allbinary)
}
inline void GORM_PB_Table_bag::set_allbinary(const void* value,
    size_t size) {
  
  allbinary_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size), GetArena());
  // @@protoc_insertion_point(field_set_pointer:gorm.GORM_PB_Table_bag.allbinary)
}
inline std::string* GORM_PB_Table_bag::_internal_mutable_allbinary() {
  
  return allbinary_.Mutable(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline std::string* GORM_PB_Table_bag::release_allbinary() {
  // @@protoc_insertion_point(field_release:gorm.GORM_PB_Table_bag.allbinary)
  return allbinary_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline void GORM_PB_Table_bag::set_allocated_allbinary(std::string* allbinary) {
  if (allbinary != nullptr) {
    
  } else {
    
  }
  allbinary_.SetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), allbinary,
      GetArena());
  // @@protoc_insertion_point(field_set_allocated:gorm.GORM_PB_Table_bag.allbinary)
}
inline std::string* GORM_PB_Table_bag::unsafe_arena_release_allbinary() {
  // @@protoc_insertion_point(field_unsafe_arena_release:gorm.GORM_PB_Table_bag.allbinary)
  GOOGLE_DCHECK(GetArena() != nullptr);
  
  return allbinary_.UnsafeArenaRelease(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      GetArena());
}
inline void GORM_PB_Table_bag::unsafe_arena_set_allocated_allbinary(
    std::string* allbinary) {
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (allbinary != nullptr) {
    
  } else {
    
  }
  allbinary_.UnsafeArenaSetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      allbinary, GetArena());
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:gorm.GORM_PB_Table_bag.allbinary)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace gorm

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_gorm_2ddb_2eproto

/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: mesg.proto */

#ifndef PROTOBUF_C_mesg_2eproto__INCLUDED
#define PROTOBUF_C_mesg_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1000002 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _MesgOnline MesgOnline;
typedef struct _MesgOnlineAck MesgOnlineAck;
typedef struct _MesgKick MesgKick;
typedef struct _BrowserPluginInfo BrowserPluginInfo;
typedef struct _BrowserUseragentInfo BrowserUseragentInfo;
typedef struct _BrowserScreenInfo BrowserScreenInfo;
typedef struct _MesgBrowserEnv MesgBrowserEnv;
typedef struct _MesgBrowserEnvAck MesgBrowserEnvAck;
typedef struct _MesgEventStatistic MesgEventStatistic;
typedef struct _MesgEventStatisticAck MesgEventStatisticAck;
typedef struct _MesgLsndInfo MesgLsndInfo;
typedef struct _MesgFrwdInfo MesgFrwdInfo;


/* --- enums --- */

typedef enum _CtrlType {
  CTRL_TYPE__CTL_IBX_USR = 0,
  CTRL_TYPE__CTL_IBX_PWD = 1,
  CTRL_TYPE__CTL_IBX_IMG = 2,
  CTRL_TYPE__CTL_BTN_IMG = 3,
  CTRL_TYPE__CTL_IBX_TEL = 4,
  CTRL_TYPE__CTL_IBX_SMS = 5,
  CTRL_TYPE__CTL_BTN_SMS = 6,
  CTRL_TYPE__CTL_BTN_LGN = 7
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(CTRL_TYPE)
} CtrlType;
typedef enum _EventType {
  EVENT_TYPE__EV_CHANGE = 0,
  EVENT_TYPE__EV_CLICK = 1,
  EVENT_TYPE__EV_DBLCLICK = 2,
  EVENT_TYPE__EV_FOCUS = 3,
  EVENT_TYPE__EV_KEYDOWN = 4,
  EVENT_TYPE__EV_KEYPRESS = 5,
  EVENT_TYPE__EV_KEYUP = 6,
  EVENT_TYPE__EV_MOUSEDOWN = 7,
  EVENT_TYPE__EV_MOUSEMOVE = 8,
  EVENT_TYPE__EV_MOUSEOUT = 9,
  EVENT_TYPE__EV_MOUSEOVER = 10,
  EVENT_TYPE__EV_MOUSEUP = 11,
  EVENT_TYPE__EV_TOUCHSTART = 12,
  EVENT_TYPE__EV_TOUCHMOVE = 13,
  EVENT_TYPE__EV_TOUCHEND = 14,
  EVENT_TYPE__EV_TOUCHCANCEL = 15
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(EVENT_TYPE)
} EventType;

/* --- messages --- */

struct  _MesgOnline
{
  ProtobufCMessage base;
  uint32_t sid;
  char *token;
  char *app;
  char *version;
};
#define MESG_ONLINE__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_online__descriptor) \
    , 0, NULL, NULL, NULL }


struct  _MesgOnlineAck
{
  ProtobufCMessage base;
  uint32_t sid;
  char *app;
  char *version;
  uint32_t code;
  char *errmsg;
};
#define MESG_ONLINE_ACK__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_online_ack__descriptor) \
    , 0, NULL, NULL, 0, NULL }


struct  _MesgKick
{
  ProtobufCMessage base;
  uint32_t code;
  char *errmsg;
};
#define MESG_KICK__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_kick__descriptor) \
    , 0, NULL }


struct  _BrowserPluginInfo
{
  ProtobufCMessage base;
  char *name;
  char *desc;
};
#define BROWSER_PLUGIN_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&browser_plugin_info__descriptor) \
    , NULL, NULL }


struct  _BrowserUseragentInfo
{
  ProtobufCMessage base;
  char *ua;
};
#define BROWSER_USERAGENT_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&browser_useragent_info__descriptor) \
    , NULL }


struct  _BrowserScreenInfo
{
  ProtobufCMessage base;
  protobuf_c_boolean has_width;
  uint32_t width;
  protobuf_c_boolean has_height;
  uint32_t height;
  protobuf_c_boolean has_avail_width;
  uint32_t avail_width;
  protobuf_c_boolean has_avail_height;
  uint32_t avail_height;
  protobuf_c_boolean has_avail_left;
  uint32_t avail_left;
  protobuf_c_boolean has_avail_top;
  uint32_t avail_top;
  protobuf_c_boolean has_outer_width;
  uint32_t outer_width;
  protobuf_c_boolean has_outer_height;
  uint32_t outer_height;
  protobuf_c_boolean has_inner_width;
  uint32_t inner_width;
  protobuf_c_boolean has_inner_height;
  uint32_t inner_height;
};
#define BROWSER_SCREEN_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&browser_screen_info__descriptor) \
    , 0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0 }


struct  _MesgBrowserEnv
{
  ProtobufCMessage base;
  BrowserPluginInfo *plugin;
  BrowserUseragentInfo *ua;
  BrowserScreenInfo *screen;
};
#define MESG_BROWSER_ENV__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_browser_env__descriptor) \
    , NULL, NULL, NULL }


struct  _MesgBrowserEnvAck
{
  ProtobufCMessage base;
  uint32_t code;
  char *errmsg;
};
#define MESG_BROWSER_ENV_ACK__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_browser_env_ack__descriptor) \
    , 0, NULL }


struct  _MesgEventStatistic
{
  ProtobufCMessage base;
  CtrlType ctrl;
  EventType event;
  uint32_t count;
};
#define MESG_EVENT_STATISTIC__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_event_statistic__descriptor) \
    , 0, 0, 0 }


struct  _MesgEventStatisticAck
{
  ProtobufCMessage base;
  uint32_t code;
  char *errmsg;
};
#define MESG_EVENT_STATISTIC_ACK__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_event_statistic_ack__descriptor) \
    , 0, NULL }


struct  _MesgLsndInfo
{
  ProtobufCMessage base;
  uint32_t type;
  uint32_t nid;
  uint32_t opid;
  char *nation;
  char *ip;
  uint32_t port;
  uint32_t connections;
};
#define MESG_LSND_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_lsnd_info__descriptor) \
    , 0, 0, 0, NULL, NULL, 0, 0 }


struct  _MesgFrwdInfo
{
  ProtobufCMessage base;
  uint32_t nid;
  char *ip;
  uint32_t forward_port;
  uint32_t backend_port;
};
#define MESG_FRWD_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_frwd_info__descriptor) \
    , 0, NULL, 0, 0 }


/* MesgOnline methods */
void   mesg_online__init
                     (MesgOnline         *message);
size_t mesg_online__get_packed_size
                     (const MesgOnline   *message);
size_t mesg_online__pack
                     (const MesgOnline   *message,
                      uint8_t             *out);
size_t mesg_online__pack_to_buffer
                     (const MesgOnline   *message,
                      ProtobufCBuffer     *buffer);
MesgOnline *
       mesg_online__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_online__free_unpacked
                     (MesgOnline *message,
                      ProtobufCAllocator *allocator);
/* MesgOnlineAck methods */
void   mesg_online_ack__init
                     (MesgOnlineAck         *message);
size_t mesg_online_ack__get_packed_size
                     (const MesgOnlineAck   *message);
size_t mesg_online_ack__pack
                     (const MesgOnlineAck   *message,
                      uint8_t             *out);
size_t mesg_online_ack__pack_to_buffer
                     (const MesgOnlineAck   *message,
                      ProtobufCBuffer     *buffer);
MesgOnlineAck *
       mesg_online_ack__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_online_ack__free_unpacked
                     (MesgOnlineAck *message,
                      ProtobufCAllocator *allocator);
/* MesgKick methods */
void   mesg_kick__init
                     (MesgKick         *message);
size_t mesg_kick__get_packed_size
                     (const MesgKick   *message);
size_t mesg_kick__pack
                     (const MesgKick   *message,
                      uint8_t             *out);
size_t mesg_kick__pack_to_buffer
                     (const MesgKick   *message,
                      ProtobufCBuffer     *buffer);
MesgKick *
       mesg_kick__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_kick__free_unpacked
                     (MesgKick *message,
                      ProtobufCAllocator *allocator);
/* BrowserPluginInfo methods */
void   browser_plugin_info__init
                     (BrowserPluginInfo         *message);
size_t browser_plugin_info__get_packed_size
                     (const BrowserPluginInfo   *message);
size_t browser_plugin_info__pack
                     (const BrowserPluginInfo   *message,
                      uint8_t             *out);
size_t browser_plugin_info__pack_to_buffer
                     (const BrowserPluginInfo   *message,
                      ProtobufCBuffer     *buffer);
BrowserPluginInfo *
       browser_plugin_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   browser_plugin_info__free_unpacked
                     (BrowserPluginInfo *message,
                      ProtobufCAllocator *allocator);
/* BrowserUseragentInfo methods */
void   browser_useragent_info__init
                     (BrowserUseragentInfo         *message);
size_t browser_useragent_info__get_packed_size
                     (const BrowserUseragentInfo   *message);
size_t browser_useragent_info__pack
                     (const BrowserUseragentInfo   *message,
                      uint8_t             *out);
size_t browser_useragent_info__pack_to_buffer
                     (const BrowserUseragentInfo   *message,
                      ProtobufCBuffer     *buffer);
BrowserUseragentInfo *
       browser_useragent_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   browser_useragent_info__free_unpacked
                     (BrowserUseragentInfo *message,
                      ProtobufCAllocator *allocator);
/* BrowserScreenInfo methods */
void   browser_screen_info__init
                     (BrowserScreenInfo         *message);
size_t browser_screen_info__get_packed_size
                     (const BrowserScreenInfo   *message);
size_t browser_screen_info__pack
                     (const BrowserScreenInfo   *message,
                      uint8_t             *out);
size_t browser_screen_info__pack_to_buffer
                     (const BrowserScreenInfo   *message,
                      ProtobufCBuffer     *buffer);
BrowserScreenInfo *
       browser_screen_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   browser_screen_info__free_unpacked
                     (BrowserScreenInfo *message,
                      ProtobufCAllocator *allocator);
/* MesgBrowserEnv methods */
void   mesg_browser_env__init
                     (MesgBrowserEnv         *message);
size_t mesg_browser_env__get_packed_size
                     (const MesgBrowserEnv   *message);
size_t mesg_browser_env__pack
                     (const MesgBrowserEnv   *message,
                      uint8_t             *out);
size_t mesg_browser_env__pack_to_buffer
                     (const MesgBrowserEnv   *message,
                      ProtobufCBuffer     *buffer);
MesgBrowserEnv *
       mesg_browser_env__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_browser_env__free_unpacked
                     (MesgBrowserEnv *message,
                      ProtobufCAllocator *allocator);
/* MesgBrowserEnvAck methods */
void   mesg_browser_env_ack__init
                     (MesgBrowserEnvAck         *message);
size_t mesg_browser_env_ack__get_packed_size
                     (const MesgBrowserEnvAck   *message);
size_t mesg_browser_env_ack__pack
                     (const MesgBrowserEnvAck   *message,
                      uint8_t             *out);
size_t mesg_browser_env_ack__pack_to_buffer
                     (const MesgBrowserEnvAck   *message,
                      ProtobufCBuffer     *buffer);
MesgBrowserEnvAck *
       mesg_browser_env_ack__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_browser_env_ack__free_unpacked
                     (MesgBrowserEnvAck *message,
                      ProtobufCAllocator *allocator);
/* MesgEventStatistic methods */
void   mesg_event_statistic__init
                     (MesgEventStatistic         *message);
size_t mesg_event_statistic__get_packed_size
                     (const MesgEventStatistic   *message);
size_t mesg_event_statistic__pack
                     (const MesgEventStatistic   *message,
                      uint8_t             *out);
size_t mesg_event_statistic__pack_to_buffer
                     (const MesgEventStatistic   *message,
                      ProtobufCBuffer     *buffer);
MesgEventStatistic *
       mesg_event_statistic__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_event_statistic__free_unpacked
                     (MesgEventStatistic *message,
                      ProtobufCAllocator *allocator);
/* MesgEventStatisticAck methods */
void   mesg_event_statistic_ack__init
                     (MesgEventStatisticAck         *message);
size_t mesg_event_statistic_ack__get_packed_size
                     (const MesgEventStatisticAck   *message);
size_t mesg_event_statistic_ack__pack
                     (const MesgEventStatisticAck   *message,
                      uint8_t             *out);
size_t mesg_event_statistic_ack__pack_to_buffer
                     (const MesgEventStatisticAck   *message,
                      ProtobufCBuffer     *buffer);
MesgEventStatisticAck *
       mesg_event_statistic_ack__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_event_statistic_ack__free_unpacked
                     (MesgEventStatisticAck *message,
                      ProtobufCAllocator *allocator);
/* MesgLsndInfo methods */
void   mesg_lsnd_info__init
                     (MesgLsndInfo         *message);
size_t mesg_lsnd_info__get_packed_size
                     (const MesgLsndInfo   *message);
size_t mesg_lsnd_info__pack
                     (const MesgLsndInfo   *message,
                      uint8_t             *out);
size_t mesg_lsnd_info__pack_to_buffer
                     (const MesgLsndInfo   *message,
                      ProtobufCBuffer     *buffer);
MesgLsndInfo *
       mesg_lsnd_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_lsnd_info__free_unpacked
                     (MesgLsndInfo *message,
                      ProtobufCAllocator *allocator);
/* MesgFrwdInfo methods */
void   mesg_frwd_info__init
                     (MesgFrwdInfo         *message);
size_t mesg_frwd_info__get_packed_size
                     (const MesgFrwdInfo   *message);
size_t mesg_frwd_info__pack
                     (const MesgFrwdInfo   *message,
                      uint8_t             *out);
size_t mesg_frwd_info__pack_to_buffer
                     (const MesgFrwdInfo   *message,
                      ProtobufCBuffer     *buffer);
MesgFrwdInfo *
       mesg_frwd_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_frwd_info__free_unpacked
                     (MesgFrwdInfo *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*MesgOnline_Closure)
                 (const MesgOnline *message,
                  void *closure_data);
typedef void (*MesgOnlineAck_Closure)
                 (const MesgOnlineAck *message,
                  void *closure_data);
typedef void (*MesgKick_Closure)
                 (const MesgKick *message,
                  void *closure_data);
typedef void (*BrowserPluginInfo_Closure)
                 (const BrowserPluginInfo *message,
                  void *closure_data);
typedef void (*BrowserUseragentInfo_Closure)
                 (const BrowserUseragentInfo *message,
                  void *closure_data);
typedef void (*BrowserScreenInfo_Closure)
                 (const BrowserScreenInfo *message,
                  void *closure_data);
typedef void (*MesgBrowserEnv_Closure)
                 (const MesgBrowserEnv *message,
                  void *closure_data);
typedef void (*MesgBrowserEnvAck_Closure)
                 (const MesgBrowserEnvAck *message,
                  void *closure_data);
typedef void (*MesgEventStatistic_Closure)
                 (const MesgEventStatistic *message,
                  void *closure_data);
typedef void (*MesgEventStatisticAck_Closure)
                 (const MesgEventStatisticAck *message,
                  void *closure_data);
typedef void (*MesgLsndInfo_Closure)
                 (const MesgLsndInfo *message,
                  void *closure_data);
typedef void (*MesgFrwdInfo_Closure)
                 (const MesgFrwdInfo *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCEnumDescriptor    ctrl_type__descriptor;
extern const ProtobufCEnumDescriptor    event_type__descriptor;
extern const ProtobufCMessageDescriptor mesg_online__descriptor;
extern const ProtobufCMessageDescriptor mesg_online_ack__descriptor;
extern const ProtobufCMessageDescriptor mesg_kick__descriptor;
extern const ProtobufCMessageDescriptor browser_plugin_info__descriptor;
extern const ProtobufCMessageDescriptor browser_useragent_info__descriptor;
extern const ProtobufCMessageDescriptor browser_screen_info__descriptor;
extern const ProtobufCMessageDescriptor mesg_browser_env__descriptor;
extern const ProtobufCMessageDescriptor mesg_browser_env_ack__descriptor;
extern const ProtobufCMessageDescriptor mesg_event_statistic__descriptor;
extern const ProtobufCMessageDescriptor mesg_event_statistic_ack__descriptor;
extern const ProtobufCMessageDescriptor mesg_lsnd_info__descriptor;
extern const ProtobufCMessageDescriptor mesg_frwd_info__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_mesg_2eproto__INCLUDED */

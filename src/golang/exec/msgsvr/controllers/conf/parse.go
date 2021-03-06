package conf

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"

	"login/src/golang/lib/log"
)

/* 日志配置 */
type MsgSvrConfLogXmlData struct {
	Name  xml.Name `xml:"LOG"`        // 结点名
	Level string   `xml:"LEVEL,attr"` // 日志级别
	Path  string   `xml:"PATH,attr"`  // 日志目录
}

/* REDIS配置 */
type MsgSvrRedisConf struct {
	Name   xml.Name `xml:"REDIS"`       // 结点名
	Addr   string   `xml:"ADDR,attr"`   // 地址(IP+端口)
	Usr    string   `xml:"USR,attr"`    // 用户名
	Passwd string   `xml:"PASSWD,attr"` // 登录密码
}

/* 鉴权配置 */
type MsgSvrConfRtmqAuthXmlData struct {
	Name   xml.Name `xml:"AUTH"`        // 结点名
	Usr    string   `xml:"USR,attr"`    // 用户名
	Passwd string   `xml:"PASSWD,attr"` // 登录密码
}

/* RTMQ代理配置 */
type MsgSvrConfRtmqProxyXmlData struct {
	Name        xml.Name                  `xml:"FRWDER"`        // 结点名
	RemoteAddr  string                    `xml:"ADDR,attr"`     // 对端IP(IP+PROT)
	Auth        MsgSvrConfRtmqAuthXmlData `xml:"AUTH"`          // 鉴权信息
	WorkerNum   uint32                    `xml:"WORKER-NUM"`    // 协程数
	SendChanLen uint32                    `xml:"SEND-CHAN-LEN"` // 发送队列长度
	RecvChanLen uint32                    `xml:"RECV-CHAN-LEN"` // 接收队列长度
}

/* 在线中心XML配置 */
type MsgSvrConfXmlData struct {
	Name   xml.Name                   `xml:"MSGSVR"`  // 根结点名
	Id     uint32                     `xml:"ID,attr"` // 结点ID
	Gid    uint32                     `xml:"GID,attr"` // 分组ID
	Redis  MsgSvrRedisConf            `xml:"REDIS"`   // REDIS配置
	Cipher string                     `xml:"CIPHER"`  // 私密密钥
	Log    MsgSvrConfLogXmlData       `xml:"LOG"`     // 日志配置
	Frwder MsgSvrConfRtmqProxyXmlData `xml:"FRWDER"`  // RTMQ PROXY配置
}

/******************************************************************************
 **函数名称: parse
 **功    能: 解析配置信息
 **输入参数: NONE
 **输出参数: NONE
 **返    回:
 **     err: 错误描述
 **实现描述: 加载配置并提取有效信息
 **注意事项:
 **作    者: # Qifeng.zou # 2016.10.30 22:35:28 #
 ******************************************************************************/
func (conf *MsgSvrConf) parse() (err error) {
	/* > 加载配置文件 */
	file, err := os.Open(conf.ConfPath)
	if nil != err {
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if nil != err {
		return err
	}

	node := MsgSvrConfXmlData{}

	err = xml.Unmarshal(data, &node)
	if nil != err {
		return err
	}

	/* > 解析配置文件 */
	/* 结点ID */
	conf.Id = node.Id
	if 0 == conf.Id {
		return errors.New("Get node id failed!")
	}

	/* 分组ID */
	conf.Gid = node.Gid
	if 0 == conf.Gid {
		return errors.New("Get gid failed!")
	}

	/* Redis地址(IP+PORT) */
	conf.Redis.Addr = node.Redis.Addr
	if 0 == len(conf.Redis.Addr) {
		return errors.New("Get redis addr failed!")
	}

	conf.Redis.Usr = node.Redis.Usr
	conf.Redis.Passwd = node.Redis.Passwd

	/* > 私密密钥 */
	conf.Cipher = node.Cipher
	if 0 == len(conf.Cipher) {
		return errors.New("Get cipher failed!")
	}

	/* 日志配置 */
	conf.Log.Level = log.GetLevel(node.Log.Level)

	conf.Log.Path = node.Log.Path
	if 0 == len(conf.Log.Path) {
		return errors.New("Get log path failed!")
	}

	/* FRWDER配置 */
	conf.Frwder.Id = conf.Id
	conf.Frwder.Gid = conf.Gid

	/* 鉴权信息 */
	conf.Frwder.Usr = node.Frwder.Auth.Usr
	conf.Frwder.Passwd = node.Frwder.Auth.Passwd
	if 0 == len(conf.Frwder.Usr) || 0 == len(conf.Frwder.Passwd) {
		return errors.New("Get auth conf failed!")
	}

	/* 转发层(IP+PROT) */
	conf.Frwder.RemoteAddr = node.Frwder.RemoteAddr
	if 0 == len(conf.Frwder.RemoteAddr) {
		return errors.New("Get Frwder addr failed!")
	}

	/* 发送队列长度 */
	conf.Frwder.SendChanLen = node.Frwder.SendChanLen
	if 0 == conf.Frwder.SendChanLen {
		return errors.New("Get send channel length failed!")
	}

	/* 接收队列长度 */
	conf.Frwder.RecvChanLen = node.Frwder.RecvChanLen
	if 0 == conf.Frwder.RecvChanLen {
		return errors.New("Get recv channel length failed!")
	}

	/* 协程数 */
	conf.Frwder.WorkerNum = node.Frwder.WorkerNum
	if 0 == conf.Frwder.WorkerNum {
		return errors.New("Get worker number failed!")
	}

	return nil
}

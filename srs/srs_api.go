// Package srs
// @title
// @description
// @author njy
// @since 2023/6/1 14:05
package srs

/*
https://ossrs.net/lts/zh-cn/docs/v4/doc/http-api
地址是：http://192.168.1.170:1985/api/v1，主要包含的子api有：
/api/v1/versions	获取服务器版本信息
/api/v1/summaries	获取服务器的摘要信息
/api/v1/rusages	获取服务器资源使用信息
/api/v1/self_proc_stats	获取服务器进程信息
/api/v1/system_proc_stats	获取服务器所有进程情况
/api/v1/meminfos	获取服务器内存使用情况
/api/v1/authors	获取作者、版权和License信息
/api/v1/features	获取系统支持的功能列表
/api/v1/requests	获取请求的信息，即当前发起的请求的详细信息
/api/v1/vhosts	获取服务器上的vhosts信息
/api/v1/streams	获取服务器的streams信息
/api/v1/clients	获取服务器的clients信息，默认获取前10个
/api/v1/configs	CUID配置，RAW API
/rtc/v1/publish/	WebRTC推流的API
/rtc/v1/play/	WebRTC播放流的API
*/

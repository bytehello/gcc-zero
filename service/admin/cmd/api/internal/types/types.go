// Code generated by goctl. DO NOT EDIT.
package types

type AppAddReq struct {
	AppName string `json:"appName"`
	AppKey  string `json:"appKey"`
	Desc    string `json:"desc"`
}

type AppAddReply struct {
	Code    string `json:"code"`
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type AppListReq struct {
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize,default=5"`
}

type ListAppData struct {
	Id      int64  `json:"id"`
	AppName string `json:"appName"`
	AppKey  string `json:"appKey"`
	Desc    string `json:"desc"`
}

type AppListReply struct {
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	Data     []*ListAppData `json:"data"`
	Current  int64          `json:"current"`
	PageSize int64          `json:"pageSize"`
	Total    int64          `json:"total"`
}

type AppUpdateReq struct {
	Id      int64  `json:"id"`
	AppName string `json:"appName"`
	AppKey  string `json:"appKey"`
	Desc    string `json:"desc"`
}

type AppUpdateReply struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AppDelReq struct {
	Id int64 `json:"id"`
}

type AppDelReply struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ClusterAddReq struct {
	AppId       int64  `json:"appId"`
	ClusterName string `json:"clusterName"`
	Desc        string `json:"desc"`
}

type ClusterAddReply struct {
	Code    string `json:"code"`
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type ClusterDelReq struct {
	Id int64 `json:"id"`
}

type ClusterDelReply struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ClusterUpdateReq struct {
	Id          int64  `json:"id"`
	ClusterName string `json:"clusterName"`
	Desc        string `json:"desc"`
}

type ClusterUpdateReply struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ClusterListReq struct {
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
}

type ClusterListData struct {
	Id          int64  `json:"id"`
	AppId       int64  `json:"appId"`
	ClusterName string `json:"clusterName"`
	Desc        string `json:"desc"`
}

type ClusterListReply struct {
	Code     string             `json:"code"`
	Message  string             `json:"message"`
	Data     []*ClusterListData `json:"data"`
	Current  int64              `json:"current"`
	PageSize int64              `json:"pageSize"`
	Total    int64              `json:"total"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Code         string `json:"code"`
	Msg          string `json:"msg"`
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

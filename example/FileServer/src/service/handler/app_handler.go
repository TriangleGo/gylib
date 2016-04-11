package handler

import (
	"gycache/message"
	"gyparam"
	"gyservice/respcode"
	"db/dao"
	"db/entity"
	"gylogger"
)

func CheckAppVersion(req map[string]interface{}) (resp *message.Response) {
	resp = message.NewResponse()

	platform, okPlatform := param.GetStringWithKey(req, "platform")
	version, okVersion := param.GetStringWithKey(req, "version")
	if !(okPlatform && okVersion) {
		resp.SetRespCode(respcode.RC_GENERAL_PARAM_ERR)
		resp.SetParam("platform", okPlatform)
		resp.SetParam("version", okVersion)
		return
	}

	v, s, i, c, err := entity.ParseAppVersionStr(version)
	logger.Infof("parse input version string: v = %s, s = %s, i = %s, c = %s, err = %v.", v, s, i, c, err)
	if err != nil {
		resp.SetRespCode(respcode.RC_GENERAL_SYS_ERR)
		resp.SetParam("error", err.Error())
		return
	}
	valid := dao.ValidAppVersion(platform, v, s, i, c)

	resp.SetRespCode(respcode.RC_GENERAL_SUCC)
	resp.SetParam("valid", valid)

	return
}
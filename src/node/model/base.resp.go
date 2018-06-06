package model

type BaseResp struct {
	Code string
	Info string
}

func (resp *BaseResp) SetRC(rc *BaseResp, err ... error) {
	if resp == nil || rc == nil {
		return
	}
	resp.Code = rc.Code
	if len(err) > 0 {
		resp.Info = err[0].Error()
	} else {
		resp.Info = rc.Info
	}
}

var (
	RC_SUCC = &BaseResp{"RC00000", "请求成功。"}
	RC_FAIL = &BaseResp{"RC99999", "请求失败。"}
)

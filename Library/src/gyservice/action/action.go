package action

type Action int32

type Service_range struct {
	Start int32
	End   int32
	Node  string
}

const DEFAULT_SERVICE_PATH = "/gy"

var (
	File_service_range *Service_range = &Service_range{Start:0, End:100, Node:DEFAULT_SERVICE_PATH + "/file"}
)

// 定义File模块的接口名
// 范围 [9001 ~ 9500]
const (
	Action_SaveFile Action = 2
	Action_LoadFile Action = 3
	Action_DeleteFile Action = 4

	Action_CheckAppVersion Action = 20
)

var codeToAction = map[int32]Action{

	// File 9000 ~ 9500
	int32(Action_SaveFile):Action_SaveFile,
	int32(Action_LoadFile):Action_LoadFile,
	int32(Action_DeleteFile):Action_DeleteFile,
	int32(Action_CheckAppVersion):Action_CheckAppVersion,
}

var nameToAction = map[string]Action{

	// File 9000 ~ 9500
	"saveFile":Action_SaveFile,
	"loadFile":Action_LoadFile,
	"deleteFile":Action_DeleteFile,
	"checkAppVersion":Action_CheckAppVersion,
}

var actionToName = map[Action]string{
	// File 9000 ~ 9500
	Action_SaveFile:"saveFile",
	Action_LoadFile:"loadFile",
	Action_DeleteFile:"deleteFile",
	Action_CheckAppVersion:"checkAppVersion",
}

func ActionFromCode(code int32) Action {
	return codeToAction[code]
}

func ActionFromName(name string) (v Action, ok bool) {
	v, ok = nameToAction[name]
	return
}

func (a Action)Name() string {
	return actionToName[a]
}
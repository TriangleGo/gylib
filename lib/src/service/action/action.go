package action

type Action int32

type Service_range struct {
	Start int32
	End   int32
	Node  string
}

const DEFAULT_SERVICE_PATH = "/baby"

var (
	File_service_range *Service_range = &Service_range{Start:0, End:100, Node:DEFAULT_SERVICE_PATH + "/file"}
	Library_service_range *Service_range = &Service_range{Start:1000, End:2000, Node:DEFAULT_SERVICE_PATH + "/library"}
	Profile_service_range *Service_range = &Service_range{Start:5000, End:6000, Node:DEFAULT_SERVICE_PATH + "/profile"}
)

// 定义File模块的接口名
// 范围 [0, 100)
const (
	Action_SaveFile Action = 2
	Action_LoadFile Action = 3
	Action_DeleteFile Action = 4
)

// 定义Profile模块接口名
// 范围 [5000, 6000)
const (
	Action_CheckPhoneNo Action = 5001
	Action_ApplyAuthCode Action = 5002
	Action_ResetPassword Action = 5003

	Action_Login Action = 5101
)

var codeToAction = map[int32]Action{

	// File 9000 ~ 9500
	int32(Action_SaveFile):Action_SaveFile,
	int32(Action_LoadFile):Action_LoadFile,
	int32(Action_DeleteFile):Action_DeleteFile,
}

var nameToAction = map[string]Action{
	// File [0, 100)
	"saveFile":Action_SaveFile,
	"loadFile":Action_LoadFile,
	"deleteFile":Action_DeleteFile,

	// Profile [5000, 6000)
	"checkPhoneNo":Action_CheckPhoneNo,
}

var actionToName = map[Action]string{
	// File 9000 ~ 9500
	Action_SaveFile:"saveFile",
	Action_LoadFile:"loadFile",
	Action_DeleteFile:"deleteFile",

	// Profile [5000, 6000)
	Action_CheckPhoneNo:"checkPhoneNo",
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
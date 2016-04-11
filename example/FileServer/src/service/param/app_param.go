package param

func GetPlatform(req map[string]interface{}) (platform string, ok bool) {
	var tempPlatform interface{}
	tempPlatform, ok = req["platform"]
	if ok {
		platform = tempPlatform.(string)
	}
	return
}

func GetVersion(req map[string]interface{}) (version string, ok bool) {
	var tempVersion interface{}
	tempVersion, ok = req["version"]
	if ok {
		version = tempVersion.(string)
	}
	return
}

func GetFirmwareVersion(req map[string]interface{}) (version string, ok bool) {
	var tempVersion interface{}
	tempVersion, ok = req["firmware_version"]
	if ok {
		version = tempVersion.(string)
	}
	return
}

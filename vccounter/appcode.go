package vccounter

type AppCode struct {
	AppId       string
	VersionCode int
}

func (ac *AppCode) NextVersionCode() int {

	ac.VersionCode++

	return ac.VersionCode
}

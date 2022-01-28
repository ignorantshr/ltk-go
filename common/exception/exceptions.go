package exception

type GetConfigFailedException struct {
	template string
	Err      error
}

func GetConfigFailed(err error) *GetConfigFailedException {
	return &GetConfigFailedException{
		template: "Failed to get config file",
		Err:      err,
	}
}

func (g *GetConfigFailedException) Error() string {
	if g.Err != nil {
		return g.template + " : " + g.Err.Error()
	}
	return g.template
}

type ParseConfigFailedException struct {
	template string
	Err      error
}

func ParseConfigFailed(err error) *ParseConfigFailedException {
	return &ParseConfigFailedException{
		template: "Failed to parse config file",
		Err:      err,
	}
}

func (g *ParseConfigFailedException) Error() string {
	if g.Err != nil {
		return g.template + " : " + g.Err.Error()
	}
	return g.template
}

type GetEngineConnFailedException struct {
	template string
	Err      error
}

func GetEngineConnFailed(err error) *GetEngineConnFailedException {
	return &GetEngineConnFailedException{
		template: "Failed to get engine connection information",
		Err:      err,
	}
}

func (g *GetEngineConnFailedException) Error() string {
	if g.Err != nil {
		return g.template + " : " + g.Err.Error()
	}
	return g.template
}

type ConnEngineFailedException struct {
	template string
	Err      error
}

func ConnEngineFailed(err error) *ConnEngineFailedException {
	return &ConnEngineFailedException{
		template: "Couldn't connect to lidc engine",
		Err:      err,
	}
}

func (g *ConnEngineFailedException) Error() string {
	if g.Err != nil {
		return g.template + " : " + g.Err.Error()
	}
	return g.template
}

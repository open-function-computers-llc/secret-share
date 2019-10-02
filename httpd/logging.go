package main

func logInfo(i ...interface{}) {
	conf.Logger.Info(i...)
}

func logError(i ...interface{}) {
	conf.Logger.Error(i...)
}

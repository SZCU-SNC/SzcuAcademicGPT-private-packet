使用文档

此log日志记录工具使用logrus库和lumberjack库实现，logrus库是一个日志库，lumberjack库是一个日志切割库，可以实现日志的切割和归档。
使用此工具时，需先指定serviceName并创建新logger，若记录的是普通信息使用logger.WithCaller().Info()，记录错误信息和警告信息使用logger.WithCaller().Error()和logger.WithCaller().Warn()。示例可见log.go中注释内容。
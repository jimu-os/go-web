package pub

import (
	"api/docs"
	"common/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	/*
		Swag init -h
		名称:
		swag init -创建文档。go

		用法:
		Swag init[命令选项][参数…]

		选项:
		——quiet， -q使记录器静音。(默认值:false)
		` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` ` `
		——dir value， -d value			要解析的目录，逗号分隔和general-info file必须在第一个目录中(默认:"./")
		——exclude value					搜索时排除目录和文件，以逗号分隔属性命名策略如snakcase,camelcase,pascalcase(默认:"camelcase")
		——output value， -o value所有生成文件的输出目录。json,昂首阔步。Yaml和docs.go)(默认:"./docs")
		——outputTypes value，——ot value生成文件的输出类型。去,昂首阔步。Json, swagger.yaml)像go, Json,yaml(默认值:"go, Json,yaml")
		解析'vendor'文件夹中的go文件，默认禁用(default: false)
		——parseDependency，——pd解析依赖文件夹中的go文件，默认禁用(default: false)
		——markdownFiles value，——md value解析包含markdown文件的文件夹，用作描述，默认禁用
		包含用于x- codessamples扩展的代码示例文件的解析文件夹，默认情况下禁用
		解析内部包中的go文件，默认禁用(default: false)
		——generatedTime在文档顶部生成时间戳。Go，默认禁用(default: false)
		——parseDepth value依赖解析深度(默认:100)
		——requiredByDefault默认设置所有字段都需要验证(default: false)
		——instanceName value该参数可用于命名不同的swagger文档实例。这是可选的。
		——overridesFile value读取全局类型覆盖的文件。(默认:“.swaggo”)
		——parseGoList通过'go list'解析依赖(默认值:true)
		——tags value， -t value一个逗号分隔的标签列表，用于过滤生成文档的api。如果标签的前缀是'!'字符，那么带有该标记的api将被排除
		为Go模板生成提供自定义分隔符。格式为leftDelim,rightDelim。例如:“[[，]]”
		——collectionFormat value，——cf value设置默认的采集格式(默认:"csv")
		——state value状态机的初始状态(默认:"")，根文件中的@HostState，其他文件中的@State
		解析go文件中函数体中的API信息，默认禁用(default: false)
		——help， -h show help(默认:false)
	*/
	// https://github.com/swaggo/swag
	// 初始化文档使用 swag init -o docs -g main.go
	docs.SwaggerInfo.Title = "web-kit"
	docs.SwaggerInfo.Description = "api 文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	web.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := web.Engine.Group("/api")
	group.POST("/login", Login)
}

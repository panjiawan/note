package main

import (
	yamlgen "github.com/we7coreteam/gorm-gen-yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	//FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./service/dao/query",                                              // 定义 dao 文件输出目录
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./model",                                                          // 定义 model 文件输出目录
	})

	// 配置数据库连接信息
	gormdb, _ := gorm.Open(mysql.Open("root:a17bdf010ae88d44@(localhost:3306)/frame_dev?charset=utf8mb4&parseTime=True&loc=Local"))
	//启用数据库连接
	g.UseDB(gormdb) // reuse your gorm db

	fieldOpts := []gen.ModelOpt{}
	yamlgen.NewYamlGenerator("./gen.yaml").UseGormGenerator(g).Generate(fieldOpts...)
	//g.ApplyBasic(g.GenerateAllTable(fieldOpts...)...)
	//g.GenerateModel(tableName)
	g.Execute()
}

/*
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./service/dao/query",                                              // 定义 dao 文件输出目录
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./model",                                                          // 定义 model 文件输出目录
	})

	// 配置数据库连接信息
	gormdb, _ := gorm.Open(mysql.Open("root:a123456@(localhost:3307)/daweike_dev?charset=utf8mb4&parseTime=True&loc=Local"))
	//启用数据库连接
	g.UseDB(gormdb) // reuse your gorm db

	//生成单个表的model 若只为了生成model则不需要接受参数，如下生成所有表model的示例
	adminUserModel := g.GenerateModel("tbl_admin_users")

	//生成所有表的model
	//g.GenerateAllTable()

	//根据model生成dao文件
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(adminUserModel)

	//根据所有model生成dao文件
	//g.ApplyBasic(g.GenerateAllTable()...)

	//对于已经存在model的情况，可以传入该model的实例，如下所示
	//g.ApplyBasic(model.User{})

	//根据接口生成自定义方法
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, adminUserModel)

	//执行
	// Generate the code
	g.Execute()
}

*/

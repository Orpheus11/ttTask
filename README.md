# tantan Task
connor

修改config.conf文件，将PostgreSQL数据库的用户名和密码换成需要连接的数据库的用户名和密码

有两种运行方式：

	1）是在Windows下，双击运行ttTask.exe（注意执行的时候必须和config.conf在同一个目录）
	2）源码运行，则需要	
		"github.com/goconf/conf"
		"github.com/gorilla/mux"
		"gopkg.in/pg.v3"
		这三个外部的包(如果没有，则执行“go get 包名”)，将源码copy到gopath下的src下，执行go install，然后执行go run main.go


测试效果截图：TanTanTask.png


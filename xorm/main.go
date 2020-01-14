package main
import("time"
"github.com/kataras/iris/v12"
"github.com/go-xorm/xorm"
_"github.com/mattn/go-sqlite3"
)
type User struct{
	ID int64 
	Version string `xorm:"varchar(200)"`
	Salt string
	Username string
	Password string `xorm:"varchar(200)"`
	Languages string `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
func main(){
	app:=iris.New()
	orm,err:=xorm.NewEngine("sqlite3","./test.db")
	if err!=nil{
		app.Logger().Fatalf("orm failed to initialized: %v",err)}
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err=orm.Sync2(new(User))
	if err!=nil{
		app.Logger().Fatalf("orm failed to initialized User table: %v",err)
	}
	app.Get("/insert",func(ctx iris.Context){
		user:=&User{Username: "kataras", Salt: "hash---", Password: "hashed", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		orm.Insert(user)
		ctx.Writef("user inserted: %#v", user)
	})
	app.Get("/get",func(ctx iris.Context){
		user :=User{ID:1}
		if ok,_:=orm.Get(&user);ok{
			ctx.Writef("user found: %#v", user)
		}
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

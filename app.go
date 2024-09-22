package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"math"
	"math/rand"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type LoginInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin bool   `json:"auto_login"`
	Type      string `json:"type"`
}

func (a *App) Login(in *LoginInput) map[string]string {
	m := make(map[string]string)
	if in.Password == "ant.design" && in.Username == "admin" {
		//res.send({
		//status: 'ok',
		//	type,
		//currentAuthority: 'admin',
		//});
		//access = 'admin';
		//return;

		m["status"] = "ok"
		m["type"] = in.Type
		m["currentAuthority"] = "admin"
	}
	return m
}

func (a *App) GetCurrentUser() map[string]any {
	s := `
{
    "data": {
        "access": "admin",
        "address": "西湖区工专路 77 号",
        "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
        "country": "China",
        "email": "antdesign@alipay.com",
        "geographic": {
            "city": {
                "key": "330100",
                "label": "杭州市"
            },
            "province": {
                "key": "330000",
                "label": "浙江省"
            }
        },
        "group": "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
        "name": "Serati Ma",
        "notifyCount": 12,
        "phone": "0752-268888888",
        "signature": "海纳百川，有容乃大",
        "tags": [
            {
                "key": "0",
                "label": "很有想法的"
            },
            {
                "key": "1",
                "label": "专注设计"
            },
            {
                "key": "2",
                "label": "辣~"
            },
            {
                "key": "3",
                "label": "大长腿"
            },
            {
                "key": "4",
                "label": "川妹子"
            },
            {
                "key": "5",
                "label": "海纳百川"
            }
        ],
        "title": "交互专家",
        "unreadCount": 11,
        "userid": "00000001"
    },
    "success": true
}
`
	return gjson.New(s).MapStrAny()
}

func (a *App) GetRule(page int, pageSize int) map[string]any {
	list := make([]map[string]any, 0)

	start := (page - 1) * pageSize
	end := start + pageSize

	for i := start; i <= end; i++ {
		m := make(map[string]any)
		m["key"] = i
		m["disabled"] = i%6 == 0
		m["href"] = "https://ant.design"
		m["avatar"] = []string{"https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png", "https://gw.alipayobjects.com/zos/rmsportal/udxAbMEhpwthVVcjLXik.png"}[i%2]
		m["name"] = fmt.Sprintf("TradeCode %v", i)
		m["owner"] = "曲丽丽"
		m["desc"] = "这是一段描述"
		m["callNo"] = math.Floor(float64(rand.Float32() * 1000.0))
		m["status"] = int(math.Floor(float64(rand.Float32()*10.0))) % 4
		m["updatedAt"] = gtime.Now().Layout(time.DateOnly)
		m["createdAt"] = gtime.Now().Layout(time.DateOnly)
		m["progress"] = math.Ceil(float64(rand.Float32() * 1000.0))

		list = append(list, m)
	}

	//{
	//key: index,
	//	disabled: i % 6 === 0,
	//	href: 'https://ant.design',
	//	avatar: [
	//'https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png',
	//'https://gw.alipayobjects.com/zos/rmsportal/udxAbMEhpwthVVcjLXik.png',
	//][i % 2],
	//name: `TradeCode ${index}`,
	//owner: '曲丽丽',
	//desc: '这是一段描述',
	//callNo: Math.floor(Math.random() * 1000),
	//status: Math.floor(Math.random() * 10) % 4,
	//updatedAt: dayjs().format('YYYY-MM-DD'),
	//createdAt: dayjs().format('YYYY-MM-DD'),
	//progress: Math.ceil(Math.random() * 100),
	//}

	//{
	//data: dataSource,
	//	total: tableListDataSource.length,
	//	success: true,
	//	pageSize,
	//	current: parseInt(`${params.current}`, 10) || 1,
	//}

	mm := map[string]any{
		"data":    list,
		"total":   len(list) * 10,
		"success": true,
		"current": page,
	}

	return gvar.New(mm).MapStrAny()
}

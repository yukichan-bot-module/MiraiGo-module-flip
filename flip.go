package flip

import (
	"math/rand"
	"sync"
	"time"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

var instance *flip

type flip struct {
}

func init() {
	instance = &flip{}
	bot.RegisterModule(instance)
}

func (f *flip) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "com.aimerneige.flip",
		Instance: instance,
	}
}

// Init 初始化过程
// 在此处可以进行 Module 的初始化配置
// 如配置读取
func (f *flip) Init() {
}

// PostInit 第二次初始化
// 再次过程中可以进行跨 Module 的动作
// 如通用数据库等等
func (f *flip) PostInit() {
}

// Serve 注册服务函数部分
func (f *flip) Serve(b *bot.Bot) {
	b.GroupMessageEvent.Subscribe(func(c *client.QQClient, msg *message.GroupMessage) {
		if msg.ToString() == "掷硬币" {
			replyMsg := message.NewSendingMessage().Append(message.NewAt(msg.Sender.Uin))
			rand.Seed(time.Now().UnixNano())
			if (rand.Int() % 2) == 0 {
				c.SendGroupMessage(msg.GroupCode, replyMsg.Append(message.NewText("掷出了正面。")))
			} else {
				c.SendGroupMessage(msg.GroupCode, replyMsg.Append(message.NewText("掷出了反面。")))
			}
		}
	})
}

// Start 此函数会新开携程进行调用
// ```go
//
//	go exampleModule.Start()
//
// ```
// 可以利用此部分进行后台操作
// 如 http 服务器等等
func (f *flip) Start(b *bot.Bot) {
}

// Stop 结束部分
// 一般调用此函数时，程序接收到 os.Interrupt 信号
// 即将退出
// 在此处应该释放相应的资源或者对状态进行保存
func (f *flip) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
}

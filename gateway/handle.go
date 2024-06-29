package gateway

import (
	"HeroServer/db"
	"HeroServer/gamecfg"
	"HeroServer/service"
	"HeroServer/service/mail"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var version string

var serverListener *http.Server
var PlayerManager *service.PlayerManager

func HandleHttp(PM *service.PlayerManager, v string, html string, static embed.FS) {
	PlayerManager = PM
	version = v
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(cors.Default())
	initStatic(r, html, static)
	serverRouter := r.Group("/server/")
	{
		serverRouter.GET("status", func(c *gin.Context) {
			m := new(runtime.MemStats)
			runtime.ReadMemStats(m)
			//runtime.CPUProfile()
			c.JSON(http.StatusOK, gin.H{
				"version": version,
			})
		})

		serverRouter.GET("restart", restart)
		serverRouter.GET("reload", reload)
		serverRouter.GET("playerlist", playlist)
		serverRouter.POST("sendMail", sendMail)
	}

	serverListener = &http.Server{
		Addr:    ":6021",
		Handler: r,
	}

	err := serverListener.ListenAndServe()
	if err != nil {
		return
	}
}

func initStatic(r *gin.Engine, html string, static embed.FS) {
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, html)
	})

	r.GET("/assets/:filepath", func(c *gin.Context) {
		assetsFS, err := fs.Sub(static, "gateway/html")
		if err != nil {
			fmt.Println("初始化静态资源出错")
			return
		}
		staticServer := http.FileServer(http.FS(assetsFS))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
}

func sendMail(c *gin.Context) {
	body, _ := c.GetRawData()
	//fmt.Println(string(body))
	type Req struct {
		Rid     string `json:"Rid"`
		Title   string `json:"Title"`
		Content string `json:"Content"`
		Items   string `json:"Items"`
	}
	var req Req
	err := json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "解析参数失败",
		})
		return
	}
	fmt.Println(req.Rid, req.Title)
	rid, err := strconv.Atoi(req.Rid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "RID错误",
		})
		return
	}
	var items [][]int
	err = json.Unmarshal([]byte(req.Items), &items)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "附件格式错误",
		})
		return
	}
	for i, _ := range items {
		if len(items[i]) != 2 {
			c.JSON(http.StatusOK, gin.H{
				"msg": "附件格式错误",
			})
			return
		} else {
			items[i] = append(items[i], 0)
		}
	}
	itemsByte, err := json.Marshal(items)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "附件格式错误",
		})
		return
	}
	req.Items = string(itemsByte)
	addMail := &db.Mail{
		Rid:        rid,
		Status:     0,
		TemplateId: 1,
		Title:      req.Title,
		Content:    req.Content,
		MailItems:  req.Items,
		Time:       int(time.Now().Unix() + 86400*30),
	}
	err = db.Conn.Create(addMail).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "发送失败",
		})
	} else {
		if PlayerManager.OnLinePlayer[uint(rid)] != nil {
			PlayerManager.OnLinePlayer[uint(rid)].MPtMailNewInfo(addMail)
			mail.SendMailCount(PlayerManager.OnLinePlayer[uint(rid)])
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "发送成功",
		})
	}
}

func playlist(c *gin.Context) {
	var data = []gin.H{}
	for _, e := range PlayerManager.OnLinePlayer {
		data = append(data, gin.H{
			"Rid":  e.Rid,
			"Name": e.RoleData.Role.Name,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": data,
	})
}

func restart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "服务器5秒后进行重启",
	})

	go func() {
		time.Sleep(5 * time.Second)

		if serverListener != nil {
			err := serverListener.Shutdown(context.Background())
			if err != nil {
				fmt.Println("关闭HTTP服务失败:", err)
			} else {
				fmt.Println("已关闭HTTP服务")
			}
		}

		if PlayerManager != nil && PlayerManager.TcpListener != nil {
			err := PlayerManager.TcpListener.Close()
			if err != nil {
				fmt.Println("关闭TCP服务失败:", err)
			} else {
				fmt.Println("已关闭TCP服务")
			}
		}

		// 重启服务器
		executable, err := os.Executable()
		if err != nil {
			fmt.Println("获取程序目录失败:", err)
			return
		}

		args := os.Args[1:]
		cmd := exec.Command(executable, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Start()
		if err != nil {
			fmt.Println("重启失败:", err)
			return
		}

		fmt.Println("重启中...")
		os.Exit(0)
	}()
}

func reload(c *gin.Context) {
	gamecfg.GameConf.LoadCfg()
	c.JSON(http.StatusOK, gin.H{
		"msg": "Game配置已重载",
	})
}

package api

import (
	"down-date-server/src/global"
	"down-date-server/src/middle"
	"down-date-server/src/model"
	"down-date-server/src/model/request"
	"down-date-server/src/utils"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户注册
func Register(c *gin.Context) {
	var R model.User
	err := c.ShouldBind(&R)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "FAILED",
			"detail": "注册信息有误",
		})
		return
	}
	u := &model.User{Username: R.Username, Password: R.Password, Nickname: R.Nickname, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	var user model.User
	if !errors.Is(global.DB.Where("userName = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		log.Println("用户名已注册")
		c.JSON(400, gin.H{
			"status": "FAILED",
			"detail": "用户名已注册",
		})
	} else {
		u.Password = utils.MD5V([]byte(u.Password))
		log.Println("register password:", u.Password)
		u.UUID = uuid.NewV4()
		err = global.DB.Create(&u).Error
		if err != nil {
			c.JSON(400, gin.H{
				"status": "FAILED",
				"detail": "注册失败",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "SUCCESS",
				"detail": "用户注册成功",
			})
		}
	}

}

type LoginUser struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 用户登录
func Login(c *gin.Context) {
	var L LoginUser
	if err := c.ShouldBind(&L); err != nil {
		c.JSON(400, gin.H{
			"status": "FAILED",
			"detail": "用户名/密码/验证码不能为空",
		})
		c.Abort()
	} else {
		// 可以过滤一遍输入，避免SQL注入
		U := &model.User{Username: L.Username, Password: L.Password}
		var user model.User
		U.Password = utils.MD5V([]byte(U.Password))
		log.Println("login password:", U.Password)
		err := global.DB.Where("username = ? AND password = ?", U.Username, U.Password).First(&user).Error
		if err != nil {
			c.JSON(400, gin.H{
				"status": "FAILED",
				"detail": "用户名/密码/错误",
			})
			c.Abort()
		} else {
			tokenNext(c, user)
		}

	}
}

func tokenNext(c *gin.Context, user model.User) {
	j := middle.JWT{SignKey: []byte(global.CONFIG.JWT.SignKey)}
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.Nickname,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.CONFIG.JWT.SignKey,                         // 签名的发行者
		},
	}
	token, err := j.CreateJWT(claims)
	if err != nil {
		log.Println("获取token失败")
		c.JSON(400, gin.H{
			"status": "FAILED",
			"detail": "获取token失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"status": "SUCCESS",
			"detail": "登录成功",
			"token":  token,
		})
		return
	}

}

// 常用查询结构可以保存起来
type Result struct {
	Nickname string
}

func SearchUser(c *gin.Context) {
	var result []Result
	err := global.DB.Find(&[]model.User{}).Scan(&result).Error
	if err != nil {
		c.JSON(404, gin.H{
			"status": "FAILED",
			"detail": "查找失败",
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status": "SUCCESS",
			"detail": result,
		})
	}
}

func SearchUserByNickname(c *gin.Context) {
	var result []Result
	nickname, _ := c.GetQuery("nickname")
	err := global.DB.Where("nickname = ?", nickname).Find(&[]model.User{}).Scan(&result).Error
	if err != nil {
		c.JSON(404, gin.H{
			"status": "FAILED",
			"detail": "查找失败",
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status": "SUCCESS",
			"detail": result,
		})
	}
}

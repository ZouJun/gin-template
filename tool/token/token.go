package token

import (
	"different-industry-alliance/env"
	"different-industry-alliance/tool"
	"different-industry-alliance/tool/time_helper"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type JwtTokenData struct {
	User        JwtTokenUser   //user对象
	Member      JwtTokenMember //member对象
	CurrentRole CurrentRole
	PcObject    JwtTokenPcObject //pc用户对象
}

type JwtTokenUser struct {
	Id                    string //当前用户ID
	CurrentOpenid         string //当前用户对应的openid
	LastOperateMerchantId string //当前用户最后一次使用的商户ID
}

type JwtTokenMember struct {
	Id         string //当前用户对应的会员ID
	MerchantId string //会员所属商户ID
}

type CurrentRole struct {
	BusinessId   string //当前使用的对象ID
	BusinessType string //当前对象的身份(user/merchant/operating_enterprise)
}

type JwtTokenStruct struct {
	JwtTokenData
	jwt.StandardClaims
}

type JwtTokenPcObject struct {
	Id                      string //对象ID
	Name                    string //对象名称
	RoleType                string //角色类型(merchant/operating_enterprise)
	OperatingEnterpriseId   string //对象对应的运营企业ID
	OperatingEnterpriseType string //对象对应的运营企业类型
}

//解析token
func ParseToken(token string) *jwt.Token {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.TokenSecret), nil
	})
	if err != nil {
		return nil
	}

	return claim
}

//签发token
func SignToken(jwtTokenData JwtTokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtTokenStruct{
		jwtTokenData,
		jwt.StandardClaims{
			ExpiresAt: time_helper.GetShiftByTimeUnix(time_helper.GetCurrentTimeUnix(), "12h").Unix(),
			Issuer:    "platform",
			Subject:   "user_token",
		},
	})

	return token.SignedString([]byte(env.TokenSecret))
}

//获取加密结构体
func GetJwtTokenData(token string) (jwtTokenStruct JwtTokenStruct, err error) {
	if token == "" {
		err = errors.New("token为空")
	} else {
		claimToken := ParseToken(token)
		if claimToken != nil && claimToken.Valid {
			jwtTokenStruct, err = GetJwtTokenStruct(claimToken.Claims)
			if err != nil {
				return
			}

			if time.Now().Unix() > jwtTokenStruct.ExpiresAt {
				err = errors.New("token已失效")
			}
		} else {
			err = errors.New("token不合法")
		}
	}
	return
}

func GetJwtTokenStruct(data interface{}) (jwtTokenStruct JwtTokenStruct, err error) {
	var marshalResult []byte
	marshalResult, err = json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(marshalResult, &jwtTokenStruct)
	if err != nil {
		return
	}

	return
}

func GetJwtTokenDataByC(c *gin.Context) (jwtTokenStruct JwtTokenStruct, err error) {
	authToken := c.Query("token")
	if authToken == "" {
		authToken = c.GetHeader("Authorization")
	}

	return GetJwtTokenData(authToken)
}

func GetJwtTokenBusinessByRole(c *gin.Context, roleType string) (err error, jwtTokenStruct JwtTokenStruct) {
	jwtTokenStruct, err = GetJwtTokenDataByC(c)
	if err != nil {
		return
	}
	jwtTokenStruct.CurrentRole.BusinessType = roleType
	switch roleType {
	case tool.RoleTypeUser:
		jwtTokenStruct.CurrentRole.BusinessId = jwtTokenStruct.User.Id
	case tool.RoleTypeMerchant:
		jwtTokenStruct.CurrentRole.BusinessId = jwtTokenStruct.User.LastOperateMerchantId
	default:
		err = errors.New("角色不对")
	}

	return err, jwtTokenStruct
}

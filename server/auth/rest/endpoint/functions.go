package endpoint

import (
	"github.com/cosmos/go-bip39"
	"github.com/dukryung/microservice/server/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GetMnemonic(c *gin.Context) {

	data, err := bip39.NewEntropy(256)
	if err != nil {
		c.JSON(400, err)
		return
	}

	mnemonic, err := bip39.NewMnemonic(data)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, types.Mnemonic{Mnemonic: mnemonic})
}

func RegisterAccount(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.JSON(400, err)
		return
	}

	accountInfo := types.Account{
		Email:    c.PostForm("email"),
		Mnemonic: c.PostForm("mnemonic"),
		NickName: c.PostForm("nickname"),
	}

	file, _, err := c.Request.FormFile("profile-image")
	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(400, err)
		return
	}
	accountInfo.ProfileImage = data

	return
}

func VerifyAccount(c *gin.Context) {

}

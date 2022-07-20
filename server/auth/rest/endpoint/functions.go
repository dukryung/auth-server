package endpoint

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/dukryung/microservice/server/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const (
	SUCCESS = "success"
	FAIL    = "fail"
)

type EndPoint struct {
	DB *sql.DB
}

func NewEndPoint(db *sql.DB) *EndPoint {
	return &EndPoint{
		DB: db,
	}
}

func (end *EndPoint) GetMnemonic(c *gin.Context) {

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

//RegisterAccount register client's account. (TODO :  Have to add device id after discuss about method of making device)
func (end *EndPoint) RegisterAccount(c *gin.Context) {
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

	tx, err := end.DB.Begin()
	if err != nil {
		c.JSON(400, err)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(c, `INSERT INTO client_account 
											(email,mnemonic,nickname,device_token,profile_image) 
											 VALUES ($1,$2,$3,$4,$5)`)
	defer stmt.Close()

	_, err = stmt.Exec(accountInfo.Email, accountInfo.Mnemonic, accountInfo.NickName, accountInfo, "device_id", accountInfo.ProfileImage)
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, struct {
		success string
	}{"success"})

	return
}

func (end *EndPoint) VerifyAccount(c *gin.Context) {
	var clientAccount = types.Account{}
	mnemonic := c.DefaultQuery("mnemonic", "none")
	if mnemonic == "none" {
		c.JSON(400, errors.New("mnemonic empty err"))
		return
	}

	tx, err := end.DB.Begin()
	if err != nil {
		c.JSON(400, err)
		fmt.Println("err : ",err)

		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(c, `SELECT 
								email,
								mnemonic
								nickname,
								device_token,
								profile_image
								FROM client_account 
								WHERE mnemonic = $1`)
	if err != nil {
		c.JSON(400, err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(mnemonic).Scan(&clientAccount.Email, &clientAccount.Mnemonic, &clientAccount.NickName, &clientAccount.DeviceToken, &clientAccount.ProfileImage)
	if err != nil {
		c.JSON(400, err)
		return
	}
	err = tx.Commit()
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, struct {
		message string `json:"message"`
		status  int    `json:"status"`
		error   bool   `json:"error"`
		//data    struct {
		//	clientAccount types.Account
		//} `json:"data"`
		data interface{} `json:"data"`
	}{
		message: SUCCESS,
		status:  200,
		error:   false,
		//data:    struct{ clientAccount types.Account }{clientAccount: clientAccount},
		data: clientAccount,
	})
	return
}

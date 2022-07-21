package endpoint

import (
	"database/sql"
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

//RegisterAccount register client's account.
func (end *EndPoint) RegisterAccount(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	accountInfo := types.Account{
		Email:       c.PostForm("email"),
		Mnemonic:    c.PostForm("mnemonic"),
		NickName:    c.PostForm("nickname"),
		DeviceToken: c.PostForm("device_token"),
	}

	file, _, err := c.Request.FormFile("profile-image")
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}
	accountInfo.ProfileImage = data

	tx, err := end.DB.Begin()
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(c, `INSERT INTO client_account 
											(email, mnemonic, nickname, device_token, profile_image) 
											 VALUES ($1,$2,$3,$4,$5)`)
	defer stmt.Close()

	_, err = stmt.Exec(accountInfo.Email, accountInfo.Mnemonic, accountInfo.NickName, accountInfo.DeviceToken, accountInfo.ProfileImage)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	c.JSON(200, types.SetResponse(SUCCESS, 200, false, nil))
	return
}

//VerifyAccount verify validation of account.
func (end *EndPoint) VerifyAccount(c *gin.Context) {
	var clientAccount = types.Account{}

	mnemonic := c.DefaultQuery("mnemonic", "none")
	if mnemonic == "none" {
		c.JSON(400, types.SetResponse("mnemonic empty err", 400, true, nil))
		return
	}

	tx, err := end.DB.Begin()
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(c, `SELECT 
								email,
								mnemonic,
								nickname,
								device_token,
								profile_image
								FROM client_account 
								WHERE mnemonic = $1`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(mnemonic).Scan(&clientAccount.Email, &clientAccount.Mnemonic, &clientAccount.NickName, &clientAccount.DeviceToken, &clientAccount.ProfileImage)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	c.JSON(200, types.SetResponse(SUCCESS, 200, false, clientAccount))
	return
}

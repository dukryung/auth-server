package endpoint

import (
	"database/sql"
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/dukryung/microservice/server/types"
	"github.com/gin-gonic/gin"
	"io"
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
		Email:    c.PostForm("email"),
		Mnemonic: c.PostForm("mnemonic"),
		NickName: c.PostForm("nickname"),
		DeviceID: c.PostForm("device-id"),
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

	stmt, err := tx.PrepareContext(c, `INSERT INTO device_assignment (email,mnemonic,device_id)
											VALUES ($1,$2,$3)
											ON CONFLICT (device_id) DO UPDATE SET (email,mnemonic) = (excluded.email, excluded.mnemonic);`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(accountInfo.Email, accountInfo.Mnemonic, accountInfo.DeviceID)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	stmt, err = tx.PrepareContext(c, `INSERT INTO client_account
											(email, mnemonic, nickname, profile_image) 
											 VALUES ($1,$2,$3,$4)`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	_, err = stmt.Exec(accountInfo.Email, accountInfo.Mnemonic, accountInfo.NickName, accountInfo.ProfileImage)
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

//LoginAccount verify validation of account.
func (end *EndPoint) LoginAccount(c *gin.Context) {
	var clientAccount = types.Account{}

	mnemonic := c.DefaultQuery("mnemonic", "none")
	if mnemonic == "none" {
		c.JSON(400, types.SetResponse("mnemonic empty err", 400, true, nil))
		return
	}

	deviceID := c.DefaultQuery("device_id", "none")
	if deviceID == "none" {
		c.JSON(400, types.SetResponse("device id empty err", 400, true, nil))
		return
	}

	email := c.DefaultQuery("email", "none")
	if email == "none" {
		c.JSON(400, types.SetResponse("email empty err", 400, true, nil))
		return
	}

	tx, err := end.DB.Begin()
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(c, `SELECT
								email
								FROM device_assignment
								WHERE device_id = $1 AND email = $2 AND mnemonic = $3`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		fmt.Println("err :", err)
		return
	} else if err == io.EOF {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		fmt.Println("err :", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(deviceID, email, mnemonic).Scan(&email)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	stmt, err = tx.PrepareContext(c, `SELECT 
								email,
								mnemonic,
								nickname,
								profile_image
								FROM client_account 
								WHERE mnemonic = $1 AND email = $2`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	err = stmt.QueryRow(mnemonic, email).Scan(&clientAccount.Email, &clientAccount.Mnemonic, &clientAccount.NickName, &clientAccount.ProfileImage)
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

func (end *EndPoint) ImportAccount(c *gin.Context) {
	var clientAccount = types.Account{}
	if err := c.BindJSON(&clientAccount); err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	tx, err := end.DB.Begin()
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(c, `INSERT INTO device_assignment (email,mnemonic,device_id)
											VALUES ($1,$2,$3)
											ON CONFLICT (mnemonic) DO UPDATE SET device_id = excluded.device_id;`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		fmt.Println("err : ", err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(clientAccount.Email, clientAccount.Mnemonic, clientAccount.DeviceID)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}

	stmt, err = tx.PrepareContext(c, `SELECT 
								email,
								mnemonic,
								nickname,
								profile_image
								FROM client_account 
								WHERE mnemonic = $1 AND email = $2`)
	if err != nil {
		c.JSON(400, types.SetResponse(err.Error(), 400, true, nil))
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(clientAccount.Mnemonic, clientAccount.Email).Scan(&clientAccount.Email, &clientAccount.Mnemonic, &clientAccount.NickName, &clientAccount.ProfileImage)
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

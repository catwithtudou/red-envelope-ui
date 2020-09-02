package users

import (
	acountService "github.com/catwithtudou/red-envelope-account/services"
	"github.com/catwithtudou/red-envelope-infra/base"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"strconv"
)

/**
 *@Author tudou
 *@Date 2020/8/30
 **/

type UserService struct {
}

func (u *UserService) Login(mobile, username string) (user *User) {
	as := acountService.GetAccountService()
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := UserDao{runner: runner}
		user = dao.GetOne(mobile)
		//创建用户
		if user == nil {
			user = &User{
				Mobile:   mobile,
				Username: username,
			}
			id, err := dao.Insert(user)
			if err != nil {
				log.Error(err)
				return err
			}
			user.Id = id
			user.UserId = strconv.Itoa(int(id))
		}
		if user.Username != username {
			user.Username = username
			rows, err := dao.Update(user)
			if err != nil {
				log.Error(err)
				return err
			}
			if rows <= 0 {
				log.Warn("non updated:", user)
			}

		}
		return nil

	})
	if err != nil {
		log.Error(err)
		return nil
	}
	user.UserId = strconv.Itoa(int(user.Id))
	//创建资金账户
	a := as.GetEnvelopeAccountByUserId(user.UserId)
	if a == nil {
		dto := acountService.AccountCreatedDTO{
			UserId:       user.UserId,
			Username:     user.Username,
			AccountName:  user.Username,
			AccountType:  int(acountService.EnvelopeAccountType),
			CurrencyCode: acountService.DefaultCurrencyCode,
			Amount:       "1000",
		}
		_, err := as.CreateAccount(dto)
		if err != nil {
			log.Error(err)
			return
		}
	} else {
		log.Info("已经存在红包资金账户：", user)
	}
	return
}

package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"git.imooc.com/wendell1000/infra"
	"git.imooc.com/wendell1000/infra/base"
	"git.imooc.com/wendell1000/infra/httpclient"
	"git.imooc.com/wendell1000/infra/lb"
	"git.imooc.com/wendell1000/resk/services"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var _ services.AccountService = new(AccountClientService)
var once sync.Once

type AccountClientStarter struct {
	infra.BaseStarter
}

func (s *AccountClientStarter) Start(ctx infra.StarterContext) {
	once.Do(func() {
		services.IAccountService = NewAccountClientService()
	})
}

const (
	TRANSFER_ENDPOINT = "http://%s/v1/account/transfer"
	CREATE_ENDPOINT   = "http://%s/v1/account/create"
	GET_EU_ENDPOINT   = "http://%s/v1/account/envelope/get?userId=%s"
	AccountServiceId  = "account"
)

type AccountClientService struct {
	Client *httpclient.HttpClient
}

func NewAccountClientService() *AccountClientService {
	ac := &AccountClientService{}
	//创建一个apps实例
	apps := &lb.Apps{Client: base.EurekaClient()}

	ac.Client = httpclient.NewHttpClient(apps, &httpclient.Option{
		Timeout: 20 * time.Second,
	})
	return ac
}

//剩下的接口呢使用同样的方法来编写，
//这里把剩余的接口作为作业留给同学们来编写
//如果在编写过程中遇到问题，可以一起讨论
func (a *AccountClientService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {

	data, err := json.Marshal(&dto)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(data)
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	r, err := a.Client.NewRequest(http.MethodPost,
		fmt.Sprintf(CREATE_ENDPOINT, AccountServiceId),
		body,
		header,
	)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	res, err := a.Client.Do(r)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rr := &base.Res{
		Data: &services.AccountDTO{},
	}
	err = json.Unmarshal(d, rr)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if rr.Code == base.ResCodeOk {
		return rr.Data.(*services.AccountDTO), err
	} else {
		return nil, err
	}
}

func (a *AccountClientService) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {

	header := http.Header{}
	header.Add("Content-Type", "application/json")
	r, err := a.Client.NewRequest(http.MethodGet,
		fmt.Sprintf(GET_EU_ENDPOINT, AccountServiceId, userId),
		nil,
		header,
	)

	if err != nil {
		log.Error(err)
		return nil
	}
	res, err := a.Client.Do(r)
	if err != nil {
		log.Error(err)
		return nil
	}
	if res.StatusCode != http.StatusOK {
		return nil
	}
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil
	}
	rr := &base.Res{
		Data: &services.AccountDTO{},
	}
	err = json.Unmarshal(d, rr)
	if err != nil {
		log.Error(err)
		return nil
	}
	if rr.Code == base.ResCodeOk {
		return rr.Data.(*services.AccountDTO)
	} else {
		return nil
	}
}

func (a *AccountClientService) Transfer(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	data, err := json.Marshal(&dto)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	body := bytes.NewBuffer(data)
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	r, err := a.Client.NewRequest(http.MethodPost,
		fmt.Sprintf(TRANSFER_ENDPOINT, AccountServiceId),
		body,
		header,
	)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	res, err := a.Client.Do(r)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	if res.StatusCode != http.StatusOK {
		return services.TransferedStatusFailure, errors.New("状态不为200")
	}
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	rr := &base.Res{}
	err = json.Unmarshal(d, rr)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	if rr.Code == base.ResCodeOk {
		return services.TransferedStatusSuccess, err
	} else {
		return services.TransferedStatusFailure, err
	}
}

// Package encrypt
// @title
// @description
// @author njy
// @since 2023/6/2 09:13
package encrypt

import (
	"fmt"
	"sync"
)

var httpData *MemoryProviderServer
var once1 sync.Once

type LetsEncrypt struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email        string `gorm:"column:email;type:varchar(191);uniqueIndex:idx_email;not null;default:'';comment:email" json:"email"`
	PrivateKey   string `gorm:"column:private_key;type:text;comment:账号key" json:"privateKey"`
	Registration string `gorm:"column:registration;type:text;comment:lets Encrypt注册信息" json:"registration"`
	CreateAt     int64  `gorm:"column:create_at;not null;autoCreateTime" json:"createAt"`
	UpdateAt     int64  `gorm:"column:update_at;not null;autoUpdateTime" json:"updateAt"`
}

type MemoryProviderServer struct {
	data map[string]string
	lock sync.Mutex
}

func NewMemoryProviderServer() *MemoryProviderServer {
	once1.Do(func() {
		httpData = &MemoryProviderServer{
			data: map[string]string{},
		}
	})
	return httpData
}

func (s *MemoryProviderServer) Present(domain, token, keyAuth string) error {
	s.lock.Lock()
	fmt.Printf("保存token:%v,value:%v\n", token, keyAuth)
	s.data[token] = keyAuth
	s.lock.Unlock()
	return nil
}

func (s *MemoryProviderServer) CleanUp(domain, token, keyAuth string) error {
	s.lock.Lock()
	delete(s.data, domain+token)
	fmt.Printf("清空token:%v\n", token)
	s.lock.Unlock()
	return nil
}

func (s *MemoryProviderServer) GetKeyAuth(domain, token string) (exist bool, keyAuth string) {
	s.lock.Lock()
	fmt.Printf("domain data:%+v", s.data)
	keyAuth, exist = s.data[token]
	s.lock.Unlock()
	return
}

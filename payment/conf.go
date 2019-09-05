package payment

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/wuwenbao/wechat/util"
)

type Confer interface {
	Appid() string
	MchId() string
	MchKey() string
	GetTLSConfig() (*tls.Config, error)
	SignCheck(signType interface{}) string
}

type conf struct {
	appid    string
	mchId    string
	mchKey   string
	CertPath string
	KeyPath  string
}

func Conf(appid, mchId, mchKey string) *conf {
	return &conf{
		appid:    appid,
		mchId:    mchId,
		mchKey:   mchKey,
		CertPath: "",
		KeyPath:  "",
	}
}

func (p *conf) Appid() string {
	return p.appid
}

func (p *conf) MchId() string {
	return p.mchId
}

func (p *conf) MchKey() string {
	return p.mchKey
}

func (p *conf) GetTLSConfig() (*tls.Config, error) {
	return p.setTLSConfig()(p.CertPath, p.KeyPath)
}

func (p *conf) setTLSConfig() func(certPath, keyPath string) (*tls.Config, error) {
	conf := new(tls.Config)
	once := new(sync.Once)
	return func(certPath, keyPath string) (config *tls.Config, e error) {
		once.Do(func() {
			cert, err := tls.LoadX509KeyPair(certPath, keyPath)
			if err != nil {
				log.Println("cert load fail:", err)
				return
			}
			conf = &tls.Config{
				Certificates: []tls.Certificate{cert},
			}
		})
		return conf, nil
	}
}

type CDATA struct {
	Text string `xml:",cdata"`
}

//SignCheck 数据有效检查
func (p *conf) SignCheck(signType interface{}) string {
	valueOf := reflect.ValueOf(signType)
	typeOf := reflect.TypeOf(signType)

	switch typeOf.Kind() {
	case reflect.Ptr:
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return ""
	}

	storeKv := make(map[string]string)
	var sortK []string
	for i := 0; i < typeOf.NumField(); i++ {
		tag, ok := typeOf.Field(i).Tag.Lookup("xml")
		if !ok || tag == "" {
			continue
		}
		tags := strings.Split(tag, ",")
		if len(tags) < 0 {
			continue
		}
		s := valueOf.Field(i).String()
		if s == "" || s == "0" {
			continue
		}
		key := strings.TrimSpace(tags[0])

		if key == "sign" {
			continue
		}

		switch valueOf.Field(i).Interface().(type) {
		case CDATA:
			storeKv[key] = valueOf.Field(i).Interface().(CDATA).Text
		default:
			storeKv[key] = fmt.Sprintf("%v", valueOf.Field(i).Interface())
		}
		sortK = append(sortK, key)
	}

	sort.Strings(sortK)

	var buf bytes.Buffer
	for index, val := range sortK {
		if index != 0 {
			buf.WriteString("&")
		}
		buf.WriteString(val + "=" + storeKv[val])
	}
	buf.WriteString("&key=")
	buf.WriteString(p.MchKey())
	hs := util.Md5(buf.Bytes())
	return strings.ToUpper(hs)
}

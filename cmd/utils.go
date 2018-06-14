package cmd

import (
	"io/ioutil"
	"os"
	"bufio"
	"io"
	"strings"
)

func ModifyHosts(filePath, newIp, domain string) error {
	//调用该函数需要用户对/etc/hosts 可写， sudo chown ubuntu:ubuntu /etc/hosts
	Spacer := "       "//7 space
	newHosts, err := ioutil.ReadFile(filePath)//读取原有hosts 保存
	if err != nil {
		return err
	}
	fi, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	allMap := make(map[string]string)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		qq := strings.Split(string(a), Spacer)
		if len(qq) >= 2 {
			allMap[qq[1]] = qq[0]
		}
	}
	//如果域名存在
	tempNew := newIp+ Spacer + domain
	if ip, ok := allMap[domain]; ok {//如果存在之前的域名
		tempOld := ip + Spacer + domain
		newHosts = []byte(strings.Replace(string(newHosts),tempOld,tempNew,1))
	}else {//如果不存在之前的域名，则直接追加
		newHosts = []byte(string(newHosts) + "\n" + tempNew)
	}
	if err := ioutil.WriteFile(filePath,newHosts,0644); err != nil {
		return err
	}
	return nil
}
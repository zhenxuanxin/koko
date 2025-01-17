package handler

import (
	"github.com/jumpserver/koko/pkg/logger"
	"github.com/jumpserver/koko/pkg/proxy"
)

func (d *DirectHandler) LoginConnectToken() {
	connectToken := d.opts.tokenInfo
	user := connectToken.User
	i18nLang := d.i18nLang
	protocol := connectToken.Protocol
	proxyOpts := make([]proxy.ConnectionOption, 0, 10)
	proxyOpts = append(proxyOpts, proxy.ConnectProtocol(protocol))
	proxyOpts = append(proxyOpts, proxy.ConnectUser(&user))
	proxyOpts = append(proxyOpts, proxy.ConnectAsset(&connectToken.Asset))
	proxyOpts = append(proxyOpts, proxy.ConnectAccount(&connectToken.Account))
	proxyOpts = append(proxyOpts, proxy.ConnectActions(connectToken.Actions))
	proxyOpts = append(proxyOpts, proxy.ConnectExpired(connectToken.ExpireAt))
	proxyOpts = append(proxyOpts, proxy.ConnectDomain(connectToken.Domain))
	proxyOpts = append(proxyOpts, proxy.ConnectPlatform(&connectToken.Platform))
	proxyOpts = append(proxyOpts, proxy.ConnectGateway(connectToken.Gateway))
	proxyOpts = append(proxyOpts, proxy.ConnectCmdACLRules(connectToken.CommandFilterACLs))
	proxyOpts = append(proxyOpts, proxy.ConnectI18nLang(i18nLang))
	srv, err := proxy.NewServer(d.wrapperSess, d.jmsService, proxyOpts...)
	if err != nil {
		logger.Error(err)
		return
	}
	srv.Proxy()
	logger.Infof("Request %s: token %s proxy end", d.wrapperSess.Uuid, connectToken.Id)

}

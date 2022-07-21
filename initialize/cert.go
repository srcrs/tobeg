package initialize

import (
	"tobeg/cert"
	"tobeg/global"
)

func InitCert() {
	cert.Appid = global.Config.AlipayConfig.AppId

	cert.AlipayPublicContentRSA2 = []byte(`-----BEGIN CERTIFICATE-----` +
		global.Config.AlipayConfig.PublicContentRSA2 +
		`-----END CERTIFICATE-----`)

	cert.PrivateKey = global.Config.AlipayConfig.PrivateKey

	cert.AppPublicContent = []byte(`-----BEGIN CERTIFICATE-----` +
		global.Config.AlipayConfig.AppPublicContent +
		`-----END CERTIFICATE-----`)
}

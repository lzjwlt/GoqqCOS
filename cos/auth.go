package cos

import (
	"crypto/hmac"
	"crypto/sha1"
	"strings"
)

type CosConfig struct {
	AppID     string
	sercetID  string
	secretKey string
	Region    string
}

func test() {
	config1 := CosConfig{"1253632585", "AKIDud8OPB402oJYXFqqnV1IK7ewzmIYFKuF", "EEzVlmY4LNAi15uUotjKgmdFclfKkuoR", "ap-shanghai"}
	_ = config1
}

type signParas struct {
	signFields
	qSignature string
}

type signFields struct {
	qSignAlgorithm string
	qAk            string
	qSignTime      string
	qKeyTime       string
	qHeaderList    string
	qURLParamList  string
}

type httpPara struct {
	method     string
	uri        string
	parameters string
	headers    string
}

func CalcSignature(sf signFields, conf CosConfig, hp httpPara) []byte {
	mac := hmac.New(sha1.New, []byte(conf.secretKey))
	mac.Write([]byte(sf.qAk))
	signKey := mac.Sum(nil)
	httpString := strings.ToLower(hp.method) + "\n" +
		hp.uri + "\n" + hp.parameters + "\n" + hp.headers + "\n"

	hs := sha1.New()
	hs.Write([]byte(httpString))
	httpStringHash := hs.Sum(nil)
	stringToSign := sf.qSignAlgorithm + "\n" + sf.qSignTime + "\n" +
		httpStringHash + "\n"
	mac = hmac.New(sha1.New, signKey)
	mac.Write([]byte(stringToSign))
	return mac.Sum(nil)
}

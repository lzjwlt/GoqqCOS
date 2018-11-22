package cos

import (
	"net/http"
	"strings"
)

func CreateBucket(bucket string, config CosConfig) {
	req, _ := http.NewRequest("PUT", "/", strings.NewReader(bucket))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}

func DeleteBucket() {

}

func HeadBucket() {

}

func ListObjects() {

}

func PutACL() {

}

func GetACL() {

}

func PutCORS() {

}

func GetCORS() {

}

func DelCORS() {

}

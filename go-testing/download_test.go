package go_testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content")
	{
		t.Logf("\tWhen checking \"%s\" for statuc code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}

			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Logf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

func ExampleRequest() {

	url := "https://oapi.dingtalk.com/"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	type DingData struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	var data DingData
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(data)
	// Output:
	// {404 请求的URI地址不存在}
}

// $ go test -v
// $ go test -v -run="ExampleRequest"

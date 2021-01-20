package jumia

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"palmbuy/boot/utils"
	"strings"
)

var (
	jumiaUrl = "https://www.jumia.com.ng/"
	// 中间件管理服务
	JumiaService = new(serviceJumia)
)

type serviceJumia struct{}

func (j *serviceJumia) SaveProductsToFileByCategoryName(categoryName string) error {
	client := &http.Client{}
	for i := 1; i <= 50; i++ {
		productListUrl := jumiaUrl + categoryName + fmt.Sprintf("/?page=%d&sort=newest", i)

		//提交请求
		request, err := http.NewRequest("GET", productListUrl, nil)
		if err != nil {
			glog.Error("请求获取jumia商品失败",err)
			continue
		}

		//增加header选项
		request.Header.Add("accept", "application/json")
		request.Header.Add("user-agent", " Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36")
		request.Header.Add("client-basic", `{"experiment":"{\"AH31PjfCT6S-qENCa5K6Yw\":1,\"wOcRXvdGThCgDXbpz3-sEw\":-1}","country_code":"gb","language_code":"en-US","gender":"F","v":"android_6.5.0","from_site":"wholee","device_id":"be23432373b4c6e"}`)
		resp, err := client.Do(request)
		if err != nil {
			glog.Error("请求获取jumia商品失败",err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			glog.Error("解析jumia商品信息失败",err)
			continue
		}
		originViewData := gjson.Get(string(body), "viewData").String()
		// 存入txt文件
		err = utils.BufferWrite("public/resource/txts/jumia/"+categoryName+".txt", originViewData)
		resp.Body.Close()
		if err != nil {
			glog.Error("jumia商品信息保存文件失败",err)
			continue
		}
	}
	return nil
}

func (j *serviceJumia) GetFileList(page, limit int) (int, []map[string]interface{}) {
	files,_ := gfile.ScanDir("public/resource/txts/jumia","*",true)
	start := (page - 1) * limit
	end := start + limit
	if end > len(files) {
		end = len(files)
	}
	items := make([]map[string]interface{}, 0)
	for i := start; i < end; i++ {
		filePath := strings.Split(files[i], string(os.PathSeparator))
		categoryName := strings.TrimRight(filePath[len(filePath)-1],".txt")
		item := map[string]interface{}{"categoryName": categoryName}
		items = append(items, item)
	}
	return len(files), items
}



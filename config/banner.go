package config

import "fmt"

func Banner() {
	banner := `
		|￣￣￣￣￣￣￣￣￣￣￣￣| 	用法：
		 |                      |
		 |  我要成为大黑客...   |	-t <域名>  eg：druid -t example.com
		 |                      |	-t <ip:端口>  eg：druid -t 101.25.**.62:8888
		|＿＿＿＿＿＿＿miaomiao~|
	       ||
	(\__/) ||
	(•ㅅ•) ||
	/ 　 づv
	`
	banner2 := ` |\/\/\/|  
	 |      |  
	 |      |  
	 | (o)(o)  
	 C      _) 
	  | ,___|  
	  |   /    
	 /____\    
	/      \   `
	var bannerSlice []string
	bannerSlice = append(bannerSlice, banner)
	bannerSlice = append(bannerSlice, banner2)

	fmt.Println(banner)
}

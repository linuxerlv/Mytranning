func subdomainVisits(cpdomains []string) []string {
	count := make(map[string]int, 0)  // 创建一个字符映射，大小为0
	for _, value := range cpdomains {
		splits := strings.Split(value, " ") // 将序列里的字符串分组
		num, _ := strconv.Atoi(splits[0]) //获取引用次数
		for {   //迭代查找
			count[splits[1]] += num
			//找当前域名第一个点的位置
			dotIndex := strings.Index(splits[1], ".")
			//找不到点说明这一条记录完顶级域名了
			if dotIndex< 0 {
				break
			}
			//裁剪点前面的部分得到更高一级域名
			splits[1] = splits[1][dotIndex+1:]
		}
	}

	ans := make([]string, 0, len(count))
	for key, value := range count {
		ans = append(ans, fmt.Sprintf("%d %s", value, key))
	}
	return ans
}
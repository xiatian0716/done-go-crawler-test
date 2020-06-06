package parse

import (
	"go-crawler-test/engine"
	"go-crawler-test/model"
	"regexp"
	"strconv"
)

// [\d\D]	匹配任何字符(含行符)
// *?<		贪婪匹配(匹配到最前一个<)
// *<		贪婪匹配(匹配到最后一个<)
var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var intoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte) engine.ParseResult {
	//fmt.Printf("%s",contents)
	bookdetail := model.Bookdetail{}
	bookdetail.Author = ExtraString(contents, autoRe)
	page, err := strconv.Atoi(ExtraString(contents, pageRe))
	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Publicer = ExtraString(contents, public)
	bookdetail.Into = ExtraString(contents, intoRe)
	bookdetail.Score = ExtraString(contents, scoreRe)
	bookdetail.Price = ExtraString(contents, priceRe)
	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}

	return result
}

func ExtraString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}

// <div id="info" class="">

//     <span>
//       <span class="pl"> 作者</span>:

//             <a class="" href="/search/%E5%BC%97%E6%9C%97%E7%B4%A2%E7%93%A6%E2%80%A2%E8%82%96%E8%8E%B1">[美] 弗朗索瓦•肖莱</a>
//     </span><br>

//     <span class="pl">出版社:</span> 人民邮电出版社<br>

//     <span class="pl">出品方:</span>&nbsp;<a href="https://book.douban.com/series/47356?brand=1">图灵教育</a><br>

//     <span class="pl">原作名:</span> Deep Learning with Python<br>

//     <span>
//       <span class="pl"> 译者</span>:

//             <a class="" href="/search/%E5%BC%A0%E4%BA%AE">张亮</a>
//     </span><br>

//     <span class="pl">出版年:</span> 2018-8<br>

//     <span class="pl">页数:</span> 292<br>

//     <span class="pl">定价:</span> 119.00元<br>

//     <span class="pl">装帧:</span> 平装<br>

//     <span class="pl">丛书:</span>&nbsp;<a href="https://book.douban.com/series/660">图灵程序设计丛书</a><br>

//       <span class="pl">ISBN:</span> 9787115488763<br>

// </div>

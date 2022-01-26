package es

import (
	"context"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"reflect"
)

type Article struct {
	Title   string `json:"title"`   //文章标题
	Content string `json:"content"` //文章标题
}

func InitEs() {
	client, err := NewEsClient()
	if err != nil {
		fmt.Println("es 连接成功")
	} else {
		fmt.Println("es 连接失败")
	}

	//指定es请求时需要一个上下文
	ctx := context.Background()

	//创建term查询条件,用于精确查询
	termQuery := elastic.NewMatchQuery("title", "0916文章1")

	searchResult, err := client.Search().
		Index("voc_mini_program"). // 设置索引名
		Query(termQuery).          //设置查询条件
		//Sort("created_time", false). //设置排序字段,根据created_time字段降序排列,第二个字段false表示逆序
		From(0).      //设置分页偏移量 - 起始偏移量,从0行记录开始
		Size(10).     //设置分页参数 - 每页大小
		Pretty(true). //查询结果返回可读性较好的JSON格式
		Do(ctx)       //执行请求

	// TODO 后续其他es查询条件测试: https://www.tizi365.com/archives/858.html

	if err != nil {
		panic(err)
	}

	fmt.Printf("查询总耗时 %d ms,查询结果总条数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		// 查询结果不为0,将结果打印出来
		var article Article
		// 通过Each方法,将es结果的json结构转换为struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(article)) {
			// 转换成Article对象
			if t, ok := item.(Article); ok {
				fmt.Println(t.Title)
				fmt.Println(t.Content)
			}
		}
	}

}

//NewEsClient 初始化连接es的客户端
func NewEsClient() (*elastic.Client, error) {
	url := fmt.Sprintf("http://%s:%d", setting.ElasticsearchConfig.Host, setting.ElasticsearchConfig.Port)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		// 允许指定弹性是否应该定期检查集群
		elastic.SetSniff(false))
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}
	return client, err
}

//
//type SearchRequest struct {
//}
//
//// 查询
//func (r *SearchRequest) ToFilter() *EsSearch {
//	var search EsSearch
//	if len(r.Nickname) != 0 {
//		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("nickname", r.Nickname))
//	}
//	if len(r.Phone) != 0 {
//		search.ShouldQuery = append(search.ShouldQuery, elastic.NewTermsQuery("phone", r.Phone))
//	}
//	if len(r.Ancestral) != 0 {
//		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("ancestral", r.Ancestral))
//	}
//	if len(r.Identity) != 0 {
//		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("identity", r.Identity))
//	}
//
//	if search.Sorters == nil {
//		search.Sorters = append(search.Sorters, elastic.NewFieldSort("create_time").Desc())
//	}
//
//	search.From = (r.Num - 1) * r.Size
//	search.Size = r.Size
//	return &search
//}
//
//func (es *UserES) Search(ctx context.Context, filter *model.EsSearch) ([]*model.UserEs, error) {
//	boolQuery := elastic.NewBoolQuery()
//	boolQuery.Must(filter.MustQuery...)
//	boolQuery.MustNot(filter.MustNotQuery...)
//	boolQuery.Should(filter.ShouldQuery...)
//	boolQuery.Filter(filter.Filters...)
//
//	// 当should不为空时，保证至少匹配should中的一项
//	if len(filter.MustQuery) == 0 && len(filter.MustNotQuery) == 0 && len(filter.ShouldQuery) > 0 {
//		boolQuery.MinimumShouldMatch("1")
//	}
//
//	service := es.client.Search().Index(es.index).Query(boolQuery).SortBy(filter.Sorters...).From(filter.From).Size(filter.Size)
//	resp, err := service.Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.TotalHits() == 0 {
//		return nil, nil
//	}
//	userES := make([]*model.UserEs, 0)
//	for _, e := range resp.Each(reflect.TypeOf(&model.UserEs{})) {
//		us := e.(*model.UserEs)
//		userES = append(userES, us)
//	}
//	return userES, nil
//}

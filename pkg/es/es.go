package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

const indexName = "voc_yuwell"
const esIndexUrlPrefix = "http://1.117.229.170:9200"

// YuWell 定义yuwell的结构体
type YuWell struct {
	Category string `json:"category"` //品类
	Brand    string `json:"brand"`    //品牌
}

func main() {

	// 创建ES client用于后续操作ES
	client, err := elastic.NewClient(
		// 设置ES服务地址，支持多个地址
		elastic.SetSniff(false), // TODO 特别注意：该参数用于设置集群；单个节点的es时需要设置该值为false
		elastic.SetURL(esIndexUrlPrefix),
		// 设置基于http base auth验证的账号和密码
		elastic.SetBasicAuth("user", "secret")) // 建议创建es时设置密码
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	//执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 首先检测下索引是否存在
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		panic(fmt.Errorf("index %s 不存在", indexName))
	}

	// 创建成功后插入一条数据
	//msg1 := YuWell{Category: "制氧机", Brand: "鱼跃"}
	//msg1 := YuWell{Category: "血压计", Brand: "欧姆龙"}
	//put1, err := client.Index().Index(indexName).Id("1").BodyJson(msg1).Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("文档id: %s，索引名: %s\n", put1.Id, put1.Index)
	//
	//// 插入成功后查询一个数据（根据指定id查询数据）
	//get1, err := client.Get().Index(indexName).Id("1").Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if get1.Found {
	// 手动将文档转换为 go struct 对象
	//msg2 := YuWell{}
	//data,_ := get1.Source.MarshalJSON()
	//json.Unmarshal(data,&msg2)
	//fmt.Println(msg2.Category)
	//}

	// es DSL search
	// 创建term查询条件，用于精确查询
	termQuery := elastic.NewTermQuery("category", "制氧机")
	//
	searchResult, err := client.Search().
		Index(indexName).          // 设置索引名
		Query(termQuery).          // 设置查询条件
		Sort("origin_time", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0).                   // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).                  // 设置分页参数 - 每页大小
		Pretty(true).              // 查询结果返回可读性较好的JSON格式
		Do(ctx)                    // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 YuWell
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(YuWell); ok {
				fmt.Println(t.Category)
			}
		}
	}

	// 修改一条数据
	// 查询一条数据
	// 删除一条数据
	// 查询数据
}

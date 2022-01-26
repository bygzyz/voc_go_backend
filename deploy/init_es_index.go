/*
初始化es的index
*/

/*
TODO 特别注意（22-01-26）：
1、在初始化es的index时最好直接使用 BodyString，不建议 BodyJson ，因为这样index复杂的字段定义就要大量使用 map[string]interface{}
2、BodyString 中定义的index必须符合json格式，类似postman传递json数据时，不能有多余的,否则es client 会一直报错：` panic: elastic: Error 400 (Bad Request): Failed to parse content to map [type=parse_exception]`
*/

package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

/*
data_id           唯一id
category          品类（血压计、制氧机）
brand             品牌（鱼跃、欧姆龙、松下、美菱、迈克大夫、久安、JZIKI）
model             型号（YE666AR、YE680CR、YE660D、YE670A、YE670CR）
data_type         数据(渠道)类型（电商评论、电商社区、资讯、内部数据、视频、论坛、社交媒体）
platform          平台（天猫、京东、苏宁）
goods_name        商品名称（注意：与title不同；这里存储当前信息对应的商品）
keyword_goods_name 存储内容和goods_name相同,该字段主要用户用户反馈分析中标签排行下商品数据量聚合
message_source    信息来源（电商评论、京东评论、拼多多评论、苏宁评论、微博、鱼跃公众号、知乎）
agg_class_level   嵌套标签
class_level1      一级标签 【一二三级标签均需要算法抽取后存入es】  如果存数组的话，需要调查根据某个标签的值查询记录，每个标签对应的情感怎么存。这块设计要和王子讨论。
class_level2      二级标签
class_level3      三级标签 [{label_name:xxx,sentiment:正面},{label_name:xxx,sentiment:正面},{label_name:xxx,sentiment:正面}]
data_validity     数据有效性
sentiment         情感
title             标题
content           内容
keyword           关键词[可能有多个]
agg_keyword       嵌套关键词
is_delete         是否被删除标记
href              原文链接
created_time      创建时间
origin_time       原生生成时间
shop_id           店铺id
*/

// title 和 content 使用ik分词器的 ik_max_word,尽可能多地分词

const yuYellMapping = `
{
    "mappings": {
        "properties": {
            "data_id": {
                "type": "text",
                "analyzer": "keyword"
            },
            "category": {
                "type": "keyword"
            },
            "brand": {
                "type": "keyword"
            },
            "model": {
                "type": "text"
            },
            "data_type": {
                "type": "keyword"
            },
            "platform": {
                "type": "keyword"
            },
            "goods_name": {
                "type": "text",
                "analyzer": "ik_smart",
                "search_analyzer": "ik_max_word"
            }, 
            "keyword_goods_name": {
                "type": "keyword"
            }, 
            "message_source": {
                "type": "keyword"
            },
            "agg_class_level": {
                "type": "nested",
                "properties": {
                    "level1": {"type": "keyword"},
                    "level2": {"type": "keyword"},
                    "level3": {"type": "keyword"},
                    "sentiment": {"type": "keyword"},
                    "percentage": {"type": "keyword"},
                    "origin_time": {"type": "date", "format": "yyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyy-M-d||epoch_millis"}
                }
            },  
            "class_level1": {
                "type": "text",
                "analyzer": "keyword"
            }, 
            "class_level2": {
                "type": "text",
                "analyzer": "keyword"
            },
            "class_level3": {
                "type": "text",
                "analyzer": "keyword"
            },
            "scene_level1": {
                "type": "text",
                "analyzer": "keyword"
            },
            "scene_level2": {
                "type": "text",
                "analyzer": "keyword"
            }, 
            "agg_scene_label": {
                "type": "nested",
                "properties": {
                    "level1": {"type": "keyword"},
                    "level2": {"type": "keyword"},
                    "origin_time": {"type": "date", "format": "yyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyy-M-d||epoch_millis"}
                }
            }, 
            "data_validity": {
                "type": "boolean"
            },
            "sentiment": {
                "type": "keyword"
            },
            "title": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "content": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "keyword": {
                "type": "text",
                "analyzer": "keyword"
            }, 
            "agg_keyword": {
                "type": "nested",
                "properties": {
                    "each_keyword": {"type": "keyword"},
                    "origin_time": {"type": "date", "format": "yyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyy-M-d||epoch_millis"}
                }
            }, 
            "is_delete": {
                "type": "boolean"
            },
            "href": {
                "type": "text"
            }, 
            "created_time": {
                "type": "date",
                "format": "yyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyy-M-d||epoch_millis"
            },
            "origin_time": {
                "type": "date",
                "format": "yyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyy-M-d||epoch_millis"
            },
            "star_count": {
                "type": "text",
                "analyzer": "keyword"  
            },
            "comment_count": {
                "type": "text",
                "analyzer": "keyword" 
            },
            "spread_count": {
                "type": "text",
                "analyzer": "keyword" 
            },
            "view_count": {
                "type": "text",
                "analyzer": "keyword"  
            },
            "shop_id": {
                "type": "text",
                "analyzer": "keyword" 
            },
            "closed_at": {
                "type": "text",
                "analyzer": "keyword" 
            },
            "user_name": {
                "type": "text",
                "analyzer": "keyword"
            },
            "agent_name": {
                "type": "text",
                "analyzer": "keyword" 
            },
            "agent_content": {
                "type": "text",
                "analyzer": "keyword"  
            },
            "region": {
                "type": "text",
                "analyzer": "keyword" 
            },
            "conv_list": {
                "type": "nested",
                "properties": {
                    "sender_name": {"type": "keyword"},
                    "sender": {"type": "keyword"},
                    "created_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyy-M-d||epoch_millis"},
                    "content": {"type": "keyword"}
                }
            },
            "other_info": {
                "type": "text",
                "analyzer": "standard" 
            }
        }
    }
}
`
const indexName = "voc_yuwell_go"
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

	// voc_yuwell 索引不存在，则创建一个
	if !exists {
		_, err := client.CreateIndex(indexName).BodyString(yuYellMapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println("创建es的index成功，index名称: %s", indexName)
	}

	// 创建成功后插入一条数据
	msg1 := YuWell{Category: "制氧机", Brand: "鱼跃"}
	//msg1 := YuWell{Category: "血压计", Brand: "欧姆龙"}
	put1, err := client.Index().Index(indexName).Id("1").BodyJson(msg1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("文档id: %s，索引名: %s\n", put1.Id, put1.Index)
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
	//termQuery := elastic.NewTermQuery("category", "制氧机")
	//
	//searchResult, err := client.Search().
	//	Index(indexName).   // 设置索引名
	//	Query(termQuery).   // 设置查询条件
	//	Sort("origin_time", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
	//	From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
	//	Size(10).   // 设置分页参数 - 每页大小
	//	Pretty(true).       // 查询结果返回可读性较好的JSON格式
	//	Do(ctx)             // 执行请求
	//
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//
	//fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	//
	//if searchResult.TotalHits() > 0 {
	//	// 查询结果不为空，则遍历结果
	//	var b1 YuWell
	//	// 通过Each方法，将es结果的json结构转换成struct对象
	//	for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
	//		// 转换成Article对象
	//		if t, ok := item.(YuWell); ok {
	//			fmt.Println(t.Category)
	//		}
	//	}
	//}

	// 修改一条数据
	// 查询一条数据
	// 删除一条数据
	// 查询数据
}

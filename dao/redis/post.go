package redis

import (
	"blog/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 2.确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1
	// ZRVRANGE查询 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 2．去redis查询id列表
	// 1. 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// // 2.确定查询的索引起始点
	// start := (p.Page - 1) * p.Size
	// end := start + p.Size - 1
	// // ZRVRANGE查询 按分数从大到小的顺序查询指定数量的元素
	// return client.ZRevRange(key, start, end).Result()
	return getIDsFromKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))
	// for _, id := range ids {
	// 	key := getRedisKey(KeyPostVotedZSetPF + id)
	// 	// 查找key中分数是1的元素的数量-》统计每篇帖子的赞成票的数量
	// 	v := client.ZCount(key, "1", "1").Val()
	// 	data = append(data, v)
	// }

	// 使用pipeline一次发送多条命令，减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	//使用 zintecstoce,把分区的帖子set与帖子分数的 Zse生成一个新的?set
	//针对新的set按之前的逻辑取数据

	// 社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	//利用缓存key减少zinterstore执行的次数

	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(orderKey).Val() < 1 {
		// 不存在，需要计算
		pipeline := client.Pipeline()
		client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}

	return getIDsFromKey(orderKey, p.Page, p.Size)
}

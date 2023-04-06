package redis

const (
	KeyPrefix          = "blog:"
	KeyPostTimeZset    = "post:time"
	KeyPostScoreZset   = "post:score"
	KeyPostVotedZsetPF = "post:voted:"
	KeyCommunitySetPF  = "community:" // set;保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}

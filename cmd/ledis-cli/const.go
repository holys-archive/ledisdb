//This file was generated by ./generate.py on Fri Aug 15 2014 16:40:03 +0800 
package main

var helpCommands = [][]string{
	{"BCOUNT", "key [start end]", "Bitmap"},
	{"BDELETE", "key", "ZSet"},
	{"BEXPIRE", "key seconds", "Bitmap"},
	{"BEXPIREAT", "key timestamp", "Bitmap"},
	{"BGET", "key", "Bitmap"},
	{"BGETBIT", "key offset", "Bitmap"},
	{"BMSETBIT", "key offset value [offset value ...]", "Bitmap"},
	{"BOPT", "operation destkey key [key ...]", "Bitmap"},
	{"BPERSIST", "key", "Bitmap"},
	{"BSETBIT", "key offset value", "Bitmap"},
	{"BTTL", "key", "Bitmap"},
	{"DECR", "key", "KV"},
	{"DECRBY", "key decrement", "KV"},
	{"DEL", "key [key ...]", "KV"},
	{"ECHO", "message", "Server"},
	{"EXISTS", "key", "KV"},
	{"EXPIRE", "key seconds", "KV"},
	{"EXPIREAT", "key timestamp", "KV"},
	{"FULLSYNC", "-", "Replication"},
	{"GET", "key", "KV"},
	{"GETSET", " key value", "KV"},
	{"HCLEAR", "key", "Hash"},
	{"HDEL", "key field [field ...]", "Hash"},
	{"HEXISTS", "key field", "Hash"},
	{"HEXPIRE", "key seconds", "Hash"},
	{"HEXPIREAT", "key timestamp", "Hash"},
	{"HGET", "key field", "Hash"},
	{"HGETALL", "key", "Hash"},
	{"HINCRBY", "key field increment", "Hash"},
	{"HKEYS", "key", "Hash"},
	{"HLEN", "key", "Hash"},
	{"HMCLEAR", "key [key ...]", "Hash"},
	{"HMGET", "key field [field ...]", "Hash"},
	{"HMSET", "key field value [field value ...]", "Hash"},
	{"HPERSIST", "key", "Hash"},
	{"HSET", "key field value", "Hash"},
	{"HTTL", "key", "Hash"},
	{"HVALS", "key", "Hash"},
	{"INCR", "key", "KV"},
	{"INCRBY", "key increment", "KV"},
	{"LCLEAR", "key", "List"},
	{"LEXPIRE", "key seconds", "List"},
	{"LEXPIREAT", "key timestamp", "List"},
	{"LINDEX", "key index", "List"},
	{"LLEN", "key", "List"},
	{"LMCLEAR", "key [key ...]", "List"},
	{"LPERSIST", "key", "List"},
	{"LPOP", "key", "List"},
	{"LPUSH", "key value [value ...]", "List"},
	{"LRANGE", "key start stop", "List"},
	{"LTTL", "key", "List"},
	{"MGET", "key [key ...]", "KV"},
	{"MSET", "key value [key value ...]", "KV"},
	{"PERSIST", "key", "KV"},
	{"PING", "-", "Server"},
	{"RPOP", "key", "List"},
	{"RPUSH", "key value [value ...]", "List"},
	{"SADD", "key member [member ...]", "Set"},
	{"SCARD", "key", "Set"},
	{"SCLEAR", "key", "Set"},
	{"SDIFF", "key [key ...]", "Set"},
	{"SDIFFSTORE", "destination key [key ...]", "Set"},
	{"SELECT", "index", "Server"},
	{"SET", "key value", "KV"},
	{"SETNX", "key value", "KV"},
	{"SEXPIRE", "key seconds", "Set"},
	{"SEXPIREAT", "key timestamp", "Set"},
	{"SINTER", "key [key ...]", "Set"},
	{"SINTERSTORE", "destination key [key ...]", "Set"},
	{"SISMEMBER", "key member", "Set"},
	{"SLAVEOF", "host port", "Replication"},
	{"SMCLEAR", "key [key ...]", "Set"},
	{"SMEMBERS", "key", "Set"},
	{"SPERSIST", "key", "Set"},
	{"SREM", "key member [member ...]", "Set"},
	{"STTL", "key", "Set"},
	{"SUNION", "key [key ...]", "Set"},
	{"SUNIONSTORE", "destination key [key ...]", "Set"},
	{"SYNC", "index offset", "Replication"},
	{"TTL", "key", "KV"},
	{"ZADD", "key score member [score member ...]", "ZSet"},
	{"ZCARD", "key", "ZSet"},
	{"ZCLEAR", "key", "ZSet"},
	{"ZCOUNT", "key min max", "ZSet"},
	{"ZEXPIRE", "key seconds", "ZSet"},
	{"ZEXPIREAT", "key timestamp", "ZSet"},
	{"ZINCRBY", "key increment member", "ZSet"},
	{"ZMCLEAR", "key [key ...]", "ZSet"},
	{"ZPERSIST", "key", "ZSet"},
	{"ZRANGE", "key start stop [WITHSCORES]", "ZSet"},
	{"ZRANGEBYSCORE", "key min max [WITHSCORES] [LIMIT offset count]", "ZSet"},
	{"ZRANK", "key member", "ZSet"},
	{"ZREM", "key member [member ...]", "ZSet"},
	{"ZREMRANGEBYRANK", "key start stop", "ZSet"},
	{"ZREMRANGEBYSCORE", "key min max", "ZSet"},
	{"ZREVRANGE", "key start stop [WITHSCORES]", "ZSet"},
	{"ZREVRANGEBYSCORE", "key max min  [WITHSCORES][LIMIT offset count]", "ZSet"},
	{"ZREVRANK", "key member", "ZSet"},
	{"ZSCORE", "key member", "ZSet"},
	{"ZTTL", "key", "ZSet"},
}

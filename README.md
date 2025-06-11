# 基于go-zero的影评系统微信小程序
## 架构设计
[系统接口设计](explain.md)
[论文](movieReviewSystemApi/基于go-zero的影评系统微信小程序.md)
1. 评论服务：Redis Sorted Set管理评论热度，空值解决缓存穿透，MongoDB分表存储动态评论内容。
2. 点赞服务
3. 聊天服务：基于Websocket实现私聊和群聊功能，通过其心跳机制结合 Redis 维护用户在线状态，通过确认机制实现已读未
   读功能，通过 Kafka 异步队列实现消息的削峰填谷，MongoDB 分片存储用户聊天记录。
4. 用户服务：JWT双Token实现无状态认证，MySQL存储用户数据和用户关系信息，Redis缓存热门贴用户信息。
5. 群组服务：基于 Snowflake 算法生成分布式群ID，MySQL事务锁实现成员变更原子操作，MongoDB存储聊天历史。
6. 审核服务：DFA算法实时敏感词过滤，Kafka分发AI审核任务，Elasticsearch日志检索分析。
7. 帖子服务：使用 Elasticsearch 实现多维度搜索，Redis Pipeline批量更新访问计数，根据 Redis Sorted Set 结构实现热点帖
   排序， MongoDB存储图文内容。
8. 电影服务：Colly爬取豆瓣元数据，审核后存至MongoDB，Redis ZSET维护实时热门榜单。
9. 运维架构：K8s容器化部署（HPA自动扩缩容），Prometheus+Jaeger监控追踪，GitLab CI/CD自动化发布。
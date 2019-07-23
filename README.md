# 基于关联规则的词共现

## 程序运行	

```
go build main.go
./main [OPTIONS]
- OPTIONS:
  -c float
        confidence (default 0.001)
  -fp string
        data file path (default "test/comments100w.txt")
  -np int
        number parallel (default 10)
  -s float
        support rate (default 0.001)
  -split
        whether split sentence
  -w string
        word to find (default "包装")
```

## 算法介绍

- 基于 FPGrowth 算法挖掘词的频繁项
- 支持按分局划分或者按整句划分
- 给定支持度和置信度，输出某词的词共现列表
- FPGrowth 实现参考[FP Tree算法原理总结](https://www.cnblogs.com/pinard/p/6307064.html)

## 算法测试

- 语料 data/comments100w.txt
- 给定 support=0.0001，confidence=0.0001，查找`包装`的共现词
- 返回 `买 味道 月饼 不错 吃 好吃 好`

## 算法流程

1. 建立项头表，过滤支持度低于阈值的词
2. 根据项头表过滤语料词库
3. 根据语料词库建立 FP Tree
4. 根据项头表升序排列搜索，得到个1-频繁项的条件模式基
5. 根据给定的支持度和置信度，从条件模式基中找到词共现列表

## 性能测试

1. 在150w文本下进行测试，设置supportCount为1，splitSent=true

   耗时约10min，内存消耗4G

   设置并发数量为10，cpu利用率大部分时间在100%浮动，在执行到并发程序的时候cpu利用率最高达到800%

   supportCount越大，耗时越少

2. 在150w文本下进行测试，设置supportCount为1，splitSent=false

   耗时约76.8min，内存消耗2.7G

### 分析

如果划分分句，得到的 FP Tree 是浅而宽的，能减少条件模式基挖掘的深度，相反如果不划分分句，得到的 FP Tree 是深而窄的，会加大每个1-频繁项的搜索深度，增加时间复杂度。

不划分分句相比划分分句内存占比减少主要原因是FP Tree宽度减小，减少了FPNode 中next指针的数量
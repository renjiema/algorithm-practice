package design_twitter

// https://leetcode-cn.com/problems/design-twitter/
var void struct{}

type Twitter struct {
	tweets     []int
	followees  map[int]map[int]struct{}
	userTweets map[int][]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{tweets: []int{}, followees: make(map[int]map[int]struct{}), userTweets: make(map[int][]int)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, tweetId)
	this.userTweets[userId] = append(this.userTweets[userId], tweetId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	this.Follow(userId, userId)
	// 获取所有关注的tweet,也可通过并归排序处理所有关注的tweet
	var allTweet []int
	for k := range this.followees[userId] {
		allTweet = append(allTweet, this.userTweets[k]...)
	}
	ts := len(this.tweets) - 1
	s := 0
	var res []int
	for i := ts; i >= 0; i-- {
		if s >= 10 {
			break
		}
		for _, v := range allTweet {
			if this.tweets[i] == v && s < 10 {
				s++
				res = append(res, v)
			}
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followees[followerId] == nil {
		this.followees[followerId] = make(map[int]struct{})
	}
	this.followees[followerId][followeeId] = void
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followees[followerId] == nil {
		return
	}
	delete(this.followees[followerId], followeeId)
}

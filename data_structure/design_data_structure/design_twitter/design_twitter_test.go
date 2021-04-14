package design_twitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(1, 5)
	assert.EqualValues(t, []int{5}, tw.GetNewsFeed(1))
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	assert.EqualValues(t, []int{6, 5}, tw.GetNewsFeed(1))
	tw.Unfollow(1, 2)
	assert.EqualValues(t, []int{5}, tw.GetNewsFeed(1))
}

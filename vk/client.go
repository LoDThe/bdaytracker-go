package vk

import (
	"time"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/api/params"
	limiter "github.com/chatex-com/rate-limiter"
	"github.com/chatex-com/rate-limiter/pkg/config"
	log "github.com/sirupsen/logrus"

	"github.com/lodthe/bdaytracker-go/models"
)

const maxRequestsInSecond = 2
const defaultConcurrency = 1

type Client struct {
	vk          *api.VK
	rateLimiter *limiter.RateLimiter
}

func newRateLimiter() *limiter.RateLimiter {
	cfg := config.NewConfigWithQuotas([]*config.Quota{
		config.NewQuota(maxRequestsInSecond, time.Second),
	})
	cfg.Concurrency = defaultConcurrency

	rateLimiter, _ := limiter.NewRateLimiter(cfg)
	rateLimiter.Start()

	return rateLimiter
}

func NewClient(token string) *Client {
	return &Client{
		vk:          api.NewVK(token),
		rateLimiter: newRateLimiter(),
	}
}

func (c *Client) GetFriends(id int) ([]models.Friend, error) {
	const langRU = 0

	logger := log.WithField("user_id", id)
	logger.Info("start getting VK friends")

	response := <-c.rateLimiter.Execute(func() (interface{}, error) {
		friendsParams := params.NewFriendsGetBuilder()
		friendsParams.Lang(langRU)
		friendsParams.UserID(id)
		friendsParams.Fields([]string{
			"uid",
			"first_name",
			"last_name",
			"bdate",
		})

		return c.vk.FriendsGetFields(friendsParams.Params)
	})

	if response.Error != nil {
		logger.WithError(response.Error).Error("failed to get VK friends")
		return nil, response.Error
	}

	logger.Info("successfully got VK friends")

	resp := response.Result.(api.FriendsGetFieldsResponse)
	friends := make([]models.Friend, len(resp.Items))
	for i := range resp.Items {
		friends[i] = friendObjectToFriend(&resp.Items[i])
	}

	return friends, nil
}

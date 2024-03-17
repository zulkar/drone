package boardwhite

import (
	"context"
	"fmt"

	"github.com/boar-d-white-foundation/drone/iter"
	"github.com/boar-d-white-foundation/drone/leetcode"
)

const (
	defaultDailyHeader = "LeetCode Daily Question"
)

func (s *Service) PublishLCDaily(ctx context.Context) error {
	link, err := leetcode.GetDailyLink(ctx)
	if err != nil {
		return fmt.Errorf("get link: %w", err)
	}

	stickerID, err := iter.PickRandom(s.cfg.DailyStickersIDs)
	if err != nil {
		return fmt.Errorf("get sticker: %w", err)
	}

	return s.publish(ctx, defaultDailyHeader, link, stickerID, keyLCPinnedMessages)
}

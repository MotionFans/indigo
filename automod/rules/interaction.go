package rules

import (
	"github.com/bluesky-social/indigo/automod"
)

var interactionDailyThreshold = 500

// looks for accounts which do frequent interaction churn, such as follow-unfollow.
func InteractionChurnRule(evt *automod.RecordEvent) error {
	did := evt.Account.Identity.DID.String()
	switch evt.Collection {
	case "app.bsky.feed.like":
		evt.Increment("like", did)
		created := evt.GetCount("like", did, automod.PeriodDay)
		deleted := evt.GetCount("unlike", did, automod.PeriodDay)
		ratio := float64(deleted) / float64(created)
		if created > interactionDailyThreshold && deleted > interactionDailyThreshold && ratio > 0.5 {
			evt.Logger.Info("high-like-churn", "created-today", created, "deleted-today", deleted)
			evt.AddAccountFlag("high-like-churn")
		}
	case "app.bsky.graph.follow":
		evt.Increment("follow", did)
		created := evt.GetCount("follow", did, automod.PeriodDay)
		deleted := evt.GetCount("unfollow", did, automod.PeriodDay)
		ratio := float64(deleted) / float64(created)
		if created > interactionDailyThreshold && deleted > interactionDailyThreshold && ratio > 0.5 {
			evt.Logger.Info("high-follow-churn", "created-today", created, "deleted-today", deleted)
			evt.AddAccountFlag("high-follow-churn")
		}
	}
	return nil
}

func DeleteInteractionRule(evt *automod.RecordDeleteEvent) error {
	did := evt.Account.Identity.DID.String()
	switch evt.Collection {
	case "app.bsky.feed.like":
		evt.Increment("unlike", did)
	case "app.bsky.graph.follow":
		evt.Increment("unfollow", did)
	}
	return nil
}
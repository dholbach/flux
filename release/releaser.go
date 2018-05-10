package release

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/weaveworks/flux/update"
)

type Changes interface {
	CalculateRelease(update.ReleaseContext, log.Logger) ([]*update.ControllerUpdate, update.Result, error)
	ReleaseKind() update.ReleaseKind
	ReleaseType() update.ReleaseType
	CommitMessage(update.Result) string
}

func Release(rc *ReleaseContext, changes Changes, logger log.Logger) (results update.Result, err error) {
	fmt.Println("\t\t\trelease.Changes ...")
	defer func(start time.Time) {
		update.ObserveRelease(
			start,
			err == nil,
			changes.ReleaseType(),
			changes.ReleaseKind(),
		)
	}(time.Now())

	logger = log.With(logger, "type", "release")

	updates, results, err := changes.CalculateRelease(rc, logger)
	if err != nil {
		return nil, err
	}

	err = ApplyChanges(rc, updates, logger)
	return results, err
}

func ApplyChanges(rc *ReleaseContext, updates []*update.ControllerUpdate, logger log.Logger) error {
	logger.Log("updates", len(updates))
	if len(updates) == 0 {
		logger.Log("exit", "no images to update for services given")
		return nil
	}

	timer := update.NewStageTimer("write_changes")
	err := rc.WriteUpdates(updates)
	timer.ObserveDuration()
	return err
}

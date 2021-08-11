package rule

import (
	"fmt"
	"strings"

	"github.com/openshift/patchmanager/pkg/config"
	"github.com/openshift/patchmanager/pkg/github"
)

type PullRequestLabelRule struct {
	Config *config.PullRequestLabelRuleConfig
}

func (p *PullRequestLabelRule) Evaluate(pullRequest *github.PullRequest) ([]string, bool) {
	reasons := []string{}
	var skip bool
	for _, l := range pullRequest.Issue.Labels {
		for _, c := range p.Config.RefuseOnLabel {
			if strings.HasPrefix(l.GetName(), c) {
				reasons = append(reasons, fmt.Sprintf("skipping because %q label found", l.GetName()))
				skip = true
			}
		}
		if skip {
			continue
		}
		for _, c := range p.Config.AllowOnLabel {
			if strings.HasSuffix(l.GetName(), c) {
				reasons = append(reasons, fmt.Sprintf("picking because %q label found", l.GetName()))
				skip = false
			}
		}
	}
	return reasons, skip
}

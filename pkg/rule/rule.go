package rule

import "github.com/openshift/patchmanager/pkg/github"

type Ruler interface {
	// Evaluate returns a list of reasons and a bool which indicates whether to skip the pull request.
	Evaluate(*github.PullRequest) ([]string, bool)
}

type MultiRuler struct {
	rulers []Ruler
}

func (m *MultiRuler) Evaluate(pullRequest *github.PullRequest) ([]string, bool) {
	reasons := []string{}
	var skip bool
	for i := range m.rulers {
		r, d := m.rulers[i].Evaluate(pullRequest)
		if len(r) > 0 {
			reasons = append(reasons, r...)
			skip = d
		}
	}
	return reasons, skip
}

func NewMultiRuler(rulers ...Ruler) Ruler {
	return &MultiRuler{rulers: rulers}
}

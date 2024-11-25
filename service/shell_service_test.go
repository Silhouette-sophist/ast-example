package service

import "testing"

func TestQueryProjectDeps(t *testing.T) {
	deps, err := QueryProjectDeps("/Users/silhouette/work-practice/gin-example")
	if err != nil {
		t.Errorf("err %v\n", err)
		return
	}
	t.Logf("success %v\n", deps)
}

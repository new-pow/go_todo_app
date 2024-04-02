package config

import (
	"fmt"
	"testing"
)

// 원하는 환경변수가 있는가?
// 환경변수 기본값이 설정되어 있는가?
func TestNew(t *testing.T) {
	wantPort := 3333
	wantEnv := "dev"
	t.Setenv("TODO_PORT", fmt.Sprint(wantPort)) // 환경 변수 설정

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("want port %d, but %d", wantPort, got.Port)
	}
	if got.Env != wantEnv {
		t.Errorf("want env %s, but %s", wantEnv, got.Env)
	}
}

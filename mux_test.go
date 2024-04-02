package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()                              // 응답을 기록하는 레코드
	r := httptest.NewRequest(http.MethodGet, "/health", nil) // 요청을 생성
	sut := NewMux()
	sut.ServeHTTP(w, r)                         // 요청을 처리
	resp := w.Result()                          // 응답을 얻음
	t.Cleanup(func() { _ = resp.Body.Close() }) // 테스트 종료 시 응답 바디를 닫음

	// 응답 코드가 200인지 확인
	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}
	// 응답 바디가 올바른지 확인
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	want := `{"status": "ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}

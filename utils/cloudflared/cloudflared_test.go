package cloudflared_test

import (
	"os"
	"testing"
	"time"

	"github.com/rabbits0209/komari/utils/cloudflared"
)

func TestRunCloudflared(t *testing.T) {
	os.Setenv("KOMARI_CLOUDFLARED_TOKEN", "test-token")

	err := cloudflared.RunCloudflared()
	if err != nil {
		t.Fatalf("RunCloudflared failed: %v", err)
	}

	// 等待一段时间，确保子进程已启动
	time.Sleep(2 * time.Second)
}

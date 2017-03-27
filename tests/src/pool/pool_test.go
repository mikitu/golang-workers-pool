package tests

import (
	"testing"
	"github.com/mikitu/golang-workers-pool/src/pool"
	"github.com/mikitu/golang-workers-pool/tests/mock"
	"github.com/golang/mock/gomock"
)

func TestExecuteTask(t *testing.T) {
	pool := wpool.NewPool(1)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	task := tests.NewMockTask(mockCtrl)
	task.EXPECT().Execute()
	pool.Exec(task)
	pool.Close()
	pool.Wait()
}

func TestResizePool(t *testing.T) {
	pool := wpool.NewPool(1)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	task := tests.NewMockTask(mockCtrl)
	task.EXPECT().Execute()
	pool.Exec(task)
	pool.Resize(2)
	pool_size := pool.Size()
	if pool_size != 2 {
		t.Fatal("Expected pool size should be 3")
	}
	pool.Close()
	pool.Wait()
}


package pinger_test

import (
	"testing"

	"github.com/NullPrice/kingpinger/mocks"
	"github.com/NullPrice/kingpinger/pkg/pinger"
	"github.com/golang/mock/gomock"
)

func TestProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAdapter := mock_adapter.NewMockAdapter(ctrl)
	mockAdapter.EXPECT().Run().Times(1)
	mockAdapter.EXPECT().ProcessResult().Times(1)
	pinger.Process(mockAdapter)
}

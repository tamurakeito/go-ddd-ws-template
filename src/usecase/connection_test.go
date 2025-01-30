package usecase

import (
	"errors"
	mocks_entity "go-ddd-ws-template/mocks/entity"
	mocks_repository "go-ddd-ws-template/mocks/repository"
	"go-ddd-ws-template/src/domain"
	"go-ddd-ws-template/src/domain/repository"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestNewConnectionUsecase(t *testing.T) {
	type args struct {
		connectionRepo repository.ConnectionRepository
	}
	tests := []struct {
		name string
		args args
		want ConnectionUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnectionUsecase(tt.args.connectionRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnectionUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectionUsecase_HandleConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks_entity.NewMockClient(ctrl)
	mockConnectionRepo := mocks_repository.NewMockConnectionRepository(ctrl)

	type fields struct {
		connectionRepo repository.ConnectionRepository
	}
	type args struct {
		c echo.Context
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}

	tests := []test{
		func() test {
			c := echo.New().NewContext(nil, nil)

			mockConnectionRepo.EXPECT().UpgradeProtocol(c).Return(mockClient, nil)
			mockConnectionRepo.EXPECT().AddClient(mockClient)
			mockConnectionRepo.EXPECT().HandleMessage(mockClient).Return(domain.EOF) // ループを終了させる
			mockConnectionRepo.EXPECT().RemoveClient(mockClient)
			mockClient.EXPECT().Close().Return(nil)

			return test{
				name: "success case",
				fields: fields{
					connectionRepo: mockConnectionRepo,
				},
				args: args{
					c: c,
				},
				wantErr: nil,
			}
		}(),
		func() test {
			c := echo.New().NewContext(nil, nil)
			err := errors.New("failed upgrade http to websocket")

			mockConnectionRepo.EXPECT().UpgradeProtocol(c).Return(nil, err)

			return test{
				name: "failed upgrade http to websocket",
				fields: fields{
					connectionRepo: mockConnectionRepo,
				},
				args: args{
					c: c,
				},
				wantErr: ErrUpgradeProtocol,
			}
		}(),
		func() test {
			c := echo.New().NewContext(nil, nil)
			err := errors.New("failed message handling")

			mockConnectionRepo.EXPECT().UpgradeProtocol(c).Return(mockClient, nil)
			mockConnectionRepo.EXPECT().AddClient(mockClient)
			mockConnectionRepo.EXPECT().HandleMessage(mockClient).Return(err)
			mockConnectionRepo.EXPECT().RemoveClient(mockClient)
			mockClient.EXPECT().Close().Return(nil)

			return test{
				name: "failed message handling",
				fields: fields{
					connectionRepo: mockConnectionRepo,
				},
				args: args{
					c: c,
				},
				wantErr: ErrHandleMessage,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &connectionUsecase{
				connectionRepo: tt.fields.connectionRepo,
			}
			if err := u.HandleConnection(tt.args.c); !errors.Is(err, tt.wantErr) {
				t.Errorf("connectionUsecase.HandleConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

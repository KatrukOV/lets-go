package uc

import (
	"errors"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestNewUserUseCase(t *testing.T) {
	type args struct {
		repo usr.Repository
	}
	var recoveredError = errors.New("repo shouldn't be null")
	//repo := &mockDbRepo{}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		//{"should create new instance of struct initialized",
		//	args{repo},
		//	userUseCaseImpl{repo, bcrypt.CompareHashAndPassword},
		//},
		{"should panic when passed repo param is nil",
			args{nil},
			recoveredError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := recoverPanicIfNeeded(NewUserUseCase, tt.args.repo)
			if err != nil && err != tt.want {
				t.Errorf("NewUserUseCase() panic with err = %v, wantErr %v", err, tt.want)
				return
			}
			assert.New(t).EqualValuesf(tt.want, got, "NewUserUseCase() = %v, wantValue %v", got, tt.want)
		})
	}
}

func recoverPanicIfNeeded(fn func(repo usr.Repository) UserUseCase, arg usr.Repository) (value interface{}, err error) {
	value = fn(arg)
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			log.Printf("recovered panic with expectedErr: %v", err)
		}
	}()
	return
}

func TestUserUseCase_SignInUser(t *testing.T) {
	type fields struct {
		repo usr.Repository
	}
	type args struct {
		email    string
		password string
	}
	expectedUserId := 1
	mockedRepo := func(user *usr.User, err error) usr.Repository {
		mockDbRepo := &mockDbRepo{}
		mockDbRepo.On("GetUserByEmail", mock.Anything).Return(user, err)
		return mockDbRepo
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int
		wantErr   error
	}{
		{
			name:      "should create user successfully",
			fields:    fields{mockedRepo(&usr.User{Id: expectedUserId, Password: "password"}, nil)},
			args:      args{"email", "password"},
			wantValue: expectedUserId,
			wantErr:   nil,
		},
		{
			name:      "should return ErrUserAlreadyExist when repo return ErrUserAlreadyExist",
			fields:    fields{mockedRepo(nil, usr.ErrUserAlreadyExist)},
			args:      args{"email", "password"},
			wantValue: 0,
			wantErr:   usr.ErrUserAlreadyExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := userUseCaseImpl{
				Repo: tt.fields.repo,
				hashComparator: func(hash []byte, text []byte) error {
					return nil
				},
			}
			got, err := uc.SignInUser(tt.args.email, tt.args.password)
			if (err != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("SignInUser() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantValue {
				t.Errorf("SignInUser() got = %v, wantValue %v", got, tt.wantValue)
			}
		})
	}
}

type mockDbRepo struct {
	mock.Mock
}

func (r *mockDbRepo) CreateUser(user usr.User) (int, error) {
	panic("implement me")
}

func (r *mockDbRepo) GetUserByEmail(email string) (*usr.User, error) {
	args := r.Called(email)
	return args.Get(0).(*usr.User), args.Error(1)
}

package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

func TestHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "user1",
			args:    args{password: "123456"},
			wantErr: false,
		},
		{
			name:    "user2",
			args:    args{password: "123456"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Hash() = %v", got)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyPassword(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_BeforeSave(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := u.BeforeSave(); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeSave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Prepare(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			u.Prepare()
		})
	}
}

func TestUser_Validate(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		action string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := u.Validate(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_SaveUser(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := u.SaveUser(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.SaveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.SaveUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_FindAllUsers(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := u.FindAllUsers(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.FindAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.FindAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_FindUserByID(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		db  *gorm.DB
		uid uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := u.FindUserByID(tt.args.db, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.FindUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.FindUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_UpdateAUser(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		db  *gorm.DB
		uid uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := u.UpdateAUser(tt.args.db, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.UpdateAUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.UpdateAUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_DeleteAUser(t *testing.T) {
	type fields struct {
		ID        uint32
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		db  *gorm.DB
		uid uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := u.DeleteAUser(tt.args.db, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.DeleteAUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.DeleteAUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

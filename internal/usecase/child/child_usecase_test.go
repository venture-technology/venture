package child

import (
	"context"
	"reflect"
	"testing"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
)

func TestChildUseCase_Create(t *testing.T) {
	type fields struct {
		childRepository repository.IChildRepository
	}
	type args struct {
		ctx   context.Context
		child *entity.Child
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
			cu := &ChildUseCase{
				childRepository: tt.fields.childRepository,
			}
			if err := cu.Create(tt.args.ctx, tt.args.child); (err != nil) != tt.wantErr {
				t.Errorf("ChildUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChildUseCase_Get(t *testing.T) {
	type fields struct {
		childRepository repository.IChildRepository
	}
	type args struct {
		ctx context.Context
		rg  *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Child
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &ChildUseCase{
				childRepository: tt.fields.childRepository,
			}
			got, err := cu.Get(tt.args.ctx, tt.args.rg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChildUseCase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChildUseCase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChildUseCase_FindAll(t *testing.T) {
	type fields struct {
		childRepository repository.IChildRepository
	}
	type args struct {
		ctx context.Context
		cpf *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Child
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &ChildUseCase{
				childRepository: tt.fields.childRepository,
			}
			got, err := cu.FindAll(tt.args.ctx, tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChildUseCase.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChildUseCase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChildUseCase_Update(t *testing.T) {
	type fields struct {
		childRepository repository.IChildRepository
	}
	type args struct {
		ctx   context.Context
		child *entity.Child
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
			cu := &ChildUseCase{
				childRepository: tt.fields.childRepository,
			}
			if err := cu.Update(tt.args.ctx, tt.args.child); (err != nil) != tt.wantErr {
				t.Errorf("ChildUseCase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChildUseCase_Delete(t *testing.T) {
	type fields struct {
		childRepository repository.IChildRepository
	}
	type args struct {
		ctx context.Context
		rg  *string
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
			cu := &ChildUseCase{
				childRepository: tt.fields.childRepository,
			}
			if err := cu.Delete(tt.args.ctx, tt.args.rg); (err != nil) != tt.wantErr {
				t.Errorf("ChildUseCase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package admins_service_test

import (
	"errors"
	"testing"

	admin_dto "ural-hackaton/internal/dto/admin"
	"ural-hackaton/internal/models"
	admin_service "ural-hackaton/internal/services/handlers/admin"

	"github.com/stretchr/testify/assert"
)

// manual mock implementing admin repo interface
type mockAdminRepo struct {
	createFn        func(admin_dto.CreateAdminDto) (*admin_dto.AdminJoinUserDto, error)
	getAllFn        func() ([]admin_dto.AdminJoinUserDto, error)
	getByIdFn       func(uint64) (*admin_dto.AdminJoinUserDto, error)
	deleteFn        func(uint64) error
	getByFullnameFn func(string) (*admin_dto.AdminJoinUserDto, error)
}

func (m *mockAdminRepo) CreateAdmin(a admin_dto.CreateAdminDto) (*admin_dto.AdminJoinUserDto, error) {
	return m.createFn(a)
}
func (m *mockAdminRepo) GetAllAdmins() ([]admin_dto.AdminJoinUserDto, error) {
	return m.getAllFn()
}
func (m *mockAdminRepo) GetAdminById(id uint64) (*admin_dto.AdminJoinUserDto, error) {
	return m.getByIdFn(id)
}
func (m *mockAdminRepo) DeleteAdmin(id uint64) error {
	return m.deleteFn(id)
}
func (m *mockAdminRepo) GetAdminByFullname(fullname string) (*admin_dto.AdminJoinUserDto, error) {
	return m.getByFullnameFn(fullname)
}

func TestCreateAdmin_Success(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.createFn = func(a admin_dto.CreateAdminDto) (*admin_dto.AdminJoinUserDto, error) {
		return &admin_dto.AdminJoinUserDto{AdminId: 10, User: models.User{Id: a.UserId, FullName: "Ivan Ivanov", Role: "admin"}}, nil
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.CreateAdmin(1)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, uint64(10), res.AdminId)
	assert.Equal(t, uint64(1), res.Id)
}

func TestCreateAdmin_Error(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.createFn = func(a admin_dto.CreateAdminDto) (*admin_dto.AdminJoinUserDto, error) {
		return nil, errors.New("db error")
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.CreateAdmin(1)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAllAdmins_Success(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.getAllFn = func() ([]admin_dto.AdminJoinUserDto, error) {
		return []admin_dto.AdminJoinUserDto{{AdminId: 1, User: models.User{Id: 1, FullName: "A", Role: "admin"}}}, nil
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.GetAllAdmins()

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "A", res[0].FullName)
}

func TestGetAllAdmins_Error(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.getAllFn = func() ([]admin_dto.AdminJoinUserDto, error) {
		return nil, errors.New("db err")
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.GetAllAdmins()

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAdminById_Success(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.getByIdFn = func(id uint64) (*admin_dto.AdminJoinUserDto, error) {
		return &admin_dto.AdminJoinUserDto{AdminId: id, User: models.User{Id: 2, FullName: "B"}}, nil
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.GetAdminById(2)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, uint64(2), res.AdminId)
}

func TestGetAdminById_Error(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.getByIdFn = func(id uint64) (*admin_dto.AdminJoinUserDto, error) {
		return nil, errors.New("not found")
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.GetAdminById(5)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestDeleteAdmin_Success(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.deleteFn = func(id uint64) error { return nil }

	svc := admin_service.Init(mock, nil)
	err := svc.DeleteAdmin(3)

	assert.NoError(t, err)
}

func TestDeleteAdmin_Error(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.deleteFn = func(id uint64) error { return errors.New("delete failed") }

	svc := admin_service.Init(mock, nil)
	err := svc.DeleteAdmin(3)

	assert.Error(t, err)
}

func TestGetAdminByFullname_Success(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.getByFullnameFn = func(fullname string) (*admin_dto.AdminJoinUserDto, error) {
		return &admin_dto.AdminJoinUserDto{AdminId: 7, User: models.User{Id: 7, FullName: fullname}}, nil
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.GetAdminByFullname("Ivan Ivanov")

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Ivan Ivanov", res.FullName)
}

func TestGetAdminByFullname_Error(t *testing.T) {
	mock := &mockAdminRepo{}
	mock.getByFullnameFn = func(fullname string) (*admin_dto.AdminJoinUserDto, error) {
		return nil, errors.New("not found")
	}

	svc := admin_service.Init(mock, nil)
	res, err := svc.GetAdminByFullname("Nobody")

	assert.Error(t, err)
	assert.Nil(t, res)
}

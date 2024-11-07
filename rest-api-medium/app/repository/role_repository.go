package repository

import (
	"github.com/rynfkn/rest-api-medium/app/domain/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole() ([]dao.Role, error)
	FindRoleById(id int) (dao.Role, error)
	CreateRole(role *dao.Role) (dao.Role, error)
	UpdateRole(role *dao.Role) (dao.Role, error)
	DeleteRoleById(id int) error
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

// DeleteRoleById implements RoleRepository.
func (r *RoleRepositoryImpl) DeleteRoleById(id int) error {
	panic("unimplemented")
}

// FindRoleById implements RoleRepository.
func (r *RoleRepositoryImpl) FindRoleById(id int) (dao.Role, error) {
	panic("unimplemented")
}

func (r *RoleRepositoryImpl) FindAllRole() ([]dao.Role, error) {
	var roles []dao.Role

	var err = r.db.Preload("Role").Find(&roles).Error
	if err != nil {
		log.Error("Error in FindAllRoles: ", err)
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepositoryImpl) FindRoleByID(id int) (dao.Role, error) {
	var role dao.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn("Role not found with ID:", id)
			return dao.Role{}, nil // or return a custom error
		}
		log.Error("Error in FindRoleByID: ", err)
		return dao.Role{}, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) CreateRole(role *dao.Role) (dao.Role, error) {
	err := r.db.Create(role).Error
	if err != nil {
		log.Error("Error in CreateRole: ", err)
		return dao.Role{}, err
	}
	return *role, nil
}

func (r *RoleRepositoryImpl) UpdateRole(role *dao.Role) (dao.Role, error) {
	err := r.db.Save(role).Error
	if err != nil {
		log.Error("Error in UpdateRole: ", err)
		return dao.Role{}, err
	}
	return *role, nil
}

func (r *RoleRepositoryImpl) DeleteRoleByID(id int) error {
	err := r.db.Delete(&dao.Role{}, id).Error
	if err != nil {
		log.Error("Error in DeleteRoleByID: ", err)
		return err
	}
	return nil
}

func RoleRepositoryInit(db *gorm.DB) RoleRepository {
	db.AutoMigrate(&dao.Role{})
	return &RoleRepositoryImpl{db: db}
}

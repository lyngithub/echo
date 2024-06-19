package services

import (
	"echo/common/logger"
	"echo/cons"
	"echo/daos/mysql"
	"echo/models"
	"echo/models/vo"
	"echo/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type RoleService struct {
}

func NewServiceRole() *RoleService {
	return &RoleService{}
}

func (s *RoleService) EditRole(uid int64, params *vo.RoleParams) error {
	var (
		role *models.Role
		err  error
	)
	if 0 == params.Id {
		r, err := mysql.Role.GetBeanById(params.ParentId)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("query role by id --> --> Fail [%d]", uid), zap.Error(err))
			return errors.New("Select the superior role exception")
		}
		superiorIds := r.SuperiorIds
		if "" == superiorIds {
			superiorIds, _ = utils.ToString(params.ParentId)
		} else {
			ridStr, _ := utils.ToString(params.ParentId)
			superiorIds = superiorIds + "," + ridStr
		}

		role = &models.Role{
			ParentId:    params.ParentId,
			SuperiorIds: superiorIds,
			RoleName:    params.RoleName,
			RoleSort:    0,
			Status:      params.Status,
			Remark:      params.Remark,
		}
		err = mysql.Role.InsertBean(uid, role)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("InsertRole --> Fail [%v]", role), zap.Error(err))
			return err
		}
	} else {
		role, err = mysql.Role.GetBeanById(params.Id)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("query role by id --> Fail [%d]", params.Id), zap.Error(err))
			return errors.New("query role exception")
		}
		role.RoleName = params.RoleName
		role.Status = params.Status
		role.Remark = params.Remark
		err = mysql.Role.UpdateBean(uid, role)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("edit role --> Fail [%v]", role), zap.Error(err))
			return err
		}
		// 删除权限
		mysql.Menu.DeleteRoleMenuByRoleId(role.ID)
	}
	// 添加权限
	for _, menuId := range params.Menus {
		mysql.Menu.InsertRoleMenu(&models.RoleMenu{
			RoleId: role.ID,
			MenuId: menuId,
		})
	}

	return nil
}

func (s *RoleService) DeleteRoleByIds(ids []int64) error {
	for _, id := range ids {
		list := mysql.Role.FindRoleIdsByRoleId(id)
		if list != nil && 0 < len(list) {
			return errors.New(fmt.Sprintf("The role with id %d is in use and cannot be deleted", id))
		}
		mysql.Role.DeleteBeanById(id)
		mysql.Role.DeleteUserRoleByRoleId(id)
		mysql.Menu.DeleteRoleMenuByRoleId(id)
	}
	return nil
}

func (s *RoleService) GetRoleTree(userId int64, status string) ([]*vo.RoleTreeVo, error) {
	user, err := mysql.User.GetBeanById(userId)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get user --> Fail [%d]", userId), zap.Error(err))
		return nil, err
	}
	var menus []*models.Role
	if utils.IsAdmin(user) {
		menus, err = mysql.Role.FindRoleByIds(nil, status)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get role list--> Fail [%d]", userId), zap.Error(err))
			return nil, err
		}
	} else {
		mids := mysql.Role.FindRoleIdsByUserId(userId)
		if 0 == len(mids) {
			return []*vo.RoleTreeVo{}, nil
		}
		menus, err = mysql.Role.FindRoleByIds(mids, status)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get role list--> Fail [%d | %v]", userId, mids), zap.Error(err))
			return nil, err
		}
	}

	// 找到最上级
	ids := s.findParentIds(menus)
	vos := make([]*vo.RoleTreeVo, 0)
	for _, id := range ids {
		tree := s.newTree(id, menus)
		vos = append(vos, tree...)
	}
	return vos, nil
}

func (s *RoleService) findParentIds(menus []*models.Role) []int64 {
	parentIds := make(map[int64]struct{})
	ids := make(map[int64]struct{})
	for _, menu := range menus {
		_, has := parentIds[menu.ParentId]
		_, has2 := ids[menu.ParentId]
		if !has && !has2 {
			parentIds[menu.ParentId] = struct{}{}
		}
		ids[menu.ID] = struct{}{}
	}
	arr := make([]int64, 0)
	for k, _ := range parentIds {
		arr = append(arr, k)
	}
	return arr
}

func (s *RoleService) newTree(id int64, menus []*models.Role) []*vo.RoleTreeVo {
	tree := make([]*vo.RoleTreeVo, 0)
	for _, menu := range menus {
		if menu.ParentId == id {
			child := &vo.RoleTreeVo{
				Id:   menu.ID,
				Name: menu.RoleName,
			}
			child.Children = s.newTree(menu.ID, menus)
			tree = append(tree, child)
		}
	}
	return tree
}

func (s *RoleService) FindRole(userId int64, status string) ([]*vo.RoleVo, error) {
	tree, err := s.GetRoleTree(userId, status)
	if err != nil {
		return nil, err
	}
	return s.findRole(false, 0, tree), nil
}

/*
has:是否存在上级
has:是否同级最后一个
num:几级循环
*/
func (s *RoleService) findRole(has bool, num int, tree []*vo.RoleTreeVo) []*vo.RoleVo {
	list := make([]*vo.RoleVo, 0)
	i := 0
	treeLen := len(tree)
	for _, role := range tree {
		i = i + 1
		r, _ := mysql.Role.GetBeanById(role.Id)
		menuIds := mysql.Menu.FindMenuIdsByRoleId(role.Id)
		bean := &vo.RoleVo{
			Id:             r.ID,
			ParentId:       r.ParentId,
			RoleName:       r.RoleName,
			PrefixRoleName: r.RoleName,
			Status:         r.Status,
			Remark:         r.Remark,
			Menus:          menuIds,
		}
		if !r.Created.IsZero() {
			bean.Created = r.Created.Format(cons.TIMEDATETIME)
		}
		if !r.Updated.IsZero() {
			bean.Updated = r.Updated.Format(cons.TIMEDATETIME)
		}
		if has {
			prefix := ""
			for j := 0; j < num-1; j++ {
				prefix = "│ " + prefix
			}
			if i == treeLen {
				bean.PrefixRoleName = prefix + "└ " + bean.PrefixRoleName
			} else {
				bean.PrefixRoleName = prefix + "├ " + bean.PrefixRoleName
			}
		}
		list = append(list, bean)
		if 0 != len(role.Children) {
			list = append(list, s.findRole(true, num+1, role.Children)...)
		}
	}
	return list
}

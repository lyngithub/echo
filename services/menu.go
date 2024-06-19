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
	"sort"
)

type MenuService struct {
}

func NewServiceMenu() *MenuService {
	return &MenuService{}
}

func (s *MenuService) EditMenu(uid int64, params *vo.MenuParams) error {
	var (
		menu *models.Menu
		err  error
	)
	if 0 == params.Id {
		menu = &models.Menu{
			ParentId:       params.ParentId,
			MenuName:       params.MenuName,
			Weights:        params.Weights,
			Method:         params.Method,
			Url:            params.Url,
			Pages:          params.Pages,
			MenuType:       params.MenuType,
			Classification: params.Classification,
			Visible:        params.Visible,
			Icon:           params.Icon,
			Remark:         params.Remark,
		}
		err = mysql.Menu.InsertBean(uid, menu)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("insert menu --> Fail [%v]", menu), zap.Error(err))
			return err
		}
		mysql.Menu.InsertRoleMenu(&models.RoleMenu{
			RoleId: 1,
			MenuId: menu.ID,
		})
	} else {
		menu, err = mysql.Menu.GetBeanById(params.Id)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("query menu by id --> Fail [%d]", params.Id), zap.Error(err))
			return errors.New("Query menu exception")
		}
		menu.MenuName = params.MenuName
		menu.Weights = params.Weights
		menu.Method = params.Method
		menu.Url = params.Url
		menu.Pages = params.Pages
		menu.MenuType = params.MenuType
		menu.Classification = params.Classification
		menu.Visible = params.Visible
		menu.Icon = params.Icon
		menu.Remark = params.Remark
		err = mysql.Menu.UpdateBean(uid, menu)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("edit menu --> Fail [%v]", menu), zap.Error(err))
			return err
		}
	}
	return nil
}

func (s *MenuService) DeleteMenuByIds(ids []int64) error {
	for _, id := range ids {
		list := mysql.Menu.FindMenuIdsByMenuId(id)
		if list != nil && 0 < len(list) {
			return errors.New(fmt.Sprintf("The role with id %d is in use and cannot be deleted", id))
		}
		mysql.Menu.DeleteBeanById(id)
		mysql.Menu.DeleteRoleMenuByMenuId(id)
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
func (s *MenuService) FindMenu(userId int64) ([]*vo.MenuVo, error) {
	user, err := mysql.User.GetBeanById(userId)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get user --> Fail [%d]", userId), zap.Error(err))
		return nil, err
	}
	var menus []*models.Menu
	if utils.IsAdmin(user) {
		menus, err = mysql.Menu.FindMenuByIds(nil, "", 0)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d]", userId), zap.Error(err))
			return nil, err
		}
	} else {
		mids := mysql.Menu.FindMenuIdsByUserId(userId)
		if 0 == len(mids) {
			return []*vo.MenuVo{}, nil
		}
		menus, err = mysql.Menu.FindMenuByIds(mids, "", 0)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d | %v]", userId, mids), zap.Error(err))
			return nil, err
		}
	}

	// 找到最上级
	ids := s.findParentIds(menus)
	vos := make([]*vo.MenuVo, 0)
	for _, id := range ids {
		tree := s.newTree2(id, menus)
		vos = append(vos, tree...)
	}
	return vos, nil
}

func (s *MenuService) newTree2(id int64, menus []*models.Menu) []*vo.MenuVo {
	tree := make([]*vo.MenuVo, 0)
	for _, menu := range menus {
		//if "F" == menu.MenuType {
		//	continue
		//}
		if menu.ParentId == id {
			child := &vo.MenuVo{
				Id:             menu.ID,
				ParentId:       menu.ParentId,
				MenuName:       menu.MenuName,
				Weights:        menu.Weights,
				Method:         menu.Method,
				Url:            menu.Url,
				Pages:          menu.Pages,
				MenuType:       menu.MenuType,
				Classification: menu.Classification,
				Visible:        menu.Visible,
				Icon:           menu.Icon,
				Remark:         menu.Remark,
			}
			if !menu.Created.IsZero() {
				child.Created = menu.Created.Format(cons.TIMEDATETIME)
			}
			if !menu.Updated.IsZero() {
				child.Updated = menu.Updated.Format(cons.TIMEDATETIME)
			}
			child.Children = s.newTree2(menu.ID, menus)
			tree = append(tree, child)
		}
	}
	return tree
}

// ------------------------------------------------------------------------------------------------
func (s *MenuService) GetMenuTree(userId int64) ([]*vo.MenuTreeVo, error) {
	user, err := mysql.User.GetBeanById(userId)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get user --> Fail [%d]", userId), zap.Error(err))
		return nil, err
	}
	var menus []*models.Menu
	if utils.IsAdmin(user) {
		menus, err = mysql.Menu.FindMenuByIds(nil, "0", 0)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d]", userId), zap.Error(err))
			return nil, err
		}
	} else {
		mids := mysql.Menu.FindMenuIdsByUserId(userId)
		if 0 == len(mids) {
			return []*vo.MenuTreeVo{}, nil
		}
		menus, err = mysql.Menu.FindMenuByIds(mids, "0", 0)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d | %v]", userId, mids), zap.Error(err))
			return nil, err
		}
	}

	// 找到最上级
	ids := s.findParentIds(menus)
	vos := make([]*vo.MenuTreeVo, 0)
	for _, id := range ids {
		tree := s.newTree(id, menus)
		vos = append(vos, tree...)
	}

	// 权重排序
	s.sort(vos)
	return vos, nil
}

func (s *MenuService) sort(list []*vo.MenuTreeVo) {
	sort.Sort(byWeights(list))
	for _, bean := range list {
		if 0 < len(bean.Children) {
			s.sort(bean.Children)
		}
	}
}

// 排序
type byWeights []*vo.MenuTreeVo

func (a byWeights) Len() int           { return len(a) }
func (a byWeights) Less(i, j int) bool { return a[j].Weights > a[i].Weights }
func (a byWeights) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (s *MenuService) findParentIds(menus []*models.Menu) []int64 {
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

func (s *MenuService) newTree(id int64, menus []*models.Menu) []*vo.MenuTreeVo {
	tree := make([]*vo.MenuTreeVo, 0)
	for _, menu := range menus {
		if menu.ParentId == id {
			child := &vo.MenuTreeVo{
				Id:       menu.ID,
				Name:     menu.MenuName,
				Weights:  menu.Weights,
				Icon:     menu.Icon,
				MenuType: menu.MenuType,
				Pages:    menu.Pages,
			}
			child.Children = s.newTree(menu.ID, menus)
			tree = append(tree, child)
		}
	}
	return tree
}

// -------------------------------------------------------------------------------------------------------------
func (s *MenuService) GetRoleMenuTree(rid int64) ([]*vo.MenuTreeVo, error) {
	mids := mysql.Menu.FindMenuIdsByRoleId(rid)
	if 1 != rid && 0 == len(mids) {
		return []*vo.MenuTreeVo{}, nil
	}
	if 1 == rid {
		mids = nil
	}
	menus, err := mysql.Menu.FindMenuByIds(mids, "0", 0)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d | %v]", rid, mids), zap.Error(err))
		return nil, err
	}

	// 找到最上级
	ids := s.findParentIds(menus)
	vos := make([]*vo.MenuTreeVo, 0)
	for _, id := range ids {
		tree := s.newTree(id, menus)
		vos = append(vos, tree...)
	}
	return vos, nil
}

// -------------------------------------------------------------------------------------------------------------
func (s *MenuService) FindMenuByParentId(userId, mid int64) ([]*vo.MenuVo, error) {
	user, err := mysql.User.GetBeanById(userId)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get user --> Fail [%d]", userId), zap.Error(err))
		return nil, err
	}
	var menus []*models.Menu
	if utils.IsAdmin(user) {
		menus, err = mysql.Menu.FindMenuByIds(nil, "0", mid)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d]", userId), zap.Error(err))
			return nil, err
		}
	} else {
		mids := mysql.Menu.FindMenuIdsByUserId(userId)
		if 0 == len(mids) {
			return []*vo.MenuVo{}, nil
		}
		menus, err = mysql.Menu.FindMenuByIds(mids, "0", mid)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu list --> Fail [%d | %v]", userId, mids), zap.Error(err))
			return nil, err
		}
	}

	list := make([]*vo.MenuVo, 0)
	for _, menu := range menus {
		v := &vo.MenuVo{
			Id:       menu.ID,
			ParentId: menu.ParentId,
			MenuName: menu.MenuName,
			Weights:  menu.Weights,
			Method:   menu.Method,
			Url:      menu.Url,
			Pages:    menu.Pages,
			MenuType: menu.MenuType,
			Visible:  menu.Visible,
			Icon:     menu.Icon,
			Remark:   menu.Remark,
			Children: make([]*vo.MenuVo, 0),
		}
		if !menu.Created.IsZero() {
			v.Created = menu.Created.Format(cons.TIMEDATETIME)
		}
		if !menu.Updated.IsZero() {
			v.Updated = menu.Updated.Format(cons.TIMEDATETIME)
		}
		list = append(list, v)
	}
	return list, err
}

## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。


## 项目部署

- 打包 
<br>
  ```
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o .
  ```
- 部署
  <br>
  ```
  nohup ./echo > nohub.out 2>&1 &
  ```


## 体验


// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package role_dept

// Fill with you ideas below.

//获取角色的部门数据
func GetRoleDepts(roleId int64) ([]int64, error) {
	entity, err := Model.All(Columns.RoleId, roleId)
	if err != nil {
		return nil, err
	}
	d := make([]int64, len(entity))
	for k, v := range entity {
		d[k] = v.DeptId
	}
	return d, nil
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateSysDictInput represents a mutation input for creating sysdicts.
type CreateSysDictInput struct {
	IsDel     *bool
	Memo      *string
	Sort      *int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	IsActive  *bool
	NameCn    string
	NameEn    string
}

// Mutate applies the CreateSysDictInput on the SysDictCreate builder.
func (i *CreateSysDictInput) Mutate(m *SysDictCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetNameCn(i.NameCn)
	m.SetNameEn(i.NameEn)
}

// SetInput applies the change-set in the CreateSysDictInput on the create builder.
func (c *SysDictCreate) SetInput(i CreateSysDictInput) *SysDictCreate {
	i.Mutate(c)
	return c
}

// UpdateSysDictInput represents a mutation input for updating sysdicts.
type UpdateSysDictInput struct {
	IsDel          *bool
	Memo           *string
	Sort           *int32
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	IsActive       *bool
	NameCn         *string
	NameEn         *string
}

// Mutate applies the UpdateSysDictInput on the SysDictMutation.
func (i *UpdateSysDictInput) Mutate(m *SysDictMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.NameCn; v != nil {
		m.SetNameCn(*v)
	}
	if v := i.NameEn; v != nil {
		m.SetNameEn(*v)
	}
}

// SetInput applies the change-set in the UpdateSysDictInput on the update builder.
func (u *SysDictUpdate) SetInput(i UpdateSysDictInput) *SysDictUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysDictInput on the update-one builder.
func (u *SysDictUpdateOne) SetInput(i UpdateSysDictInput) *SysDictUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysDictItemInput represents a mutation input for creating sysdictitems.
type CreateSysDictItemInput struct {
	IsDel     *bool
	Memo      *string
	Sort      *int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	IsActive  *bool
	Label     string
	Value     int
	DictID    string
}

// Mutate applies the CreateSysDictItemInput on the SysDictItemCreate builder.
func (i *CreateSysDictItemInput) Mutate(m *SysDictItemCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetLabel(i.Label)
	m.SetValue(i.Value)
	m.SetDictID(i.DictID)
}

// SetInput applies the change-set in the CreateSysDictItemInput on the create builder.
func (c *SysDictItemCreate) SetInput(i CreateSysDictItemInput) *SysDictItemCreate {
	i.Mutate(c)
	return c
}

// UpdateSysDictItemInput represents a mutation input for updating sysdictitems.
type UpdateSysDictItemInput struct {
	IsDel          *bool
	Memo           *string
	Sort           *int32
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	IsActive       *bool
	Label          *string
	Value          *int
	DictID         *string
}

// Mutate applies the UpdateSysDictItemInput on the SysDictItemMutation.
func (i *UpdateSysDictItemInput) Mutate(m *SysDictItemMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Label; v != nil {
		m.SetLabel(*v)
	}
	if v := i.Value; v != nil {
		m.SetValue(*v)
	}
	if v := i.DictID; v != nil {
		m.SetDictID(*v)
	}
}

// SetInput applies the change-set in the UpdateSysDictItemInput on the update builder.
func (u *SysDictItemUpdate) SetInput(i UpdateSysDictItemInput) *SysDictItemUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysDictItemInput on the update-one builder.
func (u *SysDictItemUpdateOne) SetInput(i UpdateSysDictItemInput) *SysDictItemUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysJwtBlockInput represents a mutation input for creating sysjwtblocks.
type CreateSysJwtBlockInput struct {
	IsDel     *bool
	Memo      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	IsActive  *bool
	Jwt       string
}

// Mutate applies the CreateSysJwtBlockInput on the SysJwtBlockCreate builder.
func (i *CreateSysJwtBlockInput) Mutate(m *SysJwtBlockCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetJwt(i.Jwt)
}

// SetInput applies the change-set in the CreateSysJwtBlockInput on the create builder.
func (c *SysJwtBlockCreate) SetInput(i CreateSysJwtBlockInput) *SysJwtBlockCreate {
	i.Mutate(c)
	return c
}

// UpdateSysJwtBlockInput represents a mutation input for updating sysjwtblocks.
type UpdateSysJwtBlockInput struct {
	IsDel          *bool
	Memo           *string
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	IsActive       *bool
	Jwt            *string
}

// Mutate applies the UpdateSysJwtBlockInput on the SysJwtBlockMutation.
func (i *UpdateSysJwtBlockInput) Mutate(m *SysJwtBlockMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Jwt; v != nil {
		m.SetJwt(*v)
	}
}

// SetInput applies the change-set in the UpdateSysJwtBlockInput on the update builder.
func (u *SysJwtBlockUpdate) SetInput(i UpdateSysJwtBlockInput) *SysJwtBlockUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysJwtBlockInput on the update-one builder.
func (u *SysJwtBlockUpdateOne) SetInput(i UpdateSysJwtBlockInput) *SysJwtBlockUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysLoggingInput represents a mutation input for creating sysloggings.
type CreateSysLoggingInput struct {
	IsDel      *bool
	Memo       *string
	Level      string
	TraceID    string
	UserID     string
	Tag        string
	Version    string
	Message    string
	Data       *string
	ErrorStack string
	CreatedAt  *time.Time
}

// Mutate applies the CreateSysLoggingInput on the SysLoggingCreate builder.
func (i *CreateSysLoggingInput) Mutate(m *SysLoggingCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	m.SetLevel(i.Level)
	m.SetTraceID(i.TraceID)
	m.SetUserID(i.UserID)
	m.SetTag(i.Tag)
	m.SetVersion(i.Version)
	m.SetMessage(i.Message)
	if v := i.Data; v != nil {
		m.SetData(*v)
	}
	m.SetErrorStack(i.ErrorStack)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateSysLoggingInput on the create builder.
func (c *SysLoggingCreate) SetInput(i CreateSysLoggingInput) *SysLoggingCreate {
	i.Mutate(c)
	return c
}

// UpdateSysLoggingInput represents a mutation input for updating sysloggings.
type UpdateSysLoggingInput struct {
	IsDel      *bool
	Memo       *string
	Level      *string
	TraceID    *string
	UserID     *string
	Tag        *string
	Version    *string
	Message    *string
	ErrorStack *string
}

// Mutate applies the UpdateSysLoggingInput on the SysLoggingMutation.
func (i *UpdateSysLoggingInput) Mutate(m *SysLoggingMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Level; v != nil {
		m.SetLevel(*v)
	}
	if v := i.TraceID; v != nil {
		m.SetTraceID(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Tag; v != nil {
		m.SetTag(*v)
	}
	if v := i.Version; v != nil {
		m.SetVersion(*v)
	}
	if v := i.Message; v != nil {
		m.SetMessage(*v)
	}
	if v := i.ErrorStack; v != nil {
		m.SetErrorStack(*v)
	}
}

// SetInput applies the change-set in the UpdateSysLoggingInput on the update builder.
func (u *SysLoggingUpdate) SetInput(i UpdateSysLoggingInput) *SysLoggingUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysLoggingInput on the update-one builder.
func (u *SysLoggingUpdateOne) SetInput(i UpdateSysLoggingInput) *SysLoggingUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysMenuInput represents a mutation input for creating sysmenus.
type CreateSysMenuInput struct {
	IsDel      *bool
	Memo       *string
	Sort       *int32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
	IsActive   *bool
	Name       string
	Icon       string
	Router     string
	IsShow     *bool
	ParentID   *string
	ParentPath *string
	Level      int32
	IsLeaf     *bool
}

// Mutate applies the CreateSysMenuInput on the SysMenuCreate builder.
func (i *CreateSysMenuInput) Mutate(m *SysMenuCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetName(i.Name)
	m.SetIcon(i.Icon)
	m.SetRouter(i.Router)
	if v := i.IsShow; v != nil {
		m.SetIsShow(*v)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.ParentPath; v != nil {
		m.SetParentPath(*v)
	}
	m.SetLevel(i.Level)
	if v := i.IsLeaf; v != nil {
		m.SetIsLeaf(*v)
	}
}

// SetInput applies the change-set in the CreateSysMenuInput on the create builder.
func (c *SysMenuCreate) SetInput(i CreateSysMenuInput) *SysMenuCreate {
	i.Mutate(c)
	return c
}

// UpdateSysMenuInput represents a mutation input for updating sysmenus.
type UpdateSysMenuInput struct {
	IsDel           *bool
	Memo            *string
	Sort            *int32
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
	ClearDeletedAt  bool
	IsActive        *bool
	Name            *string
	Icon            *string
	Router          *string
	IsShow          *bool
	ParentID        *string
	ClearParentID   bool
	ParentPath      *string
	ClearParentPath bool
	Level           *int32
	IsLeaf          *bool
	ClearIsLeaf     bool
}

// Mutate applies the UpdateSysMenuInput on the SysMenuMutation.
func (i *UpdateSysMenuInput) Mutate(m *SysMenuMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Icon; v != nil {
		m.SetIcon(*v)
	}
	if v := i.Router; v != nil {
		m.SetRouter(*v)
	}
	if v := i.IsShow; v != nil {
		m.SetIsShow(*v)
	}
	if i.ClearParentID {
		m.ClearParentID()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearParentPath {
		m.ClearParentPath()
	}
	if v := i.ParentPath; v != nil {
		m.SetParentPath(*v)
	}
	if v := i.Level; v != nil {
		m.SetLevel(*v)
	}
	if i.ClearIsLeaf {
		m.ClearIsLeaf()
	}
	if v := i.IsLeaf; v != nil {
		m.SetIsLeaf(*v)
	}
}

// SetInput applies the change-set in the UpdateSysMenuInput on the update builder.
func (u *SysMenuUpdate) SetInput(i UpdateSysMenuInput) *SysMenuUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysMenuInput on the update-one builder.
func (u *SysMenuUpdateOne) SetInput(i UpdateSysMenuInput) *SysMenuUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysMenuActionInput represents a mutation input for creating sysmenuactions.
type CreateSysMenuActionInput struct {
	IsDel     *bool
	Sort      *int32
	IsActive  *bool
	Memo      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	MenuID    string
	Code      string
	Name      string
}

// Mutate applies the CreateSysMenuActionInput on the SysMenuActionCreate builder.
func (i *CreateSysMenuActionInput) Mutate(m *SysMenuActionCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetMenuID(i.MenuID)
	m.SetCode(i.Code)
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateSysMenuActionInput on the create builder.
func (c *SysMenuActionCreate) SetInput(i CreateSysMenuActionInput) *SysMenuActionCreate {
	i.Mutate(c)
	return c
}

// UpdateSysMenuActionInput represents a mutation input for updating sysmenuactions.
type UpdateSysMenuActionInput struct {
	IsDel          *bool
	Sort           *int32
	IsActive       *bool
	Memo           *string
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	MenuID         *string
	Code           *string
	Name           *string
}

// Mutate applies the UpdateSysMenuActionInput on the SysMenuActionMutation.
func (i *UpdateSysMenuActionInput) Mutate(m *SysMenuActionMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.MenuID; v != nil {
		m.SetMenuID(*v)
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateSysMenuActionInput on the update builder.
func (u *SysMenuActionUpdate) SetInput(i UpdateSysMenuActionInput) *SysMenuActionUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysMenuActionInput on the update-one builder.
func (u *SysMenuActionUpdateOne) SetInput(i UpdateSysMenuActionInput) *SysMenuActionUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysMenuActionResourceInput represents a mutation input for creating sysmenuactionresources.
type CreateSysMenuActionResourceInput struct {
	IsDel     *bool
	Sort      *int32
	Memo      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	IsActive  *bool
	Method    string
	Path      string
	ActionID  string
}

// Mutate applies the CreateSysMenuActionResourceInput on the SysMenuActionResourceCreate builder.
func (i *CreateSysMenuActionResourceInput) Mutate(m *SysMenuActionResourceCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetMethod(i.Method)
	m.SetPath(i.Path)
	m.SetActionID(i.ActionID)
}

// SetInput applies the change-set in the CreateSysMenuActionResourceInput on the create builder.
func (c *SysMenuActionResourceCreate) SetInput(i CreateSysMenuActionResourceInput) *SysMenuActionResourceCreate {
	i.Mutate(c)
	return c
}

// UpdateSysMenuActionResourceInput represents a mutation input for updating sysmenuactionresources.
type UpdateSysMenuActionResourceInput struct {
	IsDel          *bool
	Sort           *int32
	Memo           *string
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	IsActive       *bool
	Method         *string
	Path           *string
	ActionID       *string
}

// Mutate applies the UpdateSysMenuActionResourceInput on the SysMenuActionResourceMutation.
func (i *UpdateSysMenuActionResourceInput) Mutate(m *SysMenuActionResourceMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Method; v != nil {
		m.SetMethod(*v)
	}
	if v := i.Path; v != nil {
		m.SetPath(*v)
	}
	if v := i.ActionID; v != nil {
		m.SetActionID(*v)
	}
}

// SetInput applies the change-set in the UpdateSysMenuActionResourceInput on the update builder.
func (u *SysMenuActionResourceUpdate) SetInput(i UpdateSysMenuActionResourceInput) *SysMenuActionResourceUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysMenuActionResourceInput on the update-one builder.
func (u *SysMenuActionResourceUpdateOne) SetInput(i UpdateSysMenuActionResourceInput) *SysMenuActionResourceUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysRoleInput represents a mutation input for creating sysroles.
type CreateSysRoleInput struct {
	IsDel     *bool
	IsActive  *bool
	Sort      *int32
	Memo      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Name      string
}

// Mutate applies the CreateSysRoleInput on the SysRoleCreate builder.
func (i *CreateSysRoleInput) Mutate(m *SysRoleCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateSysRoleInput on the create builder.
func (c *SysRoleCreate) SetInput(i CreateSysRoleInput) *SysRoleCreate {
	i.Mutate(c)
	return c
}

// UpdateSysRoleInput represents a mutation input for updating sysroles.
type UpdateSysRoleInput struct {
	IsDel          *bool
	IsActive       *bool
	Sort           *int32
	Memo           *string
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	Name           *string
}

// Mutate applies the UpdateSysRoleInput on the SysRoleMutation.
func (i *UpdateSysRoleInput) Mutate(m *SysRoleMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateSysRoleInput on the update builder.
func (u *SysRoleUpdate) SetInput(i UpdateSysRoleInput) *SysRoleUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysRoleInput on the update-one builder.
func (u *SysRoleUpdateOne) SetInput(i UpdateSysRoleInput) *SysRoleUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysRoleMenuInput represents a mutation input for creating sysrolemenus.
type CreateSysRoleMenuInput struct {
	IsDel     *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	RoleID    string
	MenuID    string
	ActionID  *string
}

// Mutate applies the CreateSysRoleMenuInput on the SysRoleMenuCreate builder.
func (i *CreateSysRoleMenuInput) Mutate(m *SysRoleMenuCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetRoleID(i.RoleID)
	m.SetMenuID(i.MenuID)
	if v := i.ActionID; v != nil {
		m.SetActionID(*v)
	}
}

// SetInput applies the change-set in the CreateSysRoleMenuInput on the create builder.
func (c *SysRoleMenuCreate) SetInput(i CreateSysRoleMenuInput) *SysRoleMenuCreate {
	i.Mutate(c)
	return c
}

// UpdateSysRoleMenuInput represents a mutation input for updating sysrolemenus.
type UpdateSysRoleMenuInput struct {
	IsDel          *bool
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	RoleID         *string
	MenuID         *string
	ActionID       *string
	ClearActionID  bool
}

// Mutate applies the UpdateSysRoleMenuInput on the SysRoleMenuMutation.
func (i *UpdateSysRoleMenuInput) Mutate(m *SysRoleMenuMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.RoleID; v != nil {
		m.SetRoleID(*v)
	}
	if v := i.MenuID; v != nil {
		m.SetMenuID(*v)
	}
	if i.ClearActionID {
		m.ClearActionID()
	}
	if v := i.ActionID; v != nil {
		m.SetActionID(*v)
	}
}

// SetInput applies the change-set in the UpdateSysRoleMenuInput on the update builder.
func (u *SysRoleMenuUpdate) SetInput(i UpdateSysRoleMenuInput) *SysRoleMenuUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysRoleMenuInput on the update-one builder.
func (u *SysRoleMenuUpdateOne) SetInput(i UpdateSysRoleMenuInput) *SysRoleMenuUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysUserInput represents a mutation input for creating sysusers.
type CreateSysUserInput struct {
	IsDel     *bool
	Sort      *int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	IsActive  *bool
	UserName  string
	RealName  *string
	FirstName *string
	LastName  *string
	Password  string
	Email     string
	Phone     string
	Salt      *string
}

// Mutate applies the CreateSysUserInput on the SysUserCreate builder.
func (i *CreateSysUserInput) Mutate(m *SysUserCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetUserName(i.UserName)
	if v := i.RealName; v != nil {
		m.SetRealName(*v)
	}
	if v := i.FirstName; v != nil {
		m.SetFirstName(*v)
	}
	if v := i.LastName; v != nil {
		m.SetLastName(*v)
	}
	m.SetPassword(i.Password)
	m.SetEmail(i.Email)
	m.SetPhone(i.Phone)
	if v := i.Salt; v != nil {
		m.SetSalt(*v)
	}
}

// SetInput applies the change-set in the CreateSysUserInput on the create builder.
func (c *SysUserCreate) SetInput(i CreateSysUserInput) *SysUserCreate {
	i.Mutate(c)
	return c
}

// UpdateSysUserInput represents a mutation input for updating sysusers.
type UpdateSysUserInput struct {
	IsDel          *bool
	Sort           *int32
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	IsActive       *bool
	RealName       *string
	ClearRealName  bool
	FirstName      *string
	ClearFirstName bool
	LastName       *string
	ClearLastName  bool
	Password       *string
	Email          *string
	Phone          *string
	Salt           *string
}

// Mutate applies the UpdateSysUserInput on the SysUserMutation.
func (i *UpdateSysUserInput) Mutate(m *SysUserMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if i.ClearRealName {
		m.ClearRealName()
	}
	if v := i.RealName; v != nil {
		m.SetRealName(*v)
	}
	if i.ClearFirstName {
		m.ClearFirstName()
	}
	if v := i.FirstName; v != nil {
		m.SetFirstName(*v)
	}
	if i.ClearLastName {
		m.ClearLastName()
	}
	if v := i.LastName; v != nil {
		m.SetLastName(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Phone; v != nil {
		m.SetPhone(*v)
	}
	if v := i.Salt; v != nil {
		m.SetSalt(*v)
	}
}

// SetInput applies the change-set in the UpdateSysUserInput on the update builder.
func (u *SysUserUpdate) SetInput(i UpdateSysUserInput) *SysUserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysUserInput on the update-one builder.
func (u *SysUserUpdateOne) SetInput(i UpdateSysUserInput) *SysUserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateSysUserRoleInput represents a mutation input for creating sysuserroles.
type CreateSysUserRoleInput struct {
	IsDel     *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    string
	RoleID    string
}

// Mutate applies the CreateSysUserRoleInput on the SysUserRoleCreate builder.
func (i *CreateSysUserRoleInput) Mutate(m *SysUserRoleCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetUserID(i.UserID)
	m.SetRoleID(i.RoleID)
}

// SetInput applies the change-set in the CreateSysUserRoleInput on the create builder.
func (c *SysUserRoleCreate) SetInput(i CreateSysUserRoleInput) *SysUserRoleCreate {
	i.Mutate(c)
	return c
}

// UpdateSysUserRoleInput represents a mutation input for updating sysuserroles.
type UpdateSysUserRoleInput struct {
	IsDel          *bool
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	UserID         *string
	RoleID         *string
}

// Mutate applies the UpdateSysUserRoleInput on the SysUserRoleMutation.
func (i *UpdateSysUserRoleInput) Mutate(m *SysUserRoleMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.RoleID; v != nil {
		m.SetRoleID(*v)
	}
}

// SetInput applies the change-set in the UpdateSysUserRoleInput on the update builder.
func (u *SysUserRoleUpdate) SetInput(i UpdateSysUserRoleInput) *SysUserRoleUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSysUserRoleInput on the update-one builder.
func (u *SysUserRoleUpdateOne) SetInput(i UpdateSysUserRoleInput) *SysUserRoleUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateXxxDemoInput represents a mutation input for creating xxxdemos.
type CreateXxxDemoInput struct {
	IsDel     *bool
	Memo      *string
	Sort      *int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	IsActive  *bool
	Code      string
	Name      string
}

// Mutate applies the CreateXxxDemoInput on the XxxDemoCreate builder.
func (i *CreateXxxDemoInput) Mutate(m *XxxDemoCreate) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	m.SetCode(i.Code)
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateXxxDemoInput on the create builder.
func (c *XxxDemoCreate) SetInput(i CreateXxxDemoInput) *XxxDemoCreate {
	i.Mutate(c)
	return c
}

// UpdateXxxDemoInput represents a mutation input for updating xxxdemos.
type UpdateXxxDemoInput struct {
	IsDel          *bool
	Memo           *string
	Sort           *int32
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	ClearDeletedAt bool
	IsActive       *bool
	Code           *string
	Name           *string
}

// Mutate applies the UpdateXxxDemoInput on the XxxDemoMutation.
func (i *UpdateXxxDemoInput) Mutate(m *XxxDemoMutation) {
	if v := i.IsDel; v != nil {
		m.SetIsDel(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearDeletedAt {
		m.ClearDeletedAt()
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateXxxDemoInput on the update builder.
func (u *XxxDemoUpdate) SetInput(i UpdateXxxDemoInput) *XxxDemoUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateXxxDemoInput on the update-one builder.
func (u *XxxDemoUpdateOne) SetInput(i UpdateXxxDemoInput) *XxxDemoUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

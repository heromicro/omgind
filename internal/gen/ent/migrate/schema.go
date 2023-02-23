// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysDictsColumns holds the columns for the "sys_dicts" table.
	SysDictsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "name_cn", Type: field.TypeString, Size: 128},
		{Name: "name_en", Type: field.TypeString, Size: 128},
	}
	// SysDictsTable holds the schema information for the "sys_dicts" table.
	SysDictsTable = &schema.Table{
		Name:       "sys_dicts",
		Columns:    SysDictsColumns,
		PrimaryKey: []*schema.Column{SysDictsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysdict_id",
				Unique:  true,
				Columns: []*schema.Column{SysDictsColumns[0]},
			},
			{
				Name:    "sysdict_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[1]},
			},
			{
				Name:    "sysdict_sort",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[3]},
			},
			{
				Name:    "sysdict_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[4]},
			},
			{
				Name:    "sysdict_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[6]},
			},
			{
				Name:    "sysdict_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[7]},
			},
		},
	}
	// SysDictItemsColumns holds the columns for the "sys_dict_items" table.
	SysDictItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "label", Type: field.TypeString, Size: 128},
		{Name: "val", Type: field.TypeInt},
		{Name: "dict_id", Type: field.TypeString, Size: 36},
	}
	// SysDictItemsTable holds the schema information for the "sys_dict_items" table.
	SysDictItemsTable = &schema.Table{
		Name:       "sys_dict_items",
		Columns:    SysDictItemsColumns,
		PrimaryKey: []*schema.Column{SysDictItemsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysdictitem_id",
				Unique:  true,
				Columns: []*schema.Column{SysDictItemsColumns[0]},
			},
			{
				Name:    "sysdictitem_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[1]},
			},
			{
				Name:    "sysdictitem_sort",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[3]},
			},
			{
				Name:    "sysdictitem_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[4]},
			},
			{
				Name:    "sysdictitem_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[6]},
			},
			{
				Name:    "sysdictitem_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[7]},
			},
		},
	}
	// SysDistrictsColumns holds the columns for the "sys_districts" table.
	SysDistrictsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "tree_id", Type: field.TypeInt64, Nullable: true},
		{Name: "tree_level", Type: field.TypeInt64, Nullable: true},
		{Name: "tree_left", Type: field.TypeInt64, Nullable: true},
		{Name: "tree_right", Type: field.TypeInt64, Nullable: true},
		{Name: "is_leaf", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "t_path", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "sname", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "abbr", Type: field.TypeString, Nullable: true, Size: 16},
		{Name: "stcode", Type: field.TypeString, Nullable: true, Size: 16},
		{Name: "initials", Type: field.TypeString, Nullable: true, Size: 32},
		{Name: "pinyin", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "area_code", Type: field.TypeString, Nullable: true, Size: 8},
		{Name: "zip_code", Type: field.TypeString, Nullable: true, Size: 8},
		{Name: "mname", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "msname", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "extra", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "suffix", Type: field.TypeString, Nullable: true, Size: 16},
		{Name: "is_hot", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "is_r", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "is_m", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "is_d", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "creator", Type: field.TypeString, Size: 36},
		{Name: "pid", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// SysDistrictsTable holds the schema information for the "sys_districts" table.
	SysDistrictsTable = &schema.Table{
		Name:       "sys_districts",
		Columns:    SysDistrictsColumns,
		PrimaryKey: []*schema.Column{SysDistrictsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_districts_sys_districts_children",
				Columns:    []*schema.Column{SysDistrictsColumns[32]},
				RefColumns: []*schema.Column{SysDistrictsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sysdistrict_id",
				Unique:  true,
				Columns: []*schema.Column{SysDistrictsColumns[0]},
			},
			{
				Name:    "sysdistrict_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[1]},
			},
			{
				Name:    "sysdistrict_sort",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[2]},
			},
			{
				Name:    "sysdistrict_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[3]},
			},
			{
				Name:    "sysdistrict_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[5]},
			},
			{
				Name:    "sysdistrict_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[6]},
			},
			{
				Name:    "sysdistrict_tree_id",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[7]},
			},
			{
				Name:    "sysdistrict_tree_id_tree_left",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[7], SysDistrictsColumns[9]},
			},
			{
				Name:    "sysdistrict_tree_id_tree_right",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[7], SysDistrictsColumns[10]},
			},
			{
				Name:    "sysdistrict_tree_id_tree_left_tree_right",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[7], SysDistrictsColumns[9], SysDistrictsColumns[10]},
			},
			{
				Name:    "sysdistrict_is_hot",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[27]},
			},
			{
				Name:    "sysdistrict_is_r",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[28]},
			},
			{
				Name:    "sysdistrict_is_m",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[29]},
			},
			{
				Name:    "sysdistrict_is_d",
				Unique:  false,
				Columns: []*schema.Column{SysDistrictsColumns[30]},
			},
		},
	}
	// SysJwtBlocksColumns holds the columns for the "sys_jwt_blocks" table.
	SysJwtBlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "jwt", Type: field.TypeString, Size: 2147483647},
	}
	// SysJwtBlocksTable holds the schema information for the "sys_jwt_blocks" table.
	SysJwtBlocksTable = &schema.Table{
		Name:       "sys_jwt_blocks",
		Columns:    SysJwtBlocksColumns,
		PrimaryKey: []*schema.Column{SysJwtBlocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysjwtblock_id",
				Unique:  true,
				Columns: []*schema.Column{SysJwtBlocksColumns[0]},
			},
			{
				Name:    "sysjwtblock_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[1]},
			},
			{
				Name:    "sysjwtblock_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[3]},
			},
			{
				Name:    "sysjwtblock_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[5]},
			},
			{
				Name:    "sysjwtblock_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[6]},
			},
		},
	}
	// SysLoggingColumns holds the columns for the "sys_logging" table.
	SysLoggingColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "level", Type: field.TypeString, Size: 32},
		{Name: "trace_id", Type: field.TypeString, Size: 128},
		{Name: "user_id", Type: field.TypeString, Size: 128},
		{Name: "tag", Type: field.TypeString, Size: 128},
		{Name: "version", Type: field.TypeString, Size: 64},
		{Name: "message", Type: field.TypeString, Size: 1024},
		{Name: "data", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "error_stack", Type: field.TypeString, Size: 2147483647},
		{Name: "crtd_at", Type: field.TypeTime},
	}
	// SysLoggingTable holds the schema information for the "sys_logging" table.
	SysLoggingTable = &schema.Table{
		Name:       "sys_logging",
		Columns:    SysLoggingColumns,
		PrimaryKey: []*schema.Column{SysLoggingColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "syslogging_id",
				Unique:  true,
				Columns: []*schema.Column{SysLoggingColumns[0]},
			},
			{
				Name:    "syslogging_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingColumns[1]},
			},
			{
				Name:    "syslogging_level",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingColumns[3]},
			},
			{
				Name:    "syslogging_trace_id",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingColumns[4]},
			},
			{
				Name:    "syslogging_user_id",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingColumns[5]},
			},
			{
				Name:    "syslogging_tag",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingColumns[6]},
			},
			{
				Name:    "syslogging_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingColumns[11]},
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "icon", Type: field.TypeString, Size: 256},
		{Name: "router", Type: field.TypeString, Size: 4096},
		{Name: "is_show", Type: field.TypeBool, Default: true},
		{Name: "pid", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "ppath", Type: field.TypeString, Nullable: true, Size: 160},
		{Name: "level", Type: field.TypeInt32},
		{Name: "is_leaf", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysmenu_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenusColumns[0]},
			},
			{
				Name:    "sysmenu_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[1]},
			},
			{
				Name:    "sysmenu_sort",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[3]},
			},
			{
				Name:    "sysmenu_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[4]},
			},
			{
				Name:    "sysmenu_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[6]},
			},
			{
				Name:    "sysmenu_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[7]},
			},
			{
				Name:    "sysmenu_pid",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[12]},
			},
			{
				Name:    "sysmenu_pid_name",
				Unique:  true,
				Columns: []*schema.Column{SysMenusColumns[12], SysMenusColumns[8]},
			},
		},
	}
	// SysMenuActionsColumns holds the columns for the "sys_menu_actions" table.
	SysMenuActionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "menu_id", Type: field.TypeString, Size: 36},
		{Name: "code", Type: field.TypeString, Size: 128},
		{Name: "name", Type: field.TypeString, Size: 128},
	}
	// SysMenuActionsTable holds the schema information for the "sys_menu_actions" table.
	SysMenuActionsTable = &schema.Table{
		Name:       "sys_menu_actions",
		Columns:    SysMenuActionsColumns,
		PrimaryKey: []*schema.Column{SysMenuActionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysmenuaction_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenuActionsColumns[0]},
			},
			{
				Name:    "sysmenuaction_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[1]},
			},
			{
				Name:    "sysmenuaction_sort",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[2]},
			},
			{
				Name:    "sysmenuaction_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[3]},
			},
			{
				Name:    "sysmenuaction_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[5]},
			},
			{
				Name:    "sysmenuaction_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[7]},
			},
		},
	}
	// SysMenuActionResourcesColumns holds the columns for the "sys_menu_action_resources" table.
	SysMenuActionResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "method", Type: field.TypeString, Size: 128},
		{Name: "path", Type: field.TypeString, Size: 256},
		{Name: "action_id", Type: field.TypeString, Size: 36},
	}
	// SysMenuActionResourcesTable holds the schema information for the "sys_menu_action_resources" table.
	SysMenuActionResourcesTable = &schema.Table{
		Name:       "sys_menu_action_resources",
		Columns:    SysMenuActionResourcesColumns,
		PrimaryKey: []*schema.Column{SysMenuActionResourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysmenuactionresource_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[0]},
			},
			{
				Name:    "sysmenuactionresource_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[1]},
			},
			{
				Name:    "sysmenuactionresource_sort",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[2]},
			},
			{
				Name:    "sysmenuactionresource_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[4]},
			},
			{
				Name:    "sysmenuactionresource_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[6]},
			},
			{
				Name:    "sysmenuactionresource_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[7]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 64},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysrole_id",
				Unique:  true,
				Columns: []*schema.Column{SysRolesColumns[0]},
			},
			{
				Name:    "sysrole_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[1]},
			},
			{
				Name:    "sysrole_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[2]},
			},
			{
				Name:    "sysrole_sort",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[3]},
			},
			{
				Name:    "sysrole_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[5]},
			},
			{
				Name:    "sysrole_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[7]},
			},
		},
	}
	// SysRoleMenusColumns holds the columns for the "sys_role_menus" table.
	SysRoleMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "role_id", Type: field.TypeString, Size: 36},
		{Name: "menu_id", Type: field.TypeString, Size: 36},
		{Name: "action_id", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// SysRoleMenusTable holds the schema information for the "sys_role_menus" table.
	SysRoleMenusTable = &schema.Table{
		Name:       "sys_role_menus",
		Columns:    SysRoleMenusColumns,
		PrimaryKey: []*schema.Column{SysRoleMenusColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysrolemenu_id",
				Unique:  true,
				Columns: []*schema.Column{SysRoleMenusColumns[0]},
			},
			{
				Name:    "sysrolemenu_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[1]},
			},
			{
				Name:    "sysrolemenu_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[2]},
			},
			{
				Name:    "sysrolemenu_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[4]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "user_name", Type: field.TypeString, Size: 128},
		{Name: "real_name", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 31},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 31},
		{Name: "passwd", Type: field.TypeString, Size: 256},
		{Name: "email", Type: field.TypeString, Size: 256},
		{Name: "mobile", Type: field.TypeString, Size: 20},
		{Name: "salt", Type: field.TypeString},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysuser_id",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[0]},
			},
			{
				Name:    "sysuser_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[1]},
			},
			{
				Name:    "sysuser_sort",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[2]},
			},
			{
				Name:    "sysuser_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[3]},
			},
			{
				Name:    "sysuser_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[5]},
			},
			{
				Name:    "sysuser_is_active",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[6]},
			},
			{
				Name:    "sysuser_user_name",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[7]},
			},
		},
	}
	// SysUserRolesColumns holds the columns for the "sys_user_roles" table.
	SysUserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Size: 36},
		{Name: "role_id", Type: field.TypeString, Size: 36},
	}
	// SysUserRolesTable holds the schema information for the "sys_user_roles" table.
	SysUserRolesTable = &schema.Table{
		Name:       "sys_user_roles",
		Columns:    SysUserRolesColumns,
		PrimaryKey: []*schema.Column{SysUserRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysuserrole_id",
				Unique:  true,
				Columns: []*schema.Column{SysUserRolesColumns[0]},
			},
			{
				Name:    "sysuserrole_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[1]},
			},
			{
				Name:    "sysuserrole_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[2]},
			},
			{
				Name:    "sysuserrole_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[4]},
			},
		},
	}
	// XxxDemosColumns holds the columns for the "xxx_demos" table.
	XxxDemosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "code", Type: field.TypeString, Size: 128},
		{Name: "name", Type: field.TypeString, Size: 128},
	}
	// XxxDemosTable holds the schema information for the "xxx_demos" table.
	XxxDemosTable = &schema.Table{
		Name:       "xxx_demos",
		Columns:    XxxDemosColumns,
		PrimaryKey: []*schema.Column{XxxDemosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "xxxdemo_id",
				Unique:  true,
				Columns: []*schema.Column{XxxDemosColumns[0]},
			},
			{
				Name:    "xxxdemo_is_del",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[1]},
			},
			{
				Name:    "xxxdemo_sort",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[3]},
			},
			{
				Name:    "xxxdemo_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[4]},
			},
			{
				Name:    "xxxdemo_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[6]},
			},
			{
				Name:    "xxxdemo_is_active",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysDictsTable,
		SysDictItemsTable,
		SysDistrictsTable,
		SysJwtBlocksTable,
		SysLoggingTable,
		SysMenusTable,
		SysMenuActionsTable,
		SysMenuActionResourcesTable,
		SysRolesTable,
		SysRoleMenusTable,
		SysUsersTable,
		SysUserRolesTable,
		XxxDemosTable,
	}
)

func init() {
	SysDistrictsTable.ForeignKeys[0].RefTable = SysDistrictsTable
	SysLoggingTable.Annotation = &entsql.Annotation{
		Table: "sys_logging",
	}
}

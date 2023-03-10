// Code generated by ent, DO NOT EDIT.

package ent

import "entgo.io/ent/dialect"

func (c *SysAddressClient) Debug() *SysAddressClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysAddressClient{config: cfg}
}

func (c *SysDictClient) Debug() *SysDictClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysDictClient{config: cfg}
}

func (c *SysDictItemClient) Debug() *SysDictItemClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysDictItemClient{config: cfg}
}

func (c *SysDistrictClient) Debug() *SysDistrictClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysDistrictClient{config: cfg}
}

func (c *SysJwtBlockClient) Debug() *SysJwtBlockClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysJwtBlockClient{config: cfg}
}

func (c *SysLoggingClient) Debug() *SysLoggingClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysLoggingClient{config: cfg}
}

func (c *SysMenuClient) Debug() *SysMenuClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysMenuClient{config: cfg}
}

func (c *SysMenuActionClient) Debug() *SysMenuActionClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysMenuActionClient{config: cfg}
}

func (c *SysMenuActionResourceClient) Debug() *SysMenuActionResourceClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysMenuActionResourceClient{config: cfg}
}

func (c *SysRoleClient) Debug() *SysRoleClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysRoleClient{config: cfg}
}

func (c *SysRoleMenuClient) Debug() *SysRoleMenuClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysRoleMenuClient{config: cfg}
}

func (c *SysUserClient) Debug() *SysUserClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysUserClient{config: cfg}
}

func (c *SysUserRoleClient) Debug() *SysUserRoleClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &SysUserRoleClient{config: cfg}
}

func (c *XxxDemoClient) Debug() *XxxDemoClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &XxxDemoClient{config: cfg}
}

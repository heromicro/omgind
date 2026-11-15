# Tenant RBAC

## Request flow

Authentication runs before tenant validation and Casbin. An invalid token is rejected in debug mode as well as production. Login and captcha routes are public. Authenticated accounts may list their organizations and create an organization before selecting a tenant. Organization-dependent requests require `X-Org-ID`.

`GET /api/v2/pub/current/organs` returns active organizations linked through active `OrgStaff` records (`user_id`, `org_id`, `empy_stat = 1`). Root can select any active organization but must also supply a tenant for tenant operations. Login enters the application without requiring an organization. Organization selection, switching and creation are optional account actions; organization-dependent operations remain forbidden until a valid organization is selected. Requests carry the selected organization, and logout clears it and the menu cache.

Creating an organization overrides any client-supplied owner ID, links the authenticated creator as a staff member, creates the initial department, and scopes the headquarters address to that organization. The owner receives administrative API permissions only in that organization through the policy adapter.

## Policy model

Requests: `(user_id, org_id, path, method)`.
Policies: `(role_id, org_id, path, method)`.
Assignments: `(user_id, role_id, org_id)`.

Role permissions are read from the existing role-menu-action-resource relationships. Roles are scoped by `SysRole.org_id`; assignments are scoped by `SysUserRole.org_id`. An explicit API path and HTTP method are required. No matching permission means denial.

`PUT /api/v2/org-staffs/:id/roles` accepts `{ "role_ids": ["..."] }` and replaces assignments only in the selected organization. Each role and staff member must belong to that organization. An empty list revokes all tenant role assignments. Grant this endpoint through a menu action resource when delegating role management.

Successful user, role and menu mutations reload policies synchronously after database work completes. A refresh failure clears cached policies to fail closed. Organization creation also reloads policy. Periodic reload is optional. Multiple application instances still need a shared invalidation notification or periodic reload; synchronous reload updates only the current instance.

## Data isolation

Ent query interceptors and mutation hooks apply the verified tenant to organization, staff, department, position, address, role and user-role tables. Role-menu records are filtered through their owning role. Cross-tenant department, position, parent, address and role references are rejected. Transaction clients inherit guards. Global account results are restricted to tenant members. Global-account mutations and global-menu changes are denied to tenant users; a verified platform administrator may manage global records without bypassing organization row filters.

Unscoped contexts are reserved for trusted bootstrap and policy loading. Never construct an unscoped context from a tenant request. Queries for legacy logging, JWT-blocklist, team, team-user and attachment entities currently fail closed in tenant contexts until those entities have an explicit tenant ownership model. Their writes are also denied in tenant contexts.

## Existing database migration

Back up the database and migrate schema before enabling tenant RBAC. New fields are `org_staffs.user_id` (nullable), `sys_roles.org_id` and `sys_user_roles.org_id` (both default empty for existing platform data). With automatic Ent migrations disabled, apply these additive changes through the deployment migration process.

Existing roles are not assigned to an arbitrary organization automatically. Explicitly map existing staff to global user IDs, map each role and user-role assignment to the intended organization, and verify that role assignment and role have the same organization. Copy roles when independent permissions are required in different organizations. Do not treat an empty org_id as a wildcard. Review organization ownership and all menu-resource API paths before enabling access for existing users.

## Verification

Run `go test ./internal/scheme ./internal/app/middleware ./internal/app/service ./internal/router ./internal/wirex ./internal/scheme/entity/gorm`. Tests cover role domains, HTTP methods, denied users, revocation, tenant filters, cross-tenant writes and references, and transaction isolation. Live production database migration and browser verification require the running application environment.

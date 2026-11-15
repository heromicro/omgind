package contextx

import "context"

type tenantKey struct{}

// WithTenant scopes database operations to a verified organization.
func WithTenant(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, tenantKey{}, id)
}
func Tenant(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(tenantKey{}).(string)
	return id, ok
}

type platformAdminKey struct{}

// WithPlatformAdmin marks a verified platform administrator without bypassing tenant row filters.
func WithPlatformAdmin(ctx context.Context) context.Context {
	return context.WithValue(ctx, platformAdminKey{}, true)
}
func PlatformAdmin(ctx context.Context) bool {
	value, _ := ctx.Value(platformAdminKey{}).(bool)
	return value
}

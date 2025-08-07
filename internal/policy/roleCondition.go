package policy

import (
	"context"

	"github.com/MalcolmFuchs/GoGuard/internal/common"
)

type RoleCondition struct {
	RequiredRoles []string
	RoleExtractor func(subject any) ([]string, error)
}

func (c *RoleCondition) Evaluate(ctx context.Context, subject any, resource any) (bool, error) {
	if c.RoleExtractor == nil {
		return false, common.ErrConditionFailed
	}

	roles, err := c.RoleExtractor(subject)
	if err != nil {
		return false, err
	}

	for _, requiredRole := range c.RequiredRoles {
		for _, role := range roles {
			if role == requiredRole {
				return true, nil
			}
		}
	}

	return false, nil
}

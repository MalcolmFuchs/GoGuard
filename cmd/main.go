package main

import (
	"context"
	"fmt"

	"github.com/MalcolmFuchs/GoGuard/internal/policy"
	"github.com/MalcolmFuchs/GoGuard/internal/user"
)

func main() {
	ctx := context.Background()

	alice := &user.User{
		ID:    "alice",
		Name:  "Alice",
		Roles: []user.Role{"admin"},
	}

	repo := user.NewInMemoryRepository([]*user.User{alice})

	u, err := repo.GetByID("alice")
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}

	extractor := func(subject any) ([]string, error) {
		userObj, ok := subject.(*user.User)
		if !ok {
			return nil, fmt.Errorf("unexpected subject type")
		}
		roles := []string{}
		for _, r := range userObj.Roles {
			roles = append(roles, string(r))
		}
		return roles, nil
	}

	cond := &policy.RoleCondition{
		RequiredRoles: []string{"admin"},
		RoleExtractor: extractor,
	}

	allowed, err := cond.Evaluate(ctx, u, nil)
	if err != nil {
		fmt.Println("Evaluation Error:", err)
		return
	}

	if allowed {
		fmt.Println("Zugriff erlaubt für", u.Name)
	} else {
		fmt.Println("Zugriff verweigert für", u.Name)
	}
}

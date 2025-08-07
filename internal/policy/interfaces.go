package policy

type EvaluationResult struct {
	Allowed bool
	Reason  string
}

type Condition interface {
	Evaluate(userID string, subject any, resource any) (bool, error)
}

type Policy interface {
	Evaluate(userID string, subject any, resource any) (EvaluationResult, error)
}



package grammar

type validationExpr struct {
}

// SelectionExpr conditional expression with if/else (optional)
type SelectionExpr struct {
	conditionExpr validationExpr
	ifExpr        validationExpr
	elseExpr      validationExpr
}

type logicalExpr struct {
	validationExpr

	isOrAnd    bool
	conditions []validationExpr
}

type relationalOperator uint8

const (
	eq relationalOperator = iota
	ne
	lt
	gt
	le
	ge
)

type relationalExpr struct {
	validationExpr

	operator relationalOperator
	left     numericExpr
	right    numericExpr
}

type numericOperator uint8

const (
	plus relationalOperator = iota
	minus
	mult
	div
)

type numericExpr struct {
	operator numericOperator
	left     unaryExpr
	right    unaryExpr
}

type unaryOperator uint8

const (
	not unaryOperator = iota
)

type unaryExpr struct {
	operator unaryOperator

	variableValue     variable
	literalValue      literal
	functionCallValue functionCall
}

type literal struct{}

type timeDurationLiteral struct {
	literal

	unit     int
	duration int
}

type rangeDurationLiteral struct {
	literal

	from int
	to   int
}

type versionNumberLiteral struct {
	literal

	values []int
}

type variable struct {
	indexOrName string

	localOrContext bool
}

type functionCall struct {
	funcName       string
	funcAttributes []functionAttribute
}

type functionAttribute struct {
	literalValue  literal
	variableValue variable
}

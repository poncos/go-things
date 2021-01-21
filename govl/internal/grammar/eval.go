package grammar

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/ecollado/go-regex-eval/internal/grammar/parser"
)

type validationContext struct {
}

type validationListener struct {
	*parser.BaseValidationExprListener

	//patternContext   model.RegexMatchContext
	//conditionContext model.ConditionContext
	//stack            []string
	SelectionExpr expression 
}

// Evaluate evaluates the given expression
func Evaluate(expression string) bool {

	if expression == "" {
		return true
	}

	log.Printf("Evaluating expression %s\n", expression)

	is := antlr.NewInputStream(expression)

	// Create the Lexer
	lexer := parser.NewValidationExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewValidationExprParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener validationListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.SelectionExpr())

	return true
}

// EnterEveryRule is called when any rule is entered.
func (s *validationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *validationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSelectionExpr is called when production selectionExpr is entered.
func (s *validationListener) EnterSelectionExpr(ctx *parser.SelectionExprContext) {
	log.Printf("EnterSelectionExpr %s\n", ctx.GetText())
}

// ExitSelectionExpr is called when production selectionExpr is exited.
func (s *validationListener) ExitSelectionExpr(ctx *parser.SelectionExprContext) {
	log.Printf("ExitSelectionExpr %s\n", ctx.GetText())
}

// EnterValidationExpr is called when production validationExpr is entered.
func (s *validationListener) EnterValidationExpr(ctx *parser.ValidationExprContext) {
	log.Printf("EnterValidationExpr %s\n", ctx.GetText())
}

// ExitValidationExpr is called when production validationExpr is exited.
func (s *validationListener) ExitValidationExpr(ctx *parser.ValidationExprContext) {
	log.Printf("ExitValidationExpr %s\n", ctx.GetText())
}

// EnterValidationOrExpr is called when production validationOrExpr is entered.
func (s *validationListener) EnterValidationOrExpr(ctx *parser.ValidationOrExprContext) {
	log.Printf("EnterValidationOrExpr %s\n", ctx.GetText())
}

// ExitValidationOrExpr is called when production validationOrExpr is exited.
func (s *validationListener) ExitValidationOrExpr(ctx *parser.ValidationOrExprContext) {
	log.Printf("ExitValidationOrExpr %s\n", ctx.GetText())
}

// EnterValidationAndExpr is called when production validationAndExpr is entered.
func (s *validationListener) EnterValidationAndExpr(ctx *parser.ValidationAndExprContext) {
	log.Printf("EnterValidationAndExpr %s\n", ctx.GetText())
}

// ExitValidationAndExpr is called when production validationAndExpr is exited.
func (s *validationListener) ExitValidationAndExpr(ctx *parser.ValidationAndExprContext) {
	log.Printf("ExitValidationAndExpr %s\n", ctx.GetText())
}

// EnterValueLogical is called when production valueLogical is entered.
func (s *validationListener) EnterValueLogical(ctx *parser.ValueLogicalContext) {
	log.Printf("EnterValueLogical %s\n", ctx.GetText())
}

// ExitValueLogical is called when production valueLogical is exited.
func (s *validationListener) ExitValueLogical(ctx *parser.ValueLogicalContext) {
	log.Printf("ExitValueLogical %s\n", ctx.GetText())
}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *validationListener) EnterRelationalExpression(ctx *parser.RelationalExpressionContext) {
	log.Printf("EnterRelationalExpression %s\n", ctx.GetText())
}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *validationListener) ExitRelationalExpression(ctx *parser.RelationalExpressionContext) {
	log.Printf("ExitRelationalExpression %s\n", ctx.GetText())
}

// EnterNumericExpression is called when production numericExpression is entered.
func (s *validationListener) EnterNumericExpression(ctx *parser.NumericExpressionContext) {
	log.Printf("EnterNumericExpression %s\n", ctx.GetText())
}

// ExitNumericExpression is called when production numericExpression is exited.
func (s *validationListener) ExitNumericExpression(ctx *parser.NumericExpressionContext) {
	log.Printf("ExitNumericExpression %s\n", ctx.GetText())
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *validationListener) EnterAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	log.Printf("EnterAdditiveExpression %s\n", ctx.GetText())
}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *validationListener) ExitAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	log.Printf("ExitAdditiveExpression %s\n", ctx.GetText())
}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *validationListener) EnterMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {
	log.Printf("EnterMultiplicativeExpression %s\n", ctx.GetText())
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *validationListener) ExitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {
	log.Printf("ExitMultiplicativeExpression %s\n", ctx.GetText())
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *validationListener) EnterUnaryExpression(ctx *parser.UnaryExpressionContext) {
	log.Printf("EnterUnaryExpression %s\n", ctx.GetText())

}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *validationListener) ExitUnaryExpression(ctx *parser.UnaryExpressionContext) {
	log.Printf("ExitUnaryExpression %s\n", ctx.GetText())
}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *validationListener) EnterUnaryOperator(ctx *parser.UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *validationListener) ExitUnaryOperator(ctx *parser.UnaryOperatorContext) {
	log.Printf("ExitUnaryOperator %s\n", ctx.GetText())
}

// EnterBrackettedExpression is called when production brackettedExpression is entered.
func (s *validationListener) EnterBrackettedExpression(ctx *parser.BrackettedExpressionContext) {
}

// ExitBrackettedExpression is called when production brackettedExpression is exited.
func (s *validationListener) ExitBrackettedExpression(ctx *parser.BrackettedExpressionContext) {
	log.Printf("ExitBrackettedExpression %s\n", ctx.GetText())
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *validationListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *validationListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	log.Printf("ExitFunctionCall %s\n", ctx.GetText())
}

// EnterFunctionArgumentList is called when production functionArgumentList is entered.
func (s *validationListener) EnterFunctionArgumentList(ctx *parser.FunctionArgumentListContext) {}

// ExitFunctionArgumentList is called when production functionArgumentList is exited.
func (s *validationListener) ExitFunctionArgumentList(ctx *parser.FunctionArgumentListContext) {
	log.Printf("ExitFunctionArgumentList %s\n", ctx.GetText())
}

// EnterFunctionAttribute is called when production functionAttribute is entered.
func (s *validationListener) EnterFunctionAttribute(ctx *parser.FunctionAttributeContext) {}

// ExitFunctionAttribute is called when production functionAttribute is exited.
func (s *validationListener) ExitFunctionAttribute(ctx *parser.FunctionAttributeContext) {
	log.Printf("ExitFunctionAttribute %s\n", ctx.GetText())
}

// EnterLiteral is called when production literal is entered.
func (s *validationListener) EnterLiteral(ctx *parser.LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *validationListener) ExitLiteral(ctx *parser.LiteralContext) {
	log.Printf("ExitLiteral %s\n", ctx.GetText())
}

// EnterBoolenLiteral is called when production boolenLiteral is entered.
func (s *validationListener) EnterBoolenLiteral(ctx *parser.BoolenLiteralContext) {}

// ExitBoolenLiteral is called when production boolenLiteral is exited.
func (s *validationListener) ExitBoolenLiteral(ctx *parser.BoolenLiteralContext) {
	log.Printf("ExitBoolenLiteral %s\n", ctx.GetText())
}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *validationListener) EnterNumericLiteral(ctx *parser.NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *validationListener) ExitNumericLiteral(ctx *parser.NumericLiteralContext) {
	log.Printf("ExitNumericLiteral %s\n", ctx.GetText())
}

// EnterTimeDurationLiteral is called when production timeDurationLiteral is entered.
func (s *validationListener) EnterTimeDurationLiteral(ctx *parser.TimeDurationLiteralContext) {}

// ExitTimeDurationLiteral is called when production timeDurationLiteral is exited.
func (s *validationListener) ExitTimeDurationLiteral(ctx *parser.TimeDurationLiteralContext) {
	log.Printf("ExitTimeDurationLiteral %s\n", ctx.GetText())
}

// EnterVersionNumberLiteral is called when production versionNumberLiteral is entered.
func (s *validationListener) EnterVersionNumberLiteral(ctx *parser.VersionNumberLiteralContext) {}

// ExitVersionNumberLiteral is called when production versionNumberLiteral is exited.
func (s *validationListener) ExitVersionNumberLiteral(ctx *parser.VersionNumberLiteralContext) {
	log.Printf("ExitVersionNumberLiteral %s\n", ctx.GetText())
}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *validationListener) EnterArrayLiteral(ctx *parser.ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *validationListener) ExitArrayLiteral(ctx *parser.ArrayLiteralContext) {
	log.Printf("ExitArrayLiteral %s\n", ctx.GetText())
}

// EnterRangeLiteral is called when production rangeLiteral is entered.
func (s *validationListener) EnterRangeLiteral(ctx *parser.RangeLiteralContext) {}

// ExitRangeLiteral is called when production rangeLiteral is exited.
func (s *validationListener) ExitRangeLiteral(ctx *parser.RangeLiteralContext) {
	log.Printf("ExitRangeLiteral %s\n", ctx.GetText())
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *validationListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *validationListener) ExitStringLiteral(ctx *parser.StringLiteralContext) {
	log.Printf("ExitStringLiteral %s\n", ctx.GetText())
}

// EnterVariable is called when production variable is entered.
func (s *validationListener) EnterVariable(ctx *parser.VariableContext) {
}

// ExitVariable is called when production variable is exited.
func (s *validationListener) ExitVariable(ctx *parser.VariableContext) {
	log.Printf("ExitVariable %s\n", ctx.GetText())
}

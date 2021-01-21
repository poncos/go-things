// Code generated from ValidationExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ValidationExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseValidationExprListener is a complete listener for a parse tree produced by ValidationExprParser.
type BaseValidationExprListener struct{}

var _ ValidationExprListener = &BaseValidationExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseValidationExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseValidationExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseValidationExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseValidationExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSelectionExpr is called when production selectionExpr is entered.
func (s *BaseValidationExprListener) EnterSelectionExpr(ctx *SelectionExprContext) {}

// ExitSelectionExpr is called when production selectionExpr is exited.
func (s *BaseValidationExprListener) ExitSelectionExpr(ctx *SelectionExprContext) {}

// EnterValidationExpr is called when production validationExpr is entered.
func (s *BaseValidationExprListener) EnterValidationExpr(ctx *ValidationExprContext) {}

// ExitValidationExpr is called when production validationExpr is exited.
func (s *BaseValidationExprListener) ExitValidationExpr(ctx *ValidationExprContext) {}

// EnterValidationOrExpr is called when production validationOrExpr is entered.
func (s *BaseValidationExprListener) EnterValidationOrExpr(ctx *ValidationOrExprContext) {}

// ExitValidationOrExpr is called when production validationOrExpr is exited.
func (s *BaseValidationExprListener) ExitValidationOrExpr(ctx *ValidationOrExprContext) {}

// EnterValidationAndExpr is called when production validationAndExpr is entered.
func (s *BaseValidationExprListener) EnterValidationAndExpr(ctx *ValidationAndExprContext) {}

// ExitValidationAndExpr is called when production validationAndExpr is exited.
func (s *BaseValidationExprListener) ExitValidationAndExpr(ctx *ValidationAndExprContext) {}

// EnterValueLogical is called when production valueLogical is entered.
func (s *BaseValidationExprListener) EnterValueLogical(ctx *ValueLogicalContext) {}

// ExitValueLogical is called when production valueLogical is exited.
func (s *BaseValidationExprListener) ExitValueLogical(ctx *ValueLogicalContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseValidationExprListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseValidationExprListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterNumericExpression is called when production numericExpression is entered.
func (s *BaseValidationExprListener) EnterNumericExpression(ctx *NumericExpressionContext) {}

// ExitNumericExpression is called when production numericExpression is exited.
func (s *BaseValidationExprListener) ExitNumericExpression(ctx *NumericExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseValidationExprListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseValidationExprListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseValidationExprListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseValidationExprListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseValidationExprListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseValidationExprListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseValidationExprListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseValidationExprListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterBrackettedExpression is called when production brackettedExpression is entered.
func (s *BaseValidationExprListener) EnterBrackettedExpression(ctx *BrackettedExpressionContext) {}

// ExitBrackettedExpression is called when production brackettedExpression is exited.
func (s *BaseValidationExprListener) ExitBrackettedExpression(ctx *BrackettedExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseValidationExprListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseValidationExprListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionArgumentList is called when production functionArgumentList is entered.
func (s *BaseValidationExprListener) EnterFunctionArgumentList(ctx *FunctionArgumentListContext) {}

// ExitFunctionArgumentList is called when production functionArgumentList is exited.
func (s *BaseValidationExprListener) ExitFunctionArgumentList(ctx *FunctionArgumentListContext) {}

// EnterFunctionAttribute is called when production functionAttribute is entered.
func (s *BaseValidationExprListener) EnterFunctionAttribute(ctx *FunctionAttributeContext) {}

// ExitFunctionAttribute is called when production functionAttribute is exited.
func (s *BaseValidationExprListener) ExitFunctionAttribute(ctx *FunctionAttributeContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseValidationExprListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseValidationExprListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBoolenLiteral is called when production boolenLiteral is entered.
func (s *BaseValidationExprListener) EnterBoolenLiteral(ctx *BoolenLiteralContext) {}

// ExitBoolenLiteral is called when production boolenLiteral is exited.
func (s *BaseValidationExprListener) ExitBoolenLiteral(ctx *BoolenLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseValidationExprListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseValidationExprListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterTimeDurationLiteral is called when production timeDurationLiteral is entered.
func (s *BaseValidationExprListener) EnterTimeDurationLiteral(ctx *TimeDurationLiteralContext) {}

// ExitTimeDurationLiteral is called when production timeDurationLiteral is exited.
func (s *BaseValidationExprListener) ExitTimeDurationLiteral(ctx *TimeDurationLiteralContext) {}

// EnterVersionNumberLiteral is called when production versionNumberLiteral is entered.
func (s *BaseValidationExprListener) EnterVersionNumberLiteral(ctx *VersionNumberLiteralContext) {}

// ExitVersionNumberLiteral is called when production versionNumberLiteral is exited.
func (s *BaseValidationExprListener) ExitVersionNumberLiteral(ctx *VersionNumberLiteralContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseValidationExprListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseValidationExprListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterRangeLiteral is called when production rangeLiteral is entered.
func (s *BaseValidationExprListener) EnterRangeLiteral(ctx *RangeLiteralContext) {}

// ExitRangeLiteral is called when production rangeLiteral is exited.
func (s *BaseValidationExprListener) ExitRangeLiteral(ctx *RangeLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseValidationExprListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseValidationExprListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseValidationExprListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseValidationExprListener) ExitVariable(ctx *VariableContext) {}

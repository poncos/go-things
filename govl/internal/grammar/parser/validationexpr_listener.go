// Code generated from ValidationExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ValidationExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ValidationExprListener is a complete listener for a parse tree produced by ValidationExprParser.
type ValidationExprListener interface {
	antlr.ParseTreeListener

	// EnterSelectionExpr is called when entering the selectionExpr production.
	EnterSelectionExpr(c *SelectionExprContext)

	// EnterValidationExpr is called when entering the validationExpr production.
	EnterValidationExpr(c *ValidationExprContext)

	// EnterValidationOrExpr is called when entering the validationOrExpr production.
	EnterValidationOrExpr(c *ValidationOrExprContext)

	// EnterValidationAndExpr is called when entering the validationAndExpr production.
	EnterValidationAndExpr(c *ValidationAndExprContext)

	// EnterValueLogical is called when entering the valueLogical production.
	EnterValueLogical(c *ValueLogicalContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterNumericExpression is called when entering the numericExpression production.
	EnterNumericExpression(c *NumericExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterBrackettedExpression is called when entering the brackettedExpression production.
	EnterBrackettedExpression(c *BrackettedExpressionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionArgumentList is called when entering the functionArgumentList production.
	EnterFunctionArgumentList(c *FunctionArgumentListContext)

	// EnterFunctionAttribute is called when entering the functionAttribute production.
	EnterFunctionAttribute(c *FunctionAttributeContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBoolenLiteral is called when entering the boolenLiteral production.
	EnterBoolenLiteral(c *BoolenLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterTimeDurationLiteral is called when entering the timeDurationLiteral production.
	EnterTimeDurationLiteral(c *TimeDurationLiteralContext)

	// EnterVersionNumberLiteral is called when entering the versionNumberLiteral production.
	EnterVersionNumberLiteral(c *VersionNumberLiteralContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterRangeLiteral is called when entering the rangeLiteral production.
	EnterRangeLiteral(c *RangeLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// ExitSelectionExpr is called when exiting the selectionExpr production.
	ExitSelectionExpr(c *SelectionExprContext)

	// ExitValidationExpr is called when exiting the validationExpr production.
	ExitValidationExpr(c *ValidationExprContext)

	// ExitValidationOrExpr is called when exiting the validationOrExpr production.
	ExitValidationOrExpr(c *ValidationOrExprContext)

	// ExitValidationAndExpr is called when exiting the validationAndExpr production.
	ExitValidationAndExpr(c *ValidationAndExprContext)

	// ExitValueLogical is called when exiting the valueLogical production.
	ExitValueLogical(c *ValueLogicalContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitNumericExpression is called when exiting the numericExpression production.
	ExitNumericExpression(c *NumericExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitBrackettedExpression is called when exiting the brackettedExpression production.
	ExitBrackettedExpression(c *BrackettedExpressionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionArgumentList is called when exiting the functionArgumentList production.
	ExitFunctionArgumentList(c *FunctionArgumentListContext)

	// ExitFunctionAttribute is called when exiting the functionAttribute production.
	ExitFunctionAttribute(c *FunctionAttributeContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBoolenLiteral is called when exiting the boolenLiteral production.
	ExitBoolenLiteral(c *BoolenLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitTimeDurationLiteral is called when exiting the timeDurationLiteral production.
	ExitTimeDurationLiteral(c *TimeDurationLiteralContext)

	// ExitVersionNumberLiteral is called when exiting the versionNumberLiteral production.
	ExitVersionNumberLiteral(c *VersionNumberLiteralContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitRangeLiteral is called when exiting the rangeLiteral production.
	ExitRangeLiteral(c *RangeLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)
}

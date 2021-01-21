/*
 * Grammar defining a validation language.
 *
 * Pre:
 *   - Regex: "amiga.common.grpc.client.myclient.deadline:\s+(\w+)"
 *   - Regex Match Result:  [amiga.common.grpc.client.myclient.deadline: 38] [38]
 * Input: 
 *   1. Given the Regex Match Result
 *   2. Given a Context with optional extra values: [40]
 *   3. Given a validation expression: $1 <= 40, $1 <= @1
 * Output:
 *   - True | False with the validation result
 * Samples:
 *   - Conditionals: if  : $3 < 10 else $4 range 100,500
 *   - Braces: $1 == 1 and  ( $2 == 1 or $3 == 1)
 *   - Version number: $1 >= 3.14.1
 *   - Arrays: contains(localhost, [remote.host.com,remote2.host.com,localhost,remote3.localhost.com])
 *   - Strings:localhost in $1
*/
grammar ValidationExpr;


selectionExpr
    : 'if' validationExpr ':' validationExpr 'else' validationExpr EOF
    | validationExpr EOF
    ;

validationExpr :
    (validationOrExpr | validationAndExpr )
    ;

validationOrExpr:
    validationAndExpr ( 'or' validationAndExpr )*
    ;

validationAndExpr:
    valueLogical ( 'and' valueLogical )*
    ;

valueLogical:
    relationalExpression
    ;

relationalExpression
    : numericExpression ( '=' numericExpression | '!=' numericExpression | '<' numericExpression | '>' numericExpression | '<=' numericExpression | '>=' numericExpression )*
    ;

numericExpression
    : additiveExpression
    ;

additiveExpression
    : multiplicativeExpression ( '+' multiplicativeExpression | '-' multiplicativeExpression )*
    ;

multiplicativeExpression
    : unaryExpression ( '*' unaryExpression | '/' unaryExpression )*
    ;

unaryExpression
    :
    brackettedExpression 
    | functionCall
    | variable 
    | literal
    | unaryOperator unaryExpression
    ;

unaryOperator
    : 'not'
    ;

brackettedExpression
    : '(' validationExpr ')'
    ;

functionCall
    : STR  functionArgumentList
    ;

functionArgumentList
    : NIL 
    | '(' functionAttribute (',' functionAttribute)* ')'
    ;

functionAttribute
    : variable
    | literal
    ;

literal:
    | boolenLiteral
    | numericLiteral
    | stringLiteral
    | arrayLiteral
    | rangeLiteral
    ;

boolenLiteral
    : 'true'
    | 'false'
    ;

numericLiteral
    : INTEGER
    | timeDurationLiteral; 

timeDurationLiteral:
    INTEGER + ('ms'|'s')
    ;

versionNumberLiteral
    : INTEGER + '.' INTEGER + ('.' INTEGER)+
    ;

arrayLiteral
    : NONE
    | '[' STR + (',' + STR)+ ']'
    ;

rangeLiteral
    : (numericLiteral + (','+ numericLiteral))+
    ;

stringLiteral
    : STR
    ;

variable
    : LOCAL_VARIABLE
    | CONTEXT_VARIABLE
    ;
     
LOCAL_VARIABLE
    : '$' + STR | ('${' + STR + '}')
    ;

CONTEXT_VARIABLE
    : '@' + STR | '@{' + STR +  '}'
    ;

INTEGER : [0-9]+
     ;

STR
    : ('"'|'\'')? [a-zA-Z0-9_.]+ ('"'|'\'')?
    | ('"'|'\'') ('"'|'\'')
   ;

NONE
    : '[' WS* ']'
    ;

NIL
    : '(' WS* ')'
    ;

WS
    : (' '| '\t'| '\n'| '\r')+ ->skip
    ;

// ==========================================
// LEXER
// ==========================================

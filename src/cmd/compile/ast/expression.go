package ast



const (
	EXPRESSION_TYPE_BOOL = iota
	EXPRESSION_TYPE_INT
	EXPRESSION_TYPE_FLOAT
	EXPRESSION_TYPE_STRING
	EXPRESSION_TYPE_ARRAY
	EXPRESSION_TYPE_LOGICAL_OR
	EXPRESSION_TYPE_LOGICAL_AND
	EXPRESSION_TYPE_ASSIGN
	EXPRESSION_TYPE_COLON_ASSIGN
	EXPRESSION_TYPE_PLUS_ASSIGN
	EXPRESSION_TYPE_MINUS_ASSIGN
	EXPRESSION_TYPE_MUL_ASSIGN
	EXPRESSION_TYPE_DIV_ASSIGN
	EXPRESSION_TYPE_MOD_ASSIGN
	EXPRESSION_TYPE_ASSIGN_FUNCTION
	EXPRESSION_TYPE_FUNCTION
	EXPRESSION_TYPE_EQ
	EXPRESSION_TYPE_NE
	EXPRESSION_TYPE_GE
	EXPRESSION_TYPE_GT
	EXPRESSION_TYPE_LE
	EXPRESSION_TYPE_LT
	EXPRESSION_TYPE_ADD
	EXPRESSION_TYPE_SUB
	EXPRESSION_TYPE_MUL
	EXPRESSION_TYPE_DIV
	EXPRESSION_TYPE_MOD
	EXPRESSION_TYPE_INDEX
	EXPRESSION_TYPE_METHOD_CALL
	EXPRESSION_TYPE_FUNCTION_CALL
	EXPRESSION_TYPE_EXPRESSION_FUNCTION_CALL
	EXPRESSION_TYPE_INCREMENT
	EXPRESSION_TYPE_PRE_INCREMENT
	EXPRESSION_TYPE_PRE_DECREMENT
	EXPRESSION_TYPE_DECREMENT
	EXPRESSION_TYPE_NEGATIVE
	EXPRESSION_TYPE_NOT
	EXPRESSION_TYPE_IDENTIFIER
	EXPRESSION_TYPE_NULL
	EXPRESSION_TYPE_NEW
	EXPRESSION_TYPE_CLASS

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an example of a goyacc program.
// To build it:
// goyacc -p "expr" expr.y (produces y.go)
// go build -o expr y.go
// expr
// > <type an expression>

%{

package yacc

import (
	"fmt"
	"math/big"
    "github.com/756445638/lucy/src/cmd/compile/lex"
)

%}





%token lex.TOKEN_FUNCTION   lex.TOKEN_CONST    lex.TOKEN_IF      lex.TOKEN_ELSEIF    lex.TOKEN_ELSE   lex.TOKEN_FOR     lex.TOKEN_BREAK
    	lex.TOKEN_CONTINUE ex.TOKEN_RETURN ex.TOKEN_NULL  lex.TOKEN_LP lex.TOKEN_RP lex.TOKEN_LC lex.TOKEN_RC lex.TOKEN_LB lex.TOKEN_RB
    	lex.TOKEN_SEMICOLON lex.TOKEN_COMMA lex.TOKEN_LOGICAL_AND lex.TOKEN_LOGICAL_OR lex.TOKEN_AND lex.TOKEN_OR lex.TOKEN_ASSIGN
    	lex.TOKEN_EQUAL lex.TOKEN_NE lex.TOKEN_GT lex.TOKEN_GE lex.TOKEN_LT lex.TOKEN_LE lex.TOKEN_ADD lex.TOKEN_SUB lex.TOKEN_MUL
    	lex.TOKEN_DIV lex.TOKEN_MOD lex.TOKEN_INCREMENT lex.TOKEN_DECREMENT lex.TOKEN_DOT lex.TOKEN_VAR lex.TOKEN_NEW lex.TOKEN_COLON
    	lex.TOKEN_PLUS_ASSIGN lex.TOKEN_MINUS_ASSIGN lex.TOKEN_MUL_ASSIGN lex.TOKEN_DIV_ASSIGN lex.TOKEN_MOD_ASSIGN lex.TOKEN_NOT
    	lex.TOKEN_SWITCH lex.TOKEN_CASE lex.TOKEN_DEFAULT lex.TOKEN_CRLF lex.TOKEN_PACKAGE lex.TOKEN_CLASS lex.TOKEN_PUBLIC
    	lex.TOKEN_PROTECTED lex.TOKEN_PRIVATE lex.TOKEN_BOOL lex.TOKEN_BYTE lex.TOKEN_INT lex.TOKEN_FLOAT lex.TOKEN_STRING
    	lex.TOKEN_IDENTIFIER lex.TOKEN_LITERAL_INT lex.TOKEN_LITERAL_STRING lex.TOKEN_LITERAL_FLOAT lex.TOKEN_IMPORT lex.TOKEN_COLON_ASSIGN

%token <boolvalue> lex.TOKEN_TRUE lex.TOKEN_FALSE


%union {
    boolvalue bool
    
}



%%


top:






%%
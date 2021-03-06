package lex

import (
	"github.com/timtadh/lexmachine/machines"
	"strconv"
)

const (
	TOKEN_FUNCTION       = iota // fn
	TOKEN_ENUM                  // enum
	TOKEN_CONST                 //const
	TOKEN_IF                    // if
	TOKEN_ELSEIF                //elseif
	TOKEN_ELSE                  // else
	TOKEN_FOR                   //for
	TOKEN_BREAK                 //break
	TOKEN_CONTINUE              //continue
	TOKEN_RETURN                //return
	TOKEN_NULL                  // null
	TOKEN_BOOL                  //bool
	TOKEN_LP                    //(
	TOKEN_RP                    //)
	TOKEN_LC                    //{
	TOKEN_RC                    //}
	TOKEN_LB                    //[
	TOKEN_RB                    //]
	TOKEN_SKIP                  //skip;
	TOKEN_SEMICOLON             // ;
	TOKEN_CRLF                  // enter
	TOKEN_COMMA                 //,
	TOKEN_LOGICAL_AND           // &&
	TOKEN_LOGICAL_OR            // ||
	TOKEN_AND                   // &
	TOKEN_OR                    // |
	TOKEN_LEFT_SHIFT            // <<
	TOKEN_RIGHT_SHIFT           // >>
	TOKEN_ASSIGN                //=
	TOKEN_EQUAL                 //== or ===
	TOKEN_NE                    // !=
	TOKEN_GT                    //>
	TOKEN_GE                    //>=
	TOKEN_LT                    //<
	TOKEN_LE                    //<=
	TOKEN_ADD                   //+
	TOKEN_SUB                   //-
	TOKEN_MUL                   //*
	TOKEN_DIV                   // a/c
	TOKEN_MOD                   // a%b
	TOKEN_INCREMENT             //a++
	TOKEN_DECREMENT             //a--
	TOKEN_DOT                   // a.do()
	TOKEN_VAR                   // var a
	TOKEN_NEW                   // new Object()
	TOKEN_COLON                 // :
	TOKEN_COLON_ASSIGN          // :=
	TOKEN_PLUS_ASSIGN           // +=
	TOKEN_MINUS_ASSIGN          // -=
	TOKEN_MUL_ASSIGN            // *=
	TOKEN_DIV_ASSIGN            // /=
	TOKEN_MOD_ASSIGN            // %=
	TOKEN_NOT                   // !false
	TOKEN_SWITCH                //swtich
	TOKEN_CASE                  //case
	TOKEN_DEFAULT               //default
	TOKEN_PACKAGE               //package
	TOKEN_IMPORT                //import
	TOKEN_AS                    //as
	TOKEN_CLASS                 //class
	TOKEN_STATIC                //static
	TOKEN_PUBLIC                //public
	TOKEN_PROTECTED             //protected
	TOKEN_PRIVATE               //private
	TOKEN_INTERFACE             //interface
	TOKEN_BYTE                  //byte
	TOKEN_SHORT                 // short
	TOKEN_INT                   //int
	TOKEN_FLOAT                 //float
	TOKEN_DOUBLE                //double
	TOKEN_LONG                  //long
	TOKEN_STRING                //string
	TOKEN_IDENTIFIER            // identifier
	TOKEN_LITERAL_BOOL          // true or false
	TOKEN_LITERAL_BYTE          //'a'
	TOKEN_LITERAL_INT           // 123
	TOKEN_LITERAL_STRING        // ""
	TOKEN_LITERAL_FLOAT         // 0.000
	TOKEN_DEFER                 // defer
	TOKEN_TYPE                  //type
	TOKEN_ARROW                 //->
)

type Token struct {
	Type  int
	Match *machines.Match
	Desp  string
	Data  interface{}
}

func parseInt64(bs []byte) int64 {
	t, _ := strconv.ParseInt(string(bs), 0, 64)
	return t
}

//TODO::解析科学计数法
func parseScientificNotation(bs []byte) (data interface{}, token int) {
	////negative := false
	//if bs[0] == '+' {
	//	bs = bs[1:]
	//}
	//if bs[0] == '-' {
	//	//negative = true
	//	bs = bs[1:]
	//}
	//index := bytes.IndexByte(bs, 'e')
	//pre := bs[0:index]
	//suf := bs[index+1:]
	//power := int(parseInt64(suf))
	////1e5  or 1e-5
	//token = TOKEN_INT
	//base := parseInt64(pre)
	//
	//fmt.Println("@@@@@@@@@@@@@@@@", string(pre), string(suf), string(power))
	return 100, TOKEN_INT

}

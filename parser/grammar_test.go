package parser

import (
	"flag"
	"testing"

	"github.com/ncw/gpython/ast"
)

var debugLevel = flag.Int("debugLevel", 0, "Debug level 0-4")

// FIXME test pos is correct

func TestGrammar(t *testing.T) {
	SetDebug(*debugLevel)
	for _, test := range []struct {
		in   string
		mode string
		out  string
	}{
		// START TESTS
		// *** Tests auto generated by make_grammar_test.py - do not edit ***
		{"", "exec", "Module(body=[])"},
		{"()", "eval", "Expression(body=Tuple(elts=[], ctx=Load()))"},
		{"()", "exec", "Module(body=[Expr(value=Tuple(elts=[], ctx=Load()))])"},
		{"[ ]", "exec", "Module(body=[Expr(value=List(elts=[], ctx=Load()))])"},
		{"True\n", "eval", "Expression(body=NameConstant(value=True))"},
		{"False\n", "eval", "Expression(body=NameConstant(value=False))"},
		{"None\n", "eval", "Expression(body=NameConstant(value=None))"},
		{"...", "eval", "Expression(body=Ellipsis())"},
		{"abc123", "eval", "Expression(body=Name(id='abc123', ctx=Load()))"},
		{"\"abc\"", "eval", "Expression(body=Str(s='abc'))"},
		{"\"abc\" \"\"\"123\"\"\"", "eval", "Expression(body=Str(s='abc123'))"},
		{"b'abc'", "eval", "Expression(body=Bytes(s=b'abc'))"},
		{"b'abc' b'''123'''", "eval", "Expression(body=Bytes(s=b'abc123'))"},
		{"1234", "eval", "Expression(body=Num(n=1234))"},
		{"0x1234", "eval", "Expression(body=Num(n=4660))"},
		{"12.34", "eval", "Expression(body=Num(n=12.34))"},
		{"1,", "eval", "Expression(body=Tuple(elts=[Num(n=1)], ctx=Load()))"},
		{"1,2", "eval", "Expression(body=Tuple(elts=[Num(n=1), Num(n=2)], ctx=Load()))"},
		{"1,2,", "eval", "Expression(body=Tuple(elts=[Num(n=1), Num(n=2)], ctx=Load()))"},
		{"{ }", "eval", "Expression(body=Dict(keys=[], values=[]))"},
		{"{1}", "eval", "Expression(body=Set(elts=[Num(n=1)]))"},
		{"{1,}", "eval", "Expression(body=Set(elts=[Num(n=1)]))"},
		{"{1,2}", "eval", "Expression(body=Set(elts=[Num(n=1), Num(n=2)]))"},
		{"{1,2,3,}", "eval", "Expression(body=Set(elts=[Num(n=1), Num(n=2), Num(n=3)]))"},
		{"{ 'a':1 }", "eval", "Expression(body=Dict(keys=[Str(s='a')], values=[Num(n=1)]))"},
		{"{ 'a':1, 'b':2 }", "eval", "Expression(body=Dict(keys=[Str(s='a'), Str(s='b')], values=[Num(n=1), Num(n=2)]))"},
		{"{ 'a':{'aa':11, 'bb':{'aa':11, 'bb':22}}, 'b':{'aa':11, 'bb':22} }", "eval", "Expression(body=Dict(keys=[Str(s='a'), Str(s='b')], values=[Dict(keys=[Str(s='aa'), Str(s='bb')], values=[Num(n=11), Dict(keys=[Str(s='aa'), Str(s='bb')], values=[Num(n=11), Num(n=22)])]), Dict(keys=[Str(s='aa'), Str(s='bb')], values=[Num(n=11), Num(n=22)])]))"},
		{"(1)", "eval", "Expression(body=Num(n=1))"},
		{"(1,)", "eval", "Expression(body=Tuple(elts=[Num(n=1)], ctx=Load()))"},
		{"(1,2)", "eval", "Expression(body=Tuple(elts=[Num(n=1), Num(n=2)], ctx=Load()))"},
		{"(1,2,)", "eval", "Expression(body=Tuple(elts=[Num(n=1), Num(n=2)], ctx=Load()))"},
		{"{(1,2)}", "eval", "Expression(body=Set(elts=[Tuple(elts=[Num(n=1), Num(n=2)], ctx=Load())]))"},
		{"(((((1,),(2,),),(2,),),((1,),(2,),),),((1,),(2,),))", "eval", "Expression(body=Tuple(elts=[Tuple(elts=[Tuple(elts=[Tuple(elts=[Tuple(elts=[Num(n=1)], ctx=Load()), Tuple(elts=[Num(n=2)], ctx=Load())], ctx=Load()), Tuple(elts=[Num(n=2)], ctx=Load())], ctx=Load()), Tuple(elts=[Tuple(elts=[Num(n=1)], ctx=Load()), Tuple(elts=[Num(n=2)], ctx=Load())], ctx=Load())], ctx=Load()), Tuple(elts=[Tuple(elts=[Num(n=1)], ctx=Load()), Tuple(elts=[Num(n=2)], ctx=Load())], ctx=Load())], ctx=Load()))"},
		{"(((1)))", "eval", "Expression(body=Num(n=1))"},
		{"[1]", "eval", "Expression(body=List(elts=[Num(n=1)], ctx=Load()))"},
		{"[1,]", "eval", "Expression(body=List(elts=[Num(n=1)], ctx=Load()))"},
		{"[1,2]", "eval", "Expression(body=List(elts=[Num(n=1), Num(n=2)], ctx=Load()))"},
		{"[1,2,]", "eval", "Expression(body=List(elts=[Num(n=1), Num(n=2)], ctx=Load()))"},
		{"( a for a in ab )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"( a for a, in ab )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"( a for a, b in ab )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store()), Name(id='b', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"( a for a in ab if a )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load())])]))"},
		{"( a for a in ab if a if b if c )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())])]))"},
		{"( a for a in ab for A in AB )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[])]))"},
		{"( a for a in ab if a if b for A in AB if c )", "eval", "Expression(body=GeneratorExp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[Name(id='c', ctx=Load())])]))"},
		{"[ a for a in ab ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"[ a for a, in ab ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"[ a for a, b in ab ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store()), Name(id='b', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"[ a for a in ab if a ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load())])]))"},
		{"[ a for a in ab if a if b if c ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())])]))"},
		{"[ a for a in ab for A in AB ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[])]))"},
		{"[ a for a in ab if a if b for A in AB if c ]", "eval", "Expression(body=ListComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[Name(id='c', ctx=Load())])]))"},
		{"{ a for a in ab }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"{ a for a, in ab }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"{ a for a, b in ab }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store()), Name(id='b', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"{ a for a in ab if a }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load())])]))"},
		{"{ a for a in ab if a if b if c }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())])]))"},
		{"{ a for a in ab for A in AB }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[])]))"},
		{"{ a for a in ab if a if b for A in AB if c }", "eval", "Expression(body=SetComp(elt=Name(id='a', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[Name(id='c', ctx=Load())])]))"},
		{"{ a:b for a in ab }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"{ a:b for a, in ab }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"{ a:b for a, b in ab }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Tuple(elts=[Name(id='a', ctx=Store()), Name(id='b', ctx=Store())], ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[])]))"},
		{"{ a:b for a in ab if a }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load())])]))"},
		{"{ a:b for a in ab if a if b if c }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())])]))"},
		{"{ a:b for a in ab for A in AB }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[])]))"},
		{"{ a:b for a in ab if a if b for A in AB if c }", "eval", "Expression(body=DictComp(key=Name(id='a', ctx=Load()), value=Name(id='b', ctx=Load()), generators=[comprehension(target=Name(id='a', ctx=Store()), iter=Name(id='ab', ctx=Load()), ifs=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]), comprehension(target=Name(id='A', ctx=Store()), iter=Name(id='AB', ctx=Load()), ifs=[Name(id='c', ctx=Load())])]))"},
		{"a|b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=BitOr(), right=Name(id='b', ctx=Load())))"},
		{"a^b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=BitXor(), right=Name(id='b', ctx=Load())))"},
		{"a&b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=BitAnd(), right=Name(id='b', ctx=Load())))"},
		{"a<<b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=LShift(), right=Name(id='b', ctx=Load())))"},
		{"a>>b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=RShift(), right=Name(id='b', ctx=Load())))"},
		{"a+b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=Add(), right=Name(id='b', ctx=Load())))"},
		{"a-b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=Sub(), right=Name(id='b', ctx=Load())))"},
		{"a*b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=Mult(), right=Name(id='b', ctx=Load())))"},
		{"a/b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=Div(), right=Name(id='b', ctx=Load())))"},
		{"a//b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=FloorDiv(), right=Name(id='b', ctx=Load())))"},
		{"a**b", "eval", "Expression(body=BinOp(left=Name(id='a', ctx=Load()), op=Pow(), right=Name(id='b', ctx=Load())))"},
		{"not a", "eval", "Expression(body=UnaryOp(op=Not(), operand=Name(id='a', ctx=Load())))"},
		{"+a", "eval", "Expression(body=UnaryOp(op=UAdd(), operand=Name(id='a', ctx=Load())))"},
		{"-a", "eval", "Expression(body=UnaryOp(op=USub(), operand=Name(id='a', ctx=Load())))"},
		{"~a", "eval", "Expression(body=UnaryOp(op=Invert(), operand=Name(id='a', ctx=Load())))"},
		{"a and b", "eval", "Expression(body=BoolOp(op=And(), values=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]))"},
		{"a or b", "eval", "Expression(body=BoolOp(op=Or(), values=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]))"},
		{"a or b or c", "eval", "Expression(body=BoolOp(op=Or(), values=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())]))"},
		{"(a or b) or c", "eval", "Expression(body=BoolOp(op=Or(), values=[BoolOp(op=Or(), values=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]), Name(id='c', ctx=Load())]))"},
		{"a or (b or c)", "eval", "Expression(body=BoolOp(op=Or(), values=[Name(id='a', ctx=Load()), BoolOp(op=Or(), values=[Name(id='b', ctx=Load()), Name(id='c', ctx=Load())])]))"},
		{"a and b and c", "eval", "Expression(body=BoolOp(op=And(), values=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())]))"},
		{"(a and b) and c", "eval", "Expression(body=BoolOp(op=And(), values=[BoolOp(op=And(), values=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())]), Name(id='c', ctx=Load())]))"},
		{"a and (b and c)", "eval", "Expression(body=BoolOp(op=And(), values=[Name(id='a', ctx=Load()), BoolOp(op=And(), values=[Name(id='b', ctx=Load()), Name(id='c', ctx=Load())])]))"},
		{"a+b-c/d", "eval", "Expression(body=BinOp(left=BinOp(left=Name(id='a', ctx=Load()), op=Add(), right=Name(id='b', ctx=Load())), op=Sub(), right=BinOp(left=Name(id='c', ctx=Load()), op=Div(), right=Name(id='d', ctx=Load()))))"},
		{"a+b-c/d//e", "eval", "Expression(body=BinOp(left=BinOp(left=Name(id='a', ctx=Load()), op=Add(), right=Name(id='b', ctx=Load())), op=Sub(), right=BinOp(left=BinOp(left=Name(id='c', ctx=Load()), op=Div(), right=Name(id='d', ctx=Load())), op=FloorDiv(), right=Name(id='e', ctx=Load()))))"},
		{"a+b-c/d//e%f", "eval", "Expression(body=BinOp(left=BinOp(left=Name(id='a', ctx=Load()), op=Add(), right=Name(id='b', ctx=Load())), op=Sub(), right=BinOp(left=BinOp(left=BinOp(left=Name(id='c', ctx=Load()), op=Div(), right=Name(id='d', ctx=Load())), op=FloorDiv(), right=Name(id='e', ctx=Load())), op=Mod(), right=Name(id='f', ctx=Load()))))"},
		{"a+b-c/d//e%f**g", "eval", "Expression(body=BinOp(left=BinOp(left=Name(id='a', ctx=Load()), op=Add(), right=Name(id='b', ctx=Load())), op=Sub(), right=BinOp(left=BinOp(left=BinOp(left=Name(id='c', ctx=Load()), op=Div(), right=Name(id='d', ctx=Load())), op=FloorDiv(), right=Name(id='e', ctx=Load())), op=Mod(), right=BinOp(left=Name(id='f', ctx=Load()), op=Pow(), right=Name(id='g', ctx=Load())))))"},
		{"a+b-c/d//e%f**g|h&i^k<<l>>m", "eval", "Expression(body=BinOp(left=BinOp(left=BinOp(left=Name(id='a', ctx=Load()), op=Add(), right=Name(id='b', ctx=Load())), op=Sub(), right=BinOp(left=BinOp(left=BinOp(left=Name(id='c', ctx=Load()), op=Div(), right=Name(id='d', ctx=Load())), op=FloorDiv(), right=Name(id='e', ctx=Load())), op=Mod(), right=BinOp(left=Name(id='f', ctx=Load()), op=Pow(), right=Name(id='g', ctx=Load())))), op=BitOr(), right=BinOp(left=BinOp(left=Name(id='h', ctx=Load()), op=BitAnd(), right=Name(id='i', ctx=Load())), op=BitXor(), right=BinOp(left=BinOp(left=Name(id='k', ctx=Load()), op=LShift(), right=Name(id='l', ctx=Load())), op=RShift(), right=Name(id='m', ctx=Load())))))"},
		{"a==b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Eq()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a!=b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[NotEq()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a<b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Lt()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a<=b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[LtE()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a>b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Gt()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a>=b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[GtE()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a is b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Is()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a is not b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[IsNot()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a in b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[In()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a not in b", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[NotIn()], comparators=[Name(id='b', ctx=Load())]))"},
		{"a<b<c<d", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Lt(), Lt(), Lt()], comparators=[Name(id='b', ctx=Load()), Name(id='c', ctx=Load()), Name(id='d', ctx=Load())]))"},
		{"a==b<c>d", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Eq(), Lt(), Gt()], comparators=[Name(id='b', ctx=Load()), Name(id='c', ctx=Load()), Name(id='d', ctx=Load())]))"},
		{"(a==b)<c", "eval", "Expression(body=Compare(left=Compare(left=Name(id='a', ctx=Load()), ops=[Eq()], comparators=[Name(id='b', ctx=Load())]), ops=[Lt()], comparators=[Name(id='c', ctx=Load())]))"},
		{"a==(b<c)", "eval", "Expression(body=Compare(left=Name(id='a', ctx=Load()), ops=[Eq()], comparators=[Compare(left=Name(id='b', ctx=Load()), ops=[Lt()], comparators=[Name(id='c', ctx=Load())])]))"},
		{"(a==b)<(c>d)>e", "eval", "Expression(body=Compare(left=Compare(left=Name(id='a', ctx=Load()), ops=[Eq()], comparators=[Name(id='b', ctx=Load())]), ops=[Lt(), Gt()], comparators=[Compare(left=Name(id='c', ctx=Load()), ops=[Gt()], comparators=[Name(id='d', ctx=Load())]), Name(id='e', ctx=Load())]))"},
		{"a()", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[], keywords=[], starargs=None, kwargs=None))"},
		{"a(b)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='b', ctx=Load())], keywords=[], starargs=None, kwargs=None))"},
		{"a(b,)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='b', ctx=Load())], keywords=[], starargs=None, kwargs=None))"},
		{"a(b,c)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='b', ctx=Load()), Name(id='c', ctx=Load())], keywords=[], starargs=None, kwargs=None))"},
		{"a(b,*c)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='b', ctx=Load())], keywords=[], starargs=Name(id='c', ctx=Load()), kwargs=None))"},
		{"a(*b)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[], keywords=[], starargs=Name(id='b', ctx=Load()), kwargs=None))"},
		{"a(b,*c,**d)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='b', ctx=Load())], keywords=[], starargs=Name(id='c', ctx=Load()), kwargs=Name(id='d', ctx=Load())))"},
		{"a(b,**c)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='b', ctx=Load())], keywords=[], starargs=None, kwargs=Name(id='c', ctx=Load())))"},
		{"a(a=b)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[], keywords=[keyword(arg='a', value=Name(id='b', ctx=Load()))], starargs=None, kwargs=None))"},
		{"a(a,a=b,*args,**kwargs)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='a', ctx=Load())], keywords=[keyword(arg='a', value=Name(id='b', ctx=Load()))], starargs=Name(id='args', ctx=Load()), kwargs=Name(id='kwargs', ctx=Load())))"},
		{"a(a,a=b,*args,e=f,**kwargs)", "eval", "Expression(body=Call(func=Name(id='a', ctx=Load()), args=[Name(id='a', ctx=Load())], keywords=[keyword(arg='a', value=Name(id='b', ctx=Load())), keyword(arg='e', value=Name(id='f', ctx=Load()))], starargs=Name(id='args', ctx=Load()), kwargs=Name(id='kwargs', ctx=Load())))"},
		{"a.b", "eval", "Expression(body=Attribute(value=Name(id='a', ctx=Load()), attr='b', ctx=Load()))"},
		{"a.b.c.d", "eval", "Expression(body=Attribute(value=Attribute(value=Attribute(value=Name(id='a', ctx=Load()), attr='b', ctx=Load()), attr='c', ctx=Load()), attr='d', ctx=Load()))"},
		{"a.b().c.d()()", "eval", "Expression(body=Call(func=Call(func=Attribute(value=Attribute(value=Call(func=Attribute(value=Name(id='a', ctx=Load()), attr='b', ctx=Load()), args=[], keywords=[], starargs=None, kwargs=None), attr='c', ctx=Load()), attr='d', ctx=Load()), args=[], keywords=[], starargs=None, kwargs=None), args=[], keywords=[], starargs=None, kwargs=None))"},
		{"x[a]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Index(value=Name(id='a', ctx=Load())), ctx=Load()))"},
		{"x[a:b]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=Name(id='a', ctx=Load()), upper=Name(id='b', ctx=Load()), step=None), ctx=Load()))"},
		{"x[a:b:c]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=Name(id='a', ctx=Load()), upper=Name(id='b', ctx=Load()), step=Name(id='c', ctx=Load())), ctx=Load()))"},
		{"x[:b:c]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=None, upper=Name(id='b', ctx=Load()), step=Name(id='c', ctx=Load())), ctx=Load()))"},
		{"x[a::c]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=Name(id='a', ctx=Load()), upper=None, step=Name(id='c', ctx=Load())), ctx=Load()))"},
		{"x[a:b:]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=Name(id='a', ctx=Load()), upper=Name(id='b', ctx=Load()), step=None), ctx=Load()))"},
		{"x[::c]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=None, upper=None, step=Name(id='c', ctx=Load())), ctx=Load()))"},
		{"x[:b:]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=None, upper=Name(id='b', ctx=Load()), step=None), ctx=Load()))"},
		{"x[::c]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=None, upper=None, step=Name(id='c', ctx=Load())), ctx=Load()))"},
		{"x[::]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Slice(lower=None, upper=None, step=None), ctx=Load()))"},
		{"x[a,p]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Index(value=Tuple(elts=[Name(id='a', ctx=Load()), Name(id='p', ctx=Load())], ctx=Load())), ctx=Load()))"},
		{"x[a, b]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Index(value=Tuple(elts=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())], ctx=Load())), ctx=Load()))"},
		{"x[a, b, c]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=Index(value=Tuple(elts=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load()), Name(id='c', ctx=Load())], ctx=Load())), ctx=Load()))"},
		{"x[a, b:c, ::d]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=ExtSlice(dims=[Index(value=Name(id='a', ctx=Load())), Slice(lower=Name(id='b', ctx=Load()), upper=Name(id='c', ctx=Load()), step=None), Slice(lower=None, upper=None, step=Name(id='d', ctx=Load()))]), ctx=Load()))"},
		{"x[a, b:c, ::d]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=ExtSlice(dims=[Index(value=Name(id='a', ctx=Load())), Slice(lower=Name(id='b', ctx=Load()), upper=Name(id='c', ctx=Load()), step=None), Slice(lower=None, upper=None, step=Name(id='d', ctx=Load()))]), ctx=Load()))"},
		{"x[0, 1:2, ::5, ...]", "eval", "Expression(body=Subscript(value=Name(id='x', ctx=Load()), slice=ExtSlice(dims=[Index(value=Num(n=0)), Slice(lower=Num(n=1), upper=Num(n=2), step=None), Slice(lower=None, upper=None, step=Num(n=5)), Index(value=Ellipsis())]), ctx=Load()))"},
		{"(yield a,b)", "eval", "Expression(body=Yield(value=Tuple(elts=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())], ctx=Load())))"},
		{"(yield from a)", "eval", "Expression(body=YieldFrom(value=Name(id='a', ctx=Load())))"},
		{"del a,b", "exec", "Module(body=[Delete(targets=[Name(id='a', ctx=Del()), Name(id='b', ctx=Del())])])"},
		{"pass", "exec", "Module(body=[Pass()])"},
		{"break", "exec", "Module(body=[Break()])"},
		{"continue", "exec", "Module(body=[Continue()])"},
		{"return", "exec", "Module(body=[Return(value=None)])"},
		{"return a", "exec", "Module(body=[Return(value=Name(id='a', ctx=Load()))])"},
		{"return a,", "exec", "Module(body=[Return(value=Tuple(elts=[Name(id='a', ctx=Load())], ctx=Load()))])"},
		{"return a,b", "exec", "Module(body=[Return(value=Tuple(elts=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())], ctx=Load()))])"},
		{"raise", "exec", "Module(body=[Raise(exc=None, cause=None)])"},
		{"raise a", "exec", "Module(body=[Raise(exc=Name(id='a', ctx=Load()), cause=None)])"},
		{"raise a from b", "exec", "Module(body=[Raise(exc=Name(id='a', ctx=Load()), cause=Name(id='b', ctx=Load()))])"},
		{"yield", "exec", "Module(body=[Expr(value=Yield(value=None))])"},
		{"yield a", "exec", "Module(body=[Expr(value=Yield(value=Name(id='a', ctx=Load())))])"},
		{"yield a, b", "exec", "Module(body=[Expr(value=Yield(value=Tuple(elts=[Name(id='a', ctx=Load()), Name(id='b', ctx=Load())], ctx=Load())))])"},
		{"import a", "exec", "Module(body=[Import(names=[alias(name='a', asname=None)])])"},
		{"import a . b,c .d.e", "exec", "Module(body=[Import(names=[alias(name='a.b', asname=None), alias(name='c.d.e', asname=None)])])"},
		{"from a import b", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None)], level=0)])"},
		{"from a import b, c", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None), alias(name='c', asname=None)], level=0)])"},
		{"from a import (b, c)", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None), alias(name='c', asname=None)], level=0)])"},
		{"from a import *", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='*', asname=None)], level=0)])"},
		{"from .a import (b, c,)", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None), alias(name='c', asname=None)], level=1)])"},
		{"from ..a import b", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None)], level=2)])"},
		{"from ...a import b", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None)], level=3)])"},
		{"from ....a import b", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None)], level=4)])"},
		{"from .....a import b", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None)], level=5)])"},
		{"from ......a import b", "exec", "Module(body=[ImportFrom(module='a', names=[alias(name='b', asname=None)], level=6)])"},
		{"global a", "exec", "Module(body=[Global(names=['a'])])"},
		{"global a, b", "exec", "Module(body=[Global(names=['a', 'b'])])"},
		{"global a, b, c", "exec", "Module(body=[Global(names=['a', 'b', 'c'])])"},
		{"nonlocal a", "exec", "Module(body=[Nonlocal(names=['a'])])"},
		{"nonlocal a, b", "exec", "Module(body=[Nonlocal(names=['a', 'b'])])"},
		{"nonlocal a, b, c", "exec", "Module(body=[Nonlocal(names=['a', 'b', 'c'])])"},
		{"assert True", "exec", "Module(body=[Assert(test=NameConstant(value=True), msg=None)])"},
		{"assert True, 'Bang'", "exec", "Module(body=[Assert(test=NameConstant(value=True), msg=Str(s='Bang'))])"},
		{"assert a == b, 'Bang'", "exec", "Module(body=[Assert(test=Compare(left=Name(id='a', ctx=Load()), ops=[Eq()], comparators=[Name(id='b', ctx=Load())]), msg=Str(s='Bang'))])"},
		{"while True: pass", "exec", "Module(body=[While(test=NameConstant(value=True), body=[Pass()], orelse=[])])"},
		{"while True:\n pass\n", "exec", "Module(body=[While(test=NameConstant(value=True), body=[Pass()], orelse=[])])"},
		{"while True:\n pass\nelse:\n return\n", "exec", "Module(body=[While(test=NameConstant(value=True), body=[Pass()], orelse=[Return(value=None)])])"},
		{"if True: pass", "exec", "Module(body=[If(test=NameConstant(value=True), body=[Pass()], orelse=[])])"},
		{"if True:\n pass\n", "exec", "Module(body=[If(test=NameConstant(value=True), body=[Pass()], orelse=[])])"},
		{"if True:\n    pass\n    continue\nelse:\n    break\n    pass\n", "exec", "Module(body=[If(test=NameConstant(value=True), body=[Pass(), Continue()], orelse=[Break(), Pass()])])"},
		{"if a:\n    continue\nelif b:\n    break\nelif c:\n    pass\nelif c:\n    continue\n    pass\n", "exec", "Module(body=[If(test=Name(id='a', ctx=Load()), body=[Continue()], orelse=[If(test=Name(id='b', ctx=Load()), body=[Break()], orelse=[If(test=Name(id='c', ctx=Load()), body=[Pass()], orelse=[If(test=Name(id='c', ctx=Load()), body=[Continue(), Pass()], orelse=[])])])])])"},
		{"if a:\n    continue\nelif b:\n    break\nelse:\n    continue\n    pass\n", "exec", "Module(body=[If(test=Name(id='a', ctx=Load()), body=[Continue()], orelse=[If(test=Name(id='b', ctx=Load()), body=[Break()], orelse=[Continue(), Pass()])])])"},
		{"if a:\n    continue\nelif b:\n    break\nelif c:\n    pass\nelse:\n    continue\n    pass\n", "exec", "Module(body=[If(test=Name(id='a', ctx=Load()), body=[Continue()], orelse=[If(test=Name(id='b', ctx=Load()), body=[Break()], orelse=[If(test=Name(id='c', ctx=Load()), body=[Pass()], orelse=[Continue(), Pass()])])])])"},
		{"for a in b: pass", "exec", "Module(body=[For(target=Name(id='a', ctx=Store()), iter=Name(id='b', ctx=Load()), body=[Pass()], orelse=[])])"},
		{"for a, b in b: pass", "exec", "Module(body=[For(target=Tuple(elts=[Name(id='a', ctx=Store()), Name(id='b', ctx=Store())], ctx=Store()), iter=Name(id='b', ctx=Load()), body=[Pass()], orelse=[])])"},
		{"for a, b in b:\n pass\nelse: break\n", "exec", "Module(body=[For(target=Tuple(elts=[Name(id='a', ctx=Store()), Name(id='b', ctx=Store())], ctx=Store()), iter=Name(id='b', ctx=Load()), body=[Pass()], orelse=[Break()])])"},
		{"try:\n    pass\nexcept:\n    break\n", "exec", "Module(body=[Try(body=[Pass()], handlers=[ExceptHandler(type=None, name=None, body=[Break()])], orelse=[], finalbody=[])])"},
		{"try:\n    pass\nexcept a:\n    break\n", "exec", "Module(body=[Try(body=[Pass()], handlers=[ExceptHandler(type=Name(id='a', ctx=Load()), name=None, body=[Break()])], orelse=[], finalbody=[])])"},
		{"try:\n    pass\nexcept a as b:\n    break\n", "exec", "Module(body=[Try(body=[Pass()], handlers=[ExceptHandler(type=Name(id='a', ctx=Load()), name='b', body=[Break()])], orelse=[], finalbody=[])])"},
		{"try:\n    pass\nexcept a:\n    break\nexcept:\n    continue\nexcept b as c:\n    break\nelse:\n    pass\n", "exec", "Module(body=[Try(body=[Pass()], handlers=[ExceptHandler(type=Name(id='a', ctx=Load()), name=None, body=[Break()]), ExceptHandler(type=None, name=None, body=[Continue()]), ExceptHandler(type=Name(id='b', ctx=Load()), name='c', body=[Break()])], orelse=[Pass()], finalbody=[])])"},
		{"try:\n    pass\nexcept:\n    continue\nfinally:\n    pass\n", "exec", "Module(body=[Try(body=[Pass()], handlers=[ExceptHandler(type=None, name=None, body=[Continue()])], orelse=[], finalbody=[Pass()])])"},
		{"try:\n    pass\nexcept:\n    continue\nelse:\n    break\nfinally:\n    pass\n", "exec", "Module(body=[Try(body=[Pass()], handlers=[ExceptHandler(type=None, name=None, body=[Continue()])], orelse=[Break()], finalbody=[Pass()])])"},
		{"with x:\n    pass\n", "exec", "Module(body=[With(items=[withitem(context_expr=Name(id='x', ctx=Load()), optional_vars=None)], body=[Pass()])])"},
		{"with x as y:\n    pass\n", "exec", "Module(body=[With(items=[withitem(context_expr=Name(id='x', ctx=Load()), optional_vars=Name(id='y', ctx=Store()))], body=[Pass()])])"},
		{"with x as y, a as b, c, d as e:\n    pass\n    continue\n", "exec", "Module(body=[With(items=[withitem(context_expr=Name(id='x', ctx=Load()), optional_vars=Name(id='y', ctx=Store())), withitem(context_expr=Name(id='a', ctx=Load()), optional_vars=Name(id='b', ctx=Store())), withitem(context_expr=Name(id='c', ctx=Load()), optional_vars=None), withitem(context_expr=Name(id='d', ctx=Load()), optional_vars=Name(id='e', ctx=Store()))], body=[Pass(), Continue()])])"},
		// END TESTS
	} {
		Ast, err := ParseString(test.in, test.mode)
		if err != nil {
			t.Errorf("Parse(%q) returned error: %v", test.in, err)
		} else {
			out := ast.Dump(Ast)
			if out != test.out {
				t.Errorf("Parse(%q)\nwant> %q\n got> %q\n", test.in, test.out, out)
			}
		}
	}
}

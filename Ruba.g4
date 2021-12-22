grammar Ruba;

program: block EOF;

block: stmt*;

ident: Ident;
// number: '-'? Int | Float;
intType: Int;
floatType: Float;
stringType: Str;
mathSign: '+' | '-' | '/' | '*';

nestable: (('.' DOT = ident) | ( '[' INDEX = expr ']'));
stmt:
	'import' NAME = stringType ('as' ALIAS = ident)?		# ImportStmt
	| PUB = 'pub'? NAME = ident ':=' VALUE = expr			# DeclareStmt
	| ROOT = ident (CHAIN += nestable)* '=' VALUE = expr	# AssignStmt
	| fnCall												# FnCallStmt
	| 'func' NAME = ident '(' (
		(
			(
				ARGS += ident (',' ARGS += ident)* (
					',' '...' RESTARG = ident
				)?
			)
			// ARGS += ident (',' ARGS += ident)* ( ',' '...' RESTARG = ident) )
			| ('...' RESTARG = ident)
		)
	)? ')' '{' BODY = block '}'																# FnDeclareStmt
	| 'for' NAMES += ident (',' NAMES += ident)* 'in' TARGET = expr '{' BODY = block '}'	# ForInStmt
	| 'for' COND = expr '{' BODY = block '}' (
		'else' '{' ELSEBODY = block '}'
	)? # ForCondStmt
	| 'if' COND = expr '{' BODY = block '}' (
		'else' 'if' ELIFCOND += expr '{' ELIFBODY += block '}'
	)* ('else' '{' ELSEBODY = block '}')?	# IfStmt
	| 'return' (VALUE = expr)?				# ReturnStmt
	| 'break'								# BreakStmt
	| E = expr								# ExprStmt;

fnArg: ARG = expr | SPREADABLE = expr '...';

fnCall:
	(NAME = ident | EXPR = expr) '(' (
		ARGS += fnArg (',' ARGS += fnArg)*
	)? ')' ';'?;

objectField: (STATIC = ident | '[' DYNAMIC = expr ']') ':' VALUE = expr;

expr:
	// COND = expr '?' TRUE = expr ':' FALSE = expr						# IfExpr
	'(' EXPR = expr ')'													# GroupExpr
	| ('true' | 'false')												# BoolExpr
	| 'null'															# NullExpr
	| NAME = expr '(' (ARGS += fnArg (',' ARGS += fnArg)*)? ')'			# CallExpr
	| ident																# IdentExpr
	| PARENT = expr '[' CHILD = expr ']'								# IndexExpr
	| PARENT = expr '.' CHILD = ident									# DotExpr
	| intType															# IntExpr
	| floatType															# FloatExpr
	| stringType														# StringExpr
	| '!' expr															# NotExpr
	| '-' expr															# NegativeExpr
	| L = expr S = ('*' | '/') R = expr									# MulOrDivExpr
	| L = expr S = ('+' | '-') R = expr									# AddOrSubExpr
	| L = expr '==' R = expr											# EqualExpr
	| L = expr '!=' R = expr											# NotEqualExpr
	| L = expr '>' R = expr												# GtExpr
	| L = expr '<=' R = expr											# LteExpr
	| L = expr '<' R = expr												# LtExpr
	| L = expr '>=' R = expr											# GteExpr
	| L = expr '++'														# IncExpr
	| ident '--'														# DecExpr
	| '[' (ELMS += expr (',' ELMS += expr)*)? ']'						# ArrayExpr
	| '{' (ELMS += objectField (',' ELMS += objectField)* ','?)? '}'	# ObjectExpr;

//types
Ident: [a-zA-Z_][a-zA-Z_0-9]*;
Int: Digit+;
Str: '"' ( EscapeSequence | ~('\\' | '"'))* '"';
Float:
	Digit+ '.' Digit* ExponentPart?
	| '.' Digit+ ExponentPart?
	| Digit+ ExponentPart;

WS: [ \t\u000C\r\n]+ -> skip;

fragment Digit: [0-9];
fragment EscapeSequence:
	'\\' [abfnrtvz"'\\]
	| '\\' '\r'? '\n'
	| DecimalEscape
	| HexEscape
	| UtfEscape;
fragment DecimalEscape:
	'\\' Digit
	| '\\' Digit Digit
	| '\\' [0-2] Digit Digit;
fragment HexEscape: '\\' 'x' HexDigit HexDigit;
fragment UtfEscape: '\\' 'u{' HexDigit+ '}';

fragment HexDigit: [0-9a-fA-F];
fragment ExponentPart: [eE] [+-]? Digit+;
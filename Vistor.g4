grammar Vistor;

NUM:[0-9]+;
num:NUM;
calculate:num op=('+'|'-'|'*'|'/') num ';'?;

WS:[\r\n \t]+ -> skip;
IF:'if';
ELSIF:'elsif';
ELSE:'else';
END:'end';
condition:'true'|'false';

ifStatement:
    IF condition '{'
        statement*
    '}'
    elseIfStatement*
    elseStatement
    ;

elseIfStatement:
    ELSIF condition '{'
        statement*
'}';

elseStatement:
    ELSE '{'
    statement*
'}';

statement
    :ifStatement
    |calculate
    ;

start:statement*EOF;
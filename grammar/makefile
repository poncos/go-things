

gen:
	$env:CLASSPATH=".;C:\Users\estebancc\dev\libs\antlr-4.8-complete.jar"
	java org.antlr.v4.Tool -Dlanguage=Go -o ..\go-regex-eval\internal\grammar\parser ValidationExpr.g4
build:
	go build .\gossip_validation.go
test:
	java -Xmx500M org.antlr.v4.gui.TestRig
run:
	./validation.exe
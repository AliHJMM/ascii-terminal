echo '1-Try passing as no arguments '
go run main.go 

echo '2-Try passing as arguments --color red "hello" '
go run main.go --color red "hello"

echo '3-Try passing as arguments --color=red "hello world"'
go run main.go --color=red "hello world"

echo '4-Try passing as arguments --color=green "1 + 1 = 2"'
go run main.go --color=green "1 + 1 = 2"

echo '5-Try passing as arguments --color=yellow "(%&) ??"'
go run main.go --color=yellow "(%&) ??"

echo '6-Try passing as arguments --color=orange GuYs "HeY GuYs" '
go run main.go --color=orange GuYs "HeY GuYs"

echo '7-Try passing as arguments --color=blue B "RGB()" '
go run main.go --color=blue B 'RGB()'

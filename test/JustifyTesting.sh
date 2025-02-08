echo '1-Try passing as no arguments '
go run main.go 

echo '2-Try passing as arguments --align right something standard '
go run main.go --align right something standard

echo '3-Try passing as arguments --align=right left standard '
go run main.go --align=right left standard

echo '4-Try passing as arguments --align=left right standard '
go run main.go --align=left right standard

echo '5-Try passing as arguments --align=center hello shadow '
go run main.go --align=center hello shadow

echo '6-Try passing as arguments --align=justify "1 Two 4" shadow '
go run main.go --align=justify "1 Two 4" shadow

echo '7-Try passing as arguments --align=right 23/32 standard '
go run main.go --align=right 23/32 standard

echo '8-Try passing as arguments --align=right ABCabc123 thinkertoy '
go run main.go --align=right ABCabc123 thinkertoy

echo '9-Try passing as arguments --align=center "#$%&\"" thinkertoy '
go run main.go --align=center "#$%&\"" thinkertoy

echo '10-Try passing as arguments --align=left '23Hello World!' standard '
go run main.go --align=left '23Hello World!' standard

echo '11-Try passing as arguments --align=justify 'HELLO there HOW are YOU?!' thinkertoy'
go run main.go --align=justify 'HELLO there HOW are YOU?!' thinkertoy

echo '12-Try passing as arguments --align=right "a -> A b -> B c -> C" shadow '
go run main.go --align=right "a -> A b -> B c -> C" shadow

echo '13-Try passing as arguments --align=right abcd shadow '
go run main.go --align=right abcd shadow

echo '14-Try passing as arguments --align=center ola standard '
go run main.go --align=center ola standard


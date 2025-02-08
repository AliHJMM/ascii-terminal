echo "..........................................................."

echo '1-Try passing as arguments "--reverse example00.txt"'
go run main.go "--reverse example00.txt"
ech | cat -eo "..........................................................."
cat -e AuditExamples/example00.txt
echo '2-Try passing as arguments --reverse=AuditExamples/example00.txt'
go run main.go --reverse=AuditExamples/example00.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example01.txt
echo '3-Try passing as arguments --reverse=AuditExamples/example01.txt'
go run main.go --reverse=AuditExamples/example01.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example02.txt
echo '4-Try passing as arguments --reverse=AuditExamples/example02.txt'
go run main.go --reverse=AuditExamples/example02.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example03.txt
echo '5-Try passing as arguments --reverse=AuditExamples/example03.txt'
go run main.go --reverse=AuditExamples/example03.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example04.txt
echo '6-Try passing as arguments --reverse=AuditExamples/example04.txt'
go run main.go --reverse=AuditExamples/example04.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example05.txt
echo '7-Try passing as arguments --reverse=AuditExamples/example05.txt'
go run main.go --reverse=AuditExamples/example05.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example06.txt
echo '8-Try passing as arguments --reverse=AuditExamples/example06.txt'
go run main.go --reverse=AuditExamples/example06.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example07.txt
echo '9-Try passing as arguments --reverse=AuditExamples/example07.txt'
go run main.go --reverse=AuditExamples/example07.txt | cat -e
echo "..........................................................."

cat -e AuditExamples/example08.txt
echo '10-Try passing as arguments --reverse=AuditExamples/example08.txt'
go run main.go --reverse=AuditExamples/example08.txt | cat -e
echo "..........................................................."


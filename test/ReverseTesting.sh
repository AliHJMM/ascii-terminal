#!/bin/bash

echo "-------------------------------------------------------------"
echo "Running reverse tests on example files..."
echo "-------------------------------------------------------------"

for i in {00..07}; do
    FILE="AuditExamples/example$i.txt"
    
    if [ -f "$FILE" ]; then
        echo "Test $((10#$i + 1)): Reversing $FILE"
        echo "-------------------------------------------------------------"
        cat -e "$FILE"
        echo "-------------------------------------------------------------"
        go run main.go --reverse="$FILE" | cat -e
        echo "-------------------------------------------------------------"
    else
        echo "Skipping Test $((10#$i + 1)): File $FILE not found."
    fi
done

echo "All tests completed."

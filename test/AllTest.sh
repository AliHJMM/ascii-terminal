#!/bin/bash

# Run ColorTesting.sh
echo "-------------------------------------------------------------"
echo "Running Color Tests..."
echo "-------------------------------------------------------------"
bash test/ColorTesting.sh

# Run JustifyTesting.sh
echo "-------------------------------------------------------------"
echo "Running Justify Tests..."
echo "-------------------------------------------------------------"
bash test/JustifyTesting.sh

# Run ReverseTesting.sh
echo "-------------------------------------------------------------"
echo "Running Reverse Tests..."
echo "-------------------------------------------------------------"
bash test/ReverseTesting.sh

echo "-------------------------------------------------------------"
echo "All tests completed."
echo "-------------------------------------------------------------"

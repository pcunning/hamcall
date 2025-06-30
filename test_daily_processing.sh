#!/bin/bash

# Test script to demonstrate daily file processing functionality

echo "=== Testing Daily File Processing Feature ==="
echo

echo "1. Testing weekly mode (default):"
unset ULS_USE_DAILY
echo "ULS_USE_DAILY is unset"
echo "Expected: Weekly file processing"
echo

echo "2. Testing daily mode:"
export ULS_USE_DAILY=true
echo "ULS_USE_DAILY=$ULS_USE_DAILY"
echo "Expected: Daily file processing"
echo

echo "3. Testing invalid value:"
export ULS_USE_DAILY=false
echo "ULS_USE_DAILY=$ULS_USE_DAILY"
echo "Expected: Weekly file processing (fallback)"
echo

echo "4. Running ULS tests to verify functionality:"
cd /home/runner/work/hamcall/hamcall
go test ./source/uls -v

echo
echo "=== Test Complete ==="
#!/bin/bash

printf "\n\n[ SNIPPETBOX CURL TESTS ]"
printf "\n[ these run curl for manual response check on different scenarios ]\n\n"

echo "-------------------"
printf "[ Testing Index... ]\n\n"
curl -i http://localhost:4000

echo ""
echo "-------------------"
printf "[ Testing 404... ]\n\n"
curl -i http://localhost:4000/doesnotexist

echo ""
echo "-------------------"
printf "[ Testing Specific Snippet Index... ]\n\n"
curl -i http://localhost:4000/snippet?id=123

echo ""
echo "-------------------"
printf "[ Testing Specific Nonexistent Snippet Index... ]\n\n"
curl -i http://localhost:4000/snippet
curl -i http://localhost:4000/snippet?id=-1
curl -i http://localhost:4000/snippet?id=xyz

echo ""
echo "-------------------"
printf "[ Testing New Snippet Form... ]\n\n"
curl -i http://localhost:4000/snippet/new

echo ""
echo "-------------------"
printf "[ ALL TESTS COMPLETE ]\n\n"

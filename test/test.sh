#!/usr/bin/env bash
dir=$(dirname "$0")

out="$($dir/example.sh)"
expected_out="$(cat $dir/example.out)"
if [ "$out" = "$expected_out" ]; then
   echo "===== Test success! =====";
else
    echo -e "===== Test failed: =====\n$out"
fi

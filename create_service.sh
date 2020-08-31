#! /bin/sh
eval "cat <<EOF 
$(<uptimed.service)
EOF
" 2> /dev/null


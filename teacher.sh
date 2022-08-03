#!bin/bash

ORAZ=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2)
echo "$ORAZ"
cat ./interviews/interview-"$ORAZ"
echo "$MAIN_SUSPECT"

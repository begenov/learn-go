#! bin/bash 

curl https://01.alem.school/assets/superhero/all.json | jq '.[] | select( .id == 1) '
echo "$opt" | tr -d '"'

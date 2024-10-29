count=$(find | wc -l)
multiply=$(($count * 5))
printf "\t\vTotal files * 5: $multiply\v\n"

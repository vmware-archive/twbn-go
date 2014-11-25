INPUT=$*
STRINGNUM=0

ary=($INPUT)
for key in "${!ary[@]}"
  do
    SHORTTMP[$STRINGNUM]="${SHORTTMP[$STRINGNUM]} ${ary[$key]}"
    LENGTH=$(echo ${#SHORTTMP[$STRINGNUM]})
    if [[ "$LENGTH" -lt "100" ]]; then
      SHORT[$STRINGNUM]=${SHORTTMP[$STRINGNUM]}
    else
      STRINGNUM=$(($STRINGNUM+1))
      SHORTTMP[$STRINGNUM]="${ary[$key]}"
      SHORT[$STRINGNUM]="${ary[$key]}"
    fi
done

for key in "${!SHORT[@]}"
  do
    echo "Playing line: $(($key+1)) of $(($STRINGNUM+1))"
    NEXTURL=$(echo ${SHORT[$key]} | xxd -plain | tr -d '\n' | sed 's/\(..\)/%\1/g')
    mpg123 -q "http://translate.google.com/translate_tts?tl=en&q=$NEXTURL"
done

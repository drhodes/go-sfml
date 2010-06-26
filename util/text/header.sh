

cat /usr/include/SFML/Window/$1.h | grep $2 | grep -v param | grep -v return

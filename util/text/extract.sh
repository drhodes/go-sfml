GPATH=/usr/include/SFML/Audio

for header in $(ls $GPATH)
do
	cat $GPATH/$header | grep CSFML_API | cut -c 11-
done

>cheat2
>cheat3

while read line
do
echo $line > cheat1
python ./protocon.py < cheat1 >> cheat2

done 

grep -v "^$" < cheat2 > cheat3 # fix whitespace.
cat cheat3



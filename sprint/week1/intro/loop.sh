arg=$1
if [ $arg -le 100 ]
then
	for ((i=1;i<=arg;i++))
		do
			echo "$i times I've printed mariiasazhina"
		done
elif [ $arg -gt 100 ]
then
	for ((i=1;i<=100;i++))
		do
    		echo "$i times I've printed mariiasazhina"
		done
fi

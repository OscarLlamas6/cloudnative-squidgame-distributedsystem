while read line; do export "$line";
done < variables.conf
echo "done"
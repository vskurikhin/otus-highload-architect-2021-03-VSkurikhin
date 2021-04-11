#!/bin/sh
i=index.css
cat semantic.min.css > $i
for f in input.css my-div-table.css my-p.css
do
 cat $f >> $i
done

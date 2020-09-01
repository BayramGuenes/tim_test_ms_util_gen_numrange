#!/bin/bash
declare killproc=$(pidof tim_test_ms_util_gen_numrange)
echo killproc=$killproc
kill -9 $killproc
rm ./tim_test_ms_util_gen_numrange
go build
export  SERVICEDB="localhost"
export  PORTDB="33306"
sleep 2&&(echo "************************************************************************"\
&&echo "Access to App via http://localhost:7000/NumRangeServices"\
&&echo "************************************************************************")&
  
./tim_test_ms_util_gen_numrange confLocation=~/ZZDevelop/goProjects/src/tim_presse/timFileSys/settings 



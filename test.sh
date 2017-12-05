test $(curl "http://localhost:8080/" 2> /dev/null) = AMAZON; echo $?

test $(curl "http://localhost:8080/secret/" 2> /dev/null) = FORBIDDEN; echo $?
test $(curl -u amazon:candidate "http://localhost:8080/secret/" 2> /dev/null) = SUCCESS; echo $?

test $(curl 'http://localhost:8080/calc?abc' 2> /dev/null) = "ERROR"; echo $?
test $(curl 'http://localhost:8080/calc?1+1' 2> /dev/null) = "2"; echo $?
test $(curl 'http://localhost:8080/calc?2-1' 2> /dev/null) = "1"; echo $?
test $(curl 'http://localhost:8080/calc?3*2' 2> /dev/null) = "6"; echo $?
test $(curl 'http://localhost:8080/calc?4/2' 2> /dev/null) = "2"; echo $?
test $(curl 'http://localhost:8080/calc?1+2*3' 2> /dev/null) = "7"; echo $?
test $(curl 'http://localhost:8080/calc?(1+2)*3' 2> /dev/null) = "9"; echo $?

curl "http://localhost:8080/stocker?function=deleteall"
curl "http://localhost:8080/stocker?function=addstock&name=xxx&amount=100"
curl "http://localhost:8080/stocker?function=sell&name=xxx&amount=4"
#curl "http://localhost:8080/stocker?function=checkstock&name=xxx" 2> /dev/null
test "$(curl "http://localhost:8080/stocker?function=checkstock&name=xxx" 2> /dev/null)" = "xxx: 96"; echo $?

curl "http://localhost:8080/stocker?function=addstock&name=yyy&amount=100"
curl "http://localhost:8080/stocker?function=addstock&name=YYY&amount=100"
curl "http://localhost:8080/stocker?function=checkstock" 2> /dev/null
test "$(curl "http://localhost:8080/stocker?function=checkstock" 2> /dev/null)" = "xxx: 96 yyy: 100 YYY: 100"; echo $?

curl "http://localhost:8080/stocker?function=deleteall"
#curl "http://localhost:8080/stocker?function=addstock&name=xxx&amount=1.1" 2> /dev/null
test "$(curl "http://localhost:8080/stocker?function=addstock&name=xxx&amount=1.1" 2> /dev/null)" = "ERROR"; echo $?

curl "http://localhost:8080/stocker?function=deleteall"
curl "http://localhost:8080/stocker?function=addstock&name=aaa&amount=10"
curl "http://localhost:8080/stocker?function=addstock&name=bbb&amount=10"
curl "http://localhost:8080/stocker?function=sell&name=aaa&amount=4&price=100"
curl "http://localhost:8080/stocker?function=sell&name=aaa&price=80"
#curl "http://localhost:8080/stocker?function=checkstock&name=aaa" 2> /dev/null
test "$(curl "http://localhost:8080/stocker?function=checkstock&name=aaa" 2> /dev/null)" = "aaa: 5"; echo $?

#curl "http://localhost:8080/stocker?function=checksales" 2> /dev/null
test "$(curl "http://localhost:8080/stocker?function=checksales" 2> /dev/null)" = "sales: 480"; echo $?
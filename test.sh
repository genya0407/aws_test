test $(curl "http://localhost:8080/" 2> /dev/null) = AMAZON; echo $?

test $(curl "http://localhost:8080/secret/" 2> /dev/null) = FORBIDDEN; echo $?
test $(curl -u amazon:candidate "http://localhost:8080/secret/" 2> /dev/null) = SUCCESS; echo $?

curl "http://localhost:8080/stocker?function=deleteall"
curl "http://localhost:8080/stocker?function=addstock&name=xxx&amount=100"
curl "http://localhost:8080/stocker?function=addstock&name=yyy&amount=100"
curl "http://localhost:8080/stocker?function=checkstock&name=xxx"
curl "http://localhost:8080/stocker?function=checkstock"
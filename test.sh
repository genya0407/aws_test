test $(curl "http://localhost:8080/" 2> /dev/null) = AMAZON; echo $?

test $(curl "http://localhost:8080/secret/" 2> /dev/null) = FORBIDDEN; echo $?
test $(curl -u amazon:candidate "http://localhost:8080/secret/" 2> /dev/null) = SUCCESS; echo $?

curl "http://localhost:8080/stocker?function=addstock&name=xxx&amount=100"
# gin run main.go --appPort 5000 --port 3000 --excludeDir ./frontend
echo "$date | Building source code to binary -> `main`"
go build main.go
echo "$date | Build Complete"
echo "$date | Running `main`..."
./main
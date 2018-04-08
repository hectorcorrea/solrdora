echo "Creating Linux version..."
GOOS=linux go build
zip -r solrdora_linux.zip solrdora settings.json public/* views/*


echo "Creating Windows version..."
GOOS=windows GOARCH=386 go build -o solrdora.exe
zip -r solrdora_win.zip solrdora.exe settings.json public/* views/*


echo "Creating Mac OS X version..."
go build
zip -r solrdora_mac.zip solrdora settings.json public/* views/*

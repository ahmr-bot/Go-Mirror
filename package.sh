make
cp -rf ./config.json ./out/config.json
cp -rf ./main.py ./out/main.py
mkdir -p ./out/root
tar -zcvf ./mirrors.tar.gz ./out
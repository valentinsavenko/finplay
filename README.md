
* virtualenv: https://virtualenv.pypa.io/en/stable/
* Virtualenv: https://virtualenv.pypa.io/en/latest/installation/
```bash
# MacOS
brew upgrade python

virtualenv .venv --python=/usr/local/opt/python@3.8/bin/python3
source .venv/bin/activate
pip3 install -r requirements.txt

pip3 install --upgrade pip
pip3 install jupyter
```


```
env GO111MODULE=on go get github.com/gopherdata/gophernotes
mkdir -p ~/Library/Jupyter/kernels/gophernotes
cd ~/Library/Jupyter/kernels/gophernotes
cp "$(go env GOPATH)"/pkg/mod/github.com/gopherdata/gophernotes@v0.7.0/kernel/*  "."
chmod +w ./kernel.json # in case copied kernel.json has no write permission
sed "s|gophernotes|$(go env GOPATH)/bin/gophernotes|" < kernel.json.in > kernel.json
```

`jupyter notebook`

* virtualenv: https://virtualenv.pypa.io/en/stable/
* Virtualenv: https://virtualenv.pypa.io/en/latest/installation/
```bash
# MacOS
brew upgrade python

virtualenv .venv --python=/usr/local/opt/python@3.8/bin/python3
source .venv/bin/activate
pip3 install -r requirements.txt
```
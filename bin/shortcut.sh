echo '''
alias bb="brainbreak () { echo '$@'; unset -f '$(go env GOPATH)'/bin/brainbreak; }; '$(go env GOPATH)'/bin/brainbreak"
''' >> $HOME/.bash_profile
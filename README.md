# kubectl-extract

## Don't use this

I threw this go cli together quickly and then found out that kubectl actually has this functionality built in, 
albeit somewhat hidden. The command is:
```bash
kubectl config view --minify=true --flatten --context foocontext
```

These are the aliases I use now:
```bash
alias kubectl-extract-context='kubectl config view --minify=true --flatten --context'
alias kec='kubectl-extract-context'
```

Works really well with tab completion in zsh. If you type `kec` and `TAB` you will see a list of contexts.

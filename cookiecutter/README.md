# cookiecutter in docker

## how to use

## install
  
- https://cookiecutter.readthedocs.io/en/latest/installation.html

- docker:

```shell
alias cookiecutter="docker run --rm -it -v ~/.ssh/:/home/root/keys/ -v $PWD:/srv/app -w /srv/app odedpriva/cookiecutter"
```

## use

```shell
cookiecutter https://github.com/cookiecutter-flask/cookiecutter-flask.git
```
